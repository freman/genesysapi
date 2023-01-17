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

// NewGetOutboundSettingsParams creates a new GetOutboundSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOutboundSettingsParams() *GetOutboundSettingsParams {
	return &GetOutboundSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundSettingsParamsWithTimeout creates a new GetOutboundSettingsParams object
// with the ability to set a timeout on a request.
func NewGetOutboundSettingsParamsWithTimeout(timeout time.Duration) *GetOutboundSettingsParams {
	return &GetOutboundSettingsParams{
		timeout: timeout,
	}
}

// NewGetOutboundSettingsParamsWithContext creates a new GetOutboundSettingsParams object
// with the ability to set a context for a request.
func NewGetOutboundSettingsParamsWithContext(ctx context.Context) *GetOutboundSettingsParams {
	return &GetOutboundSettingsParams{
		Context: ctx,
	}
}

// NewGetOutboundSettingsParamsWithHTTPClient creates a new GetOutboundSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOutboundSettingsParamsWithHTTPClient(client *http.Client) *GetOutboundSettingsParams {
	return &GetOutboundSettingsParams{
		HTTPClient: client,
	}
}

/*
GetOutboundSettingsParams contains all the parameters to send to the API endpoint

	for the get outbound settings operation.

	Typically these are written to a http.Request.
*/
type GetOutboundSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get outbound settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundSettingsParams) WithDefaults() *GetOutboundSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get outbound settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get outbound settings params
func (o *GetOutboundSettingsParams) WithTimeout(timeout time.Duration) *GetOutboundSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound settings params
func (o *GetOutboundSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound settings params
func (o *GetOutboundSettingsParams) WithContext(ctx context.Context) *GetOutboundSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound settings params
func (o *GetOutboundSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound settings params
func (o *GetOutboundSettingsParams) WithHTTPClient(client *http.Client) *GetOutboundSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound settings params
func (o *GetOutboundSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
