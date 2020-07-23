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

// NewGetOauthAuthorizationParams creates a new GetOauthAuthorizationParams object
// with the default values initialized.
func NewGetOauthAuthorizationParams() *GetOauthAuthorizationParams {
	var ()
	return &GetOauthAuthorizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOauthAuthorizationParamsWithTimeout creates a new GetOauthAuthorizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOauthAuthorizationParamsWithTimeout(timeout time.Duration) *GetOauthAuthorizationParams {
	var ()
	return &GetOauthAuthorizationParams{

		timeout: timeout,
	}
}

// NewGetOauthAuthorizationParamsWithContext creates a new GetOauthAuthorizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOauthAuthorizationParamsWithContext(ctx context.Context) *GetOauthAuthorizationParams {
	var ()
	return &GetOauthAuthorizationParams{

		Context: ctx,
	}
}

// NewGetOauthAuthorizationParamsWithHTTPClient creates a new GetOauthAuthorizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOauthAuthorizationParamsWithHTTPClient(client *http.Client) *GetOauthAuthorizationParams {
	var ()
	return &GetOauthAuthorizationParams{
		HTTPClient: client,
	}
}

/*GetOauthAuthorizationParams contains all the parameters to send to the API endpoint
for the get oauth authorization operation typically these are written to a http.Request
*/
type GetOauthAuthorizationParams struct {

	/*ClientID
	  The ID of client

	*/
	ClientID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get oauth authorization params
func (o *GetOauthAuthorizationParams) WithTimeout(timeout time.Duration) *GetOauthAuthorizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oauth authorization params
func (o *GetOauthAuthorizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oauth authorization params
func (o *GetOauthAuthorizationParams) WithContext(ctx context.Context) *GetOauthAuthorizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oauth authorization params
func (o *GetOauthAuthorizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get oauth authorization params
func (o *GetOauthAuthorizationParams) WithHTTPClient(client *http.Client) *GetOauthAuthorizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get oauth authorization params
func (o *GetOauthAuthorizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the get oauth authorization params
func (o *GetOauthAuthorizationParams) WithClientID(clientID string) *GetOauthAuthorizationParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the get oauth authorization params
func (o *GetOauthAuthorizationParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOauthAuthorizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clientId
	if err := r.SetPathParam("clientId", o.ClientID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
