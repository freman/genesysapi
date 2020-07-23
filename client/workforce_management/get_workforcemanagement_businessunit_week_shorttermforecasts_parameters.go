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

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsParams creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams object
// with the default values initialized.
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsParams() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams {
	var ()
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsParamsWithTimeout creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams {
	var ()
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsParamsWithContext creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsParamsWithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams {
	var ()
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsParamsWithHTTPClient creates a new GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams {
	var ()
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams contains all the parameters to send to the API endpoint
for the get workforcemanagement businessunit week shorttermforecasts operation typically these are written to a http.Request
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams struct {

	/*BusinessUnitID
	  The business unit ID of the business unit to which the forecast belongs

	*/
	BusinessUnitID string
	/*WeekDateID
	  The week start date of the forecast in yyyy-MM-dd format or 'recent' to fetch recent forecasts

	*/
	WeekDateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement businessunit week shorttermforecasts params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement businessunit week shorttermforecasts params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement businessunit week shorttermforecasts params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams) WithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement businessunit week shorttermforecasts params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement businessunit week shorttermforecasts params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement businessunit week shorttermforecasts params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBusinessUnitID adds the businessUnitID to the get workforcemanagement businessunit week shorttermforecasts params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams) WithBusinessUnitID(businessUnitID string) *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the get workforcemanagement businessunit week shorttermforecasts params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithWeekDateID adds the weekDateID to the get workforcemanagement businessunit week shorttermforecasts params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams) WithWeekDateID(weekDateID string) *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams {
	o.SetWeekDateID(weekDateID)
	return o
}

// SetWeekDateID adds the weekDateId to the get workforcemanagement businessunit week shorttermforecasts params
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams) SetWeekDateID(weekDateID string) {
	o.WeekDateID = weekDateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param businessUnitId
	if err := r.SetPathParam("businessUnitId", o.BusinessUnitID); err != nil {
		return err
	}

	// path param weekDateId
	if err := r.SetPathParam("weekDateId", o.WeekDateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
