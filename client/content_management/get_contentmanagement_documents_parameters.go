// Code generated by go-swagger; DO NOT EDIT.

package content_management

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

// NewGetContentmanagementDocumentsParams creates a new GetContentmanagementDocumentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContentmanagementDocumentsParams() *GetContentmanagementDocumentsParams {
	return &GetContentmanagementDocumentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentmanagementDocumentsParamsWithTimeout creates a new GetContentmanagementDocumentsParams object
// with the ability to set a timeout on a request.
func NewGetContentmanagementDocumentsParamsWithTimeout(timeout time.Duration) *GetContentmanagementDocumentsParams {
	return &GetContentmanagementDocumentsParams{
		timeout: timeout,
	}
}

// NewGetContentmanagementDocumentsParamsWithContext creates a new GetContentmanagementDocumentsParams object
// with the ability to set a context for a request.
func NewGetContentmanagementDocumentsParamsWithContext(ctx context.Context) *GetContentmanagementDocumentsParams {
	return &GetContentmanagementDocumentsParams{
		Context: ctx,
	}
}

// NewGetContentmanagementDocumentsParamsWithHTTPClient creates a new GetContentmanagementDocumentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContentmanagementDocumentsParamsWithHTTPClient(client *http.Client) *GetContentmanagementDocumentsParams {
	return &GetContentmanagementDocumentsParams{
		HTTPClient: client,
	}
}

/*
GetContentmanagementDocumentsParams contains all the parameters to send to the API endpoint

	for the get contentmanagement documents operation.

	Typically these are written to a http.Request.
*/
type GetContentmanagementDocumentsParams struct {

	/* Expand.

	   Which fields, if any, to expand.
	*/
	Expand []string

	/* Name.

	   Name
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

	   name or dateCreated
	*/
	SortBy *string

	/* SortOrder.

	   ascending or descending

	   Default: "ascending"
	*/
	SortOrder *string

	/* WorkspaceID.

	   Workspace ID
	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contentmanagement documents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentmanagementDocumentsParams) WithDefaults() *GetContentmanagementDocumentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contentmanagement documents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentmanagementDocumentsParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		sortOrderDefault = string("ascending")
	)

	val := GetContentmanagementDocumentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) WithTimeout(timeout time.Duration) *GetContentmanagementDocumentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) WithContext(ctx context.Context) *GetContentmanagementDocumentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) WithHTTPClient(client *http.Client) *GetContentmanagementDocumentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) WithExpand(expand []string) *GetContentmanagementDocumentsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithName adds the name to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) WithName(name *string) *GetContentmanagementDocumentsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) WithPageNumber(pageNumber *int32) *GetContentmanagementDocumentsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) WithPageSize(pageSize *int32) *GetContentmanagementDocumentsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) WithSortBy(sortBy *string) *GetContentmanagementDocumentsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) WithSortOrder(sortOrder *string) *GetContentmanagementDocumentsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithWorkspaceID adds the workspaceID to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) WithWorkspaceID(workspaceID string) *GetContentmanagementDocumentsParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get contentmanagement documents params
func (o *GetContentmanagementDocumentsParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentmanagementDocumentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param workspaceId
	qrWorkspaceID := o.WorkspaceID
	qWorkspaceID := qrWorkspaceID
	if qWorkspaceID != "" {

		if err := r.SetQueryParam("workspaceId", qWorkspaceID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetContentmanagementDocuments binds the parameter expand
func (o *GetContentmanagementDocumentsParams) bindParamExpand(formats strfmt.Registry) []string {
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
