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
)

// NewDeleteUserStationDefaultstationParams creates a new DeleteUserStationDefaultstationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserStationDefaultstationParams() *DeleteUserStationDefaultstationParams {
	return &DeleteUserStationDefaultstationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserStationDefaultstationParamsWithTimeout creates a new DeleteUserStationDefaultstationParams object
// with the ability to set a timeout on a request.
func NewDeleteUserStationDefaultstationParamsWithTimeout(timeout time.Duration) *DeleteUserStationDefaultstationParams {
	return &DeleteUserStationDefaultstationParams{
		timeout: timeout,
	}
}

// NewDeleteUserStationDefaultstationParamsWithContext creates a new DeleteUserStationDefaultstationParams object
// with the ability to set a context for a request.
func NewDeleteUserStationDefaultstationParamsWithContext(ctx context.Context) *DeleteUserStationDefaultstationParams {
	return &DeleteUserStationDefaultstationParams{
		Context: ctx,
	}
}

// NewDeleteUserStationDefaultstationParamsWithHTTPClient creates a new DeleteUserStationDefaultstationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserStationDefaultstationParamsWithHTTPClient(client *http.Client) *DeleteUserStationDefaultstationParams {
	return &DeleteUserStationDefaultstationParams{
		HTTPClient: client,
	}
}

/*
DeleteUserStationDefaultstationParams contains all the parameters to send to the API endpoint

	for the delete user station defaultstation operation.

	Typically these are written to a http.Request.
*/
type DeleteUserStationDefaultstationParams struct {

	/* UserID.

	   User ID
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user station defaultstation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserStationDefaultstationParams) WithDefaults() *DeleteUserStationDefaultstationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user station defaultstation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserStationDefaultstationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user station defaultstation params
func (o *DeleteUserStationDefaultstationParams) WithTimeout(timeout time.Duration) *DeleteUserStationDefaultstationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user station defaultstation params
func (o *DeleteUserStationDefaultstationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user station defaultstation params
func (o *DeleteUserStationDefaultstationParams) WithContext(ctx context.Context) *DeleteUserStationDefaultstationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user station defaultstation params
func (o *DeleteUserStationDefaultstationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user station defaultstation params
func (o *DeleteUserStationDefaultstationParams) WithHTTPClient(client *http.Client) *DeleteUserStationDefaultstationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user station defaultstation params
func (o *DeleteUserStationDefaultstationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the delete user station defaultstation params
func (o *DeleteUserStationDefaultstationParams) WithUserID(userID string) *DeleteUserStationDefaultstationParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user station defaultstation params
func (o *DeleteUserStationDefaultstationParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserStationDefaultstationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
