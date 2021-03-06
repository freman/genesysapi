// Code generated by go-swagger; DO NOT EDIT.

package recording

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

// NewGetRecordingCrossplatformMediaretentionpoliciesParams creates a new GetRecordingCrossplatformMediaretentionpoliciesParams object
// with the default values initialized.
func NewGetRecordingCrossplatformMediaretentionpoliciesParams() *GetRecordingCrossplatformMediaretentionpoliciesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		summaryDefault    = bool(false)
	)
	return &GetRecordingCrossplatformMediaretentionpoliciesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		Summary:    &summaryDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecordingCrossplatformMediaretentionpoliciesParamsWithTimeout creates a new GetRecordingCrossplatformMediaretentionpoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecordingCrossplatformMediaretentionpoliciesParamsWithTimeout(timeout time.Duration) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		summaryDefault    = bool(false)
	)
	return &GetRecordingCrossplatformMediaretentionpoliciesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		Summary:    &summaryDefault,

		timeout: timeout,
	}
}

// NewGetRecordingCrossplatformMediaretentionpoliciesParamsWithContext creates a new GetRecordingCrossplatformMediaretentionpoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecordingCrossplatformMediaretentionpoliciesParamsWithContext(ctx context.Context) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		summaryDefault    = bool(false)
	)
	return &GetRecordingCrossplatformMediaretentionpoliciesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		Summary:    &summaryDefault,

		Context: ctx,
	}
}

// NewGetRecordingCrossplatformMediaretentionpoliciesParamsWithHTTPClient creates a new GetRecordingCrossplatformMediaretentionpoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecordingCrossplatformMediaretentionpoliciesParamsWithHTTPClient(client *http.Client) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		summaryDefault    = bool(false)
	)
	return &GetRecordingCrossplatformMediaretentionpoliciesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		Summary:    &summaryDefault,
		HTTPClient: client,
	}
}

/*GetRecordingCrossplatformMediaretentionpoliciesParams contains all the parameters to send to the API endpoint
for the get recording crossplatform mediaretentionpolicies operation typically these are written to a http.Request
*/
type GetRecordingCrossplatformMediaretentionpoliciesParams struct {

	/*Enabled
	  checks to see if policy is enabled - use enabled = true or enabled = false

	*/
	Enabled *bool
	/*Expand
	  variable name requested by expand list

	*/
	Expand []string
	/*HasErrors
	  provides a way to fetch all policies with errors or policies that do not have errors

	*/
	HasErrors *bool
	/*Name
	  the policy name - used for filtering results in searches.

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
	/*SortBy
	  variable name requested to sort by

	*/
	SortBy *string
	/*Summary
	  provides a less verbose response of policy lists.

	*/
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithTimeout(timeout time.Duration) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithContext(ctx context.Context) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithHTTPClient(client *http.Client) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithEnabled(enabled *bool) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithExpand adds the expand to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithExpand(expand []string) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithHasErrors adds the hasErrors to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithHasErrors(hasErrors *bool) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetHasErrors(hasErrors)
	return o
}

// SetHasErrors adds the hasErrors to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetHasErrors(hasErrors *bool) {
	o.HasErrors = hasErrors
}

// WithName adds the name to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithName(name *string) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetName(name *string) {
	o.Name = name
}

// WithNextPage adds the nextPage to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithNextPage(nextPage *string) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithPageNumber adds the pageNumber to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithPageNumber(pageNumber *int32) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithPageSize(pageSize *int32) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPreviousPage adds the previousPage to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithPreviousPage(previousPage *string) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetPreviousPage(previousPage)
	return o
}

// SetPreviousPage adds the previousPage to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetPreviousPage(previousPage *string) {
	o.PreviousPage = previousPage
}

// WithSortBy adds the sortBy to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithSortBy(sortBy *string) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSummary adds the summary to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WithSummary(summary *bool) *GetRecordingCrossplatformMediaretentionpoliciesParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the get recording crossplatform mediaretentionpolicies params
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecordingCrossplatformMediaretentionpoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled bool
		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {
			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}

	}

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	if o.HasErrors != nil {

		// query param hasErrors
		var qrHasErrors bool
		if o.HasErrors != nil {
			qrHasErrors = *o.HasErrors
		}
		qHasErrors := swag.FormatBool(qrHasErrors)
		if qHasErrors != "" {
			if err := r.SetQueryParam("hasErrors", qHasErrors); err != nil {
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

	if o.Summary != nil {

		// query param summary
		var qrSummary bool
		if o.Summary != nil {
			qrSummary = *o.Summary
		}
		qSummary := swag.FormatBool(qrSummary)
		if qSummary != "" {
			if err := r.SetQueryParam("summary", qSummary); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
