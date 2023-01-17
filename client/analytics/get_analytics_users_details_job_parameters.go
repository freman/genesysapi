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

// NewGetAnalyticsUsersDetailsJobParams creates a new GetAnalyticsUsersDetailsJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAnalyticsUsersDetailsJobParams() *GetAnalyticsUsersDetailsJobParams {
	return &GetAnalyticsUsersDetailsJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnalyticsUsersDetailsJobParamsWithTimeout creates a new GetAnalyticsUsersDetailsJobParams object
// with the ability to set a timeout on a request.
func NewGetAnalyticsUsersDetailsJobParamsWithTimeout(timeout time.Duration) *GetAnalyticsUsersDetailsJobParams {
	return &GetAnalyticsUsersDetailsJobParams{
		timeout: timeout,
	}
}

// NewGetAnalyticsUsersDetailsJobParamsWithContext creates a new GetAnalyticsUsersDetailsJobParams object
// with the ability to set a context for a request.
func NewGetAnalyticsUsersDetailsJobParamsWithContext(ctx context.Context) *GetAnalyticsUsersDetailsJobParams {
	return &GetAnalyticsUsersDetailsJobParams{
		Context: ctx,
	}
}

// NewGetAnalyticsUsersDetailsJobParamsWithHTTPClient creates a new GetAnalyticsUsersDetailsJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAnalyticsUsersDetailsJobParamsWithHTTPClient(client *http.Client) *GetAnalyticsUsersDetailsJobParams {
	return &GetAnalyticsUsersDetailsJobParams{
		HTTPClient: client,
	}
}

/*
GetAnalyticsUsersDetailsJobParams contains all the parameters to send to the API endpoint

	for the get analytics users details job operation.

	Typically these are written to a http.Request.
*/
type GetAnalyticsUsersDetailsJobParams struct {

	/* JobID.

	   jobId
	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get analytics users details job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnalyticsUsersDetailsJobParams) WithDefaults() *GetAnalyticsUsersDetailsJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get analytics users details job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnalyticsUsersDetailsJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get analytics users details job params
func (o *GetAnalyticsUsersDetailsJobParams) WithTimeout(timeout time.Duration) *GetAnalyticsUsersDetailsJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get analytics users details job params
func (o *GetAnalyticsUsersDetailsJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get analytics users details job params
func (o *GetAnalyticsUsersDetailsJobParams) WithContext(ctx context.Context) *GetAnalyticsUsersDetailsJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get analytics users details job params
func (o *GetAnalyticsUsersDetailsJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get analytics users details job params
func (o *GetAnalyticsUsersDetailsJobParams) WithHTTPClient(client *http.Client) *GetAnalyticsUsersDetailsJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get analytics users details job params
func (o *GetAnalyticsUsersDetailsJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the get analytics users details job params
func (o *GetAnalyticsUsersDetailsJobParams) WithJobID(jobID string) *GetAnalyticsUsersDetailsJobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get analytics users details job params
func (o *GetAnalyticsUsersDetailsJobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnalyticsUsersDetailsJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param jobId
	if err := r.SetPathParam("jobId", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
