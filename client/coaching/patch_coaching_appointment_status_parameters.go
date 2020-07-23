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

// NewPatchCoachingAppointmentStatusParams creates a new PatchCoachingAppointmentStatusParams object
// with the default values initialized.
func NewPatchCoachingAppointmentStatusParams() *PatchCoachingAppointmentStatusParams {
	var ()
	return &PatchCoachingAppointmentStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCoachingAppointmentStatusParamsWithTimeout creates a new PatchCoachingAppointmentStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCoachingAppointmentStatusParamsWithTimeout(timeout time.Duration) *PatchCoachingAppointmentStatusParams {
	var ()
	return &PatchCoachingAppointmentStatusParams{

		timeout: timeout,
	}
}

// NewPatchCoachingAppointmentStatusParamsWithContext creates a new PatchCoachingAppointmentStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCoachingAppointmentStatusParamsWithContext(ctx context.Context) *PatchCoachingAppointmentStatusParams {
	var ()
	return &PatchCoachingAppointmentStatusParams{

		Context: ctx,
	}
}

// NewPatchCoachingAppointmentStatusParamsWithHTTPClient creates a new PatchCoachingAppointmentStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCoachingAppointmentStatusParamsWithHTTPClient(client *http.Client) *PatchCoachingAppointmentStatusParams {
	var ()
	return &PatchCoachingAppointmentStatusParams{
		HTTPClient: client,
	}
}

/*PatchCoachingAppointmentStatusParams contains all the parameters to send to the API endpoint
for the patch coaching appointment status operation typically these are written to a http.Request
*/
type PatchCoachingAppointmentStatusParams struct {

	/*AppointmentID
	  The ID of the coaching appointment.

	*/
	AppointmentID string
	/*Body
	  Updated status of the coaching appointment

	*/
	Body *models.CoachingAppointmentStatusDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch coaching appointment status params
func (o *PatchCoachingAppointmentStatusParams) WithTimeout(timeout time.Duration) *PatchCoachingAppointmentStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch coaching appointment status params
func (o *PatchCoachingAppointmentStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch coaching appointment status params
func (o *PatchCoachingAppointmentStatusParams) WithContext(ctx context.Context) *PatchCoachingAppointmentStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch coaching appointment status params
func (o *PatchCoachingAppointmentStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch coaching appointment status params
func (o *PatchCoachingAppointmentStatusParams) WithHTTPClient(client *http.Client) *PatchCoachingAppointmentStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch coaching appointment status params
func (o *PatchCoachingAppointmentStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppointmentID adds the appointmentID to the patch coaching appointment status params
func (o *PatchCoachingAppointmentStatusParams) WithAppointmentID(appointmentID string) *PatchCoachingAppointmentStatusParams {
	o.SetAppointmentID(appointmentID)
	return o
}

// SetAppointmentID adds the appointmentId to the patch coaching appointment status params
func (o *PatchCoachingAppointmentStatusParams) SetAppointmentID(appointmentID string) {
	o.AppointmentID = appointmentID
}

// WithBody adds the body to the patch coaching appointment status params
func (o *PatchCoachingAppointmentStatusParams) WithBody(body *models.CoachingAppointmentStatusDto) *PatchCoachingAppointmentStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch coaching appointment status params
func (o *PatchCoachingAppointmentStatusParams) SetBody(body *models.CoachingAppointmentStatusDto) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCoachingAppointmentStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
