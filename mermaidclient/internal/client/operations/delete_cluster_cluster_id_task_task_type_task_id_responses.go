// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteClusterClusterIDTaskTaskTypeTaskIDReader is a Reader for the DeleteClusterClusterIDTaskTaskTypeTaskID structure.
type DeleteClusterClusterIDTaskTaskTypeTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClusterClusterIDTaskTaskTypeTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		body := response.Body()
		defer body.Close()

		var m json.RawMessage
		if err := json.NewDecoder(body).Decode(&m); err != nil {
			return nil, err
		}

		return nil, runtime.NewAPIError("API error", m, response.Code())
	}
}

// NewDeleteClusterClusterIDTaskTaskTypeTaskIDOK creates a DeleteClusterClusterIDTaskTaskTypeTaskIDOK with default headers values
func NewDeleteClusterClusterIDTaskTaskTypeTaskIDOK() *DeleteClusterClusterIDTaskTaskTypeTaskIDOK {
	return &DeleteClusterClusterIDTaskTaskTypeTaskIDOK{}
}

/*DeleteClusterClusterIDTaskTaskTypeTaskIDOK handles this case with default header values.

OK
*/
type DeleteClusterClusterIDTaskTaskTypeTaskIDOK struct {
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDOK) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}/task/{task_type}/{task_id}][%d] deleteClusterClusterIdTaskTaskTypeTaskIdOK ", 200)
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
