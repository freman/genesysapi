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

// NewGetRoutingQueuesParams creates a new GetRoutingQueuesParams object
// with the default values initialized.
func NewGetRoutingQueuesParams() *GetRoutingQueuesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("asc")
	)
	return &GetRoutingQueuesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingQueuesParamsWithTimeout creates a new GetRoutingQueuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoutingQueuesParamsWithTimeout(timeout time.Duration) *GetRoutingQueuesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("asc")
	)
	return &GetRoutingQueuesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetRoutingQueuesParamsWithContext creates a new GetRoutingQueuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoutingQueuesParamsWithContext(ctx context.Context) *GetRoutingQueuesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("asc")
	)
	return &GetRoutingQueuesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetRoutingQueuesParamsWithHTTPClient creates a new GetRoutingQueuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoutingQueuesParamsWithHTTPClient(client *http.Client) *GetRoutingQueuesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("asc")
	)
	return &GetRoutingQueuesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetRoutingQueuesParams contains all the parameters to send to the API endpoint
for the get routing queues operation typically these are written to a http.Request
*/
type GetRoutingQueuesParams struct {

	/*DivisionID
	  Filter by queue division ID(s)

	*/
	DivisionID []string
	/*ID
	  Filter by queue ID(s)

	*/
	ID []string
	/*Name
	  Filter by queue name

	*/
	Name *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*SortOrder
	  Note: results are sorted by name.

	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routing queues params
func (o *GetRoutingQueuesParams) WithTimeout(timeout time.Duration) *GetRoutingQueuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing queues params
func (o *GetRoutingQueuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing queues params
func (o *GetRoutingQueuesParams) WithContext(ctx context.Context) *GetRoutingQueuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing queues params
func (o *GetRoutingQueuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing queues params
func (o *GetRoutingQueuesParams) WithHTTPClient(client *http.Client) *GetRoutingQueuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing queues params
func (o *GetRoutingQueuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDivisionID adds the divisionID to the get routing queues params
func (o *GetRoutingQueuesParams) WithDivisionID(divisionID []string) *GetRoutingQueuesParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the get routing queues params
func (o *GetRoutingQueuesParams) SetDivisionID(divisionID []string) {
	o.DivisionID = divisionID
}

// WithID adds the id to the get routing queues params
func (o *GetRoutingQueuesParams) WithID(id []string) *GetRoutingQueuesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get routing queues params
func (o *GetRoutingQueuesParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the get routing queues params
func (o *GetRoutingQueuesParams) WithName(name *string) *GetRoutingQueuesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get routing queues params
func (o *GetRoutingQueuesParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get routing queues params
func (o *GetRoutingQueuesParams) WithPageNumber(pageNumber *int32) *GetRoutingQueuesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get routing queues params
func (o *GetRoutingQueuesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get routing queues params
func (o *GetRoutingQueuesParams) WithPageSize(pageSize *int32) *GetRoutingQueuesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get routing queues params
func (o *GetRoutingQueuesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortOrder adds the sortOrder to the get routing queues params
func (o *GetRoutingQueuesParams) WithSortOrder(sortOrder *string) *GetRoutingQueuesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get routing queues params
func (o *GetRoutingQueuesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingQueuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDivisionID := o.DivisionID

	joinedDivisionID := swag.JoinByFormat(valuesDivisionID, "multi")
	// query array param divisionId
	if err := r.SetQueryParam("divisionId", joinedDivisionID...); err != nil {
		return err
	}

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

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
