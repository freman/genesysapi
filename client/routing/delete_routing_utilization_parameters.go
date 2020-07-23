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

// NewDeleteRoutingUtilizationParams creates a new DeleteRoutingUtilizationParams object
// with the default values initialized.
func NewDeleteRoutingUtilizationParams() *DeleteRoutingUtilizationParams {

	return &DeleteRoutingUtilizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoutingUtilizationParamsWithTimeout creates a new DeleteRoutingUtilizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRoutingUtilizationParamsWithTimeout(timeout time.Duration) *DeleteRoutingUtilizationParams {

	return &DeleteRoutingUtilizationParams{

		timeout: timeout,
	}
}

// NewDeleteRoutingUtilizationParamsWithContext creates a new DeleteRoutingUtilizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRoutingUtilizationParamsWithContext(ctx context.Context) *DeleteRoutingUtilizationParams {

	return &DeleteRoutingUtilizationParams{

		Context: ctx,
	}
}

// NewDeleteRoutingUtilizationParamsWithHTTPClient creates a new DeleteRoutingUtilizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRoutingUtilizationParamsWithHTTPClient(client *http.Client) *DeleteRoutingUtilizationParams {

	return &DeleteRoutingUtilizationParams{
		HTTPClient: client,
	}
}

/*DeleteRoutingUtilizationParams contains all the parameters to send to the API endpoint
for the delete routing utilization operation typically these are written to a http.Request
*/
type DeleteRoutingUtilizationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete routing utilization params
func (o *DeleteRoutingUtilizationParams) WithTimeout(timeout time.Duration) *DeleteRoutingUtilizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete routing utilization params
func (o *DeleteRoutingUtilizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete routing utilization params
func (o *DeleteRoutingUtilizationParams) WithContext(ctx context.Context) *DeleteRoutingUtilizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete routing utilization params
func (o *DeleteRoutingUtilizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete routing utilization params
func (o *DeleteRoutingUtilizationParams) WithHTTPClient(client *http.Client) *DeleteRoutingUtilizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete routing utilization params
func (o *DeleteRoutingUtilizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoutingUtilizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
