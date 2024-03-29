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

// NewGetFlowParams creates a new GetFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFlowParams() *GetFlowParams {
	return &GetFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowParamsWithTimeout creates a new GetFlowParams object
// with the ability to set a timeout on a request.
func NewGetFlowParamsWithTimeout(timeout time.Duration) *GetFlowParams {
	return &GetFlowParams{
		timeout: timeout,
	}
}

// NewGetFlowParamsWithContext creates a new GetFlowParams object
// with the ability to set a context for a request.
func NewGetFlowParamsWithContext(ctx context.Context) *GetFlowParams {
	return &GetFlowParams{
		Context: ctx,
	}
}

// NewGetFlowParamsWithHTTPClient creates a new GetFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFlowParamsWithHTTPClient(client *http.Client) *GetFlowParams {
	return &GetFlowParams{
		HTTPClient: client,
	}
}

/*
GetFlowParams contains all the parameters to send to the API endpoint

	for the get flow operation.

	Typically these are written to a http.Request.
*/
type GetFlowParams struct {

	/* Deleted.

	   Deleted flows
	*/
	Deleted *bool

	/* FlowID.

	   Flow ID
	*/
	FlowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowParams) WithDefaults() *GetFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowParams) SetDefaults() {
	var (
		deletedDefault = bool(false)
	)

	val := GetFlowParams{
		Deleted: &deletedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get flow params
func (o *GetFlowParams) WithTimeout(timeout time.Duration) *GetFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flow params
func (o *GetFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flow params
func (o *GetFlowParams) WithContext(ctx context.Context) *GetFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flow params
func (o *GetFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flow params
func (o *GetFlowParams) WithHTTPClient(client *http.Client) *GetFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flow params
func (o *GetFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleted adds the deleted to the get flow params
func (o *GetFlowParams) WithDeleted(deleted *bool) *GetFlowParams {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the get flow params
func (o *GetFlowParams) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithFlowID adds the flowID to the get flow params
func (o *GetFlowParams) WithFlowID(flowID string) *GetFlowParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the get flow params
func (o *GetFlowParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool

		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {

			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}
	}

	// path param flowId
	if err := r.SetPathParam("flowId", o.FlowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
