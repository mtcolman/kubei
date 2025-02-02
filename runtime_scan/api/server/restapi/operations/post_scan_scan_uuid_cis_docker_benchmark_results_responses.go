// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cisco-open/kubei/runtime_scan/api/server/models"
)

// PostScanScanUUIDCisDockerBenchmarkResultsCreatedCode is the HTTP code returned for type PostScanScanUUIDCisDockerBenchmarkResultsCreated
const PostScanScanUUIDCisDockerBenchmarkResultsCreatedCode int = 201

/*PostScanScanUUIDCisDockerBenchmarkResultsCreated CIS docker benchmark scan successfully reported.

swagger:response postScanScanUuidCisDockerBenchmarkResultsCreated
*/
type PostScanScanUUIDCisDockerBenchmarkResultsCreated struct {
}

// NewPostScanScanUUIDCisDockerBenchmarkResultsCreated creates PostScanScanUUIDCisDockerBenchmarkResultsCreated with default headers values
func NewPostScanScanUUIDCisDockerBenchmarkResultsCreated() *PostScanScanUUIDCisDockerBenchmarkResultsCreated {

	return &PostScanScanUUIDCisDockerBenchmarkResultsCreated{}
}

// WriteResponse to the client
func (o *PostScanScanUUIDCisDockerBenchmarkResultsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*PostScanScanUUIDCisDockerBenchmarkResultsDefault unknown error

swagger:response postScanScanUuidCisDockerBenchmarkResultsDefault
*/
type PostScanScanUUIDCisDockerBenchmarkResultsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPostScanScanUUIDCisDockerBenchmarkResultsDefault creates PostScanScanUUIDCisDockerBenchmarkResultsDefault with default headers values
func NewPostScanScanUUIDCisDockerBenchmarkResultsDefault(code int) *PostScanScanUUIDCisDockerBenchmarkResultsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostScanScanUUIDCisDockerBenchmarkResultsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post scan scan UUID cis docker benchmark results default response
func (o *PostScanScanUUIDCisDockerBenchmarkResultsDefault) WithStatusCode(code int) *PostScanScanUUIDCisDockerBenchmarkResultsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post scan scan UUID cis docker benchmark results default response
func (o *PostScanScanUUIDCisDockerBenchmarkResultsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post scan scan UUID cis docker benchmark results default response
func (o *PostScanScanUUIDCisDockerBenchmarkResultsDefault) WithPayload(payload *models.APIResponse) *PostScanScanUUIDCisDockerBenchmarkResultsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post scan scan UUID cis docker benchmark results default response
func (o *PostScanScanUUIDCisDockerBenchmarkResultsDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostScanScanUUIDCisDockerBenchmarkResultsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
