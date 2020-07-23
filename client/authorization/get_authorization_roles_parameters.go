// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// NewGetAuthorizationRolesParams creates a new GetAuthorizationRolesParams object
// with the default values initialized.
func NewGetAuthorizationRolesParams() *GetAuthorizationRolesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		userCountDefault  = bool(true)
	)
	return &GetAuthorizationRolesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		UserCount:  &userCountDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationRolesParamsWithTimeout creates a new GetAuthorizationRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthorizationRolesParamsWithTimeout(timeout time.Duration) *GetAuthorizationRolesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		userCountDefault  = bool(true)
	)
	return &GetAuthorizationRolesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		UserCount:  &userCountDefault,

		timeout: timeout,
	}
}

// NewGetAuthorizationRolesParamsWithContext creates a new GetAuthorizationRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthorizationRolesParamsWithContext(ctx context.Context) *GetAuthorizationRolesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		userCountDefault  = bool(true)
	)
	return &GetAuthorizationRolesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		UserCount:  &userCountDefault,

		Context: ctx,
	}
}

// NewGetAuthorizationRolesParamsWithHTTPClient creates a new GetAuthorizationRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthorizationRolesParamsWithHTTPClient(client *http.Client) *GetAuthorizationRolesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		userCountDefault  = bool(true)
	)
	return &GetAuthorizationRolesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		UserCount:  &userCountDefault,
		HTTPClient: client,
	}
}

/*GetAuthorizationRolesParams contains all the parameters to send to the API endpoint
for the get authorization roles operation typically these are written to a http.Request
*/
type GetAuthorizationRolesParams struct {

	/*DefaultRoleID*/
	DefaultRoleID []string
	/*Expand
	  variable name requested by expand list

	*/
	Expand []string
	/*ID
	  id

	*/
	ID []string
	/*Name*/
	Name *string
	/*NextPage
	  next page token

	*/
	NextPage *string
	/*PageNumber
	  The page number requested

	*/
	PageNumber *int32
	/*PageSize
	  The total page size requested

	*/
	PageSize *int32
	/*Permission*/
	Permission []string
	/*PreviousPage
	  Previous page token

	*/
	PreviousPage *string
	/*SortBy
	  variable name requested to sort by

	*/
	SortBy *string
	/*UserCount*/
	UserCount *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithTimeout(timeout time.Duration) *GetAuthorizationRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithContext(ctx context.Context) *GetAuthorizationRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithHTTPClient(client *http.Client) *GetAuthorizationRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDefaultRoleID adds the defaultRoleID to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithDefaultRoleID(defaultRoleID []string) *GetAuthorizationRolesParams {
	o.SetDefaultRoleID(defaultRoleID)
	return o
}

// SetDefaultRoleID adds the defaultRoleId to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetDefaultRoleID(defaultRoleID []string) {
	o.DefaultRoleID = defaultRoleID
}

// WithExpand adds the expand to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithExpand(expand []string) *GetAuthorizationRolesParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithID adds the id to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithID(id []string) *GetAuthorizationRolesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithName(name *string) *GetAuthorizationRolesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetName(name *string) {
	o.Name = name
}

// WithNextPage adds the nextPage to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithNextPage(nextPage *string) *GetAuthorizationRolesParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithPageNumber adds the pageNumber to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithPageNumber(pageNumber *int32) *GetAuthorizationRolesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithPageSize(pageSize *int32) *GetAuthorizationRolesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPermission adds the permission to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithPermission(permission []string) *GetAuthorizationRolesParams {
	o.SetPermission(permission)
	return o
}

// SetPermission adds the permission to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetPermission(permission []string) {
	o.Permission = permission
}

// WithPreviousPage adds the previousPage to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithPreviousPage(previousPage *string) *GetAuthorizationRolesParams {
	o.SetPreviousPage(previousPage)
	return o
}

// SetPreviousPage adds the previousPage to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetPreviousPage(previousPage *string) {
	o.PreviousPage = previousPage
}

// WithSortBy adds the sortBy to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithSortBy(sortBy *string) *GetAuthorizationRolesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithUserCount adds the userCount to the get authorization roles params
func (o *GetAuthorizationRolesParams) WithUserCount(userCount *bool) *GetAuthorizationRolesParams {
	o.SetUserCount(userCount)
	return o
}

// SetUserCount adds the userCount to the get authorization roles params
func (o *GetAuthorizationRolesParams) SetUserCount(userCount *bool) {
	o.UserCount = userCount
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDefaultRoleID := o.DefaultRoleID

	joinedDefaultRoleID := swag.JoinByFormat(valuesDefaultRoleID, "multi")
	// query array param defaultRoleId
	if err := r.SetQueryParam("defaultRoleId", joinedDefaultRoleID...); err != nil {
		return err
	}

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
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

	if o.NextPage != nil {

		// query param nextPage
		var qrNextPage string
		if o.NextPage != nil {
			qrNextPage = *o.NextPage
		}
		qNextPage := qrNextPage
		if qNextPage != "" {
			if err := r.SetQueryParam("nextPage", qNextPage); err != nil {
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

	valuesPermission := o.Permission

	joinedPermission := swag.JoinByFormat(valuesPermission, "multi")
	// query array param permission
	if err := r.SetQueryParam("permission", joinedPermission...); err != nil {
		return err
	}

	if o.PreviousPage != nil {

		// query param previousPage
		var qrPreviousPage string
		if o.PreviousPage != nil {
			qrPreviousPage = *o.PreviousPage
		}
		qPreviousPage := qrPreviousPage
		if qPreviousPage != "" {
			if err := r.SetQueryParam("previousPage", qPreviousPage); err != nil {
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

	if o.UserCount != nil {

		// query param userCount
		var qrUserCount bool
		if o.UserCount != nil {
			qrUserCount = *o.UserCount
		}
		qUserCount := swag.FormatBool(qrUserCount)
		if qUserCount != "" {
			if err := r.SetQueryParam("userCount", qUserCount); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
