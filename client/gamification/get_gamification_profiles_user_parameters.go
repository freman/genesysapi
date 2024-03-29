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

// NewGetGamificationProfilesUserParams creates a new GetGamificationProfilesUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGamificationProfilesUserParams() *GetGamificationProfilesUserParams {
	return &GetGamificationProfilesUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGamificationProfilesUserParamsWithTimeout creates a new GetGamificationProfilesUserParams object
// with the ability to set a timeout on a request.
func NewGetGamificationProfilesUserParamsWithTimeout(timeout time.Duration) *GetGamificationProfilesUserParams {
	return &GetGamificationProfilesUserParams{
		timeout: timeout,
	}
}

// NewGetGamificationProfilesUserParamsWithContext creates a new GetGamificationProfilesUserParams object
// with the ability to set a context for a request.
func NewGetGamificationProfilesUserParamsWithContext(ctx context.Context) *GetGamificationProfilesUserParams {
	return &GetGamificationProfilesUserParams{
		Context: ctx,
	}
}

// NewGetGamificationProfilesUserParamsWithHTTPClient creates a new GetGamificationProfilesUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGamificationProfilesUserParamsWithHTTPClient(client *http.Client) *GetGamificationProfilesUserParams {
	return &GetGamificationProfilesUserParams{
		HTTPClient: client,
	}
}

/*
GetGamificationProfilesUserParams contains all the parameters to send to the API endpoint

	for the get gamification profiles user operation.

	Typically these are written to a http.Request.
*/
type GetGamificationProfilesUserParams struct {

	// UserID.
	UserID string

	/* Workday.

	   Target querying workday. If not provided, then queries the current performance profile. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	   Format: date
	*/
	Workday *strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get gamification profiles user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGamificationProfilesUserParams) WithDefaults() *GetGamificationProfilesUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gamification profiles user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGamificationProfilesUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get gamification profiles user params
func (o *GetGamificationProfilesUserParams) WithTimeout(timeout time.Duration) *GetGamificationProfilesUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gamification profiles user params
func (o *GetGamificationProfilesUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gamification profiles user params
func (o *GetGamificationProfilesUserParams) WithContext(ctx context.Context) *GetGamificationProfilesUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gamification profiles user params
func (o *GetGamificationProfilesUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gamification profiles user params
func (o *GetGamificationProfilesUserParams) WithHTTPClient(client *http.Client) *GetGamificationProfilesUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gamification profiles user params
func (o *GetGamificationProfilesUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get gamification profiles user params
func (o *GetGamificationProfilesUserParams) WithUserID(userID string) *GetGamificationProfilesUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get gamification profiles user params
func (o *GetGamificationProfilesUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithWorkday adds the workday to the get gamification profiles user params
func (o *GetGamificationProfilesUserParams) WithWorkday(workday *strfmt.Date) *GetGamificationProfilesUserParams {
	o.SetWorkday(workday)
	return o
}

// SetWorkday adds the workday to the get gamification profiles user params
func (o *GetGamificationProfilesUserParams) SetWorkday(workday *strfmt.Date) {
	o.Workday = workday
}

// WriteToRequest writes these params to a swagger request
func (o *GetGamificationProfilesUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if o.Workday != nil {

		// query param workday
		var qrWorkday strfmt.Date

		if o.Workday != nil {
			qrWorkday = *o.Workday
		}
		qWorkday := qrWorkday.String()
		if qWorkday != "" {

			if err := r.SetQueryParam("workday", qWorkday); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
