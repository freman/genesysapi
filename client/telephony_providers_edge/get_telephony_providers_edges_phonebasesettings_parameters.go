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

// NewGetTelephonyProvidersEdgesPhonebasesettingsParams creates a new GetTelephonyProvidersEdgesPhonebasesettingsParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesPhonebasesettingsParams() *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("ASC")
	)
	return &GetTelephonyProvidersEdgesPhonebasesettingsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsParamsWithTimeout creates a new GetTelephonyProvidersEdgesPhonebasesettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesPhonebasesettingsParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("ASC")
	)
	return &GetTelephonyProvidersEdgesPhonebasesettingsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsParamsWithContext creates a new GetTelephonyProvidersEdgesPhonebasesettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesPhonebasesettingsParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("ASC")
	)
	return &GetTelephonyProvidersEdgesPhonebasesettingsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesPhonebasesettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesPhonebasesettingsParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("ASC")
	)
	return &GetTelephonyProvidersEdgesPhonebasesettingsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsParams contains all the parameters to send to the API endpoint
for the get telephony providers edges phonebasesettings operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesPhonebasesettingsParams struct {

	/*Expand
	  Fields to expand in the response, comma-separated

	*/
	Expand []string
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
	  Value by which to sort

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

// WithTimeout adds the timeout to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) WithExpand(expand []string) *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithName adds the name to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) WithName(name *string) *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) WithPageNumber(pageNumber *int32) *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) WithPageSize(pageSize *int32) *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) WithSortBy(sortBy *string) *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) WithSortOrder(sortOrder *string) *GetTelephonyProvidersEdgesPhonebasesettingsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get telephony providers edges phonebasesettings params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesPhonebasesettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
