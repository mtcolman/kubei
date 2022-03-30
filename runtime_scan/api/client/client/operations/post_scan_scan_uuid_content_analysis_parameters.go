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

	"wwwin-github.cisco.com/eti/scan-gazr/runtime_scan/api/client/models"
)

// NewPostScanScanUUIDContentAnalysisParams creates a new PostScanScanUUIDContentAnalysisParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostScanScanUUIDContentAnalysisParams() *PostScanScanUUIDContentAnalysisParams {
	return &PostScanScanUUIDContentAnalysisParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostScanScanUUIDContentAnalysisParamsWithTimeout creates a new PostScanScanUUIDContentAnalysisParams object
// with the ability to set a timeout on a request.
func NewPostScanScanUUIDContentAnalysisParamsWithTimeout(timeout time.Duration) *PostScanScanUUIDContentAnalysisParams {
	return &PostScanScanUUIDContentAnalysisParams{
		timeout: timeout,
	}
}

// NewPostScanScanUUIDContentAnalysisParamsWithContext creates a new PostScanScanUUIDContentAnalysisParams object
// with the ability to set a context for a request.
func NewPostScanScanUUIDContentAnalysisParamsWithContext(ctx context.Context) *PostScanScanUUIDContentAnalysisParams {
	return &PostScanScanUUIDContentAnalysisParams{
		Context: ctx,
	}
}

// NewPostScanScanUUIDContentAnalysisParamsWithHTTPClient creates a new PostScanScanUUIDContentAnalysisParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostScanScanUUIDContentAnalysisParamsWithHTTPClient(client *http.Client) *PostScanScanUUIDContentAnalysisParams {
	return &PostScanScanUUIDContentAnalysisParams{
		HTTPClient: client,
	}
}

/* PostScanScanUUIDContentAnalysisParams contains all the parameters to send to the API endpoint
   for the post scan scan UUID content analysis operation.

   Typically these are written to a http.Request.
*/
type PostScanScanUUIDContentAnalysisParams struct {

	// Body.
	Body *models.ImageContentAnalysis

	// ScanUUID.
	//
	// Format: uuid
	ScanUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post scan scan UUID content analysis params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScanScanUUIDContentAnalysisParams) WithDefaults() *PostScanScanUUIDContentAnalysisParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post scan scan UUID content analysis params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScanScanUUIDContentAnalysisParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post scan scan UUID content analysis params
func (o *PostScanScanUUIDContentAnalysisParams) WithTimeout(timeout time.Duration) *PostScanScanUUIDContentAnalysisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post scan scan UUID content analysis params
func (o *PostScanScanUUIDContentAnalysisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post scan scan UUID content analysis params
func (o *PostScanScanUUIDContentAnalysisParams) WithContext(ctx context.Context) *PostScanScanUUIDContentAnalysisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post scan scan UUID content analysis params
func (o *PostScanScanUUIDContentAnalysisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post scan scan UUID content analysis params
func (o *PostScanScanUUIDContentAnalysisParams) WithHTTPClient(client *http.Client) *PostScanScanUUIDContentAnalysisParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post scan scan UUID content analysis params
func (o *PostScanScanUUIDContentAnalysisParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post scan scan UUID content analysis params
func (o *PostScanScanUUIDContentAnalysisParams) WithBody(body *models.ImageContentAnalysis) *PostScanScanUUIDContentAnalysisParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post scan scan UUID content analysis params
func (o *PostScanScanUUIDContentAnalysisParams) SetBody(body *models.ImageContentAnalysis) {
	o.Body = body
}

// WithScanUUID adds the scanUUID to the post scan scan UUID content analysis params
func (o *PostScanScanUUIDContentAnalysisParams) WithScanUUID(scanUUID strfmt.UUID) *PostScanScanUUIDContentAnalysisParams {
	o.SetScanUUID(scanUUID)
	return o
}

// SetScanUUID adds the scanUuid to the post scan scan UUID content analysis params
func (o *PostScanScanUUIDContentAnalysisParams) SetScanUUID(scanUUID strfmt.UUID) {
	o.ScanUUID = scanUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PostScanScanUUIDContentAnalysisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param scan-uuid
	if err := r.SetPathParam("scan-uuid", o.ScanUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
