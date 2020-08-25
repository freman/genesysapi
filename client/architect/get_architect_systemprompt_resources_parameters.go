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

// NewGetArchitectSystempromptResourcesParams creates a new GetArchitectSystempromptResourcesParams object
// with the default values initialized.
func NewGetArchitectSystempromptResourcesParams() *GetArchitectSystempromptResourcesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("id")
		sortOrderDefault  = string("asc")
	)
	return &GetArchitectSystempromptResourcesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectSystempromptResourcesParamsWithTimeout creates a new GetArchitectSystempromptResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArchitectSystempromptResourcesParamsWithTimeout(timeout time.Duration) *GetArchitectSystempromptResourcesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("id")
		sortOrderDefault  = string("asc")
	)
	return &GetArchitectSystempromptResourcesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetArchitectSystempromptResourcesParamsWithContext creates a new GetArchitectSystempromptResourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArchitectSystempromptResourcesParamsWithContext(ctx context.Context) *GetArchitectSystempromptResourcesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("id")
		sortOrderDefault  = string("asc")
	)
	return &GetArchitectSystempromptResourcesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetArchitectSystempromptResourcesParamsWithHTTPClient creates a new GetArchitectSystempromptResourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArchitectSystempromptResourcesParamsWithHTTPClient(client *http.Client) *GetArchitectSystempromptResourcesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("id")
		sortOrderDefault  = string("asc")
	)
	return &GetArchitectSystempromptResourcesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetArchitectSystempromptResourcesParams contains all the parameters to send to the API endpoint
for the get architect systemprompt resources operation typically these are written to a http.Request
*/
type GetArchitectSystempromptResourcesParams struct {

	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*PromptID
	  Prompt ID

	*/
	PromptID string
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

// WithTimeout adds the timeout to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) WithTimeout(timeout time.Duration) *GetArchitectSystempromptResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) WithContext(ctx context.Context) *GetArchitectSystempromptResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) WithHTTPClient(client *http.Client) *GetArchitectSystempromptResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) WithPageNumber(pageNumber *int32) *GetArchitectSystempromptResourcesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) WithPageSize(pageSize *int32) *GetArchitectSystempromptResourcesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPromptID adds the promptID to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) WithPromptID(promptID string) *GetArchitectSystempromptResourcesParams {
	o.SetPromptID(promptID)
	return o
}

// SetPromptID adds the promptId to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) SetPromptID(promptID string) {
	o.PromptID = promptID
}

// WithSortBy adds the sortBy to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) WithSortBy(sortBy *string) *GetArchitectSystempromptResourcesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) WithSortOrder(sortOrder *string) *GetArchitectSystempromptResourcesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get architect systemprompt resources params
func (o *GetArchitectSystempromptResourcesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectSystempromptResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param promptId
	if err := r.SetPathParam("promptId", o.PromptID); err != nil {
		return err
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