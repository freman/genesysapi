// Code generated by go-swagger; DO NOT EDIT.

package scripts

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

// NewGetScriptsPublishedScriptIDVariablesParams creates a new GetScriptsPublishedScriptIDVariablesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScriptsPublishedScriptIDVariablesParams() *GetScriptsPublishedScriptIDVariablesParams {
	return &GetScriptsPublishedScriptIDVariablesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScriptsPublishedScriptIDVariablesParamsWithTimeout creates a new GetScriptsPublishedScriptIDVariablesParams object
// with the ability to set a timeout on a request.
func NewGetScriptsPublishedScriptIDVariablesParamsWithTimeout(timeout time.Duration) *GetScriptsPublishedScriptIDVariablesParams {
	return &GetScriptsPublishedScriptIDVariablesParams{
		timeout: timeout,
	}
}

// NewGetScriptsPublishedScriptIDVariablesParamsWithContext creates a new GetScriptsPublishedScriptIDVariablesParams object
// with the ability to set a context for a request.
func NewGetScriptsPublishedScriptIDVariablesParamsWithContext(ctx context.Context) *GetScriptsPublishedScriptIDVariablesParams {
	return &GetScriptsPublishedScriptIDVariablesParams{
		Context: ctx,
	}
}

// NewGetScriptsPublishedScriptIDVariablesParamsWithHTTPClient creates a new GetScriptsPublishedScriptIDVariablesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScriptsPublishedScriptIDVariablesParamsWithHTTPClient(client *http.Client) *GetScriptsPublishedScriptIDVariablesParams {
	return &GetScriptsPublishedScriptIDVariablesParams{
		HTTPClient: client,
	}
}

/*
GetScriptsPublishedScriptIDVariablesParams contains all the parameters to send to the API endpoint

	for the get scripts published script Id variables operation.

	Typically these are written to a http.Request.
*/
type GetScriptsPublishedScriptIDVariablesParams struct {

	/* Input.

	   input
	*/
	Input *string

	/* Output.

	   output
	*/
	Output *string

	/* ScriptDataVersion.

	   Advanced usage - controls the data version of the script
	*/
	ScriptDataVersion *string

	/* ScriptID.

	   Script ID
	*/
	ScriptID string

	/* Type.

	   type
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get scripts published script Id variables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScriptsPublishedScriptIDVariablesParams) WithDefaults() *GetScriptsPublishedScriptIDVariablesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scripts published script Id variables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScriptsPublishedScriptIDVariablesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) WithTimeout(timeout time.Duration) *GetScriptsPublishedScriptIDVariablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) WithContext(ctx context.Context) *GetScriptsPublishedScriptIDVariablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) WithHTTPClient(client *http.Client) *GetScriptsPublishedScriptIDVariablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) WithInput(input *string) *GetScriptsPublishedScriptIDVariablesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) SetInput(input *string) {
	o.Input = input
}

// WithOutput adds the output to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) WithOutput(output *string) *GetScriptsPublishedScriptIDVariablesParams {
	o.SetOutput(output)
	return o
}

// SetOutput adds the output to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) SetOutput(output *string) {
	o.Output = output
}

// WithScriptDataVersion adds the scriptDataVersion to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) WithScriptDataVersion(scriptDataVersion *string) *GetScriptsPublishedScriptIDVariablesParams {
	o.SetScriptDataVersion(scriptDataVersion)
	return o
}

// SetScriptDataVersion adds the scriptDataVersion to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) SetScriptDataVersion(scriptDataVersion *string) {
	o.ScriptDataVersion = scriptDataVersion
}

// WithScriptID adds the scriptID to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) WithScriptID(scriptID string) *GetScriptsPublishedScriptIDVariablesParams {
	o.SetScriptID(scriptID)
	return o
}

// SetScriptID adds the scriptId to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) SetScriptID(scriptID string) {
	o.ScriptID = scriptID
}

// WithType adds the typeVar to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) WithType(typeVar *string) *GetScriptsPublishedScriptIDVariablesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get scripts published script Id variables params
func (o *GetScriptsPublishedScriptIDVariablesParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetScriptsPublishedScriptIDVariablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Input != nil {

		// query param input
		var qrInput string

		if o.Input != nil {
			qrInput = *o.Input
		}
		qInput := qrInput
		if qInput != "" {

			if err := r.SetQueryParam("input", qInput); err != nil {
				return err
			}
		}
	}

	if o.Output != nil {

		// query param output
		var qrOutput string

		if o.Output != nil {
			qrOutput = *o.Output
		}
		qOutput := qrOutput
		if qOutput != "" {

			if err := r.SetQueryParam("output", qOutput); err != nil {
				return err
			}
		}
	}

	if o.ScriptDataVersion != nil {

		// query param scriptDataVersion
		var qrScriptDataVersion string

		if o.ScriptDataVersion != nil {
			qrScriptDataVersion = *o.ScriptDataVersion
		}
		qScriptDataVersion := qrScriptDataVersion
		if qScriptDataVersion != "" {

			if err := r.SetQueryParam("scriptDataVersion", qScriptDataVersion); err != nil {
				return err
			}
		}
	}

	// path param scriptId
	if err := r.SetPathParam("scriptId", o.ScriptID); err != nil {
		return err
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
