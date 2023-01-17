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

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams() *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParamsWithTimeout creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams object
// with the ability to set a timeout on a request.
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams{
		timeout: timeout,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParamsWithContext creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams object
// with the ability to set a context for a request.
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParamsWithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams{
		Context: ctx,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParamsWithHTTPClient creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams{
		HTTPClient: client,
	}
}

/*
GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams contains all the parameters to send to the API endpoint

	for the get workforcemanagement businessunit week shorttermforecast generationresults operation.

	Typically these are written to a http.Request.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams struct {

	/* BusinessUnitID.

	   The ID of the business unit to which the forecast belongs
	*/
	BusinessUnitID string

	/* ForecastID.

	   The ID of the forecast
	*/
	ForecastID string

	/* WeekDateID.

	   The week start date of the forecast in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	   Format: date
	*/
	WeekDateID strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workforcemanagement businessunit week shorttermforecast generationresults params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) WithDefaults() *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workforcemanagement businessunit week shorttermforecast generationresults params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workforcemanagement businessunit week shorttermforecast generationresults params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement businessunit week shorttermforecast generationresults params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement businessunit week shorttermforecast generationresults params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) WithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement businessunit week shorttermforecast generationresults params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement businessunit week shorttermforecast generationresults params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement businessunit week shorttermforecast generationresults params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBusinessUnitID adds the businessUnitID to the get workforcemanagement businessunit week shorttermforecast generationresults params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) WithBusinessUnitID(businessUnitID string) *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the get workforcemanagement businessunit week shorttermforecast generationresults params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithForecastID adds the forecastID to the get workforcemanagement businessunit week shorttermforecast generationresults params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) WithForecastID(forecastID string) *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams {
	o.SetForecastID(forecastID)
	return o
}

// SetForecastID adds the forecastId to the get workforcemanagement businessunit week shorttermforecast generationresults params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) SetForecastID(forecastID string) {
	o.ForecastID = forecastID
}

// WithWeekDateID adds the weekDateID to the get workforcemanagement businessunit week shorttermforecast generationresults params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) WithWeekDateID(weekDateID strfmt.Date) *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams {
	o.SetWeekDateID(weekDateID)
	return o
}

// SetWeekDateID adds the weekDateId to the get workforcemanagement businessunit week shorttermforecast generationresults params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) SetWeekDateID(weekDateID strfmt.Date) {
	o.WeekDateID = weekDateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
