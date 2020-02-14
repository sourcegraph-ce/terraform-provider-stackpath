// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteBucketReader is a Reader for the DeleteBucket structure.
type DeleteBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteBucketNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBucketNoContent creates a DeleteBucketNoContent with default headers values
func NewDeleteBucketNoContent() *DeleteBucketNoContent {
	return &DeleteBucketNoContent{}
}

/*DeleteBucketNoContent handles this case with default header values.

No content
*/
type DeleteBucketNoContent struct {
}

func (o *DeleteBucketNoContent) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] deleteBucketNoContent ", 204)
}

func (o *DeleteBucketNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBucketDefault creates a DeleteBucketDefault with default headers values
func NewDeleteBucketDefault(code int) *DeleteBucketDefault {
	return &DeleteBucketDefault{
		_statusCode: code,
	}
}

/*DeleteBucketDefault handles this case with default header values.

Default error structure.
*/
type DeleteBucketDefault struct {
	_statusCode int

	Payload *DeleteBucketDefaultBody
}

// Code gets the status code for the delete bucket default response
func (o *DeleteBucketDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBucketDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/v1/stacks/{stack_id}/buckets/{bucket_id}][%d] DeleteBucket default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteBucketDefault) GetPayload() *DeleteBucketDefaultBody {
	return o.Payload
}

func (o *DeleteBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteBucketDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteBucketDefaultBody delete bucket default body
swagger:model DeleteBucketDefaultBody
*/
type DeleteBucketDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete bucket default body
func (o *DeleteBucketDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteBucketDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteBucketDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteBucketDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
