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

// NewDeleteMobiledeviceParams creates a new DeleteMobiledeviceParams object
// with the default values initialized.
func NewDeleteMobiledeviceParams() *DeleteMobiledeviceParams {
	var ()
	return &DeleteMobiledeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMobiledeviceParamsWithTimeout creates a new DeleteMobiledeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMobiledeviceParamsWithTimeout(timeout time.Duration) *DeleteMobiledeviceParams {
	var ()
	return &DeleteMobiledeviceParams{

		timeout: timeout,
	}
}

// NewDeleteMobiledeviceParamsWithContext creates a new DeleteMobiledeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMobiledeviceParamsWithContext(ctx context.Context) *DeleteMobiledeviceParams {
	var ()
	return &DeleteMobiledeviceParams{

		Context: ctx,
	}
}

// NewDeleteMobiledeviceParamsWithHTTPClient creates a new DeleteMobiledeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMobiledeviceParamsWithHTTPClient(client *http.Client) *DeleteMobiledeviceParams {
	var ()
	return &DeleteMobiledeviceParams{
		HTTPClient: client,
	}
}

/*DeleteMobiledeviceParams contains all the parameters to send to the API endpoint
for the delete mobiledevice operation typically these are written to a http.Request
*/
type DeleteMobiledeviceParams struct {

	/*DeviceID
	  Device ID

	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete mobiledevice params
func (o *DeleteMobiledeviceParams) WithTimeout(timeout time.Duration) *DeleteMobiledeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete mobiledevice params
func (o *DeleteMobiledeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete mobiledevice params
func (o *DeleteMobiledeviceParams) WithContext(ctx context.Context) *DeleteMobiledeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete mobiledevice params
func (o *DeleteMobiledeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete mobiledevice params
func (o *DeleteMobiledeviceParams) WithHTTPClient(client *http.Client) *DeleteMobiledeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete mobiledevice params
func (o *DeleteMobiledeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the delete mobiledevice params
func (o *DeleteMobiledeviceParams) WithDeviceID(deviceID string) *DeleteMobiledeviceParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the delete mobiledevice params
func (o *DeleteMobiledeviceParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMobiledeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
