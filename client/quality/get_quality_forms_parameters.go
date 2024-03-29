// Code generated by go-swagger; DO NOT EDIT.

package quality

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

// NewGetQualityFormsParams creates a new GetQualityFormsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetQualityFormsParams() *GetQualityFormsParams {
	return &GetQualityFormsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetQualityFormsParamsWithTimeout creates a new GetQualityFormsParams object
// with the ability to set a timeout on a request.
func NewGetQualityFormsParamsWithTimeout(timeout time.Duration) *GetQualityFormsParams {
	return &GetQualityFormsParams{
		timeout: timeout,
	}
}

// NewGetQualityFormsParamsWithContext creates a new GetQualityFormsParams object
// with the ability to set a context for a request.
func NewGetQualityFormsParamsWithContext(ctx context.Context) *GetQualityFormsParams {
	return &GetQualityFormsParams{
		Context: ctx,
	}
}

// NewGetQualityFormsParamsWithHTTPClient creates a new GetQualityFormsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetQualityFormsParamsWithHTTPClient(client *http.Client) *GetQualityFormsParams {
	return &GetQualityFormsParams{
		HTTPClient: client,
	}
}

/*
GetQualityFormsParams contains all the parameters to send to the API endpoint

	for the get quality forms operation.

	Typically these are written to a http.Request.
*/
type GetQualityFormsParams struct {

	/* Expand.

	   If 'expand=publishHistory', then each unpublished evaluation form includes a listing of its published versions
	*/
	Expand *string

	/* Name.

	   Name
	*/
	Name *string

	/* NextPage.

	   next page token
	*/
	NextPage *string

	/* PageNumber.

	   The page number requested

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   The total page size requested

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	/* PreviousPage.

	   Previous page token
	*/
	PreviousPage *string

	/* SortBy.

	   variable name requested to sort by
	*/
	SortBy *string

	/* SortOrder.

	   Order to sort results, either asc or desc
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get quality forms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQualityFormsParams) WithDefaults() *GetQualityFormsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get quality forms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQualityFormsParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetQualityFormsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get quality forms params
func (o *GetQualityFormsParams) WithTimeout(timeout time.Duration) *GetQualityFormsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get quality forms params
func (o *GetQualityFormsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get quality forms params
func (o *GetQualityFormsParams) WithContext(ctx context.Context) *GetQualityFormsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get quality forms params
func (o *GetQualityFormsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get quality forms params
func (o *GetQualityFormsParams) WithHTTPClient(client *http.Client) *GetQualityFormsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get quality forms params
func (o *GetQualityFormsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get quality forms params
func (o *GetQualityFormsParams) WithExpand(expand *string) *GetQualityFormsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get quality forms params
func (o *GetQualityFormsParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithName adds the name to the get quality forms params
func (o *GetQualityFormsParams) WithName(name *string) *GetQualityFormsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get quality forms params
func (o *GetQualityFormsParams) SetName(name *string) {
	o.Name = name
}

// WithNextPage adds the nextPage to the get quality forms params
func (o *GetQualityFormsParams) WithNextPage(nextPage *string) *GetQualityFormsParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the get quality forms params
func (o *GetQualityFormsParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithPageNumber adds the pageNumber to the get quality forms params
func (o *GetQualityFormsParams) WithPageNumber(pageNumber *int32) *GetQualityFormsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get quality forms params
func (o *GetQualityFormsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get quality forms params
func (o *GetQualityFormsParams) WithPageSize(pageSize *int32) *GetQualityFormsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get quality forms params
func (o *GetQualityFormsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPreviousPage adds the previousPage to the get quality forms params
func (o *GetQualityFormsParams) WithPreviousPage(previousPage *string) *GetQualityFormsParams {
	o.SetPreviousPage(previousPage)
	return o
}

// SetPreviousPage adds the previousPage to the get quality forms params
func (o *GetQualityFormsParams) SetPreviousPage(previousPage *string) {
	o.PreviousPage = previousPage
}

// WithSortBy adds the sortBy to the get quality forms params
func (o *GetQualityFormsParams) WithSortBy(sortBy *string) *GetQualityFormsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get quality forms params
func (o *GetQualityFormsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get quality forms params
func (o *GetQualityFormsParams) WithSortOrder(sortOrder *string) *GetQualityFormsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get quality forms params
func (o *GetQualityFormsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetQualityFormsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// query param expand
		var qrExpand string

		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {

			if err := r.SetQueryParam("expand", qExpand); err != nil {
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

	if o.NextPage != nil {

		// query param nextPage
		var qrNextPage string

		if o.NextPage != nil {
			qrNextPage = *o.NextPage
		}
		qNextPage := qrNextPage
		if qNextPage != "" {

			if err := r.SetQueryParam("nextPage", qNextPage); err != nil {
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

	if o.PreviousPage != nil {

		// query param previousPage
		var qrPreviousPage string

		if o.PreviousPage != nil {
			qrPreviousPage = *o.PreviousPage
		}
		qPreviousPage := qrPreviousPage
		if qPreviousPage != "" {

			if err := r.SetQueryParam("previousPage", qPreviousPage); err != nil {
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
