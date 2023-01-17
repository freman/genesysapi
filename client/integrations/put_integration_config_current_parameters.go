// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewPutIntegrationConfigCurrentParams creates a new PutIntegrationConfigCurrentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutIntegrationConfigCurrentParams() *PutIntegrationConfigCurrentParams {
	return &PutIntegrationConfigCurrentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutIntegrationConfigCurrentParamsWithTimeout creates a new PutIntegrationConfigCurrentParams object
// with the ability to set a timeout on a request.
func NewPutIntegrationConfigCurrentParamsWithTimeout(timeout time.Duration) *PutIntegrationConfigCurrentParams {
	return &PutIntegrationConfigCurrentParams{
		timeout: timeout,
	}
}

// NewPutIntegrationConfigCurrentParamsWithContext creates a new PutIntegrationConfigCurrentParams object
// with the ability to set a context for a request.
func NewPutIntegrationConfigCurrentParamsWithContext(ctx context.Context) *PutIntegrationConfigCurrentParams {
	return &PutIntegrationConfigCurrentParams{
		Context: ctx,
	}
}

// NewPutIntegrationConfigCurrentParamsWithHTTPClient creates a new PutIntegrationConfigCurrentParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutIntegrationConfigCurrentParamsWithHTTPClient(client *http.Client) *PutIntegrationConfigCurrentParams {
	return &PutIntegrationConfigCurrentParams{
		HTTPClient: client,
	}
}

/*
PutIntegrationConfigCurrentParams contains all the parameters to send to the API endpoint

	for the put integration config current operation.

	Typically these are written to a http.Request.
*/
type PutIntegrationConfigCurrentParams struct {

	/* Body.

	   Integration Configuration
	*/
	Body *models.IntegrationConfiguration

	/* IntegrationID.

	   Integration Id
	*/
	IntegrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put integration config current params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutIntegrationConfigCurrentParams) WithDefaults() *PutIntegrationConfigCurrentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put integration config current params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutIntegrationConfigCurrentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put integration config current params
func (o *PutIntegrationConfigCurrentParams) WithTimeout(timeout time.Duration) *PutIntegrationConfigCurrentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put integration config current params
func (o *PutIntegrationConfigCurrentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put integration config current params
func (o *PutIntegrationConfigCurrentParams) WithContext(ctx context.Context) *PutIntegrationConfigCurrentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put integration config current params
func (o *PutIntegrationConfigCurrentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put integration config current params
func (o *PutIntegrationConfigCurrentParams) WithHTTPClient(client *http.Client) *PutIntegrationConfigCurrentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put integration config current params
func (o *PutIntegrationConfigCurrentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put integration config current params
func (o *PutIntegrationConfigCurrentParams) WithBody(body *models.IntegrationConfiguration) *PutIntegrationConfigCurrentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put integration config current params
func (o *PutIntegrationConfigCurrentParams) SetBody(body *models.IntegrationConfiguration) {
	o.Body = body
}

// WithIntegrationID adds the integrationID to the put integration config current params
func (o *PutIntegrationConfigCurrentParams) WithIntegrationID(integrationID string) *PutIntegrationConfigCurrentParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the put integration config current params
func (o *PutIntegrationConfigCurrentParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *PutIntegrationConfigCurrentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param integrationId
	if err := r.SetPathParam("integrationId", o.IntegrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
