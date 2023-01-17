// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetIntegrationsActionSchemaParams creates a new GetIntegrationsActionSchemaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIntegrationsActionSchemaParams() *GetIntegrationsActionSchemaParams {
	return &GetIntegrationsActionSchemaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationsActionSchemaParamsWithTimeout creates a new GetIntegrationsActionSchemaParams object
// with the ability to set a timeout on a request.
func NewGetIntegrationsActionSchemaParamsWithTimeout(timeout time.Duration) *GetIntegrationsActionSchemaParams {
	return &GetIntegrationsActionSchemaParams{
		timeout: timeout,
	}
}

// NewGetIntegrationsActionSchemaParamsWithContext creates a new GetIntegrationsActionSchemaParams object
// with the ability to set a context for a request.
func NewGetIntegrationsActionSchemaParamsWithContext(ctx context.Context) *GetIntegrationsActionSchemaParams {
	return &GetIntegrationsActionSchemaParams{
		Context: ctx,
	}
}

// NewGetIntegrationsActionSchemaParamsWithHTTPClient creates a new GetIntegrationsActionSchemaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIntegrationsActionSchemaParamsWithHTTPClient(client *http.Client) *GetIntegrationsActionSchemaParams {
	return &GetIntegrationsActionSchemaParams{
		HTTPClient: client,
	}
}

/*
GetIntegrationsActionSchemaParams contains all the parameters to send to the API endpoint

	for the get integrations action schema operation.

	Typically these are written to a http.Request.
*/
type GetIntegrationsActionSchemaParams struct {

	/* ActionID.

	   actionId
	*/
	ActionID string

	/* FileName.

	   Name of schema file to be retrieved for this action.
	*/
	FileName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get integrations action schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationsActionSchemaParams) WithDefaults() *GetIntegrationsActionSchemaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get integrations action schema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationsActionSchemaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get integrations action schema params
func (o *GetIntegrationsActionSchemaParams) WithTimeout(timeout time.Duration) *GetIntegrationsActionSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrations action schema params
func (o *GetIntegrationsActionSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrations action schema params
func (o *GetIntegrationsActionSchemaParams) WithContext(ctx context.Context) *GetIntegrationsActionSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrations action schema params
func (o *GetIntegrationsActionSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrations action schema params
func (o *GetIntegrationsActionSchemaParams) WithHTTPClient(client *http.Client) *GetIntegrationsActionSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrations action schema params
func (o *GetIntegrationsActionSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the get integrations action schema params
func (o *GetIntegrationsActionSchemaParams) WithActionID(actionID string) *GetIntegrationsActionSchemaParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the get integrations action schema params
func (o *GetIntegrationsActionSchemaParams) SetActionID(actionID string) {
	o.ActionID = actionID
}

// WithFileName adds the fileName to the get integrations action schema params
func (o *GetIntegrationsActionSchemaParams) WithFileName(fileName string) *GetIntegrationsActionSchemaParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the get integrations action schema params
func (o *GetIntegrationsActionSchemaParams) SetFileName(fileName string) {
	o.FileName = fileName
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationsActionSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID); err != nil {
		return err
	}

	// path param fileName
	if err := r.SetPathParam("fileName", o.FileName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
