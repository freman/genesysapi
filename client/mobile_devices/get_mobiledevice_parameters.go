// Code generated by go-swagger; DO NOT EDIT.

package mobile_devices

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

// NewGetMobiledeviceParams creates a new GetMobiledeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMobiledeviceParams() *GetMobiledeviceParams {
	return &GetMobiledeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMobiledeviceParamsWithTimeout creates a new GetMobiledeviceParams object
// with the ability to set a timeout on a request.
func NewGetMobiledeviceParamsWithTimeout(timeout time.Duration) *GetMobiledeviceParams {
	return &GetMobiledeviceParams{
		timeout: timeout,
	}
}

// NewGetMobiledeviceParamsWithContext creates a new GetMobiledeviceParams object
// with the ability to set a context for a request.
func NewGetMobiledeviceParamsWithContext(ctx context.Context) *GetMobiledeviceParams {
	return &GetMobiledeviceParams{
		Context: ctx,
	}
}

// NewGetMobiledeviceParamsWithHTTPClient creates a new GetMobiledeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMobiledeviceParamsWithHTTPClient(client *http.Client) *GetMobiledeviceParams {
	return &GetMobiledeviceParams{
		HTTPClient: client,
	}
}

/*
GetMobiledeviceParams contains all the parameters to send to the API endpoint

	for the get mobiledevice operation.

	Typically these are written to a http.Request.
*/
type GetMobiledeviceParams struct {

	/* DeviceID.

	   Device ID
	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get mobiledevice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMobiledeviceParams) WithDefaults() *GetMobiledeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get mobiledevice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMobiledeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get mobiledevice params
func (o *GetMobiledeviceParams) WithTimeout(timeout time.Duration) *GetMobiledeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mobiledevice params
func (o *GetMobiledeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mobiledevice params
func (o *GetMobiledeviceParams) WithContext(ctx context.Context) *GetMobiledeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mobiledevice params
func (o *GetMobiledeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mobiledevice params
func (o *GetMobiledeviceParams) WithHTTPClient(client *http.Client) *GetMobiledeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mobiledevice params
func (o *GetMobiledeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get mobiledevice params
func (o *GetMobiledeviceParams) WithDeviceID(deviceID string) *GetMobiledeviceParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get mobiledevice params
func (o *GetMobiledeviceParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMobiledeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
