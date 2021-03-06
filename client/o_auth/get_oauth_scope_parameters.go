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

// NewGetOauthScopeParams creates a new GetOauthScopeParams object
// with the default values initialized.
func NewGetOauthScopeParams() *GetOauthScopeParams {
	var (
		acceptLanguageDefault = string("en-us")
	)
	return &GetOauthScopeParams{
		AcceptLanguage: &acceptLanguageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOauthScopeParamsWithTimeout creates a new GetOauthScopeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOauthScopeParamsWithTimeout(timeout time.Duration) *GetOauthScopeParams {
	var (
		acceptLanguageDefault = string("en-us")
	)
	return &GetOauthScopeParams{
		AcceptLanguage: &acceptLanguageDefault,

		timeout: timeout,
	}
}

// NewGetOauthScopeParamsWithContext creates a new GetOauthScopeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOauthScopeParamsWithContext(ctx context.Context) *GetOauthScopeParams {
	var (
		acceptLanguageDefault = string("en-us")
	)
	return &GetOauthScopeParams{
		AcceptLanguage: &acceptLanguageDefault,

		Context: ctx,
	}
}

// NewGetOauthScopeParamsWithHTTPClient creates a new GetOauthScopeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOauthScopeParamsWithHTTPClient(client *http.Client) *GetOauthScopeParams {
	var (
		acceptLanguageDefault = string("en-us")
	)
	return &GetOauthScopeParams{
		AcceptLanguage: &acceptLanguageDefault,
		HTTPClient:     client,
	}
}

/*GetOauthScopeParams contains all the parameters to send to the API endpoint
for the get oauth scope operation typically these are written to a http.Request
*/
type GetOauthScopeParams struct {

	/*AcceptLanguage
	  The language with which to display the scope description.

	*/
	AcceptLanguage *string
	/*ScopeID
	  Scope ID

	*/
	ScopeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get oauth scope params
func (o *GetOauthScopeParams) WithTimeout(timeout time.Duration) *GetOauthScopeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oauth scope params
func (o *GetOauthScopeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oauth scope params
func (o *GetOauthScopeParams) WithContext(ctx context.Context) *GetOauthScopeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oauth scope params
func (o *GetOauthScopeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get oauth scope params
func (o *GetOauthScopeParams) WithHTTPClient(client *http.Client) *GetOauthScopeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get oauth scope params
func (o *GetOauthScopeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get oauth scope params
func (o *GetOauthScopeParams) WithAcceptLanguage(acceptLanguage *string) *GetOauthScopeParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get oauth scope params
func (o *GetOauthScopeParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithScopeID adds the scopeID to the get oauth scope params
func (o *GetOauthScopeParams) WithScopeID(scopeID string) *GetOauthScopeParams {
	o.SetScopeID(scopeID)
	return o
}

// SetScopeID adds the scopeId to the get oauth scope params
func (o *GetOauthScopeParams) SetScopeID(scopeID string) {
	o.ScopeID = scopeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOauthScopeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// header param Accept-Language
		if err := r.SetHeaderParam("Accept-Language", *o.AcceptLanguage); err != nil {
			return err
		}

	}

	// path param scopeId
	if err := r.SetPathParam("scopeId", o.ScopeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
