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

// NewGetTelephonyProvidersEdgesDidParams creates a new GetTelephonyProvidersEdgesDidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTelephonyProvidersEdgesDidParams() *GetTelephonyProvidersEdgesDidParams {
	return &GetTelephonyProvidersEdgesDidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesDidParamsWithTimeout creates a new GetTelephonyProvidersEdgesDidParams object
// with the ability to set a timeout on a request.
func NewGetTelephonyProvidersEdgesDidParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesDidParams {
	return &GetTelephonyProvidersEdgesDidParams{
		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesDidParamsWithContext creates a new GetTelephonyProvidersEdgesDidParams object
// with the ability to set a context for a request.
func NewGetTelephonyProvidersEdgesDidParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesDidParams {
	return &GetTelephonyProvidersEdgesDidParams{
		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesDidParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesDidParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTelephonyProvidersEdgesDidParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesDidParams {
	return &GetTelephonyProvidersEdgesDidParams{
		HTTPClient: client,
	}
}

/*
GetTelephonyProvidersEdgesDidParams contains all the parameters to send to the API endpoint

	for the get telephony providers edges did operation.

	Typically these are written to a http.Request.
*/
type GetTelephonyProvidersEdgesDidParams struct {

	/* DidID.

	   DID ID
	*/
	DidID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get telephony providers edges did params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesDidParams) WithDefaults() *GetTelephonyProvidersEdgesDidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get telephony providers edges did params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesDidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get telephony providers edges did params
func (o *GetTelephonyProvidersEdgesDidParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesDidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges did params
func (o *GetTelephonyProvidersEdgesDidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges did params
func (o *GetTelephonyProvidersEdgesDidParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesDidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges did params
func (o *GetTelephonyProvidersEdgesDidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges did params
func (o *GetTelephonyProvidersEdgesDidParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesDidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges did params
func (o *GetTelephonyProvidersEdgesDidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDidID adds the didID to the get telephony providers edges did params
func (o *GetTelephonyProvidersEdgesDidParams) WithDidID(didID string) *GetTelephonyProvidersEdgesDidParams {
	o.SetDidID(didID)
	return o
}

// SetDidID adds the didId to the get telephony providers edges did params
func (o *GetTelephonyProvidersEdgesDidParams) SetDidID(didID string) {
	o.DidID = didID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesDidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param didId
	if err := r.SetPathParam("didId", o.DidID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
