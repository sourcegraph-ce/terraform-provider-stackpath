// Code generated by go-swagger; DO NOT EDIT.

package workloads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/workload/workload_models"
)

// GetWorkloadReader is a Reader for the GetWorkload structure.
type GetWorkloadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkloadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWorkloadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkloadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWorkloadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkloadOK creates a GetWorkloadOK with default headers values
func NewGetWorkloadOK() *GetWorkloadOK {
	return &GetWorkloadOK{}
}

/*GetWorkloadOK handles this case with default header values.

GetWorkloadOK get workload o k
*/
type GetWorkloadOK struct {
	Payload *workload_models.V1GetWorkloadResponse
}

func (o *GetWorkloadOK) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}][%d] getWorkloadOK  %+v", 200, o.Payload)
}

func (o *GetWorkloadOK) GetPayload() *workload_models.V1GetWorkloadResponse {
	return o.Payload
}

func (o *GetWorkloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.V1GetWorkloadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadUnauthorized creates a GetWorkloadUnauthorized with default headers values
func NewGetWorkloadUnauthorized() *GetWorkloadUnauthorized {
	return &GetWorkloadUnauthorized{}
}

/*GetWorkloadUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type GetWorkloadUnauthorized struct {
	Payload *workload_models.StackpathapiStatus
}

func (o *GetWorkloadUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}][%d] getWorkloadUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkloadUnauthorized) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetWorkloadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadInternalServerError creates a GetWorkloadInternalServerError with default headers values
func NewGetWorkloadInternalServerError() *GetWorkloadInternalServerError {
	return &GetWorkloadInternalServerError{}
}

/*GetWorkloadInternalServerError handles this case with default header values.

Internal server error.
*/
type GetWorkloadInternalServerError struct {
	Payload *workload_models.StackpathapiStatus
}

func (o *GetWorkloadInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}][%d] getWorkloadInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkloadInternalServerError) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetWorkloadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadDefault creates a GetWorkloadDefault with default headers values
func NewGetWorkloadDefault(code int) *GetWorkloadDefault {
	return &GetWorkloadDefault{
		_statusCode: code,
	}
}

/*GetWorkloadDefault handles this case with default header values.

Default error structure.
*/
type GetWorkloadDefault struct {
	_statusCode int

	Payload *workload_models.StackpathapiStatus
}

// Code gets the status code for the get workload default response
func (o *GetWorkloadDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkloadDefault) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads/{workload_id}][%d] GetWorkload default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkloadDefault) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetWorkloadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
