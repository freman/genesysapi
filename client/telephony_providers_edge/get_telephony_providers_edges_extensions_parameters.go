// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

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

// NewGetTelephonyProvidersEdgesExtensionsParams creates a new GetTelephonyProvidersEdgesExtensionsParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesExtensionsParams() *GetTelephonyProvidersEdgesExtensionsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("number")
		sortOrderDefault  = string("ASC")
	)
	return &GetTelephonyProvidersEdgesExtensionsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesExtensionsParamsWithTimeout creates a new GetTelephonyProvidersEdgesExtensionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesExtensionsParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesExtensionsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("number")
		sortOrderDefault  = string("ASC")
	)
	return &GetTelephonyProvidersEdgesExtensionsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesExtensionsParamsWithContext creates a new GetTelephonyProvidersEdgesExtensionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesExtensionsParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesExtensionsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("number")
		sortOrderDefault  = string("ASC")
	)
	return &GetTelephonyProvidersEdgesExtensionsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesExtensionsParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesExtensionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesExtensionsParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesExtensionsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("number")
		sortOrderDefault  = string("ASC")
	)
	return &GetTelephonyProvidersEdgesExtensionsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesExtensionsParams contains all the parameters to send to the API endpoint
for the get telephony providers edges extensions operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesExtensionsParams struct {

	/*Number
	  Filter by number

	*/
	Number *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*SortBy
	  Sort by

	*/
	SortBy *string
	/*SortOrder
	  Sort order

	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesExtensionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesExtensionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesExtensionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNumber adds the number to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) WithNumber(number *string) *GetTelephonyProvidersEdgesExtensionsParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) SetNumber(number *string) {
	o.Number = number
}

// WithPageNumber adds the pageNumber to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) WithPageNumber(pageNumber *int32) *GetTelephonyProvidersEdgesExtensionsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) WithPageSize(pageSize *int32) *GetTelephonyProvidersEdgesExtensionsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) WithSortBy(sortBy *string) *GetTelephonyProvidersEdgesExtensionsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) WithSortOrder(sortOrder *string) *GetTelephonyProvidersEdgesExtensionsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get telephony providers edges extensions params
func (o *GetTelephonyProvidersEdgesExtensionsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesExtensionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Number != nil {

		// query param number
		var qrNumber string
		if o.Number != nil {
			qrNumber = *o.Number
		}
		qNumber := qrNumber
		if qNumber != "" {
			if err := r.SetQueryParam("number", qNumber); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
