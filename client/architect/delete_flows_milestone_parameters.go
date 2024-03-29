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

// NewDeleteFlowsMilestoneParams creates a new DeleteFlowsMilestoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteFlowsMilestoneParams() *DeleteFlowsMilestoneParams {
	return &DeleteFlowsMilestoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFlowsMilestoneParamsWithTimeout creates a new DeleteFlowsMilestoneParams object
// with the ability to set a timeout on a request.
func NewDeleteFlowsMilestoneParamsWithTimeout(timeout time.Duration) *DeleteFlowsMilestoneParams {
	return &DeleteFlowsMilestoneParams{
		timeout: timeout,
	}
}

// NewDeleteFlowsMilestoneParamsWithContext creates a new DeleteFlowsMilestoneParams object
// with the ability to set a context for a request.
func NewDeleteFlowsMilestoneParamsWithContext(ctx context.Context) *DeleteFlowsMilestoneParams {
	return &DeleteFlowsMilestoneParams{
		Context: ctx,
	}
}

// NewDeleteFlowsMilestoneParamsWithHTTPClient creates a new DeleteFlowsMilestoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteFlowsMilestoneParamsWithHTTPClient(client *http.Client) *DeleteFlowsMilestoneParams {
	return &DeleteFlowsMilestoneParams{
		HTTPClient: client,
	}
}

/*
DeleteFlowsMilestoneParams contains all the parameters to send to the API endpoint

	for the delete flows milestone operation.

	Typically these are written to a http.Request.
*/
type DeleteFlowsMilestoneParams struct {

	/* MilestoneID.

	   flow milestone ID
	*/
	MilestoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete flows milestone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFlowsMilestoneParams) WithDefaults() *DeleteFlowsMilestoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete flows milestone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFlowsMilestoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete flows milestone params
func (o *DeleteFlowsMilestoneParams) WithTimeout(timeout time.Duration) *DeleteFlowsMilestoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete flows milestone params
func (o *DeleteFlowsMilestoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete flows milestone params
func (o *DeleteFlowsMilestoneParams) WithContext(ctx context.Context) *DeleteFlowsMilestoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete flows milestone params
func (o *DeleteFlowsMilestoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete flows milestone params
func (o *DeleteFlowsMilestoneParams) WithHTTPClient(client *http.Client) *DeleteFlowsMilestoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete flows milestone params
func (o *DeleteFlowsMilestoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMilestoneID adds the milestoneID to the delete flows milestone params
func (o *DeleteFlowsMilestoneParams) WithMilestoneID(milestoneID string) *DeleteFlowsMilestoneParams {
	o.SetMilestoneID(milestoneID)
	return o
}

// SetMilestoneID adds the milestoneId to the delete flows milestone params
func (o *DeleteFlowsMilestoneParams) SetMilestoneID(milestoneID string) {
	o.MilestoneID = milestoneID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFlowsMilestoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
