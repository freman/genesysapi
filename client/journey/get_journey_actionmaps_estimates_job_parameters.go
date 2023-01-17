// Code generated by go-swagger; DO NOT EDIT.

package journey

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

// NewGetJourneyActionmapsEstimatesJobParams creates a new GetJourneyActionmapsEstimatesJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetJourneyActionmapsEstimatesJobParams() *GetJourneyActionmapsEstimatesJobParams {
	return &GetJourneyActionmapsEstimatesJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetJourneyActionmapsEstimatesJobParamsWithTimeout creates a new GetJourneyActionmapsEstimatesJobParams object
// with the ability to set a timeout on a request.
func NewGetJourneyActionmapsEstimatesJobParamsWithTimeout(timeout time.Duration) *GetJourneyActionmapsEstimatesJobParams {
	return &GetJourneyActionmapsEstimatesJobParams{
		timeout: timeout,
	}
}

// NewGetJourneyActionmapsEstimatesJobParamsWithContext creates a new GetJourneyActionmapsEstimatesJobParams object
// with the ability to set a context for a request.
func NewGetJourneyActionmapsEstimatesJobParamsWithContext(ctx context.Context) *GetJourneyActionmapsEstimatesJobParams {
	return &GetJourneyActionmapsEstimatesJobParams{
		Context: ctx,
	}
}

// NewGetJourneyActionmapsEstimatesJobParamsWithHTTPClient creates a new GetJourneyActionmapsEstimatesJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetJourneyActionmapsEstimatesJobParamsWithHTTPClient(client *http.Client) *GetJourneyActionmapsEstimatesJobParams {
	return &GetJourneyActionmapsEstimatesJobParams{
		HTTPClient: client,
	}
}

/*
GetJourneyActionmapsEstimatesJobParams contains all the parameters to send to the API endpoint

	for the get journey actionmaps estimates job operation.

	Typically these are written to a http.Request.
*/
type GetJourneyActionmapsEstimatesJobParams struct {

	/* JobID.

	   ID of the job.
	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get journey actionmaps estimates job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJourneyActionmapsEstimatesJobParams) WithDefaults() *GetJourneyActionmapsEstimatesJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get journey actionmaps estimates job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJourneyActionmapsEstimatesJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get journey actionmaps estimates job params
func (o *GetJourneyActionmapsEstimatesJobParams) WithTimeout(timeout time.Duration) *GetJourneyActionmapsEstimatesJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get journey actionmaps estimates job params
func (o *GetJourneyActionmapsEstimatesJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get journey actionmaps estimates job params
func (o *GetJourneyActionmapsEstimatesJobParams) WithContext(ctx context.Context) *GetJourneyActionmapsEstimatesJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get journey actionmaps estimates job params
func (o *GetJourneyActionmapsEstimatesJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get journey actionmaps estimates job params
func (o *GetJourneyActionmapsEstimatesJobParams) WithHTTPClient(client *http.Client) *GetJourneyActionmapsEstimatesJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get journey actionmaps estimates job params
func (o *GetJourneyActionmapsEstimatesJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the get journey actionmaps estimates job params
func (o *GetJourneyActionmapsEstimatesJobParams) WithJobID(jobID string) *GetJourneyActionmapsEstimatesJobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get journey actionmaps estimates job params
func (o *GetJourneyActionmapsEstimatesJobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetJourneyActionmapsEstimatesJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
