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

// NewGetIntegrationsCredentialsTypesParams creates a new GetIntegrationsCredentialsTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIntegrationsCredentialsTypesParams() *GetIntegrationsCredentialsTypesParams {
	return &GetIntegrationsCredentialsTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationsCredentialsTypesParamsWithTimeout creates a new GetIntegrationsCredentialsTypesParams object
// with the ability to set a timeout on a request.
func NewGetIntegrationsCredentialsTypesParamsWithTimeout(timeout time.Duration) *GetIntegrationsCredentialsTypesParams {
	return &GetIntegrationsCredentialsTypesParams{
		timeout: timeout,
	}
}

// NewGetIntegrationsCredentialsTypesParamsWithContext creates a new GetIntegrationsCredentialsTypesParams object
// with the ability to set a context for a request.
func NewGetIntegrationsCredentialsTypesParamsWithContext(ctx context.Context) *GetIntegrationsCredentialsTypesParams {
	return &GetIntegrationsCredentialsTypesParams{
		Context: ctx,
	}
}

// NewGetIntegrationsCredentialsTypesParamsWithHTTPClient creates a new GetIntegrationsCredentialsTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIntegrationsCredentialsTypesParamsWithHTTPClient(client *http.Client) *GetIntegrationsCredentialsTypesParams {
	return &GetIntegrationsCredentialsTypesParams{
		HTTPClient: client,
	}
}

/*
GetIntegrationsCredentialsTypesParams contains all the parameters to send to the API endpoint

	for the get integrations credentials types operation.

	Typically these are written to a http.Request.
*/
type GetIntegrationsCredentialsTypesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get integrations credentials types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationsCredentialsTypesParams) WithDefaults() *GetIntegrationsCredentialsTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get integrations credentials types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationsCredentialsTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get integrations credentials types params
func (o *GetIntegrationsCredentialsTypesParams) WithTimeout(timeout time.Duration) *GetIntegrationsCredentialsTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrations credentials types params
func (o *GetIntegrationsCredentialsTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrations credentials types params
func (o *GetIntegrationsCredentialsTypesParams) WithContext(ctx context.Context) *GetIntegrationsCredentialsTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrations credentials types params
func (o *GetIntegrationsCredentialsTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrations credentials types params
func (o *GetIntegrationsCredentialsTypesParams) WithHTTPClient(client *http.Client) *GetIntegrationsCredentialsTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrations credentials types params
func (o *GetIntegrationsCredentialsTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationsCredentialsTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
