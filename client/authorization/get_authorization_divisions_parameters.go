// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// NewGetAuthorizationDivisionsParams creates a new GetAuthorizationDivisionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAuthorizationDivisionsParams() *GetAuthorizationDivisionsParams {
	return &GetAuthorizationDivisionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationDivisionsParamsWithTimeout creates a new GetAuthorizationDivisionsParams object
// with the ability to set a timeout on a request.
func NewGetAuthorizationDivisionsParamsWithTimeout(timeout time.Duration) *GetAuthorizationDivisionsParams {
	return &GetAuthorizationDivisionsParams{
		timeout: timeout,
	}
}

// NewGetAuthorizationDivisionsParamsWithContext creates a new GetAuthorizationDivisionsParams object
// with the ability to set a context for a request.
func NewGetAuthorizationDivisionsParamsWithContext(ctx context.Context) *GetAuthorizationDivisionsParams {
	return &GetAuthorizationDivisionsParams{
		Context: ctx,
	}
}

// NewGetAuthorizationDivisionsParamsWithHTTPClient creates a new GetAuthorizationDivisionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAuthorizationDivisionsParamsWithHTTPClient(client *http.Client) *GetAuthorizationDivisionsParams {
	return &GetAuthorizationDivisionsParams{
		HTTPClient: client,
	}
}

/*
GetAuthorizationDivisionsParams contains all the parameters to send to the API endpoint

	for the get authorization divisions operation.

	Typically these are written to a http.Request.
*/
type GetAuthorizationDivisionsParams struct {

	/* Expand.

	   variable name requested by expand list
	*/
	Expand []string

	/* ID.

	   Optionally request specific divisions by their IDs
	*/
	ID []string

	/* Name.

	   Search term to filter by division name
	*/
	Name *string

	/* NextPage.

	   next page token
	*/
	NextPage *string

	/* ObjectCount.

	   Include the count of objects contained in the division
	*/
	ObjectCount *bool

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

// WithDefaults hydrates default values in the get authorization divisions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationDivisionsParams) WithDefaults() *GetAuthorizationDivisionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get authorization divisions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationDivisionsParams) SetDefaults() {
	var (
		objectCountDefault = bool(false)

		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetAuthorizationDivisionsParams{
		ObjectCount: &objectCountDefault,
		PageNumber:  &pageNumberDefault,
		PageSize:    &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) WithTimeout(timeout time.Duration) *GetAuthorizationDivisionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) WithContext(ctx context.Context) *GetAuthorizationDivisionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) WithHTTPClient(client *http.Client) *GetAuthorizationDivisionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) WithExpand(expand []string) *GetAuthorizationDivisionsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithID adds the id to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) WithID(id []string) *GetAuthorizationDivisionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) WithName(name *string) *GetAuthorizationDivisionsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) SetName(name *string) {
	o.Name = name
}

// WithNextPage adds the nextPage to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) WithNextPage(nextPage *string) *GetAuthorizationDivisionsParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithObjectCount adds the objectCount to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) WithObjectCount(objectCount *bool) *GetAuthorizationDivisionsParams {
	o.SetObjectCount(objectCount)
	return o
}

// SetObjectCount adds the objectCount to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) SetObjectCount(objectCount *bool) {
	o.ObjectCount = objectCount
}

// WithPageNumber adds the pageNumber to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) WithPageNumber(pageNumber *int32) *GetAuthorizationDivisionsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) WithPageSize(pageSize *int32) *GetAuthorizationDivisionsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPreviousPage adds the previousPage to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) WithPreviousPage(previousPage *string) *GetAuthorizationDivisionsParams {
	o.SetPreviousPage(previousPage)
	return o
}

// SetPreviousPage adds the previousPage to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) SetPreviousPage(previousPage *string) {
	o.PreviousPage = previousPage
}

// WithSortBy adds the sortBy to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) WithSortBy(sortBy *string) *GetAuthorizationDivisionsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get authorization divisions params
func (o *GetAuthorizationDivisionsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationDivisionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
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

	if o.ObjectCount != nil {

		// query param objectCount
		var qrObjectCount bool

		if o.ObjectCount != nil {
			qrObjectCount = *o.ObjectCount
		}
		qObjectCount := swag.FormatBool(qrObjectCount)
		if qObjectCount != "" {

			if err := r.SetQueryParam("objectCount", qObjectCount); err != nil {
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

// bindParamGetAuthorizationDivisions binds the parameter expand
func (o *GetAuthorizationDivisionsParams) bindParamExpand(formats strfmt.Registry) []string {
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

// bindParamGetAuthorizationDivisions binds the parameter id
func (o *GetAuthorizationDivisionsParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}
