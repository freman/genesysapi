// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

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

// NewGetSpeechandtextanalyticsTopicsPublishjobParams creates a new GetSpeechandtextanalyticsTopicsPublishjobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSpeechandtextanalyticsTopicsPublishjobParams() *GetSpeechandtextanalyticsTopicsPublishjobParams {
	return &GetSpeechandtextanalyticsTopicsPublishjobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpeechandtextanalyticsTopicsPublishjobParamsWithTimeout creates a new GetSpeechandtextanalyticsTopicsPublishjobParams object
// with the ability to set a timeout on a request.
func NewGetSpeechandtextanalyticsTopicsPublishjobParamsWithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsTopicsPublishjobParams {
	return &GetSpeechandtextanalyticsTopicsPublishjobParams{
		timeout: timeout,
	}
}

// NewGetSpeechandtextanalyticsTopicsPublishjobParamsWithContext creates a new GetSpeechandtextanalyticsTopicsPublishjobParams object
// with the ability to set a context for a request.
func NewGetSpeechandtextanalyticsTopicsPublishjobParamsWithContext(ctx context.Context) *GetSpeechandtextanalyticsTopicsPublishjobParams {
	return &GetSpeechandtextanalyticsTopicsPublishjobParams{
		Context: ctx,
	}
}

// NewGetSpeechandtextanalyticsTopicsPublishjobParamsWithHTTPClient creates a new GetSpeechandtextanalyticsTopicsPublishjobParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSpeechandtextanalyticsTopicsPublishjobParamsWithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsTopicsPublishjobParams {
	return &GetSpeechandtextanalyticsTopicsPublishjobParams{
		HTTPClient: client,
	}
}

/*
GetSpeechandtextanalyticsTopicsPublishjobParams contains all the parameters to send to the API endpoint

	for the get speechandtextanalytics topics publishjob operation.

	Typically these are written to a http.Request.
*/
type GetSpeechandtextanalyticsTopicsPublishjobParams struct {

	/* JobID.

	   The id of the publish topics job
	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get speechandtextanalytics topics publishjob params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpeechandtextanalyticsTopicsPublishjobParams) WithDefaults() *GetSpeechandtextanalyticsTopicsPublishjobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get speechandtextanalytics topics publishjob params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpeechandtextanalyticsTopicsPublishjobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get speechandtextanalytics topics publishjob params
func (o *GetSpeechandtextanalyticsTopicsPublishjobParams) WithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsTopicsPublishjobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get speechandtextanalytics topics publishjob params
func (o *GetSpeechandtextanalyticsTopicsPublishjobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get speechandtextanalytics topics publishjob params
func (o *GetSpeechandtextanalyticsTopicsPublishjobParams) WithContext(ctx context.Context) *GetSpeechandtextanalyticsTopicsPublishjobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get speechandtextanalytics topics publishjob params
func (o *GetSpeechandtextanalyticsTopicsPublishjobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get speechandtextanalytics topics publishjob params
func (o *GetSpeechandtextanalyticsTopicsPublishjobParams) WithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsTopicsPublishjobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get speechandtextanalytics topics publishjob params
func (o *GetSpeechandtextanalyticsTopicsPublishjobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the get speechandtextanalytics topics publishjob params
func (o *GetSpeechandtextanalyticsTopicsPublishjobParams) WithJobID(jobID string) *GetSpeechandtextanalyticsTopicsPublishjobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get speechandtextanalytics topics publishjob params
func (o *GetSpeechandtextanalyticsTopicsPublishjobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpeechandtextanalyticsTopicsPublishjobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
