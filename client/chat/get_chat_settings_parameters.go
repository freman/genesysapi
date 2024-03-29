// Code generated by go-swagger; DO NOT EDIT.

package chat

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

// NewGetChatSettingsParams creates a new GetChatSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetChatSettingsParams() *GetChatSettingsParams {
	return &GetChatSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetChatSettingsParamsWithTimeout creates a new GetChatSettingsParams object
// with the ability to set a timeout on a request.
func NewGetChatSettingsParamsWithTimeout(timeout time.Duration) *GetChatSettingsParams {
	return &GetChatSettingsParams{
		timeout: timeout,
	}
}

// NewGetChatSettingsParamsWithContext creates a new GetChatSettingsParams object
// with the ability to set a context for a request.
func NewGetChatSettingsParamsWithContext(ctx context.Context) *GetChatSettingsParams {
	return &GetChatSettingsParams{
		Context: ctx,
	}
}

// NewGetChatSettingsParamsWithHTTPClient creates a new GetChatSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetChatSettingsParamsWithHTTPClient(client *http.Client) *GetChatSettingsParams {
	return &GetChatSettingsParams{
		HTTPClient: client,
	}
}

/*
GetChatSettingsParams contains all the parameters to send to the API endpoint

	for the get chat settings operation.

	Typically these are written to a http.Request.
*/
type GetChatSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get chat settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChatSettingsParams) WithDefaults() *GetChatSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get chat settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChatSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get chat settings params
func (o *GetChatSettingsParams) WithTimeout(timeout time.Duration) *GetChatSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get chat settings params
func (o *GetChatSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get chat settings params
func (o *GetChatSettingsParams) WithContext(ctx context.Context) *GetChatSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get chat settings params
func (o *GetChatSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get chat settings params
func (o *GetChatSettingsParams) WithHTTPClient(client *http.Client) *GetChatSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get chat settings params
func (o *GetChatSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetChatSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
