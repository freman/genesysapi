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

	"github.com/freman/genesysapi/models"
)

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams creates a new PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParamsWithTimeout creates a new PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParamsWithContext creates a new PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParamsWithContext(ctx context.Context) *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParamsWithHTTPClient creates a new PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams {
	return &PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement businessunit week shorttermforecasts generate operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams struct {

	// Body.
	Body *models.GenerateBuForecastRequest

	/* BusinessUnitID.

	   The ID of the business unit to which the forecast belongs
	*/
	BusinessUnitID string

	/* ForceAsync.

	   Force the result of this operation to be sent asynchronously via notification.  For testing/app development purposes
	*/
	ForceAsync *bool

	/* WeekDateID.

	   The week start date of the forecast in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	   Format: date
	*/
	WeekDateID strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workforcemanagement businessunit week shorttermforecasts generate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) WithDefaults() *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement businessunit week shorttermforecasts generate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) WithContext(ctx context.Context) *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) WithBody(body *models.GenerateBuForecastRequest) *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) SetBody(body *models.GenerateBuForecastRequest) {
	o.Body = body
}

// WithBusinessUnitID adds the businessUnitID to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) WithBusinessUnitID(businessUnitID string) *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithForceAsync adds the forceAsync to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) WithForceAsync(forceAsync *bool) *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams {
	o.SetForceAsync(forceAsync)
	return o
}

// SetForceAsync adds the forceAsync to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) SetForceAsync(forceAsync *bool) {
	o.ForceAsync = forceAsync
}

// WithWeekDateID adds the weekDateID to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) WithWeekDateID(weekDateID strfmt.Date) *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams {
	o.SetWeekDateID(weekDateID)
	return o
}

// SetWeekDateID adds the weekDateId to the post workforcemanagement businessunit week shorttermforecasts generate params
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) SetWeekDateID(weekDateID strfmt.Date) {
	o.WeekDateID = weekDateID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ForceAsync != nil {

		// query param forceAsync
		var qrForceAsync bool

		if o.ForceAsync != nil {
			qrForceAsync = *o.ForceAsync
		}
		qForceAsync := swag.FormatBool(qrForceAsync)
		if qForceAsync != "" {

			if err := r.SetQueryParam("forceAsync", qForceAsync); err != nil {
				return err
			}
		}
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
