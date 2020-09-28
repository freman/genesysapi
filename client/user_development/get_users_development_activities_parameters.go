// Code generated by go-swagger; DO NOT EDIT.

package user_development

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

// NewGetUsersDevelopmentActivitiesParams creates a new GetUsersDevelopmentActivitiesParams object
// with the default values initialized.
func NewGetUsersDevelopmentActivitiesParams() *GetUsersDevelopmentActivitiesParams {
	var (
		overdueDefault    = string("Any")
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("Desc")
	)
	return &GetUsersDevelopmentActivitiesParams{
		Overdue:    &overdueDefault,
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersDevelopmentActivitiesParamsWithTimeout creates a new GetUsersDevelopmentActivitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersDevelopmentActivitiesParamsWithTimeout(timeout time.Duration) *GetUsersDevelopmentActivitiesParams {
	var (
		overdueDefault    = string("Any")
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("Desc")
	)
	return &GetUsersDevelopmentActivitiesParams{
		Overdue:    &overdueDefault,
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetUsersDevelopmentActivitiesParamsWithContext creates a new GetUsersDevelopmentActivitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersDevelopmentActivitiesParamsWithContext(ctx context.Context) *GetUsersDevelopmentActivitiesParams {
	var (
		overdueDefault    = string("Any")
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("Desc")
	)
	return &GetUsersDevelopmentActivitiesParams{
		Overdue:    &overdueDefault,
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetUsersDevelopmentActivitiesParamsWithHTTPClient creates a new GetUsersDevelopmentActivitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersDevelopmentActivitiesParamsWithHTTPClient(client *http.Client) *GetUsersDevelopmentActivitiesParams {
	var (
		overdueDefault    = string("Any")
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("Desc")
	)
	return &GetUsersDevelopmentActivitiesParams{
		Overdue:    &overdueDefault,
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetUsersDevelopmentActivitiesParams contains all the parameters to send to the API endpoint
for the get users development activities operation typically these are written to a http.Request
*/
type GetUsersDevelopmentActivitiesParams struct {

	/*CompletionInterval
	  Specifies the range of completion dates to be used for filtering. A maximum of 365 days can be specified in the range. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss

	*/
	CompletionInterval *string
	/*Interval
	  Specifies the dateDue range to be queried. Milliseconds will be truncated. A maximum of 365 days can be specified in the range. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss

	*/
	Interval *string
	/*ModuleID
	  Specifies the ID of the learning module.

	*/
	ModuleID *string
	/*Overdue
	  Specifies if non-overdue, overdue, or all activities are returned. If not specified, all activities are returned

	*/
	Overdue *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*Relationship
	  Specifies how the current user relation should be interpreted, and filters the activities returned to only those that have the specified relationship. If not specified, all relationships are returned.

	*/
	Relationship []string
	/*SortOrder
	  Specifies result set sort order sorted by the date due; if not specified, default sort order is descending (Desc)

	*/
	SortOrder *string
	/*Statuses
	  Specifies the activity statuses to filter by

	*/
	Statuses []string
	/*Types
	  Specifies the activity types.

	*/
	Types []string
	/*UserID
	  Specifies the list of user IDs to be queried, up to 100 user IDs. It searches for any relationship for the userId.

	*/
	UserID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithTimeout(timeout time.Duration) *GetUsersDevelopmentActivitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithContext(ctx context.Context) *GetUsersDevelopmentActivitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithHTTPClient(client *http.Client) *GetUsersDevelopmentActivitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompletionInterval adds the completionInterval to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithCompletionInterval(completionInterval *string) *GetUsersDevelopmentActivitiesParams {
	o.SetCompletionInterval(completionInterval)
	return o
}

// SetCompletionInterval adds the completionInterval to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetCompletionInterval(completionInterval *string) {
	o.CompletionInterval = completionInterval
}

// WithInterval adds the interval to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithInterval(interval *string) *GetUsersDevelopmentActivitiesParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetInterval(interval *string) {
	o.Interval = interval
}

// WithModuleID adds the moduleID to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithModuleID(moduleID *string) *GetUsersDevelopmentActivitiesParams {
	o.SetModuleID(moduleID)
	return o
}

// SetModuleID adds the moduleId to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetModuleID(moduleID *string) {
	o.ModuleID = moduleID
}

// WithOverdue adds the overdue to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithOverdue(overdue *string) *GetUsersDevelopmentActivitiesParams {
	o.SetOverdue(overdue)
	return o
}

// SetOverdue adds the overdue to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetOverdue(overdue *string) {
	o.Overdue = overdue
}

// WithPageNumber adds the pageNumber to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithPageNumber(pageNumber *int32) *GetUsersDevelopmentActivitiesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithPageSize(pageSize *int32) *GetUsersDevelopmentActivitiesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithRelationship adds the relationship to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithRelationship(relationship []string) *GetUsersDevelopmentActivitiesParams {
	o.SetRelationship(relationship)
	return o
}

// SetRelationship adds the relationship to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetRelationship(relationship []string) {
	o.Relationship = relationship
}

// WithSortOrder adds the sortOrder to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithSortOrder(sortOrder *string) *GetUsersDevelopmentActivitiesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithStatuses adds the statuses to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithStatuses(statuses []string) *GetUsersDevelopmentActivitiesParams {
	o.SetStatuses(statuses)
	return o
}

// SetStatuses adds the statuses to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetStatuses(statuses []string) {
	o.Statuses = statuses
}

// WithTypes adds the types to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithTypes(types []string) *GetUsersDevelopmentActivitiesParams {
	o.SetTypes(types)
	return o
}

// SetTypes adds the types to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetTypes(types []string) {
	o.Types = types
}

// WithUserID adds the userID to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) WithUserID(userID []string) *GetUsersDevelopmentActivitiesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get users development activities params
func (o *GetUsersDevelopmentActivitiesParams) SetUserID(userID []string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersDevelopmentActivitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesRelationship := o.Relationship

	joinedRelationship := swag.JoinByFormat(valuesRelationship, "multi")
	// query array param relationship
	if err := r.SetQueryParam("relationship", joinedRelationship...); err != nil {
		return err
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

	valuesStatuses := o.Statuses

	joinedStatuses := swag.JoinByFormat(valuesStatuses, "multi")
	// query array param statuses
	if err := r.SetQueryParam("statuses", joinedStatuses...); err != nil {
		return err
	}

	valuesTypes := o.Types

	joinedTypes := swag.JoinByFormat(valuesTypes, "multi")
	// query array param types
	if err := r.SetQueryParam("types", joinedTypes...); err != nil {
		return err
	}

	valuesUserID := o.UserID

	joinedUserID := swag.JoinByFormat(valuesUserID, "multi")
	// query array param userId
	if err := r.SetQueryParam("userId", joinedUserID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
