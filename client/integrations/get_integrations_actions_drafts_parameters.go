// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetIntegrationsActionsDraftsParams creates a new GetIntegrationsActionsDraftsParams object
// with the default values initialized.
func NewGetIntegrationsActionsDraftsParams() *GetIntegrationsActionsDraftsParams {
	var (
		includeAuthActionsDefault = string("false")
		pageNumberDefault         = int32(1)
		pageSizeDefault           = int32(25)
		sortOrderDefault          = string("asc")
	)
	return &GetIntegrationsActionsDraftsParams{
		IncludeAuthActions: &includeAuthActionsDefault,
		PageNumber:         &pageNumberDefault,
		PageSize:           &pageSizeDefault,
		SortOrder:          &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationsActionsDraftsParamsWithTimeout creates a new GetIntegrationsActionsDraftsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIntegrationsActionsDraftsParamsWithTimeout(timeout time.Duration) *GetIntegrationsActionsDraftsParams {
	var (
		includeAuthActionsDefault = string("false")
		pageNumberDefault         = int32(1)
		pageSizeDefault           = int32(25)
		sortOrderDefault          = string("asc")
	)
	return &GetIntegrationsActionsDraftsParams{
		IncludeAuthActions: &includeAuthActionsDefault,
		PageNumber:         &pageNumberDefault,
		PageSize:           &pageSizeDefault,
		SortOrder:          &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetIntegrationsActionsDraftsParamsWithContext creates a new GetIntegrationsActionsDraftsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIntegrationsActionsDraftsParamsWithContext(ctx context.Context) *GetIntegrationsActionsDraftsParams {
	var (
		includeAuthActionsDefault = string("false")
		pageNumberDefault         = int32(1)
		pageSizeDefault           = int32(25)
		sortOrderDefault          = string("asc")
	)
	return &GetIntegrationsActionsDraftsParams{
		IncludeAuthActions: &includeAuthActionsDefault,
		PageNumber:         &pageNumberDefault,
		PageSize:           &pageSizeDefault,
		SortOrder:          &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetIntegrationsActionsDraftsParamsWithHTTPClient creates a new GetIntegrationsActionsDraftsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIntegrationsActionsDraftsParamsWithHTTPClient(client *http.Client) *GetIntegrationsActionsDraftsParams {
	var (
		includeAuthActionsDefault = string("false")
		pageNumberDefault         = int32(1)
		pageSizeDefault           = int32(25)
		sortOrderDefault          = string("asc")
	)
	return &GetIntegrationsActionsDraftsParams{
		IncludeAuthActions: &includeAuthActionsDefault,
		PageNumber:         &pageNumberDefault,
		PageSize:           &pageSizeDefault,
		SortOrder:          &sortOrderDefault,
		HTTPClient:         client,
	}
}

/*GetIntegrationsActionsDraftsParams contains all the parameters to send to the API endpoint
for the get integrations actions drafts operation typically these are written to a http.Request
*/
type GetIntegrationsActionsDraftsParams struct {

	/*Category
	  Filter by category name

	*/
	Category *string
	/*IncludeAuthActions
	  Whether or not to include authentication actions in the response. These actions are not directly executable. Some integrations create them and will run them as needed to refresh authentication information for other actions.

	*/
	IncludeAuthActions *string
	/*Name
	  Filter by action name. Provide full or just the first part of name.

	*/
	Name *string
	/*NextPage
	  next page token

	*/
	NextPage *string
	/*PageNumber
	  The page number requested

	*/
	PageNumber *int32
	/*PageSize
	  The total page size requested

	*/
	PageSize *int32
	/*PreviousPage
	  Previous page token

	*/
	PreviousPage *string
	/*Secure
	  Filter to only include secure actions. True will only include actions marked secured. False will include only unsecure actions. Do not use filter if you want all Actions.

	*/
	Secure *string
	/*SortBy
	  Root level field name to sort on.

	*/
	SortBy *string
	/*SortOrder
	  Direction to sort 'sortBy' field.

	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithTimeout(timeout time.Duration) *GetIntegrationsActionsDraftsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithContext(ctx context.Context) *GetIntegrationsActionsDraftsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithHTTPClient(client *http.Client) *GetIntegrationsActionsDraftsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategory adds the category to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithCategory(category *string) *GetIntegrationsActionsDraftsParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetCategory(category *string) {
	o.Category = category
}

// WithIncludeAuthActions adds the includeAuthActions to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithIncludeAuthActions(includeAuthActions *string) *GetIntegrationsActionsDraftsParams {
	o.SetIncludeAuthActions(includeAuthActions)
	return o
}

// SetIncludeAuthActions adds the includeAuthActions to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetIncludeAuthActions(includeAuthActions *string) {
	o.IncludeAuthActions = includeAuthActions
}

// WithName adds the name to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithName(name *string) *GetIntegrationsActionsDraftsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetName(name *string) {
	o.Name = name
}

// WithNextPage adds the nextPage to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithNextPage(nextPage *string) *GetIntegrationsActionsDraftsParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithPageNumber adds the pageNumber to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithPageNumber(pageNumber *int32) *GetIntegrationsActionsDraftsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithPageSize(pageSize *int32) *GetIntegrationsActionsDraftsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPreviousPage adds the previousPage to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithPreviousPage(previousPage *string) *GetIntegrationsActionsDraftsParams {
	o.SetPreviousPage(previousPage)
	return o
}

// SetPreviousPage adds the previousPage to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetPreviousPage(previousPage *string) {
	o.PreviousPage = previousPage
}

// WithSecure adds the secure to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithSecure(secure *string) *GetIntegrationsActionsDraftsParams {
	o.SetSecure(secure)
	return o
}

// SetSecure adds the secure to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetSecure(secure *string) {
	o.Secure = secure
}

// WithSortBy adds the sortBy to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithSortBy(sortBy *string) *GetIntegrationsActionsDraftsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) WithSortOrder(sortOrder *string) *GetIntegrationsActionsDraftsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get integrations actions drafts params
func (o *GetIntegrationsActionsDraftsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationsActionsDraftsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Category != nil {

		// query param category
		var qrCategory string
		if o.Category != nil {
			qrCategory = *o.Category
		}
		qCategory := qrCategory
		if qCategory != "" {
			if err := r.SetQueryParam("category", qCategory); err != nil {
				return err
			}
		}

	}

	if o.IncludeAuthActions != nil {

		// query param includeAuthActions
		var qrIncludeAuthActions string
		if o.IncludeAuthActions != nil {
			qrIncludeAuthActions = *o.IncludeAuthActions
		}
		qIncludeAuthActions := qrIncludeAuthActions
		if qIncludeAuthActions != "" {
			if err := r.SetQueryParam("includeAuthActions", qIncludeAuthActions); err != nil {
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

	if o.Secure != nil {

		// query param secure
		var qrSecure string
		if o.Secure != nil {
			qrSecure = *o.Secure
		}
		qSecure := qrSecure
		if qSecure != "" {
			if err := r.SetQueryParam("secure", qSecure); err != nil {
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
