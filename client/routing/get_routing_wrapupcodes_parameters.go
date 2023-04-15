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

// NewGetRoutingWrapupcodesParams creates a new GetRoutingWrapupcodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoutingWrapupcodesParams() *GetRoutingWrapupcodesParams {
	return &GetRoutingWrapupcodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingWrapupcodesParamsWithTimeout creates a new GetRoutingWrapupcodesParams object
// with the ability to set a timeout on a request.
func NewGetRoutingWrapupcodesParamsWithTimeout(timeout time.Duration) *GetRoutingWrapupcodesParams {
	return &GetRoutingWrapupcodesParams{
		timeout: timeout,
	}
}

// NewGetRoutingWrapupcodesParamsWithContext creates a new GetRoutingWrapupcodesParams object
// with the ability to set a context for a request.
func NewGetRoutingWrapupcodesParamsWithContext(ctx context.Context) *GetRoutingWrapupcodesParams {
	return &GetRoutingWrapupcodesParams{
		Context: ctx,
	}
}

// NewGetRoutingWrapupcodesParamsWithHTTPClient creates a new GetRoutingWrapupcodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoutingWrapupcodesParamsWithHTTPClient(client *http.Client) *GetRoutingWrapupcodesParams {
	return &GetRoutingWrapupcodesParams{
		HTTPClient: client,
	}
}

/*
GetRoutingWrapupcodesParams contains all the parameters to send to the API endpoint

	for the get routing wrapupcodes operation.

	Typically these are written to a http.Request.
*/
type GetRoutingWrapupcodesParams struct {

	/* DivisionID.

	   Filter by division ID(s)
	*/
	DivisionID []string

	/* ID.

	   Filter by wrapup code ID(s)
	*/
	ID []string

	/* Name.

	   Wrapup code's name ('Sort by' param is ignored unless this field is provided)
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

	   Default: "ascending"
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get routing wrapupcodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingWrapupcodesParams) WithDefaults() *GetRoutingWrapupcodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get routing wrapupcodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingWrapupcodesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		sortByDefault = string("name")

		sortOrderDefault = string("ascending")
	)

	val := GetRoutingWrapupcodesParams{
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

// WithTimeout adds the timeout to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) WithTimeout(timeout time.Duration) *GetRoutingWrapupcodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) WithContext(ctx context.Context) *GetRoutingWrapupcodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) WithHTTPClient(client *http.Client) *GetRoutingWrapupcodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDivisionID adds the divisionID to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) WithDivisionID(divisionID []string) *GetRoutingWrapupcodesParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) SetDivisionID(divisionID []string) {
	o.DivisionID = divisionID
}

// WithID adds the id to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) WithID(id []string) *GetRoutingWrapupcodesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) WithName(name *string) *GetRoutingWrapupcodesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) WithPageNumber(pageNumber *int32) *GetRoutingWrapupcodesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) WithPageSize(pageSize *int32) *GetRoutingWrapupcodesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) WithSortBy(sortBy *string) *GetRoutingWrapupcodesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) WithSortOrder(sortOrder *string) *GetRoutingWrapupcodesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get routing wrapupcodes params
func (o *GetRoutingWrapupcodesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingWrapupcodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
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

// bindParamGetRoutingWrapupcodes binds the parameter divisionId
func (o *GetRoutingWrapupcodesParams) bindParamDivisionID(formats strfmt.Registry) []string {
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

// bindParamGetRoutingWrapupcodes binds the parameter id
func (o *GetRoutingWrapupcodesParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}
