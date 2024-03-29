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

// NewGetAuthorizationDivisionsHomeParams creates a new GetAuthorizationDivisionsHomeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAuthorizationDivisionsHomeParams() *GetAuthorizationDivisionsHomeParams {
	return &GetAuthorizationDivisionsHomeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationDivisionsHomeParamsWithTimeout creates a new GetAuthorizationDivisionsHomeParams object
// with the ability to set a timeout on a request.
func NewGetAuthorizationDivisionsHomeParamsWithTimeout(timeout time.Duration) *GetAuthorizationDivisionsHomeParams {
	return &GetAuthorizationDivisionsHomeParams{
		timeout: timeout,
	}
}

// NewGetAuthorizationDivisionsHomeParamsWithContext creates a new GetAuthorizationDivisionsHomeParams object
// with the ability to set a context for a request.
func NewGetAuthorizationDivisionsHomeParamsWithContext(ctx context.Context) *GetAuthorizationDivisionsHomeParams {
	return &GetAuthorizationDivisionsHomeParams{
		Context: ctx,
	}
}

// NewGetAuthorizationDivisionsHomeParamsWithHTTPClient creates a new GetAuthorizationDivisionsHomeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAuthorizationDivisionsHomeParamsWithHTTPClient(client *http.Client) *GetAuthorizationDivisionsHomeParams {
	return &GetAuthorizationDivisionsHomeParams{
		HTTPClient: client,
	}
}

/*
GetAuthorizationDivisionsHomeParams contains all the parameters to send to the API endpoint

	for the get authorization divisions home operation.

	Typically these are written to a http.Request.
*/
type GetAuthorizationDivisionsHomeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get authorization divisions home params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationDivisionsHomeParams) WithDefaults() *GetAuthorizationDivisionsHomeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get authorization divisions home params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationDivisionsHomeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get authorization divisions home params
func (o *GetAuthorizationDivisionsHomeParams) WithTimeout(timeout time.Duration) *GetAuthorizationDivisionsHomeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization divisions home params
func (o *GetAuthorizationDivisionsHomeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization divisions home params
func (o *GetAuthorizationDivisionsHomeParams) WithContext(ctx context.Context) *GetAuthorizationDivisionsHomeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization divisions home params
func (o *GetAuthorizationDivisionsHomeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization divisions home params
func (o *GetAuthorizationDivisionsHomeParams) WithHTTPClient(client *http.Client) *GetAuthorizationDivisionsHomeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization divisions home params
func (o *GetAuthorizationDivisionsHomeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationDivisionsHomeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
