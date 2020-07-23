// Code generated by go-swagger; DO NOT EDIT.

package o_auth

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

// NewGetOauthAuthorizationsParams creates a new GetOauthAuthorizationsParams object
// with the default values initialized.
func NewGetOauthAuthorizationsParams() *GetOauthAuthorizationsParams {

	return &GetOauthAuthorizationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOauthAuthorizationsParamsWithTimeout creates a new GetOauthAuthorizationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOauthAuthorizationsParamsWithTimeout(timeout time.Duration) *GetOauthAuthorizationsParams {

	return &GetOauthAuthorizationsParams{

		timeout: timeout,
	}
}

// NewGetOauthAuthorizationsParamsWithContext creates a new GetOauthAuthorizationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOauthAuthorizationsParamsWithContext(ctx context.Context) *GetOauthAuthorizationsParams {

	return &GetOauthAuthorizationsParams{

		Context: ctx,
	}
}

// NewGetOauthAuthorizationsParamsWithHTTPClient creates a new GetOauthAuthorizationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOauthAuthorizationsParamsWithHTTPClient(client *http.Client) *GetOauthAuthorizationsParams {

	return &GetOauthAuthorizationsParams{
		HTTPClient: client,
	}
}

/*GetOauthAuthorizationsParams contains all the parameters to send to the API endpoint
for the get oauth authorizations operation typically these are written to a http.Request
*/
type GetOauthAuthorizationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get oauth authorizations params
func (o *GetOauthAuthorizationsParams) WithTimeout(timeout time.Duration) *GetOauthAuthorizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oauth authorizations params
func (o *GetOauthAuthorizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oauth authorizations params
func (o *GetOauthAuthorizationsParams) WithContext(ctx context.Context) *GetOauthAuthorizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oauth authorizations params
func (o *GetOauthAuthorizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get oauth authorizations params
func (o *GetOauthAuthorizationsParams) WithHTTPClient(client *http.Client) *GetOauthAuthorizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get oauth authorizations params
func (o *GetOauthAuthorizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOauthAuthorizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
