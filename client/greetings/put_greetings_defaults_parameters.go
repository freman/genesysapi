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

	"github.com/freman/genesysapi/models"
)

// NewPutGreetingsDefaultsParams creates a new PutGreetingsDefaultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutGreetingsDefaultsParams() *PutGreetingsDefaultsParams {
	return &PutGreetingsDefaultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutGreetingsDefaultsParamsWithTimeout creates a new PutGreetingsDefaultsParams object
// with the ability to set a timeout on a request.
func NewPutGreetingsDefaultsParamsWithTimeout(timeout time.Duration) *PutGreetingsDefaultsParams {
	return &PutGreetingsDefaultsParams{
		timeout: timeout,
	}
}

// NewPutGreetingsDefaultsParamsWithContext creates a new PutGreetingsDefaultsParams object
// with the ability to set a context for a request.
func NewPutGreetingsDefaultsParamsWithContext(ctx context.Context) *PutGreetingsDefaultsParams {
	return &PutGreetingsDefaultsParams{
		Context: ctx,
	}
}

// NewPutGreetingsDefaultsParamsWithHTTPClient creates a new PutGreetingsDefaultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutGreetingsDefaultsParamsWithHTTPClient(client *http.Client) *PutGreetingsDefaultsParams {
	return &PutGreetingsDefaultsParams{
		HTTPClient: client,
	}
}

/*
PutGreetingsDefaultsParams contains all the parameters to send to the API endpoint

	for the put greetings defaults operation.

	Typically these are written to a http.Request.
*/
type PutGreetingsDefaultsParams struct {

	/* Body.

	   The updated defaultGreetingList
	*/
	Body *models.DefaultGreetingList

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put greetings defaults params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutGreetingsDefaultsParams) WithDefaults() *PutGreetingsDefaultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put greetings defaults params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutGreetingsDefaultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put greetings defaults params
func (o *PutGreetingsDefaultsParams) WithTimeout(timeout time.Duration) *PutGreetingsDefaultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put greetings defaults params
func (o *PutGreetingsDefaultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put greetings defaults params
func (o *PutGreetingsDefaultsParams) WithContext(ctx context.Context) *PutGreetingsDefaultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put greetings defaults params
func (o *PutGreetingsDefaultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put greetings defaults params
func (o *PutGreetingsDefaultsParams) WithHTTPClient(client *http.Client) *PutGreetingsDefaultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put greetings defaults params
func (o *PutGreetingsDefaultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put greetings defaults params
func (o *PutGreetingsDefaultsParams) WithBody(body *models.DefaultGreetingList) *PutGreetingsDefaultsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put greetings defaults params
func (o *PutGreetingsDefaultsParams) SetBody(body *models.DefaultGreetingList) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutGreetingsDefaultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
