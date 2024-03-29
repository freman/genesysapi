// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

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

// NewGetScimSchemasParams creates a new GetScimSchemasParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScimSchemasParams() *GetScimSchemasParams {
	return &GetScimSchemasParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScimSchemasParamsWithTimeout creates a new GetScimSchemasParams object
// with the ability to set a timeout on a request.
func NewGetScimSchemasParamsWithTimeout(timeout time.Duration) *GetScimSchemasParams {
	return &GetScimSchemasParams{
		timeout: timeout,
	}
}

// NewGetScimSchemasParamsWithContext creates a new GetScimSchemasParams object
// with the ability to set a context for a request.
func NewGetScimSchemasParamsWithContext(ctx context.Context) *GetScimSchemasParams {
	return &GetScimSchemasParams{
		Context: ctx,
	}
}

// NewGetScimSchemasParamsWithHTTPClient creates a new GetScimSchemasParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScimSchemasParamsWithHTTPClient(client *http.Client) *GetScimSchemasParams {
	return &GetScimSchemasParams{
		HTTPClient: client,
	}
}

/*
GetScimSchemasParams contains all the parameters to send to the API endpoint

	for the get scim schemas operation.

	Typically these are written to a http.Request.
*/
type GetScimSchemasParams struct {

	/* Filter.

	   Filtered results are invalid and return 403 Unauthorized.
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get scim schemas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScimSchemasParams) WithDefaults() *GetScimSchemasParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scim schemas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScimSchemasParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get scim schemas params
func (o *GetScimSchemasParams) WithTimeout(timeout time.Duration) *GetScimSchemasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scim schemas params
func (o *GetScimSchemasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scim schemas params
func (o *GetScimSchemasParams) WithContext(ctx context.Context) *GetScimSchemasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scim schemas params
func (o *GetScimSchemasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scim schemas params
func (o *GetScimSchemasParams) WithHTTPClient(client *http.Client) *GetScimSchemasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scim schemas params
func (o *GetScimSchemasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get scim schemas params
func (o *GetScimSchemasParams) WithFilter(filter *string) *GetScimSchemasParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get scim schemas params
func (o *GetScimSchemasParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *GetScimSchemasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
