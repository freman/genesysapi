// Code generated by go-swagger; DO NOT EDIT.

package greetings

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

// NewGetGreetingsDefaultsParams creates a new GetGreetingsDefaultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGreetingsDefaultsParams() *GetGreetingsDefaultsParams {
	return &GetGreetingsDefaultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGreetingsDefaultsParamsWithTimeout creates a new GetGreetingsDefaultsParams object
// with the ability to set a timeout on a request.
func NewGetGreetingsDefaultsParamsWithTimeout(timeout time.Duration) *GetGreetingsDefaultsParams {
	return &GetGreetingsDefaultsParams{
		timeout: timeout,
	}
}

// NewGetGreetingsDefaultsParamsWithContext creates a new GetGreetingsDefaultsParams object
// with the ability to set a context for a request.
func NewGetGreetingsDefaultsParamsWithContext(ctx context.Context) *GetGreetingsDefaultsParams {
	return &GetGreetingsDefaultsParams{
		Context: ctx,
	}
}

// NewGetGreetingsDefaultsParamsWithHTTPClient creates a new GetGreetingsDefaultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGreetingsDefaultsParamsWithHTTPClient(client *http.Client) *GetGreetingsDefaultsParams {
	return &GetGreetingsDefaultsParams{
		HTTPClient: client,
	}
}

/*
GetGreetingsDefaultsParams contains all the parameters to send to the API endpoint

	for the get greetings defaults operation.

	Typically these are written to a http.Request.
*/
type GetGreetingsDefaultsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get greetings defaults params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGreetingsDefaultsParams) WithDefaults() *GetGreetingsDefaultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get greetings defaults params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGreetingsDefaultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get greetings defaults params
func (o *GetGreetingsDefaultsParams) WithTimeout(timeout time.Duration) *GetGreetingsDefaultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get greetings defaults params
func (o *GetGreetingsDefaultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get greetings defaults params
func (o *GetGreetingsDefaultsParams) WithContext(ctx context.Context) *GetGreetingsDefaultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get greetings defaults params
func (o *GetGreetingsDefaultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get greetings defaults params
func (o *GetGreetingsDefaultsParams) WithHTTPClient(client *http.Client) *GetGreetingsDefaultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get greetings defaults params
func (o *GetGreetingsDefaultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetGreetingsDefaultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
