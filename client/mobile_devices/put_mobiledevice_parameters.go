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

	"github.com/freman/genesysapi/models"
)

// NewPutMobiledeviceParams creates a new PutMobiledeviceParams object
// with the default values initialized.
func NewPutMobiledeviceParams() *PutMobiledeviceParams {
	var ()
	return &PutMobiledeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutMobiledeviceParamsWithTimeout creates a new PutMobiledeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutMobiledeviceParamsWithTimeout(timeout time.Duration) *PutMobiledeviceParams {
	var ()
	return &PutMobiledeviceParams{

		timeout: timeout,
	}
}

// NewPutMobiledeviceParamsWithContext creates a new PutMobiledeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutMobiledeviceParamsWithContext(ctx context.Context) *PutMobiledeviceParams {
	var ()
	return &PutMobiledeviceParams{

		Context: ctx,
	}
}

// NewPutMobiledeviceParamsWithHTTPClient creates a new PutMobiledeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutMobiledeviceParamsWithHTTPClient(client *http.Client) *PutMobiledeviceParams {
	var ()
	return &PutMobiledeviceParams{
		HTTPClient: client,
	}
}

/*PutMobiledeviceParams contains all the parameters to send to the API endpoint
for the put mobiledevice operation typically these are written to a http.Request
*/
type PutMobiledeviceParams struct {

	/*Body
	  Device

	*/
	Body *models.UserDevice
	/*DeviceID
	  Device ID

	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put mobiledevice params
func (o *PutMobiledeviceParams) WithTimeout(timeout time.Duration) *PutMobiledeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put mobiledevice params
func (o *PutMobiledeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put mobiledevice params
func (o *PutMobiledeviceParams) WithContext(ctx context.Context) *PutMobiledeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put mobiledevice params
func (o *PutMobiledeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put mobiledevice params
func (o *PutMobiledeviceParams) WithHTTPClient(client *http.Client) *PutMobiledeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put mobiledevice params
func (o *PutMobiledeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put mobiledevice params
func (o *PutMobiledeviceParams) WithBody(body *models.UserDevice) *PutMobiledeviceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put mobiledevice params
func (o *PutMobiledeviceParams) SetBody(body *models.UserDevice) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the put mobiledevice params
func (o *PutMobiledeviceParams) WithDeviceID(deviceID string) *PutMobiledeviceParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the put mobiledevice params
func (o *PutMobiledeviceParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *PutMobiledeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}