// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewDeleteRoutingSettingsParams creates a new DeleteRoutingSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRoutingSettingsParams() *DeleteRoutingSettingsParams {
	return &DeleteRoutingSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoutingSettingsParamsWithTimeout creates a new DeleteRoutingSettingsParams object
// with the ability to set a timeout on a request.
func NewDeleteRoutingSettingsParamsWithTimeout(timeout time.Duration) *DeleteRoutingSettingsParams {
	return &DeleteRoutingSettingsParams{
		timeout: timeout,
	}
}

// NewDeleteRoutingSettingsParamsWithContext creates a new DeleteRoutingSettingsParams object
// with the ability to set a context for a request.
func NewDeleteRoutingSettingsParamsWithContext(ctx context.Context) *DeleteRoutingSettingsParams {
	return &DeleteRoutingSettingsParams{
		Context: ctx,
	}
}

// NewDeleteRoutingSettingsParamsWithHTTPClient creates a new DeleteRoutingSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRoutingSettingsParamsWithHTTPClient(client *http.Client) *DeleteRoutingSettingsParams {
	return &DeleteRoutingSettingsParams{
		HTTPClient: client,
	}
}

/*
DeleteRoutingSettingsParams contains all the parameters to send to the API endpoint

	for the delete routing settings operation.

	Typically these are written to a http.Request.
*/
type DeleteRoutingSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete routing settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRoutingSettingsParams) WithDefaults() *DeleteRoutingSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete routing settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRoutingSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete routing settings params
func (o *DeleteRoutingSettingsParams) WithTimeout(timeout time.Duration) *DeleteRoutingSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete routing settings params
func (o *DeleteRoutingSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete routing settings params
func (o *DeleteRoutingSettingsParams) WithContext(ctx context.Context) *DeleteRoutingSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete routing settings params
func (o *DeleteRoutingSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete routing settings params
func (o *DeleteRoutingSettingsParams) WithHTTPClient(client *http.Client) *DeleteRoutingSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete routing settings params
func (o *DeleteRoutingSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoutingSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
