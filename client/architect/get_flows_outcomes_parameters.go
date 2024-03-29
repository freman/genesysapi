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

// NewGetFlowsOutcomesParams creates a new GetFlowsOutcomesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFlowsOutcomesParams() *GetFlowsOutcomesParams {
	return &GetFlowsOutcomesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowsOutcomesParamsWithTimeout creates a new GetFlowsOutcomesParams object
// with the ability to set a timeout on a request.
func NewGetFlowsOutcomesParamsWithTimeout(timeout time.Duration) *GetFlowsOutcomesParams {
	return &GetFlowsOutcomesParams{
		timeout: timeout,
	}
}

// NewGetFlowsOutcomesParamsWithContext creates a new GetFlowsOutcomesParams object
// with the ability to set a context for a request.
func NewGetFlowsOutcomesParamsWithContext(ctx context.Context) *GetFlowsOutcomesParams {
	return &GetFlowsOutcomesParams{
		Context: ctx,
	}
}

// NewGetFlowsOutcomesParamsWithHTTPClient creates a new GetFlowsOutcomesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFlowsOutcomesParamsWithHTTPClient(client *http.Client) *GetFlowsOutcomesParams {
	return &GetFlowsOutcomesParams{
		HTTPClient: client,
	}
}

/*
GetFlowsOutcomesParams contains all the parameters to send to the API endpoint

	for the get flows outcomes operation.

	Typically these are written to a http.Request.
*/
type GetFlowsOutcomesParams struct {

	/* Description.

	   Description
	*/
	Description *string

	/* DivisionID.

	   division ID(s)
	*/
	DivisionID []string

	/* ID.

	   ID
	*/
	ID []string

	/* Name.

	   Name
	*/
	Name *string

	/* NameOrDescription.

	   Name or description
	*/
	NameOrDescription *string

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

	   Default: "id"
	*/
	SortBy *string

	/* SortOrder.

	   Sort order

	   Default: "asc"
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get flows outcomes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowsOutcomesParams) WithDefaults() *GetFlowsOutcomesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get flows outcomes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowsOutcomesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		sortByDefault = string("id")

		sortOrderDefault = string("asc")
	)

	val := GetFlowsOutcomesParams{
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

// WithTimeout adds the timeout to the get flows outcomes params
func (o *GetFlowsOutcomesParams) WithTimeout(timeout time.Duration) *GetFlowsOutcomesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flows outcomes params
func (o *GetFlowsOutcomesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flows outcomes params
func (o *GetFlowsOutcomesParams) WithContext(ctx context.Context) *GetFlowsOutcomesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flows outcomes params
func (o *GetFlowsOutcomesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flows outcomes params
func (o *GetFlowsOutcomesParams) WithHTTPClient(client *http.Client) *GetFlowsOutcomesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flows outcomes params
func (o *GetFlowsOutcomesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the get flows outcomes params
func (o *GetFlowsOutcomesParams) WithDescription(description *string) *GetFlowsOutcomesParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the get flows outcomes params
func (o *GetFlowsOutcomesParams) SetDescription(description *string) {
	o.Description = description
}

// WithDivisionID adds the divisionID to the get flows outcomes params
func (o *GetFlowsOutcomesParams) WithDivisionID(divisionID []string) *GetFlowsOutcomesParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the get flows outcomes params
func (o *GetFlowsOutcomesParams) SetDivisionID(divisionID []string) {
	o.DivisionID = divisionID
}

// WithID adds the id to the get flows outcomes params
func (o *GetFlowsOutcomesParams) WithID(id []string) *GetFlowsOutcomesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get flows outcomes params
func (o *GetFlowsOutcomesParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the get flows outcomes params
func (o *GetFlowsOutcomesParams) WithName(name *string) *GetFlowsOutcomesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get flows outcomes params
func (o *GetFlowsOutcomesParams) SetName(name *string) {
	o.Name = name
}

// WithNameOrDescription adds the nameOrDescription to the get flows outcomes params
func (o *GetFlowsOutcomesParams) WithNameOrDescription(nameOrDescription *string) *GetFlowsOutcomesParams {
	o.SetNameOrDescription(nameOrDescription)
	return o
}

// SetNameOrDescription adds the nameOrDescription to the get flows outcomes params
func (o *GetFlowsOutcomesParams) SetNameOrDescription(nameOrDescription *string) {
	o.NameOrDescription = nameOrDescription
}

// WithPageNumber adds the pageNumber to the get flows outcomes params
func (o *GetFlowsOutcomesParams) WithPageNumber(pageNumber *int32) *GetFlowsOutcomesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get flows outcomes params
func (o *GetFlowsOutcomesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get flows outcomes params
func (o *GetFlowsOutcomesParams) WithPageSize(pageSize *int32) *GetFlowsOutcomesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get flows outcomes params
func (o *GetFlowsOutcomesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get flows outcomes params
func (o *GetFlowsOutcomesParams) WithSortBy(sortBy *string) *GetFlowsOutcomesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get flows outcomes params
func (o *GetFlowsOutcomesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get flows outcomes params
func (o *GetFlowsOutcomesParams) WithSortOrder(sortOrder *string) *GetFlowsOutcomesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get flows outcomes params
func (o *GetFlowsOutcomesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowsOutcomesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

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

	if o.NameOrDescription != nil {

		// query param nameOrDescription
		var qrNameOrDescription string

		if o.NameOrDescription != nil {
			qrNameOrDescription = *o.NameOrDescription
		}
		qNameOrDescription := qrNameOrDescription
		if qNameOrDescription != "" {

			if err := r.SetQueryParam("nameOrDescription", qNameOrDescription); err != nil {
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

// bindParamGetFlowsOutcomes binds the parameter divisionId
func (o *GetFlowsOutcomesParams) bindParamDivisionID(formats strfmt.Registry) []string {
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

// bindParamGetFlowsOutcomes binds the parameter id
func (o *GetFlowsOutcomesParams) bindParamID(formats strfmt.Registry) []string {
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
