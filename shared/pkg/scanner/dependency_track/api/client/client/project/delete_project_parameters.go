// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteProjectParams creates a new DeleteProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProjectParams() *DeleteProjectParams {
	return &DeleteProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectParamsWithTimeout creates a new DeleteProjectParams object
// with the ability to set a timeout on a request.
func NewDeleteProjectParamsWithTimeout(timeout time.Duration) *DeleteProjectParams {
	return &DeleteProjectParams{
		timeout: timeout,
	}
}

// NewDeleteProjectParamsWithContext creates a new DeleteProjectParams object
// with the ability to set a context for a request.
func NewDeleteProjectParamsWithContext(ctx context.Context) *DeleteProjectParams {
	return &DeleteProjectParams{
		Context: ctx,
	}
}

// NewDeleteProjectParamsWithHTTPClient creates a new DeleteProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProjectParamsWithHTTPClient(client *http.Client) *DeleteProjectParams {
	return &DeleteProjectParams{
		HTTPClient: client,
	}
}

/* DeleteProjectParams contains all the parameters to send to the API endpoint
   for the delete project operation.

   Typically these are written to a http.Request.
*/
type DeleteProjectParams struct {

	/* UUID.

	   The UUID of the project to delete
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectParams) WithDefaults() *DeleteProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete project params
func (o *DeleteProjectParams) WithTimeout(timeout time.Duration) *DeleteProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project params
func (o *DeleteProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project params
func (o *DeleteProjectParams) WithContext(ctx context.Context) *DeleteProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project params
func (o *DeleteProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project params
func (o *DeleteProjectParams) WithHTTPClient(client *http.Client) *DeleteProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project params
func (o *DeleteProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the delete project params
func (o *DeleteProjectParams) WithUUID(uuid string) *DeleteProjectParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete project params
func (o *DeleteProjectParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
