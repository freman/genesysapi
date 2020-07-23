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
)

// NewGetUserrecordingMediaParams creates a new GetUserrecordingMediaParams object
// with the default values initialized.
func NewGetUserrecordingMediaParams() *GetUserrecordingMediaParams {
	var (
		formatIDDefault = string("WEBM")
	)
	return &GetUserrecordingMediaParams{
		FormatID: &formatIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserrecordingMediaParamsWithTimeout creates a new GetUserrecordingMediaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserrecordingMediaParamsWithTimeout(timeout time.Duration) *GetUserrecordingMediaParams {
	var (
		formatIDDefault = string("WEBM")
	)
	return &GetUserrecordingMediaParams{
		FormatID: &formatIDDefault,

		timeout: timeout,
	}
}

// NewGetUserrecordingMediaParamsWithContext creates a new GetUserrecordingMediaParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserrecordingMediaParamsWithContext(ctx context.Context) *GetUserrecordingMediaParams {
	var (
		formatIdDefault = string("WEBM")
	)
	return &GetUserrecordingMediaParams{
		FormatID: &formatIdDefault,

		Context: ctx,
	}
}

// NewGetUserrecordingMediaParamsWithHTTPClient creates a new GetUserrecordingMediaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserrecordingMediaParamsWithHTTPClient(client *http.Client) *GetUserrecordingMediaParams {
	var (
		formatIdDefault = string("WEBM")
	)
	return &GetUserrecordingMediaParams{
		FormatID:   &formatIdDefault,
		HTTPClient: client,
	}
}

/*GetUserrecordingMediaParams contains all the parameters to send to the API endpoint
for the get userrecording media operation typically these are written to a http.Request
*/
type GetUserrecordingMediaParams struct {

	/*FormatID
	  The desired media format.

	*/
	FormatID *string
	/*RecordingID
	  User Recording ID

	*/
	RecordingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get userrecording media params
func (o *GetUserrecordingMediaParams) WithTimeout(timeout time.Duration) *GetUserrecordingMediaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get userrecording media params
func (o *GetUserrecordingMediaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get userrecording media params
func (o *GetUserrecordingMediaParams) WithContext(ctx context.Context) *GetUserrecordingMediaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get userrecording media params
func (o *GetUserrecordingMediaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get userrecording media params
func (o *GetUserrecordingMediaParams) WithHTTPClient(client *http.Client) *GetUserrecordingMediaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get userrecording media params
func (o *GetUserrecordingMediaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormatID adds the formatID to the get userrecording media params
func (o *GetUserrecordingMediaParams) WithFormatID(formatID *string) *GetUserrecordingMediaParams {
	o.SetFormatID(formatID)
	return o
}

// SetFormatID adds the formatId to the get userrecording media params
func (o *GetUserrecordingMediaParams) SetFormatID(formatID *string) {
	o.FormatID = formatID
}

// WithRecordingID adds the recordingID to the get userrecording media params
func (o *GetUserrecordingMediaParams) WithRecordingID(recordingID string) *GetUserrecordingMediaParams {
	o.SetRecordingID(recordingID)
	return o
}

// SetRecordingID adds the recordingId to the get userrecording media params
func (o *GetUserrecordingMediaParams) SetRecordingID(recordingID string) {
	o.RecordingID = recordingID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserrecordingMediaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FormatID != nil {

		// query param formatId
		var qrFormatID string
		if o.FormatID != nil {
			qrFormatID = *o.FormatID
		}
		qFormatID := qrFormatID
		if qFormatID != "" {
			if err := r.SetQueryParam("formatId", qFormatID); err != nil {
				return err
			}
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
