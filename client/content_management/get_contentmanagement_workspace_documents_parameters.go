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

// NewGetContentmanagementWorkspaceDocumentsParams creates a new GetContentmanagementWorkspaceDocumentsParams object
// with the default values initialized.
func NewGetContentmanagementWorkspaceDocumentsParams() *GetContentmanagementWorkspaceDocumentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ascending")
	)
	return &GetContentmanagementWorkspaceDocumentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentmanagementWorkspaceDocumentsParamsWithTimeout creates a new GetContentmanagementWorkspaceDocumentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContentmanagementWorkspaceDocumentsParamsWithTimeout(timeout time.Duration) *GetContentmanagementWorkspaceDocumentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ascending")
	)
	return &GetContentmanagementWorkspaceDocumentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetContentmanagementWorkspaceDocumentsParamsWithContext creates a new GetContentmanagementWorkspaceDocumentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContentmanagementWorkspaceDocumentsParamsWithContext(ctx context.Context) *GetContentmanagementWorkspaceDocumentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ascending")
	)
	return &GetContentmanagementWorkspaceDocumentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetContentmanagementWorkspaceDocumentsParamsWithHTTPClient creates a new GetContentmanagementWorkspaceDocumentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContentmanagementWorkspaceDocumentsParamsWithHTTPClient(client *http.Client) *GetContentmanagementWorkspaceDocumentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ascending")
	)
	return &GetContentmanagementWorkspaceDocumentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetContentmanagementWorkspaceDocumentsParams contains all the parameters to send to the API endpoint
for the get contentmanagement workspace documents operation typically these are written to a http.Request
*/
type GetContentmanagementWorkspaceDocumentsParams struct {

	/*Expand
	  Which fields, if any, to expand.

	*/
	Expand []string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*SortBy
	  name or dateCreated

	*/
	SortBy *string
	/*SortOrder
	  ascending or descending

	*/
	SortOrder *string
	/*WorkspaceID
	  Workspace ID

	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) WithTimeout(timeout time.Duration) *GetContentmanagementWorkspaceDocumentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) WithContext(ctx context.Context) *GetContentmanagementWorkspaceDocumentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) WithHTTPClient(client *http.Client) *GetContentmanagementWorkspaceDocumentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) WithExpand(expand []string) *GetContentmanagementWorkspaceDocumentsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithPageNumber adds the pageNumber to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) WithPageNumber(pageNumber *int32) *GetContentmanagementWorkspaceDocumentsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) WithPageSize(pageSize *int32) *GetContentmanagementWorkspaceDocumentsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) WithSortBy(sortBy *string) *GetContentmanagementWorkspaceDocumentsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) WithSortOrder(sortOrder *string) *GetContentmanagementWorkspaceDocumentsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithWorkspaceID adds the workspaceID to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) WithWorkspaceID(workspaceID string) *GetContentmanagementWorkspaceDocumentsParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get contentmanagement workspace documents params
func (o *GetContentmanagementWorkspaceDocumentsParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentmanagementWorkspaceDocumentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
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

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", o.WorkspaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}