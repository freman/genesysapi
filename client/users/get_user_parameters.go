// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetUserParams creates a new GetUserParams object
// with the default values initialized.
func NewGetUserParams() *GetUserParams {
	var (
		stateDefault = string("active")
	)
	return &GetUserParams{
		State: &stateDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserParamsWithTimeout creates a new GetUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserParamsWithTimeout(timeout time.Duration) *GetUserParams {
	var (
		stateDefault = string("active")
	)
	return &GetUserParams{
		State: &stateDefault,

		timeout: timeout,
	}
}

// NewGetUserParamsWithContext creates a new GetUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserParamsWithContext(ctx context.Context) *GetUserParams {
	var (
		stateDefault = string("active")
	)
	return &GetUserParams{
		State: &stateDefault,

		Context: ctx,
	}
}

// NewGetUserParamsWithHTTPClient creates a new GetUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserParamsWithHTTPClient(client *http.Client) *GetUserParams {
	var (
		stateDefault = string("active")
	)
	return &GetUserParams{
		State:      &stateDefault,
		HTTPClient: client,
	}
}

/*GetUserParams contains all the parameters to send to the API endpoint
for the get user operation typically these are written to a http.Request
*/
type GetUserParams struct {

	/*Expand
	  Which fields, if any, to expand

	*/
	Expand []string
	/*State
	  Search for a user with this state

	*/
	State *string
	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user params
func (o *GetUserParams) WithTimeout(timeout time.Duration) *GetUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user params
func (o *GetUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user params
func (o *GetUserParams) WithContext(ctx context.Context) *GetUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user params
func (o *GetUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user params
func (o *GetUserParams) WithHTTPClient(client *http.Client) *GetUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user params
func (o *GetUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get user params
func (o *GetUserParams) WithExpand(expand []string) *GetUserParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get user params
func (o *GetUserParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithState adds the state to the get user params
func (o *GetUserParams) WithState(state *string) *GetUserParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get user params
func (o *GetUserParams) SetState(state *string) {
	o.State = state
}

// WithUserID adds the userID to the get user params
func (o *GetUserParams) WithUserID(userID string) *GetUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user params
func (o *GetUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
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
