// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"wwwin-github.cisco.com/eti/scan-gazr/api/client/models"
)

// GetDashboardMostVulnerableReader is a Reader for the GetDashboardMostVulnerable structure.
type GetDashboardMostVulnerableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardMostVulnerableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardMostVulnerableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDashboardMostVulnerableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDashboardMostVulnerableOK creates a GetDashboardMostVulnerableOK with default headers values
func NewGetDashboardMostVulnerableOK() *GetDashboardMostVulnerableOK {
	return &GetDashboardMostVulnerableOK{}
}

/* GetDashboardMostVulnerableOK describes a response with status code 200, with default header values.

Success
*/
type GetDashboardMostVulnerableOK struct {
	Payload *models.MostVulnerable
}

func (o *GetDashboardMostVulnerableOK) Error() string {
	return fmt.Sprintf("[GET /dashboard/mostVulnerable][%d] getDashboardMostVulnerableOK  %+v", 200, o.Payload)
}
func (o *GetDashboardMostVulnerableOK) GetPayload() *models.MostVulnerable {
	return o.Payload
}

func (o *GetDashboardMostVulnerableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MostVulnerable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardMostVulnerableDefault creates a GetDashboardMostVulnerableDefault with default headers values
func NewGetDashboardMostVulnerableDefault(code int) *GetDashboardMostVulnerableDefault {
	return &GetDashboardMostVulnerableDefault{
		_statusCode: code,
	}
}

/* GetDashboardMostVulnerableDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetDashboardMostVulnerableDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get dashboard most vulnerable default response
func (o *GetDashboardMostVulnerableDefault) Code() int {
	return o._statusCode
}

func (o *GetDashboardMostVulnerableDefault) Error() string {
	return fmt.Sprintf("[GET /dashboard/mostVulnerable][%d] GetDashboardMostVulnerable default  %+v", o._statusCode, o.Payload)
}
func (o *GetDashboardMostVulnerableDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetDashboardMostVulnerableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
