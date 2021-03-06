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

// NewPutTelephonyProvidersEdgesPhonebasesettingParams creates a new PutTelephonyProvidersEdgesPhonebasesettingParams object
// with the default values initialized.
func NewPutTelephonyProvidersEdgesPhonebasesettingParams() *PutTelephonyProvidersEdgesPhonebasesettingParams {
	var ()
	return &PutTelephonyProvidersEdgesPhonebasesettingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTelephonyProvidersEdgesPhonebasesettingParamsWithTimeout creates a new PutTelephonyProvidersEdgesPhonebasesettingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTelephonyProvidersEdgesPhonebasesettingParamsWithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesPhonebasesettingParams {
	var ()
	return &PutTelephonyProvidersEdgesPhonebasesettingParams{

		timeout: timeout,
	}
}

// NewPutTelephonyProvidersEdgesPhonebasesettingParamsWithContext creates a new PutTelephonyProvidersEdgesPhonebasesettingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTelephonyProvidersEdgesPhonebasesettingParamsWithContext(ctx context.Context) *PutTelephonyProvidersEdgesPhonebasesettingParams {
	var ()
	return &PutTelephonyProvidersEdgesPhonebasesettingParams{

		Context: ctx,
	}
}

// NewPutTelephonyProvidersEdgesPhonebasesettingParamsWithHTTPClient creates a new PutTelephonyProvidersEdgesPhonebasesettingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTelephonyProvidersEdgesPhonebasesettingParamsWithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesPhonebasesettingParams {
	var ()
	return &PutTelephonyProvidersEdgesPhonebasesettingParams{
		HTTPClient: client,
	}
}

/*PutTelephonyProvidersEdgesPhonebasesettingParams contains all the parameters to send to the API endpoint
for the put telephony providers edges phonebasesetting operation typically these are written to a http.Request
*/
type PutTelephonyProvidersEdgesPhonebasesettingParams struct {

	/*Body
	  Phone base settings

	*/
	Body *models.PhoneBase
	/*PhoneBaseID
	  Phone base ID

	*/
	PhoneBaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put telephony providers edges phonebasesetting params
func (o *PutTelephonyProvidersEdgesPhonebasesettingParams) WithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesPhonebasesettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put telephony providers edges phonebasesetting params
func (o *PutTelephonyProvidersEdgesPhonebasesettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put telephony providers edges phonebasesetting params
func (o *PutTelephonyProvidersEdgesPhonebasesettingParams) WithContext(ctx context.Context) *PutTelephonyProvidersEdgesPhonebasesettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put telephony providers edges phonebasesetting params
func (o *PutTelephonyProvidersEdgesPhonebasesettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put telephony providers edges phonebasesetting params
func (o *PutTelephonyProvidersEdgesPhonebasesettingParams) WithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesPhonebasesettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put telephony providers edges phonebasesetting params
func (o *PutTelephonyProvidersEdgesPhonebasesettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put telephony providers edges phonebasesetting params
func (o *PutTelephonyProvidersEdgesPhonebasesettingParams) WithBody(body *models.PhoneBase) *PutTelephonyProvidersEdgesPhonebasesettingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put telephony providers edges phonebasesetting params
func (o *PutTelephonyProvidersEdgesPhonebasesettingParams) SetBody(body *models.PhoneBase) {
	o.Body = body
}

// WithPhoneBaseID adds the phoneBaseID to the put telephony providers edges phonebasesetting params
func (o *PutTelephonyProvidersEdgesPhonebasesettingParams) WithPhoneBaseID(phoneBaseID string) *PutTelephonyProvidersEdgesPhonebasesettingParams {
	o.SetPhoneBaseID(phoneBaseID)
	return o
}

// SetPhoneBaseID adds the phoneBaseId to the put telephony providers edges phonebasesetting params
func (o *PutTelephonyProvidersEdgesPhonebasesettingParams) SetPhoneBaseID(phoneBaseID string) {
	o.PhoneBaseID = phoneBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *PutTelephonyProvidersEdgesPhonebasesettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param phoneBaseId
	if err := r.SetPathParam("phoneBaseId", o.PhoneBaseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
