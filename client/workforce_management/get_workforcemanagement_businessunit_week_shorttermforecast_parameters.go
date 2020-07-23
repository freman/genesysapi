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
	"github.com/go-openapi/swag"
)

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastParams creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastParams object
// with the default values initialized.
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastParams() *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams {
	var ()
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastParamsWithTimeout creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams {
	var ()
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastParamsWithContext creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastParamsWithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams {
	var ()
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastParamsWithHTTPClient creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams {
	var ()
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastParams contains all the parameters to send to the API endpoint
for the get workforcemanagement businessunit week shorttermforecast operation typically these are written to a http.Request
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastParams struct {

	/*BusinessUnitID
	  The business unit ID of the business unit to which the forecast belongs

	*/
	BusinessUnitID string
	/*Expand*/
	Expand []string
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

// WithTimeout adds the timeout to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) WithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBusinessUnitID adds the businessUnitID to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) WithBusinessUnitID(businessUnitID string) *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithExpand adds the expand to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) WithExpand(expand []string) *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithForecastID adds the forecastID to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) WithForecastID(forecastID string) *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams {
	o.SetForecastID(forecastID)
	return o
}

// SetForecastID adds the forecastId to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) SetForecastID(forecastID string) {
	o.ForecastID = forecastID
}

// WithWeekDateID adds the weekDateID to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) WithWeekDateID(weekDateID strfmt.Date) *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams {
	o.SetWeekDateID(weekDateID)
	return o
}

// SetWeekDateID adds the weekDateId to the get workforcemanagement businessunit week shorttermforecast params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) SetWeekDateID(weekDateID strfmt.Date) {
	o.WeekDateID = weekDateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param businessUnitId
	if err := r.SetPathParam("businessUnitId", o.BusinessUnitID); err != nil {
		return err
	}

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
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
