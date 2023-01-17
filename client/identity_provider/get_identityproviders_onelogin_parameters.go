// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

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

// NewGetIdentityprovidersOneloginParams creates a new GetIdentityprovidersOneloginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIdentityprovidersOneloginParams() *GetIdentityprovidersOneloginParams {
	return &GetIdentityprovidersOneloginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentityprovidersOneloginParamsWithTimeout creates a new GetIdentityprovidersOneloginParams object
// with the ability to set a timeout on a request.
func NewGetIdentityprovidersOneloginParamsWithTimeout(timeout time.Duration) *GetIdentityprovidersOneloginParams {
	return &GetIdentityprovidersOneloginParams{
		timeout: timeout,
	}
}

// NewGetIdentityprovidersOneloginParamsWithContext creates a new GetIdentityprovidersOneloginParams object
// with the ability to set a context for a request.
func NewGetIdentityprovidersOneloginParamsWithContext(ctx context.Context) *GetIdentityprovidersOneloginParams {
	return &GetIdentityprovidersOneloginParams{
		Context: ctx,
	}
}

// NewGetIdentityprovidersOneloginParamsWithHTTPClient creates a new GetIdentityprovidersOneloginParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIdentityprovidersOneloginParamsWithHTTPClient(client *http.Client) *GetIdentityprovidersOneloginParams {
	return &GetIdentityprovidersOneloginParams{
		HTTPClient: client,
	}
}

/*
GetIdentityprovidersOneloginParams contains all the parameters to send to the API endpoint

	for the get identityproviders onelogin operation.

	Typically these are written to a http.Request.
*/
type GetIdentityprovidersOneloginParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get identityproviders onelogin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentityprovidersOneloginParams) WithDefaults() *GetIdentityprovidersOneloginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get identityproviders onelogin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentityprovidersOneloginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get identityproviders onelogin params
func (o *GetIdentityprovidersOneloginParams) WithTimeout(timeout time.Duration) *GetIdentityprovidersOneloginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identityproviders onelogin params
func (o *GetIdentityprovidersOneloginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identityproviders onelogin params
func (o *GetIdentityprovidersOneloginParams) WithContext(ctx context.Context) *GetIdentityprovidersOneloginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identityproviders onelogin params
func (o *GetIdentityprovidersOneloginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identityproviders onelogin params
func (o *GetIdentityprovidersOneloginParams) WithHTTPClient(client *http.Client) *GetIdentityprovidersOneloginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identityproviders onelogin params
func (o *GetIdentityprovidersOneloginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentityprovidersOneloginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
