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

// NewGetArchitectPromptsParams creates a new GetArchitectPromptsParams object
// with the default values initialized.
func NewGetArchitectPromptsParams() *GetArchitectPromptsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("id")
		sortOrderDefault  = string("asc")
	)
	return &GetArchitectPromptsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectPromptsParamsWithTimeout creates a new GetArchitectPromptsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArchitectPromptsParamsWithTimeout(timeout time.Duration) *GetArchitectPromptsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("id")
		sortOrderDefault  = string("asc")
	)
	return &GetArchitectPromptsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetArchitectPromptsParamsWithContext creates a new GetArchitectPromptsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArchitectPromptsParamsWithContext(ctx context.Context) *GetArchitectPromptsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("id")
		sortOrderDefault  = string("asc")
	)
	return &GetArchitectPromptsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetArchitectPromptsParamsWithHTTPClient creates a new GetArchitectPromptsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArchitectPromptsParamsWithHTTPClient(client *http.Client) *GetArchitectPromptsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("id")
		sortOrderDefault  = string("asc")
	)
	return &GetArchitectPromptsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetArchitectPromptsParams contains all the parameters to send to the API endpoint
for the get architect prompts operation typically these are written to a http.Request
*/
type GetArchitectPromptsParams struct {

	/*Description
	  Description

	*/
	Description *string
	/*Name
	  Name

	*/
	Name []string
	/*NameOrDescription
	  Name or description

	*/
	NameOrDescription *string
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

// WithTimeout adds the timeout to the get architect prompts params
func (o *GetArchitectPromptsParams) WithTimeout(timeout time.Duration) *GetArchitectPromptsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect prompts params
func (o *GetArchitectPromptsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect prompts params
func (o *GetArchitectPromptsParams) WithContext(ctx context.Context) *GetArchitectPromptsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect prompts params
func (o *GetArchitectPromptsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect prompts params
func (o *GetArchitectPromptsParams) WithHTTPClient(client *http.Client) *GetArchitectPromptsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect prompts params
func (o *GetArchitectPromptsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the get architect prompts params
func (o *GetArchitectPromptsParams) WithDescription(description *string) *GetArchitectPromptsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the get architect prompts params
func (o *GetArchitectPromptsParams) SetDescription(description *string) {
	o.Description = description
}

// WithName adds the name to the get architect prompts params
func (o *GetArchitectPromptsParams) WithName(name []string) *GetArchitectPromptsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get architect prompts params
func (o *GetArchitectPromptsParams) SetName(name []string) {
	o.Name = name
}

// WithNameOrDescription adds the nameOrDescription to the get architect prompts params
func (o *GetArchitectPromptsParams) WithNameOrDescription(nameOrDescription *string) *GetArchitectPromptsParams {
	o.SetNameOrDescription(nameOrDescription)
	return o
}

// SetNameOrDescription adds the nameOrDescription to the get architect prompts params
func (o *GetArchitectPromptsParams) SetNameOrDescription(nameOrDescription *string) {
	o.NameOrDescription = nameOrDescription
}

// WithPageNumber adds the pageNumber to the get architect prompts params
func (o *GetArchitectPromptsParams) WithPageNumber(pageNumber *int32) *GetArchitectPromptsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get architect prompts params
func (o *GetArchitectPromptsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get architect prompts params
func (o *GetArchitectPromptsParams) WithPageSize(pageSize *int32) *GetArchitectPromptsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get architect prompts params
func (o *GetArchitectPromptsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get architect prompts params
func (o *GetArchitectPromptsParams) WithSortBy(sortBy *string) *GetArchitectPromptsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get architect prompts params
func (o *GetArchitectPromptsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get architect prompts params
func (o *GetArchitectPromptsParams) WithSortOrder(sortOrder *string) *GetArchitectPromptsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get architect prompts params
func (o *GetArchitectPromptsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectPromptsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// query param description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}

	}

	valuesName := o.Name

	joinedName := swag.JoinByFormat(valuesName, "multi")
	// query array param name
	if err := r.SetQueryParam("name", joinedName...); err != nil {
		return err
	}

	if o.NameOrDescription != nil {

		// query param nameOrDescription
		var qrNameOrDescription string
		if o.NameOrDescription != nil {
			qrNameOrDescription = *o.NameOrDescription
		}
		qNameOrDescription := qrNameOrDescription
		if qNameOrDescription != "" {
			if err := r.SetQueryParam("nameOrDescription", qNameOrDescription); err != nil {
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