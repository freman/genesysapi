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

// NewGetScriptPageParams creates a new GetScriptPageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScriptPageParams() *GetScriptPageParams {
	return &GetScriptPageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScriptPageParamsWithTimeout creates a new GetScriptPageParams object
// with the ability to set a timeout on a request.
func NewGetScriptPageParamsWithTimeout(timeout time.Duration) *GetScriptPageParams {
	return &GetScriptPageParams{
		timeout: timeout,
	}
}

// NewGetScriptPageParamsWithContext creates a new GetScriptPageParams object
// with the ability to set a context for a request.
func NewGetScriptPageParamsWithContext(ctx context.Context) *GetScriptPageParams {
	return &GetScriptPageParams{
		Context: ctx,
	}
}

// NewGetScriptPageParamsWithHTTPClient creates a new GetScriptPageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScriptPageParamsWithHTTPClient(client *http.Client) *GetScriptPageParams {
	return &GetScriptPageParams{
		HTTPClient: client,
	}
}

/*
GetScriptPageParams contains all the parameters to send to the API endpoint

	for the get script page operation.

	Typically these are written to a http.Request.
*/
type GetScriptPageParams struct {

	/* PageID.

	   Page ID
	*/
	PageID string

	/* ScriptDataVersion.

	   Advanced usage - controls the data version of the script
	*/
	ScriptDataVersion *string

	/* ScriptID.

	   Script ID
	*/
	ScriptID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get script page params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScriptPageParams) WithDefaults() *GetScriptPageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get script page params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScriptPageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get script page params
func (o *GetScriptPageParams) WithTimeout(timeout time.Duration) *GetScriptPageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get script page params
func (o *GetScriptPageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get script page params
func (o *GetScriptPageParams) WithContext(ctx context.Context) *GetScriptPageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get script page params
func (o *GetScriptPageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get script page params
func (o *GetScriptPageParams) WithHTTPClient(client *http.Client) *GetScriptPageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get script page params
func (o *GetScriptPageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageID adds the pageID to the get script page params
func (o *GetScriptPageParams) WithPageID(pageID string) *GetScriptPageParams {
	o.SetPageID(pageID)
	return o
}

// SetPageID adds the pageId to the get script page params
func (o *GetScriptPageParams) SetPageID(pageID string) {
	o.PageID = pageID
}

// WithScriptDataVersion adds the scriptDataVersion to the get script page params
func (o *GetScriptPageParams) WithScriptDataVersion(scriptDataVersion *string) *GetScriptPageParams {
	o.SetScriptDataVersion(scriptDataVersion)
	return o
}

// SetScriptDataVersion adds the scriptDataVersion to the get script page params
func (o *GetScriptPageParams) SetScriptDataVersion(scriptDataVersion *string) {
	o.ScriptDataVersion = scriptDataVersion
}

// WithScriptID adds the scriptID to the get script page params
func (o *GetScriptPageParams) WithScriptID(scriptID string) *GetScriptPageParams {
	o.SetScriptID(scriptID)
	return o
}

// SetScriptID adds the scriptId to the get script page params
func (o *GetScriptPageParams) SetScriptID(scriptID string) {
	o.ScriptID = scriptID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScriptPageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pageId
	if err := r.SetPathParam("pageId", o.PageID); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
