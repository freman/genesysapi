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

// NewPutUserStationAssociatedstationStationIDParams creates a new PutUserStationAssociatedstationStationIDParams object
// with the default values initialized.
func NewPutUserStationAssociatedstationStationIDParams() *PutUserStationAssociatedstationStationIDParams {
	var ()
	return &PutUserStationAssociatedstationStationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutUserStationAssociatedstationStationIDParamsWithTimeout creates a new PutUserStationAssociatedstationStationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutUserStationAssociatedstationStationIDParamsWithTimeout(timeout time.Duration) *PutUserStationAssociatedstationStationIDParams {
	var ()
	return &PutUserStationAssociatedstationStationIDParams{

		timeout: timeout,
	}
}

// NewPutUserStationAssociatedstationStationIDParamsWithContext creates a new PutUserStationAssociatedstationStationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutUserStationAssociatedstationStationIDParamsWithContext(ctx context.Context) *PutUserStationAssociatedstationStationIDParams {
	var ()
	return &PutUserStationAssociatedstationStationIDParams{

		Context: ctx,
	}
}

// NewPutUserStationAssociatedstationStationIDParamsWithHTTPClient creates a new PutUserStationAssociatedstationStationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutUserStationAssociatedstationStationIDParamsWithHTTPClient(client *http.Client) *PutUserStationAssociatedstationStationIDParams {
	var ()
	return &PutUserStationAssociatedstationStationIDParams{
		HTTPClient: client,
	}
}

/*PutUserStationAssociatedstationStationIDParams contains all the parameters to send to the API endpoint
for the put user station associatedstation station Id operation typically these are written to a http.Request
*/
type PutUserStationAssociatedstationStationIDParams struct {

	/*StationID
	  stationId

	*/
	StationID string
	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put user station associatedstation station Id params
func (o *PutUserStationAssociatedstationStationIDParams) WithTimeout(timeout time.Duration) *PutUserStationAssociatedstationStationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put user station associatedstation station Id params
func (o *PutUserStationAssociatedstationStationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put user station associatedstation station Id params
func (o *PutUserStationAssociatedstationStationIDParams) WithContext(ctx context.Context) *PutUserStationAssociatedstationStationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put user station associatedstation station Id params
func (o *PutUserStationAssociatedstationStationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put user station associatedstation station Id params
func (o *PutUserStationAssociatedstationStationIDParams) WithHTTPClient(client *http.Client) *PutUserStationAssociatedstationStationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put user station associatedstation station Id params
func (o *PutUserStationAssociatedstationStationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStationID adds the stationID to the put user station associatedstation station Id params
func (o *PutUserStationAssociatedstationStationIDParams) WithStationID(stationID string) *PutUserStationAssociatedstationStationIDParams {
	o.SetStationID(stationID)
	return o
}

// SetStationID adds the stationId to the put user station associatedstation station Id params
func (o *PutUserStationAssociatedstationStationIDParams) SetStationID(stationID string) {
	o.StationID = stationID
}

// WithUserID adds the userID to the put user station associatedstation station Id params
func (o *PutUserStationAssociatedstationStationIDParams) WithUserID(userID string) *PutUserStationAssociatedstationStationIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the put user station associatedstation station Id params
func (o *PutUserStationAssociatedstationStationIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PutUserStationAssociatedstationStationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stationId
	if err := r.SetPathParam("stationId", o.StationID); err != nil {
		return err
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
