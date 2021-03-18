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

// NewGetGamificationLeaderboardParams creates a new GetGamificationLeaderboardParams object
// with the default values initialized.
func NewGetGamificationLeaderboardParams() *GetGamificationLeaderboardParams {
	var ()
	return &GetGamificationLeaderboardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGamificationLeaderboardParamsWithTimeout creates a new GetGamificationLeaderboardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGamificationLeaderboardParamsWithTimeout(timeout time.Duration) *GetGamificationLeaderboardParams {
	var ()
	return &GetGamificationLeaderboardParams{

		timeout: timeout,
	}
}

// NewGetGamificationLeaderboardParamsWithContext creates a new GetGamificationLeaderboardParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGamificationLeaderboardParamsWithContext(ctx context.Context) *GetGamificationLeaderboardParams {
	var ()
	return &GetGamificationLeaderboardParams{

		Context: ctx,
	}
}

// NewGetGamificationLeaderboardParamsWithHTTPClient creates a new GetGamificationLeaderboardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGamificationLeaderboardParamsWithHTTPClient(client *http.Client) *GetGamificationLeaderboardParams {
	var ()
	return &GetGamificationLeaderboardParams{
		HTTPClient: client,
	}
}

/*GetGamificationLeaderboardParams contains all the parameters to send to the API endpoint
for the get gamification leaderboard operation typically these are written to a http.Request
*/
type GetGamificationLeaderboardParams struct {

	/*EndWorkday
	  End workday to retrieve for the leaderboard. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	*/
	EndWorkday strfmt.Date
	/*MetricID
	  Metric Id for which the leaderboard is to be generated. The total points is used if nothing is given.

	*/
	MetricID *string
	/*StartWorkday
	  Start workday to retrieve for the leaderboard. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	*/
	StartWorkday strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gamification leaderboard params
func (o *GetGamificationLeaderboardParams) WithTimeout(timeout time.Duration) *GetGamificationLeaderboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gamification leaderboard params
func (o *GetGamificationLeaderboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gamification leaderboard params
func (o *GetGamificationLeaderboardParams) WithContext(ctx context.Context) *GetGamificationLeaderboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gamification leaderboard params
func (o *GetGamificationLeaderboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gamification leaderboard params
func (o *GetGamificationLeaderboardParams) WithHTTPClient(client *http.Client) *GetGamificationLeaderboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gamification leaderboard params
func (o *GetGamificationLeaderboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndWorkday adds the endWorkday to the get gamification leaderboard params
func (o *GetGamificationLeaderboardParams) WithEndWorkday(endWorkday strfmt.Date) *GetGamificationLeaderboardParams {
	o.SetEndWorkday(endWorkday)
	return o
}

// SetEndWorkday adds the endWorkday to the get gamification leaderboard params
func (o *GetGamificationLeaderboardParams) SetEndWorkday(endWorkday strfmt.Date) {
	o.EndWorkday = endWorkday
}

// WithMetricID adds the metricID to the get gamification leaderboard params
func (o *GetGamificationLeaderboardParams) WithMetricID(metricID *string) *GetGamificationLeaderboardParams {
	o.SetMetricID(metricID)
	return o
}

// SetMetricID adds the metricId to the get gamification leaderboard params
func (o *GetGamificationLeaderboardParams) SetMetricID(metricID *string) {
	o.MetricID = metricID
}

// WithStartWorkday adds the startWorkday to the get gamification leaderboard params
func (o *GetGamificationLeaderboardParams) WithStartWorkday(startWorkday strfmt.Date) *GetGamificationLeaderboardParams {
	o.SetStartWorkday(startWorkday)
	return o
}

// SetStartWorkday adds the startWorkday to the get gamification leaderboard params
func (o *GetGamificationLeaderboardParams) SetStartWorkday(startWorkday strfmt.Date) {
	o.StartWorkday = startWorkday
}

// WriteToRequest writes these params to a swagger request
func (o *GetGamificationLeaderboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.MetricID != nil {

		// query param metricId
		var qrMetricID string
		if o.MetricID != nil {
			qrMetricID = *o.MetricID
		}
		qMetricID := qrMetricID
		if qMetricID != "" {
			if err := r.SetQueryParam("metricId", qMetricID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}