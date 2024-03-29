// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

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

	"github.com/freman/genesysapi/models"
)

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateParams creates a new PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateParams() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateParamsWithTimeout creates a new PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateParamsWithContext creates a new PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateParamsWithContext(ctx context.Context) *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateParamsWithHTTPClient creates a new PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementBusinessunitWeekScheduleUpdateParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams {
	return &PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement businessunit week schedule update operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams struct {

	/* Body.

	   body
	*/
	Body *models.ProcessScheduleUpdateUploadRequest

	/* BusinessUnitID.

	   The ID of the business unit
	*/
	BusinessUnitID string

	/* ScheduleID.

	   The ID of the schedule
	*/
	ScheduleID string

	/* WeekID.

	   First day of schedule week in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	   Format: date
	*/
	WeekID strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workforcemanagement businessunit week schedule update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) WithDefaults() *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement businessunit week schedule update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) WithContext(ctx context.Context) *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) WithBody(body *models.ProcessScheduleUpdateUploadRequest) *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) SetBody(body *models.ProcessScheduleUpdateUploadRequest) {
	o.Body = body
}

// WithBusinessUnitID adds the businessUnitID to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) WithBusinessUnitID(businessUnitID string) *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithScheduleID adds the scheduleID to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) WithScheduleID(scheduleID string) *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) SetScheduleID(scheduleID string) {
	o.ScheduleID = scheduleID
}

// WithWeekID adds the weekID to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) WithWeekID(weekID strfmt.Date) *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams {
	o.SetWeekID(weekID)
	return o
}

// SetWeekID adds the weekId to the post workforcemanagement businessunit week schedule update params
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) SetWeekID(weekID strfmt.Date) {
	o.WeekID = weekID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementBusinessunitWeekScheduleUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param businessUnitId
	if err := r.SetPathParam("businessUnitId", o.BusinessUnitID); err != nil {
		return err
	}

	// path param scheduleId
	if err := r.SetPathParam("scheduleId", o.ScheduleID); err != nil {
		return err
	}

	// path param weekId
	if err := r.SetPathParam("weekId", o.WeekID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
