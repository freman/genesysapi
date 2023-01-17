// Code generated by go-swagger; DO NOT EDIT.

package languages

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

// NewGetLanguagesTranslationsOrganizationParams creates a new GetLanguagesTranslationsOrganizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLanguagesTranslationsOrganizationParams() *GetLanguagesTranslationsOrganizationParams {
	return &GetLanguagesTranslationsOrganizationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLanguagesTranslationsOrganizationParamsWithTimeout creates a new GetLanguagesTranslationsOrganizationParams object
// with the ability to set a timeout on a request.
func NewGetLanguagesTranslationsOrganizationParamsWithTimeout(timeout time.Duration) *GetLanguagesTranslationsOrganizationParams {
	return &GetLanguagesTranslationsOrganizationParams{
		timeout: timeout,
	}
}

// NewGetLanguagesTranslationsOrganizationParamsWithContext creates a new GetLanguagesTranslationsOrganizationParams object
// with the ability to set a context for a request.
func NewGetLanguagesTranslationsOrganizationParamsWithContext(ctx context.Context) *GetLanguagesTranslationsOrganizationParams {
	return &GetLanguagesTranslationsOrganizationParams{
		Context: ctx,
	}
}

// NewGetLanguagesTranslationsOrganizationParamsWithHTTPClient creates a new GetLanguagesTranslationsOrganizationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLanguagesTranslationsOrganizationParamsWithHTTPClient(client *http.Client) *GetLanguagesTranslationsOrganizationParams {
	return &GetLanguagesTranslationsOrganizationParams{
		HTTPClient: client,
	}
}

/*
GetLanguagesTranslationsOrganizationParams contains all the parameters to send to the API endpoint

	for the get languages translations organization operation.

	Typically these are written to a http.Request.
*/
type GetLanguagesTranslationsOrganizationParams struct {

	/* Language.

	   The language of the translation to retrieve for the organization
	*/
	Language string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get languages translations organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLanguagesTranslationsOrganizationParams) WithDefaults() *GetLanguagesTranslationsOrganizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get languages translations organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLanguagesTranslationsOrganizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get languages translations organization params
func (o *GetLanguagesTranslationsOrganizationParams) WithTimeout(timeout time.Duration) *GetLanguagesTranslationsOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get languages translations organization params
func (o *GetLanguagesTranslationsOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get languages translations organization params
func (o *GetLanguagesTranslationsOrganizationParams) WithContext(ctx context.Context) *GetLanguagesTranslationsOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get languages translations organization params
func (o *GetLanguagesTranslationsOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get languages translations organization params
func (o *GetLanguagesTranslationsOrganizationParams) WithHTTPClient(client *http.Client) *GetLanguagesTranslationsOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get languages translations organization params
func (o *GetLanguagesTranslationsOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguage adds the language to the get languages translations organization params
func (o *GetLanguagesTranslationsOrganizationParams) WithLanguage(language string) *GetLanguagesTranslationsOrganizationParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get languages translations organization params
func (o *GetLanguagesTranslationsOrganizationParams) SetLanguage(language string) {
	o.Language = language
}

// WriteToRequest writes these params to a swagger request
func (o *GetLanguagesTranslationsOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param language
	qrLanguage := o.Language
	qLanguage := qrLanguage
	if qLanguage != "" {

		if err := r.SetQueryParam("language", qLanguage); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
