// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetRuntimeScanResultsParams creates a new GetRuntimeScanResultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRuntimeScanResultsParams() *GetRuntimeScanResultsParams {
	return &GetRuntimeScanResultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRuntimeScanResultsParamsWithTimeout creates a new GetRuntimeScanResultsParams object
// with the ability to set a timeout on a request.
func NewGetRuntimeScanResultsParamsWithTimeout(timeout time.Duration) *GetRuntimeScanResultsParams {
	return &GetRuntimeScanResultsParams{
		timeout: timeout,
	}
}

// NewGetRuntimeScanResultsParamsWithContext creates a new GetRuntimeScanResultsParams object
// with the ability to set a context for a request.
func NewGetRuntimeScanResultsParamsWithContext(ctx context.Context) *GetRuntimeScanResultsParams {
	return &GetRuntimeScanResultsParams{
		Context: ctx,
	}
}

// NewGetRuntimeScanResultsParamsWithHTTPClient creates a new GetRuntimeScanResultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRuntimeScanResultsParamsWithHTTPClient(client *http.Client) *GetRuntimeScanResultsParams {
	return &GetRuntimeScanResultsParams{
		HTTPClient: client,
	}
}

/* GetRuntimeScanResultsParams contains all the parameters to send to the API endpoint
   for the get runtime scan results operation.

   Typically these are written to a http.Request.
*/
type GetRuntimeScanResultsParams struct {

	// VulnerabilitySeverityGte.
	VulnerabilitySeverityGte *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get runtime scan results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuntimeScanResultsParams) WithDefaults() *GetRuntimeScanResultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get runtime scan results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRuntimeScanResultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get runtime scan results params
func (o *GetRuntimeScanResultsParams) WithTimeout(timeout time.Duration) *GetRuntimeScanResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runtime scan results params
func (o *GetRuntimeScanResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runtime scan results params
func (o *GetRuntimeScanResultsParams) WithContext(ctx context.Context) *GetRuntimeScanResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runtime scan results params
func (o *GetRuntimeScanResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runtime scan results params
func (o *GetRuntimeScanResultsParams) WithHTTPClient(client *http.Client) *GetRuntimeScanResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runtime scan results params
func (o *GetRuntimeScanResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVulnerabilitySeverityGte adds the vulnerabilitySeverityGte to the get runtime scan results params
func (o *GetRuntimeScanResultsParams) WithVulnerabilitySeverityGte(vulnerabilitySeverityGte *string) *GetRuntimeScanResultsParams {
	o.SetVulnerabilitySeverityGte(vulnerabilitySeverityGte)
	return o
}

// SetVulnerabilitySeverityGte adds the vulnerabilitySeverityGte to the get runtime scan results params
func (o *GetRuntimeScanResultsParams) SetVulnerabilitySeverityGte(vulnerabilitySeverityGte *string) {
	o.VulnerabilitySeverityGte = vulnerabilitySeverityGte
}

// WriteToRequest writes these params to a swagger request
func (o *GetRuntimeScanResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.VulnerabilitySeverityGte != nil {

		// query param vulnerabilitySeverity[gte]
		var qrVulnerabilitySeverityGte string

		if o.VulnerabilitySeverityGte != nil {
			qrVulnerabilitySeverityGte = *o.VulnerabilitySeverityGte
		}
		qVulnerabilitySeverityGte := qrVulnerabilitySeverityGte
		if qVulnerabilitySeverityGte != "" {

			if err := r.SetQueryParam("vulnerabilitySeverity[gte]", qVulnerabilitySeverityGte); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
