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

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams creates a new PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateParamsWithTimeout creates a new PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateParamsWithContext creates a new PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateParamsWithContext(ctx context.Context) *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateParamsWithHTTPClient creates a new PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementBusinessunitWeekSchedulesGenerateParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams {
	return &PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement businessunit week schedules generate operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams struct {

	/* Body.

	   body
	*/
	Body *models.BuGenerateScheduleRequest

	/* BusinessUnitID.

	   The ID of the business unit
	*/
	BusinessUnitID string

	/* WeekID.

	   First day of schedule week in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	   Format: date
	*/
	WeekID strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workforcemanagement businessunit week schedules generate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) WithDefaults() *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement businessunit week schedules generate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workforcemanagement businessunit week schedules generate params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement businessunit week schedules generate params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement businessunit week schedules generate params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) WithContext(ctx context.Context) *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement businessunit week schedules generate params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement businessunit week schedules generate params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement businessunit week schedules generate params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement businessunit week schedules generate params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) WithBody(body *models.BuGenerateScheduleRequest) *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement businessunit week schedules generate params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) SetBody(body *models.BuGenerateScheduleRequest) {
	o.Body = body
}

// WithBusinessUnitID adds the businessUnitID to the post workforcemanagement businessunit week schedules generate params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) WithBusinessUnitID(businessUnitID string) *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the post workforcemanagement businessunit week schedules generate params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithWeekID adds the weekID to the post workforcemanagement businessunit week schedules generate params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) WithWeekID(weekID strfmt.Date) *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams {
	o.SetWeekID(weekID)
	return o
}

// SetWeekID adds the weekId to the post workforcemanagement businessunit week schedules generate params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) SetWeekID(weekID strfmt.Date) {
	o.WeekID = weekID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesGenerateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param weekId
	if err := r.SetPathParam("weekId", o.WeekID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
