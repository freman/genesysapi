// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

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

// NewGetTelephonyProvidersEdgesDidpoolParams creates a new GetTelephonyProvidersEdgesDidpoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTelephonyProvidersEdgesDidpoolParams() *GetTelephonyProvidersEdgesDidpoolParams {
	return &GetTelephonyProvidersEdgesDidpoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesDidpoolParamsWithTimeout creates a new GetTelephonyProvidersEdgesDidpoolParams object
// with the ability to set a timeout on a request.
func NewGetTelephonyProvidersEdgesDidpoolParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesDidpoolParams {
	return &GetTelephonyProvidersEdgesDidpoolParams{
		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesDidpoolParamsWithContext creates a new GetTelephonyProvidersEdgesDidpoolParams object
// with the ability to set a context for a request.
func NewGetTelephonyProvidersEdgesDidpoolParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesDidpoolParams {
	return &GetTelephonyProvidersEdgesDidpoolParams{
		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesDidpoolParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesDidpoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTelephonyProvidersEdgesDidpoolParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesDidpoolParams {
	return &GetTelephonyProvidersEdgesDidpoolParams{
		HTTPClient: client,
	}
}

/*
GetTelephonyProvidersEdgesDidpoolParams contains all the parameters to send to the API endpoint

	for the get telephony providers edges didpool operation.

	Typically these are written to a http.Request.
*/
type GetTelephonyProvidersEdgesDidpoolParams struct {

	/* DidPoolID.

	   DID pool ID
	*/
	DidPoolID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get telephony providers edges didpool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesDidpoolParams) WithDefaults() *GetTelephonyProvidersEdgesDidpoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get telephony providers edges didpool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesDidpoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get telephony providers edges didpool params
func (o *GetTelephonyProvidersEdgesDidpoolParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesDidpoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges didpool params
func (o *GetTelephonyProvidersEdgesDidpoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges didpool params
func (o *GetTelephonyProvidersEdgesDidpoolParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesDidpoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges didpool params
func (o *GetTelephonyProvidersEdgesDidpoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges didpool params
func (o *GetTelephonyProvidersEdgesDidpoolParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesDidpoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges didpool params
func (o *GetTelephonyProvidersEdgesDidpoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDidPoolID adds the didPoolID to the get telephony providers edges didpool params
func (o *GetTelephonyProvidersEdgesDidpoolParams) WithDidPoolID(didPoolID string) *GetTelephonyProvidersEdgesDidpoolParams {
	o.SetDidPoolID(didPoolID)
	return o
}

// SetDidPoolID adds the didPoolId to the get telephony providers edges didpool params
func (o *GetTelephonyProvidersEdgesDidpoolParams) SetDidPoolID(didPoolID string) {
	o.DidPoolID = didPoolID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesDidpoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param didPoolId
	if err := r.SetPathParam("didPoolId", o.DidPoolID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
