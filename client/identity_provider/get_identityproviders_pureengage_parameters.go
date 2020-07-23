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

// NewGetIdentityprovidersPureengageParams creates a new GetIdentityprovidersPureengageParams object
// with the default values initialized.
func NewGetIdentityprovidersPureengageParams() *GetIdentityprovidersPureengageParams {

	return &GetIdentityprovidersPureengageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentityprovidersPureengageParamsWithTimeout creates a new GetIdentityprovidersPureengageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIdentityprovidersPureengageParamsWithTimeout(timeout time.Duration) *GetIdentityprovidersPureengageParams {

	return &GetIdentityprovidersPureengageParams{

		timeout: timeout,
	}
}

// NewGetIdentityprovidersPureengageParamsWithContext creates a new GetIdentityprovidersPureengageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIdentityprovidersPureengageParamsWithContext(ctx context.Context) *GetIdentityprovidersPureengageParams {

	return &GetIdentityprovidersPureengageParams{

		Context: ctx,
	}
}

// NewGetIdentityprovidersPureengageParamsWithHTTPClient creates a new GetIdentityprovidersPureengageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIdentityprovidersPureengageParamsWithHTTPClient(client *http.Client) *GetIdentityprovidersPureengageParams {

	return &GetIdentityprovidersPureengageParams{
		HTTPClient: client,
	}
}

/*GetIdentityprovidersPureengageParams contains all the parameters to send to the API endpoint
for the get identityproviders pureengage operation typically these are written to a http.Request
*/
type GetIdentityprovidersPureengageParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get identityproviders pureengage params
func (o *GetIdentityprovidersPureengageParams) WithTimeout(timeout time.Duration) *GetIdentityprovidersPureengageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identityproviders pureengage params
func (o *GetIdentityprovidersPureengageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identityproviders pureengage params
func (o *GetIdentityprovidersPureengageParams) WithContext(ctx context.Context) *GetIdentityprovidersPureengageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identityproviders pureengage params
func (o *GetIdentityprovidersPureengageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identityproviders pureengage params
func (o *GetIdentityprovidersPureengageParams) WithHTTPClient(client *http.Client) *GetIdentityprovidersPureengageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identityproviders pureengage params
func (o *GetIdentityprovidersPureengageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentityprovidersPureengageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
