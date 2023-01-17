// Code generated by go-swagger; DO NOT EDIT.

package user_recordings

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
	"github.com/go-openapi/swag"

	"github.com/freman/genesysapi/models"
)

// NewPutUserrecordingParams creates a new PutUserrecordingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutUserrecordingParams() *PutUserrecordingParams {
	return &PutUserrecordingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutUserrecordingParamsWithTimeout creates a new PutUserrecordingParams object
// with the ability to set a timeout on a request.
func NewPutUserrecordingParamsWithTimeout(timeout time.Duration) *PutUserrecordingParams {
	return &PutUserrecordingParams{
		timeout: timeout,
	}
}

// NewPutUserrecordingParamsWithContext creates a new PutUserrecordingParams object
// with the ability to set a context for a request.
func NewPutUserrecordingParamsWithContext(ctx context.Context) *PutUserrecordingParams {
	return &PutUserrecordingParams{
		Context: ctx,
	}
}

// NewPutUserrecordingParamsWithHTTPClient creates a new PutUserrecordingParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutUserrecordingParamsWithHTTPClient(client *http.Client) *PutUserrecordingParams {
	return &PutUserrecordingParams{
		HTTPClient: client,
	}
}

/*
PutUserrecordingParams contains all the parameters to send to the API endpoint

	for the put userrecording operation.

	Typically these are written to a http.Request.
*/
type PutUserrecordingParams struct {

	/* Body.

	   UserRecording
	*/
	Body *models.UserRecording

	/* Expand.

	   Which fields, if any, to expand.
	*/
	Expand []string

	/* RecordingID.

	   User Recording ID
	*/
	RecordingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put userrecording params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutUserrecordingParams) WithDefaults() *PutUserrecordingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put userrecording params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutUserrecordingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put userrecording params
func (o *PutUserrecordingParams) WithTimeout(timeout time.Duration) *PutUserrecordingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put userrecording params
func (o *PutUserrecordingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put userrecording params
func (o *PutUserrecordingParams) WithContext(ctx context.Context) *PutUserrecordingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put userrecording params
func (o *PutUserrecordingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put userrecording params
func (o *PutUserrecordingParams) WithHTTPClient(client *http.Client) *PutUserrecordingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put userrecording params
func (o *PutUserrecordingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put userrecording params
func (o *PutUserrecordingParams) WithBody(body *models.UserRecording) *PutUserrecordingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put userrecording params
func (o *PutUserrecordingParams) SetBody(body *models.UserRecording) {
	o.Body = body
}

// WithExpand adds the expand to the put userrecording params
func (o *PutUserrecordingParams) WithExpand(expand []string) *PutUserrecordingParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the put userrecording params
func (o *PutUserrecordingParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithRecordingID adds the recordingID to the put userrecording params
func (o *PutUserrecordingParams) WithRecordingID(recordingID string) *PutUserrecordingParams {
	o.SetRecordingID(recordingID)
	return o
}

// SetRecordingID adds the recordingId to the put userrecording params
func (o *PutUserrecordingParams) SetRecordingID(recordingID string) {
	o.RecordingID = recordingID
}

// WriteToRequest writes these params to a swagger request
func (o *PutUserrecordingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	// path param recordingId
	if err := r.SetPathParam("recordingId", o.RecordingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPutUserrecording binds the parameter expand
func (o *PutUserrecordingParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}
