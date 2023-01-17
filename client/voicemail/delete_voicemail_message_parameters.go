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
)

// NewDeleteVoicemailMessageParams creates a new DeleteVoicemailMessageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVoicemailMessageParams() *DeleteVoicemailMessageParams {
	return &DeleteVoicemailMessageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVoicemailMessageParamsWithTimeout creates a new DeleteVoicemailMessageParams object
// with the ability to set a timeout on a request.
func NewDeleteVoicemailMessageParamsWithTimeout(timeout time.Duration) *DeleteVoicemailMessageParams {
	return &DeleteVoicemailMessageParams{
		timeout: timeout,
	}
}

// NewDeleteVoicemailMessageParamsWithContext creates a new DeleteVoicemailMessageParams object
// with the ability to set a context for a request.
func NewDeleteVoicemailMessageParamsWithContext(ctx context.Context) *DeleteVoicemailMessageParams {
	return &DeleteVoicemailMessageParams{
		Context: ctx,
	}
}

// NewDeleteVoicemailMessageParamsWithHTTPClient creates a new DeleteVoicemailMessageParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVoicemailMessageParamsWithHTTPClient(client *http.Client) *DeleteVoicemailMessageParams {
	return &DeleteVoicemailMessageParams{
		HTTPClient: client,
	}
}

/*
DeleteVoicemailMessageParams contains all the parameters to send to the API endpoint

	for the delete voicemail message operation.

	Typically these are written to a http.Request.
*/
type DeleteVoicemailMessageParams struct {

	/* MessageID.

	   Message ID
	*/
	MessageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete voicemail message params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVoicemailMessageParams) WithDefaults() *DeleteVoicemailMessageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete voicemail message params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVoicemailMessageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete voicemail message params
func (o *DeleteVoicemailMessageParams) WithTimeout(timeout time.Duration) *DeleteVoicemailMessageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete voicemail message params
func (o *DeleteVoicemailMessageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete voicemail message params
func (o *DeleteVoicemailMessageParams) WithContext(ctx context.Context) *DeleteVoicemailMessageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete voicemail message params
func (o *DeleteVoicemailMessageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete voicemail message params
func (o *DeleteVoicemailMessageParams) WithHTTPClient(client *http.Client) *DeleteVoicemailMessageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete voicemail message params
func (o *DeleteVoicemailMessageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessageID adds the messageID to the delete voicemail message params
func (o *DeleteVoicemailMessageParams) WithMessageID(messageID string) *DeleteVoicemailMessageParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the delete voicemail message params
func (o *DeleteVoicemailMessageParams) SetMessageID(messageID string) {
	o.MessageID = messageID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVoicemailMessageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param messageId
	if err := r.SetPathParam("messageId", o.MessageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
