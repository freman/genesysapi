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

// NewGetArchitectEmergencygroupsParams creates a new GetArchitectEmergencygroupsParams object
// with the default values initialized.
func NewGetArchitectEmergencygroupsParams() *GetArchitectEmergencygroupsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("ASC")
	)
	return &GetArchitectEmergencygroupsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectEmergencygroupsParamsWithTimeout creates a new GetArchitectEmergencygroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArchitectEmergencygroupsParamsWithTimeout(timeout time.Duration) *GetArchitectEmergencygroupsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("ASC")
	)
	return &GetArchitectEmergencygroupsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetArchitectEmergencygroupsParamsWithContext creates a new GetArchitectEmergencygroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArchitectEmergencygroupsParamsWithContext(ctx context.Context) *GetArchitectEmergencygroupsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("ASC")
	)
	return &GetArchitectEmergencygroupsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetArchitectEmergencygroupsParamsWithHTTPClient creates a new GetArchitectEmergencygroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArchitectEmergencygroupsParamsWithHTTPClient(client *http.Client) *GetArchitectEmergencygroupsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("ASC")
	)
	return &GetArchitectEmergencygroupsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetArchitectEmergencygroupsParams contains all the parameters to send to the API endpoint
for the get architect emergencygroups operation typically these are written to a http.Request
*/
type GetArchitectEmergencygroupsParams struct {

	/*Name
	  Name of the Emergency Group to filter by.

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
	/*SortOrder
	  Sort order

	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) WithTimeout(timeout time.Duration) *GetArchitectEmergencygroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) WithContext(ctx context.Context) *GetArchitectEmergencygroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) WithHTTPClient(client *http.Client) *GetArchitectEmergencygroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) WithName(name *string) *GetArchitectEmergencygroupsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) WithPageNumber(pageNumber *int32) *GetArchitectEmergencygroupsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) WithPageSize(pageSize *int32) *GetArchitectEmergencygroupsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) WithSortBy(sortBy *string) *GetArchitectEmergencygroupsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) WithSortOrder(sortOrder *string) *GetArchitectEmergencygroupsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get architect emergencygroups params
func (o *GetArchitectEmergencygroupsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectEmergencygroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
