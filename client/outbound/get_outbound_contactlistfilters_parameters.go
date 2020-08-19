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

// NewGetOutboundContactlistfiltersParams creates a new GetOutboundContactlistfiltersParams object
// with the default values initialized.
func NewGetOutboundContactlistfiltersParams() *GetOutboundContactlistfiltersParams {
	var (
		allowEmptyResultDefault = bool(false)
		filterTypeDefault       = string("Prefix")
		pageNumberDefault       = int32(1)
		pageSizeDefault         = int32(25)
		sortOrderDefault        = string("a")
	)
	return &GetOutboundContactlistfiltersParams{
		AllowEmptyResult: &allowEmptyResultDefault,
		FilterType:       &filterTypeDefault,
		PageNumber:       &pageNumberDefault,
		PageSize:         &pageSizeDefault,
		SortOrder:        &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundContactlistfiltersParamsWithTimeout creates a new GetOutboundContactlistfiltersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutboundContactlistfiltersParamsWithTimeout(timeout time.Duration) *GetOutboundContactlistfiltersParams {
	var (
		allowEmptyResultDefault = bool(false)
		filterTypeDefault       = string("Prefix")
		pageNumberDefault       = int32(1)
		pageSizeDefault         = int32(25)
		sortOrderDefault        = string("a")
	)
	return &GetOutboundContactlistfiltersParams{
		AllowEmptyResult: &allowEmptyResultDefault,
		FilterType:       &filterTypeDefault,
		PageNumber:       &pageNumberDefault,
		PageSize:         &pageSizeDefault,
		SortOrder:        &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetOutboundContactlistfiltersParamsWithContext creates a new GetOutboundContactlistfiltersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutboundContactlistfiltersParamsWithContext(ctx context.Context) *GetOutboundContactlistfiltersParams {
	var (
		allowEmptyResultDefault = bool(false)
		filterTypeDefault       = string("Prefix")
		pageNumberDefault       = int32(1)
		pageSizeDefault         = int32(25)
		sortOrderDefault        = string("a")
	)
	return &GetOutboundContactlistfiltersParams{
		AllowEmptyResult: &allowEmptyResultDefault,
		FilterType:       &filterTypeDefault,
		PageNumber:       &pageNumberDefault,
		PageSize:         &pageSizeDefault,
		SortOrder:        &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetOutboundContactlistfiltersParamsWithHTTPClient creates a new GetOutboundContactlistfiltersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutboundContactlistfiltersParamsWithHTTPClient(client *http.Client) *GetOutboundContactlistfiltersParams {
	var (
		allowEmptyResultDefault = bool(false)
		filterTypeDefault       = string("Prefix")
		pageNumberDefault       = int32(1)
		pageSizeDefault         = int32(25)
		sortOrderDefault        = string("a")
	)
	return &GetOutboundContactlistfiltersParams{
		AllowEmptyResult: &allowEmptyResultDefault,
		FilterType:       &filterTypeDefault,
		PageNumber:       &pageNumberDefault,
		PageSize:         &pageSizeDefault,
		SortOrder:        &sortOrderDefault,
		HTTPClient:       client,
	}
}

/*GetOutboundContactlistfiltersParams contains all the parameters to send to the API endpoint
for the get outbound contactlistfilters operation typically these are written to a http.Request
*/
type GetOutboundContactlistfiltersParams struct {

	/*AllowEmptyResult
	  Whether to return an empty page when there are no results for that page

	*/
	AllowEmptyResult *bool
	/*ContactListID
	  Contact List ID

	*/
	ContactListID *string
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

// WithTimeout adds the timeout to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) WithTimeout(timeout time.Duration) *GetOutboundContactlistfiltersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) WithContext(ctx context.Context) *GetOutboundContactlistfiltersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) WithHTTPClient(client *http.Client) *GetOutboundContactlistfiltersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowEmptyResult adds the allowEmptyResult to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) WithAllowEmptyResult(allowEmptyResult *bool) *GetOutboundContactlistfiltersParams {
	o.SetAllowEmptyResult(allowEmptyResult)
	return o
}

// SetAllowEmptyResult adds the allowEmptyResult to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) SetAllowEmptyResult(allowEmptyResult *bool) {
	o.AllowEmptyResult = allowEmptyResult
}

// WithContactListID adds the contactListID to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) WithContactListID(contactListID *string) *GetOutboundContactlistfiltersParams {
	o.SetContactListID(contactListID)
	return o
}

// SetContactListID adds the contactListId to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) SetContactListID(contactListID *string) {
	o.ContactListID = contactListID
}

// WithFilterType adds the filterType to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) WithFilterType(filterType *string) *GetOutboundContactlistfiltersParams {
	o.SetFilterType(filterType)
	return o
}

// SetFilterType adds the filterType to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) SetFilterType(filterType *string) {
	o.FilterType = filterType
}

// WithName adds the name to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) WithName(name *string) *GetOutboundContactlistfiltersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) WithPageNumber(pageNumber *int32) *GetOutboundContactlistfiltersParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) WithPageSize(pageSize *int32) *GetOutboundContactlistfiltersParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) WithSortBy(sortBy *string) *GetOutboundContactlistfiltersParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) WithSortOrder(sortOrder *string) *GetOutboundContactlistfiltersParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get outbound contactlistfilters params
func (o *GetOutboundContactlistfiltersParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundContactlistfiltersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ContactListID != nil {

		// query param contactListId
		var qrContactListID string
		if o.ContactListID != nil {
			qrContactListID = *o.ContactListID
		}
		qContactListID := qrContactListID
		if qContactListID != "" {
			if err := r.SetQueryParam("contactListId", qContactListID); err != nil {
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
