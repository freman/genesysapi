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

// NewGetRoutingQueueComparisonperiodsParams creates a new GetRoutingQueueComparisonperiodsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoutingQueueComparisonperiodsParams() *GetRoutingQueueComparisonperiodsParams {
	return &GetRoutingQueueComparisonperiodsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingQueueComparisonperiodsParamsWithTimeout creates a new GetRoutingQueueComparisonperiodsParams object
// with the ability to set a timeout on a request.
func NewGetRoutingQueueComparisonperiodsParamsWithTimeout(timeout time.Duration) *GetRoutingQueueComparisonperiodsParams {
	return &GetRoutingQueueComparisonperiodsParams{
		timeout: timeout,
	}
}

// NewGetRoutingQueueComparisonperiodsParamsWithContext creates a new GetRoutingQueueComparisonperiodsParams object
// with the ability to set a context for a request.
func NewGetRoutingQueueComparisonperiodsParamsWithContext(ctx context.Context) *GetRoutingQueueComparisonperiodsParams {
	return &GetRoutingQueueComparisonperiodsParams{
		Context: ctx,
	}
}

// NewGetRoutingQueueComparisonperiodsParamsWithHTTPClient creates a new GetRoutingQueueComparisonperiodsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoutingQueueComparisonperiodsParamsWithHTTPClient(client *http.Client) *GetRoutingQueueComparisonperiodsParams {
	return &GetRoutingQueueComparisonperiodsParams{
		HTTPClient: client,
	}
}

/*
GetRoutingQueueComparisonperiodsParams contains all the parameters to send to the API endpoint

	for the get routing queue comparisonperiods operation.

	Typically these are written to a http.Request.
*/
type GetRoutingQueueComparisonperiodsParams struct {

	/* QueueID.

	   Queue id
	*/
	QueueID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get routing queue comparisonperiods params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingQueueComparisonperiodsParams) WithDefaults() *GetRoutingQueueComparisonperiodsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get routing queue comparisonperiods params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingQueueComparisonperiodsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get routing queue comparisonperiods params
func (o *GetRoutingQueueComparisonperiodsParams) WithTimeout(timeout time.Duration) *GetRoutingQueueComparisonperiodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing queue comparisonperiods params
func (o *GetRoutingQueueComparisonperiodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing queue comparisonperiods params
func (o *GetRoutingQueueComparisonperiodsParams) WithContext(ctx context.Context) *GetRoutingQueueComparisonperiodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing queue comparisonperiods params
func (o *GetRoutingQueueComparisonperiodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing queue comparisonperiods params
func (o *GetRoutingQueueComparisonperiodsParams) WithHTTPClient(client *http.Client) *GetRoutingQueueComparisonperiodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing queue comparisonperiods params
func (o *GetRoutingQueueComparisonperiodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQueueID adds the queueID to the get routing queue comparisonperiods params
func (o *GetRoutingQueueComparisonperiodsParams) WithQueueID(queueID string) *GetRoutingQueueComparisonperiodsParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the get routing queue comparisonperiods params
func (o *GetRoutingQueueComparisonperiodsParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingQueueComparisonperiodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param queueId
	if err := r.SetPathParam("queueId", o.QueueID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
