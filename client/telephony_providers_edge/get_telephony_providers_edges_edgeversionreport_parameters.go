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

// NewGetTelephonyProvidersEdgesEdgeversionreportParams creates a new GetTelephonyProvidersEdgesEdgeversionreportParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesEdgeversionreportParams() *GetTelephonyProvidersEdgesEdgeversionreportParams {

	return &GetTelephonyProvidersEdgesEdgeversionreportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesEdgeversionreportParamsWithTimeout creates a new GetTelephonyProvidersEdgesEdgeversionreportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesEdgeversionreportParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesEdgeversionreportParams {

	return &GetTelephonyProvidersEdgesEdgeversionreportParams{

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesEdgeversionreportParamsWithContext creates a new GetTelephonyProvidersEdgesEdgeversionreportParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesEdgeversionreportParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesEdgeversionreportParams {

	return &GetTelephonyProvidersEdgesEdgeversionreportParams{

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesEdgeversionreportParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesEdgeversionreportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesEdgeversionreportParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesEdgeversionreportParams {

	return &GetTelephonyProvidersEdgesEdgeversionreportParams{
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesEdgeversionreportParams contains all the parameters to send to the API endpoint
for the get telephony providers edges edgeversionreport operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesEdgeversionreportParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edges edgeversionreport params
func (o *GetTelephonyProvidersEdgesEdgeversionreportParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesEdgeversionreportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges edgeversionreport params
func (o *GetTelephonyProvidersEdgesEdgeversionreportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges edgeversionreport params
func (o *GetTelephonyProvidersEdgesEdgeversionreportParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesEdgeversionreportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges edgeversionreport params
func (o *GetTelephonyProvidersEdgesEdgeversionreportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges edgeversionreport params
func (o *GetTelephonyProvidersEdgesEdgeversionreportParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesEdgeversionreportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges edgeversionreport params
func (o *GetTelephonyProvidersEdgesEdgeversionreportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesEdgeversionreportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}