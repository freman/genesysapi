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

// NewPostSearchSuggestParams creates a new PostSearchSuggestParams object
// with the default values initialized.
func NewPostSearchSuggestParams() *PostSearchSuggestParams {
	var (
		profileDefault = bool(true)
	)
	return &PostSearchSuggestParams{
		Profile: &profileDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSearchSuggestParamsWithTimeout creates a new PostSearchSuggestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSearchSuggestParamsWithTimeout(timeout time.Duration) *PostSearchSuggestParams {
	var (
		profileDefault = bool(true)
	)
	return &PostSearchSuggestParams{
		Profile: &profileDefault,

		timeout: timeout,
	}
}

// NewPostSearchSuggestParamsWithContext creates a new PostSearchSuggestParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSearchSuggestParamsWithContext(ctx context.Context) *PostSearchSuggestParams {
	var (
		profileDefault = bool(true)
	)
	return &PostSearchSuggestParams{
		Profile: &profileDefault,

		Context: ctx,
	}
}

// NewPostSearchSuggestParamsWithHTTPClient creates a new PostSearchSuggestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSearchSuggestParamsWithHTTPClient(client *http.Client) *PostSearchSuggestParams {
	var (
		profileDefault = bool(true)
	)
	return &PostSearchSuggestParams{
		Profile:    &profileDefault,
		HTTPClient: client,
	}
}

/*PostSearchSuggestParams contains all the parameters to send to the API endpoint
for the post search suggest operation typically these are written to a http.Request
*/
type PostSearchSuggestParams struct {

	/*Body
	  Search request options

	*/
	Body *models.SuggestSearchRequest
	/*Profile
	  profile

	*/
	Profile *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post search suggest params
func (o *PostSearchSuggestParams) WithTimeout(timeout time.Duration) *PostSearchSuggestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post search suggest params
func (o *PostSearchSuggestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post search suggest params
func (o *PostSearchSuggestParams) WithContext(ctx context.Context) *PostSearchSuggestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post search suggest params
func (o *PostSearchSuggestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post search suggest params
func (o *PostSearchSuggestParams) WithHTTPClient(client *http.Client) *PostSearchSuggestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post search suggest params
func (o *PostSearchSuggestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post search suggest params
func (o *PostSearchSuggestParams) WithBody(body *models.SuggestSearchRequest) *PostSearchSuggestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post search suggest params
func (o *PostSearchSuggestParams) SetBody(body *models.SuggestSearchRequest) {
	o.Body = body
}

// WithProfile adds the profile to the post search suggest params
func (o *PostSearchSuggestParams) WithProfile(profile *bool) *PostSearchSuggestParams {
	o.SetProfile(profile)
	return o
}

// SetProfile adds the profile to the post search suggest params
func (o *PostSearchSuggestParams) SetProfile(profile *bool) {
	o.Profile = profile
}

// WriteToRequest writes these params to a swagger request
func (o *PostSearchSuggestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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