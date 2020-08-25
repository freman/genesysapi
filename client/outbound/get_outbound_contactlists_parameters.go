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

// NewGetOutboundContactlistsParams creates a new GetOutboundContactlistsParams object
// with the default values initialized.
func NewGetOutboundContactlistsParams() *GetOutboundContactlistsParams {
	var (
		allowEmptyResultDefault    = bool(false)
		filterTypeDefault          = string("Prefix")
		includeImportStatusDefault = bool(false)
		includeSizeDefault         = bool(false)
		pageNumberDefault          = int32(1)
		pageSizeDefault            = int32(25)
		sortOrderDefault           = string("a")
	)
	return &GetOutboundContactlistsParams{
		AllowEmptyResult:    &allowEmptyResultDefault,
		FilterType:          &filterTypeDefault,
		IncludeImportStatus: &includeImportStatusDefault,
		IncludeSize:         &includeSizeDefault,
		PageNumber:          &pageNumberDefault,
		PageSize:            &pageSizeDefault,
		SortOrder:           &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundContactlistsParamsWithTimeout creates a new GetOutboundContactlistsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutboundContactlistsParamsWithTimeout(timeout time.Duration) *GetOutboundContactlistsParams {
	var (
		allowEmptyResultDefault    = bool(false)
		filterTypeDefault          = string("Prefix")
		includeImportStatusDefault = bool(false)
		includeSizeDefault         = bool(false)
		pageNumberDefault          = int32(1)
		pageSizeDefault            = int32(25)
		sortOrderDefault           = string("a")
	)
	return &GetOutboundContactlistsParams{
		AllowEmptyResult:    &allowEmptyResultDefault,
		FilterType:          &filterTypeDefault,
		IncludeImportStatus: &includeImportStatusDefault,
		IncludeSize:         &includeSizeDefault,
		PageNumber:          &pageNumberDefault,
		PageSize:            &pageSizeDefault,
		SortOrder:           &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetOutboundContactlistsParamsWithContext creates a new GetOutboundContactlistsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutboundContactlistsParamsWithContext(ctx context.Context) *GetOutboundContactlistsParams {
	var (
		allowEmptyResultDefault    = bool(false)
		filterTypeDefault          = string("Prefix")
		includeImportStatusDefault = bool(false)
		includeSizeDefault         = bool(false)
		pageNumberDefault          = int32(1)
		pageSizeDefault            = int32(25)
		sortOrderDefault           = string("a")
	)
	return &GetOutboundContactlistsParams{
		AllowEmptyResult:    &allowEmptyResultDefault,
		FilterType:          &filterTypeDefault,
		IncludeImportStatus: &includeImportStatusDefault,
		IncludeSize:         &includeSizeDefault,
		PageNumber:          &pageNumberDefault,
		PageSize:            &pageSizeDefault,
		SortOrder:           &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetOutboundContactlistsParamsWithHTTPClient creates a new GetOutboundContactlistsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutboundContactlistsParamsWithHTTPClient(client *http.Client) *GetOutboundContactlistsParams {
	var (
		allowEmptyResultDefault    = bool(false)
		filterTypeDefault          = string("Prefix")
		includeImportStatusDefault = bool(false)
		includeSizeDefault         = bool(false)
		pageNumberDefault          = int32(1)
		pageSizeDefault            = int32(25)
		sortOrderDefault           = string("a")
	)
	return &GetOutboundContactlistsParams{
		AllowEmptyResult:    &allowEmptyResultDefault,
		FilterType:          &filterTypeDefault,
		IncludeImportStatus: &includeImportStatusDefault,
		IncludeSize:         &includeSizeDefault,
		PageNumber:          &pageNumberDefault,
		PageSize:            &pageSizeDefault,
		SortOrder:           &sortOrderDefault,
		HTTPClient:          client,
	}
}

/*GetOutboundContactlistsParams contains all the parameters to send to the API endpoint
for the get outbound contactlists operation typically these are written to a http.Request
*/
type GetOutboundContactlistsParams struct {

	/*AllowEmptyResult
	  Whether to return an empty page when there are no results for that page

	*/
	AllowEmptyResult *bool
	/*DivisionID
	  Division ID(s)

	*/
	DivisionID []string
	/*FilterType
	  Filter type

	*/
	FilterType *string
	/*ID
	  id

	*/
	ID []string
	/*IncludeImportStatus
	  Include import status

	*/
	IncludeImportStatus *bool
	/*IncludeSize
	  Include size

	*/
	IncludeSize *bool
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

// WithTimeout adds the timeout to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithTimeout(timeout time.Duration) *GetOutboundContactlistsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithContext(ctx context.Context) *GetOutboundContactlistsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithHTTPClient(client *http.Client) *GetOutboundContactlistsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowEmptyResult adds the allowEmptyResult to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithAllowEmptyResult(allowEmptyResult *bool) *GetOutboundContactlistsParams {
	o.SetAllowEmptyResult(allowEmptyResult)
	return o
}

// SetAllowEmptyResult adds the allowEmptyResult to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetAllowEmptyResult(allowEmptyResult *bool) {
	o.AllowEmptyResult = allowEmptyResult
}

// WithDivisionID adds the divisionID to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithDivisionID(divisionID []string) *GetOutboundContactlistsParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetDivisionID(divisionID []string) {
	o.DivisionID = divisionID
}

// WithFilterType adds the filterType to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithFilterType(filterType *string) *GetOutboundContactlistsParams {
	o.SetFilterType(filterType)
	return o
}

// SetFilterType adds the filterType to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetFilterType(filterType *string) {
	o.FilterType = filterType
}

// WithID adds the id to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithID(id []string) *GetOutboundContactlistsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetID(id []string) {
	o.ID = id
}

// WithIncludeImportStatus adds the includeImportStatus to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithIncludeImportStatus(includeImportStatus *bool) *GetOutboundContactlistsParams {
	o.SetIncludeImportStatus(includeImportStatus)
	return o
}

// SetIncludeImportStatus adds the includeImportStatus to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetIncludeImportStatus(includeImportStatus *bool) {
	o.IncludeImportStatus = includeImportStatus
}

// WithIncludeSize adds the includeSize to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithIncludeSize(includeSize *bool) *GetOutboundContactlistsParams {
	o.SetIncludeSize(includeSize)
	return o
}

// SetIncludeSize adds the includeSize to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetIncludeSize(includeSize *bool) {
	o.IncludeSize = includeSize
}

// WithName adds the name to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithName(name *string) *GetOutboundContactlistsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithPageNumber(pageNumber *int32) *GetOutboundContactlistsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithPageSize(pageSize *int32) *GetOutboundContactlistsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithSortBy(sortBy *string) *GetOutboundContactlistsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) WithSortOrder(sortOrder *string) *GetOutboundContactlistsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get outbound contactlists params
func (o *GetOutboundContactlistsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundContactlistsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesDivisionID := o.DivisionID

	joinedDivisionID := swag.JoinByFormat(valuesDivisionID, "multi")
	// query array param divisionId
	if err := r.SetQueryParam("divisionId", joinedDivisionID...); err != nil {
		return err
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

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
		return err
	}

	if o.IncludeImportStatus != nil {

		// query param includeImportStatus
		var qrIncludeImportStatus bool
		if o.IncludeImportStatus != nil {
			qrIncludeImportStatus = *o.IncludeImportStatus
		}
		qIncludeImportStatus := swag.FormatBool(qrIncludeImportStatus)
		if qIncludeImportStatus != "" {
			if err := r.SetQueryParam("includeImportStatus", qIncludeImportStatus); err != nil {
				return err
			}
		}

	}

	if o.IncludeSize != nil {

		// query param includeSize
		var qrIncludeSize bool
		if o.IncludeSize != nil {
			qrIncludeSize = *o.IncludeSize
		}
		qIncludeSize := swag.FormatBool(qrIncludeSize)
		if qIncludeSize != "" {
			if err := r.SetQueryParam("includeSize", qIncludeSize); err != nil {
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