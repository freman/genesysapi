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

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams object
// with the default values initialized.
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams() *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams {
	var ()
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParamsWithTimeout creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams {
	var ()
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParamsWithContext creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParamsWithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams {
	var ()
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParamsWithHTTPClient creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams {
	var ()
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams contains all the parameters to send to the API endpoint
for the get workforcemanagement businessunit week shorttermforecast planninggroups operation typically these are written to a http.Request
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams struct {

	/*BusinessUnitID
	  The business unit ID of the business unit to which the forecast belongs

	*/
	BusinessUnitID string
	/*ForecastID
	  The ID of the forecast

	*/
	ForecastID string
	/*WeekDateID
	  The week start date of the forecast in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	*/
	WeekDateID strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement businessunit week shorttermforecast planninggroups params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement businessunit week shorttermforecast planninggroups params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement businessunit week shorttermforecast planninggroups params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) WithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement businessunit week shorttermforecast planninggroups params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement businessunit week shorttermforecast planninggroups params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement businessunit week shorttermforecast planninggroups params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBusinessUnitID adds the businessUnitID to the get workforcemanagement businessunit week shorttermforecast planninggroups params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) WithBusinessUnitID(businessUnitID string) *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the get workforcemanagement businessunit week shorttermforecast planninggroups params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithForecastID adds the forecastID to the get workforcemanagement businessunit week shorttermforecast planninggroups params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) WithForecastID(forecastID string) *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams {
	o.SetForecastID(forecastID)
	return o
}

// SetForecastID adds the forecastId to the get workforcemanagement businessunit week shorttermforecast planninggroups params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) SetForecastID(forecastID string) {
	o.ForecastID = forecastID
}

// WithWeekDateID adds the weekDateID to the get workforcemanagement businessunit week shorttermforecast planninggroups params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) WithWeekDateID(weekDateID strfmt.Date) *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams {
	o.SetWeekDateID(weekDateID)
	return o
}

// SetWeekDateID adds the weekDateId to the get workforcemanagement businessunit week shorttermforecast planninggroups params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) SetWeekDateID(weekDateID strfmt.Date) {
	o.WeekDateID = weekDateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param businessUnitId
	if err := r.SetPathParam("businessUnitId", o.BusinessUnitID); err != nil {
		return err
	}

	// path param forecastId
	if err := r.SetPathParam("forecastId", o.ForecastID); err != nil {
		return err
	}

	// path param weekDateId
	if err := r.SetPathParam("weekDateId", o.WeekDateID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
