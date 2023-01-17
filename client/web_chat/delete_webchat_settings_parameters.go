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

// NewDeleteWebchatSettingsParams creates a new DeleteWebchatSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteWebchatSettingsParams() *DeleteWebchatSettingsParams {
	return &DeleteWebchatSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWebchatSettingsParamsWithTimeout creates a new DeleteWebchatSettingsParams object
// with the ability to set a timeout on a request.
func NewDeleteWebchatSettingsParamsWithTimeout(timeout time.Duration) *DeleteWebchatSettingsParams {
	return &DeleteWebchatSettingsParams{
		timeout: timeout,
	}
}

// NewDeleteWebchatSettingsParamsWithContext creates a new DeleteWebchatSettingsParams object
// with the ability to set a context for a request.
func NewDeleteWebchatSettingsParamsWithContext(ctx context.Context) *DeleteWebchatSettingsParams {
	return &DeleteWebchatSettingsParams{
		Context: ctx,
	}
}

// NewDeleteWebchatSettingsParamsWithHTTPClient creates a new DeleteWebchatSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteWebchatSettingsParamsWithHTTPClient(client *http.Client) *DeleteWebchatSettingsParams {
	return &DeleteWebchatSettingsParams{
		HTTPClient: client,
	}
}

/*
DeleteWebchatSettingsParams contains all the parameters to send to the API endpoint

	for the delete webchat settings operation.

	Typically these are written to a http.Request.
*/
type DeleteWebchatSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete webchat settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWebchatSettingsParams) WithDefaults() *DeleteWebchatSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete webchat settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWebchatSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete webchat settings params
func (o *DeleteWebchatSettingsParams) WithTimeout(timeout time.Duration) *DeleteWebchatSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete webchat settings params
func (o *DeleteWebchatSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete webchat settings params
func (o *DeleteWebchatSettingsParams) WithContext(ctx context.Context) *DeleteWebchatSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete webchat settings params
func (o *DeleteWebchatSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete webchat settings params
func (o *DeleteWebchatSettingsParams) WithHTTPClient(client *http.Client) *DeleteWebchatSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete webchat settings params
func (o *DeleteWebchatSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWebchatSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
