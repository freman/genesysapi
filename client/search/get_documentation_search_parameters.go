// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewGetDocumentationSearchParams creates a new GetDocumentationSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDocumentationSearchParams() *GetDocumentationSearchParams {
	return &GetDocumentationSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDocumentationSearchParamsWithTimeout creates a new GetDocumentationSearchParams object
// with the ability to set a timeout on a request.
func NewGetDocumentationSearchParamsWithTimeout(timeout time.Duration) *GetDocumentationSearchParams {
	return &GetDocumentationSearchParams{
		timeout: timeout,
	}
}

// NewGetDocumentationSearchParamsWithContext creates a new GetDocumentationSearchParams object
// with the ability to set a context for a request.
func NewGetDocumentationSearchParamsWithContext(ctx context.Context) *GetDocumentationSearchParams {
	return &GetDocumentationSearchParams{
		Context: ctx,
	}
}

// NewGetDocumentationSearchParamsWithHTTPClient creates a new GetDocumentationSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDocumentationSearchParamsWithHTTPClient(client *http.Client) *GetDocumentationSearchParams {
	return &GetDocumentationSearchParams{
		HTTPClient: client,
	}
}

/*
GetDocumentationSearchParams contains all the parameters to send to the API endpoint

	for the get documentation search operation.

	Typically these are written to a http.Request.
*/
type GetDocumentationSearchParams struct {

	/* Q64.

	   q64
	*/
	Q64 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get documentation search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDocumentationSearchParams) WithDefaults() *GetDocumentationSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get documentation search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDocumentationSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get documentation search params
func (o *GetDocumentationSearchParams) WithTimeout(timeout time.Duration) *GetDocumentationSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get documentation search params
func (o *GetDocumentationSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get documentation search params
func (o *GetDocumentationSearchParams) WithContext(ctx context.Context) *GetDocumentationSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get documentation search params
func (o *GetDocumentationSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get documentation search params
func (o *GetDocumentationSearchParams) WithHTTPClient(client *http.Client) *GetDocumentationSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get documentation search params
func (o *GetDocumentationSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQ64 adds the q64 to the get documentation search params
func (o *GetDocumentationSearchParams) WithQ64(q64 string) *GetDocumentationSearchParams {
	o.SetQ64(q64)
	return o
}

// SetQ64 adds the q64 to the get documentation search params
func (o *GetDocumentationSearchParams) SetQ64(q64 string) {
	o.Q64 = q64
}

// WriteToRequest writes these params to a swagger request
func (o *GetDocumentationSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param q64
	qrQ64 := o.Q64
	qQ64 := qrQ64
	if qQ64 != "" {

		if err := r.SetQueryParam("q64", qQ64); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
