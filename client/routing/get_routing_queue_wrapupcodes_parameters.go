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

// NewGetRoutingQueueWrapupcodesParams creates a new GetRoutingQueueWrapupcodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoutingQueueWrapupcodesParams() *GetRoutingQueueWrapupcodesParams {
	return &GetRoutingQueueWrapupcodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingQueueWrapupcodesParamsWithTimeout creates a new GetRoutingQueueWrapupcodesParams object
// with the ability to set a timeout on a request.
func NewGetRoutingQueueWrapupcodesParamsWithTimeout(timeout time.Duration) *GetRoutingQueueWrapupcodesParams {
	return &GetRoutingQueueWrapupcodesParams{
		timeout: timeout,
	}
}

// NewGetRoutingQueueWrapupcodesParamsWithContext creates a new GetRoutingQueueWrapupcodesParams object
// with the ability to set a context for a request.
func NewGetRoutingQueueWrapupcodesParamsWithContext(ctx context.Context) *GetRoutingQueueWrapupcodesParams {
	return &GetRoutingQueueWrapupcodesParams{
		Context: ctx,
	}
}

// NewGetRoutingQueueWrapupcodesParamsWithHTTPClient creates a new GetRoutingQueueWrapupcodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoutingQueueWrapupcodesParamsWithHTTPClient(client *http.Client) *GetRoutingQueueWrapupcodesParams {
	return &GetRoutingQueueWrapupcodesParams{
		HTTPClient: client,
	}
}

/*
GetRoutingQueueWrapupcodesParams contains all the parameters to send to the API endpoint

	for the get routing queue wrapupcodes operation.

	Typically these are written to a http.Request.
*/
type GetRoutingQueueWrapupcodesParams struct {

	/* PageNumber.

	   Page number

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   Page size

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	/* QueueID.

	   Queue ID
	*/
	QueueID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get routing queue wrapupcodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingQueueWrapupcodesParams) WithDefaults() *GetRoutingQueueWrapupcodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get routing queue wrapupcodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingQueueWrapupcodesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetRoutingQueueWrapupcodesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get routing queue wrapupcodes params
func (o *GetRoutingQueueWrapupcodesParams) WithTimeout(timeout time.Duration) *GetRoutingQueueWrapupcodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing queue wrapupcodes params
func (o *GetRoutingQueueWrapupcodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing queue wrapupcodes params
func (o *GetRoutingQueueWrapupcodesParams) WithContext(ctx context.Context) *GetRoutingQueueWrapupcodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing queue wrapupcodes params
func (o *GetRoutingQueueWrapupcodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing queue wrapupcodes params
func (o *GetRoutingQueueWrapupcodesParams) WithHTTPClient(client *http.Client) *GetRoutingQueueWrapupcodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing queue wrapupcodes params
func (o *GetRoutingQueueWrapupcodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get routing queue wrapupcodes params
func (o *GetRoutingQueueWrapupcodesParams) WithPageNumber(pageNumber *int32) *GetRoutingQueueWrapupcodesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get routing queue wrapupcodes params
func (o *GetRoutingQueueWrapupcodesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get routing queue wrapupcodes params
func (o *GetRoutingQueueWrapupcodesParams) WithPageSize(pageSize *int32) *GetRoutingQueueWrapupcodesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get routing queue wrapupcodes params
func (o *GetRoutingQueueWrapupcodesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithQueueID adds the queueID to the get routing queue wrapupcodes params
func (o *GetRoutingQueueWrapupcodesParams) WithQueueID(queueID string) *GetRoutingQueueWrapupcodesParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the get routing queue wrapupcodes params
func (o *GetRoutingQueueWrapupcodesParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingQueueWrapupcodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
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
