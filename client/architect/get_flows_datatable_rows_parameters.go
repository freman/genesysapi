// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewGetFlowsDatatableRowsParams creates a new GetFlowsDatatableRowsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFlowsDatatableRowsParams() *GetFlowsDatatableRowsParams {
	return &GetFlowsDatatableRowsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowsDatatableRowsParamsWithTimeout creates a new GetFlowsDatatableRowsParams object
// with the ability to set a timeout on a request.
func NewGetFlowsDatatableRowsParamsWithTimeout(timeout time.Duration) *GetFlowsDatatableRowsParams {
	return &GetFlowsDatatableRowsParams{
		timeout: timeout,
	}
}

// NewGetFlowsDatatableRowsParamsWithContext creates a new GetFlowsDatatableRowsParams object
// with the ability to set a context for a request.
func NewGetFlowsDatatableRowsParamsWithContext(ctx context.Context) *GetFlowsDatatableRowsParams {
	return &GetFlowsDatatableRowsParams{
		Context: ctx,
	}
}

// NewGetFlowsDatatableRowsParamsWithHTTPClient creates a new GetFlowsDatatableRowsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFlowsDatatableRowsParamsWithHTTPClient(client *http.Client) *GetFlowsDatatableRowsParams {
	return &GetFlowsDatatableRowsParams{
		HTTPClient: client,
	}
}

/*
GetFlowsDatatableRowsParams contains all the parameters to send to the API endpoint

	for the get flows datatable rows operation.

	Typically these are written to a http.Request.
*/
type GetFlowsDatatableRowsParams struct {

	/* DatatableID.

	   id of datatable
	*/
	DatatableID string

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

	/* Showbrief.

	   If true returns just the key value of the row

	   Default: true
	*/
	Showbrief *bool

	/* SortOrder.

	   Sort order

	   Default: "ascending"
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get flows datatable rows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowsDatatableRowsParams) WithDefaults() *GetFlowsDatatableRowsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get flows datatable rows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowsDatatableRowsParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		showbriefDefault = bool(true)

		sortOrderDefault = string("ascending")
	)

	val := GetFlowsDatatableRowsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		Showbrief:  &showbriefDefault,
		SortOrder:  &sortOrderDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) WithTimeout(timeout time.Duration) *GetFlowsDatatableRowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) WithContext(ctx context.Context) *GetFlowsDatatableRowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) WithHTTPClient(client *http.Client) *GetFlowsDatatableRowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatatableID adds the datatableID to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) WithDatatableID(datatableID string) *GetFlowsDatatableRowsParams {
	o.SetDatatableID(datatableID)
	return o
}

// SetDatatableID adds the datatableId to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) SetDatatableID(datatableID string) {
	o.DatatableID = datatableID
}

// WithPageNumber adds the pageNumber to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) WithPageNumber(pageNumber *int32) *GetFlowsDatatableRowsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) WithPageSize(pageSize *int32) *GetFlowsDatatableRowsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithShowbrief adds the showbrief to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) WithShowbrief(showbrief *bool) *GetFlowsDatatableRowsParams {
	o.SetShowbrief(showbrief)
	return o
}

// SetShowbrief adds the showbrief to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) SetShowbrief(showbrief *bool) {
	o.Showbrief = showbrief
}

// WithSortOrder adds the sortOrder to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) WithSortOrder(sortOrder *string) *GetFlowsDatatableRowsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get flows datatable rows params
func (o *GetFlowsDatatableRowsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowsDatatableRowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datatableId
	if err := r.SetPathParam("datatableId", o.DatatableID); err != nil {
		return err
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

	if o.Showbrief != nil {

		// query param showbrief
		var qrShowbrief bool

		if o.Showbrief != nil {
			qrShowbrief = *o.Showbrief
		}
		qShowbrief := swag.FormatBool(qrShowbrief)
		if qShowbrief != "" {

			if err := r.SetQueryParam("showbrief", qShowbrief); err != nil {
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
