// Code generated by go-swagger; DO NOT EDIT.

package instance_logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/workload/workload_models"
)

// GetLogsReader is a Reader for the GetLogs structure.
type GetLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLogsOK creates a GetLogsOK with default headers values
func NewGetLogsOK() *GetLogsOK {
	return &GetLogsOK{}
}

/*GetLogsOK handles this case with default header values.

(streaming responses)
*/
type GetLogsOK struct {
	Payload *workload_models.V1LogChunk
}

func (o *GetLogsOK) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/logs][%d] getLogsOK  %+v", 200, o.Payload)
}

func (o *GetLogsOK) GetPayload() *workload_models.V1LogChunk {
	return o.Payload
}

func (o *GetLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.V1LogChunk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsUnauthorized creates a GetLogsUnauthorized with default headers values
func NewGetLogsUnauthorized() *GetLogsUnauthorized {
	return &GetLogsUnauthorized{}
}

/*GetLogsUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type GetLogsUnauthorized struct {
	Payload *workload_models.StackpathapiStatus
}

func (o *GetLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/logs][%d] getLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLogsUnauthorized) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsInternalServerError creates a GetLogsInternalServerError with default headers values
func NewGetLogsInternalServerError() *GetLogsInternalServerError {
	return &GetLogsInternalServerError{}
}

/*GetLogsInternalServerError handles this case with default header values.

Internal server error.
*/
type GetLogsInternalServerError struct {
	Payload *workload_models.StackpathapiStatus
}

func (o *GetLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/logs][%d] getLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLogsInternalServerError) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogsDefault creates a GetLogsDefault with default headers values
func NewGetLogsDefault(code int) *GetLogsDefault {
	return &GetLogsDefault{
		_statusCode: code,
	}
}

/*GetLogsDefault handles this case with default header values.

Default error structure.
*/
type GetLogsDefault struct {
	_statusCode int

	Payload *workload_models.StackpathapiStatus
}

// Code gets the status code for the get logs default response
func (o *GetLogsDefault) Code() int {
	return o._statusCode
}

func (o *GetLogsDefault) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances/{instance_name}/logs][%d] GetLogs default  %+v", o._statusCode, o.Payload)
}

func (o *GetLogsDefault) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
