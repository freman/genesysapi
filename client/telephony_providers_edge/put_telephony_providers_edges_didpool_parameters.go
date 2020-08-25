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

// NewPutTelephonyProvidersEdgesDidpoolParams creates a new PutTelephonyProvidersEdgesDidpoolParams object
// with the default values initialized.
func NewPutTelephonyProvidersEdgesDidpoolParams() *PutTelephonyProvidersEdgesDidpoolParams {
	var ()
	return &PutTelephonyProvidersEdgesDidpoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTelephonyProvidersEdgesDidpoolParamsWithTimeout creates a new PutTelephonyProvidersEdgesDidpoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTelephonyProvidersEdgesDidpoolParamsWithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesDidpoolParams {
	var ()
	return &PutTelephonyProvidersEdgesDidpoolParams{

		timeout: timeout,
	}
}

// NewPutTelephonyProvidersEdgesDidpoolParamsWithContext creates a new PutTelephonyProvidersEdgesDidpoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTelephonyProvidersEdgesDidpoolParamsWithContext(ctx context.Context) *PutTelephonyProvidersEdgesDidpoolParams {
	var ()
	return &PutTelephonyProvidersEdgesDidpoolParams{

		Context: ctx,
	}
}

// NewPutTelephonyProvidersEdgesDidpoolParamsWithHTTPClient creates a new PutTelephonyProvidersEdgesDidpoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTelephonyProvidersEdgesDidpoolParamsWithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesDidpoolParams {
	var ()
	return &PutTelephonyProvidersEdgesDidpoolParams{
		HTTPClient: client,
	}
}

/*PutTelephonyProvidersEdgesDidpoolParams contains all the parameters to send to the API endpoint
for the put telephony providers edges didpool operation typically these are written to a http.Request
*/
type PutTelephonyProvidersEdgesDidpoolParams struct {

	/*Body
	  DID pool

	*/
	Body *models.DIDPool
	/*DidPoolID
	  DID pool ID

	*/
	DidPoolID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put telephony providers edges didpool params
func (o *PutTelephonyProvidersEdgesDidpoolParams) WithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesDidpoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put telephony providers edges didpool params
func (o *PutTelephonyProvidersEdgesDidpoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put telephony providers edges didpool params
func (o *PutTelephonyProvidersEdgesDidpoolParams) WithContext(ctx context.Context) *PutTelephonyProvidersEdgesDidpoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put telephony providers edges didpool params
func (o *PutTelephonyProvidersEdgesDidpoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put telephony providers edges didpool params
func (o *PutTelephonyProvidersEdgesDidpoolParams) WithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesDidpoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put telephony providers edges didpool params
func (o *PutTelephonyProvidersEdgesDidpoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put telephony providers edges didpool params
func (o *PutTelephonyProvidersEdgesDidpoolParams) WithBody(body *models.DIDPool) *PutTelephonyProvidersEdgesDidpoolParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put telephony providers edges didpool params
func (o *PutTelephonyProvidersEdgesDidpoolParams) SetBody(body *models.DIDPool) {
	o.Body = body
}

// WithDidPoolID adds the didPoolID to the put telephony providers edges didpool params
func (o *PutTelephonyProvidersEdgesDidpoolParams) WithDidPoolID(didPoolID string) *PutTelephonyProvidersEdgesDidpoolParams {
	o.SetDidPoolID(didPoolID)
	return o
}

// SetDidPoolID adds the didPoolId to the put telephony providers edges didpool params
func (o *PutTelephonyProvidersEdgesDidpoolParams) SetDidPoolID(didPoolID string) {
	o.DidPoolID = didPoolID
}

// WriteToRequest writes these params to a swagger request
func (o *PutTelephonyProvidersEdgesDidpoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param didPoolId
	if err := r.SetPathParam("didPoolId", o.DidPoolID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}