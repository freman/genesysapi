// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewPutRoutingSmsPhonenumberParams creates a new PutRoutingSmsPhonenumberParams object
// with the default values initialized.
func NewPutRoutingSmsPhonenumberParams() *PutRoutingSmsPhonenumberParams {
	var ()
	return &PutRoutingSmsPhonenumberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRoutingSmsPhonenumberParamsWithTimeout creates a new PutRoutingSmsPhonenumberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRoutingSmsPhonenumberParamsWithTimeout(timeout time.Duration) *PutRoutingSmsPhonenumberParams {
	var ()
	return &PutRoutingSmsPhonenumberParams{

		timeout: timeout,
	}
}

// NewPutRoutingSmsPhonenumberParamsWithContext creates a new PutRoutingSmsPhonenumberParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRoutingSmsPhonenumberParamsWithContext(ctx context.Context) *PutRoutingSmsPhonenumberParams {
	var ()
	return &PutRoutingSmsPhonenumberParams{

		Context: ctx,
	}
}

// NewPutRoutingSmsPhonenumberParamsWithHTTPClient creates a new PutRoutingSmsPhonenumberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRoutingSmsPhonenumberParamsWithHTTPClient(client *http.Client) *PutRoutingSmsPhonenumberParams {
	var ()
	return &PutRoutingSmsPhonenumberParams{
		HTTPClient: client,
	}
}

/*PutRoutingSmsPhonenumberParams contains all the parameters to send to the API endpoint
for the put routing sms phonenumber operation typically these are written to a http.Request
*/
type PutRoutingSmsPhonenumberParams struct {

	/*AddressID
	  Address ID

	*/
	AddressID string
	/*Body
	  SmsPhoneNumber

	*/
	Body *models.SmsPhoneNumber

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put routing sms phonenumber params
func (o *PutRoutingSmsPhonenumberParams) WithTimeout(timeout time.Duration) *PutRoutingSmsPhonenumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put routing sms phonenumber params
func (o *PutRoutingSmsPhonenumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put routing sms phonenumber params
func (o *PutRoutingSmsPhonenumberParams) WithContext(ctx context.Context) *PutRoutingSmsPhonenumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put routing sms phonenumber params
func (o *PutRoutingSmsPhonenumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put routing sms phonenumber params
func (o *PutRoutingSmsPhonenumberParams) WithHTTPClient(client *http.Client) *PutRoutingSmsPhonenumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put routing sms phonenumber params
func (o *PutRoutingSmsPhonenumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressID adds the addressID to the put routing sms phonenumber params
func (o *PutRoutingSmsPhonenumberParams) WithAddressID(addressID string) *PutRoutingSmsPhonenumberParams {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the put routing sms phonenumber params
func (o *PutRoutingSmsPhonenumberParams) SetAddressID(addressID string) {
	o.AddressID = addressID
}

// WithBody adds the body to the put routing sms phonenumber params
func (o *PutRoutingSmsPhonenumberParams) WithBody(body *models.SmsPhoneNumber) *PutRoutingSmsPhonenumberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put routing sms phonenumber params
func (o *PutRoutingSmsPhonenumberParams) SetBody(body *models.SmsPhoneNumber) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutRoutingSmsPhonenumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addressId
	if err := r.SetPathParam("addressId", o.AddressID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
