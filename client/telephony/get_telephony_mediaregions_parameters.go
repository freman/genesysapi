// Code generated by go-swagger; DO NOT EDIT.

package telephony

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

// NewGetTelephonyMediaregionsParams creates a new GetTelephonyMediaregionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTelephonyMediaregionsParams() *GetTelephonyMediaregionsParams {
	return &GetTelephonyMediaregionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyMediaregionsParamsWithTimeout creates a new GetTelephonyMediaregionsParams object
// with the ability to set a timeout on a request.
func NewGetTelephonyMediaregionsParamsWithTimeout(timeout time.Duration) *GetTelephonyMediaregionsParams {
	return &GetTelephonyMediaregionsParams{
		timeout: timeout,
	}
}

// NewGetTelephonyMediaregionsParamsWithContext creates a new GetTelephonyMediaregionsParams object
// with the ability to set a context for a request.
func NewGetTelephonyMediaregionsParamsWithContext(ctx context.Context) *GetTelephonyMediaregionsParams {
	return &GetTelephonyMediaregionsParams{
		Context: ctx,
	}
}

// NewGetTelephonyMediaregionsParamsWithHTTPClient creates a new GetTelephonyMediaregionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTelephonyMediaregionsParamsWithHTTPClient(client *http.Client) *GetTelephonyMediaregionsParams {
	return &GetTelephonyMediaregionsParams{
		HTTPClient: client,
	}
}

/*
GetTelephonyMediaregionsParams contains all the parameters to send to the API endpoint

	for the get telephony mediaregions operation.

	Typically these are written to a http.Request.
*/
type GetTelephonyMediaregionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get telephony mediaregions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyMediaregionsParams) WithDefaults() *GetTelephonyMediaregionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get telephony mediaregions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyMediaregionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get telephony mediaregions params
func (o *GetTelephonyMediaregionsParams) WithTimeout(timeout time.Duration) *GetTelephonyMediaregionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony mediaregions params
func (o *GetTelephonyMediaregionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony mediaregions params
func (o *GetTelephonyMediaregionsParams) WithContext(ctx context.Context) *GetTelephonyMediaregionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony mediaregions params
func (o *GetTelephonyMediaregionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony mediaregions params
func (o *GetTelephonyMediaregionsParams) WithHTTPClient(client *http.Client) *GetTelephonyMediaregionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony mediaregions params
func (o *GetTelephonyMediaregionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyMediaregionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
