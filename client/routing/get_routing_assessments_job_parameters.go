// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewGetRoutingAssessmentsJobParams creates a new GetRoutingAssessmentsJobParams object
// with the default values initialized.
func NewGetRoutingAssessmentsJobParams() *GetRoutingAssessmentsJobParams {
	var ()
	return &GetRoutingAssessmentsJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingAssessmentsJobParamsWithTimeout creates a new GetRoutingAssessmentsJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoutingAssessmentsJobParamsWithTimeout(timeout time.Duration) *GetRoutingAssessmentsJobParams {
	var ()
	return &GetRoutingAssessmentsJobParams{

		timeout: timeout,
	}
}

// NewGetRoutingAssessmentsJobParamsWithContext creates a new GetRoutingAssessmentsJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoutingAssessmentsJobParamsWithContext(ctx context.Context) *GetRoutingAssessmentsJobParams {
	var ()
	return &GetRoutingAssessmentsJobParams{

		Context: ctx,
	}
}

// NewGetRoutingAssessmentsJobParamsWithHTTPClient creates a new GetRoutingAssessmentsJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoutingAssessmentsJobParamsWithHTTPClient(client *http.Client) *GetRoutingAssessmentsJobParams {
	var ()
	return &GetRoutingAssessmentsJobParams{
		HTTPClient: client,
	}
}

/*GetRoutingAssessmentsJobParams contains all the parameters to send to the API endpoint
for the get routing assessments job operation typically these are written to a http.Request
*/
type GetRoutingAssessmentsJobParams struct {

	/*JobID
	  Benefit Assessment Job ID

	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routing assessments job params
func (o *GetRoutingAssessmentsJobParams) WithTimeout(timeout time.Duration) *GetRoutingAssessmentsJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing assessments job params
func (o *GetRoutingAssessmentsJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing assessments job params
func (o *GetRoutingAssessmentsJobParams) WithContext(ctx context.Context) *GetRoutingAssessmentsJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing assessments job params
func (o *GetRoutingAssessmentsJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing assessments job params
func (o *GetRoutingAssessmentsJobParams) WithHTTPClient(client *http.Client) *GetRoutingAssessmentsJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing assessments job params
func (o *GetRoutingAssessmentsJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the get routing assessments job params
func (o *GetRoutingAssessmentsJobParams) WithJobID(jobID string) *GetRoutingAssessmentsJobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get routing assessments job params
func (o *GetRoutingAssessmentsJobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingAssessmentsJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
