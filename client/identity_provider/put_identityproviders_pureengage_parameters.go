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

	"github.com/freman/genesysapi/models"
)

// NewPutIdentityprovidersPureengageParams creates a new PutIdentityprovidersPureengageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutIdentityprovidersPureengageParams() *PutIdentityprovidersPureengageParams {
	return &PutIdentityprovidersPureengageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutIdentityprovidersPureengageParamsWithTimeout creates a new PutIdentityprovidersPureengageParams object
// with the ability to set a timeout on a request.
func NewPutIdentityprovidersPureengageParamsWithTimeout(timeout time.Duration) *PutIdentityprovidersPureengageParams {
	return &PutIdentityprovidersPureengageParams{
		timeout: timeout,
	}
}

// NewPutIdentityprovidersPureengageParamsWithContext creates a new PutIdentityprovidersPureengageParams object
// with the ability to set a context for a request.
func NewPutIdentityprovidersPureengageParamsWithContext(ctx context.Context) *PutIdentityprovidersPureengageParams {
	return &PutIdentityprovidersPureengageParams{
		Context: ctx,
	}
}

// NewPutIdentityprovidersPureengageParamsWithHTTPClient creates a new PutIdentityprovidersPureengageParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutIdentityprovidersPureengageParamsWithHTTPClient(client *http.Client) *PutIdentityprovidersPureengageParams {
	return &PutIdentityprovidersPureengageParams{
		HTTPClient: client,
	}
}

/*
PutIdentityprovidersPureengageParams contains all the parameters to send to the API endpoint

	for the put identityproviders pureengage operation.

	Typically these are written to a http.Request.
*/
type PutIdentityprovidersPureengageParams struct {

	/* Body.

	   Provider
	*/
	Body *models.PureEngage

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put identityproviders pureengage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutIdentityprovidersPureengageParams) WithDefaults() *PutIdentityprovidersPureengageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put identityproviders pureengage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutIdentityprovidersPureengageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put identityproviders pureengage params
func (o *PutIdentityprovidersPureengageParams) WithTimeout(timeout time.Duration) *PutIdentityprovidersPureengageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put identityproviders pureengage params
func (o *PutIdentityprovidersPureengageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put identityproviders pureengage params
func (o *PutIdentityprovidersPureengageParams) WithContext(ctx context.Context) *PutIdentityprovidersPureengageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put identityproviders pureengage params
func (o *PutIdentityprovidersPureengageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put identityproviders pureengage params
func (o *PutIdentityprovidersPureengageParams) WithHTTPClient(client *http.Client) *PutIdentityprovidersPureengageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put identityproviders pureengage params
func (o *PutIdentityprovidersPureengageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put identityproviders pureengage params
func (o *PutIdentityprovidersPureengageParams) WithBody(body *models.PureEngage) *PutIdentityprovidersPureengageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put identityproviders pureengage params
func (o *PutIdentityprovidersPureengageParams) SetBody(body *models.PureEngage) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutIdentityprovidersPureengageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
