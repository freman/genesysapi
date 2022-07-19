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

// NewGetGamificationScorecardsProfileMetricUserValuesTrendsParams creates a new GetGamificationScorecardsProfileMetricUserValuesTrendsParams object
// with the default values initialized.
func NewGetGamificationScorecardsProfileMetricUserValuesTrendsParams() *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &GetGamificationScorecardsProfileMetricUserValuesTrendsParams{
		TimeZone: &timeZoneDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGamificationScorecardsProfileMetricUserValuesTrendsParamsWithTimeout creates a new GetGamificationScorecardsProfileMetricUserValuesTrendsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGamificationScorecardsProfileMetricUserValuesTrendsParamsWithTimeout(timeout time.Duration) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &GetGamificationScorecardsProfileMetricUserValuesTrendsParams{
		TimeZone: &timeZoneDefault,

		timeout: timeout,
	}
}

// NewGetGamificationScorecardsProfileMetricUserValuesTrendsParamsWithContext creates a new GetGamificationScorecardsProfileMetricUserValuesTrendsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGamificationScorecardsProfileMetricUserValuesTrendsParamsWithContext(ctx context.Context) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &GetGamificationScorecardsProfileMetricUserValuesTrendsParams{
		TimeZone: &timeZoneDefault,

		Context: ctx,
	}
}

// NewGetGamificationScorecardsProfileMetricUserValuesTrendsParamsWithHTTPClient creates a new GetGamificationScorecardsProfileMetricUserValuesTrendsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGamificationScorecardsProfileMetricUserValuesTrendsParamsWithHTTPClient(client *http.Client) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	var (
		timeZoneDefault = string("UTC")
	)
	return &GetGamificationScorecardsProfileMetricUserValuesTrendsParams{
		TimeZone:   &timeZoneDefault,
		HTTPClient: client,
	}
}

/*GetGamificationScorecardsProfileMetricUserValuesTrendsParams contains all the parameters to send to the API endpoint
for the get gamification scorecards profile metric user values trends operation typically these are written to a http.Request
*/
type GetGamificationScorecardsProfileMetricUserValuesTrendsParams struct {

	/*EndWorkday
	  End workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	*/
	EndWorkday strfmt.Date
	/*MetricID
	  metricId

	*/
	MetricID string
	/*ProfileID
	  performanceProfileId

	*/
	ProfileID string
	/*ReferenceWorkday
	  Reference workday for the trend. Used to determine the associated metric definition. If not set, then the value of endWorkday is used. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	*/
	ReferenceWorkday *strfmt.Date
	/*StartWorkday
	  Start workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	*/
	StartWorkday strfmt.Date
	/*TimeZone
	  Timezone for the workday. Defaults to UTC

	*/
	TimeZone *string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) WithTimeout(timeout time.Duration) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) WithContext(ctx context.Context) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) WithHTTPClient(client *http.Client) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndWorkday adds the endWorkday to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) WithEndWorkday(endWorkday strfmt.Date) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	o.SetEndWorkday(endWorkday)
	return o
}

// SetEndWorkday adds the endWorkday to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) SetEndWorkday(endWorkday strfmt.Date) {
	o.EndWorkday = endWorkday
}

// WithMetricID adds the metricID to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) WithMetricID(metricID string) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	o.SetMetricID(metricID)
	return o
}

// SetMetricID adds the metricId to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) SetMetricID(metricID string) {
	o.MetricID = metricID
}

// WithProfileID adds the profileID to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) WithProfileID(profileID string) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) SetProfileID(profileID string) {
	o.ProfileID = profileID
}

// WithReferenceWorkday adds the referenceWorkday to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) WithReferenceWorkday(referenceWorkday *strfmt.Date) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	o.SetReferenceWorkday(referenceWorkday)
	return o
}

// SetReferenceWorkday adds the referenceWorkday to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) SetReferenceWorkday(referenceWorkday *strfmt.Date) {
	o.ReferenceWorkday = referenceWorkday
}

// WithStartWorkday adds the startWorkday to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) WithStartWorkday(startWorkday strfmt.Date) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	o.SetStartWorkday(startWorkday)
	return o
}

// SetStartWorkday adds the startWorkday to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) SetStartWorkday(startWorkday strfmt.Date) {
	o.StartWorkday = startWorkday
}

// WithTimeZone adds the timeZone to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) WithTimeZone(timeZone *string) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	o.SetTimeZone(timeZone)
	return o
}

// SetTimeZone adds the timeZone to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) SetTimeZone(timeZone *string) {
	o.TimeZone = timeZone
}

// WithUserID adds the userID to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) WithUserID(userID string) *GetGamificationScorecardsProfileMetricUserValuesTrendsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get gamification scorecards profile metric user values trends params
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGamificationScorecardsProfileMetricUserValuesTrendsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
