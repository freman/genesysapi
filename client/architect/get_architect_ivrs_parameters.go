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

// NewGetArchitectIvrsParams creates a new GetArchitectIvrsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetArchitectIvrsParams() *GetArchitectIvrsParams {
	return &GetArchitectIvrsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectIvrsParamsWithTimeout creates a new GetArchitectIvrsParams object
// with the ability to set a timeout on a request.
func NewGetArchitectIvrsParamsWithTimeout(timeout time.Duration) *GetArchitectIvrsParams {
	return &GetArchitectIvrsParams{
		timeout: timeout,
	}
}

// NewGetArchitectIvrsParamsWithContext creates a new GetArchitectIvrsParams object
// with the ability to set a context for a request.
func NewGetArchitectIvrsParamsWithContext(ctx context.Context) *GetArchitectIvrsParams {
	return &GetArchitectIvrsParams{
		Context: ctx,
	}
}

// NewGetArchitectIvrsParamsWithHTTPClient creates a new GetArchitectIvrsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetArchitectIvrsParamsWithHTTPClient(client *http.Client) *GetArchitectIvrsParams {
	return &GetArchitectIvrsParams{
		HTTPClient: client,
	}
}

/*
GetArchitectIvrsParams contains all the parameters to send to the API endpoint

	for the get architect ivrs operation.

	Typically these are written to a http.Request.
*/
type GetArchitectIvrsParams struct {

	/* Dnis.

	   The phone number of the IVR to filter by.
	*/
	Dnis *string

	/* Name.

	   Name of the IVR to filter by.
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

	/* ScheduleGroup.

	   The Schedule Group of the IVR to filter by.
	*/
	ScheduleGroup *string

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

// WithDefaults hydrates default values in the get architect ivrs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectIvrsParams) WithDefaults() *GetArchitectIvrsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get architect ivrs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectIvrsParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		sortByDefault = string("name")

		sortOrderDefault = string("ASC")
	)

	val := GetArchitectIvrsParams{
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

// WithTimeout adds the timeout to the get architect ivrs params
func (o *GetArchitectIvrsParams) WithTimeout(timeout time.Duration) *GetArchitectIvrsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect ivrs params
func (o *GetArchitectIvrsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect ivrs params
func (o *GetArchitectIvrsParams) WithContext(ctx context.Context) *GetArchitectIvrsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect ivrs params
func (o *GetArchitectIvrsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect ivrs params
func (o *GetArchitectIvrsParams) WithHTTPClient(client *http.Client) *GetArchitectIvrsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect ivrs params
func (o *GetArchitectIvrsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDnis adds the dnis to the get architect ivrs params
func (o *GetArchitectIvrsParams) WithDnis(dnis *string) *GetArchitectIvrsParams {
	o.SetDnis(dnis)
	return o
}

// SetDnis adds the dnis to the get architect ivrs params
func (o *GetArchitectIvrsParams) SetDnis(dnis *string) {
	o.Dnis = dnis
}

// WithName adds the name to the get architect ivrs params
func (o *GetArchitectIvrsParams) WithName(name *string) *GetArchitectIvrsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get architect ivrs params
func (o *GetArchitectIvrsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get architect ivrs params
func (o *GetArchitectIvrsParams) WithPageNumber(pageNumber *int32) *GetArchitectIvrsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get architect ivrs params
func (o *GetArchitectIvrsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get architect ivrs params
func (o *GetArchitectIvrsParams) WithPageSize(pageSize *int32) *GetArchitectIvrsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get architect ivrs params
func (o *GetArchitectIvrsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithScheduleGroup adds the scheduleGroup to the get architect ivrs params
func (o *GetArchitectIvrsParams) WithScheduleGroup(scheduleGroup *string) *GetArchitectIvrsParams {
	o.SetScheduleGroup(scheduleGroup)
	return o
}

// SetScheduleGroup adds the scheduleGroup to the get architect ivrs params
func (o *GetArchitectIvrsParams) SetScheduleGroup(scheduleGroup *string) {
	o.ScheduleGroup = scheduleGroup
}

// WithSortBy adds the sortBy to the get architect ivrs params
func (o *GetArchitectIvrsParams) WithSortBy(sortBy *string) *GetArchitectIvrsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get architect ivrs params
func (o *GetArchitectIvrsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get architect ivrs params
func (o *GetArchitectIvrsParams) WithSortOrder(sortOrder *string) *GetArchitectIvrsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get architect ivrs params
func (o *GetArchitectIvrsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectIvrsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Dnis != nil {

		// query param dnis
		var qrDnis string

		if o.Dnis != nil {
			qrDnis = *o.Dnis
		}
		qDnis := qrDnis
		if qDnis != "" {

			if err := r.SetQueryParam("dnis", qDnis); err != nil {
				return err
			}
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

	if o.ScheduleGroup != nil {

		// query param scheduleGroup
		var qrScheduleGroup string

		if o.ScheduleGroup != nil {
			qrScheduleGroup = *o.ScheduleGroup
		}
		qScheduleGroup := qrScheduleGroup
		if qScheduleGroup != "" {

			if err := r.SetQueryParam("scheduleGroup", qScheduleGroup); err != nil {
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
