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

// NewGetRoutingQueueComparisonperiodParams creates a new GetRoutingQueueComparisonperiodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoutingQueueComparisonperiodParams() *GetRoutingQueueComparisonperiodParams {
	return &GetRoutingQueueComparisonperiodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingQueueComparisonperiodParamsWithTimeout creates a new GetRoutingQueueComparisonperiodParams object
// with the ability to set a timeout on a request.
func NewGetRoutingQueueComparisonperiodParamsWithTimeout(timeout time.Duration) *GetRoutingQueueComparisonperiodParams {
	return &GetRoutingQueueComparisonperiodParams{
		timeout: timeout,
	}
}

// NewGetRoutingQueueComparisonperiodParamsWithContext creates a new GetRoutingQueueComparisonperiodParams object
// with the ability to set a context for a request.
func NewGetRoutingQueueComparisonperiodParamsWithContext(ctx context.Context) *GetRoutingQueueComparisonperiodParams {
	return &GetRoutingQueueComparisonperiodParams{
		Context: ctx,
	}
}

// NewGetRoutingQueueComparisonperiodParamsWithHTTPClient creates a new GetRoutingQueueComparisonperiodParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoutingQueueComparisonperiodParamsWithHTTPClient(client *http.Client) *GetRoutingQueueComparisonperiodParams {
	return &GetRoutingQueueComparisonperiodParams{
		HTTPClient: client,
	}
}

/*
GetRoutingQueueComparisonperiodParams contains all the parameters to send to the API endpoint

	for the get routing queue comparisonperiod operation.

	Typically these are written to a http.Request.
*/
type GetRoutingQueueComparisonperiodParams struct {

	/* ComparisonPeriodID.

	   ComparisonPeriod id
	*/
	ComparisonPeriodID string

	/* QueueID.

	   Queue id
	*/
	QueueID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get routing queue comparisonperiod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingQueueComparisonperiodParams) WithDefaults() *GetRoutingQueueComparisonperiodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get routing queue comparisonperiod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingQueueComparisonperiodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get routing queue comparisonperiod params
func (o *GetRoutingQueueComparisonperiodParams) WithTimeout(timeout time.Duration) *GetRoutingQueueComparisonperiodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing queue comparisonperiod params
func (o *GetRoutingQueueComparisonperiodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing queue comparisonperiod params
func (o *GetRoutingQueueComparisonperiodParams) WithContext(ctx context.Context) *GetRoutingQueueComparisonperiodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing queue comparisonperiod params
func (o *GetRoutingQueueComparisonperiodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing queue comparisonperiod params
func (o *GetRoutingQueueComparisonperiodParams) WithHTTPClient(client *http.Client) *GetRoutingQueueComparisonperiodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing queue comparisonperiod params
func (o *GetRoutingQueueComparisonperiodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComparisonPeriodID adds the comparisonPeriodID to the get routing queue comparisonperiod params
func (o *GetRoutingQueueComparisonperiodParams) WithComparisonPeriodID(comparisonPeriodID string) *GetRoutingQueueComparisonperiodParams {
	o.SetComparisonPeriodID(comparisonPeriodID)
	return o
}

// SetComparisonPeriodID adds the comparisonPeriodId to the get routing queue comparisonperiod params
func (o *GetRoutingQueueComparisonperiodParams) SetComparisonPeriodID(comparisonPeriodID string) {
	o.ComparisonPeriodID = comparisonPeriodID
}

// WithQueueID adds the queueID to the get routing queue comparisonperiod params
func (o *GetRoutingQueueComparisonperiodParams) WithQueueID(queueID string) *GetRoutingQueueComparisonperiodParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the get routing queue comparisonperiod params
func (o *GetRoutingQueueComparisonperiodParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingQueueComparisonperiodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param comparisonPeriodId
	if err := r.SetPathParam("comparisonPeriodId", o.ComparisonPeriodID); err != nil {
		return err
	}

	// path param queueId
	if err := r.SetPathParam("queueId", o.QueueID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
