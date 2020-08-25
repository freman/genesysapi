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

// NewGetAuthorizationRoleSubjectgrantsParams creates a new GetAuthorizationRoleSubjectgrantsParams object
// with the default values initialized.
func NewGetAuthorizationRoleSubjectgrantsParams() *GetAuthorizationRoleSubjectgrantsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetAuthorizationRoleSubjectgrantsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationRoleSubjectgrantsParamsWithTimeout creates a new GetAuthorizationRoleSubjectgrantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthorizationRoleSubjectgrantsParamsWithTimeout(timeout time.Duration) *GetAuthorizationRoleSubjectgrantsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetAuthorizationRoleSubjectgrantsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetAuthorizationRoleSubjectgrantsParamsWithContext creates a new GetAuthorizationRoleSubjectgrantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthorizationRoleSubjectgrantsParamsWithContext(ctx context.Context) *GetAuthorizationRoleSubjectgrantsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetAuthorizationRoleSubjectgrantsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetAuthorizationRoleSubjectgrantsParamsWithHTTPClient creates a new GetAuthorizationRoleSubjectgrantsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthorizationRoleSubjectgrantsParamsWithHTTPClient(client *http.Client) *GetAuthorizationRoleSubjectgrantsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetAuthorizationRoleSubjectgrantsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetAuthorizationRoleSubjectgrantsParams contains all the parameters to send to the API endpoint
for the get authorization role subjectgrants operation typically these are written to a http.Request
*/
type GetAuthorizationRoleSubjectgrantsParams struct {

	/*Expand
	  variable name requested by expand list

	*/
	Expand []string
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
	/*PreviousPage
	  Previous page token

	*/
	PreviousPage *string
	/*RoleID
	  Role ID

	*/
	RoleID string
	/*SortBy
	  variable name requested to sort by

	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) WithTimeout(timeout time.Duration) *GetAuthorizationRoleSubjectgrantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) WithContext(ctx context.Context) *GetAuthorizationRoleSubjectgrantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) WithHTTPClient(client *http.Client) *GetAuthorizationRoleSubjectgrantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) WithExpand(expand []string) *GetAuthorizationRoleSubjectgrantsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithNextPage adds the nextPage to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) WithNextPage(nextPage *string) *GetAuthorizationRoleSubjectgrantsParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithPageNumber adds the pageNumber to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) WithPageNumber(pageNumber *int32) *GetAuthorizationRoleSubjectgrantsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) WithPageSize(pageSize *int32) *GetAuthorizationRoleSubjectgrantsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPreviousPage adds the previousPage to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) WithPreviousPage(previousPage *string) *GetAuthorizationRoleSubjectgrantsParams {
	o.SetPreviousPage(previousPage)
	return o
}

// SetPreviousPage adds the previousPage to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) SetPreviousPage(previousPage *string) {
	o.PreviousPage = previousPage
}

// WithRoleID adds the roleID to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) WithRoleID(roleID string) *GetAuthorizationRoleSubjectgrantsParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WithSortBy adds the sortBy to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) WithSortBy(sortBy *string) *GetAuthorizationRoleSubjectgrantsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get authorization role subjectgrants params
func (o *GetAuthorizationRoleSubjectgrantsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationRoleSubjectgrantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param roleId
	if err := r.SetPathParam("roleId", o.RoleID); err != nil {
		return err
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