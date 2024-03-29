// Code generated by go-swagger; DO NOT EDIT.

package tokens

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

// NewHeadTokensMeParams creates a new HeadTokensMeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadTokensMeParams() *HeadTokensMeParams {
	return &HeadTokensMeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadTokensMeParamsWithTimeout creates a new HeadTokensMeParams object
// with the ability to set a timeout on a request.
func NewHeadTokensMeParamsWithTimeout(timeout time.Duration) *HeadTokensMeParams {
	return &HeadTokensMeParams{
		timeout: timeout,
	}
}

// NewHeadTokensMeParamsWithContext creates a new HeadTokensMeParams object
// with the ability to set a context for a request.
func NewHeadTokensMeParamsWithContext(ctx context.Context) *HeadTokensMeParams {
	return &HeadTokensMeParams{
		Context: ctx,
	}
}

// NewHeadTokensMeParamsWithHTTPClient creates a new HeadTokensMeParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadTokensMeParamsWithHTTPClient(client *http.Client) *HeadTokensMeParams {
	return &HeadTokensMeParams{
		HTTPClient: client,
	}
}

/*
HeadTokensMeParams contains all the parameters to send to the API endpoint

	for the head tokens me operation.

	Typically these are written to a http.Request.
*/
type HeadTokensMeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the head tokens me params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadTokensMeParams) WithDefaults() *HeadTokensMeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the head tokens me params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadTokensMeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the head tokens me params
func (o *HeadTokensMeParams) WithTimeout(timeout time.Duration) *HeadTokensMeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head tokens me params
func (o *HeadTokensMeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head tokens me params
func (o *HeadTokensMeParams) WithContext(ctx context.Context) *HeadTokensMeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head tokens me params
func (o *HeadTokensMeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head tokens me params
func (o *HeadTokensMeParams) WithHTTPClient(client *http.Client) *HeadTokensMeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head tokens me params
func (o *HeadTokensMeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HeadTokensMeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
