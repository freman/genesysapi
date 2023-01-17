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

// NewGetExternalcontactsOrganizationNotesParams creates a new GetExternalcontactsOrganizationNotesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExternalcontactsOrganizationNotesParams() *GetExternalcontactsOrganizationNotesParams {
	return &GetExternalcontactsOrganizationNotesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalcontactsOrganizationNotesParamsWithTimeout creates a new GetExternalcontactsOrganizationNotesParams object
// with the ability to set a timeout on a request.
func NewGetExternalcontactsOrganizationNotesParamsWithTimeout(timeout time.Duration) *GetExternalcontactsOrganizationNotesParams {
	return &GetExternalcontactsOrganizationNotesParams{
		timeout: timeout,
	}
}

// NewGetExternalcontactsOrganizationNotesParamsWithContext creates a new GetExternalcontactsOrganizationNotesParams object
// with the ability to set a context for a request.
func NewGetExternalcontactsOrganizationNotesParamsWithContext(ctx context.Context) *GetExternalcontactsOrganizationNotesParams {
	return &GetExternalcontactsOrganizationNotesParams{
		Context: ctx,
	}
}

// NewGetExternalcontactsOrganizationNotesParamsWithHTTPClient creates a new GetExternalcontactsOrganizationNotesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExternalcontactsOrganizationNotesParamsWithHTTPClient(client *http.Client) *GetExternalcontactsOrganizationNotesParams {
	return &GetExternalcontactsOrganizationNotesParams{
		HTTPClient: client,
	}
}

/*
GetExternalcontactsOrganizationNotesParams contains all the parameters to send to the API endpoint

	for the get externalcontacts organization notes operation.

	Typically these are written to a http.Request.
*/
type GetExternalcontactsOrganizationNotesParams struct {

	/* Expand.

	   which fields, if any, to expand
	*/
	Expand []string

	/* ExternalOrganizationID.

	   External Organization Id
	*/
	ExternalOrganizationID string

	/* PageNumber.

	   Page number (limited to fetching first 1,000 records; pageNumber * pageSize must be <= 1,000)

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   Page size (limited to fetching first 1,000 records; pageNumber * pageSize must be <= 1,000)

	   Format: int32
	   Default: 20
	*/
	PageSize *int32

	/* SortOrder.

	   Sort order
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get externalcontacts organization notes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalcontactsOrganizationNotesParams) WithDefaults() *GetExternalcontactsOrganizationNotesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get externalcontacts organization notes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalcontactsOrganizationNotesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(20)
	)

	val := GetExternalcontactsOrganizationNotesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) WithTimeout(timeout time.Duration) *GetExternalcontactsOrganizationNotesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) WithContext(ctx context.Context) *GetExternalcontactsOrganizationNotesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) WithHTTPClient(client *http.Client) *GetExternalcontactsOrganizationNotesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) WithExpand(expand []string) *GetExternalcontactsOrganizationNotesParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithExternalOrganizationID adds the externalOrganizationID to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) WithExternalOrganizationID(externalOrganizationID string) *GetExternalcontactsOrganizationNotesParams {
	o.SetExternalOrganizationID(externalOrganizationID)
	return o
}

// SetExternalOrganizationID adds the externalOrganizationId to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) SetExternalOrganizationID(externalOrganizationID string) {
	o.ExternalOrganizationID = externalOrganizationID
}

// WithPageNumber adds the pageNumber to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) WithPageNumber(pageNumber *int32) *GetExternalcontactsOrganizationNotesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) WithPageSize(pageSize *int32) *GetExternalcontactsOrganizationNotesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortOrder adds the sortOrder to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) WithSortOrder(sortOrder *string) *GetExternalcontactsOrganizationNotesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get externalcontacts organization notes params
func (o *GetExternalcontactsOrganizationNotesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalcontactsOrganizationNotesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetExternalcontactsOrganizationNotes binds the parameter expand
func (o *GetExternalcontactsOrganizationNotesParams) bindParamExpand(formats strfmt.Registry) []string {
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
