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

// NewGetArchitectSchedulesParams creates a new GetArchitectSchedulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetArchitectSchedulesParams() *GetArchitectSchedulesParams {
	return &GetArchitectSchedulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectSchedulesParamsWithTimeout creates a new GetArchitectSchedulesParams object
// with the ability to set a timeout on a request.
func NewGetArchitectSchedulesParamsWithTimeout(timeout time.Duration) *GetArchitectSchedulesParams {
	return &GetArchitectSchedulesParams{
		timeout: timeout,
	}
}

// NewGetArchitectSchedulesParamsWithContext creates a new GetArchitectSchedulesParams object
// with the ability to set a context for a request.
func NewGetArchitectSchedulesParamsWithContext(ctx context.Context) *GetArchitectSchedulesParams {
	return &GetArchitectSchedulesParams{
		Context: ctx,
	}
}

// NewGetArchitectSchedulesParamsWithHTTPClient creates a new GetArchitectSchedulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetArchitectSchedulesParamsWithHTTPClient(client *http.Client) *GetArchitectSchedulesParams {
	return &GetArchitectSchedulesParams{
		HTTPClient: client,
	}
}

/*
GetArchitectSchedulesParams contains all the parameters to send to the API endpoint

	for the get architect schedules operation.

	Typically these are written to a http.Request.
*/
type GetArchitectSchedulesParams struct {

	/* DivisionID.

	   List of divisionIds on which to filter.
	*/
	DivisionID []string

	/* Name.

	   Name of the Schedule to filter by.
	*/
	Name *string

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

	/* SortBy.

	   Sort by

	   Default: "name"
	*/
	SortBy *string

	/* SortOrder.

	   Sort order

	   Default: "ASC"
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get architect schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectSchedulesParams) WithDefaults() *GetArchitectSchedulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get architect schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectSchedulesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		sortByDefault = string("name")

		sortOrderDefault = string("ASC")
	)

	val := GetArchitectSchedulesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get architect schedules params
func (o *GetArchitectSchedulesParams) WithTimeout(timeout time.Duration) *GetArchitectSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect schedules params
func (o *GetArchitectSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect schedules params
func (o *GetArchitectSchedulesParams) WithContext(ctx context.Context) *GetArchitectSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect schedules params
func (o *GetArchitectSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect schedules params
func (o *GetArchitectSchedulesParams) WithHTTPClient(client *http.Client) *GetArchitectSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect schedules params
func (o *GetArchitectSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDivisionID adds the divisionID to the get architect schedules params
func (o *GetArchitectSchedulesParams) WithDivisionID(divisionID []string) *GetArchitectSchedulesParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the get architect schedules params
func (o *GetArchitectSchedulesParams) SetDivisionID(divisionID []string) {
	o.DivisionID = divisionID
}

// WithName adds the name to the get architect schedules params
func (o *GetArchitectSchedulesParams) WithName(name *string) *GetArchitectSchedulesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get architect schedules params
func (o *GetArchitectSchedulesParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get architect schedules params
func (o *GetArchitectSchedulesParams) WithPageNumber(pageNumber *int32) *GetArchitectSchedulesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get architect schedules params
func (o *GetArchitectSchedulesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get architect schedules params
func (o *GetArchitectSchedulesParams) WithPageSize(pageSize *int32) *GetArchitectSchedulesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get architect schedules params
func (o *GetArchitectSchedulesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get architect schedules params
func (o *GetArchitectSchedulesParams) WithSortBy(sortBy *string) *GetArchitectSchedulesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get architect schedules params
func (o *GetArchitectSchedulesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get architect schedules params
func (o *GetArchitectSchedulesParams) WithSortOrder(sortOrder *string) *GetArchitectSchedulesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get architect schedules params
func (o *GetArchitectSchedulesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DivisionID != nil {

		// binding items for divisionId
		joinedDivisionID := o.bindParamDivisionID(reg)

		// query array param divisionId
		if err := r.SetQueryParam("divisionId", joinedDivisionID...); err != nil {
			return err
		}
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

// bindParamGetArchitectSchedules binds the parameter divisionId
func (o *GetArchitectSchedulesParams) bindParamDivisionID(formats strfmt.Registry) []string {
	divisionIDIR := o.DivisionID

	var divisionIDIC []string
	for _, divisionIDIIR := range divisionIDIR { // explode []string

		divisionIDIIV := divisionIDIIR // string as string
		divisionIDIC = append(divisionIDIC, divisionIDIIV)
	}

	// items.CollectionFormat: "multi"
	divisionIDIS := swag.JoinByFormat(divisionIDIC, "multi")

	return divisionIDIS
}
