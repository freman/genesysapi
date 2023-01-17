// Code generated by go-swagger; DO NOT EDIT.

package data_extensions

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

// NewGetDataextensionsCoretypeParams creates a new GetDataextensionsCoretypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDataextensionsCoretypeParams() *GetDataextensionsCoretypeParams {
	return &GetDataextensionsCoretypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDataextensionsCoretypeParamsWithTimeout creates a new GetDataextensionsCoretypeParams object
// with the ability to set a timeout on a request.
func NewGetDataextensionsCoretypeParamsWithTimeout(timeout time.Duration) *GetDataextensionsCoretypeParams {
	return &GetDataextensionsCoretypeParams{
		timeout: timeout,
	}
}

// NewGetDataextensionsCoretypeParamsWithContext creates a new GetDataextensionsCoretypeParams object
// with the ability to set a context for a request.
func NewGetDataextensionsCoretypeParamsWithContext(ctx context.Context) *GetDataextensionsCoretypeParams {
	return &GetDataextensionsCoretypeParams{
		Context: ctx,
	}
}

// NewGetDataextensionsCoretypeParamsWithHTTPClient creates a new GetDataextensionsCoretypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDataextensionsCoretypeParamsWithHTTPClient(client *http.Client) *GetDataextensionsCoretypeParams {
	return &GetDataextensionsCoretypeParams{
		HTTPClient: client,
	}
}

/*
GetDataextensionsCoretypeParams contains all the parameters to send to the API endpoint

	for the get dataextensions coretype operation.

	Typically these are written to a http.Request.
*/
type GetDataextensionsCoretypeParams struct {

	/* CoretypeName.

	   The core type's name
	*/
	CoretypeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dataextensions coretype params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataextensionsCoretypeParams) WithDefaults() *GetDataextensionsCoretypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dataextensions coretype params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataextensionsCoretypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dataextensions coretype params
func (o *GetDataextensionsCoretypeParams) WithTimeout(timeout time.Duration) *GetDataextensionsCoretypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dataextensions coretype params
func (o *GetDataextensionsCoretypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dataextensions coretype params
func (o *GetDataextensionsCoretypeParams) WithContext(ctx context.Context) *GetDataextensionsCoretypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dataextensions coretype params
func (o *GetDataextensionsCoretypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dataextensions coretype params
func (o *GetDataextensionsCoretypeParams) WithHTTPClient(client *http.Client) *GetDataextensionsCoretypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dataextensions coretype params
func (o *GetDataextensionsCoretypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCoretypeName adds the coretypeName to the get dataextensions coretype params
func (o *GetDataextensionsCoretypeParams) WithCoretypeName(coretypeName string) *GetDataextensionsCoretypeParams {
	o.SetCoretypeName(coretypeName)
	return o
}

// SetCoretypeName adds the coretypeName to the get dataextensions coretype params
func (o *GetDataextensionsCoretypeParams) SetCoretypeName(coretypeName string) {
	o.CoretypeName = coretypeName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDataextensionsCoretypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param coretypeName
	if err := r.SetPathParam("coretypeName", o.CoretypeName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
