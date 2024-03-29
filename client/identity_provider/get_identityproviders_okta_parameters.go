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

// NewGetIdentityprovidersOktaParams creates a new GetIdentityprovidersOktaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIdentityprovidersOktaParams() *GetIdentityprovidersOktaParams {
	return &GetIdentityprovidersOktaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentityprovidersOktaParamsWithTimeout creates a new GetIdentityprovidersOktaParams object
// with the ability to set a timeout on a request.
func NewGetIdentityprovidersOktaParamsWithTimeout(timeout time.Duration) *GetIdentityprovidersOktaParams {
	return &GetIdentityprovidersOktaParams{
		timeout: timeout,
	}
}

// NewGetIdentityprovidersOktaParamsWithContext creates a new GetIdentityprovidersOktaParams object
// with the ability to set a context for a request.
func NewGetIdentityprovidersOktaParamsWithContext(ctx context.Context) *GetIdentityprovidersOktaParams {
	return &GetIdentityprovidersOktaParams{
		Context: ctx,
	}
}

// NewGetIdentityprovidersOktaParamsWithHTTPClient creates a new GetIdentityprovidersOktaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIdentityprovidersOktaParamsWithHTTPClient(client *http.Client) *GetIdentityprovidersOktaParams {
	return &GetIdentityprovidersOktaParams{
		HTTPClient: client,
	}
}

/*
GetIdentityprovidersOktaParams contains all the parameters to send to the API endpoint

	for the get identityproviders okta operation.

	Typically these are written to a http.Request.
*/
type GetIdentityprovidersOktaParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get identityproviders okta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentityprovidersOktaParams) WithDefaults() *GetIdentityprovidersOktaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get identityproviders okta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentityprovidersOktaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get identityproviders okta params
func (o *GetIdentityprovidersOktaParams) WithTimeout(timeout time.Duration) *GetIdentityprovidersOktaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identityproviders okta params
func (o *GetIdentityprovidersOktaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identityproviders okta params
func (o *GetIdentityprovidersOktaParams) WithContext(ctx context.Context) *GetIdentityprovidersOktaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identityproviders okta params
func (o *GetIdentityprovidersOktaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identityproviders okta params
func (o *GetIdentityprovidersOktaParams) WithHTTPClient(client *http.Client) *GetIdentityprovidersOktaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identityproviders okta params
func (o *GetIdentityprovidersOktaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentityprovidersOktaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
