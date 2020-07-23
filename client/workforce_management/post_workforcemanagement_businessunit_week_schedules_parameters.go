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

// NewPostWorkforcemanagementBusinessunitWeekSchedulesParams creates a new PostWorkforcemanagementBusinessunitWeekSchedulesParams object
// with the default values initialized.
func NewPostWorkforcemanagementBusinessunitWeekSchedulesParams() *PostWorkforcemanagementBusinessunitWeekSchedulesParams {
	var ()
	return &PostWorkforcemanagementBusinessunitWeekSchedulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesParamsWithTimeout creates a new PostWorkforcemanagementBusinessunitWeekSchedulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWorkforcemanagementBusinessunitWeekSchedulesParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementBusinessunitWeekSchedulesParams {
	var ()
	return &PostWorkforcemanagementBusinessunitWeekSchedulesParams{

		timeout: timeout,
	}
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesParamsWithContext creates a new PostWorkforcemanagementBusinessunitWeekSchedulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWorkforcemanagementBusinessunitWeekSchedulesParamsWithContext(ctx context.Context) *PostWorkforcemanagementBusinessunitWeekSchedulesParams {
	var ()
	return &PostWorkforcemanagementBusinessunitWeekSchedulesParams{

		Context: ctx,
	}
}

// NewPostWorkforcemanagementBusinessunitWeekSchedulesParamsWithHTTPClient creates a new PostWorkforcemanagementBusinessunitWeekSchedulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWorkforcemanagementBusinessunitWeekSchedulesParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementBusinessunitWeekSchedulesParams {
	var ()
	return &PostWorkforcemanagementBusinessunitWeekSchedulesParams{
		HTTPClient: client,
	}
}

/*PostWorkforcemanagementBusinessunitWeekSchedulesParams contains all the parameters to send to the API endpoint
for the post workforcemanagement businessunit week schedules operation typically these are written to a http.Request
*/
type PostWorkforcemanagementBusinessunitWeekSchedulesParams struct {

	/*Body
	  body

	*/
	Body *models.BuCreateBlankScheduleRequest
	/*BusinessUnitID
	  The ID of the business unit

	*/
	BusinessUnitID string
	/*WeekID
	  First day of schedule week in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	*/
	WeekID strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post workforcemanagement businessunit week schedules params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementBusinessunitWeekSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement businessunit week schedules params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement businessunit week schedules params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) WithContext(ctx context.Context) *PostWorkforcemanagementBusinessunitWeekSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement businessunit week schedules params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement businessunit week schedules params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementBusinessunitWeekSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement businessunit week schedules params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement businessunit week schedules params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) WithBody(body *models.BuCreateBlankScheduleRequest) *PostWorkforcemanagementBusinessunitWeekSchedulesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement businessunit week schedules params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) SetBody(body *models.BuCreateBlankScheduleRequest) {
	o.Body = body
}

// WithBusinessUnitID adds the businessUnitID to the post workforcemanagement businessunit week schedules params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) WithBusinessUnitID(businessUnitID string) *PostWorkforcemanagementBusinessunitWeekSchedulesParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the post workforcemanagement businessunit week schedules params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithWeekID adds the weekID to the post workforcemanagement businessunit week schedules params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) WithWeekID(weekID strfmt.Date) *PostWorkforcemanagementBusinessunitWeekSchedulesParams {
	o.SetWeekID(weekID)
	return o
}

// SetWeekID adds the weekId to the post workforcemanagement businessunit week schedules params
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) SetWeekID(weekID strfmt.Date) {
	o.WeekID = weekID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementBusinessunitWeekSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
