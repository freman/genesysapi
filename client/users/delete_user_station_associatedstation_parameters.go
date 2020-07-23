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

// NewDeleteUserStationAssociatedstationParams creates a new DeleteUserStationAssociatedstationParams object
// with the default values initialized.
func NewDeleteUserStationAssociatedstationParams() *DeleteUserStationAssociatedstationParams {
	var ()
	return &DeleteUserStationAssociatedstationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserStationAssociatedstationParamsWithTimeout creates a new DeleteUserStationAssociatedstationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserStationAssociatedstationParamsWithTimeout(timeout time.Duration) *DeleteUserStationAssociatedstationParams {
	var ()
	return &DeleteUserStationAssociatedstationParams{

		timeout: timeout,
	}
}

// NewDeleteUserStationAssociatedstationParamsWithContext creates a new DeleteUserStationAssociatedstationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserStationAssociatedstationParamsWithContext(ctx context.Context) *DeleteUserStationAssociatedstationParams {
	var ()
	return &DeleteUserStationAssociatedstationParams{

		Context: ctx,
	}
}

// NewDeleteUserStationAssociatedstationParamsWithHTTPClient creates a new DeleteUserStationAssociatedstationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserStationAssociatedstationParamsWithHTTPClient(client *http.Client) *DeleteUserStationAssociatedstationParams {
	var ()
	return &DeleteUserStationAssociatedstationParams{
		HTTPClient: client,
	}
}

/*DeleteUserStationAssociatedstationParams contains all the parameters to send to the API endpoint
for the delete user station associatedstation operation typically these are written to a http.Request
*/
type DeleteUserStationAssociatedstationParams struct {

	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user station associatedstation params
func (o *DeleteUserStationAssociatedstationParams) WithTimeout(timeout time.Duration) *DeleteUserStationAssociatedstationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user station associatedstation params
func (o *DeleteUserStationAssociatedstationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user station associatedstation params
func (o *DeleteUserStationAssociatedstationParams) WithContext(ctx context.Context) *DeleteUserStationAssociatedstationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user station associatedstation params
func (o *DeleteUserStationAssociatedstationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user station associatedstation params
func (o *DeleteUserStationAssociatedstationParams) WithHTTPClient(client *http.Client) *DeleteUserStationAssociatedstationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user station associatedstation params
func (o *DeleteUserStationAssociatedstationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the delete user station associatedstation params
func (o *DeleteUserStationAssociatedstationParams) WithUserID(userID string) *DeleteUserStationAssociatedstationParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user station associatedstation params
func (o *DeleteUserStationAssociatedstationParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserStationAssociatedstationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
