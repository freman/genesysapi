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

// NewGetTelephonyProvidersEdgesPhoneParams creates a new GetTelephonyProvidersEdgesPhoneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTelephonyProvidersEdgesPhoneParams() *GetTelephonyProvidersEdgesPhoneParams {
	return &GetTelephonyProvidersEdgesPhoneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesPhoneParamsWithTimeout creates a new GetTelephonyProvidersEdgesPhoneParams object
// with the ability to set a timeout on a request.
func NewGetTelephonyProvidersEdgesPhoneParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesPhoneParams {
	return &GetTelephonyProvidersEdgesPhoneParams{
		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesPhoneParamsWithContext creates a new GetTelephonyProvidersEdgesPhoneParams object
// with the ability to set a context for a request.
func NewGetTelephonyProvidersEdgesPhoneParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesPhoneParams {
	return &GetTelephonyProvidersEdgesPhoneParams{
		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesPhoneParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesPhoneParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTelephonyProvidersEdgesPhoneParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesPhoneParams {
	return &GetTelephonyProvidersEdgesPhoneParams{
		HTTPClient: client,
	}
}

/*
GetTelephonyProvidersEdgesPhoneParams contains all the parameters to send to the API endpoint

	for the get telephony providers edges phone operation.

	Typically these are written to a http.Request.
*/
type GetTelephonyProvidersEdgesPhoneParams struct {

	/* PhoneID.

	   Phone ID
	*/
	PhoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get telephony providers edges phone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesPhoneParams) WithDefaults() *GetTelephonyProvidersEdgesPhoneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get telephony providers edges phone params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesPhoneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get telephony providers edges phone params
func (o *GetTelephonyProvidersEdgesPhoneParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesPhoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges phone params
func (o *GetTelephonyProvidersEdgesPhoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges phone params
func (o *GetTelephonyProvidersEdgesPhoneParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesPhoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges phone params
func (o *GetTelephonyProvidersEdgesPhoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges phone params
func (o *GetTelephonyProvidersEdgesPhoneParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesPhoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges phone params
func (o *GetTelephonyProvidersEdgesPhoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPhoneID adds the phoneID to the get telephony providers edges phone params
func (o *GetTelephonyProvidersEdgesPhoneParams) WithPhoneID(phoneID string) *GetTelephonyProvidersEdgesPhoneParams {
	o.SetPhoneID(phoneID)
	return o
}

// SetPhoneID adds the phoneId to the get telephony providers edges phone params
func (o *GetTelephonyProvidersEdgesPhoneParams) SetPhoneID(phoneID string) {
	o.PhoneID = phoneID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesPhoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param phoneId
	if err := r.SetPathParam("phoneId", o.PhoneID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
