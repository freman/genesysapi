// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

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

// NewGetExternalcontactsOrganizationContactsParams creates a new GetExternalcontactsOrganizationContactsParams object
// with the default values initialized.
func NewGetExternalcontactsOrganizationContactsParams() *GetExternalcontactsOrganizationContactsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(20)
	)
	return &GetExternalcontactsOrganizationContactsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalcontactsOrganizationContactsParamsWithTimeout creates a new GetExternalcontactsOrganizationContactsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetExternalcontactsOrganizationContactsParamsWithTimeout(timeout time.Duration) *GetExternalcontactsOrganizationContactsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(20)
	)
	return &GetExternalcontactsOrganizationContactsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetExternalcontactsOrganizationContactsParamsWithContext creates a new GetExternalcontactsOrganizationContactsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetExternalcontactsOrganizationContactsParamsWithContext(ctx context.Context) *GetExternalcontactsOrganizationContactsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(20)
	)
	return &GetExternalcontactsOrganizationContactsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetExternalcontactsOrganizationContactsParamsWithHTTPClient creates a new GetExternalcontactsOrganizationContactsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetExternalcontactsOrganizationContactsParamsWithHTTPClient(client *http.Client) *GetExternalcontactsOrganizationContactsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(20)
	)
	return &GetExternalcontactsOrganizationContactsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetExternalcontactsOrganizationContactsParams contains all the parameters to send to the API endpoint
for the get externalcontacts organization contacts operation typically these are written to a http.Request
*/
type GetExternalcontactsOrganizationContactsParams struct {

	/*Expand
	  which fields, if any, to expand

	*/
	Expand []string
	/*ExternalOrganizationID
	  External Organization ID

	*/
	ExternalOrganizationID string
	/*PageNumber
	  Page number (limited to fetching first 1,000 records; pageNumber * pageSize must be <= 1,000)

	*/
	PageNumber *int32
	/*PageSize
	  Page size (limited to fetching first 1,000 records; pageNumber * pageSize must be <= 1,000)

	*/
	PageSize *int32
	/*Q
	  User supplied search keywords (no special syntax is currently supported)

	*/
	Q *string
	/*SortOrder
	  Sort order

	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) WithTimeout(timeout time.Duration) *GetExternalcontactsOrganizationContactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) WithContext(ctx context.Context) *GetExternalcontactsOrganizationContactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) WithHTTPClient(client *http.Client) *GetExternalcontactsOrganizationContactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) WithExpand(expand []string) *GetExternalcontactsOrganizationContactsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithExternalOrganizationID adds the externalOrganizationID to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) WithExternalOrganizationID(externalOrganizationID string) *GetExternalcontactsOrganizationContactsParams {
	o.SetExternalOrganizationID(externalOrganizationID)
	return o
}

// SetExternalOrganizationID adds the externalOrganizationId to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) SetExternalOrganizationID(externalOrganizationID string) {
	o.ExternalOrganizationID = externalOrganizationID
}

// WithPageNumber adds the pageNumber to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) WithPageNumber(pageNumber *int32) *GetExternalcontactsOrganizationContactsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) WithPageSize(pageSize *int32) *GetExternalcontactsOrganizationContactsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithQ adds the q to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) WithQ(q *string) *GetExternalcontactsOrganizationContactsParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) SetQ(q *string) {
	o.Q = q
}

// WithSortOrder adds the sortOrder to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) WithSortOrder(sortOrder *string) *GetExternalcontactsOrganizationContactsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get externalcontacts organization contacts params
func (o *GetExternalcontactsOrganizationContactsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalcontactsOrganizationContactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param externalOrganizationId
	if err := r.SetPathParam("externalOrganizationId", o.ExternalOrganizationID); err != nil {
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

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
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
