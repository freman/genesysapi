// Code generated by go-swagger; DO NOT EDIT.

package stations

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

// NewGetStationParams creates a new GetStationParams object
// with the default values initialized.
func NewGetStationParams() *GetStationParams {
	var ()
	return &GetStationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStationParamsWithTimeout creates a new GetStationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStationParamsWithTimeout(timeout time.Duration) *GetStationParams {
	var ()
	return &GetStationParams{

		timeout: timeout,
	}
}

// NewGetStationParamsWithContext creates a new GetStationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStationParamsWithContext(ctx context.Context) *GetStationParams {
	var ()
	return &GetStationParams{

		Context: ctx,
	}
}

// NewGetStationParamsWithHTTPClient creates a new GetStationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStationParamsWithHTTPClient(client *http.Client) *GetStationParams {
	var ()
	return &GetStationParams{
		HTTPClient: client,
	}
}

/*GetStationParams contains all the parameters to send to the API endpoint
for the get station operation typically these are written to a http.Request
*/
type GetStationParams struct {

	/*StationID
	  Station ID

	*/
	StationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get station params
func (o *GetStationParams) WithTimeout(timeout time.Duration) *GetStationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get station params
func (o *GetStationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get station params
func (o *GetStationParams) WithContext(ctx context.Context) *GetStationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get station params
func (o *GetStationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get station params
func (o *GetStationParams) WithHTTPClient(client *http.Client) *GetStationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get station params
func (o *GetStationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStationID adds the stationID to the get station params
func (o *GetStationParams) WithStationID(stationID string) *GetStationParams {
	o.SetStationID(stationID)
	return o
}

// SetStationID adds the stationId to the get station params
func (o *GetStationParams) SetStationID(stationID string) {
	o.StationID = stationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stationId
	if err := r.SetPathParam("stationId", o.StationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}