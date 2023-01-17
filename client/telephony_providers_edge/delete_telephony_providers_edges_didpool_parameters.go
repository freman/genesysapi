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

// NewDeleteTelephonyProvidersEdgesDidpoolParams creates a new DeleteTelephonyProvidersEdgesDidpoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTelephonyProvidersEdgesDidpoolParams() *DeleteTelephonyProvidersEdgesDidpoolParams {
	return &DeleteTelephonyProvidersEdgesDidpoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTelephonyProvidersEdgesDidpoolParamsWithTimeout creates a new DeleteTelephonyProvidersEdgesDidpoolParams object
// with the ability to set a timeout on a request.
func NewDeleteTelephonyProvidersEdgesDidpoolParamsWithTimeout(timeout time.Duration) *DeleteTelephonyProvidersEdgesDidpoolParams {
	return &DeleteTelephonyProvidersEdgesDidpoolParams{
		timeout: timeout,
	}
}

// NewDeleteTelephonyProvidersEdgesDidpoolParamsWithContext creates a new DeleteTelephonyProvidersEdgesDidpoolParams object
// with the ability to set a context for a request.
func NewDeleteTelephonyProvidersEdgesDidpoolParamsWithContext(ctx context.Context) *DeleteTelephonyProvidersEdgesDidpoolParams {
	return &DeleteTelephonyProvidersEdgesDidpoolParams{
		Context: ctx,
	}
}

// NewDeleteTelephonyProvidersEdgesDidpoolParamsWithHTTPClient creates a new DeleteTelephonyProvidersEdgesDidpoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTelephonyProvidersEdgesDidpoolParamsWithHTTPClient(client *http.Client) *DeleteTelephonyProvidersEdgesDidpoolParams {
	return &DeleteTelephonyProvidersEdgesDidpoolParams{
		HTTPClient: client,
	}
}

/*
DeleteTelephonyProvidersEdgesDidpoolParams contains all the parameters to send to the API endpoint

	for the delete telephony providers edges didpool operation.

	Typically these are written to a http.Request.
*/
type DeleteTelephonyProvidersEdgesDidpoolParams struct {

	/* DidPoolID.

	   DID pool ID
	*/
	DidPoolID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete telephony providers edges didpool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTelephonyProvidersEdgesDidpoolParams) WithDefaults() *DeleteTelephonyProvidersEdgesDidpoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete telephony providers edges didpool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTelephonyProvidersEdgesDidpoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete telephony providers edges didpool params
func (o *DeleteTelephonyProvidersEdgesDidpoolParams) WithTimeout(timeout time.Duration) *DeleteTelephonyProvidersEdgesDidpoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete telephony providers edges didpool params
func (o *DeleteTelephonyProvidersEdgesDidpoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete telephony providers edges didpool params
func (o *DeleteTelephonyProvidersEdgesDidpoolParams) WithContext(ctx context.Context) *DeleteTelephonyProvidersEdgesDidpoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete telephony providers edges didpool params
func (o *DeleteTelephonyProvidersEdgesDidpoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete telephony providers edges didpool params
func (o *DeleteTelephonyProvidersEdgesDidpoolParams) WithHTTPClient(client *http.Client) *DeleteTelephonyProvidersEdgesDidpoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete telephony providers edges didpool params
func (o *DeleteTelephonyProvidersEdgesDidpoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDidPoolID adds the didPoolID to the delete telephony providers edges didpool params
func (o *DeleteTelephonyProvidersEdgesDidpoolParams) WithDidPoolID(didPoolID string) *DeleteTelephonyProvidersEdgesDidpoolParams {
	o.SetDidPoolID(didPoolID)
	return o
}

// SetDidPoolID adds the didPoolId to the delete telephony providers edges didpool params
func (o *DeleteTelephonyProvidersEdgesDidpoolParams) SetDidPoolID(didPoolID string) {
	o.DidPoolID = didPoolID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTelephonyProvidersEdgesDidpoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
