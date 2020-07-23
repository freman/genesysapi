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

// NewGetRoutingQueuesDivisionviewsAllParams creates a new GetRoutingQueuesDivisionviewsAllParams object
// with the default values initialized.
func NewGetRoutingQueuesDivisionviewsAllParams() *GetRoutingQueuesDivisionviewsAllParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("asc")
	)
	return &GetRoutingQueuesDivisionviewsAllParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingQueuesDivisionviewsAllParamsWithTimeout creates a new GetRoutingQueuesDivisionviewsAllParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoutingQueuesDivisionviewsAllParamsWithTimeout(timeout time.Duration) *GetRoutingQueuesDivisionviewsAllParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("asc")
	)
	return &GetRoutingQueuesDivisionviewsAllParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetRoutingQueuesDivisionviewsAllParamsWithContext creates a new GetRoutingQueuesDivisionviewsAllParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoutingQueuesDivisionviewsAllParamsWithContext(ctx context.Context) *GetRoutingQueuesDivisionviewsAllParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("asc")
	)
	return &GetRoutingQueuesDivisionviewsAllParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetRoutingQueuesDivisionviewsAllParamsWithHTTPClient creates a new GetRoutingQueuesDivisionviewsAllParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoutingQueuesDivisionviewsAllParamsWithHTTPClient(client *http.Client) *GetRoutingQueuesDivisionviewsAllParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("asc")
	)
	return &GetRoutingQueuesDivisionviewsAllParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetRoutingQueuesDivisionviewsAllParams contains all the parameters to send to the API endpoint
for the get routing queues divisionviews all operation typically these are written to a http.Request
*/
type GetRoutingQueuesDivisionviewsAllParams struct {

	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size [max value is 500]

	*/
	PageSize *int32
	/*SortBy
	  Sort by

	*/
	SortBy *string
	/*SortOrder
	  Sort order

	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) WithTimeout(timeout time.Duration) *GetRoutingQueuesDivisionviewsAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) WithContext(ctx context.Context) *GetRoutingQueuesDivisionviewsAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) WithHTTPClient(client *http.Client) *GetRoutingQueuesDivisionviewsAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) WithPageNumber(pageNumber *int32) *GetRoutingQueuesDivisionviewsAllParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) WithPageSize(pageSize *int32) *GetRoutingQueuesDivisionviewsAllParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) WithSortBy(sortBy *string) *GetRoutingQueuesDivisionviewsAllParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) WithSortOrder(sortOrder *string) *GetRoutingQueuesDivisionviewsAllParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get routing queues divisionviews all params
func (o *GetRoutingQueuesDivisionviewsAllParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingQueuesDivisionviewsAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}

	}

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string
		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {
			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
