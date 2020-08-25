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

// NewGetTelephonyProvidersEdgesPhonebasesettingParams creates a new GetTelephonyProvidersEdgesPhonebasesettingParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesPhonebasesettingParams() *GetTelephonyProvidersEdgesPhonebasesettingParams {
	var ()
	return &GetTelephonyProvidersEdgesPhonebasesettingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesPhonebasesettingParamsWithTimeout creates a new GetTelephonyProvidersEdgesPhonebasesettingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesPhonebasesettingParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesPhonebasesettingParams {
	var ()
	return &GetTelephonyProvidersEdgesPhonebasesettingParams{

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesPhonebasesettingParamsWithContext creates a new GetTelephonyProvidersEdgesPhonebasesettingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesPhonebasesettingParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesPhonebasesettingParams {
	var ()
	return &GetTelephonyProvidersEdgesPhonebasesettingParams{

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesPhonebasesettingParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesPhonebasesettingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesPhonebasesettingParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesPhonebasesettingParams {
	var ()
	return &GetTelephonyProvidersEdgesPhonebasesettingParams{
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesPhonebasesettingParams contains all the parameters to send to the API endpoint
for the get telephony providers edges phonebasesetting operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesPhonebasesettingParams struct {

	/*PhoneBaseID
	  Phone base ID

	*/
	PhoneBaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edges phonebasesetting params
func (o *GetTelephonyProvidersEdgesPhonebasesettingParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesPhonebasesettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges phonebasesetting params
func (o *GetTelephonyProvidersEdgesPhonebasesettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges phonebasesetting params
func (o *GetTelephonyProvidersEdgesPhonebasesettingParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesPhonebasesettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges phonebasesetting params
func (o *GetTelephonyProvidersEdgesPhonebasesettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges phonebasesetting params
func (o *GetTelephonyProvidersEdgesPhonebasesettingParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesPhonebasesettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges phonebasesetting params
func (o *GetTelephonyProvidersEdgesPhonebasesettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPhoneBaseID adds the phoneBaseID to the get telephony providers edges phonebasesetting params
func (o *GetTelephonyProvidersEdgesPhonebasesettingParams) WithPhoneBaseID(phoneBaseID string) *GetTelephonyProvidersEdgesPhonebasesettingParams {
	o.SetPhoneBaseID(phoneBaseID)
	return o
}

// SetPhoneBaseID adds the phoneBaseId to the get telephony providers edges phonebasesetting params
func (o *GetTelephonyProvidersEdgesPhonebasesettingParams) SetPhoneBaseID(phoneBaseID string) {
	o.PhoneBaseID = phoneBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesPhonebasesettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param phoneBaseId
	if err := r.SetPathParam("phoneBaseId", o.PhoneBaseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}