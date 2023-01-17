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
	"github.com/go-openapi/swag"
)

// NewDeleteArchitectPromptsParams creates a new DeleteArchitectPromptsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteArchitectPromptsParams() *DeleteArchitectPromptsParams {
	return &DeleteArchitectPromptsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteArchitectPromptsParamsWithTimeout creates a new DeleteArchitectPromptsParams object
// with the ability to set a timeout on a request.
func NewDeleteArchitectPromptsParamsWithTimeout(timeout time.Duration) *DeleteArchitectPromptsParams {
	return &DeleteArchitectPromptsParams{
		timeout: timeout,
	}
}

// NewDeleteArchitectPromptsParamsWithContext creates a new DeleteArchitectPromptsParams object
// with the ability to set a context for a request.
func NewDeleteArchitectPromptsParamsWithContext(ctx context.Context) *DeleteArchitectPromptsParams {
	return &DeleteArchitectPromptsParams{
		Context: ctx,
	}
}

// NewDeleteArchitectPromptsParamsWithHTTPClient creates a new DeleteArchitectPromptsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteArchitectPromptsParamsWithHTTPClient(client *http.Client) *DeleteArchitectPromptsParams {
	return &DeleteArchitectPromptsParams{
		HTTPClient: client,
	}
}

/*
DeleteArchitectPromptsParams contains all the parameters to send to the API endpoint

	for the delete architect prompts operation.

	Typically these are written to a http.Request.
*/
type DeleteArchitectPromptsParams struct {

	/* ID.

	   List of Prompt IDs
	*/
	ID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete architect prompts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteArchitectPromptsParams) WithDefaults() *DeleteArchitectPromptsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete architect prompts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteArchitectPromptsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) WithTimeout(timeout time.Duration) *DeleteArchitectPromptsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) WithContext(ctx context.Context) *DeleteArchitectPromptsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) WithHTTPClient(client *http.Client) *DeleteArchitectPromptsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) WithID(id []string) *DeleteArchitectPromptsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) SetID(id []string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteArchitectPromptsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeleteArchitectPrompts binds the parameter id
func (o *DeleteArchitectPromptsParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}
