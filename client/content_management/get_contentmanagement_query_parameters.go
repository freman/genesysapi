// Code generated by go-swagger; DO NOT EDIT.

package content_management

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

// NewGetContentmanagementQueryParams creates a new GetContentmanagementQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContentmanagementQueryParams() *GetContentmanagementQueryParams {
	return &GetContentmanagementQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentmanagementQueryParamsWithTimeout creates a new GetContentmanagementQueryParams object
// with the ability to set a timeout on a request.
func NewGetContentmanagementQueryParamsWithTimeout(timeout time.Duration) *GetContentmanagementQueryParams {
	return &GetContentmanagementQueryParams{
		timeout: timeout,
	}
}

// NewGetContentmanagementQueryParamsWithContext creates a new GetContentmanagementQueryParams object
// with the ability to set a context for a request.
func NewGetContentmanagementQueryParamsWithContext(ctx context.Context) *GetContentmanagementQueryParams {
	return &GetContentmanagementQueryParams{
		Context: ctx,
	}
}

// NewGetContentmanagementQueryParamsWithHTTPClient creates a new GetContentmanagementQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContentmanagementQueryParamsWithHTTPClient(client *http.Client) *GetContentmanagementQueryParams {
	return &GetContentmanagementQueryParams{
		HTTPClient: client,
	}
}

/*
GetContentmanagementQueryParams contains all the parameters to send to the API endpoint

	for the get contentmanagement query operation.

	Typically these are written to a http.Request.
*/
type GetContentmanagementQueryParams struct {

	/* Expand.

	   Which fields, if any, to expand.
	*/
	Expand []string

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

	/* QueryPhrase.

	   Phrase tokens are ANDed together over all searchable fields
	*/
	QueryPhrase string

	/* SortBy.

	   name or dateCreated

	   Default: "name"
	*/
	SortBy *string

	/* SortOrder.

	   ascending or descending

	   Default: "ascending"
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contentmanagement query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentmanagementQueryParams) WithDefaults() *GetContentmanagementQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contentmanagement query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentmanagementQueryParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		sortByDefault = string("name")

		sortOrderDefault = string("ascending")
	)

	val := GetContentmanagementQueryParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) WithTimeout(timeout time.Duration) *GetContentmanagementQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) WithContext(ctx context.Context) *GetContentmanagementQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) WithHTTPClient(client *http.Client) *GetContentmanagementQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) WithExpand(expand []string) *GetContentmanagementQueryParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithPageNumber adds the pageNumber to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) WithPageNumber(pageNumber *int32) *GetContentmanagementQueryParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) WithPageSize(pageSize *int32) *GetContentmanagementQueryParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithQueryPhrase adds the queryPhrase to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) WithQueryPhrase(queryPhrase string) *GetContentmanagementQueryParams {
	o.SetQueryPhrase(queryPhrase)
	return o
}

// SetQueryPhrase adds the queryPhrase to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) SetQueryPhrase(queryPhrase string) {
	o.QueryPhrase = queryPhrase
}

// WithSortBy adds the sortBy to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) WithSortBy(sortBy *string) *GetContentmanagementQueryParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) WithSortOrder(sortOrder *string) *GetContentmanagementQueryParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get contentmanagement query params
func (o *GetContentmanagementQueryParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentmanagementQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
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

	// query param queryPhrase
	qrQueryPhrase := o.QueryPhrase
	qQueryPhrase := qrQueryPhrase
	if qQueryPhrase != "" {

		if err := r.SetQueryParam("queryPhrase", qQueryPhrase); err != nil {
			return err
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

// bindParamGetContentmanagementQuery binds the parameter expand
func (o *GetContentmanagementQueryParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}
