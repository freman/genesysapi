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

// NewGetConversationsCallsParams creates a new GetConversationsCallsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConversationsCallsParams() *GetConversationsCallsParams {
	return &GetConversationsCallsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsCallsParamsWithTimeout creates a new GetConversationsCallsParams object
// with the ability to set a timeout on a request.
func NewGetConversationsCallsParamsWithTimeout(timeout time.Duration) *GetConversationsCallsParams {
	return &GetConversationsCallsParams{
		timeout: timeout,
	}
}

// NewGetConversationsCallsParamsWithContext creates a new GetConversationsCallsParams object
// with the ability to set a context for a request.
func NewGetConversationsCallsParamsWithContext(ctx context.Context) *GetConversationsCallsParams {
	return &GetConversationsCallsParams{
		Context: ctx,
	}
}

// NewGetConversationsCallsParamsWithHTTPClient creates a new GetConversationsCallsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConversationsCallsParamsWithHTTPClient(client *http.Client) *GetConversationsCallsParams {
	return &GetConversationsCallsParams{
		HTTPClient: client,
	}
}

/*
GetConversationsCallsParams contains all the parameters to send to the API endpoint

	for the get conversations calls operation.

	Typically these are written to a http.Request.
*/
type GetConversationsCallsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get conversations calls params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsCallsParams) WithDefaults() *GetConversationsCallsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get conversations calls params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsCallsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get conversations calls params
func (o *GetConversationsCallsParams) WithTimeout(timeout time.Duration) *GetConversationsCallsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations calls params
func (o *GetConversationsCallsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations calls params
func (o *GetConversationsCallsParams) WithContext(ctx context.Context) *GetConversationsCallsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations calls params
func (o *GetConversationsCallsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations calls params
func (o *GetConversationsCallsParams) WithHTTPClient(client *http.Client) *GetConversationsCallsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations calls params
func (o *GetConversationsCallsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsCallsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
