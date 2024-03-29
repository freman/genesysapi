// Code generated by go-swagger; DO NOT EDIT.

package knowledge

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

// NewGetKnowledgeKnowledgebaseCategoriesParams creates a new GetKnowledgeKnowledgebaseCategoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKnowledgeKnowledgebaseCategoriesParams() *GetKnowledgeKnowledgebaseCategoriesParams {
	return &GetKnowledgeKnowledgebaseCategoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseCategoriesParamsWithTimeout creates a new GetKnowledgeKnowledgebaseCategoriesParams object
// with the ability to set a timeout on a request.
func NewGetKnowledgeKnowledgebaseCategoriesParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseCategoriesParams {
	return &GetKnowledgeKnowledgebaseCategoriesParams{
		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseCategoriesParamsWithContext creates a new GetKnowledgeKnowledgebaseCategoriesParams object
// with the ability to set a context for a request.
func NewGetKnowledgeKnowledgebaseCategoriesParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseCategoriesParams {
	return &GetKnowledgeKnowledgebaseCategoriesParams{
		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseCategoriesParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseCategoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKnowledgeKnowledgebaseCategoriesParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseCategoriesParams {
	return &GetKnowledgeKnowledgebaseCategoriesParams{
		HTTPClient: client,
	}
}

/*
GetKnowledgeKnowledgebaseCategoriesParams contains all the parameters to send to the API endpoint

	for the get knowledge knowledgebase categories operation.

	Typically these are written to a http.Request.
*/
type GetKnowledgeKnowledgebaseCategoriesParams struct {

	/* After.

	   The cursor that points to the end of the set of entities that has been returned.
	*/
	After *string

	/* Before.

	   The cursor that points to the start of the set of entities that has been returned.
	*/
	Before *string

	/* Expand.

	   The specified entity attribute will be filled. Supported value:"Ancestors": every ancestors will be filled via the parent attribute recursively,but only the id, name, parentId will be present for the ancestors.
	*/
	Expand *string

	/* IncludeDocumentCount.

	   If specified, retrieves the number of documents related to category.
	*/
	IncludeDocumentCount *bool

	/* IsRoot.

	   If specified, retrieves only the root categories.
	*/
	IsRoot *bool

	/* KnowledgeBaseID.

	   Knowledge base ID
	*/
	KnowledgeBaseID string

	/* Name.

	   Filter to return the categories that starts with the given category name.
	*/
	Name *string

	/* PageSize.

	   Number of entities to return. Maximum of 200.
	*/
	PageSize *string

	/* ParentID.

	   If specified, retrieves the children categories by parent category ID.
	*/
	ParentID *string

	/* SortBy.

	   Name: sort by category names alphabetically; Hierarchy: sort by the full path of hierarchical category names alphabetically

	   Default: "Name"
	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get knowledge knowledgebase categories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithDefaults() *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get knowledge knowledgebase categories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetDefaults() {
	var (
		sortByDefault = string("Name")
	)

	val := GetKnowledgeKnowledgebaseCategoriesParams{
		SortBy: &sortByDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithAfter(after *string) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetAfter(after *string) {
	o.After = after
}

// WithBefore adds the before to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithBefore(before *string) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetBefore(before *string) {
	o.Before = before
}

// WithExpand adds the expand to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithExpand(expand *string) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithIncludeDocumentCount adds the includeDocumentCount to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithIncludeDocumentCount(includeDocumentCount *bool) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetIncludeDocumentCount(includeDocumentCount)
	return o
}

// SetIncludeDocumentCount adds the includeDocumentCount to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetIncludeDocumentCount(includeDocumentCount *bool) {
	o.IncludeDocumentCount = includeDocumentCount
}

// WithIsRoot adds the isRoot to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithIsRoot(isRoot *bool) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetIsRoot(isRoot)
	return o
}

// SetIsRoot adds the isRoot to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetIsRoot(isRoot *bool) {
	o.IsRoot = isRoot
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithName adds the name to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithName(name *string) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetName(name *string) {
	o.Name = name
}

// WithPageSize adds the pageSize to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithPageSize(pageSize *string) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WithParentID adds the parentID to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithParentID(parentID *string) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetParentID(parentID *string) {
	o.ParentID = parentID
}

// WithSortBy adds the sortBy to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WithSortBy(sortBy *string) *GetKnowledgeKnowledgebaseCategoriesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get knowledge knowledgebase categories params
func (o *GetKnowledgeKnowledgebaseCategoriesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if o.Before != nil {

		// query param before
		var qrBefore string

		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := qrBefore
		if qBefore != "" {

			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}
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

	if o.IncludeDocumentCount != nil {

		// query param includeDocumentCount
		var qrIncludeDocumentCount bool

		if o.IncludeDocumentCount != nil {
			qrIncludeDocumentCount = *o.IncludeDocumentCount
		}
		qIncludeDocumentCount := swag.FormatBool(qrIncludeDocumentCount)
		if qIncludeDocumentCount != "" {

			if err := r.SetQueryParam("includeDocumentCount", qIncludeDocumentCount); err != nil {
				return err
			}
		}
	}

	if o.IsRoot != nil {

		// query param isRoot
		var qrIsRoot bool

		if o.IsRoot != nil {
			qrIsRoot = *o.IsRoot
		}
		qIsRoot := swag.FormatBool(qrIsRoot)
		if qIsRoot != "" {

			if err := r.SetQueryParam("isRoot", qIsRoot); err != nil {
				return err
			}
		}
	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
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

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize string

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := qrPageSize
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.ParentID != nil {

		// query param parentId
		var qrParentID string

		if o.ParentID != nil {
			qrParentID = *o.ParentID
		}
		qParentID := qrParentID
		if qParentID != "" {

			if err := r.SetQueryParam("parentId", qParentID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
