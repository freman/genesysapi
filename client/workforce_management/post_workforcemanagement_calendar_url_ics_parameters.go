// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

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

// NewPostWorkforcemanagementCalendarURLIcsParams creates a new PostWorkforcemanagementCalendarURLIcsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementCalendarURLIcsParams() *PostWorkforcemanagementCalendarURLIcsParams {
	return &PostWorkforcemanagementCalendarURLIcsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementCalendarURLIcsParamsWithTimeout creates a new PostWorkforcemanagementCalendarURLIcsParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementCalendarURLIcsParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementCalendarURLIcsParams {
	return &PostWorkforcemanagementCalendarURLIcsParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementCalendarURLIcsParamsWithContext creates a new PostWorkforcemanagementCalendarURLIcsParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementCalendarURLIcsParamsWithContext(ctx context.Context) *PostWorkforcemanagementCalendarURLIcsParams {
	return &PostWorkforcemanagementCalendarURLIcsParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementCalendarURLIcsParamsWithHTTPClient creates a new PostWorkforcemanagementCalendarURLIcsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementCalendarURLIcsParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementCalendarURLIcsParams {
	return &PostWorkforcemanagementCalendarURLIcsParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementCalendarURLIcsParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement calendar Url ics operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementCalendarURLIcsParams struct {

	/* Language.

	   A language tag (which is sometimes referred to as a "locale identifier") to use to localize default activity code names in the ics-formatted calendar

	   Default: "en-US"
	*/
	Language *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workforcemanagement calendar Url ics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementCalendarURLIcsParams) WithDefaults() *PostWorkforcemanagementCalendarURLIcsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement calendar Url ics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementCalendarURLIcsParams) SetDefaults() {
	var (
		languageDefault = string("en-US")
	)

	val := PostWorkforcemanagementCalendarURLIcsParams{
		Language: &languageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the post workforcemanagement calendar Url ics params
func (o *PostWorkforcemanagementCalendarURLIcsParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementCalendarURLIcsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement calendar Url ics params
func (o *PostWorkforcemanagementCalendarURLIcsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement calendar Url ics params
func (o *PostWorkforcemanagementCalendarURLIcsParams) WithContext(ctx context.Context) *PostWorkforcemanagementCalendarURLIcsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement calendar Url ics params
func (o *PostWorkforcemanagementCalendarURLIcsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement calendar Url ics params
func (o *PostWorkforcemanagementCalendarURLIcsParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementCalendarURLIcsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement calendar Url ics params
func (o *PostWorkforcemanagementCalendarURLIcsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguage adds the language to the post workforcemanagement calendar Url ics params
func (o *PostWorkforcemanagementCalendarURLIcsParams) WithLanguage(language *string) *PostWorkforcemanagementCalendarURLIcsParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the post workforcemanagement calendar Url ics params
func (o *PostWorkforcemanagementCalendarURLIcsParams) SetLanguage(language *string) {
	o.Language = language
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementCalendarURLIcsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Language != nil {

		// query param language
		var qrLanguage string

		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {

			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
