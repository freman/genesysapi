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

// NewGetGamificationScorecardsUserPointsAlltimeParams creates a new GetGamificationScorecardsUserPointsAlltimeParams object
// with the default values initialized.
func NewGetGamificationScorecardsUserPointsAlltimeParams() *GetGamificationScorecardsUserPointsAlltimeParams {
	var ()
	return &GetGamificationScorecardsUserPointsAlltimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGamificationScorecardsUserPointsAlltimeParamsWithTimeout creates a new GetGamificationScorecardsUserPointsAlltimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGamificationScorecardsUserPointsAlltimeParamsWithTimeout(timeout time.Duration) *GetGamificationScorecardsUserPointsAlltimeParams {
	var ()
	return &GetGamificationScorecardsUserPointsAlltimeParams{

		timeout: timeout,
	}
}

// NewGetGamificationScorecardsUserPointsAlltimeParamsWithContext creates a new GetGamificationScorecardsUserPointsAlltimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGamificationScorecardsUserPointsAlltimeParamsWithContext(ctx context.Context) *GetGamificationScorecardsUserPointsAlltimeParams {
	var ()
	return &GetGamificationScorecardsUserPointsAlltimeParams{

		Context: ctx,
	}
}

// NewGetGamificationScorecardsUserPointsAlltimeParamsWithHTTPClient creates a new GetGamificationScorecardsUserPointsAlltimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGamificationScorecardsUserPointsAlltimeParamsWithHTTPClient(client *http.Client) *GetGamificationScorecardsUserPointsAlltimeParams {
	var ()
	return &GetGamificationScorecardsUserPointsAlltimeParams{
		HTTPClient: client,
	}
}

/*GetGamificationScorecardsUserPointsAlltimeParams contains all the parameters to send to the API endpoint
for the get gamification scorecards user points alltime operation typically these are written to a http.Request
*/
type GetGamificationScorecardsUserPointsAlltimeParams struct {

	/*EndWorkday
	  End workday of querying workdays range. Workday is represented as an ISO-8601 string. For example: yyyy-MM-dd

	*/
	EndWorkday strfmt.Date
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gamification scorecards user points alltime params
func (o *GetGamificationScorecardsUserPointsAlltimeParams) WithTimeout(timeout time.Duration) *GetGamificationScorecardsUserPointsAlltimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gamification scorecards user points alltime params
func (o *GetGamificationScorecardsUserPointsAlltimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gamification scorecards user points alltime params
func (o *GetGamificationScorecardsUserPointsAlltimeParams) WithContext(ctx context.Context) *GetGamificationScorecardsUserPointsAlltimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gamification scorecards user points alltime params
func (o *GetGamificationScorecardsUserPointsAlltimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gamification scorecards user points alltime params
func (o *GetGamificationScorecardsUserPointsAlltimeParams) WithHTTPClient(client *http.Client) *GetGamificationScorecardsUserPointsAlltimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gamification scorecards user points alltime params
func (o *GetGamificationScorecardsUserPointsAlltimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndWorkday adds the endWorkday to the get gamification scorecards user points alltime params
func (o *GetGamificationScorecardsUserPointsAlltimeParams) WithEndWorkday(endWorkday strfmt.Date) *GetGamificationScorecardsUserPointsAlltimeParams {
	o.SetEndWorkday(endWorkday)
	return o
}

// SetEndWorkday adds the endWorkday to the get gamification scorecards user points alltime params
func (o *GetGamificationScorecardsUserPointsAlltimeParams) SetEndWorkday(endWorkday strfmt.Date) {
	o.EndWorkday = endWorkday
}

// WithUserID adds the userID to the get gamification scorecards user points alltime params
func (o *GetGamificationScorecardsUserPointsAlltimeParams) WithUserID(userID string) *GetGamificationScorecardsUserPointsAlltimeParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get gamification scorecards user points alltime params
func (o *GetGamificationScorecardsUserPointsAlltimeParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGamificationScorecardsUserPointsAlltimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
