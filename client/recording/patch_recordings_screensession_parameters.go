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

// NewPatchRecordingsScreensessionParams creates a new PatchRecordingsScreensessionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchRecordingsScreensessionParams() *PatchRecordingsScreensessionParams {
	return &PatchRecordingsScreensessionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchRecordingsScreensessionParamsWithTimeout creates a new PatchRecordingsScreensessionParams object
// with the ability to set a timeout on a request.
func NewPatchRecordingsScreensessionParamsWithTimeout(timeout time.Duration) *PatchRecordingsScreensessionParams {
	return &PatchRecordingsScreensessionParams{
		timeout: timeout,
	}
}

// NewPatchRecordingsScreensessionParamsWithContext creates a new PatchRecordingsScreensessionParams object
// with the ability to set a context for a request.
func NewPatchRecordingsScreensessionParamsWithContext(ctx context.Context) *PatchRecordingsScreensessionParams {
	return &PatchRecordingsScreensessionParams{
		Context: ctx,
	}
}

// NewPatchRecordingsScreensessionParamsWithHTTPClient creates a new PatchRecordingsScreensessionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchRecordingsScreensessionParamsWithHTTPClient(client *http.Client) *PatchRecordingsScreensessionParams {
	return &PatchRecordingsScreensessionParams{
		HTTPClient: client,
	}
}

/*
PatchRecordingsScreensessionParams contains all the parameters to send to the API endpoint

	for the patch recordings screensession operation.

	Typically these are written to a http.Request.
*/
type PatchRecordingsScreensessionParams struct {

	// Body.
	Body *models.ScreenRecordingSessionRequest

	/* RecordingSessionID.

	   Screen recording session ID
	*/
	RecordingSessionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch recordings screensession params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRecordingsScreensessionParams) WithDefaults() *PatchRecordingsScreensessionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch recordings screensession params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchRecordingsScreensessionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch recordings screensession params
func (o *PatchRecordingsScreensessionParams) WithTimeout(timeout time.Duration) *PatchRecordingsScreensessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch recordings screensession params
func (o *PatchRecordingsScreensessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch recordings screensession params
func (o *PatchRecordingsScreensessionParams) WithContext(ctx context.Context) *PatchRecordingsScreensessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch recordings screensession params
func (o *PatchRecordingsScreensessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch recordings screensession params
func (o *PatchRecordingsScreensessionParams) WithHTTPClient(client *http.Client) *PatchRecordingsScreensessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch recordings screensession params
func (o *PatchRecordingsScreensessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch recordings screensession params
func (o *PatchRecordingsScreensessionParams) WithBody(body *models.ScreenRecordingSessionRequest) *PatchRecordingsScreensessionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch recordings screensession params
func (o *PatchRecordingsScreensessionParams) SetBody(body *models.ScreenRecordingSessionRequest) {
	o.Body = body
}

// WithRecordingSessionID adds the recordingSessionID to the patch recordings screensession params
func (o *PatchRecordingsScreensessionParams) WithRecordingSessionID(recordingSessionID string) *PatchRecordingsScreensessionParams {
	o.SetRecordingSessionID(recordingSessionID)
	return o
}

// SetRecordingSessionID adds the recordingSessionId to the patch recordings screensession params
func (o *PatchRecordingsScreensessionParams) SetRecordingSessionID(recordingSessionID string) {
	o.RecordingSessionID = recordingSessionID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchRecordingsScreensessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param recordingSessionId
	if err := r.SetPathParam("recordingSessionId", o.RecordingSessionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
