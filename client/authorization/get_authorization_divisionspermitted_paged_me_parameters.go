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

// NewGetAuthorizationDivisionspermittedPagedMeParams creates a new GetAuthorizationDivisionspermittedPagedMeParams object
// with the default values initialized.
func NewGetAuthorizationDivisionspermittedPagedMeParams() *GetAuthorizationDivisionspermittedPagedMeParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetAuthorizationDivisionspermittedPagedMeParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationDivisionspermittedPagedMeParamsWithTimeout creates a new GetAuthorizationDivisionspermittedPagedMeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthorizationDivisionspermittedPagedMeParamsWithTimeout(timeout time.Duration) *GetAuthorizationDivisionspermittedPagedMeParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetAuthorizationDivisionspermittedPagedMeParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetAuthorizationDivisionspermittedPagedMeParamsWithContext creates a new GetAuthorizationDivisionspermittedPagedMeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthorizationDivisionspermittedPagedMeParamsWithContext(ctx context.Context) *GetAuthorizationDivisionspermittedPagedMeParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetAuthorizationDivisionspermittedPagedMeParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetAuthorizationDivisionspermittedPagedMeParamsWithHTTPClient creates a new GetAuthorizationDivisionspermittedPagedMeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthorizationDivisionspermittedPagedMeParamsWithHTTPClient(client *http.Client) *GetAuthorizationDivisionspermittedPagedMeParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetAuthorizationDivisionspermittedPagedMeParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetAuthorizationDivisionspermittedPagedMeParams contains all the parameters to send to the API endpoint
for the get authorization divisionspermitted paged me operation typically these are written to a http.Request
*/
type GetAuthorizationDivisionspermittedPagedMeParams struct {

	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*Permission
	  The permission string, including the object to access, e.g. routing:queue:view

	*/
	Permission string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get authorization divisionspermitted paged me params
func (o *GetAuthorizationDivisionspermittedPagedMeParams) WithTimeout(timeout time.Duration) *GetAuthorizationDivisionspermittedPagedMeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization divisionspermitted paged me params
func (o *GetAuthorizationDivisionspermittedPagedMeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization divisionspermitted paged me params
func (o *GetAuthorizationDivisionspermittedPagedMeParams) WithContext(ctx context.Context) *GetAuthorizationDivisionspermittedPagedMeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization divisionspermitted paged me params
func (o *GetAuthorizationDivisionspermittedPagedMeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization divisionspermitted paged me params
func (o *GetAuthorizationDivisionspermittedPagedMeParams) WithHTTPClient(client *http.Client) *GetAuthorizationDivisionspermittedPagedMeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization divisionspermitted paged me params
func (o *GetAuthorizationDivisionspermittedPagedMeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get authorization divisionspermitted paged me params
func (o *GetAuthorizationDivisionspermittedPagedMeParams) WithPageNumber(pageNumber *int32) *GetAuthorizationDivisionspermittedPagedMeParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get authorization divisionspermitted paged me params
func (o *GetAuthorizationDivisionspermittedPagedMeParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get authorization divisionspermitted paged me params
func (o *GetAuthorizationDivisionspermittedPagedMeParams) WithPageSize(pageSize *int32) *GetAuthorizationDivisionspermittedPagedMeParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get authorization divisionspermitted paged me params
func (o *GetAuthorizationDivisionspermittedPagedMeParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPermission adds the permission to the get authorization divisionspermitted paged me params
func (o *GetAuthorizationDivisionspermittedPagedMeParams) WithPermission(permission string) *GetAuthorizationDivisionspermittedPagedMeParams {
	o.SetPermission(permission)
	return o
}

// SetPermission adds the permission to the get authorization divisionspermitted paged me params
func (o *GetAuthorizationDivisionspermittedPagedMeParams) SetPermission(permission string) {
	o.Permission = permission
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationDivisionspermittedPagedMeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// query param permission
	qrPermission := o.Permission
	qPermission := qrPermission
	if qPermission != "" {
		if err := r.SetQueryParam("permission", qPermission); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}