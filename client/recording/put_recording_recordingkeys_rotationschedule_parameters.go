// Code generated by go-swagger; DO NOT EDIT.

package recording

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

// NewPutRecordingRecordingkeysRotationscheduleParams creates a new PutRecordingRecordingkeysRotationscheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRecordingRecordingkeysRotationscheduleParams() *PutRecordingRecordingkeysRotationscheduleParams {
	return &PutRecordingRecordingkeysRotationscheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRecordingRecordingkeysRotationscheduleParamsWithTimeout creates a new PutRecordingRecordingkeysRotationscheduleParams object
// with the ability to set a timeout on a request.
func NewPutRecordingRecordingkeysRotationscheduleParamsWithTimeout(timeout time.Duration) *PutRecordingRecordingkeysRotationscheduleParams {
	return &PutRecordingRecordingkeysRotationscheduleParams{
		timeout: timeout,
	}
}

// NewPutRecordingRecordingkeysRotationscheduleParamsWithContext creates a new PutRecordingRecordingkeysRotationscheduleParams object
// with the ability to set a context for a request.
func NewPutRecordingRecordingkeysRotationscheduleParamsWithContext(ctx context.Context) *PutRecordingRecordingkeysRotationscheduleParams {
	return &PutRecordingRecordingkeysRotationscheduleParams{
		Context: ctx,
	}
}

// NewPutRecordingRecordingkeysRotationscheduleParamsWithHTTPClient creates a new PutRecordingRecordingkeysRotationscheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRecordingRecordingkeysRotationscheduleParamsWithHTTPClient(client *http.Client) *PutRecordingRecordingkeysRotationscheduleParams {
	return &PutRecordingRecordingkeysRotationscheduleParams{
		HTTPClient: client,
	}
}

/*
PutRecordingRecordingkeysRotationscheduleParams contains all the parameters to send to the API endpoint

	for the put recording recordingkeys rotationschedule operation.

	Typically these are written to a http.Request.
*/
type PutRecordingRecordingkeysRotationscheduleParams struct {

	/* Body.

	   KeyRotationSchedule
	*/
	Body *models.KeyRotationSchedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put recording recordingkeys rotationschedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRecordingRecordingkeysRotationscheduleParams) WithDefaults() *PutRecordingRecordingkeysRotationscheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put recording recordingkeys rotationschedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRecordingRecordingkeysRotationscheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put recording recordingkeys rotationschedule params
func (o *PutRecordingRecordingkeysRotationscheduleParams) WithTimeout(timeout time.Duration) *PutRecordingRecordingkeysRotationscheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put recording recordingkeys rotationschedule params
func (o *PutRecordingRecordingkeysRotationscheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put recording recordingkeys rotationschedule params
func (o *PutRecordingRecordingkeysRotationscheduleParams) WithContext(ctx context.Context) *PutRecordingRecordingkeysRotationscheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put recording recordingkeys rotationschedule params
func (o *PutRecordingRecordingkeysRotationscheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put recording recordingkeys rotationschedule params
func (o *PutRecordingRecordingkeysRotationscheduleParams) WithHTTPClient(client *http.Client) *PutRecordingRecordingkeysRotationscheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put recording recordingkeys rotationschedule params
func (o *PutRecordingRecordingkeysRotationscheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put recording recordingkeys rotationschedule params
func (o *PutRecordingRecordingkeysRotationscheduleParams) WithBody(body *models.KeyRotationSchedule) *PutRecordingRecordingkeysRotationscheduleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put recording recordingkeys rotationschedule params
func (o *PutRecordingRecordingkeysRotationscheduleParams) SetBody(body *models.KeyRotationSchedule) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutRecordingRecordingkeysRotationscheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
