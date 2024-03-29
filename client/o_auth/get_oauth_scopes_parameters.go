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

// NewGetOauthScopesParams creates a new GetOauthScopesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOauthScopesParams() *GetOauthScopesParams {
	return &GetOauthScopesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOauthScopesParamsWithTimeout creates a new GetOauthScopesParams object
// with the ability to set a timeout on a request.
func NewGetOauthScopesParamsWithTimeout(timeout time.Duration) *GetOauthScopesParams {
	return &GetOauthScopesParams{
		timeout: timeout,
	}
}

// NewGetOauthScopesParamsWithContext creates a new GetOauthScopesParams object
// with the ability to set a context for a request.
func NewGetOauthScopesParamsWithContext(ctx context.Context) *GetOauthScopesParams {
	return &GetOauthScopesParams{
		Context: ctx,
	}
}

// NewGetOauthScopesParamsWithHTTPClient creates a new GetOauthScopesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOauthScopesParamsWithHTTPClient(client *http.Client) *GetOauthScopesParams {
	return &GetOauthScopesParams{
		HTTPClient: client,
	}
}

/*
GetOauthScopesParams contains all the parameters to send to the API endpoint

	for the get oauth scopes operation.

	Typically these are written to a http.Request.
*/
type GetOauthScopesParams struct {

	/* AcceptLanguage.

	   The language with which to display the scope descriptions.

	   Default: "en-us"
	*/
	AcceptLanguage *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get oauth scopes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOauthScopesParams) WithDefaults() *GetOauthScopesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get oauth scopes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOauthScopesParams) SetDefaults() {
	var (
		acceptLanguageDefault = string("en-us")
	)

	val := GetOauthScopesParams{
		AcceptLanguage: &acceptLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get oauth scopes params
func (o *GetOauthScopesParams) WithTimeout(timeout time.Duration) *GetOauthScopesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oauth scopes params
func (o *GetOauthScopesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oauth scopes params
func (o *GetOauthScopesParams) WithContext(ctx context.Context) *GetOauthScopesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oauth scopes params
func (o *GetOauthScopesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get oauth scopes params
func (o *GetOauthScopesParams) WithHTTPClient(client *http.Client) *GetOauthScopesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get oauth scopes params
func (o *GetOauthScopesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get oauth scopes params
func (o *GetOauthScopesParams) WithAcceptLanguage(acceptLanguage *string) *GetOauthScopesParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get oauth scopes params
func (o *GetOauthScopesParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WriteToRequest writes these params to a swagger request
func (o *GetOauthScopesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
