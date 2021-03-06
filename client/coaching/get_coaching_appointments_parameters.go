// Code generated by go-swagger; DO NOT EDIT.

package coaching

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

// NewGetCoachingAppointmentsParams creates a new GetCoachingAppointmentsParams object
// with the default values initialized.
func NewGetCoachingAppointmentsParams() *GetCoachingAppointmentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetCoachingAppointmentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCoachingAppointmentsParamsWithTimeout creates a new GetCoachingAppointmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCoachingAppointmentsParamsWithTimeout(timeout time.Duration) *GetCoachingAppointmentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetCoachingAppointmentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetCoachingAppointmentsParamsWithContext creates a new GetCoachingAppointmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCoachingAppointmentsParamsWithContext(ctx context.Context) *GetCoachingAppointmentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetCoachingAppointmentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetCoachingAppointmentsParamsWithHTTPClient creates a new GetCoachingAppointmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCoachingAppointmentsParamsWithHTTPClient(client *http.Client) *GetCoachingAppointmentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetCoachingAppointmentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetCoachingAppointmentsParams contains all the parameters to send to the API endpoint
for the get coaching appointments operation typically these are written to a http.Request
*/
type GetCoachingAppointmentsParams struct {

	/*CompletionInterval
	  Appointment completion start and end to filter by. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss

	*/
	CompletionInterval *string
	/*FacilitatorIds
	  The facilitator IDs for which to retrieve appointments

	*/
	FacilitatorIds []string
	/*Interval
	  Interval to filter data by. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss

	*/
	Interval *string
	/*Overdue
	  Overdue status to filter by

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
	/*Relationships
	  Relationships to filter by

	*/
	Relationships []string
	/*SortOrder
	  Sort (by due date) either Asc or Desc

	*/
	SortOrder *string
	/*Statuses
	  Appointment Statuses to filter by

	*/
	Statuses []string
	/*UserIds
	  The user IDs for which to retrieve appointments

	*/
	UserIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithTimeout(timeout time.Duration) *GetCoachingAppointmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithContext(ctx context.Context) *GetCoachingAppointmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithHTTPClient(client *http.Client) *GetCoachingAppointmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompletionInterval adds the completionInterval to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithCompletionInterval(completionInterval *string) *GetCoachingAppointmentsParams {
	o.SetCompletionInterval(completionInterval)
	return o
}

// SetCompletionInterval adds the completionInterval to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetCompletionInterval(completionInterval *string) {
	o.CompletionInterval = completionInterval
}

// WithFacilitatorIds adds the facilitatorIds to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithFacilitatorIds(facilitatorIds []string) *GetCoachingAppointmentsParams {
	o.SetFacilitatorIds(facilitatorIds)
	return o
}

// SetFacilitatorIds adds the facilitatorIds to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetFacilitatorIds(facilitatorIds []string) {
	o.FacilitatorIds = facilitatorIds
}

// WithInterval adds the interval to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithInterval(interval *string) *GetCoachingAppointmentsParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetInterval(interval *string) {
	o.Interval = interval
}

// WithOverdue adds the overdue to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithOverdue(overdue *string) *GetCoachingAppointmentsParams {
	o.SetOverdue(overdue)
	return o
}

// SetOverdue adds the overdue to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetOverdue(overdue *string) {
	o.Overdue = overdue
}

// WithPageNumber adds the pageNumber to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithPageNumber(pageNumber *int32) *GetCoachingAppointmentsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithPageSize(pageSize *int32) *GetCoachingAppointmentsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithRelationships adds the relationships to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithRelationships(relationships []string) *GetCoachingAppointmentsParams {
	o.SetRelationships(relationships)
	return o
}

// SetRelationships adds the relationships to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetRelationships(relationships []string) {
	o.Relationships = relationships
}

// WithSortOrder adds the sortOrder to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithSortOrder(sortOrder *string) *GetCoachingAppointmentsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithStatuses adds the statuses to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithStatuses(statuses []string) *GetCoachingAppointmentsParams {
	o.SetStatuses(statuses)
	return o
}

// SetStatuses adds the statuses to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetStatuses(statuses []string) {
	o.Statuses = statuses
}

// WithUserIds adds the userIds to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) WithUserIds(userIds []string) *GetCoachingAppointmentsParams {
	o.SetUserIds(userIds)
	return o
}

// SetUserIds adds the userIds to the get coaching appointments params
func (o *GetCoachingAppointmentsParams) SetUserIds(userIds []string) {
	o.UserIds = userIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetCoachingAppointmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesFacilitatorIds := o.FacilitatorIds

	joinedFacilitatorIds := swag.JoinByFormat(valuesFacilitatorIds, "multi")
	// query array param facilitatorIds
	if err := r.SetQueryParam("facilitatorIds", joinedFacilitatorIds...); err != nil {
		return err
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

	valuesRelationships := o.Relationships

	joinedRelationships := swag.JoinByFormat(valuesRelationships, "multi")
	// query array param relationships
	if err := r.SetQueryParam("relationships", joinedRelationships...); err != nil {
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

	valuesUserIds := o.UserIds

	joinedUserIds := swag.JoinByFormat(valuesUserIds, "multi")
	// query array param userIds
	if err := r.SetQueryParam("userIds", joinedUserIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
