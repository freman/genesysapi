// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetIntegrationsClientappsParams creates a new GetIntegrationsClientappsParams object
// with the default values initialized.
func NewGetIntegrationsClientappsParams() *GetIntegrationsClientappsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetIntegrationsClientappsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationsClientappsParamsWithTimeout creates a new GetIntegrationsClientappsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIntegrationsClientappsParamsWithTimeout(timeout time.Duration) *GetIntegrationsClientappsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetIntegrationsClientappsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetIntegrationsClientappsParamsWithContext creates a new GetIntegrationsClientappsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIntegrationsClientappsParamsWithContext(ctx context.Context) *GetIntegrationsClientappsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetIntegrationsClientappsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetIntegrationsClientappsParamsWithHTTPClient creates a new GetIntegrationsClientappsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIntegrationsClientappsParamsWithHTTPClient(client *http.Client) *GetIntegrationsClientappsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetIntegrationsClientappsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetIntegrationsClientappsParams contains all the parameters to send to the API endpoint
for the get integrations clientapps operation typically these are written to a http.Request
*/
type GetIntegrationsClientappsParams struct {

	/*Expand
	  variable name requested by expand list

	*/
	Expand []string
	/*NextPage
	  next page token

	*/
	NextPage *string
	/*PageNumber
	  The page number requested

	*/
	PageNumber *int32
	/*PageSize
	  The total page size requested

	*/
	PageSize *int32
	/*PreviousPage
	  Previous page token

	*/
	PreviousPage *string
	/*SortBy
	  variable name requested to sort by

	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) WithTimeout(timeout time.Duration) *GetIntegrationsClientappsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) WithContext(ctx context.Context) *GetIntegrationsClientappsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) WithHTTPClient(client *http.Client) *GetIntegrationsClientappsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) WithExpand(expand []string) *GetIntegrationsClientappsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithNextPage adds the nextPage to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) WithNextPage(nextPage *string) *GetIntegrationsClientappsParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithPageNumber adds the pageNumber to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) WithPageNumber(pageNumber *int32) *GetIntegrationsClientappsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) WithPageSize(pageSize *int32) *GetIntegrationsClientappsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPreviousPage adds the previousPage to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) WithPreviousPage(previousPage *string) *GetIntegrationsClientappsParams {
	o.SetPreviousPage(previousPage)
	return o
}

// SetPreviousPage adds the previousPage to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) SetPreviousPage(previousPage *string) {
	o.PreviousPage = previousPage
}

// WithSortBy adds the sortBy to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) WithSortBy(sortBy *string) *GetIntegrationsClientappsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get integrations clientapps params
func (o *GetIntegrationsClientappsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationsClientappsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}