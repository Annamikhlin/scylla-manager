// Code generated by MockGen. DO NOT EDIT.
// Source: repair.go

// Package mermaidmock is a generated GoMock package.
package mermaidmock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repair "github.com/scylladb/mermaid/repair"
	uuid "github.com/scylladb/mermaid/uuid"
	reflect "reflect"
)

// MockRepairService is a mock of RepairService interface
type MockRepairService struct {
	ctrl     *gomock.Controller
	recorder *MockRepairServiceMockRecorder
}

// MockRepairServiceMockRecorder is the mock recorder for MockRepairService
type MockRepairServiceMockRecorder struct {
	mock *MockRepairService
}

// NewMockRepairService creates a new mock instance
func NewMockRepairService(ctrl *gomock.Controller) *MockRepairService {
	mock := &MockRepairService{ctrl: ctrl}
	mock.recorder = &MockRepairServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepairService) EXPECT() *MockRepairServiceMockRecorder {
	return m.recorder
}

// GetUnit mocks base method
func (m *MockRepairService) GetUnit(ctx context.Context, clusterID, ID uuid.UUID) (*repair.Unit, error) {
	ret := m.ctrl.Call(m, "GetUnit", ctx, clusterID, ID)
	ret0, _ := ret[0].(*repair.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnit indicates an expected call of GetUnit
func (mr *MockRepairServiceMockRecorder) GetUnit(ctx, clusterID, ID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnit", reflect.TypeOf((*MockRepairService)(nil).GetUnit), ctx, clusterID, ID)
}

// PutUnit mocks base method
func (m *MockRepairService) PutUnit(ctx context.Context, u *repair.Unit) error {
	ret := m.ctrl.Call(m, "PutUnit", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutUnit indicates an expected call of PutUnit
func (mr *MockRepairServiceMockRecorder) PutUnit(ctx, u interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUnit", reflect.TypeOf((*MockRepairService)(nil).PutUnit), ctx, u)
}

// DeleteUnit mocks base method
func (m *MockRepairService) DeleteUnit(ctx context.Context, clusterID, ID uuid.UUID) error {
	ret := m.ctrl.Call(m, "DeleteUnit", ctx, clusterID, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUnit indicates an expected call of DeleteUnit
func (mr *MockRepairServiceMockRecorder) DeleteUnit(ctx, clusterID, ID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUnit", reflect.TypeOf((*MockRepairService)(nil).DeleteUnit), ctx, clusterID, ID)
}

// ListUnits mocks base method
func (m *MockRepairService) ListUnits(ctx context.Context, clusterID uuid.UUID) ([]*repair.Unit, error) {
	ret := m.ctrl.Call(m, "ListUnits", ctx, clusterID)
	ret0, _ := ret[0].([]*repair.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUnits indicates an expected call of ListUnits
func (mr *MockRepairServiceMockRecorder) ListUnits(ctx, clusterID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUnits", reflect.TypeOf((*MockRepairService)(nil).ListUnits), ctx, clusterID)
}

// GetConfig mocks base method
func (m *MockRepairService) GetConfig(ctx context.Context, src repair.ConfigSource) (*repair.Config, error) {
	ret := m.ctrl.Call(m, "GetConfig", ctx, src)
	ret0, _ := ret[0].(*repair.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockRepairServiceMockRecorder) GetConfig(ctx, src interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockRepairService)(nil).GetConfig), ctx, src)
}

// PutConfig mocks base method
func (m *MockRepairService) PutConfig(ctx context.Context, src repair.ConfigSource, c *repair.Config) error {
	ret := m.ctrl.Call(m, "PutConfig", ctx, src, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutConfig indicates an expected call of PutConfig
func (mr *MockRepairServiceMockRecorder) PutConfig(ctx, src, c interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutConfig", reflect.TypeOf((*MockRepairService)(nil).PutConfig), ctx, src, c)
}

// DeleteConfig mocks base method
func (m *MockRepairService) DeleteConfig(ctx context.Context, src repair.ConfigSource) error {
	ret := m.ctrl.Call(m, "DeleteConfig", ctx, src)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConfig indicates an expected call of DeleteConfig
func (mr *MockRepairServiceMockRecorder) DeleteConfig(ctx, src interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConfig", reflect.TypeOf((*MockRepairService)(nil).DeleteConfig), ctx, src)
}

// Repair mocks base method
func (m *MockRepairService) Repair(ctx context.Context, u *repair.Unit, taskID uuid.UUID) error {
	ret := m.ctrl.Call(m, "Repair", ctx, u, taskID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Repair indicates an expected call of Repair
func (mr *MockRepairServiceMockRecorder) Repair(ctx, u, taskID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Repair", reflect.TypeOf((*MockRepairService)(nil).Repair), ctx, u, taskID)
}
