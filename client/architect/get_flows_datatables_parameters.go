// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewGetFlowsDatatablesParams creates a new GetFlowsDatatablesParams object
// with the default values initialized.
func NewGetFlowsDatatablesParams() *GetFlowsDatatablesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("id")
		sortOrderDefault  = string("ascending")
	)
	return &GetFlowsDatatablesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowsDatatablesParamsWithTimeout creates a new GetFlowsDatatablesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFlowsDatatablesParamsWithTimeout(timeout time.Duration) *GetFlowsDatatablesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("id")
		sortOrderDefault  = string("ascending")
	)
	return &GetFlowsDatatablesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetFlowsDatatablesParamsWithContext creates a new GetFlowsDatatablesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFlowsDatatablesParamsWithContext(ctx context.Context) *GetFlowsDatatablesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("id")
		sortOrderDefault  = string("ascending")
	)
	return &GetFlowsDatatablesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetFlowsDatatablesParamsWithHTTPClient creates a new GetFlowsDatatablesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFlowsDatatablesParamsWithHTTPClient(client *http.Client) *GetFlowsDatatablesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("id")
		sortOrderDefault  = string("ascending")
	)
	return &GetFlowsDatatablesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetFlowsDatatablesParams contains all the parameters to send to the API endpoint
for the get flows datatables operation typically these are written to a http.Request
*/
type GetFlowsDatatablesParams struct {

	/*DivisionID
	  division ID(s)

	*/
	DivisionID []string
	/*Expand
	  Expand instructions for the result

	*/
	Expand *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

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

// WithTimeout adds the timeout to the get flows datatables params
func (o *GetFlowsDatatablesParams) WithTimeout(timeout time.Duration) *GetFlowsDatatablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flows datatables params
func (o *GetFlowsDatatablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flows datatables params
func (o *GetFlowsDatatablesParams) WithContext(ctx context.Context) *GetFlowsDatatablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flows datatables params
func (o *GetFlowsDatatablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flows datatables params
func (o *GetFlowsDatatablesParams) WithHTTPClient(client *http.Client) *GetFlowsDatatablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flows datatables params
func (o *GetFlowsDatatablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDivisionID adds the divisionID to the get flows datatables params
func (o *GetFlowsDatatablesParams) WithDivisionID(divisionID []string) *GetFlowsDatatablesParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the get flows datatables params
func (o *GetFlowsDatatablesParams) SetDivisionID(divisionID []string) {
	o.DivisionID = divisionID
}

// WithExpand adds the expand to the get flows datatables params
func (o *GetFlowsDatatablesParams) WithExpand(expand *string) *GetFlowsDatatablesParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get flows datatables params
func (o *GetFlowsDatatablesParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithPageNumber adds the pageNumber to the get flows datatables params
func (o *GetFlowsDatatablesParams) WithPageNumber(pageNumber *int32) *GetFlowsDatatablesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get flows datatables params
func (o *GetFlowsDatatablesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get flows datatables params
func (o *GetFlowsDatatablesParams) WithPageSize(pageSize *int32) *GetFlowsDatatablesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get flows datatables params
func (o *GetFlowsDatatablesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get flows datatables params
func (o *GetFlowsDatatablesParams) WithSortBy(sortBy *string) *GetFlowsDatatablesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get flows datatables params
func (o *GetFlowsDatatablesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get flows datatables params
func (o *GetFlowsDatatablesParams) WithSortOrder(sortOrder *string) *GetFlowsDatatablesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get flows datatables params
func (o *GetFlowsDatatablesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowsDatatablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Expand != nil {

		// query param expand
		var qrExpand string
		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {
			if err := r.SetQueryParam("expand", qExpand); err != nil {
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
