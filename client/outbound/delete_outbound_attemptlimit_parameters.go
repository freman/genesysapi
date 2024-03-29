// Code generated by go-swagger; DO NOT EDIT.

package outbound

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

// NewDeleteOutboundAttemptlimitParams creates a new DeleteOutboundAttemptlimitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOutboundAttemptlimitParams() *DeleteOutboundAttemptlimitParams {
	return &DeleteOutboundAttemptlimitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOutboundAttemptlimitParamsWithTimeout creates a new DeleteOutboundAttemptlimitParams object
// with the ability to set a timeout on a request.
func NewDeleteOutboundAttemptlimitParamsWithTimeout(timeout time.Duration) *DeleteOutboundAttemptlimitParams {
	return &DeleteOutboundAttemptlimitParams{
		timeout: timeout,
	}
}

// NewDeleteOutboundAttemptlimitParamsWithContext creates a new DeleteOutboundAttemptlimitParams object
// with the ability to set a context for a request.
func NewDeleteOutboundAttemptlimitParamsWithContext(ctx context.Context) *DeleteOutboundAttemptlimitParams {
	return &DeleteOutboundAttemptlimitParams{
		Context: ctx,
	}
}

// NewDeleteOutboundAttemptlimitParamsWithHTTPClient creates a new DeleteOutboundAttemptlimitParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOutboundAttemptlimitParamsWithHTTPClient(client *http.Client) *DeleteOutboundAttemptlimitParams {
	return &DeleteOutboundAttemptlimitParams{
		HTTPClient: client,
	}
}

/*
DeleteOutboundAttemptlimitParams contains all the parameters to send to the API endpoint

	for the delete outbound attemptlimit operation.

	Typically these are written to a http.Request.
*/
type DeleteOutboundAttemptlimitParams struct {

	/* AttemptLimitsID.

	   Attempt limits ID
	*/
	AttemptLimitsID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete outbound attemptlimit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOutboundAttemptlimitParams) WithDefaults() *DeleteOutboundAttemptlimitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete outbound attemptlimit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOutboundAttemptlimitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete outbound attemptlimit params
func (o *DeleteOutboundAttemptlimitParams) WithTimeout(timeout time.Duration) *DeleteOutboundAttemptlimitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete outbound attemptlimit params
func (o *DeleteOutboundAttemptlimitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete outbound attemptlimit params
func (o *DeleteOutboundAttemptlimitParams) WithContext(ctx context.Context) *DeleteOutboundAttemptlimitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete outbound attemptlimit params
func (o *DeleteOutboundAttemptlimitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete outbound attemptlimit params
func (o *DeleteOutboundAttemptlimitParams) WithHTTPClient(client *http.Client) *DeleteOutboundAttemptlimitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete outbound attemptlimit params
func (o *DeleteOutboundAttemptlimitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttemptLimitsID adds the attemptLimitsID to the delete outbound attemptlimit params
func (o *DeleteOutboundAttemptlimitParams) WithAttemptLimitsID(attemptLimitsID string) *DeleteOutboundAttemptlimitParams {
	o.SetAttemptLimitsID(attemptLimitsID)
	return o
}

// SetAttemptLimitsID adds the attemptLimitsId to the delete outbound attemptlimit params
func (o *DeleteOutboundAttemptlimitParams) SetAttemptLimitsID(attemptLimitsID string) {
	o.AttemptLimitsID = attemptLimitsID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOutboundAttemptlimitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attemptLimitsId
	if err := r.SetPathParam("attemptLimitsId", o.AttemptLimitsID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
