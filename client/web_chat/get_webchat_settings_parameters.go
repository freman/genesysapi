// Code generated by go-swagger; DO NOT EDIT.

package web_chat

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

// NewGetWebchatSettingsParams creates a new GetWebchatSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWebchatSettingsParams() *GetWebchatSettingsParams {
	return &GetWebchatSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebchatSettingsParamsWithTimeout creates a new GetWebchatSettingsParams object
// with the ability to set a timeout on a request.
func NewGetWebchatSettingsParamsWithTimeout(timeout time.Duration) *GetWebchatSettingsParams {
	return &GetWebchatSettingsParams{
		timeout: timeout,
	}
}

// NewGetWebchatSettingsParamsWithContext creates a new GetWebchatSettingsParams object
// with the ability to set a context for a request.
func NewGetWebchatSettingsParamsWithContext(ctx context.Context) *GetWebchatSettingsParams {
	return &GetWebchatSettingsParams{
		Context: ctx,
	}
}

// NewGetWebchatSettingsParamsWithHTTPClient creates a new GetWebchatSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWebchatSettingsParamsWithHTTPClient(client *http.Client) *GetWebchatSettingsParams {
	return &GetWebchatSettingsParams{
		HTTPClient: client,
	}
}

/*
GetWebchatSettingsParams contains all the parameters to send to the API endpoint

	for the get webchat settings operation.

	Typically these are written to a http.Request.
*/
type GetWebchatSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get webchat settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebchatSettingsParams) WithDefaults() *GetWebchatSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get webchat settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebchatSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get webchat settings params
func (o *GetWebchatSettingsParams) WithTimeout(timeout time.Duration) *GetWebchatSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get webchat settings params
func (o *GetWebchatSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get webchat settings params
func (o *GetWebchatSettingsParams) WithContext(ctx context.Context) *GetWebchatSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get webchat settings params
func (o *GetWebchatSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get webchat settings params
func (o *GetWebchatSettingsParams) WithHTTPClient(client *http.Client) *GetWebchatSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get webchat settings params
func (o *GetWebchatSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebchatSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
