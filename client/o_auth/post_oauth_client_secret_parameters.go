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

// NewPostOauthClientSecretParams creates a new PostOauthClientSecretParams object
// with the default values initialized.
func NewPostOauthClientSecretParams() *PostOauthClientSecretParams {
	var ()
	return &PostOauthClientSecretParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOauthClientSecretParamsWithTimeout creates a new PostOauthClientSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOauthClientSecretParamsWithTimeout(timeout time.Duration) *PostOauthClientSecretParams {
	var ()
	return &PostOauthClientSecretParams{

		timeout: timeout,
	}
}

// NewPostOauthClientSecretParamsWithContext creates a new PostOauthClientSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOauthClientSecretParamsWithContext(ctx context.Context) *PostOauthClientSecretParams {
	var ()
	return &PostOauthClientSecretParams{

		Context: ctx,
	}
}

// NewPostOauthClientSecretParamsWithHTTPClient creates a new PostOauthClientSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOauthClientSecretParamsWithHTTPClient(client *http.Client) *PostOauthClientSecretParams {
	var ()
	return &PostOauthClientSecretParams{
		HTTPClient: client,
	}
}

/*PostOauthClientSecretParams contains all the parameters to send to the API endpoint
for the post oauth client secret operation typically these are written to a http.Request
*/
type PostOauthClientSecretParams struct {

	/*ClientID
	  Client ID

	*/
	ClientID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post oauth client secret params
func (o *PostOauthClientSecretParams) WithTimeout(timeout time.Duration) *PostOauthClientSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post oauth client secret params
func (o *PostOauthClientSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post oauth client secret params
func (o *PostOauthClientSecretParams) WithContext(ctx context.Context) *PostOauthClientSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post oauth client secret params
func (o *PostOauthClientSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post oauth client secret params
func (o *PostOauthClientSecretParams) WithHTTPClient(client *http.Client) *PostOauthClientSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post oauth client secret params
func (o *PostOauthClientSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the post oauth client secret params
func (o *PostOauthClientSecretParams) WithClientID(clientID string) *PostOauthClientSecretParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the post oauth client secret params
func (o *PostOauthClientSecretParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *PostOauthClientSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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