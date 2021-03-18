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

// NewGetGamificationScorecardsUsersPointsAverageParams creates a new GetGamificationScorecardsUsersPointsAverageParams object
// with the default values initialized.
func NewGetGamificationScorecardsUsersPointsAverageParams() *GetGamificationScorecardsUsersPointsAverageParams {
	var ()
	return &GetGamificationScorecardsUsersPointsAverageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGamificationScorecardsUsersPointsAverageParamsWithTimeout creates a new GetGamificationScorecardsUsersPointsAverageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGamificationScorecardsUsersPointsAverageParamsWithTimeout(timeout time.Duration) *GetGamificationScorecardsUsersPointsAverageParams {
	var ()
	return &GetGamificationScorecardsUsersPointsAverageParams{

		timeout: timeout,
	}
}

// NewGetGamificationScorecardsUsersPointsAverageParamsWithContext creates a new GetGamificationScorecardsUsersPointsAverageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGamificationScorecardsUsersPointsAverageParamsWithContext(ctx context.Context) *GetGamificationScorecardsUsersPointsAverageParams {
	var ()
	return &GetGamificationScorecardsUsersPointsAverageParams{

		Context: ctx,
	}
}

// NewGetGamificationScorecardsUsersPointsAverageParamsWithHTTPClient creates a new GetGamificationScorecardsUsersPointsAverageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGamificationScorecardsUsersPointsAverageParamsWithHTTPClient(client *http.Client) *GetGamificationScorecardsUsersPointsAverageParams {
	var ()
	return &GetGamificationScorecardsUsersPointsAverageParams{
		HTTPClient: client,
	}
}

/*GetGamificationScorecardsUsersPointsAverageParams contains all the parameters to send to the API endpoint
for the get gamification scorecards users points average operation typically these are written to a http.Request
*/
type GetGamificationScorecardsUsersPointsAverageParams struct {

	/*FilterID
	  ID for the filter type.

	*/
	FilterID string
	/*FilterType
	  Filter type for the query request.

	*/
	FilterType string
	/*Workday
	  The target workday. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	*/
	Workday strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gamification scorecards users points average params
func (o *GetGamificationScorecardsUsersPointsAverageParams) WithTimeout(timeout time.Duration) *GetGamificationScorecardsUsersPointsAverageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gamification scorecards users points average params
func (o *GetGamificationScorecardsUsersPointsAverageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gamification scorecards users points average params
func (o *GetGamificationScorecardsUsersPointsAverageParams) WithContext(ctx context.Context) *GetGamificationScorecardsUsersPointsAverageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gamification scorecards users points average params
func (o *GetGamificationScorecardsUsersPointsAverageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gamification scorecards users points average params
func (o *GetGamificationScorecardsUsersPointsAverageParams) WithHTTPClient(client *http.Client) *GetGamificationScorecardsUsersPointsAverageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gamification scorecards users points average params
func (o *GetGamificationScorecardsUsersPointsAverageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterID adds the filterID to the get gamification scorecards users points average params
func (o *GetGamificationScorecardsUsersPointsAverageParams) WithFilterID(filterID string) *GetGamificationScorecardsUsersPointsAverageParams {
	o.SetFilterID(filterID)
	return o
}

// SetFilterID adds the filterId to the get gamification scorecards users points average params
func (o *GetGamificationScorecardsUsersPointsAverageParams) SetFilterID(filterID string) {
	o.FilterID = filterID
}

// WithFilterType adds the filterType to the get gamification scorecards users points average params
func (o *GetGamificationScorecardsUsersPointsAverageParams) WithFilterType(filterType string) *GetGamificationScorecardsUsersPointsAverageParams {
	o.SetFilterType(filterType)
	return o
}

// SetFilterType adds the filterType to the get gamification scorecards users points average params
func (o *GetGamificationScorecardsUsersPointsAverageParams) SetFilterType(filterType string) {
	o.FilterType = filterType
}

// WithWorkday adds the workday to the get gamification scorecards users points average params
func (o *GetGamificationScorecardsUsersPointsAverageParams) WithWorkday(workday strfmt.Date) *GetGamificationScorecardsUsersPointsAverageParams {
	o.SetWorkday(workday)
	return o
}

// SetWorkday adds the workday to the get gamification scorecards users points average params
func (o *GetGamificationScorecardsUsersPointsAverageParams) SetWorkday(workday strfmt.Date) {
	o.Workday = workday
}

// WriteToRequest writes these params to a swagger request
func (o *GetGamificationScorecardsUsersPointsAverageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param filterId
	qrFilterID := o.FilterID
	qFilterID := qrFilterID
	if qFilterID != "" {
		if err := r.SetQueryParam("filterId", qFilterID); err != nil {
			return err
		}
	}

	// query param filterType
	qrFilterType := o.FilterType
	qFilterType := qrFilterType
	if qFilterType != "" {
		if err := r.SetQueryParam("filterType", qFilterType); err != nil {
			return err
		}
	}

	// query param workday
	qrWorkday := o.Workday
	qWorkday := qrWorkday.String()
	if qWorkday != "" {
		if err := r.SetQueryParam("workday", qWorkday); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}