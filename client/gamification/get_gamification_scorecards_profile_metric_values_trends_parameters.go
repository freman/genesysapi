// Code generated by go-swagger; DO NOT EDIT.

package gamification

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

// NewGetGamificationScorecardsProfileMetricValuesTrendsParams creates a new GetGamificationScorecardsProfileMetricValuesTrendsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGamificationScorecardsProfileMetricValuesTrendsParams() *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	return &GetGamificationScorecardsProfileMetricValuesTrendsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGamificationScorecardsProfileMetricValuesTrendsParamsWithTimeout creates a new GetGamificationScorecardsProfileMetricValuesTrendsParams object
// with the ability to set a timeout on a request.
func NewGetGamificationScorecardsProfileMetricValuesTrendsParamsWithTimeout(timeout time.Duration) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	return &GetGamificationScorecardsProfileMetricValuesTrendsParams{
		timeout: timeout,
	}
}

// NewGetGamificationScorecardsProfileMetricValuesTrendsParamsWithContext creates a new GetGamificationScorecardsProfileMetricValuesTrendsParams object
// with the ability to set a context for a request.
func NewGetGamificationScorecardsProfileMetricValuesTrendsParamsWithContext(ctx context.Context) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	return &GetGamificationScorecardsProfileMetricValuesTrendsParams{
		Context: ctx,
	}
}

// NewGetGamificationScorecardsProfileMetricValuesTrendsParamsWithHTTPClient creates a new GetGamificationScorecardsProfileMetricValuesTrendsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGamificationScorecardsProfileMetricValuesTrendsParamsWithHTTPClient(client *http.Client) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	return &GetGamificationScorecardsProfileMetricValuesTrendsParams{
		HTTPClient: client,
	}
}

/*
GetGamificationScorecardsProfileMetricValuesTrendsParams contains all the parameters to send to the API endpoint

	for the get gamification scorecards profile metric values trends operation.

	Typically these are written to a http.Request.
*/
type GetGamificationScorecardsProfileMetricValuesTrendsParams struct {

	/* EndWorkday.

	   End workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	   Format: date
	*/
	EndWorkday strfmt.Date

	/* FilterType.

	   Filter type for the query request. If not set, returns the values trends of the requesting user
	*/
	FilterType *string

	/* MetricID.

	   metricId
	*/
	MetricID string

	/* ProfileID.

	   performanceProfileId
	*/
	ProfileID string

	/* ReferenceWorkday.

	   Reference workday for the trend. Used to determine the associated metric definition. If not set, then the value of endWorkday is used. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	   Format: date
	*/
	ReferenceWorkday *strfmt.Date

	/* StartWorkday.

	   Start workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	   Format: date
	*/
	StartWorkday strfmt.Date

	/* TimeZone.

	   Timezone for the workday. Defaults to UTC

	   Default: "UTC"
	*/
	TimeZone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get gamification scorecards profile metric values trends params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) WithDefaults() *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gamification scorecards profile metric values trends params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) SetDefaults() {
	var (
		timeZoneDefault = string("UTC")
	)

	val := GetGamificationScorecardsProfileMetricValuesTrendsParams{
		TimeZone: &timeZoneDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) WithTimeout(timeout time.Duration) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) WithContext(ctx context.Context) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) WithHTTPClient(client *http.Client) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndWorkday adds the endWorkday to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) WithEndWorkday(endWorkday strfmt.Date) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	o.SetEndWorkday(endWorkday)
	return o
}

// SetEndWorkday adds the endWorkday to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) SetEndWorkday(endWorkday strfmt.Date) {
	o.EndWorkday = endWorkday
}

// WithFilterType adds the filterType to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) WithFilterType(filterType *string) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	o.SetFilterType(filterType)
	return o
}

// SetFilterType adds the filterType to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) SetFilterType(filterType *string) {
	o.FilterType = filterType
}

// WithMetricID adds the metricID to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) WithMetricID(metricID string) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	o.SetMetricID(metricID)
	return o
}

// SetMetricID adds the metricId to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) SetMetricID(metricID string) {
	o.MetricID = metricID
}

// WithProfileID adds the profileID to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) WithProfileID(profileID string) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) SetProfileID(profileID string) {
	o.ProfileID = profileID
}

// WithReferenceWorkday adds the referenceWorkday to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) WithReferenceWorkday(referenceWorkday *strfmt.Date) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	o.SetReferenceWorkday(referenceWorkday)
	return o
}

// SetReferenceWorkday adds the referenceWorkday to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) SetReferenceWorkday(referenceWorkday *strfmt.Date) {
	o.ReferenceWorkday = referenceWorkday
}

// WithStartWorkday adds the startWorkday to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) WithStartWorkday(startWorkday strfmt.Date) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	o.SetStartWorkday(startWorkday)
	return o
}

// SetStartWorkday adds the startWorkday to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) SetStartWorkday(startWorkday strfmt.Date) {
	o.StartWorkday = startWorkday
}

// WithTimeZone adds the timeZone to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) WithTimeZone(timeZone *string) *GetGamificationScorecardsProfileMetricValuesTrendsParams {
	o.SetTimeZone(timeZone)
	return o
}

// SetTimeZone adds the timeZone to the get gamification scorecards profile metric values trends params
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) SetTimeZone(timeZone *string) {
	o.TimeZone = timeZone
}

// WriteToRequest writes these params to a swagger request
func (o *GetGamificationScorecardsProfileMetricValuesTrendsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param endWorkday
	qrEndWorkday := o.EndWorkday
	qEndWorkday := qrEndWorkday.String()
	if qEndWorkday != "" {

		if err := r.SetQueryParam("endWorkday", qEndWorkday); err != nil {
			return err
		}
	}

	if o.FilterType != nil {

		// query param filterType
		var qrFilterType string

		if o.FilterType != nil {
			qrFilterType = *o.FilterType
		}
		qFilterType := qrFilterType
		if qFilterType != "" {

			if err := r.SetQueryParam("filterType", qFilterType); err != nil {
				return err
			}
		}
	}

	// path param metricId
	if err := r.SetPathParam("metricId", o.MetricID); err != nil {
		return err
	}

	// path param profileId
	if err := r.SetPathParam("profileId", o.ProfileID); err != nil {
		return err
	}

	if o.ReferenceWorkday != nil {

		// query param referenceWorkday
		var qrReferenceWorkday strfmt.Date

		if o.ReferenceWorkday != nil {
			qrReferenceWorkday = *o.ReferenceWorkday
		}
		qReferenceWorkday := qrReferenceWorkday.String()
		if qReferenceWorkday != "" {

			if err := r.SetQueryParam("referenceWorkday", qReferenceWorkday); err != nil {
				return err
			}
		}
	}

	// query param startWorkday
	qrStartWorkday := o.StartWorkday
	qStartWorkday := qrStartWorkday.String()
	if qStartWorkday != "" {

		if err := r.SetQueryParam("startWorkday", qStartWorkday); err != nil {
			return err
		}
	}

	if o.TimeZone != nil {

		// query param timeZone
		var qrTimeZone string

		if o.TimeZone != nil {
			qrTimeZone = *o.TimeZone
		}
		qTimeZone := qrTimeZone
		if qTimeZone != "" {

			if err := r.SetQueryParam("timeZone", qTimeZone); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
