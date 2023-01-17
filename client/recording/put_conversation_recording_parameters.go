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
	"github.com/go-openapi/swag"

	"github.com/freman/genesysapi/models"
)

// NewPutConversationRecordingParams creates a new PutConversationRecordingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutConversationRecordingParams() *PutConversationRecordingParams {
	return &PutConversationRecordingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutConversationRecordingParamsWithTimeout creates a new PutConversationRecordingParams object
// with the ability to set a timeout on a request.
func NewPutConversationRecordingParamsWithTimeout(timeout time.Duration) *PutConversationRecordingParams {
	return &PutConversationRecordingParams{
		timeout: timeout,
	}
}

// NewPutConversationRecordingParamsWithContext creates a new PutConversationRecordingParams object
// with the ability to set a context for a request.
func NewPutConversationRecordingParamsWithContext(ctx context.Context) *PutConversationRecordingParams {
	return &PutConversationRecordingParams{
		Context: ctx,
	}
}

// NewPutConversationRecordingParamsWithHTTPClient creates a new PutConversationRecordingParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutConversationRecordingParamsWithHTTPClient(client *http.Client) *PutConversationRecordingParams {
	return &PutConversationRecordingParams{
		HTTPClient: client,
	}
}

/*
PutConversationRecordingParams contains all the parameters to send to the API endpoint

	for the put conversation recording operation.

	Typically these are written to a http.Request.
*/
type PutConversationRecordingParams struct {

	/* Body.

	   recording
	*/
	Body *models.Recording

	/* ClearExport.

	   Whether to clear the pending export for the recording
	*/
	ClearExport *bool

	/* ConversationID.

	   Conversation ID
	*/
	ConversationID string

	/* RecordingID.

	   Recording ID
	*/
	RecordingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put conversation recording params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutConversationRecordingParams) WithDefaults() *PutConversationRecordingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put conversation recording params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutConversationRecordingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put conversation recording params
func (o *PutConversationRecordingParams) WithTimeout(timeout time.Duration) *PutConversationRecordingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put conversation recording params
func (o *PutConversationRecordingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put conversation recording params
func (o *PutConversationRecordingParams) WithContext(ctx context.Context) *PutConversationRecordingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put conversation recording params
func (o *PutConversationRecordingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put conversation recording params
func (o *PutConversationRecordingParams) WithHTTPClient(client *http.Client) *PutConversationRecordingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put conversation recording params
func (o *PutConversationRecordingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put conversation recording params
func (o *PutConversationRecordingParams) WithBody(body *models.Recording) *PutConversationRecordingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put conversation recording params
func (o *PutConversationRecordingParams) SetBody(body *models.Recording) {
	o.Body = body
}

// WithClearExport adds the clearExport to the put conversation recording params
func (o *PutConversationRecordingParams) WithClearExport(clearExport *bool) *PutConversationRecordingParams {
	o.SetClearExport(clearExport)
	return o
}

// SetClearExport adds the clearExport to the put conversation recording params
func (o *PutConversationRecordingParams) SetClearExport(clearExport *bool) {
	o.ClearExport = clearExport
}

// WithConversationID adds the conversationID to the put conversation recording params
func (o *PutConversationRecordingParams) WithConversationID(conversationID string) *PutConversationRecordingParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the put conversation recording params
func (o *PutConversationRecordingParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithRecordingID adds the recordingID to the put conversation recording params
func (o *PutConversationRecordingParams) WithRecordingID(recordingID string) *PutConversationRecordingParams {
	o.SetRecordingID(recordingID)
	return o
}

// SetRecordingID adds the recordingId to the put conversation recording params
func (o *PutConversationRecordingParams) SetRecordingID(recordingID string) {
	o.RecordingID = recordingID
}

// WriteToRequest writes these params to a swagger request
func (o *PutConversationRecordingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ClearExport != nil {

		// query param clearExport
		var qrClearExport bool

		if o.ClearExport != nil {
			qrClearExport = *o.ClearExport
		}
		qClearExport := swag.FormatBool(qrClearExport)
		if qClearExport != "" {

			if err := r.SetQueryParam("clearExport", qClearExport); err != nil {
				return err
			}
		}
	}

	// path param conversationId
	if err := r.SetPathParam("conversationId", o.ConversationID); err != nil {
		return err
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
