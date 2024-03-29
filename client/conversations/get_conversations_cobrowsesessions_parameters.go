// Code generated by go-swagger; DO NOT EDIT.

package conversations

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

// NewGetConversationsCobrowsesessionsParams creates a new GetConversationsCobrowsesessionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConversationsCobrowsesessionsParams() *GetConversationsCobrowsesessionsParams {
	return &GetConversationsCobrowsesessionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsCobrowsesessionsParamsWithTimeout creates a new GetConversationsCobrowsesessionsParams object
// with the ability to set a timeout on a request.
func NewGetConversationsCobrowsesessionsParamsWithTimeout(timeout time.Duration) *GetConversationsCobrowsesessionsParams {
	return &GetConversationsCobrowsesessionsParams{
		timeout: timeout,
	}
}

// NewGetConversationsCobrowsesessionsParamsWithContext creates a new GetConversationsCobrowsesessionsParams object
// with the ability to set a context for a request.
func NewGetConversationsCobrowsesessionsParamsWithContext(ctx context.Context) *GetConversationsCobrowsesessionsParams {
	return &GetConversationsCobrowsesessionsParams{
		Context: ctx,
	}
}

// NewGetConversationsCobrowsesessionsParamsWithHTTPClient creates a new GetConversationsCobrowsesessionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConversationsCobrowsesessionsParamsWithHTTPClient(client *http.Client) *GetConversationsCobrowsesessionsParams {
	return &GetConversationsCobrowsesessionsParams{
		HTTPClient: client,
	}
}

/*
GetConversationsCobrowsesessionsParams contains all the parameters to send to the API endpoint

	for the get conversations cobrowsesessions operation.

	Typically these are written to a http.Request.
*/
type GetConversationsCobrowsesessionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get conversations cobrowsesessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsCobrowsesessionsParams) WithDefaults() *GetConversationsCobrowsesessionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get conversations cobrowsesessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsCobrowsesessionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get conversations cobrowsesessions params
func (o *GetConversationsCobrowsesessionsParams) WithTimeout(timeout time.Duration) *GetConversationsCobrowsesessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations cobrowsesessions params
func (o *GetConversationsCobrowsesessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations cobrowsesessions params
func (o *GetConversationsCobrowsesessionsParams) WithContext(ctx context.Context) *GetConversationsCobrowsesessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations cobrowsesessions params
func (o *GetConversationsCobrowsesessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations cobrowsesessions params
func (o *GetConversationsCobrowsesessionsParams) WithHTTPClient(client *http.Client) *GetConversationsCobrowsesessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations cobrowsesessions params
func (o *GetConversationsCobrowsesessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsCobrowsesessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
