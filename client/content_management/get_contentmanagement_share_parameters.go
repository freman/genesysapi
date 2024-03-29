// Code generated by go-swagger; DO NOT EDIT.

package content_management

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

// NewGetContentmanagementShareParams creates a new GetContentmanagementShareParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContentmanagementShareParams() *GetContentmanagementShareParams {
	return &GetContentmanagementShareParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentmanagementShareParamsWithTimeout creates a new GetContentmanagementShareParams object
// with the ability to set a timeout on a request.
func NewGetContentmanagementShareParamsWithTimeout(timeout time.Duration) *GetContentmanagementShareParams {
	return &GetContentmanagementShareParams{
		timeout: timeout,
	}
}

// NewGetContentmanagementShareParamsWithContext creates a new GetContentmanagementShareParams object
// with the ability to set a context for a request.
func NewGetContentmanagementShareParamsWithContext(ctx context.Context) *GetContentmanagementShareParams {
	return &GetContentmanagementShareParams{
		Context: ctx,
	}
}

// NewGetContentmanagementShareParamsWithHTTPClient creates a new GetContentmanagementShareParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContentmanagementShareParamsWithHTTPClient(client *http.Client) *GetContentmanagementShareParams {
	return &GetContentmanagementShareParams{
		HTTPClient: client,
	}
}

/*
GetContentmanagementShareParams contains all the parameters to send to the API endpoint

	for the get contentmanagement share operation.

	Typically these are written to a http.Request.
*/
type GetContentmanagementShareParams struct {

	/* Expand.

	   Which fields, if any, to expand.
	*/
	Expand []string

	/* ShareID.

	   Share ID
	*/
	ShareID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contentmanagement share params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentmanagementShareParams) WithDefaults() *GetContentmanagementShareParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contentmanagement share params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentmanagementShareParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get contentmanagement share params
func (o *GetContentmanagementShareParams) WithTimeout(timeout time.Duration) *GetContentmanagementShareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contentmanagement share params
func (o *GetContentmanagementShareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contentmanagement share params
func (o *GetContentmanagementShareParams) WithContext(ctx context.Context) *GetContentmanagementShareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contentmanagement share params
func (o *GetContentmanagementShareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contentmanagement share params
func (o *GetContentmanagementShareParams) WithHTTPClient(client *http.Client) *GetContentmanagementShareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contentmanagement share params
func (o *GetContentmanagementShareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get contentmanagement share params
func (o *GetContentmanagementShareParams) WithExpand(expand []string) *GetContentmanagementShareParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get contentmanagement share params
func (o *GetContentmanagementShareParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithShareID adds the shareID to the get contentmanagement share params
func (o *GetContentmanagementShareParams) WithShareID(shareID string) *GetContentmanagementShareParams {
	o.SetShareID(shareID)
	return o
}

// SetShareID adds the shareId to the get contentmanagement share params
func (o *GetContentmanagementShareParams) SetShareID(shareID string) {
	o.ShareID = shareID
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentmanagementShareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	// path param shareId
	if err := r.SetPathParam("shareId", o.ShareID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetContentmanagementShare binds the parameter expand
func (o *GetContentmanagementShareParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}
