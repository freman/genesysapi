// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

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

// NewGetTelephonyProvidersEdgesEdgegroupsParams creates a new GetTelephonyProvidersEdgesEdgegroupsParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesEdgegroupsParams() *GetTelephonyProvidersEdgesEdgegroupsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
	)
	return &GetTelephonyProvidersEdgesEdgegroupsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesEdgegroupsParamsWithTimeout creates a new GetTelephonyProvidersEdgesEdgegroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesEdgegroupsParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesEdgegroupsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
	)
	return &GetTelephonyProvidersEdgesEdgegroupsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesEdgegroupsParamsWithContext creates a new GetTelephonyProvidersEdgesEdgegroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesEdgegroupsParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesEdgegroupsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
	)
	return &GetTelephonyProvidersEdgesEdgegroupsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesEdgegroupsParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesEdgegroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesEdgegroupsParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesEdgegroupsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
	)
	return &GetTelephonyProvidersEdgesEdgegroupsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesEdgegroupsParams contains all the parameters to send to the API endpoint
for the get telephony providers edges edgegroups operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesEdgegroupsParams struct {

	/*Managed
	  Filter by managed

	*/
	Managed *bool
	/*Name
	  Name

	*/
	Name *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*SortBy
	  Sort by

	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesEdgegroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesEdgegroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesEdgegroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManaged adds the managed to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) WithManaged(managed *bool) *GetTelephonyProvidersEdgesEdgegroupsParams {
	o.SetManaged(managed)
	return o
}

// SetManaged adds the managed to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) SetManaged(managed *bool) {
	o.Managed = managed
}

// WithName adds the name to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) WithName(name *string) *GetTelephonyProvidersEdgesEdgegroupsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) WithPageNumber(pageNumber *int32) *GetTelephonyProvidersEdgesEdgegroupsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) WithPageSize(pageSize *int32) *GetTelephonyProvidersEdgesEdgegroupsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) WithSortBy(sortBy *string) *GetTelephonyProvidersEdgesEdgegroupsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get telephony providers edges edgegroups params
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesEdgegroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Managed != nil {

		// query param managed
		var qrManaged bool
		if o.Managed != nil {
			qrManaged = *o.Managed
		}
		qManaged := swag.FormatBool(qrManaged)
		if qManaged != "" {
			if err := r.SetQueryParam("managed", qManaged); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
