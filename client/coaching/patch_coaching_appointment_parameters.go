// Code generated by go-swagger; DO NOT EDIT.

package coaching

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

// NewPatchCoachingAppointmentParams creates a new PatchCoachingAppointmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchCoachingAppointmentParams() *PatchCoachingAppointmentParams {
	return &PatchCoachingAppointmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCoachingAppointmentParamsWithTimeout creates a new PatchCoachingAppointmentParams object
// with the ability to set a timeout on a request.
func NewPatchCoachingAppointmentParamsWithTimeout(timeout time.Duration) *PatchCoachingAppointmentParams {
	return &PatchCoachingAppointmentParams{
		timeout: timeout,
	}
}

// NewPatchCoachingAppointmentParamsWithContext creates a new PatchCoachingAppointmentParams object
// with the ability to set a context for a request.
func NewPatchCoachingAppointmentParamsWithContext(ctx context.Context) *PatchCoachingAppointmentParams {
	return &PatchCoachingAppointmentParams{
		Context: ctx,
	}
}

// NewPatchCoachingAppointmentParamsWithHTTPClient creates a new PatchCoachingAppointmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchCoachingAppointmentParamsWithHTTPClient(client *http.Client) *PatchCoachingAppointmentParams {
	return &PatchCoachingAppointmentParams{
		HTTPClient: client,
	}
}

/*
PatchCoachingAppointmentParams contains all the parameters to send to the API endpoint

	for the patch coaching appointment operation.

	Typically these are written to a http.Request.
*/
type PatchCoachingAppointmentParams struct {

	/* AppointmentID.

	   The ID of the coaching appointment.
	*/
	AppointmentID string

	/* Body.

	   The new version of the appointment
	*/
	Body *models.UpdateCoachingAppointmentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch coaching appointment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchCoachingAppointmentParams) WithDefaults() *PatchCoachingAppointmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch coaching appointment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchCoachingAppointmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch coaching appointment params
func (o *PatchCoachingAppointmentParams) WithTimeout(timeout time.Duration) *PatchCoachingAppointmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch coaching appointment params
func (o *PatchCoachingAppointmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch coaching appointment params
func (o *PatchCoachingAppointmentParams) WithContext(ctx context.Context) *PatchCoachingAppointmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch coaching appointment params
func (o *PatchCoachingAppointmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch coaching appointment params
func (o *PatchCoachingAppointmentParams) WithHTTPClient(client *http.Client) *PatchCoachingAppointmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch coaching appointment params
func (o *PatchCoachingAppointmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppointmentID adds the appointmentID to the patch coaching appointment params
func (o *PatchCoachingAppointmentParams) WithAppointmentID(appointmentID string) *PatchCoachingAppointmentParams {
	o.SetAppointmentID(appointmentID)
	return o
}

// SetAppointmentID adds the appointmentId to the patch coaching appointment params
func (o *PatchCoachingAppointmentParams) SetAppointmentID(appointmentID string) {
	o.AppointmentID = appointmentID
}

// WithBody adds the body to the patch coaching appointment params
func (o *PatchCoachingAppointmentParams) WithBody(body *models.UpdateCoachingAppointmentRequest) *PatchCoachingAppointmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch coaching appointment params
func (o *PatchCoachingAppointmentParams) SetBody(body *models.UpdateCoachingAppointmentRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCoachingAppointmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appointmentId
	if err := r.SetPathParam("appointmentId", o.AppointmentID); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
