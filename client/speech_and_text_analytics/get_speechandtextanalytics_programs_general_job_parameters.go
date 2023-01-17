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

// NewGetSpeechandtextanalyticsProgramsGeneralJobParams creates a new GetSpeechandtextanalyticsProgramsGeneralJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSpeechandtextanalyticsProgramsGeneralJobParams() *GetSpeechandtextanalyticsProgramsGeneralJobParams {
	return &GetSpeechandtextanalyticsProgramsGeneralJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobParamsWithTimeout creates a new GetSpeechandtextanalyticsProgramsGeneralJobParams object
// with the ability to set a timeout on a request.
func NewGetSpeechandtextanalyticsProgramsGeneralJobParamsWithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsProgramsGeneralJobParams {
	return &GetSpeechandtextanalyticsProgramsGeneralJobParams{
		timeout: timeout,
	}
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobParamsWithContext creates a new GetSpeechandtextanalyticsProgramsGeneralJobParams object
// with the ability to set a context for a request.
func NewGetSpeechandtextanalyticsProgramsGeneralJobParamsWithContext(ctx context.Context) *GetSpeechandtextanalyticsProgramsGeneralJobParams {
	return &GetSpeechandtextanalyticsProgramsGeneralJobParams{
		Context: ctx,
	}
}

// NewGetSpeechandtextanalyticsProgramsGeneralJobParamsWithHTTPClient creates a new GetSpeechandtextanalyticsProgramsGeneralJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSpeechandtextanalyticsProgramsGeneralJobParamsWithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsProgramsGeneralJobParams {
	return &GetSpeechandtextanalyticsProgramsGeneralJobParams{
		HTTPClient: client,
	}
}

/*
GetSpeechandtextanalyticsProgramsGeneralJobParams contains all the parameters to send to the API endpoint

	for the get speechandtextanalytics programs general job operation.

	Typically these are written to a http.Request.
*/
type GetSpeechandtextanalyticsProgramsGeneralJobParams struct {

	/* JobID.

	   The id of the publish programs job
	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get speechandtextanalytics programs general job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpeechandtextanalyticsProgramsGeneralJobParams) WithDefaults() *GetSpeechandtextanalyticsProgramsGeneralJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get speechandtextanalytics programs general job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpeechandtextanalyticsProgramsGeneralJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get speechandtextanalytics programs general job params
func (o *GetSpeechandtextanalyticsProgramsGeneralJobParams) WithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsProgramsGeneralJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get speechandtextanalytics programs general job params
func (o *GetSpeechandtextanalyticsProgramsGeneralJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get speechandtextanalytics programs general job params
func (o *GetSpeechandtextanalyticsProgramsGeneralJobParams) WithContext(ctx context.Context) *GetSpeechandtextanalyticsProgramsGeneralJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get speechandtextanalytics programs general job params
func (o *GetSpeechandtextanalyticsProgramsGeneralJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get speechandtextanalytics programs general job params
func (o *GetSpeechandtextanalyticsProgramsGeneralJobParams) WithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsProgramsGeneralJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get speechandtextanalytics programs general job params
func (o *GetSpeechandtextanalyticsProgramsGeneralJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the get speechandtextanalytics programs general job params
func (o *GetSpeechandtextanalyticsProgramsGeneralJobParams) WithJobID(jobID string) *GetSpeechandtextanalyticsProgramsGeneralJobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get speechandtextanalytics programs general job params
func (o *GetSpeechandtextanalyticsProgramsGeneralJobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpeechandtextanalyticsProgramsGeneralJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
