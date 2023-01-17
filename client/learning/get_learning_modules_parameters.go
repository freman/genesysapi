// Code generated by go-swagger; DO NOT EDIT.

package learning

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

// NewGetLearningModulesParams creates a new GetLearningModulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLearningModulesParams() *GetLearningModulesParams {
	return &GetLearningModulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearningModulesParamsWithTimeout creates a new GetLearningModulesParams object
// with the ability to set a timeout on a request.
func NewGetLearningModulesParamsWithTimeout(timeout time.Duration) *GetLearningModulesParams {
	return &GetLearningModulesParams{
		timeout: timeout,
	}
}

// NewGetLearningModulesParamsWithContext creates a new GetLearningModulesParams object
// with the ability to set a context for a request.
func NewGetLearningModulesParamsWithContext(ctx context.Context) *GetLearningModulesParams {
	return &GetLearningModulesParams{
		Context: ctx,
	}
}

// NewGetLearningModulesParamsWithHTTPClient creates a new GetLearningModulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLearningModulesParamsWithHTTPClient(client *http.Client) *GetLearningModulesParams {
	return &GetLearningModulesParams{
		HTTPClient: client,
	}
}

/*
GetLearningModulesParams contains all the parameters to send to the API endpoint

	for the get learning modules operation.

	Typically these are written to a http.Request.
*/
type GetLearningModulesParams struct {

	/* Expand.

	   Fields to expand in response(case insensitive)
	*/
	Expand []string

	/* IsArchived.

	   Archive status
	*/
	IsArchived *bool

	/* IsPublished.

	   Specifies if only the Unpublished (isPublished is "False") or Published (isPublished is "True") modules are returned. If isPublished is "Any" or omitted, both types are returned

	   Default: "Any"
	*/
	IsPublished *string

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

	/* SearchTerm.

	   Search Term (searchable by name)
	*/
	SearchTerm *string

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

	/* Statuses.

	   Specifies the module statuses to filter by
	*/
	Statuses []string

	/* Types.

	   Specifies the module types.
	*/
	Types []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get learning modules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLearningModulesParams) WithDefaults() *GetLearningModulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get learning modules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLearningModulesParams) SetDefaults() {
	var (
		isArchivedDefault = bool(false)

		isPublishedDefault = string("Any")

		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		sortByDefault = string("name")

		sortOrderDefault = string("ascending")
	)

	val := GetLearningModulesParams{
		IsArchived:  &isArchivedDefault,
		IsPublished: &isPublishedDefault,
		PageNumber:  &pageNumberDefault,
		PageSize:    &pageSizeDefault,
		SortBy:      &sortByDefault,
		SortOrder:   &sortOrderDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get learning modules params
func (o *GetLearningModulesParams) WithTimeout(timeout time.Duration) *GetLearningModulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learning modules params
func (o *GetLearningModulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learning modules params
func (o *GetLearningModulesParams) WithContext(ctx context.Context) *GetLearningModulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learning modules params
func (o *GetLearningModulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learning modules params
func (o *GetLearningModulesParams) WithHTTPClient(client *http.Client) *GetLearningModulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learning modules params
func (o *GetLearningModulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get learning modules params
func (o *GetLearningModulesParams) WithExpand(expand []string) *GetLearningModulesParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get learning modules params
func (o *GetLearningModulesParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithIsArchived adds the isArchived to the get learning modules params
func (o *GetLearningModulesParams) WithIsArchived(isArchived *bool) *GetLearningModulesParams {
	o.SetIsArchived(isArchived)
	return o
}

// SetIsArchived adds the isArchived to the get learning modules params
func (o *GetLearningModulesParams) SetIsArchived(isArchived *bool) {
	o.IsArchived = isArchived
}

// WithIsPublished adds the isPublished to the get learning modules params
func (o *GetLearningModulesParams) WithIsPublished(isPublished *string) *GetLearningModulesParams {
	o.SetIsPublished(isPublished)
	return o
}

// SetIsPublished adds the isPublished to the get learning modules params
func (o *GetLearningModulesParams) SetIsPublished(isPublished *string) {
	o.IsPublished = isPublished
}

// WithPageNumber adds the pageNumber to the get learning modules params
func (o *GetLearningModulesParams) WithPageNumber(pageNumber *int32) *GetLearningModulesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get learning modules params
func (o *GetLearningModulesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get learning modules params
func (o *GetLearningModulesParams) WithPageSize(pageSize *int32) *GetLearningModulesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get learning modules params
func (o *GetLearningModulesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSearchTerm adds the searchTerm to the get learning modules params
func (o *GetLearningModulesParams) WithSearchTerm(searchTerm *string) *GetLearningModulesParams {
	o.SetSearchTerm(searchTerm)
	return o
}

// SetSearchTerm adds the searchTerm to the get learning modules params
func (o *GetLearningModulesParams) SetSearchTerm(searchTerm *string) {
	o.SearchTerm = searchTerm
}

// WithSortBy adds the sortBy to the get learning modules params
func (o *GetLearningModulesParams) WithSortBy(sortBy *string) *GetLearningModulesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get learning modules params
func (o *GetLearningModulesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get learning modules params
func (o *GetLearningModulesParams) WithSortOrder(sortOrder *string) *GetLearningModulesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get learning modules params
func (o *GetLearningModulesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithStatuses adds the statuses to the get learning modules params
func (o *GetLearningModulesParams) WithStatuses(statuses []string) *GetLearningModulesParams {
	o.SetStatuses(statuses)
	return o
}

// SetStatuses adds the statuses to the get learning modules params
func (o *GetLearningModulesParams) SetStatuses(statuses []string) {
	o.Statuses = statuses
}

// WithTypes adds the types to the get learning modules params
func (o *GetLearningModulesParams) WithTypes(types []string) *GetLearningModulesParams {
	o.SetTypes(types)
	return o
}

// SetTypes adds the types to the get learning modules params
func (o *GetLearningModulesParams) SetTypes(types []string) {
	o.Types = types
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearningModulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	if o.IsArchived != nil {

		// query param isArchived
		var qrIsArchived bool

		if o.IsArchived != nil {
			qrIsArchived = *o.IsArchived
		}
		qIsArchived := swag.FormatBool(qrIsArchived)
		if qIsArchived != "" {

			if err := r.SetQueryParam("isArchived", qIsArchived); err != nil {
				return err
			}
		}
	}

	if o.IsPublished != nil {

		// query param isPublished
		var qrIsPublished string

		if o.IsPublished != nil {
			qrIsPublished = *o.IsPublished
		}
		qIsPublished := qrIsPublished
		if qIsPublished != "" {

			if err := r.SetQueryParam("isPublished", qIsPublished); err != nil {
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

	if o.SearchTerm != nil {

		// query param searchTerm
		var qrSearchTerm string

		if o.SearchTerm != nil {
			qrSearchTerm = *o.SearchTerm
		}
		qSearchTerm := qrSearchTerm
		if qSearchTerm != "" {

			if err := r.SetQueryParam("searchTerm", qSearchTerm); err != nil {
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

	if o.Statuses != nil {

		// binding items for statuses
		joinedStatuses := o.bindParamStatuses(reg)

		// query array param statuses
		if err := r.SetQueryParam("statuses", joinedStatuses...); err != nil {
			return err
		}
	}

	if o.Types != nil {

		// binding items for types
		joinedTypes := o.bindParamTypes(reg)

		// query array param types
		if err := r.SetQueryParam("types", joinedTypes...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetLearningModules binds the parameter expand
func (o *GetLearningModulesParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}

// bindParamGetLearningModules binds the parameter statuses
func (o *GetLearningModulesParams) bindParamStatuses(formats strfmt.Registry) []string {
	statusesIR := o.Statuses

	var statusesIC []string
	for _, statusesIIR := range statusesIR { // explode []string

		statusesIIV := statusesIIR // string as string
		statusesIC = append(statusesIC, statusesIIV)
	}

	// items.CollectionFormat: "multi"
	statusesIS := swag.JoinByFormat(statusesIC, "multi")

	return statusesIS
}

// bindParamGetLearningModules binds the parameter types
func (o *GetLearningModulesParams) bindParamTypes(formats strfmt.Registry) []string {
	typesIR := o.Types

	var typesIC []string
	for _, typesIIR := range typesIR { // explode []string

		typesIIV := typesIIR // string as string
		typesIC = append(typesIC, typesIIV)
	}

	// items.CollectionFormat: "multi"
	typesIS := swag.JoinByFormat(typesIC, "multi")

	return typesIS
}
