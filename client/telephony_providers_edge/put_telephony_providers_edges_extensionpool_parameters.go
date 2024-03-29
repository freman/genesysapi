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

	"github.com/freman/genesysapi/models"
)

// NewPutTelephonyProvidersEdgesExtensionpoolParams creates a new PutTelephonyProvidersEdgesExtensionpoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutTelephonyProvidersEdgesExtensionpoolParams() *PutTelephonyProvidersEdgesExtensionpoolParams {
	return &PutTelephonyProvidersEdgesExtensionpoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutTelephonyProvidersEdgesExtensionpoolParamsWithTimeout creates a new PutTelephonyProvidersEdgesExtensionpoolParams object
// with the ability to set a timeout on a request.
func NewPutTelephonyProvidersEdgesExtensionpoolParamsWithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesExtensionpoolParams {
	return &PutTelephonyProvidersEdgesExtensionpoolParams{
		timeout: timeout,
	}
}

// NewPutTelephonyProvidersEdgesExtensionpoolParamsWithContext creates a new PutTelephonyProvidersEdgesExtensionpoolParams object
// with the ability to set a context for a request.
func NewPutTelephonyProvidersEdgesExtensionpoolParamsWithContext(ctx context.Context) *PutTelephonyProvidersEdgesExtensionpoolParams {
	return &PutTelephonyProvidersEdgesExtensionpoolParams{
		Context: ctx,
	}
}

// NewPutTelephonyProvidersEdgesExtensionpoolParamsWithHTTPClient creates a new PutTelephonyProvidersEdgesExtensionpoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutTelephonyProvidersEdgesExtensionpoolParamsWithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesExtensionpoolParams {
	return &PutTelephonyProvidersEdgesExtensionpoolParams{
		HTTPClient: client,
	}
}

/*
PutTelephonyProvidersEdgesExtensionpoolParams contains all the parameters to send to the API endpoint

	for the put telephony providers edges extensionpool operation.

	Typically these are written to a http.Request.
*/
type PutTelephonyProvidersEdgesExtensionpoolParams struct {

	/* Body.

	   ExtensionPool
	*/
	Body *models.ExtensionPool

	/* ExtensionPoolID.

	   Extension pool ID
	*/
	ExtensionPoolID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put telephony providers edges extensionpool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) WithDefaults() *PutTelephonyProvidersEdgesExtensionpoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put telephony providers edges extensionpool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put telephony providers edges extensionpool params
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) WithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesExtensionpoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put telephony providers edges extensionpool params
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put telephony providers edges extensionpool params
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) WithContext(ctx context.Context) *PutTelephonyProvidersEdgesExtensionpoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put telephony providers edges extensionpool params
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put telephony providers edges extensionpool params
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) WithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesExtensionpoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put telephony providers edges extensionpool params
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put telephony providers edges extensionpool params
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) WithBody(body *models.ExtensionPool) *PutTelephonyProvidersEdgesExtensionpoolParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put telephony providers edges extensionpool params
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) SetBody(body *models.ExtensionPool) {
	o.Body = body
}

// WithExtensionPoolID adds the extensionPoolID to the put telephony providers edges extensionpool params
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) WithExtensionPoolID(extensionPoolID string) *PutTelephonyProvidersEdgesExtensionpoolParams {
	o.SetExtensionPoolID(extensionPoolID)
	return o
}

// SetExtensionPoolID adds the extensionPoolId to the put telephony providers edges extensionpool params
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) SetExtensionPoolID(extensionPoolID string) {
	o.ExtensionPoolID = extensionPoolID
}

// WriteToRequest writes these params to a swagger request
func (o *PutTelephonyProvidersEdgesExtensionpoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param extensionPoolId
	if err := r.SetPathParam("extensionPoolId", o.ExtensionPoolID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
