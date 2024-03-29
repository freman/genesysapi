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
)

// NewGetIntegrationConfigCurrentParams creates a new GetIntegrationConfigCurrentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIntegrationConfigCurrentParams() *GetIntegrationConfigCurrentParams {
	return &GetIntegrationConfigCurrentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationConfigCurrentParamsWithTimeout creates a new GetIntegrationConfigCurrentParams object
// with the ability to set a timeout on a request.
func NewGetIntegrationConfigCurrentParamsWithTimeout(timeout time.Duration) *GetIntegrationConfigCurrentParams {
	return &GetIntegrationConfigCurrentParams{
		timeout: timeout,
	}
}

// NewGetIntegrationConfigCurrentParamsWithContext creates a new GetIntegrationConfigCurrentParams object
// with the ability to set a context for a request.
func NewGetIntegrationConfigCurrentParamsWithContext(ctx context.Context) *GetIntegrationConfigCurrentParams {
	return &GetIntegrationConfigCurrentParams{
		Context: ctx,
	}
}

// NewGetIntegrationConfigCurrentParamsWithHTTPClient creates a new GetIntegrationConfigCurrentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIntegrationConfigCurrentParamsWithHTTPClient(client *http.Client) *GetIntegrationConfigCurrentParams {
	return &GetIntegrationConfigCurrentParams{
		HTTPClient: client,
	}
}

/*
GetIntegrationConfigCurrentParams contains all the parameters to send to the API endpoint

	for the get integration config current operation.

	Typically these are written to a http.Request.
*/
type GetIntegrationConfigCurrentParams struct {

	/* IntegrationID.

	   Integration Id
	*/
	IntegrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get integration config current params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationConfigCurrentParams) WithDefaults() *GetIntegrationConfigCurrentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get integration config current params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationConfigCurrentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get integration config current params
func (o *GetIntegrationConfigCurrentParams) WithTimeout(timeout time.Duration) *GetIntegrationConfigCurrentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integration config current params
func (o *GetIntegrationConfigCurrentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integration config current params
func (o *GetIntegrationConfigCurrentParams) WithContext(ctx context.Context) *GetIntegrationConfigCurrentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integration config current params
func (o *GetIntegrationConfigCurrentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integration config current params
func (o *GetIntegrationConfigCurrentParams) WithHTTPClient(client *http.Client) *GetIntegrationConfigCurrentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integration config current params
func (o *GetIntegrationConfigCurrentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntegrationID adds the integrationID to the get integration config current params
func (o *GetIntegrationConfigCurrentParams) WithIntegrationID(integrationID string) *GetIntegrationConfigCurrentParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the get integration config current params
func (o *GetIntegrationConfigCurrentParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationConfigCurrentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param integrationId
	if err := r.SetPathParam("integrationId", o.IntegrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
