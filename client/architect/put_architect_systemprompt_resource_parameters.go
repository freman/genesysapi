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

	"github.com/freman/genesysapi/models"
)

// NewPutArchitectSystempromptResourceParams creates a new PutArchitectSystempromptResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutArchitectSystempromptResourceParams() *PutArchitectSystempromptResourceParams {
	return &PutArchitectSystempromptResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutArchitectSystempromptResourceParamsWithTimeout creates a new PutArchitectSystempromptResourceParams object
// with the ability to set a timeout on a request.
func NewPutArchitectSystempromptResourceParamsWithTimeout(timeout time.Duration) *PutArchitectSystempromptResourceParams {
	return &PutArchitectSystempromptResourceParams{
		timeout: timeout,
	}
}

// NewPutArchitectSystempromptResourceParamsWithContext creates a new PutArchitectSystempromptResourceParams object
// with the ability to set a context for a request.
func NewPutArchitectSystempromptResourceParamsWithContext(ctx context.Context) *PutArchitectSystempromptResourceParams {
	return &PutArchitectSystempromptResourceParams{
		Context: ctx,
	}
}

// NewPutArchitectSystempromptResourceParamsWithHTTPClient creates a new PutArchitectSystempromptResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutArchitectSystempromptResourceParamsWithHTTPClient(client *http.Client) *PutArchitectSystempromptResourceParams {
	return &PutArchitectSystempromptResourceParams{
		HTTPClient: client,
	}
}

/*
PutArchitectSystempromptResourceParams contains all the parameters to send to the API endpoint

	for the put architect systemprompt resource operation.

	Typically these are written to a http.Request.
*/
type PutArchitectSystempromptResourceParams struct {

	// Body.
	Body *models.SystemPromptAsset

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

// WithDefaults hydrates default values in the put architect systemprompt resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutArchitectSystempromptResourceParams) WithDefaults() *PutArchitectSystempromptResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put architect systemprompt resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutArchitectSystempromptResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put architect systemprompt resource params
func (o *PutArchitectSystempromptResourceParams) WithTimeout(timeout time.Duration) *PutArchitectSystempromptResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put architect systemprompt resource params
func (o *PutArchitectSystempromptResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put architect systemprompt resource params
func (o *PutArchitectSystempromptResourceParams) WithContext(ctx context.Context) *PutArchitectSystempromptResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put architect systemprompt resource params
func (o *PutArchitectSystempromptResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put architect systemprompt resource params
func (o *PutArchitectSystempromptResourceParams) WithHTTPClient(client *http.Client) *PutArchitectSystempromptResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put architect systemprompt resource params
func (o *PutArchitectSystempromptResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put architect systemprompt resource params
func (o *PutArchitectSystempromptResourceParams) WithBody(body *models.SystemPromptAsset) *PutArchitectSystempromptResourceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put architect systemprompt resource params
func (o *PutArchitectSystempromptResourceParams) SetBody(body *models.SystemPromptAsset) {
	o.Body = body
}

// WithLanguageCode adds the languageCode to the put architect systemprompt resource params
func (o *PutArchitectSystempromptResourceParams) WithLanguageCode(languageCode string) *PutArchitectSystempromptResourceParams {
	o.SetLanguageCode(languageCode)
	return o
}

// SetLanguageCode adds the languageCode to the put architect systemprompt resource params
func (o *PutArchitectSystempromptResourceParams) SetLanguageCode(languageCode string) {
	o.LanguageCode = languageCode
}

// WithPromptID adds the promptID to the put architect systemprompt resource params
func (o *PutArchitectSystempromptResourceParams) WithPromptID(promptID string) *PutArchitectSystempromptResourceParams {
	o.SetPromptID(promptID)
	return o
}

// SetPromptID adds the promptId to the put architect systemprompt resource params
func (o *PutArchitectSystempromptResourceParams) SetPromptID(promptID string) {
	o.PromptID = promptID
}

// WriteToRequest writes these params to a swagger request
func (o *PutArchitectSystempromptResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
