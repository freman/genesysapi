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

// NewGetIntegrationParams creates a new GetIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIntegrationParams() *GetIntegrationParams {
	return &GetIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationParamsWithTimeout creates a new GetIntegrationParams object
// with the ability to set a timeout on a request.
func NewGetIntegrationParamsWithTimeout(timeout time.Duration) *GetIntegrationParams {
	return &GetIntegrationParams{
		timeout: timeout,
	}
}

// NewGetIntegrationParamsWithContext creates a new GetIntegrationParams object
// with the ability to set a context for a request.
func NewGetIntegrationParamsWithContext(ctx context.Context) *GetIntegrationParams {
	return &GetIntegrationParams{
		Context: ctx,
	}
}

// NewGetIntegrationParamsWithHTTPClient creates a new GetIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIntegrationParamsWithHTTPClient(client *http.Client) *GetIntegrationParams {
	return &GetIntegrationParams{
		HTTPClient: client,
	}
}

/*
GetIntegrationParams contains all the parameters to send to the API endpoint

	for the get integration operation.

	Typically these are written to a http.Request.
*/
type GetIntegrationParams struct {

	/* Expand.

	   variable name requested by expand list
	*/
	Expand []string

	/* IntegrationID.

	   Integration Id
	*/
	IntegrationID string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationParams) WithDefaults() *GetIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetIntegrationParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get integration params
func (o *GetIntegrationParams) WithTimeout(timeout time.Duration) *GetIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integration params
func (o *GetIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integration params
func (o *GetIntegrationParams) WithContext(ctx context.Context) *GetIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integration params
func (o *GetIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integration params
func (o *GetIntegrationParams) WithHTTPClient(client *http.Client) *GetIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integration params
func (o *GetIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get integration params
func (o *GetIntegrationParams) WithExpand(expand []string) *GetIntegrationParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get integration params
func (o *GetIntegrationParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithIntegrationID adds the integrationID to the get integration params
func (o *GetIntegrationParams) WithIntegrationID(integrationID string) *GetIntegrationParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the get integration params
func (o *GetIntegrationParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WithNextPage adds the nextPage to the get integration params
func (o *GetIntegrationParams) WithNextPage(nextPage *string) *GetIntegrationParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the get integration params
func (o *GetIntegrationParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithPageNumber adds the pageNumber to the get integration params
func (o *GetIntegrationParams) WithPageNumber(pageNumber *int32) *GetIntegrationParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get integration params
func (o *GetIntegrationParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get integration params
func (o *GetIntegrationParams) WithPageSize(pageSize *int32) *GetIntegrationParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get integration params
func (o *GetIntegrationParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPreviousPage adds the previousPage to the get integration params
func (o *GetIntegrationParams) WithPreviousPage(previousPage *string) *GetIntegrationParams {
	o.SetPreviousPage(previousPage)
	return o
}

// SetPreviousPage adds the previousPage to the get integration params
func (o *GetIntegrationParams) SetPreviousPage(previousPage *string) {
	o.PreviousPage = previousPage
}

// WithSortBy adds the sortBy to the get integration params
func (o *GetIntegrationParams) WithSortBy(sortBy *string) *GetIntegrationParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get integration params
func (o *GetIntegrationParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param integrationId
	if err := r.SetPathParam("integrationId", o.IntegrationID); err != nil {
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

// bindParamGetIntegration binds the parameter expand
func (o *GetIntegrationParams) bindParamExpand(formats strfmt.Registry) []string {
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
