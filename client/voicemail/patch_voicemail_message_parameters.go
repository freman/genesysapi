// Code generated by go-swagger; DO NOT EDIT.

package voicemail

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

// NewPatchVoicemailMessageParams creates a new PatchVoicemailMessageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchVoicemailMessageParams() *PatchVoicemailMessageParams {
	return &PatchVoicemailMessageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchVoicemailMessageParamsWithTimeout creates a new PatchVoicemailMessageParams object
// with the ability to set a timeout on a request.
func NewPatchVoicemailMessageParamsWithTimeout(timeout time.Duration) *PatchVoicemailMessageParams {
	return &PatchVoicemailMessageParams{
		timeout: timeout,
	}
}

// NewPatchVoicemailMessageParamsWithContext creates a new PatchVoicemailMessageParams object
// with the ability to set a context for a request.
func NewPatchVoicemailMessageParamsWithContext(ctx context.Context) *PatchVoicemailMessageParams {
	return &PatchVoicemailMessageParams{
		Context: ctx,
	}
}

// NewPatchVoicemailMessageParamsWithHTTPClient creates a new PatchVoicemailMessageParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchVoicemailMessageParamsWithHTTPClient(client *http.Client) *PatchVoicemailMessageParams {
	return &PatchVoicemailMessageParams{
		HTTPClient: client,
	}
}

/*
PatchVoicemailMessageParams contains all the parameters to send to the API endpoint

	for the patch voicemail message operation.

	Typically these are written to a http.Request.
*/
type PatchVoicemailMessageParams struct {

	/* Body.

	   VoicemailMessage
	*/
	Body *models.VoicemailMessage

	/* MessageID.

	   Message ID
	*/
	MessageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch voicemail message params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchVoicemailMessageParams) WithDefaults() *PatchVoicemailMessageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch voicemail message params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchVoicemailMessageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch voicemail message params
func (o *PatchVoicemailMessageParams) WithTimeout(timeout time.Duration) *PatchVoicemailMessageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch voicemail message params
func (o *PatchVoicemailMessageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch voicemail message params
func (o *PatchVoicemailMessageParams) WithContext(ctx context.Context) *PatchVoicemailMessageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch voicemail message params
func (o *PatchVoicemailMessageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch voicemail message params
func (o *PatchVoicemailMessageParams) WithHTTPClient(client *http.Client) *PatchVoicemailMessageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch voicemail message params
func (o *PatchVoicemailMessageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch voicemail message params
func (o *PatchVoicemailMessageParams) WithBody(body *models.VoicemailMessage) *PatchVoicemailMessageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch voicemail message params
func (o *PatchVoicemailMessageParams) SetBody(body *models.VoicemailMessage) {
	o.Body = body
}

// WithMessageID adds the messageID to the patch voicemail message params
func (o *PatchVoicemailMessageParams) WithMessageID(messageID string) *PatchVoicemailMessageParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the patch voicemail message params
func (o *PatchVoicemailMessageParams) SetMessageID(messageID string) {
	o.MessageID = messageID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchVoicemailMessageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param messageId
	if err := r.SetPathParam("messageId", o.MessageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
