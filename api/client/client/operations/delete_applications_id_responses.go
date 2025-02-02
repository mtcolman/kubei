// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cisco-open/kubei/api/client/models"
)

// DeleteApplicationsIDReader is a Reader for the DeleteApplicationsID structure.
type DeleteApplicationsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteApplicationsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteApplicationsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteApplicationsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteApplicationsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteApplicationsIDNoContent creates a DeleteApplicationsIDNoContent with default headers values
func NewDeleteApplicationsIDNoContent() *DeleteApplicationsIDNoContent {
	return &DeleteApplicationsIDNoContent{}
}

/* DeleteApplicationsIDNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteApplicationsIDNoContent struct {
}

func (o *DeleteApplicationsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /applications/{id}][%d] deleteApplicationsIdNoContent ", 204)
}

func (o *DeleteApplicationsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationsIDNotFound creates a DeleteApplicationsIDNotFound with default headers values
func NewDeleteApplicationsIDNotFound() *DeleteApplicationsIDNotFound {
	return &DeleteApplicationsIDNotFound{}
}

/* DeleteApplicationsIDNotFound describes a response with status code 404, with default header values.

Application not found.
*/
type DeleteApplicationsIDNotFound struct {
}

func (o *DeleteApplicationsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /applications/{id}][%d] deleteApplicationsIdNotFound ", 404)
}

func (o *DeleteApplicationsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationsIDDefault creates a DeleteApplicationsIDDefault with default headers values
func NewDeleteApplicationsIDDefault(code int) *DeleteApplicationsIDDefault {
	return &DeleteApplicationsIDDefault{
		_statusCode: code,
	}
}

/* DeleteApplicationsIDDefault describes a response with status code -1, with default header values.

unknown error
*/
type DeleteApplicationsIDDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the delete applications ID default response
func (o *DeleteApplicationsIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteApplicationsIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /applications/{id}][%d] DeleteApplicationsID default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteApplicationsIDDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *DeleteApplicationsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
