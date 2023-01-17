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
	"github.com/go-openapi/swag"
)

// NewGetRoutingPredictorsParams creates a new GetRoutingPredictorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoutingPredictorsParams() *GetRoutingPredictorsParams {
	return &GetRoutingPredictorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingPredictorsParamsWithTimeout creates a new GetRoutingPredictorsParams object
// with the ability to set a timeout on a request.
func NewGetRoutingPredictorsParamsWithTimeout(timeout time.Duration) *GetRoutingPredictorsParams {
	return &GetRoutingPredictorsParams{
		timeout: timeout,
	}
}

// NewGetRoutingPredictorsParamsWithContext creates a new GetRoutingPredictorsParams object
// with the ability to set a context for a request.
func NewGetRoutingPredictorsParamsWithContext(ctx context.Context) *GetRoutingPredictorsParams {
	return &GetRoutingPredictorsParams{
		Context: ctx,
	}
}

// NewGetRoutingPredictorsParamsWithHTTPClient creates a new GetRoutingPredictorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoutingPredictorsParamsWithHTTPClient(client *http.Client) *GetRoutingPredictorsParams {
	return &GetRoutingPredictorsParams{
		HTTPClient: client,
	}
}

/*
GetRoutingPredictorsParams contains all the parameters to send to the API endpoint

	for the get routing predictors operation.

	Typically these are written to a http.Request.
*/
type GetRoutingPredictorsParams struct {

	/* After.

	   The cursor that points to the end of the set of entities that has been returned.
	*/
	After *string

	/* Before.

	   The cursor that points to the start of the set of entities that has been returned.
	*/
	Before *string

	/* Limit.

	   Number of entities to return. Maximum of 200. Deprecated in favour of pageSize
	*/
	Limit *string

	/* PageSize.

	   Number of entities to return. Maximum of 200.
	*/
	PageSize *string

	/* QueueID.

	   Comma-separated list of queue Ids to filter by.
	*/
	QueueID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get routing predictors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingPredictorsParams) WithDefaults() *GetRoutingPredictorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get routing predictors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingPredictorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get routing predictors params
func (o *GetRoutingPredictorsParams) WithTimeout(timeout time.Duration) *GetRoutingPredictorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing predictors params
func (o *GetRoutingPredictorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing predictors params
func (o *GetRoutingPredictorsParams) WithContext(ctx context.Context) *GetRoutingPredictorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing predictors params
func (o *GetRoutingPredictorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing predictors params
func (o *GetRoutingPredictorsParams) WithHTTPClient(client *http.Client) *GetRoutingPredictorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing predictors params
func (o *GetRoutingPredictorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get routing predictors params
func (o *GetRoutingPredictorsParams) WithAfter(after *string) *GetRoutingPredictorsParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get routing predictors params
func (o *GetRoutingPredictorsParams) SetAfter(after *string) {
	o.After = after
}

// WithBefore adds the before to the get routing predictors params
func (o *GetRoutingPredictorsParams) WithBefore(before *string) *GetRoutingPredictorsParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the get routing predictors params
func (o *GetRoutingPredictorsParams) SetBefore(before *string) {
	o.Before = before
}

// WithLimit adds the limit to the get routing predictors params
func (o *GetRoutingPredictorsParams) WithLimit(limit *string) *GetRoutingPredictorsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get routing predictors params
func (o *GetRoutingPredictorsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithPageSize adds the pageSize to the get routing predictors params
func (o *GetRoutingPredictorsParams) WithPageSize(pageSize *string) *GetRoutingPredictorsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get routing predictors params
func (o *GetRoutingPredictorsParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WithQueueID adds the queueID to the get routing predictors params
func (o *GetRoutingPredictorsParams) WithQueueID(queueID []string) *GetRoutingPredictorsParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the get routing predictors params
func (o *GetRoutingPredictorsParams) SetQueueID(queueID []string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingPredictorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if o.Before != nil {

		// query param before
		var qrBefore string

		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := qrBefore
		if qBefore != "" {

			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize string

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := qrPageSize
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.QueueID != nil {

		// binding items for queueId
		joinedQueueID := o.bindParamQueueID(reg)

		// query array param queueId
		if err := r.SetQueryParam("queueId", joinedQueueID...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetRoutingPredictors binds the parameter queueId
func (o *GetRoutingPredictorsParams) bindParamQueueID(formats strfmt.Registry) []string {
	queueIDIR := o.QueueID

	var queueIDIC []string
	for _, queueIDIIR := range queueIDIR { // explode []string

		queueIDIIV := queueIDIIR // string as string
		queueIDIC = append(queueIDIC, queueIDIIV)
	}

	// items.CollectionFormat: "multi"
	queueIDIS := swag.JoinByFormat(queueIDIC, "multi")

	return queueIDIS
}
