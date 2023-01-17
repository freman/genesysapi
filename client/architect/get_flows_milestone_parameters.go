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
)

// NewGetFlowsMilestoneParams creates a new GetFlowsMilestoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFlowsMilestoneParams() *GetFlowsMilestoneParams {
	return &GetFlowsMilestoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowsMilestoneParamsWithTimeout creates a new GetFlowsMilestoneParams object
// with the ability to set a timeout on a request.
func NewGetFlowsMilestoneParamsWithTimeout(timeout time.Duration) *GetFlowsMilestoneParams {
	return &GetFlowsMilestoneParams{
		timeout: timeout,
	}
}

// NewGetFlowsMilestoneParamsWithContext creates a new GetFlowsMilestoneParams object
// with the ability to set a context for a request.
func NewGetFlowsMilestoneParamsWithContext(ctx context.Context) *GetFlowsMilestoneParams {
	return &GetFlowsMilestoneParams{
		Context: ctx,
	}
}

// NewGetFlowsMilestoneParamsWithHTTPClient creates a new GetFlowsMilestoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFlowsMilestoneParamsWithHTTPClient(client *http.Client) *GetFlowsMilestoneParams {
	return &GetFlowsMilestoneParams{
		HTTPClient: client,
	}
}

/*
GetFlowsMilestoneParams contains all the parameters to send to the API endpoint

	for the get flows milestone operation.

	Typically these are written to a http.Request.
*/
type GetFlowsMilestoneParams struct {

	/* MilestoneID.

	   flow milestone ID
	*/
	MilestoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get flows milestone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowsMilestoneParams) WithDefaults() *GetFlowsMilestoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get flows milestone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowsMilestoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get flows milestone params
func (o *GetFlowsMilestoneParams) WithTimeout(timeout time.Duration) *GetFlowsMilestoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flows milestone params
func (o *GetFlowsMilestoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flows milestone params
func (o *GetFlowsMilestoneParams) WithContext(ctx context.Context) *GetFlowsMilestoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flows milestone params
func (o *GetFlowsMilestoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flows milestone params
func (o *GetFlowsMilestoneParams) WithHTTPClient(client *http.Client) *GetFlowsMilestoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flows milestone params
func (o *GetFlowsMilestoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMilestoneID adds the milestoneID to the get flows milestone params
func (o *GetFlowsMilestoneParams) WithMilestoneID(milestoneID string) *GetFlowsMilestoneParams {
	o.SetMilestoneID(milestoneID)
	return o
}

// SetMilestoneID adds the milestoneId to the get flows milestone params
func (o *GetFlowsMilestoneParams) SetMilestoneID(milestoneID string) {
	o.MilestoneID = milestoneID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowsMilestoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param milestoneId
	if err := r.SetPathParam("milestoneId", o.MilestoneID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
