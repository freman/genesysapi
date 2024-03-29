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
)

// NewGetWorkforcemanagementBusinessunitWeekScheduleParams creates a new GetWorkforcemanagementBusinessunitWeekScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkforcemanagementBusinessunitWeekScheduleParams() *GetWorkforcemanagementBusinessunitWeekScheduleParams {
	return &GetWorkforcemanagementBusinessunitWeekScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleParamsWithTimeout creates a new GetWorkforcemanagementBusinessunitWeekScheduleParams object
// with the ability to set a timeout on a request.
func NewGetWorkforcemanagementBusinessunitWeekScheduleParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitWeekScheduleParams {
	return &GetWorkforcemanagementBusinessunitWeekScheduleParams{
		timeout: timeout,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleParamsWithContext creates a new GetWorkforcemanagementBusinessunitWeekScheduleParams object
// with the ability to set a context for a request.
func NewGetWorkforcemanagementBusinessunitWeekScheduleParamsWithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitWeekScheduleParams {
	return &GetWorkforcemanagementBusinessunitWeekScheduleParams{
		Context: ctx,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekScheduleParamsWithHTTPClient creates a new GetWorkforcemanagementBusinessunitWeekScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkforcemanagementBusinessunitWeekScheduleParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitWeekScheduleParams {
	return &GetWorkforcemanagementBusinessunitWeekScheduleParams{
		HTTPClient: client,
	}
}

/*
GetWorkforcemanagementBusinessunitWeekScheduleParams contains all the parameters to send to the API endpoint

	for the get workforcemanagement businessunit week schedule operation.

	Typically these are written to a http.Request.
*/
type GetWorkforcemanagementBusinessunitWeekScheduleParams struct {

	/* BusinessUnitID.

	   The ID of the business unit
	*/
	BusinessUnitID string

	/* Expand.

	   expand
	*/
	Expand *string

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

// WithDefaults hydrates default values in the get workforcemanagement businessunit week schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) WithDefaults() *GetWorkforcemanagementBusinessunitWeekScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workforcemanagement businessunit week schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitWeekScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) WithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitWeekScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitWeekScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBusinessUnitID adds the businessUnitID to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) WithBusinessUnitID(businessUnitID string) *GetWorkforcemanagementBusinessunitWeekScheduleParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithExpand adds the expand to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) WithExpand(expand *string) *GetWorkforcemanagementBusinessunitWeekScheduleParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithScheduleID adds the scheduleID to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) WithScheduleID(scheduleID string) *GetWorkforcemanagementBusinessunitWeekScheduleParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) SetScheduleID(scheduleID string) {
	o.ScheduleID = scheduleID
}

// WithWeekID adds the weekID to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) WithWeekID(weekID strfmt.Date) *GetWorkforcemanagementBusinessunitWeekScheduleParams {
	o.SetWeekID(weekID)
	return o
}

// SetWeekID adds the weekId to the get workforcemanagement businessunit week schedule params
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) SetWeekID(weekID strfmt.Date) {
	o.WeekID = weekID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementBusinessunitWeekScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param businessUnitId
	if err := r.SetPathParam("businessUnitId", o.BusinessUnitID); err != nil {
		return err
	}

	if o.Expand != nil {

		// query param expand
		var qrExpand string

		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {

			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}
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
