// Code generated by go-swagger; DO NOT EDIT.

package search

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
	"github.com/go-openapi/swag"

	"github.com/freman/genesysapi/models"
)

// NewPostSearchParams creates a new PostSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSearchParams() *PostSearchParams {
	return &PostSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSearchParamsWithTimeout creates a new PostSearchParams object
// with the ability to set a timeout on a request.
func NewPostSearchParamsWithTimeout(timeout time.Duration) *PostSearchParams {
	return &PostSearchParams{
		timeout: timeout,
	}
}

// NewPostSearchParamsWithContext creates a new PostSearchParams object
// with the ability to set a context for a request.
func NewPostSearchParamsWithContext(ctx context.Context) *PostSearchParams {
	return &PostSearchParams{
		Context: ctx,
	}
}

// NewPostSearchParamsWithHTTPClient creates a new PostSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSearchParamsWithHTTPClient(client *http.Client) *PostSearchParams {
	return &PostSearchParams{
		HTTPClient: client,
	}
}

/*
PostSearchParams contains all the parameters to send to the API endpoint

	for the post search operation.

	Typically these are written to a http.Request.
*/
type PostSearchParams struct {

	/* Body.

	   Search request options
	*/
	Body *models.SearchRequest

	/* Profile.

	   profile

	   Default: true
	*/
	Profile *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSearchParams) WithDefaults() *PostSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSearchParams) SetDefaults() {
	var (
		profileDefault = bool(true)
	)

	val := PostSearchParams{
		Profile: &profileDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the post search params
func (o *PostSearchParams) WithTimeout(timeout time.Duration) *PostSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post search params
func (o *PostSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post search params
func (o *PostSearchParams) WithContext(ctx context.Context) *PostSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post search params
func (o *PostSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post search params
func (o *PostSearchParams) WithHTTPClient(client *http.Client) *PostSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post search params
func (o *PostSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post search params
func (o *PostSearchParams) WithBody(body *models.SearchRequest) *PostSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post search params
func (o *PostSearchParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithProfile adds the profile to the post search params
func (o *PostSearchParams) WithProfile(profile *bool) *PostSearchParams {
	o.SetProfile(profile)
	return o
}

// SetProfile adds the profile to the post search params
func (o *PostSearchParams) SetProfile(profile *bool) {
	o.Profile = profile
}

// WriteToRequest writes these params to a swagger request
func (o *PostSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Profile != nil {

		// query param profile
		var qrProfile bool

		if o.Profile != nil {
			qrProfile = *o.Profile
		}
		qProfile := swag.FormatBool(qrProfile)
		if qProfile != "" {

			if err := r.SetQueryParam("profile", qProfile); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
