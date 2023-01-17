// Code generated by go-swagger; DO NOT EDIT.

package outbound

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

// NewGetOutboundCallanalysisresponsesetParams creates a new GetOutboundCallanalysisresponsesetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOutboundCallanalysisresponsesetParams() *GetOutboundCallanalysisresponsesetParams {
	return &GetOutboundCallanalysisresponsesetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundCallanalysisresponsesetParamsWithTimeout creates a new GetOutboundCallanalysisresponsesetParams object
// with the ability to set a timeout on a request.
func NewGetOutboundCallanalysisresponsesetParamsWithTimeout(timeout time.Duration) *GetOutboundCallanalysisresponsesetParams {
	return &GetOutboundCallanalysisresponsesetParams{
		timeout: timeout,
	}
}

// NewGetOutboundCallanalysisresponsesetParamsWithContext creates a new GetOutboundCallanalysisresponsesetParams object
// with the ability to set a context for a request.
func NewGetOutboundCallanalysisresponsesetParamsWithContext(ctx context.Context) *GetOutboundCallanalysisresponsesetParams {
	return &GetOutboundCallanalysisresponsesetParams{
		Context: ctx,
	}
}

// NewGetOutboundCallanalysisresponsesetParamsWithHTTPClient creates a new GetOutboundCallanalysisresponsesetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOutboundCallanalysisresponsesetParamsWithHTTPClient(client *http.Client) *GetOutboundCallanalysisresponsesetParams {
	return &GetOutboundCallanalysisresponsesetParams{
		HTTPClient: client,
	}
}

/*
GetOutboundCallanalysisresponsesetParams contains all the parameters to send to the API endpoint

	for the get outbound callanalysisresponseset operation.

	Typically these are written to a http.Request.
*/
type GetOutboundCallanalysisresponsesetParams struct {

	/* CallAnalysisSetID.

	   Call Analysis Response Set ID
	*/
	CallAnalysisSetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get outbound callanalysisresponseset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundCallanalysisresponsesetParams) WithDefaults() *GetOutboundCallanalysisresponsesetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get outbound callanalysisresponseset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundCallanalysisresponsesetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get outbound callanalysisresponseset params
func (o *GetOutboundCallanalysisresponsesetParams) WithTimeout(timeout time.Duration) *GetOutboundCallanalysisresponsesetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound callanalysisresponseset params
func (o *GetOutboundCallanalysisresponsesetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound callanalysisresponseset params
func (o *GetOutboundCallanalysisresponsesetParams) WithContext(ctx context.Context) *GetOutboundCallanalysisresponsesetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound callanalysisresponseset params
func (o *GetOutboundCallanalysisresponsesetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound callanalysisresponseset params
func (o *GetOutboundCallanalysisresponsesetParams) WithHTTPClient(client *http.Client) *GetOutboundCallanalysisresponsesetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound callanalysisresponseset params
func (o *GetOutboundCallanalysisresponsesetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCallAnalysisSetID adds the callAnalysisSetID to the get outbound callanalysisresponseset params
func (o *GetOutboundCallanalysisresponsesetParams) WithCallAnalysisSetID(callAnalysisSetID string) *GetOutboundCallanalysisresponsesetParams {
	o.SetCallAnalysisSetID(callAnalysisSetID)
	return o
}

// SetCallAnalysisSetID adds the callAnalysisSetId to the get outbound callanalysisresponseset params
func (o *GetOutboundCallanalysisresponsesetParams) SetCallAnalysisSetID(callAnalysisSetID string) {
	o.CallAnalysisSetID = callAnalysisSetID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundCallanalysisresponsesetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param callAnalysisSetId
	if err := r.SetPathParam("callAnalysisSetId", o.CallAnalysisSetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
