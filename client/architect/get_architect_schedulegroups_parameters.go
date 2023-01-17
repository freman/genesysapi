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

// NewGetArchitectSchedulegroupsParams creates a new GetArchitectSchedulegroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetArchitectSchedulegroupsParams() *GetArchitectSchedulegroupsParams {
	return &GetArchitectSchedulegroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectSchedulegroupsParamsWithTimeout creates a new GetArchitectSchedulegroupsParams object
// with the ability to set a timeout on a request.
func NewGetArchitectSchedulegroupsParamsWithTimeout(timeout time.Duration) *GetArchitectSchedulegroupsParams {
	return &GetArchitectSchedulegroupsParams{
		timeout: timeout,
	}
}

// NewGetArchitectSchedulegroupsParamsWithContext creates a new GetArchitectSchedulegroupsParams object
// with the ability to set a context for a request.
func NewGetArchitectSchedulegroupsParamsWithContext(ctx context.Context) *GetArchitectSchedulegroupsParams {
	return &GetArchitectSchedulegroupsParams{
		Context: ctx,
	}
}

// NewGetArchitectSchedulegroupsParamsWithHTTPClient creates a new GetArchitectSchedulegroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetArchitectSchedulegroupsParamsWithHTTPClient(client *http.Client) *GetArchitectSchedulegroupsParams {
	return &GetArchitectSchedulegroupsParams{
		HTTPClient: client,
	}
}

/*
GetArchitectSchedulegroupsParams contains all the parameters to send to the API endpoint

	for the get architect schedulegroups operation.

	Typically these are written to a http.Request.
*/
type GetArchitectSchedulegroupsParams struct {

	/* DivisionID.

	   List of divisionIds on which to filter.
	*/
	DivisionID []string

	/* Name.

	   Name of the Schedule Group to filter by.
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

	/* ScheduleIds.

	   A comma-delimited list of Schedule IDs to filter by.
	*/
	ScheduleIds *string

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

// WithDefaults hydrates default values in the get architect schedulegroups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectSchedulegroupsParams) WithDefaults() *GetArchitectSchedulegroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get architect schedulegroups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectSchedulegroupsParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		sortByDefault = string("name")

		sortOrderDefault = string("ASC")
	)

	val := GetArchitectSchedulegroupsParams{
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

// WithTimeout adds the timeout to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) WithTimeout(timeout time.Duration) *GetArchitectSchedulegroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) WithContext(ctx context.Context) *GetArchitectSchedulegroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) WithHTTPClient(client *http.Client) *GetArchitectSchedulegroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDivisionID adds the divisionID to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) WithDivisionID(divisionID []string) *GetArchitectSchedulegroupsParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) SetDivisionID(divisionID []string) {
	o.DivisionID = divisionID
}

// WithName adds the name to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) WithName(name *string) *GetArchitectSchedulegroupsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) WithPageNumber(pageNumber *int32) *GetArchitectSchedulegroupsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) WithPageSize(pageSize *int32) *GetArchitectSchedulegroupsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithScheduleIds adds the scheduleIds to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) WithScheduleIds(scheduleIds *string) *GetArchitectSchedulegroupsParams {
	o.SetScheduleIds(scheduleIds)
	return o
}

// SetScheduleIds adds the scheduleIds to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) SetScheduleIds(scheduleIds *string) {
	o.ScheduleIds = scheduleIds
}

// WithSortBy adds the sortBy to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) WithSortBy(sortBy *string) *GetArchitectSchedulegroupsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) WithSortOrder(sortOrder *string) *GetArchitectSchedulegroupsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get architect schedulegroups params
func (o *GetArchitectSchedulegroupsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectSchedulegroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ScheduleIds != nil {

		// query param scheduleIds
		var qrScheduleIds string

		if o.ScheduleIds != nil {
			qrScheduleIds = *o.ScheduleIds
		}
		qScheduleIds := qrScheduleIds
		if qScheduleIds != "" {

			if err := r.SetQueryParam("scheduleIds", qScheduleIds); err != nil {
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

// bindParamGetArchitectSchedulegroups binds the parameter divisionId
func (o *GetArchitectSchedulegroupsParams) bindParamDivisionID(formats strfmt.Registry) []string {
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
