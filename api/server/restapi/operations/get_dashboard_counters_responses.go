// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cisco-open/kubei/api/server/models"
)

// GetDashboardCountersOKCode is the HTTP code returned for type GetDashboardCountersOK
const GetDashboardCountersOKCode int = 200

/*GetDashboardCountersOK Success

swagger:response getDashboardCountersOK
*/
type GetDashboardCountersOK struct {

	/*
	  In: Body
	*/
	Payload *models.DashboardCounters `json:"body,omitempty"`
}

// NewGetDashboardCountersOK creates GetDashboardCountersOK with default headers values
func NewGetDashboardCountersOK() *GetDashboardCountersOK {

	return &GetDashboardCountersOK{}
}

// WithPayload adds the payload to the get dashboard counters o k response
func (o *GetDashboardCountersOK) WithPayload(payload *models.DashboardCounters) *GetDashboardCountersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dashboard counters o k response
func (o *GetDashboardCountersOK) SetPayload(payload *models.DashboardCounters) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDashboardCountersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetDashboardCountersDefault unknown error

swagger:response getDashboardCountersDefault
*/
type GetDashboardCountersDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetDashboardCountersDefault creates GetDashboardCountersDefault with default headers values
func NewGetDashboardCountersDefault(code int) *GetDashboardCountersDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDashboardCountersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get dashboard counters default response
func (o *GetDashboardCountersDefault) WithStatusCode(code int) *GetDashboardCountersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get dashboard counters default response
func (o *GetDashboardCountersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get dashboard counters default response
func (o *GetDashboardCountersDefault) WithPayload(payload *models.APIResponse) *GetDashboardCountersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get dashboard counters default response
func (o *GetDashboardCountersDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDashboardCountersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
