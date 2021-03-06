// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetProfilesUsersParams creates a new GetProfilesUsersParams object
// with the default values initialized.
func NewGetProfilesUsersParams() *GetProfilesUsersParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ASC")
	)
	return &GetProfilesUsersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProfilesUsersParamsWithTimeout creates a new GetProfilesUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProfilesUsersParamsWithTimeout(timeout time.Duration) *GetProfilesUsersParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ASC")
	)
	return &GetProfilesUsersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetProfilesUsersParamsWithContext creates a new GetProfilesUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProfilesUsersParamsWithContext(ctx context.Context) *GetProfilesUsersParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ASC")
	)
	return &GetProfilesUsersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetProfilesUsersParamsWithHTTPClient creates a new GetProfilesUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProfilesUsersParamsWithHTTPClient(client *http.Client) *GetProfilesUsersParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ASC")
	)
	return &GetProfilesUsersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetProfilesUsersParams contains all the parameters to send to the API endpoint
for the get profiles users operation typically these are written to a http.Request
*/
type GetProfilesUsersParams struct {

	/*Expand
	  Which fields, if any, to expand

	*/
	Expand []string
	/*ID
	  id

	*/
	ID []string
	/*IntegrationPresenceSource
	  Gets an integration presence for users instead of their defaults. This parameter will only be used when presence is provided as an "expand".

	*/
	IntegrationPresenceSource *string
	/*Jid
	  jid

	*/
	Jid []string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*SortOrder
	  Ascending or descending sort order

	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get profiles users params
func (o *GetProfilesUsersParams) WithTimeout(timeout time.Duration) *GetProfilesUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get profiles users params
func (o *GetProfilesUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get profiles users params
func (o *GetProfilesUsersParams) WithContext(ctx context.Context) *GetProfilesUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get profiles users params
func (o *GetProfilesUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get profiles users params
func (o *GetProfilesUsersParams) WithHTTPClient(client *http.Client) *GetProfilesUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get profiles users params
func (o *GetProfilesUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get profiles users params
func (o *GetProfilesUsersParams) WithExpand(expand []string) *GetProfilesUsersParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get profiles users params
func (o *GetProfilesUsersParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithID adds the id to the get profiles users params
func (o *GetProfilesUsersParams) WithID(id []string) *GetProfilesUsersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get profiles users params
func (o *GetProfilesUsersParams) SetID(id []string) {
	o.ID = id
}

// WithIntegrationPresenceSource adds the integrationPresenceSource to the get profiles users params
func (o *GetProfilesUsersParams) WithIntegrationPresenceSource(integrationPresenceSource *string) *GetProfilesUsersParams {
	o.SetIntegrationPresenceSource(integrationPresenceSource)
	return o
}

// SetIntegrationPresenceSource adds the integrationPresenceSource to the get profiles users params
func (o *GetProfilesUsersParams) SetIntegrationPresenceSource(integrationPresenceSource *string) {
	o.IntegrationPresenceSource = integrationPresenceSource
}

// WithJid adds the jid to the get profiles users params
func (o *GetProfilesUsersParams) WithJid(jid []string) *GetProfilesUsersParams {
	o.SetJid(jid)
	return o
}

// SetJid adds the jid to the get profiles users params
func (o *GetProfilesUsersParams) SetJid(jid []string) {
	o.Jid = jid
}

// WithPageNumber adds the pageNumber to the get profiles users params
func (o *GetProfilesUsersParams) WithPageNumber(pageNumber *int32) *GetProfilesUsersParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get profiles users params
func (o *GetProfilesUsersParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get profiles users params
func (o *GetProfilesUsersParams) WithPageSize(pageSize *int32) *GetProfilesUsersParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get profiles users params
func (o *GetProfilesUsersParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortOrder adds the sortOrder to the get profiles users params
func (o *GetProfilesUsersParams) WithSortOrder(sortOrder *string) *GetProfilesUsersParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get profiles users params
func (o *GetProfilesUsersParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetProfilesUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
		return err
	}

	if o.IntegrationPresenceSource != nil {

		// query param integrationPresenceSource
		var qrIntegrationPresenceSource string
		if o.IntegrationPresenceSource != nil {
			qrIntegrationPresenceSource = *o.IntegrationPresenceSource
		}
		qIntegrationPresenceSource := qrIntegrationPresenceSource
		if qIntegrationPresenceSource != "" {
			if err := r.SetQueryParam("integrationPresenceSource", qIntegrationPresenceSource); err != nil {
				return err
			}
		}

	}

	valuesJid := o.Jid

	joinedJid := swag.JoinByFormat(valuesJid, "multi")
	// query array param jid
	if err := r.SetQueryParam("jid", joinedJid...); err != nil {
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
