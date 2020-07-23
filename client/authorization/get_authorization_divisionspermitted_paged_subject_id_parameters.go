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

// NewGetAuthorizationDivisionspermittedPagedSubjectIDParams creates a new GetAuthorizationDivisionspermittedPagedSubjectIDParams object
// with the default values initialized.
func NewGetAuthorizationDivisionspermittedPagedSubjectIDParams() *GetAuthorizationDivisionspermittedPagedSubjectIDParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetAuthorizationDivisionspermittedPagedSubjectIDParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationDivisionspermittedPagedSubjectIDParamsWithTimeout creates a new GetAuthorizationDivisionspermittedPagedSubjectIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthorizationDivisionspermittedPagedSubjectIDParamsWithTimeout(timeout time.Duration) *GetAuthorizationDivisionspermittedPagedSubjectIDParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetAuthorizationDivisionspermittedPagedSubjectIDParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetAuthorizationDivisionspermittedPagedSubjectIDParamsWithContext creates a new GetAuthorizationDivisionspermittedPagedSubjectIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthorizationDivisionspermittedPagedSubjectIDParamsWithContext(ctx context.Context) *GetAuthorizationDivisionspermittedPagedSubjectIDParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetAuthorizationDivisionspermittedPagedSubjectIDParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetAuthorizationDivisionspermittedPagedSubjectIDParamsWithHTTPClient creates a new GetAuthorizationDivisionspermittedPagedSubjectIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthorizationDivisionspermittedPagedSubjectIDParamsWithHTTPClient(client *http.Client) *GetAuthorizationDivisionspermittedPagedSubjectIDParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetAuthorizationDivisionspermittedPagedSubjectIDParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetAuthorizationDivisionspermittedPagedSubjectIDParams contains all the parameters to send to the API endpoint
for the get authorization divisionspermitted paged subject Id operation typically these are written to a http.Request
*/
type GetAuthorizationDivisionspermittedPagedSubjectIDParams struct {

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
	/*SubjectID
	  Subject ID (user or group)

	*/
	SubjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) WithTimeout(timeout time.Duration) *GetAuthorizationDivisionspermittedPagedSubjectIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) WithContext(ctx context.Context) *GetAuthorizationDivisionspermittedPagedSubjectIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) WithHTTPClient(client *http.Client) *GetAuthorizationDivisionspermittedPagedSubjectIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) WithPageNumber(pageNumber *int32) *GetAuthorizationDivisionspermittedPagedSubjectIDParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) WithPageSize(pageSize *int32) *GetAuthorizationDivisionspermittedPagedSubjectIDParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPermission adds the permission to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) WithPermission(permission string) *GetAuthorizationDivisionspermittedPagedSubjectIDParams {
	o.SetPermission(permission)
	return o
}

// SetPermission adds the permission to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) SetPermission(permission string) {
	o.Permission = permission
}

// WithSubjectID adds the subjectID to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) WithSubjectID(subjectID string) *GetAuthorizationDivisionspermittedPagedSubjectIDParams {
	o.SetSubjectID(subjectID)
	return o
}

// SetSubjectID adds the subjectId to the get authorization divisionspermitted paged subject Id params
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) SetSubjectID(subjectID string) {
	o.SubjectID = subjectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationDivisionspermittedPagedSubjectIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subjectId
	if err := r.SetPathParam("subjectId", o.SubjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
