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

// NewGetLanguagesTranslationsBuiltinParams creates a new GetLanguagesTranslationsBuiltinParams object
// with the default values initialized.
func NewGetLanguagesTranslationsBuiltinParams() *GetLanguagesTranslationsBuiltinParams {
	var ()
	return &GetLanguagesTranslationsBuiltinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLanguagesTranslationsBuiltinParamsWithTimeout creates a new GetLanguagesTranslationsBuiltinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLanguagesTranslationsBuiltinParamsWithTimeout(timeout time.Duration) *GetLanguagesTranslationsBuiltinParams {
	var ()
	return &GetLanguagesTranslationsBuiltinParams{

		timeout: timeout,
	}
}

// NewGetLanguagesTranslationsBuiltinParamsWithContext creates a new GetLanguagesTranslationsBuiltinParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLanguagesTranslationsBuiltinParamsWithContext(ctx context.Context) *GetLanguagesTranslationsBuiltinParams {
	var ()
	return &GetLanguagesTranslationsBuiltinParams{

		Context: ctx,
	}
}

// NewGetLanguagesTranslationsBuiltinParamsWithHTTPClient creates a new GetLanguagesTranslationsBuiltinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLanguagesTranslationsBuiltinParamsWithHTTPClient(client *http.Client) *GetLanguagesTranslationsBuiltinParams {
	var ()
	return &GetLanguagesTranslationsBuiltinParams{
		HTTPClient: client,
	}
}

/*GetLanguagesTranslationsBuiltinParams contains all the parameters to send to the API endpoint
for the get languages translations builtin operation typically these are written to a http.Request
*/
type GetLanguagesTranslationsBuiltinParams struct {

	/*Language
	  The language of the builtin translation to retrieve

	*/
	Language string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get languages translations builtin params
func (o *GetLanguagesTranslationsBuiltinParams) WithTimeout(timeout time.Duration) *GetLanguagesTranslationsBuiltinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get languages translations builtin params
func (o *GetLanguagesTranslationsBuiltinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get languages translations builtin params
func (o *GetLanguagesTranslationsBuiltinParams) WithContext(ctx context.Context) *GetLanguagesTranslationsBuiltinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get languages translations builtin params
func (o *GetLanguagesTranslationsBuiltinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get languages translations builtin params
func (o *GetLanguagesTranslationsBuiltinParams) WithHTTPClient(client *http.Client) *GetLanguagesTranslationsBuiltinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get languages translations builtin params
func (o *GetLanguagesTranslationsBuiltinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguage adds the language to the get languages translations builtin params
func (o *GetLanguagesTranslationsBuiltinParams) WithLanguage(language string) *GetLanguagesTranslationsBuiltinParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get languages translations builtin params
func (o *GetLanguagesTranslationsBuiltinParams) SetLanguage(language string) {
	o.Language = language
}

// WriteToRequest writes these params to a swagger request
func (o *GetLanguagesTranslationsBuiltinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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