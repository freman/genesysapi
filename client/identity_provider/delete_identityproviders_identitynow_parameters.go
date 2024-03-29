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

// NewDeleteIdentityprovidersIdentitynowParams creates a new DeleteIdentityprovidersIdentitynowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteIdentityprovidersIdentitynowParams() *DeleteIdentityprovidersIdentitynowParams {
	return &DeleteIdentityprovidersIdentitynowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIdentityprovidersIdentitynowParamsWithTimeout creates a new DeleteIdentityprovidersIdentitynowParams object
// with the ability to set a timeout on a request.
func NewDeleteIdentityprovidersIdentitynowParamsWithTimeout(timeout time.Duration) *DeleteIdentityprovidersIdentitynowParams {
	return &DeleteIdentityprovidersIdentitynowParams{
		timeout: timeout,
	}
}

// NewDeleteIdentityprovidersIdentitynowParamsWithContext creates a new DeleteIdentityprovidersIdentitynowParams object
// with the ability to set a context for a request.
func NewDeleteIdentityprovidersIdentitynowParamsWithContext(ctx context.Context) *DeleteIdentityprovidersIdentitynowParams {
	return &DeleteIdentityprovidersIdentitynowParams{
		Context: ctx,
	}
}

// NewDeleteIdentityprovidersIdentitynowParamsWithHTTPClient creates a new DeleteIdentityprovidersIdentitynowParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteIdentityprovidersIdentitynowParamsWithHTTPClient(client *http.Client) *DeleteIdentityprovidersIdentitynowParams {
	return &DeleteIdentityprovidersIdentitynowParams{
		HTTPClient: client,
	}
}

/*
DeleteIdentityprovidersIdentitynowParams contains all the parameters to send to the API endpoint

	for the delete identityproviders identitynow operation.

	Typically these are written to a http.Request.
*/
type DeleteIdentityprovidersIdentitynowParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete identityproviders identitynow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIdentityprovidersIdentitynowParams) WithDefaults() *DeleteIdentityprovidersIdentitynowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete identityproviders identitynow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIdentityprovidersIdentitynowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete identityproviders identitynow params
func (o *DeleteIdentityprovidersIdentitynowParams) WithTimeout(timeout time.Duration) *DeleteIdentityprovidersIdentitynowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete identityproviders identitynow params
func (o *DeleteIdentityprovidersIdentitynowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete identityproviders identitynow params
func (o *DeleteIdentityprovidersIdentitynowParams) WithContext(ctx context.Context) *DeleteIdentityprovidersIdentitynowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete identityproviders identitynow params
func (o *DeleteIdentityprovidersIdentitynowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete identityproviders identitynow params
func (o *DeleteIdentityprovidersIdentitynowParams) WithHTTPClient(client *http.Client) *DeleteIdentityprovidersIdentitynowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete identityproviders identitynow params
func (o *DeleteIdentityprovidersIdentitynowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIdentityprovidersIdentitynowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
