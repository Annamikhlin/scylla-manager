// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v2/models"
)

// FindConfigMemoryAllocatorReader is a Reader for the FindConfigMemoryAllocator structure.
type FindConfigMemoryAllocatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigMemoryAllocatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigMemoryAllocatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigMemoryAllocatorDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigMemoryAllocatorOK creates a FindConfigMemoryAllocatorOK with default headers values
func NewFindConfigMemoryAllocatorOK() *FindConfigMemoryAllocatorOK {
	return &FindConfigMemoryAllocatorOK{}
}

/*FindConfigMemoryAllocatorOK handles this case with default header values.

Config value
*/
type FindConfigMemoryAllocatorOK struct {
	Payload string
}

func (o *FindConfigMemoryAllocatorOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigMemoryAllocatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigMemoryAllocatorDefault creates a FindConfigMemoryAllocatorDefault with default headers values
func NewFindConfigMemoryAllocatorDefault(code int) *FindConfigMemoryAllocatorDefault {
	return &FindConfigMemoryAllocatorDefault{
		_statusCode: code,
	}
}

/*FindConfigMemoryAllocatorDefault handles this case with default header values.

unexpected error
*/
type FindConfigMemoryAllocatorDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config memory allocator default response
func (o *FindConfigMemoryAllocatorDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigMemoryAllocatorDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigMemoryAllocatorDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigMemoryAllocatorDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}