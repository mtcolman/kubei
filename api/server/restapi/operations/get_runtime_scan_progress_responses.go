// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"wwwin-github.cisco.com/eti/scan-gazr/api/server/models"
)

// GetRuntimeScanProgressOKCode is the HTTP code returned for type GetRuntimeScanProgressOK
const GetRuntimeScanProgressOKCode int = 200

/*GetRuntimeScanProgressOK Success

swagger:response getRuntimeScanProgressOK
*/
type GetRuntimeScanProgressOK struct {

	/*
	  In: Body
	*/
	Payload *models.Progress `json:"body,omitempty"`
}

// NewGetRuntimeScanProgressOK creates GetRuntimeScanProgressOK with default headers values
func NewGetRuntimeScanProgressOK() *GetRuntimeScanProgressOK {

	return &GetRuntimeScanProgressOK{}
}

// WithPayload adds the payload to the get runtime scan progress o k response
func (o *GetRuntimeScanProgressOK) WithPayload(payload *models.Progress) *GetRuntimeScanProgressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime scan progress o k response
func (o *GetRuntimeScanProgressOK) SetPayload(payload *models.Progress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeScanProgressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetRuntimeScanProgressDefault unknown error

swagger:response getRuntimeScanProgressDefault
*/
type GetRuntimeScanProgressDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetRuntimeScanProgressDefault creates GetRuntimeScanProgressDefault with default headers values
func NewGetRuntimeScanProgressDefault(code int) *GetRuntimeScanProgressDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRuntimeScanProgressDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get runtime scan progress default response
func (o *GetRuntimeScanProgressDefault) WithStatusCode(code int) *GetRuntimeScanProgressDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get runtime scan progress default response
func (o *GetRuntimeScanProgressDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get runtime scan progress default response
func (o *GetRuntimeScanProgressDefault) WithPayload(payload *models.APIResponse) *GetRuntimeScanProgressDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get runtime scan progress default response
func (o *GetRuntimeScanProgressDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuntimeScanProgressDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
