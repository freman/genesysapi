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

// NewGetCoachingAppointmentsMeParams creates a new GetCoachingAppointmentsMeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCoachingAppointmentsMeParams() *GetCoachingAppointmentsMeParams {
	return &GetCoachingAppointmentsMeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCoachingAppointmentsMeParamsWithTimeout creates a new GetCoachingAppointmentsMeParams object
// with the ability to set a timeout on a request.
func NewGetCoachingAppointmentsMeParamsWithTimeout(timeout time.Duration) *GetCoachingAppointmentsMeParams {
	return &GetCoachingAppointmentsMeParams{
		timeout: timeout,
	}
}

// NewGetCoachingAppointmentsMeParamsWithContext creates a new GetCoachingAppointmentsMeParams object
// with the ability to set a context for a request.
func NewGetCoachingAppointmentsMeParamsWithContext(ctx context.Context) *GetCoachingAppointmentsMeParams {
	return &GetCoachingAppointmentsMeParams{
		Context: ctx,
	}
}

// NewGetCoachingAppointmentsMeParamsWithHTTPClient creates a new GetCoachingAppointmentsMeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCoachingAppointmentsMeParamsWithHTTPClient(client *http.Client) *GetCoachingAppointmentsMeParams {
	return &GetCoachingAppointmentsMeParams{
		HTTPClient: client,
	}
}

/*
GetCoachingAppointmentsMeParams contains all the parameters to send to the API endpoint

	for the get coaching appointments me operation.

	Typically these are written to a http.Request.
*/
type GetCoachingAppointmentsMeParams struct {

	/* CompletionInterval.

	   Appointment completion start and end to filter by. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss

	   Format: interval
	*/
	CompletionInterval *string

	/* FacilitatorIds.

	   The facilitator IDs for which to retrieve appointments
	*/
	FacilitatorIds []string

	/* Interval.

	   Interval to filter data by. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss

	   Format: interval
	*/
	Interval *string

	/* IntervalCondition.

	   Filter condition for interval
	*/
	IntervalCondition *string

	/* Overdue.

	   Overdue status to filter by
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

	/* Relationships.

	   Relationships to filter by
	*/
	Relationships []string

	/* SortOrder.

	   Sort (by due date) either Asc or Desc
	*/
	SortOrder *string

	/* Statuses.

	   Appointment Statuses to filter by
	*/
	Statuses []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get coaching appointments me params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoachingAppointmentsMeParams) WithDefaults() *GetCoachingAppointmentsMeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get coaching appointments me params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoachingAppointmentsMeParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetCoachingAppointmentsMeParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithTimeout(timeout time.Duration) *GetCoachingAppointmentsMeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithContext(ctx context.Context) *GetCoachingAppointmentsMeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithHTTPClient(client *http.Client) *GetCoachingAppointmentsMeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompletionInterval adds the completionInterval to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithCompletionInterval(completionInterval *string) *GetCoachingAppointmentsMeParams {
	o.SetCompletionInterval(completionInterval)
	return o
}

// SetCompletionInterval adds the completionInterval to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetCompletionInterval(completionInterval *string) {
	o.CompletionInterval = completionInterval
}

// WithFacilitatorIds adds the facilitatorIds to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithFacilitatorIds(facilitatorIds []string) *GetCoachingAppointmentsMeParams {
	o.SetFacilitatorIds(facilitatorIds)
	return o
}

// SetFacilitatorIds adds the facilitatorIds to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetFacilitatorIds(facilitatorIds []string) {
	o.FacilitatorIds = facilitatorIds
}

// WithInterval adds the interval to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithInterval(interval *string) *GetCoachingAppointmentsMeParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetInterval(interval *string) {
	o.Interval = interval
}

// WithIntervalCondition adds the intervalCondition to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithIntervalCondition(intervalCondition *string) *GetCoachingAppointmentsMeParams {
	o.SetIntervalCondition(intervalCondition)
	return o
}

// SetIntervalCondition adds the intervalCondition to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetIntervalCondition(intervalCondition *string) {
	o.IntervalCondition = intervalCondition
}

// WithOverdue adds the overdue to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithOverdue(overdue *string) *GetCoachingAppointmentsMeParams {
	o.SetOverdue(overdue)
	return o
}

// SetOverdue adds the overdue to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetOverdue(overdue *string) {
	o.Overdue = overdue
}

// WithPageNumber adds the pageNumber to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithPageNumber(pageNumber *int32) *GetCoachingAppointmentsMeParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithPageSize(pageSize *int32) *GetCoachingAppointmentsMeParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithRelationships adds the relationships to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithRelationships(relationships []string) *GetCoachingAppointmentsMeParams {
	o.SetRelationships(relationships)
	return o
}

// SetRelationships adds the relationships to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetRelationships(relationships []string) {
	o.Relationships = relationships
}

// WithSortOrder adds the sortOrder to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithSortOrder(sortOrder *string) *GetCoachingAppointmentsMeParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithStatuses adds the statuses to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) WithStatuses(statuses []string) *GetCoachingAppointmentsMeParams {
	o.SetStatuses(statuses)
	return o
}

// SetStatuses adds the statuses to the get coaching appointments me params
func (o *GetCoachingAppointmentsMeParams) SetStatuses(statuses []string) {
	o.Statuses = statuses
}

// WriteToRequest writes these params to a swagger request
func (o *GetCoachingAppointmentsMeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FacilitatorIds != nil {

		// binding items for facilitatorIds
		joinedFacilitatorIds := o.bindParamFacilitatorIds(reg)

		// query array param facilitatorIds
		if err := r.SetQueryParam("facilitatorIds", joinedFacilitatorIds...); err != nil {
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

	if o.IntervalCondition != nil {

		// query param intervalCondition
		var qrIntervalCondition string

		if o.IntervalCondition != nil {
			qrIntervalCondition = *o.IntervalCondition
		}
		qIntervalCondition := qrIntervalCondition
		if qIntervalCondition != "" {

			if err := r.SetQueryParam("intervalCondition", qIntervalCondition); err != nil {
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

	if o.Relationships != nil {

		// binding items for relationships
		joinedRelationships := o.bindParamRelationships(reg)

		// query array param relationships
		if err := r.SetQueryParam("relationships", joinedRelationships...); err != nil {
			return err
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

	if o.Statuses != nil {

		// binding items for statuses
		joinedStatuses := o.bindParamStatuses(reg)

		// query array param statuses
		if err := r.SetQueryParam("statuses", joinedStatuses...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetCoachingAppointmentsMe binds the parameter facilitatorIds
func (o *GetCoachingAppointmentsMeParams) bindParamFacilitatorIds(formats strfmt.Registry) []string {
	facilitatorIdsIR := o.FacilitatorIds

	var facilitatorIdsIC []string
	for _, facilitatorIdsIIR := range facilitatorIdsIR { // explode []string

		facilitatorIdsIIV := facilitatorIdsIIR // string as string
		facilitatorIdsIC = append(facilitatorIdsIC, facilitatorIdsIIV)
	}

	// items.CollectionFormat: "multi"
	facilitatorIdsIS := swag.JoinByFormat(facilitatorIdsIC, "multi")

	return facilitatorIdsIS
}

// bindParamGetCoachingAppointmentsMe binds the parameter relationships
func (o *GetCoachingAppointmentsMeParams) bindParamRelationships(formats strfmt.Registry) []string {
	relationshipsIR := o.Relationships

	var relationshipsIC []string
	for _, relationshipsIIR := range relationshipsIR { // explode []string

		relationshipsIIV := relationshipsIIR // string as string
		relationshipsIC = append(relationshipsIC, relationshipsIIV)
	}

	// items.CollectionFormat: "multi"
	relationshipsIS := swag.JoinByFormat(relationshipsIC, "multi")

	return relationshipsIS
}

// bindParamGetCoachingAppointmentsMe binds the parameter statuses
func (o *GetCoachingAppointmentsMeParams) bindParamStatuses(formats strfmt.Registry) []string {
	statusesIR := o.Statuses

	var statusesIC []string
	for _, statusesIIR := range statusesIR { // explode []string

		statusesIIV := statusesIIR // string as string
		statusesIC = append(statusesIC, statusesIIV)
	}

	// items.CollectionFormat: "multi"
	statusesIS := swag.JoinByFormat(statusesIC, "multi")

	return statusesIS
}
