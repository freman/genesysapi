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

// NewGetIdentityprovidersGsuiteParams creates a new GetIdentityprovidersGsuiteParams object
// with the default values initialized.
func NewGetIdentityprovidersGsuiteParams() *GetIdentityprovidersGsuiteParams {

	return &GetIdentityprovidersGsuiteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentityprovidersGsuiteParamsWithTimeout creates a new GetIdentityprovidersGsuiteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIdentityprovidersGsuiteParamsWithTimeout(timeout time.Duration) *GetIdentityprovidersGsuiteParams {

	return &GetIdentityprovidersGsuiteParams{

		timeout: timeout,
	}
}

// NewGetIdentityprovidersGsuiteParamsWithContext creates a new GetIdentityprovidersGsuiteParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIdentityprovidersGsuiteParamsWithContext(ctx context.Context) *GetIdentityprovidersGsuiteParams {

	return &GetIdentityprovidersGsuiteParams{

		Context: ctx,
	}
}

// NewGetIdentityprovidersGsuiteParamsWithHTTPClient creates a new GetIdentityprovidersGsuiteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIdentityprovidersGsuiteParamsWithHTTPClient(client *http.Client) *GetIdentityprovidersGsuiteParams {

	return &GetIdentityprovidersGsuiteParams{
		HTTPClient: client,
	}
}

/*GetIdentityprovidersGsuiteParams contains all the parameters to send to the API endpoint
for the get identityproviders gsuite operation typically these are written to a http.Request
*/
type GetIdentityprovidersGsuiteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get identityproviders gsuite params
func (o *GetIdentityprovidersGsuiteParams) WithTimeout(timeout time.Duration) *GetIdentityprovidersGsuiteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identityproviders gsuite params
func (o *GetIdentityprovidersGsuiteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identityproviders gsuite params
func (o *GetIdentityprovidersGsuiteParams) WithContext(ctx context.Context) *GetIdentityprovidersGsuiteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identityproviders gsuite params
func (o *GetIdentityprovidersGsuiteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identityproviders gsuite params
func (o *GetIdentityprovidersGsuiteParams) WithHTTPClient(client *http.Client) *GetIdentityprovidersGsuiteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identityproviders gsuite params
func (o *GetIdentityprovidersGsuiteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentityprovidersGsuiteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
