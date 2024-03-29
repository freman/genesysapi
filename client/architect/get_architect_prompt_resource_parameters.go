// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewGetArchitectPromptResourceParams creates a new GetArchitectPromptResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetArchitectPromptResourceParams() *GetArchitectPromptResourceParams {
	return &GetArchitectPromptResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectPromptResourceParamsWithTimeout creates a new GetArchitectPromptResourceParams object
// with the ability to set a timeout on a request.
func NewGetArchitectPromptResourceParamsWithTimeout(timeout time.Duration) *GetArchitectPromptResourceParams {
	return &GetArchitectPromptResourceParams{
		timeout: timeout,
	}
}

// NewGetArchitectPromptResourceParamsWithContext creates a new GetArchitectPromptResourceParams object
// with the ability to set a context for a request.
func NewGetArchitectPromptResourceParamsWithContext(ctx context.Context) *GetArchitectPromptResourceParams {
	return &GetArchitectPromptResourceParams{
		Context: ctx,
	}
}

// NewGetArchitectPromptResourceParamsWithHTTPClient creates a new GetArchitectPromptResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetArchitectPromptResourceParamsWithHTTPClient(client *http.Client) *GetArchitectPromptResourceParams {
	return &GetArchitectPromptResourceParams{
		HTTPClient: client,
	}
}

/*
GetArchitectPromptResourceParams contains all the parameters to send to the API endpoint

	for the get architect prompt resource operation.

	Typically these are written to a http.Request.
*/
type GetArchitectPromptResourceParams struct {

	/* LanguageCode.

	   Language
	*/
	LanguageCode string

	/* PromptID.

	   Prompt ID
	*/
	PromptID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get architect prompt resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectPromptResourceParams) WithDefaults() *GetArchitectPromptResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get architect prompt resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectPromptResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get architect prompt resource params
func (o *GetArchitectPromptResourceParams) WithTimeout(timeout time.Duration) *GetArchitectPromptResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect prompt resource params
func (o *GetArchitectPromptResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect prompt resource params
func (o *GetArchitectPromptResourceParams) WithContext(ctx context.Context) *GetArchitectPromptResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect prompt resource params
func (o *GetArchitectPromptResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect prompt resource params
func (o *GetArchitectPromptResourceParams) WithHTTPClient(client *http.Client) *GetArchitectPromptResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect prompt resource params
func (o *GetArchitectPromptResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguageCode adds the languageCode to the get architect prompt resource params
func (o *GetArchitectPromptResourceParams) WithLanguageCode(languageCode string) *GetArchitectPromptResourceParams {
	o.SetLanguageCode(languageCode)
	return o
}

// SetLanguageCode adds the languageCode to the get architect prompt resource params
func (o *GetArchitectPromptResourceParams) SetLanguageCode(languageCode string) {
	o.LanguageCode = languageCode
}

// WithPromptID adds the promptID to the get architect prompt resource params
func (o *GetArchitectPromptResourceParams) WithPromptID(promptID string) *GetArchitectPromptResourceParams {
	o.SetPromptID(promptID)
	return o
}

// SetPromptID adds the promptId to the get architect prompt resource params
func (o *GetArchitectPromptResourceParams) SetPromptID(promptID string) {
	o.PromptID = promptID
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectPromptResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param languageCode
	if err := r.SetPathParam("languageCode", o.LanguageCode); err != nil {
		return err
	}

	// path param promptId
	if err := r.SetPathParam("promptId", o.PromptID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
