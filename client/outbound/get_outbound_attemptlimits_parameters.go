// Code generated by go-swagger; DO NOT EDIT.

package outbound

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

// NewGetOutboundAttemptlimitsParams creates a new GetOutboundAttemptlimitsParams object
// with the default values initialized.
func NewGetOutboundAttemptlimitsParams() *GetOutboundAttemptlimitsParams {
	var (
		allowEmptyResultDefault = bool(false)
		filterTypeDefault       = string("Prefix")
		pageNumberDefault       = int32(1)
		pageSizeDefault         = int32(25)
		sortOrderDefault        = string("a")
	)
	return &GetOutboundAttemptlimitsParams{
		AllowEmptyResult: &allowEmptyResultDefault,
		FilterType:       &filterTypeDefault,
		PageNumber:       &pageNumberDefault,
		PageSize:         &pageSizeDefault,
		SortOrder:        &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundAttemptlimitsParamsWithTimeout creates a new GetOutboundAttemptlimitsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutboundAttemptlimitsParamsWithTimeout(timeout time.Duration) *GetOutboundAttemptlimitsParams {
	var (
		allowEmptyResultDefault = bool(false)
		filterTypeDefault       = string("Prefix")
		pageNumberDefault       = int32(1)
		pageSizeDefault         = int32(25)
		sortOrderDefault        = string("a")
	)
	return &GetOutboundAttemptlimitsParams{
		AllowEmptyResult: &allowEmptyResultDefault,
		FilterType:       &filterTypeDefault,
		PageNumber:       &pageNumberDefault,
		PageSize:         &pageSizeDefault,
		SortOrder:        &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetOutboundAttemptlimitsParamsWithContext creates a new GetOutboundAttemptlimitsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutboundAttemptlimitsParamsWithContext(ctx context.Context) *GetOutboundAttemptlimitsParams {
	var (
		allowEmptyResultDefault = bool(false)
		filterTypeDefault       = string("Prefix")
		pageNumberDefault       = int32(1)
		pageSizeDefault         = int32(25)
		sortOrderDefault        = string("a")
	)
	return &GetOutboundAttemptlimitsParams{
		AllowEmptyResult: &allowEmptyResultDefault,
		FilterType:       &filterTypeDefault,
		PageNumber:       &pageNumberDefault,
		PageSize:         &pageSizeDefault,
		SortOrder:        &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetOutboundAttemptlimitsParamsWithHTTPClient creates a new GetOutboundAttemptlimitsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutboundAttemptlimitsParamsWithHTTPClient(client *http.Client) *GetOutboundAttemptlimitsParams {
	var (
		allowEmptyResultDefault = bool(false)
		filterTypeDefault       = string("Prefix")
		pageNumberDefault       = int32(1)
		pageSizeDefault         = int32(25)
		sortOrderDefault        = string("a")
	)
	return &GetOutboundAttemptlimitsParams{
		AllowEmptyResult: &allowEmptyResultDefault,
		FilterType:       &filterTypeDefault,
		PageNumber:       &pageNumberDefault,
		PageSize:         &pageSizeDefault,
		SortOrder:        &sortOrderDefault,
		HTTPClient:       client,
	}
}

/*GetOutboundAttemptlimitsParams contains all the parameters to send to the API endpoint
for the get outbound attemptlimits operation typically these are written to a http.Request
*/
type GetOutboundAttemptlimitsParams struct {

	/*AllowEmptyResult
	  Whether to return an empty page when there are no results for that page

	*/
	AllowEmptyResult *bool
	/*FilterType
	  Filter type

	*/
	FilterType *string
	/*Name
	  Name

	*/
	Name *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size. The max that will be returned is 100.

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

// WithTimeout adds the timeout to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) WithTimeout(timeout time.Duration) *GetOutboundAttemptlimitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) WithContext(ctx context.Context) *GetOutboundAttemptlimitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) WithHTTPClient(client *http.Client) *GetOutboundAttemptlimitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowEmptyResult adds the allowEmptyResult to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) WithAllowEmptyResult(allowEmptyResult *bool) *GetOutboundAttemptlimitsParams {
	o.SetAllowEmptyResult(allowEmptyResult)
	return o
}

// SetAllowEmptyResult adds the allowEmptyResult to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) SetAllowEmptyResult(allowEmptyResult *bool) {
	o.AllowEmptyResult = allowEmptyResult
}

// WithFilterType adds the filterType to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) WithFilterType(filterType *string) *GetOutboundAttemptlimitsParams {
	o.SetFilterType(filterType)
	return o
}

// SetFilterType adds the filterType to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) SetFilterType(filterType *string) {
	o.FilterType = filterType
}

// WithName adds the name to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) WithName(name *string) *GetOutboundAttemptlimitsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) WithPageNumber(pageNumber *int32) *GetOutboundAttemptlimitsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) WithPageSize(pageSize *int32) *GetOutboundAttemptlimitsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) WithSortBy(sortBy *string) *GetOutboundAttemptlimitsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) WithSortOrder(sortOrder *string) *GetOutboundAttemptlimitsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get outbound attemptlimits params
func (o *GetOutboundAttemptlimitsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundAttemptlimitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowEmptyResult != nil {

		// query param allowEmptyResult
		var qrAllowEmptyResult bool
		if o.AllowEmptyResult != nil {
			qrAllowEmptyResult = *o.AllowEmptyResult
		}
		qAllowEmptyResult := swag.FormatBool(qrAllowEmptyResult)
		if qAllowEmptyResult != "" {
			if err := r.SetQueryParam("allowEmptyResult", qAllowEmptyResult); err != nil {
				return err
			}
		}

	}

	if o.FilterType != nil {

		// query param filterType
		var qrFilterType string
		if o.FilterType != nil {
			qrFilterType = *o.FilterType
		}
		qFilterType := qrFilterType
		if qFilterType != "" {
			if err := r.SetQueryParam("filterType", qFilterType); err != nil {
				return err
			}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}