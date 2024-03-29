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

// NewGetIdentityprovidersAdfsParams creates a new GetIdentityprovidersAdfsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIdentityprovidersAdfsParams() *GetIdentityprovidersAdfsParams {
	return &GetIdentityprovidersAdfsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentityprovidersAdfsParamsWithTimeout creates a new GetIdentityprovidersAdfsParams object
// with the ability to set a timeout on a request.
func NewGetIdentityprovidersAdfsParamsWithTimeout(timeout time.Duration) *GetIdentityprovidersAdfsParams {
	return &GetIdentityprovidersAdfsParams{
		timeout: timeout,
	}
}

// NewGetIdentityprovidersAdfsParamsWithContext creates a new GetIdentityprovidersAdfsParams object
// with the ability to set a context for a request.
func NewGetIdentityprovidersAdfsParamsWithContext(ctx context.Context) *GetIdentityprovidersAdfsParams {
	return &GetIdentityprovidersAdfsParams{
		Context: ctx,
	}
}

// NewGetIdentityprovidersAdfsParamsWithHTTPClient creates a new GetIdentityprovidersAdfsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIdentityprovidersAdfsParamsWithHTTPClient(client *http.Client) *GetIdentityprovidersAdfsParams {
	return &GetIdentityprovidersAdfsParams{
		HTTPClient: client,
	}
}

/*
GetIdentityprovidersAdfsParams contains all the parameters to send to the API endpoint

	for the get identityproviders adfs operation.

	Typically these are written to a http.Request.
*/
type GetIdentityprovidersAdfsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get identityproviders adfs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentityprovidersAdfsParams) WithDefaults() *GetIdentityprovidersAdfsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get identityproviders adfs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIdentityprovidersAdfsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get identityproviders adfs params
func (o *GetIdentityprovidersAdfsParams) WithTimeout(timeout time.Duration) *GetIdentityprovidersAdfsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identityproviders adfs params
func (o *GetIdentityprovidersAdfsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identityproviders adfs params
func (o *GetIdentityprovidersAdfsParams) WithContext(ctx context.Context) *GetIdentityprovidersAdfsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identityproviders adfs params
func (o *GetIdentityprovidersAdfsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identityproviders adfs params
func (o *GetIdentityprovidersAdfsParams) WithHTTPClient(client *http.Client) *GetIdentityprovidersAdfsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identityproviders adfs params
func (o *GetIdentityprovidersAdfsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentityprovidersAdfsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
