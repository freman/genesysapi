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

// NewPutTelephonyProvidersEdgesPhoneParams creates a new PutTelephonyProvidersEdgesPhoneParams object
// with the default values initialized.
func NewPutTelephonyProvidersEdgesPhoneParams() *PutTelephonyProvidersEdgesPhoneParams {
	var ()
	return &PutTelephonyProvidersEdgesPhoneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTelephonyProvidersEdgesPhoneParamsWithTimeout creates a new PutTelephonyProvidersEdgesPhoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTelephonyProvidersEdgesPhoneParamsWithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesPhoneParams {
	var ()
	return &PutTelephonyProvidersEdgesPhoneParams{

		timeout: timeout,
	}
}

// NewPutTelephonyProvidersEdgesPhoneParamsWithContext creates a new PutTelephonyProvidersEdgesPhoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTelephonyProvidersEdgesPhoneParamsWithContext(ctx context.Context) *PutTelephonyProvidersEdgesPhoneParams {
	var ()
	return &PutTelephonyProvidersEdgesPhoneParams{

		Context: ctx,
	}
}

// NewPutTelephonyProvidersEdgesPhoneParamsWithHTTPClient creates a new PutTelephonyProvidersEdgesPhoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTelephonyProvidersEdgesPhoneParamsWithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesPhoneParams {
	var ()
	return &PutTelephonyProvidersEdgesPhoneParams{
		HTTPClient: client,
	}
}

/*PutTelephonyProvidersEdgesPhoneParams contains all the parameters to send to the API endpoint
for the put telephony providers edges phone operation typically these are written to a http.Request
*/
type PutTelephonyProvidersEdgesPhoneParams struct {

	/*Body
	  Phone

	*/
	Body *models.Phone
	/*PhoneID
	  Phone ID

	*/
	PhoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put telephony providers edges phone params
func (o *PutTelephonyProvidersEdgesPhoneParams) WithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesPhoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put telephony providers edges phone params
func (o *PutTelephonyProvidersEdgesPhoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put telephony providers edges phone params
func (o *PutTelephonyProvidersEdgesPhoneParams) WithContext(ctx context.Context) *PutTelephonyProvidersEdgesPhoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put telephony providers edges phone params
func (o *PutTelephonyProvidersEdgesPhoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put telephony providers edges phone params
func (o *PutTelephonyProvidersEdgesPhoneParams) WithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesPhoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put telephony providers edges phone params
func (o *PutTelephonyProvidersEdgesPhoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put telephony providers edges phone params
func (o *PutTelephonyProvidersEdgesPhoneParams) WithBody(body *models.Phone) *PutTelephonyProvidersEdgesPhoneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put telephony providers edges phone params
func (o *PutTelephonyProvidersEdgesPhoneParams) SetBody(body *models.Phone) {
	o.Body = body
}

// WithPhoneID adds the phoneID to the put telephony providers edges phone params
func (o *PutTelephonyProvidersEdgesPhoneParams) WithPhoneID(phoneID string) *PutTelephonyProvidersEdgesPhoneParams {
	o.SetPhoneID(phoneID)
	return o
}

// SetPhoneID adds the phoneId to the put telephony providers edges phone params
func (o *PutTelephonyProvidersEdgesPhoneParams) SetPhoneID(phoneID string) {
	o.PhoneID = phoneID
}

// WriteToRequest writes these params to a swagger request
func (o *PutTelephonyProvidersEdgesPhoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param phoneId
	if err := r.SetPathParam("phoneId", o.PhoneID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}