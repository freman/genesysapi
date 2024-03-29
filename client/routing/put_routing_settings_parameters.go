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

// NewPutRoutingSettingsParams creates a new PutRoutingSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRoutingSettingsParams() *PutRoutingSettingsParams {
	return &PutRoutingSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRoutingSettingsParamsWithTimeout creates a new PutRoutingSettingsParams object
// with the ability to set a timeout on a request.
func NewPutRoutingSettingsParamsWithTimeout(timeout time.Duration) *PutRoutingSettingsParams {
	return &PutRoutingSettingsParams{
		timeout: timeout,
	}
}

// NewPutRoutingSettingsParamsWithContext creates a new PutRoutingSettingsParams object
// with the ability to set a context for a request.
func NewPutRoutingSettingsParamsWithContext(ctx context.Context) *PutRoutingSettingsParams {
	return &PutRoutingSettingsParams{
		Context: ctx,
	}
}

// NewPutRoutingSettingsParamsWithHTTPClient creates a new PutRoutingSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRoutingSettingsParamsWithHTTPClient(client *http.Client) *PutRoutingSettingsParams {
	return &PutRoutingSettingsParams{
		HTTPClient: client,
	}
}

/*
PutRoutingSettingsParams contains all the parameters to send to the API endpoint

	for the put routing settings operation.

	Typically these are written to a http.Request.
*/
type PutRoutingSettingsParams struct {

	/* Body.

	   Organization Settings
	*/
	Body *models.RoutingSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put routing settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRoutingSettingsParams) WithDefaults() *PutRoutingSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put routing settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRoutingSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put routing settings params
func (o *PutRoutingSettingsParams) WithTimeout(timeout time.Duration) *PutRoutingSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put routing settings params
func (o *PutRoutingSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put routing settings params
func (o *PutRoutingSettingsParams) WithContext(ctx context.Context) *PutRoutingSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put routing settings params
func (o *PutRoutingSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put routing settings params
func (o *PutRoutingSettingsParams) WithHTTPClient(client *http.Client) *PutRoutingSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put routing settings params
func (o *PutRoutingSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put routing settings params
func (o *PutRoutingSettingsParams) WithBody(body *models.RoutingSettings) *PutRoutingSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put routing settings params
func (o *PutRoutingSettingsParams) SetBody(body *models.RoutingSettings) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutRoutingSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
