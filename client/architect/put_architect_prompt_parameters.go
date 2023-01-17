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

// NewPutArchitectPromptParams creates a new PutArchitectPromptParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutArchitectPromptParams() *PutArchitectPromptParams {
	return &PutArchitectPromptParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutArchitectPromptParamsWithTimeout creates a new PutArchitectPromptParams object
// with the ability to set a timeout on a request.
func NewPutArchitectPromptParamsWithTimeout(timeout time.Duration) *PutArchitectPromptParams {
	return &PutArchitectPromptParams{
		timeout: timeout,
	}
}

// NewPutArchitectPromptParamsWithContext creates a new PutArchitectPromptParams object
// with the ability to set a context for a request.
func NewPutArchitectPromptParamsWithContext(ctx context.Context) *PutArchitectPromptParams {
	return &PutArchitectPromptParams{
		Context: ctx,
	}
}

// NewPutArchitectPromptParamsWithHTTPClient creates a new PutArchitectPromptParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutArchitectPromptParamsWithHTTPClient(client *http.Client) *PutArchitectPromptParams {
	return &PutArchitectPromptParams{
		HTTPClient: client,
	}
}

/*
PutArchitectPromptParams contains all the parameters to send to the API endpoint

	for the put architect prompt operation.

	Typically these are written to a http.Request.
*/
type PutArchitectPromptParams struct {

	// Body.
	Body *models.Prompt

	/* PromptID.

	   Prompt ID
	*/
	PromptID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put architect prompt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutArchitectPromptParams) WithDefaults() *PutArchitectPromptParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put architect prompt params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutArchitectPromptParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put architect prompt params
func (o *PutArchitectPromptParams) WithTimeout(timeout time.Duration) *PutArchitectPromptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put architect prompt params
func (o *PutArchitectPromptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put architect prompt params
func (o *PutArchitectPromptParams) WithContext(ctx context.Context) *PutArchitectPromptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put architect prompt params
func (o *PutArchitectPromptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put architect prompt params
func (o *PutArchitectPromptParams) WithHTTPClient(client *http.Client) *PutArchitectPromptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put architect prompt params
func (o *PutArchitectPromptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put architect prompt params
func (o *PutArchitectPromptParams) WithBody(body *models.Prompt) *PutArchitectPromptParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put architect prompt params
func (o *PutArchitectPromptParams) SetBody(body *models.Prompt) {
	o.Body = body
}

// WithPromptID adds the promptID to the put architect prompt params
func (o *PutArchitectPromptParams) WithPromptID(promptID string) *PutArchitectPromptParams {
	o.SetPromptID(promptID)
	return o
}

// SetPromptID adds the promptId to the put architect prompt params
func (o *PutArchitectPromptParams) SetPromptID(promptID string) {
	o.PromptID = promptID
}

// WriteToRequest writes these params to a swagger request
func (o *PutArchitectPromptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
