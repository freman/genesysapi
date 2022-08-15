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

// NewGetRoutingAvailablemediatypesParams creates a new GetRoutingAvailablemediatypesParams object
// with the default values initialized.
func NewGetRoutingAvailablemediatypesParams() *GetRoutingAvailablemediatypesParams {

	return &GetRoutingAvailablemediatypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingAvailablemediatypesParamsWithTimeout creates a new GetRoutingAvailablemediatypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoutingAvailablemediatypesParamsWithTimeout(timeout time.Duration) *GetRoutingAvailablemediatypesParams {

	return &GetRoutingAvailablemediatypesParams{

		timeout: timeout,
	}
}

// NewGetRoutingAvailablemediatypesParamsWithContext creates a new GetRoutingAvailablemediatypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoutingAvailablemediatypesParamsWithContext(ctx context.Context) *GetRoutingAvailablemediatypesParams {

	return &GetRoutingAvailablemediatypesParams{

		Context: ctx,
	}
}

// NewGetRoutingAvailablemediatypesParamsWithHTTPClient creates a new GetRoutingAvailablemediatypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoutingAvailablemediatypesParamsWithHTTPClient(client *http.Client) *GetRoutingAvailablemediatypesParams {

	return &GetRoutingAvailablemediatypesParams{
		HTTPClient: client,
	}
}

/*GetRoutingAvailablemediatypesParams contains all the parameters to send to the API endpoint
for the get routing availablemediatypes operation typically these are written to a http.Request
*/
type GetRoutingAvailablemediatypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routing availablemediatypes params
func (o *GetRoutingAvailablemediatypesParams) WithTimeout(timeout time.Duration) *GetRoutingAvailablemediatypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing availablemediatypes params
func (o *GetRoutingAvailablemediatypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing availablemediatypes params
func (o *GetRoutingAvailablemediatypesParams) WithContext(ctx context.Context) *GetRoutingAvailablemediatypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing availablemediatypes params
func (o *GetRoutingAvailablemediatypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing availablemediatypes params
func (o *GetRoutingAvailablemediatypesParams) WithHTTPClient(client *http.Client) *GetRoutingAvailablemediatypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing availablemediatypes params
func (o *GetRoutingAvailablemediatypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingAvailablemediatypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
