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

// NewGetFlowLatestconfigurationParams creates a new GetFlowLatestconfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFlowLatestconfigurationParams() *GetFlowLatestconfigurationParams {
	return &GetFlowLatestconfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowLatestconfigurationParamsWithTimeout creates a new GetFlowLatestconfigurationParams object
// with the ability to set a timeout on a request.
func NewGetFlowLatestconfigurationParamsWithTimeout(timeout time.Duration) *GetFlowLatestconfigurationParams {
	return &GetFlowLatestconfigurationParams{
		timeout: timeout,
	}
}

// NewGetFlowLatestconfigurationParamsWithContext creates a new GetFlowLatestconfigurationParams object
// with the ability to set a context for a request.
func NewGetFlowLatestconfigurationParamsWithContext(ctx context.Context) *GetFlowLatestconfigurationParams {
	return &GetFlowLatestconfigurationParams{
		Context: ctx,
	}
}

// NewGetFlowLatestconfigurationParamsWithHTTPClient creates a new GetFlowLatestconfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFlowLatestconfigurationParamsWithHTTPClient(client *http.Client) *GetFlowLatestconfigurationParams {
	return &GetFlowLatestconfigurationParams{
		HTTPClient: client,
	}
}

/*
GetFlowLatestconfigurationParams contains all the parameters to send to the API endpoint

	for the get flow latestconfiguration operation.

	Typically these are written to a http.Request.
*/
type GetFlowLatestconfigurationParams struct {

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

// WithDefaults hydrates default values in the get flow latestconfiguration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowLatestconfigurationParams) WithDefaults() *GetFlowLatestconfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get flow latestconfiguration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowLatestconfigurationParams) SetDefaults() {
	var (
		deletedDefault = bool(false)
	)

	val := GetFlowLatestconfigurationParams{
		Deleted: &deletedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get flow latestconfiguration params
func (o *GetFlowLatestconfigurationParams) WithTimeout(timeout time.Duration) *GetFlowLatestconfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flow latestconfiguration params
func (o *GetFlowLatestconfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flow latestconfiguration params
func (o *GetFlowLatestconfigurationParams) WithContext(ctx context.Context) *GetFlowLatestconfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flow latestconfiguration params
func (o *GetFlowLatestconfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flow latestconfiguration params
func (o *GetFlowLatestconfigurationParams) WithHTTPClient(client *http.Client) *GetFlowLatestconfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flow latestconfiguration params
func (o *GetFlowLatestconfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleted adds the deleted to the get flow latestconfiguration params
func (o *GetFlowLatestconfigurationParams) WithDeleted(deleted *bool) *GetFlowLatestconfigurationParams {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the get flow latestconfiguration params
func (o *GetFlowLatestconfigurationParams) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithFlowID adds the flowID to the get flow latestconfiguration params
func (o *GetFlowLatestconfigurationParams) WithFlowID(flowID string) *GetFlowLatestconfigurationParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the get flow latestconfiguration params
func (o *GetFlowLatestconfigurationParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowLatestconfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
