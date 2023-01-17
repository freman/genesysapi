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
)

// NewGetCoachingAppointmentAnnotationParams creates a new GetCoachingAppointmentAnnotationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCoachingAppointmentAnnotationParams() *GetCoachingAppointmentAnnotationParams {
	return &GetCoachingAppointmentAnnotationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCoachingAppointmentAnnotationParamsWithTimeout creates a new GetCoachingAppointmentAnnotationParams object
// with the ability to set a timeout on a request.
func NewGetCoachingAppointmentAnnotationParamsWithTimeout(timeout time.Duration) *GetCoachingAppointmentAnnotationParams {
	return &GetCoachingAppointmentAnnotationParams{
		timeout: timeout,
	}
}

// NewGetCoachingAppointmentAnnotationParamsWithContext creates a new GetCoachingAppointmentAnnotationParams object
// with the ability to set a context for a request.
func NewGetCoachingAppointmentAnnotationParamsWithContext(ctx context.Context) *GetCoachingAppointmentAnnotationParams {
	return &GetCoachingAppointmentAnnotationParams{
		Context: ctx,
	}
}

// NewGetCoachingAppointmentAnnotationParamsWithHTTPClient creates a new GetCoachingAppointmentAnnotationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCoachingAppointmentAnnotationParamsWithHTTPClient(client *http.Client) *GetCoachingAppointmentAnnotationParams {
	return &GetCoachingAppointmentAnnotationParams{
		HTTPClient: client,
	}
}

/*
GetCoachingAppointmentAnnotationParams contains all the parameters to send to the API endpoint

	for the get coaching appointment annotation operation.

	Typically these are written to a http.Request.
*/
type GetCoachingAppointmentAnnotationParams struct {

	/* AnnotationID.

	   The ID of the annotation.
	*/
	AnnotationID string

	/* AppointmentID.

	   The ID of the coaching appointment.
	*/
	AppointmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get coaching appointment annotation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoachingAppointmentAnnotationParams) WithDefaults() *GetCoachingAppointmentAnnotationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get coaching appointment annotation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoachingAppointmentAnnotationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get coaching appointment annotation params
func (o *GetCoachingAppointmentAnnotationParams) WithTimeout(timeout time.Duration) *GetCoachingAppointmentAnnotationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get coaching appointment annotation params
func (o *GetCoachingAppointmentAnnotationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get coaching appointment annotation params
func (o *GetCoachingAppointmentAnnotationParams) WithContext(ctx context.Context) *GetCoachingAppointmentAnnotationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get coaching appointment annotation params
func (o *GetCoachingAppointmentAnnotationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get coaching appointment annotation params
func (o *GetCoachingAppointmentAnnotationParams) WithHTTPClient(client *http.Client) *GetCoachingAppointmentAnnotationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get coaching appointment annotation params
func (o *GetCoachingAppointmentAnnotationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAnnotationID adds the annotationID to the get coaching appointment annotation params
func (o *GetCoachingAppointmentAnnotationParams) WithAnnotationID(annotationID string) *GetCoachingAppointmentAnnotationParams {
	o.SetAnnotationID(annotationID)
	return o
}

// SetAnnotationID adds the annotationId to the get coaching appointment annotation params
func (o *GetCoachingAppointmentAnnotationParams) SetAnnotationID(annotationID string) {
	o.AnnotationID = annotationID
}

// WithAppointmentID adds the appointmentID to the get coaching appointment annotation params
func (o *GetCoachingAppointmentAnnotationParams) WithAppointmentID(appointmentID string) *GetCoachingAppointmentAnnotationParams {
	o.SetAppointmentID(appointmentID)
	return o
}

// SetAppointmentID adds the appointmentId to the get coaching appointment annotation params
func (o *GetCoachingAppointmentAnnotationParams) SetAppointmentID(appointmentID string) {
	o.AppointmentID = appointmentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCoachingAppointmentAnnotationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param annotationId
	if err := r.SetPathParam("annotationId", o.AnnotationID); err != nil {
		return err
	}

	// path param appointmentId
	if err := r.SetPathParam("appointmentId", o.AppointmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
