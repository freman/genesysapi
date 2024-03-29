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

// NewGetAuthorizationPermissionsParams creates a new GetAuthorizationPermissionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAuthorizationPermissionsParams() *GetAuthorizationPermissionsParams {
	return &GetAuthorizationPermissionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationPermissionsParamsWithTimeout creates a new GetAuthorizationPermissionsParams object
// with the ability to set a timeout on a request.
func NewGetAuthorizationPermissionsParamsWithTimeout(timeout time.Duration) *GetAuthorizationPermissionsParams {
	return &GetAuthorizationPermissionsParams{
		timeout: timeout,
	}
}

// NewGetAuthorizationPermissionsParamsWithContext creates a new GetAuthorizationPermissionsParams object
// with the ability to set a context for a request.
func NewGetAuthorizationPermissionsParamsWithContext(ctx context.Context) *GetAuthorizationPermissionsParams {
	return &GetAuthorizationPermissionsParams{
		Context: ctx,
	}
}

// NewGetAuthorizationPermissionsParamsWithHTTPClient creates a new GetAuthorizationPermissionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAuthorizationPermissionsParamsWithHTTPClient(client *http.Client) *GetAuthorizationPermissionsParams {
	return &GetAuthorizationPermissionsParams{
		HTTPClient: client,
	}
}

/*
GetAuthorizationPermissionsParams contains all the parameters to send to the API endpoint

	for the get authorization permissions operation.

	Typically these are written to a http.Request.
*/
type GetAuthorizationPermissionsParams struct {

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

	/* Query.

	   Comma-separated list of permissions or domains to query
	*/
	Query *string

	/* QueryType.

	   Query filter type
	*/
	QueryType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get authorization permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationPermissionsParams) WithDefaults() *GetAuthorizationPermissionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get authorization permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationPermissionsParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetAuthorizationPermissionsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) WithTimeout(timeout time.Duration) *GetAuthorizationPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) WithContext(ctx context.Context) *GetAuthorizationPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) WithHTTPClient(client *http.Client) *GetAuthorizationPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) WithPageNumber(pageNumber *int32) *GetAuthorizationPermissionsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) WithPageSize(pageSize *int32) *GetAuthorizationPermissionsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithQuery adds the query to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) WithQuery(query *string) *GetAuthorizationPermissionsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) SetQuery(query *string) {
	o.Query = query
}

// WithQueryType adds the queryType to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) WithQueryType(queryType *string) *GetAuthorizationPermissionsParams {
	o.SetQueryType(queryType)
	return o
}

// SetQueryType adds the queryType to the get authorization permissions params
func (o *GetAuthorizationPermissionsParams) SetQueryType(queryType *string) {
	o.QueryType = queryType
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if o.QueryType != nil {

		// query param queryType
		var qrQueryType string

		if o.QueryType != nil {
			qrQueryType = *o.QueryType
		}
		qQueryType := qrQueryType
		if qQueryType != "" {

			if err := r.SetQueryParam("queryType", qQueryType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
