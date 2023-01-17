// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// NewGetAuthorizationDivisionsLimitParams creates a new GetAuthorizationDivisionsLimitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAuthorizationDivisionsLimitParams() *GetAuthorizationDivisionsLimitParams {
	return &GetAuthorizationDivisionsLimitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationDivisionsLimitParamsWithTimeout creates a new GetAuthorizationDivisionsLimitParams object
// with the ability to set a timeout on a request.
func NewGetAuthorizationDivisionsLimitParamsWithTimeout(timeout time.Duration) *GetAuthorizationDivisionsLimitParams {
	return &GetAuthorizationDivisionsLimitParams{
		timeout: timeout,
	}
}

// NewGetAuthorizationDivisionsLimitParamsWithContext creates a new GetAuthorizationDivisionsLimitParams object
// with the ability to set a context for a request.
func NewGetAuthorizationDivisionsLimitParamsWithContext(ctx context.Context) *GetAuthorizationDivisionsLimitParams {
	return &GetAuthorizationDivisionsLimitParams{
		Context: ctx,
	}
}

// NewGetAuthorizationDivisionsLimitParamsWithHTTPClient creates a new GetAuthorizationDivisionsLimitParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAuthorizationDivisionsLimitParamsWithHTTPClient(client *http.Client) *GetAuthorizationDivisionsLimitParams {
	return &GetAuthorizationDivisionsLimitParams{
		HTTPClient: client,
	}
}

/*
GetAuthorizationDivisionsLimitParams contains all the parameters to send to the API endpoint

	for the get authorization divisions limit operation.

	Typically these are written to a http.Request.
*/
type GetAuthorizationDivisionsLimitParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get authorization divisions limit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationDivisionsLimitParams) WithDefaults() *GetAuthorizationDivisionsLimitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get authorization divisions limit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationDivisionsLimitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get authorization divisions limit params
func (o *GetAuthorizationDivisionsLimitParams) WithTimeout(timeout time.Duration) *GetAuthorizationDivisionsLimitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization divisions limit params
func (o *GetAuthorizationDivisionsLimitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization divisions limit params
func (o *GetAuthorizationDivisionsLimitParams) WithContext(ctx context.Context) *GetAuthorizationDivisionsLimitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization divisions limit params
func (o *GetAuthorizationDivisionsLimitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization divisions limit params
func (o *GetAuthorizationDivisionsLimitParams) WithHTTPClient(client *http.Client) *GetAuthorizationDivisionsLimitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization divisions limit params
func (o *GetAuthorizationDivisionsLimitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationDivisionsLimitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
