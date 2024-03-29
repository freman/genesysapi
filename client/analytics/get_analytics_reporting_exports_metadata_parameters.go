// Code generated by go-swagger; DO NOT EDIT.

package analytics

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

// NewGetAnalyticsReportingExportsMetadataParams creates a new GetAnalyticsReportingExportsMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAnalyticsReportingExportsMetadataParams() *GetAnalyticsReportingExportsMetadataParams {
	return &GetAnalyticsReportingExportsMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnalyticsReportingExportsMetadataParamsWithTimeout creates a new GetAnalyticsReportingExportsMetadataParams object
// with the ability to set a timeout on a request.
func NewGetAnalyticsReportingExportsMetadataParamsWithTimeout(timeout time.Duration) *GetAnalyticsReportingExportsMetadataParams {
	return &GetAnalyticsReportingExportsMetadataParams{
		timeout: timeout,
	}
}

// NewGetAnalyticsReportingExportsMetadataParamsWithContext creates a new GetAnalyticsReportingExportsMetadataParams object
// with the ability to set a context for a request.
func NewGetAnalyticsReportingExportsMetadataParamsWithContext(ctx context.Context) *GetAnalyticsReportingExportsMetadataParams {
	return &GetAnalyticsReportingExportsMetadataParams{
		Context: ctx,
	}
}

// NewGetAnalyticsReportingExportsMetadataParamsWithHTTPClient creates a new GetAnalyticsReportingExportsMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAnalyticsReportingExportsMetadataParamsWithHTTPClient(client *http.Client) *GetAnalyticsReportingExportsMetadataParams {
	return &GetAnalyticsReportingExportsMetadataParams{
		HTTPClient: client,
	}
}

/*
GetAnalyticsReportingExportsMetadataParams contains all the parameters to send to the API endpoint

	for the get analytics reporting exports metadata operation.

	Typically these are written to a http.Request.
*/
type GetAnalyticsReportingExportsMetadataParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get analytics reporting exports metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnalyticsReportingExportsMetadataParams) WithDefaults() *GetAnalyticsReportingExportsMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get analytics reporting exports metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnalyticsReportingExportsMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get analytics reporting exports metadata params
func (o *GetAnalyticsReportingExportsMetadataParams) WithTimeout(timeout time.Duration) *GetAnalyticsReportingExportsMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get analytics reporting exports metadata params
func (o *GetAnalyticsReportingExportsMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get analytics reporting exports metadata params
func (o *GetAnalyticsReportingExportsMetadataParams) WithContext(ctx context.Context) *GetAnalyticsReportingExportsMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get analytics reporting exports metadata params
func (o *GetAnalyticsReportingExportsMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get analytics reporting exports metadata params
func (o *GetAnalyticsReportingExportsMetadataParams) WithHTTPClient(client *http.Client) *GetAnalyticsReportingExportsMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get analytics reporting exports metadata params
func (o *GetAnalyticsReportingExportsMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnalyticsReportingExportsMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
