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

// NewGetTelephonyProvidersEdgesExtensionParams creates a new GetTelephonyProvidersEdgesExtensionParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesExtensionParams() *GetTelephonyProvidersEdgesExtensionParams {
	var ()
	return &GetTelephonyProvidersEdgesExtensionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesExtensionParamsWithTimeout creates a new GetTelephonyProvidersEdgesExtensionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesExtensionParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesExtensionParams {
	var ()
	return &GetTelephonyProvidersEdgesExtensionParams{

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesExtensionParamsWithContext creates a new GetTelephonyProvidersEdgesExtensionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesExtensionParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesExtensionParams {
	var ()
	return &GetTelephonyProvidersEdgesExtensionParams{

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesExtensionParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesExtensionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesExtensionParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesExtensionParams {
	var ()
	return &GetTelephonyProvidersEdgesExtensionParams{
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesExtensionParams contains all the parameters to send to the API endpoint
for the get telephony providers edges extension operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesExtensionParams struct {

	/*ExtensionID
	  Extension ID

	*/
	ExtensionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edges extension params
func (o *GetTelephonyProvidersEdgesExtensionParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesExtensionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges extension params
func (o *GetTelephonyProvidersEdgesExtensionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges extension params
func (o *GetTelephonyProvidersEdgesExtensionParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesExtensionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges extension params
func (o *GetTelephonyProvidersEdgesExtensionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges extension params
func (o *GetTelephonyProvidersEdgesExtensionParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesExtensionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges extension params
func (o *GetTelephonyProvidersEdgesExtensionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtensionID adds the extensionID to the get telephony providers edges extension params
func (o *GetTelephonyProvidersEdgesExtensionParams) WithExtensionID(extensionID string) *GetTelephonyProvidersEdgesExtensionParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the get telephony providers edges extension params
func (o *GetTelephonyProvidersEdgesExtensionParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesExtensionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param extensionId
	if err := r.SetPathParam("extensionId", o.ExtensionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
