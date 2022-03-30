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

// GetRuntimeScanProgressReader is a Reader for the GetRuntimeScanProgress structure.
type GetRuntimeScanProgressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeScanProgressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuntimeScanProgressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRuntimeScanProgressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRuntimeScanProgressOK creates a GetRuntimeScanProgressOK with default headers values
func NewGetRuntimeScanProgressOK() *GetRuntimeScanProgressOK {
	return &GetRuntimeScanProgressOK{}
}

/* GetRuntimeScanProgressOK describes a response with status code 200, with default header values.

Success
*/
type GetRuntimeScanProgressOK struct {
	Payload *models.Progress
}

func (o *GetRuntimeScanProgressOK) Error() string {
	return fmt.Sprintf("[GET /runtime/scan/progress][%d] getRuntimeScanProgressOK  %+v", 200, o.Payload)
}
func (o *GetRuntimeScanProgressOK) GetPayload() *models.Progress {
	return o.Payload
}

func (o *GetRuntimeScanProgressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Progress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeScanProgressDefault creates a GetRuntimeScanProgressDefault with default headers values
func NewGetRuntimeScanProgressDefault(code int) *GetRuntimeScanProgressDefault {
	return &GetRuntimeScanProgressDefault{
		_statusCode: code,
	}
}

/* GetRuntimeScanProgressDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetRuntimeScanProgressDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get runtime scan progress default response
func (o *GetRuntimeScanProgressDefault) Code() int {
	return o._statusCode
}

func (o *GetRuntimeScanProgressDefault) Error() string {
	return fmt.Sprintf("[GET /runtime/scan/progress][%d] GetRuntimeScanProgress default  %+v", o._statusCode, o.Payload)
}
func (o *GetRuntimeScanProgressDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetRuntimeScanProgressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
