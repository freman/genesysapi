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

// NewGetAnalyticsUsersDetailsJobsAvailabilityParams creates a new GetAnalyticsUsersDetailsJobsAvailabilityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAnalyticsUsersDetailsJobsAvailabilityParams() *GetAnalyticsUsersDetailsJobsAvailabilityParams {
	return &GetAnalyticsUsersDetailsJobsAvailabilityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityParamsWithTimeout creates a new GetAnalyticsUsersDetailsJobsAvailabilityParams object
// with the ability to set a timeout on a request.
func NewGetAnalyticsUsersDetailsJobsAvailabilityParamsWithTimeout(timeout time.Duration) *GetAnalyticsUsersDetailsJobsAvailabilityParams {
	return &GetAnalyticsUsersDetailsJobsAvailabilityParams{
		timeout: timeout,
	}
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityParamsWithContext creates a new GetAnalyticsUsersDetailsJobsAvailabilityParams object
// with the ability to set a context for a request.
func NewGetAnalyticsUsersDetailsJobsAvailabilityParamsWithContext(ctx context.Context) *GetAnalyticsUsersDetailsJobsAvailabilityParams {
	return &GetAnalyticsUsersDetailsJobsAvailabilityParams{
		Context: ctx,
	}
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityParamsWithHTTPClient creates a new GetAnalyticsUsersDetailsJobsAvailabilityParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAnalyticsUsersDetailsJobsAvailabilityParamsWithHTTPClient(client *http.Client) *GetAnalyticsUsersDetailsJobsAvailabilityParams {
	return &GetAnalyticsUsersDetailsJobsAvailabilityParams{
		HTTPClient: client,
	}
}

/*
GetAnalyticsUsersDetailsJobsAvailabilityParams contains all the parameters to send to the API endpoint

	for the get analytics users details jobs availability operation.

	Typically these are written to a http.Request.
*/
type GetAnalyticsUsersDetailsJobsAvailabilityParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get analytics users details jobs availability params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnalyticsUsersDetailsJobsAvailabilityParams) WithDefaults() *GetAnalyticsUsersDetailsJobsAvailabilityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get analytics users details jobs availability params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnalyticsUsersDetailsJobsAvailabilityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get analytics users details jobs availability params
func (o *GetAnalyticsUsersDetailsJobsAvailabilityParams) WithTimeout(timeout time.Duration) *GetAnalyticsUsersDetailsJobsAvailabilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get analytics users details jobs availability params
func (o *GetAnalyticsUsersDetailsJobsAvailabilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get analytics users details jobs availability params
func (o *GetAnalyticsUsersDetailsJobsAvailabilityParams) WithContext(ctx context.Context) *GetAnalyticsUsersDetailsJobsAvailabilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get analytics users details jobs availability params
func (o *GetAnalyticsUsersDetailsJobsAvailabilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get analytics users details jobs availability params
func (o *GetAnalyticsUsersDetailsJobsAvailabilityParams) WithHTTPClient(client *http.Client) *GetAnalyticsUsersDetailsJobsAvailabilityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get analytics users details jobs availability params
func (o *GetAnalyticsUsersDetailsJobsAvailabilityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnalyticsUsersDetailsJobsAvailabilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
