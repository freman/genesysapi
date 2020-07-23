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
)

// NewDeleteRoutingSmsAddressParams creates a new DeleteRoutingSmsAddressParams object
// with the default values initialized.
func NewDeleteRoutingSmsAddressParams() *DeleteRoutingSmsAddressParams {
	var ()
	return &DeleteRoutingSmsAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoutingSmsAddressParamsWithTimeout creates a new DeleteRoutingSmsAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRoutingSmsAddressParamsWithTimeout(timeout time.Duration) *DeleteRoutingSmsAddressParams {
	var ()
	return &DeleteRoutingSmsAddressParams{

		timeout: timeout,
	}
}

// NewDeleteRoutingSmsAddressParamsWithContext creates a new DeleteRoutingSmsAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRoutingSmsAddressParamsWithContext(ctx context.Context) *DeleteRoutingSmsAddressParams {
	var ()
	return &DeleteRoutingSmsAddressParams{

		Context: ctx,
	}
}

// NewDeleteRoutingSmsAddressParamsWithHTTPClient creates a new DeleteRoutingSmsAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRoutingSmsAddressParamsWithHTTPClient(client *http.Client) *DeleteRoutingSmsAddressParams {
	var ()
	return &DeleteRoutingSmsAddressParams{
		HTTPClient: client,
	}
}

/*DeleteRoutingSmsAddressParams contains all the parameters to send to the API endpoint
for the delete routing sms address operation typically these are written to a http.Request
*/
type DeleteRoutingSmsAddressParams struct {

	/*AddressID
	  Address ID

	*/
	AddressID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete routing sms address params
func (o *DeleteRoutingSmsAddressParams) WithTimeout(timeout time.Duration) *DeleteRoutingSmsAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete routing sms address params
func (o *DeleteRoutingSmsAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete routing sms address params
func (o *DeleteRoutingSmsAddressParams) WithContext(ctx context.Context) *DeleteRoutingSmsAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete routing sms address params
func (o *DeleteRoutingSmsAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete routing sms address params
func (o *DeleteRoutingSmsAddressParams) WithHTTPClient(client *http.Client) *DeleteRoutingSmsAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete routing sms address params
func (o *DeleteRoutingSmsAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressID adds the addressID to the delete routing sms address params
func (o *DeleteRoutingSmsAddressParams) WithAddressID(addressID string) *DeleteRoutingSmsAddressParams {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the delete routing sms address params
func (o *DeleteRoutingSmsAddressParams) SetAddressID(addressID string) {
	o.AddressID = addressID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoutingSmsAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addressId
	if err := r.SetPathParam("addressId", o.AddressID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
