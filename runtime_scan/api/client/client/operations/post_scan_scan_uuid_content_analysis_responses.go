// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cisco-open/kubei/runtime_scan/api/client/models"
)

// PostScanScanUUIDContentAnalysisReader is a Reader for the PostScanScanUUIDContentAnalysis structure.
type PostScanScanUUIDContentAnalysisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostScanScanUUIDContentAnalysisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostScanScanUUIDContentAnalysisCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostScanScanUUIDContentAnalysisDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostScanScanUUIDContentAnalysisCreated creates a PostScanScanUUIDContentAnalysisCreated with default headers values
func NewPostScanScanUUIDContentAnalysisCreated() *PostScanScanUUIDContentAnalysisCreated {
	return &PostScanScanUUIDContentAnalysisCreated{}
}

/* PostScanScanUUIDContentAnalysisCreated describes a response with status code 201, with default header values.

Image content analysis successfully reported.
*/
type PostScanScanUUIDContentAnalysisCreated struct {
}

func (o *PostScanScanUUIDContentAnalysisCreated) Error() string {
	return fmt.Sprintf("[POST /scan/{scan-uuid}/contentAnalysis][%d] postScanScanUuidContentAnalysisCreated ", 201)
}

func (o *PostScanScanUUIDContentAnalysisCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostScanScanUUIDContentAnalysisDefault creates a PostScanScanUUIDContentAnalysisDefault with default headers values
func NewPostScanScanUUIDContentAnalysisDefault(code int) *PostScanScanUUIDContentAnalysisDefault {
	return &PostScanScanUUIDContentAnalysisDefault{
		_statusCode: code,
	}
}

/* PostScanScanUUIDContentAnalysisDefault describes a response with status code -1, with default header values.

unknown error
*/
type PostScanScanUUIDContentAnalysisDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the post scan scan UUID content analysis default response
func (o *PostScanScanUUIDContentAnalysisDefault) Code() int {
	return o._statusCode
}

func (o *PostScanScanUUIDContentAnalysisDefault) Error() string {
	return fmt.Sprintf("[POST /scan/{scan-uuid}/contentAnalysis][%d] PostScanScanUUIDContentAnalysis default  %+v", o._statusCode, o.Payload)
}
func (o *PostScanScanUUIDContentAnalysisDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *PostScanScanUUIDContentAnalysisDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
