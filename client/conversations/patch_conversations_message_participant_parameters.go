// Code generated by go-swagger; DO NOT EDIT.

package conversations

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

// NewPatchConversationsMessageParticipantParams creates a new PatchConversationsMessageParticipantParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchConversationsMessageParticipantParams() *PatchConversationsMessageParticipantParams {
	return &PatchConversationsMessageParticipantParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConversationsMessageParticipantParamsWithTimeout creates a new PatchConversationsMessageParticipantParams object
// with the ability to set a timeout on a request.
func NewPatchConversationsMessageParticipantParamsWithTimeout(timeout time.Duration) *PatchConversationsMessageParticipantParams {
	return &PatchConversationsMessageParticipantParams{
		timeout: timeout,
	}
}

// NewPatchConversationsMessageParticipantParamsWithContext creates a new PatchConversationsMessageParticipantParams object
// with the ability to set a context for a request.
func NewPatchConversationsMessageParticipantParamsWithContext(ctx context.Context) *PatchConversationsMessageParticipantParams {
	return &PatchConversationsMessageParticipantParams{
		Context: ctx,
	}
}

// NewPatchConversationsMessageParticipantParamsWithHTTPClient creates a new PatchConversationsMessageParticipantParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchConversationsMessageParticipantParamsWithHTTPClient(client *http.Client) *PatchConversationsMessageParticipantParams {
	return &PatchConversationsMessageParticipantParams{
		HTTPClient: client,
	}
}

/*
PatchConversationsMessageParticipantParams contains all the parameters to send to the API endpoint

	for the patch conversations message participant operation.

	Typically these are written to a http.Request.
*/
type PatchConversationsMessageParticipantParams struct {

	// Body.
	Body *models.MediaParticipantRequest

	/* ConversationID.

	   conversationId
	*/
	ConversationID string

	/* ParticipantID.

	   participantId
	*/
	ParticipantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch conversations message participant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchConversationsMessageParticipantParams) WithDefaults() *PatchConversationsMessageParticipantParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch conversations message participant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchConversationsMessageParticipantParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch conversations message participant params
func (o *PatchConversationsMessageParticipantParams) WithTimeout(timeout time.Duration) *PatchConversationsMessageParticipantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch conversations message participant params
func (o *PatchConversationsMessageParticipantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch conversations message participant params
func (o *PatchConversationsMessageParticipantParams) WithContext(ctx context.Context) *PatchConversationsMessageParticipantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch conversations message participant params
func (o *PatchConversationsMessageParticipantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch conversations message participant params
func (o *PatchConversationsMessageParticipantParams) WithHTTPClient(client *http.Client) *PatchConversationsMessageParticipantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch conversations message participant params
func (o *PatchConversationsMessageParticipantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch conversations message participant params
func (o *PatchConversationsMessageParticipantParams) WithBody(body *models.MediaParticipantRequest) *PatchConversationsMessageParticipantParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch conversations message participant params
func (o *PatchConversationsMessageParticipantParams) SetBody(body *models.MediaParticipantRequest) {
	o.Body = body
}

// WithConversationID adds the conversationID to the patch conversations message participant params
func (o *PatchConversationsMessageParticipantParams) WithConversationID(conversationID string) *PatchConversationsMessageParticipantParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch conversations message participant params
func (o *PatchConversationsMessageParticipantParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the patch conversations message participant params
func (o *PatchConversationsMessageParticipantParams) WithParticipantID(participantID string) *PatchConversationsMessageParticipantParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the patch conversations message participant params
func (o *PatchConversationsMessageParticipantParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConversationsMessageParticipantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param conversationId
	if err := r.SetPathParam("conversationId", o.ConversationID); err != nil {
		return err
	}

	// path param participantId
	if err := r.SetPathParam("participantId", o.ParticipantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
