// Code generated by go-swagger; DO NOT EDIT.

package learning

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

// NewGetLearningAssignmentsMeParams creates a new GetLearningAssignmentsMeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLearningAssignmentsMeParams() *GetLearningAssignmentsMeParams {
	return &GetLearningAssignmentsMeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearningAssignmentsMeParamsWithTimeout creates a new GetLearningAssignmentsMeParams object
// with the ability to set a timeout on a request.
func NewGetLearningAssignmentsMeParamsWithTimeout(timeout time.Duration) *GetLearningAssignmentsMeParams {
	return &GetLearningAssignmentsMeParams{
		timeout: timeout,
	}
}

// NewGetLearningAssignmentsMeParamsWithContext creates a new GetLearningAssignmentsMeParams object
// with the ability to set a context for a request.
func NewGetLearningAssignmentsMeParamsWithContext(ctx context.Context) *GetLearningAssignmentsMeParams {
	return &GetLearningAssignmentsMeParams{
		Context: ctx,
	}
}

// NewGetLearningAssignmentsMeParamsWithHTTPClient creates a new GetLearningAssignmentsMeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLearningAssignmentsMeParamsWithHTTPClient(client *http.Client) *GetLearningAssignmentsMeParams {
	return &GetLearningAssignmentsMeParams{
		HTTPClient: client,
	}
}

/*
GetLearningAssignmentsMeParams contains all the parameters to send to the API endpoint

	for the get learning assignments me operation.

	Typically these are written to a http.Request.
*/
type GetLearningAssignmentsMeParams struct {

	/* CompletionInterval.

	   Specifies the range of completion dates to be used for filtering. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss

	   Format: interval
	*/
	CompletionInterval *string

	/* Expand.

	   Specifies the expand option for returning additional information
	*/
	Expand []string

	/* Interval.

	   Specifies the range of dueDates to be queried. Milliseconds will be truncated. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss

	   Format: interval
	*/
	Interval *string

	/* MaxPercentageScore.

	   The maximum assessment score for an assignment (completed with assessment) to be included in the results (inclusive)

	   Format: float
	*/
	MaxPercentageScore *float32

	/* MinPercentageScore.

	   The minimum assessment score for an assignment (completed with assessment) to be included in the results (inclusive)

	   Format: float
	*/
	MinPercentageScore *float32

	/* ModuleID.

	   Specifies the ID of the learning module. Fetch assignments for learning module ID
	*/
	ModuleID *string

	/* Overdue.

	   Specifies if only the non-overdue (overdue is "False") or overdue (overdue is "True") assignments are returned. If overdue is "Any" or if the overdue parameter is not supplied, all assignments are returned

	   Default: "Any"
	*/
	Overdue *string

	/* PageNumber.

	   Page number

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   Page size

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	/* Pass.

	   Specifies if only the failed (pass is "False") or passed (pass is "True") assignments (completed with assessment)are returned. If pass is "Any" or if the pass parameter is not supplied, all assignments are returned

	   Default: "Any"
	*/
	Pass *string

	/* SortBy.

	   Specifies which field to sort the results by, default sort is by recommendedCompletionDate
	*/
	SortBy *string

	/* SortOrder.

	   Specifies result set sort order; if not specified, default sort order is descending (Desc)

	   Default: "Desc"
	*/
	SortOrder *string

	/* States.

	   Specifies the assignment states to filter by
	*/
	States []string

	/* Types.

	   Specifies the module types to filter by
	*/
	Types []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get learning assignments me params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLearningAssignmentsMeParams) WithDefaults() *GetLearningAssignmentsMeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get learning assignments me params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLearningAssignmentsMeParams) SetDefaults() {
	var (
		overdueDefault = string("Any")

		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		passDefault = string("Any")

		sortOrderDefault = string("Desc")
	)

	val := GetLearningAssignmentsMeParams{
		Overdue:    &overdueDefault,
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		Pass:       &passDefault,
		SortOrder:  &sortOrderDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithTimeout(timeout time.Duration) *GetLearningAssignmentsMeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithContext(ctx context.Context) *GetLearningAssignmentsMeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithHTTPClient(client *http.Client) *GetLearningAssignmentsMeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompletionInterval adds the completionInterval to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithCompletionInterval(completionInterval *string) *GetLearningAssignmentsMeParams {
	o.SetCompletionInterval(completionInterval)
	return o
}

// SetCompletionInterval adds the completionInterval to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetCompletionInterval(completionInterval *string) {
	o.CompletionInterval = completionInterval
}

// WithExpand adds the expand to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithExpand(expand []string) *GetLearningAssignmentsMeParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithInterval adds the interval to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithInterval(interval *string) *GetLearningAssignmentsMeParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetInterval(interval *string) {
	o.Interval = interval
}

// WithMaxPercentageScore adds the maxPercentageScore to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithMaxPercentageScore(maxPercentageScore *float32) *GetLearningAssignmentsMeParams {
	o.SetMaxPercentageScore(maxPercentageScore)
	return o
}

// SetMaxPercentageScore adds the maxPercentageScore to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetMaxPercentageScore(maxPercentageScore *float32) {
	o.MaxPercentageScore = maxPercentageScore
}

// WithMinPercentageScore adds the minPercentageScore to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithMinPercentageScore(minPercentageScore *float32) *GetLearningAssignmentsMeParams {
	o.SetMinPercentageScore(minPercentageScore)
	return o
}

// SetMinPercentageScore adds the minPercentageScore to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetMinPercentageScore(minPercentageScore *float32) {
	o.MinPercentageScore = minPercentageScore
}

// WithModuleID adds the moduleID to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithModuleID(moduleID *string) *GetLearningAssignmentsMeParams {
	o.SetModuleID(moduleID)
	return o
}

// SetModuleID adds the moduleId to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetModuleID(moduleID *string) {
	o.ModuleID = moduleID
}

// WithOverdue adds the overdue to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithOverdue(overdue *string) *GetLearningAssignmentsMeParams {
	o.SetOverdue(overdue)
	return o
}

// SetOverdue adds the overdue to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetOverdue(overdue *string) {
	o.Overdue = overdue
}

// WithPageNumber adds the pageNumber to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithPageNumber(pageNumber *int32) *GetLearningAssignmentsMeParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithPageSize(pageSize *int32) *GetLearningAssignmentsMeParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPass adds the pass to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithPass(pass *string) *GetLearningAssignmentsMeParams {
	o.SetPass(pass)
	return o
}

// SetPass adds the pass to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetPass(pass *string) {
	o.Pass = pass
}

// WithSortBy adds the sortBy to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithSortBy(sortBy *string) *GetLearningAssignmentsMeParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithSortOrder(sortOrder *string) *GetLearningAssignmentsMeParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithStates adds the states to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithStates(states []string) *GetLearningAssignmentsMeParams {
	o.SetStates(states)
	return o
}

// SetStates adds the states to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetStates(states []string) {
	o.States = states
}

// WithTypes adds the types to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) WithTypes(types []string) *GetLearningAssignmentsMeParams {
	o.SetTypes(types)
	return o
}

// SetTypes adds the types to the get learning assignments me params
func (o *GetLearningAssignmentsMeParams) SetTypes(types []string) {
	o.Types = types
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearningAssignmentsMeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CompletionInterval != nil {

		// query param completionInterval
		var qrCompletionInterval string

		if o.CompletionInterval != nil {
			qrCompletionInterval = *o.CompletionInterval
		}
		qCompletionInterval := qrCompletionInterval
		if qCompletionInterval != "" {

			if err := r.SetQueryParam("completionInterval", qCompletionInterval); err != nil {
				return err
			}
		}
	}

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	if o.Interval != nil {

		// query param interval
		var qrInterval string

		if o.Interval != nil {
			qrInterval = *o.Interval
		}
		qInterval := qrInterval
		if qInterval != "" {

			if err := r.SetQueryParam("interval", qInterval); err != nil {
				return err
			}
		}
	}

	if o.MaxPercentageScore != nil {

		// query param maxPercentageScore
		var qrMaxPercentageScore float32

		if o.MaxPercentageScore != nil {
			qrMaxPercentageScore = *o.MaxPercentageScore
		}
		qMaxPercentageScore := swag.FormatFloat32(qrMaxPercentageScore)
		if qMaxPercentageScore != "" {

			if err := r.SetQueryParam("maxPercentageScore", qMaxPercentageScore); err != nil {
				return err
			}
		}
	}

	if o.MinPercentageScore != nil {

		// query param minPercentageScore
		var qrMinPercentageScore float32

		if o.MinPercentageScore != nil {
			qrMinPercentageScore = *o.MinPercentageScore
		}
		qMinPercentageScore := swag.FormatFloat32(qrMinPercentageScore)
		if qMinPercentageScore != "" {

			if err := r.SetQueryParam("minPercentageScore", qMinPercentageScore); err != nil {
				return err
			}
		}
	}

	if o.ModuleID != nil {

		// query param moduleId
		var qrModuleID string

		if o.ModuleID != nil {
			qrModuleID = *o.ModuleID
		}
		qModuleID := qrModuleID
		if qModuleID != "" {

			if err := r.SetQueryParam("moduleId", qModuleID); err != nil {
				return err
			}
		}
	}

	if o.Overdue != nil {

		// query param overdue
		var qrOverdue string

		if o.Overdue != nil {
			qrOverdue = *o.Overdue
		}
		qOverdue := qrOverdue
		if qOverdue != "" {

			if err := r.SetQueryParam("overdue", qOverdue); err != nil {
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

	if o.Pass != nil {

		// query param pass
		var qrPass string

		if o.Pass != nil {
			qrPass = *o.Pass
		}
		qPass := qrPass
		if qPass != "" {

			if err := r.SetQueryParam("pass", qPass); err != nil {
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

	if o.States != nil {

		// binding items for states
		joinedStates := o.bindParamStates(reg)

		// query array param states
		if err := r.SetQueryParam("states", joinedStates...); err != nil {
			return err
		}
	}

	if o.Types != nil {

		// binding items for types
		joinedTypes := o.bindParamTypes(reg)

		// query array param types
		if err := r.SetQueryParam("types", joinedTypes...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetLearningAssignmentsMe binds the parameter expand
func (o *GetLearningAssignmentsMeParams) bindParamExpand(formats strfmt.Registry) []string {
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

// bindParamGetLearningAssignmentsMe binds the parameter states
func (o *GetLearningAssignmentsMeParams) bindParamStates(formats strfmt.Registry) []string {
	statesIR := o.States

	var statesIC []string
	for _, statesIIR := range statesIR { // explode []string

		statesIIV := statesIIR // string as string
		statesIC = append(statesIC, statesIIV)
	}

	// items.CollectionFormat: "multi"
	statesIS := swag.JoinByFormat(statesIC, "multi")

	return statesIS
}

// bindParamGetLearningAssignmentsMe binds the parameter types
func (o *GetLearningAssignmentsMeParams) bindParamTypes(formats strfmt.Registry) []string {
	typesIR := o.Types

	var typesIC []string
	for _, typesIIR := range typesIR { // explode []string

		typesIIV := typesIIR // string as string
		typesIC = append(typesIC, typesIIV)
	}

	// items.CollectionFormat: "multi"
	typesIS := swag.JoinByFormat(typesIC, "multi")

	return typesIS
}
