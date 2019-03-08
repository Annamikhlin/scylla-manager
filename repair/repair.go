// Copyright (C) 2017 ScyllaDB

package repair

import (
	"encoding/binary"
	"math"
	"sort"
	"strings"

	"github.com/cespare/xxhash"
	"github.com/pkg/errors"
	"github.com/scylladb/go-set/strset"
	"github.com/scylladb/mermaid"
	"github.com/scylladb/mermaid/internal/dht"
	"github.com/scylladb/mermaid/internal/inexlist"
	"github.com/scylladb/mermaid/scyllaclient"
	"github.com/scylladb/mermaid/uuid"
	"go.uber.org/multierr"
)

// validateHostsBelongToCluster checks that the hosts belong to the cluster.
func validateHostsBelongToCluster(dcMap map[string][]string, hosts ...string) error {
	if len(hosts) == 0 {
		return nil
	}

	all := strset.New()
	for _, dcHosts := range dcMap {
		for _, h := range dcHosts {
			all.Add(h)
		}
	}

	var missing []string
	for _, h := range hosts {
		if !all.Has(h) {
			missing = append(missing, h)
		}
	}
	if len(missing) > 0 {
		return errors.Errorf("no such hosts %s", strings.Join(missing, ", "))
	}
	return nil
}

// groupSegmentsByHost extracts a list of segments (token ranges) for hosts
// in a datacenter ds and returns a mapping from host to list of its segments.
// If host is not empty the mapping contains only token ranges for that host.
// If withHosts is not empty the mapping contains only token ranges that are
// replicated by at least one of the hosts.
func groupSegmentsByHost(dc string, host string, withHosts []string, tr TokenRangesKind, ring scyllaclient.Ring) map[string]segments {
	var (
		hostSegments  = make(map[string]segments)
		replicaFilter = strset.New(withHosts...)
	)

	for _, t := range ring.Tokens {
		// ignore segments that are not replicated by any of the withHosts
		if !replicaFilter.IsEmpty() {
			ok := false
			for _, h := range t.Replicas {
				if replicaFilter.Has(h) {
					ok = true
					break
				}
			}
			if !ok {
				continue
			}
		}

		// select replicas from dc based on token kind
		hosts := strset.New()
		switch tr {
		case DCPrimaryTokenRanges:
			for _, h := range t.Replicas {
				if ring.HostDC[h] == dc {
					hosts.Add(h)
					break
				}
			}
		case PrimaryTokenRanges:
			h := t.Replicas[0]
			if ring.HostDC[h] == dc {
				hosts.Add(h)
			}
		case NonPrimaryTokenRanges:
			for _, h := range t.Replicas[1:] {
				if ring.HostDC[h] == dc {
					hosts.Add(h)
				}
			}
		case AllTokenRanges:
			for _, h := range t.Replicas {
				if ring.HostDC[h] == dc {
					hosts.Add(h)
				}
			}
		default:
			panic("no token ranges specified") // this should never happen...
		}

		// filter replicas by host (if needed)
		if host != "" {
			if hosts.Has(host) {
				hosts = strset.New(host)
			} else {
				continue
			}
		}

		// create and add segments for every host
		hosts.Each(func(h string) bool {
			if t.StartToken > t.EndToken {
				hostSegments[h] = append(hostSegments[h],
					&segment{StartToken: dht.Murmur3MinToken, EndToken: t.EndToken},
					&segment{StartToken: t.StartToken, EndToken: dht.Murmur3MaxToken},
				)
			} else {
				hostSegments[h] = append(hostSegments[h], &segment{StartToken: t.StartToken, EndToken: t.EndToken})
			}
			return true
		})
	}

	// remove segments for withHosts as they should not be coordinator hosts
	if !replicaFilter.IsEmpty() {
		replicaFilter.Each(func(h string) bool {
			delete(hostSegments, h)
			return true
		})
	}

	return hostSegments
}

// validateShardProgress checks if run progress, possibly copied from a
// different run matches the shards.
func validateShardProgress(shards []segments, prog []*RunProgress) error {
	if len(prog) != len(shards) {
		return errors.New("length mismatch")
	}

	for i, p := range prog {
		if p.Shard != i {
			return errors.Errorf("shard %d: progress for shard %d", i, p.Shard)
		}
		if p.SegmentCount != len(shards[i]) {
			return errors.Errorf("shard %d: segment count mismatch got %d expected %d", p.Shard, p.SegmentCount, len(shards[i]))
		}
		if p.LastStartToken != 0 {
			if _, ok := shards[i].containStartToken(p.LastStartToken); !ok {
				return errors.Errorf("shard %d: no segment for start token %d", p.Shard, p.LastStartToken)
			}
		}
		for _, token := range p.SegmentErrorStartTokens {
			if _, ok := shards[i].containStartToken(token); !ok {
				return errors.Errorf("shard %d: no segment for (failed) start token %d", p.Shard, token)
			}
		}
	}

	return nil
}

// topologyHash returns hash of all the tokens.
func topologyHash(tokens []int64) uuid.UUID {
	var (
		xx = xxhash.New()
		b  = make([]byte, 8)
		u  uint64
	)
	for _, t := range tokens {
		if t >= 0 {
			u = uint64(t)
		} else {
			u = uint64(math.MaxInt64 + t)
		}
		binary.LittleEndian.PutUint64(b, u)
		xx.Write(b) // nolint
	}
	h := xx.Sum64()

	return uuid.NewFromUint64(h>>32, uint64(uint32(h)))
}

func aggregateProgress(run *Run, prog []*RunProgress) Progress {
	if len(run.Units) == 0 {
		return Progress{}
	}

	v := Progress{
		DC:          run.DC,
		TokenRanges: run.TokenRanges,
	}

	idx := 0
	for i, u := range run.Units {
		end := sort.Search(len(prog), func(j int) bool {
			return prog[j].Unit > i
		})
		p := aggregateUnitProgress(u, prog[idx:end])
		v.progress.addProgress(p.progress)
		v.Units = append(v.Units, p)
		idx = end
	}

	v.progress.calculateProgress()
	return v
}

func aggregateUnitProgress(u Unit, prog []*RunProgress) UnitProgress {
	v := UnitProgress{Unit: u}

	if len(prog) == 0 {
		return v
	}

	var (
		host   = prog[0].Host
		nprog  progress
		shards []ShardProgress
	)
	for _, p := range prog {
		if p.Host != host {
			v.Nodes = append(v.Nodes, NodeProgress{
				progress: nprog.calculateProgress(),
				Host:     host,
				Shards:   shards,
			})
			host = p.Host
			nprog = progress{}
			shards = nil
		}

		sprog := progress{
			segmentCount:   p.SegmentCount,
			segmentSuccess: p.SegmentSuccess,
			segmentError:   p.SegmentError,
		}
		nprog.addProgress(sprog)
		v.progress.addProgress(sprog)

		shards = append(shards, ShardProgress{
			progress:       sprog.calculateProgress(),
			SegmentCount:   p.SegmentCount,
			SegmentSuccess: p.SegmentSuccess,
			SegmentError:   p.SegmentError,
		})
	}
	v.Nodes = append(v.Nodes, NodeProgress{
		progress: nprog.calculateProgress(),
		Host:     host,
		Shards:   shards,
	})

	sort.Slice(v.Nodes, func(i, j int) bool {
		return v.Nodes[i].PercentComplete > v.Nodes[j].PercentComplete
	})
	v.progress.calculateProgress()
	return v
}

func sortUnits(units []Unit, inclExcl inexlist.InExList) {
	positions := make(map[string]int)
	for _, u := range units {
		min := inclExcl.Size()
		for _, t := range u.Tables {
			if p := inclExcl.FirstMatch(u.Keyspace + "." + t); p >= 0 && p < min {
				min = p
			}
		}
		positions[u.Keyspace] = min
	}

	sort.Slice(units, func(i, j int) bool {
		// order by position
		if positions[units[i].Keyspace] < positions[units[j].Keyspace] {
			return true
		} else if positions[units[i].Keyspace] > positions[units[j].Keyspace] {
			return false
		} else {
			// promote system keyspaces
			l := strings.HasPrefix(units[i].Keyspace, "system")
			r := strings.HasPrefix(units[j].Keyspace, "system")
			if l && !r {
				return true
			} else if !l && r {
				return false
			} else {
				// order by name
				return units[i].Keyspace < units[j].Keyspace
			}
		}
	})
}

func validateKeyspaceFilters(filters []string) error {
	var errs error
	for i, f := range filters {
		err := validateKeyspaceFilter(filters[i])
		if err != nil {
			errs = multierr.Append(errs, errors.Wrapf(err, "%q on position %d", f, i))
			continue
		}
	}
	return mermaid.ErrValidate(errs, "invalid filters")
}

func validateKeyspaceFilter(filter string) error {
	if filter == "*" || filter == "!*" {
		return nil
	}
	if strings.HasPrefix(filter, ".") {
		return errors.New("missing keyspace")
	}
	return nil
}

func decorateKeyspaceFilters(filters []string) []string {
	if len(filters) == 0 {
		filters = append(filters, "*.*")
	}

	for i, f := range filters {
		if strings.Contains(f, ".") {
			continue
		}
		if strings.HasSuffix(f, "*") {
			filters[i] = strings.TrimSuffix(f, "*") + "*.*"
		} else {
			filters[i] += ".*"
		}
	}

	return filters
}

func decorateDCFilters(filters []string) []string {
	if len(filters) == 0 {
		filters = append(filters, "*")
	}
	return filters
}
