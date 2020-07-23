// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewGetUserRoutinglanguagesParams creates a new GetUserRoutinglanguagesParams object
// with the default values initialized.
func NewGetUserRoutinglanguagesParams() *GetUserRoutinglanguagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ASC")
	)
	return &GetUserRoutinglanguagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserRoutinglanguagesParamsWithTimeout creates a new GetUserRoutinglanguagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserRoutinglanguagesParamsWithTimeout(timeout time.Duration) *GetUserRoutinglanguagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ASC")
	)
	return &GetUserRoutinglanguagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetUserRoutinglanguagesParamsWithContext creates a new GetUserRoutinglanguagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserRoutinglanguagesParamsWithContext(ctx context.Context) *GetUserRoutinglanguagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ASC")
	)
	return &GetUserRoutinglanguagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetUserRoutinglanguagesParamsWithHTTPClient creates a new GetUserRoutinglanguagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserRoutinglanguagesParamsWithHTTPClient(client *http.Client) *GetUserRoutinglanguagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ASC")
	)
	return &GetUserRoutinglanguagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetUserRoutinglanguagesParams contains all the parameters to send to the API endpoint
for the get user routinglanguages operation typically these are written to a http.Request
*/
type GetUserRoutinglanguagesParams struct {

	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*SortOrder
	  Ascending or descending sort order

	*/
	SortOrder *string
	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) WithTimeout(timeout time.Duration) *GetUserRoutinglanguagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) WithContext(ctx context.Context) *GetUserRoutinglanguagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) WithHTTPClient(client *http.Client) *GetUserRoutinglanguagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) WithPageNumber(pageNumber *int32) *GetUserRoutinglanguagesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) WithPageSize(pageSize *int32) *GetUserRoutinglanguagesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortOrder adds the sortOrder to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) WithSortOrder(sortOrder *string) *GetUserRoutinglanguagesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithUserID adds the userID to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) WithUserID(userID string) *GetUserRoutinglanguagesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user routinglanguages params
func (o *GetUserRoutinglanguagesParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserRoutinglanguagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
