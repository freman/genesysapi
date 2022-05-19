// Code generated by go-swagger; DO NOT EDIT.

package architect

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
	"github.com/go-openapi/swag"
)

// NewGetFlowsJobParams creates a new GetFlowsJobParams object
// with the default values initialized.
func NewGetFlowsJobParams() *GetFlowsJobParams {
	var ()
	return &GetFlowsJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowsJobParamsWithTimeout creates a new GetFlowsJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFlowsJobParamsWithTimeout(timeout time.Duration) *GetFlowsJobParams {
	var ()
	return &GetFlowsJobParams{

		timeout: timeout,
	}
}

// NewGetFlowsJobParamsWithContext creates a new GetFlowsJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFlowsJobParamsWithContext(ctx context.Context) *GetFlowsJobParams {
	var ()
	return &GetFlowsJobParams{

		Context: ctx,
	}
}

// NewGetFlowsJobParamsWithHTTPClient creates a new GetFlowsJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFlowsJobParamsWithHTTPClient(client *http.Client) *GetFlowsJobParams {
	var ()
	return &GetFlowsJobParams{
		HTTPClient: client,
	}
}

/*GetFlowsJobParams contains all the parameters to send to the API endpoint
for the get flows job operation typically these are written to a http.Request
*/
type GetFlowsJobParams struct {

	/*Expand
	  Which fields, if any, to expand.

	*/
	Expand []string
	/*JobID
	  Job ID

	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get flows job params
func (o *GetFlowsJobParams) WithTimeout(timeout time.Duration) *GetFlowsJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flows job params
func (o *GetFlowsJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flows job params
func (o *GetFlowsJobParams) WithContext(ctx context.Context) *GetFlowsJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flows job params
func (o *GetFlowsJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flows job params
func (o *GetFlowsJobParams) WithHTTPClient(client *http.Client) *GetFlowsJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flows job params
func (o *GetFlowsJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get flows job params
func (o *GetFlowsJobParams) WithExpand(expand []string) *GetFlowsJobParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get flows job params
func (o *GetFlowsJobParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithJobID adds the jobID to the get flows job params
func (o *GetFlowsJobParams) WithJobID(jobID string) *GetFlowsJobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get flows job params
func (o *GetFlowsJobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowsJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	// path param jobId
	if err := r.SetPathParam("jobId", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}