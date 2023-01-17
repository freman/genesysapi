// Code generated by go-swagger; DO NOT EDIT.

package mobile_devices

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

// NewGetMobiledevicesParams creates a new GetMobiledevicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMobiledevicesParams() *GetMobiledevicesParams {
	return &GetMobiledevicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMobiledevicesParamsWithTimeout creates a new GetMobiledevicesParams object
// with the ability to set a timeout on a request.
func NewGetMobiledevicesParamsWithTimeout(timeout time.Duration) *GetMobiledevicesParams {
	return &GetMobiledevicesParams{
		timeout: timeout,
	}
}

// NewGetMobiledevicesParamsWithContext creates a new GetMobiledevicesParams object
// with the ability to set a context for a request.
func NewGetMobiledevicesParamsWithContext(ctx context.Context) *GetMobiledevicesParams {
	return &GetMobiledevicesParams{
		Context: ctx,
	}
}

// NewGetMobiledevicesParamsWithHTTPClient creates a new GetMobiledevicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMobiledevicesParamsWithHTTPClient(client *http.Client) *GetMobiledevicesParams {
	return &GetMobiledevicesParams{
		HTTPClient: client,
	}
}

/*
GetMobiledevicesParams contains all the parameters to send to the API endpoint

	for the get mobiledevices operation.

	Typically these are written to a http.Request.
*/
type GetMobiledevicesParams struct {

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

	/* SortOrder.

	   Ascending or descending sort order

	   Default: "ascending"
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get mobiledevices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMobiledevicesParams) WithDefaults() *GetMobiledevicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get mobiledevices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMobiledevicesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		sortOrderDefault = string("ascending")
	)

	val := GetMobiledevicesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get mobiledevices params
func (o *GetMobiledevicesParams) WithTimeout(timeout time.Duration) *GetMobiledevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mobiledevices params
func (o *GetMobiledevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mobiledevices params
func (o *GetMobiledevicesParams) WithContext(ctx context.Context) *GetMobiledevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mobiledevices params
func (o *GetMobiledevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mobiledevices params
func (o *GetMobiledevicesParams) WithHTTPClient(client *http.Client) *GetMobiledevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mobiledevices params
func (o *GetMobiledevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get mobiledevices params
func (o *GetMobiledevicesParams) WithPageNumber(pageNumber *int32) *GetMobiledevicesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get mobiledevices params
func (o *GetMobiledevicesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get mobiledevices params
func (o *GetMobiledevicesParams) WithPageSize(pageSize *int32) *GetMobiledevicesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get mobiledevices params
func (o *GetMobiledevicesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortOrder adds the sortOrder to the get mobiledevices params
func (o *GetMobiledevicesParams) WithSortOrder(sortOrder *string) *GetMobiledevicesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get mobiledevices params
func (o *GetMobiledevicesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetMobiledevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
