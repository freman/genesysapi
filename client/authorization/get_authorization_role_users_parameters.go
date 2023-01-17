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

// NewGetAuthorizationRoleUsersParams creates a new GetAuthorizationRoleUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAuthorizationRoleUsersParams() *GetAuthorizationRoleUsersParams {
	return &GetAuthorizationRoleUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationRoleUsersParamsWithTimeout creates a new GetAuthorizationRoleUsersParams object
// with the ability to set a timeout on a request.
func NewGetAuthorizationRoleUsersParamsWithTimeout(timeout time.Duration) *GetAuthorizationRoleUsersParams {
	return &GetAuthorizationRoleUsersParams{
		timeout: timeout,
	}
}

// NewGetAuthorizationRoleUsersParamsWithContext creates a new GetAuthorizationRoleUsersParams object
// with the ability to set a context for a request.
func NewGetAuthorizationRoleUsersParamsWithContext(ctx context.Context) *GetAuthorizationRoleUsersParams {
	return &GetAuthorizationRoleUsersParams{
		Context: ctx,
	}
}

// NewGetAuthorizationRoleUsersParamsWithHTTPClient creates a new GetAuthorizationRoleUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAuthorizationRoleUsersParamsWithHTTPClient(client *http.Client) *GetAuthorizationRoleUsersParams {
	return &GetAuthorizationRoleUsersParams{
		HTTPClient: client,
	}
}

/*
GetAuthorizationRoleUsersParams contains all the parameters to send to the API endpoint

	for the get authorization role users operation.

	Typically these are written to a http.Request.
*/
type GetAuthorizationRoleUsersParams struct {

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

	/* RoleID.

	   Role ID
	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get authorization role users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationRoleUsersParams) WithDefaults() *GetAuthorizationRoleUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get authorization role users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationRoleUsersParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetAuthorizationRoleUsersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get authorization role users params
func (o *GetAuthorizationRoleUsersParams) WithTimeout(timeout time.Duration) *GetAuthorizationRoleUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization role users params
func (o *GetAuthorizationRoleUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization role users params
func (o *GetAuthorizationRoleUsersParams) WithContext(ctx context.Context) *GetAuthorizationRoleUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization role users params
func (o *GetAuthorizationRoleUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization role users params
func (o *GetAuthorizationRoleUsersParams) WithHTTPClient(client *http.Client) *GetAuthorizationRoleUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization role users params
func (o *GetAuthorizationRoleUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get authorization role users params
func (o *GetAuthorizationRoleUsersParams) WithPageNumber(pageNumber *int32) *GetAuthorizationRoleUsersParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get authorization role users params
func (o *GetAuthorizationRoleUsersParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get authorization role users params
func (o *GetAuthorizationRoleUsersParams) WithPageSize(pageSize *int32) *GetAuthorizationRoleUsersParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get authorization role users params
func (o *GetAuthorizationRoleUsersParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithRoleID adds the roleID to the get authorization role users params
func (o *GetAuthorizationRoleUsersParams) WithRoleID(roleID string) *GetAuthorizationRoleUsersParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the get authorization role users params
func (o *GetAuthorizationRoleUsersParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationRoleUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param roleId
	if err := r.SetPathParam("roleId", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
