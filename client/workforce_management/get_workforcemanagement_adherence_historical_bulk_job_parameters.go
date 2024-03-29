// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

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

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobParams creates a new GetWorkforcemanagementAdherenceHistoricalBulkJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobParams() *GetWorkforcemanagementAdherenceHistoricalBulkJobParams {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobParamsWithTimeout creates a new GetWorkforcemanagementAdherenceHistoricalBulkJobParams object
// with the ability to set a timeout on a request.
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementAdherenceHistoricalBulkJobParams {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobParams{
		timeout: timeout,
	}
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobParamsWithContext creates a new GetWorkforcemanagementAdherenceHistoricalBulkJobParams object
// with the ability to set a context for a request.
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobParamsWithContext(ctx context.Context) *GetWorkforcemanagementAdherenceHistoricalBulkJobParams {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobParams{
		Context: ctx,
	}
}

// NewGetWorkforcemanagementAdherenceHistoricalBulkJobParamsWithHTTPClient creates a new GetWorkforcemanagementAdherenceHistoricalBulkJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkforcemanagementAdherenceHistoricalBulkJobParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementAdherenceHistoricalBulkJobParams {
	return &GetWorkforcemanagementAdherenceHistoricalBulkJobParams{
		HTTPClient: client,
	}
}

/*
GetWorkforcemanagementAdherenceHistoricalBulkJobParams contains all the parameters to send to the API endpoint

	for the get workforcemanagement adherence historical bulk job operation.

	Typically these are written to a http.Request.
*/
type GetWorkforcemanagementAdherenceHistoricalBulkJobParams struct {

	/* JobID.

	   ID of the job to get
	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workforcemanagement adherence historical bulk job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobParams) WithDefaults() *GetWorkforcemanagementAdherenceHistoricalBulkJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workforcemanagement adherence historical bulk job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workforcemanagement adherence historical bulk job params
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementAdherenceHistoricalBulkJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement adherence historical bulk job params
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement adherence historical bulk job params
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobParams) WithContext(ctx context.Context) *GetWorkforcemanagementAdherenceHistoricalBulkJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement adherence historical bulk job params
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement adherence historical bulk job params
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementAdherenceHistoricalBulkJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement adherence historical bulk job params
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the get workforcemanagement adherence historical bulk job params
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobParams) WithJobID(jobID string) *GetWorkforcemanagementAdherenceHistoricalBulkJobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get workforcemanagement adherence historical bulk job params
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementAdherenceHistoricalBulkJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
