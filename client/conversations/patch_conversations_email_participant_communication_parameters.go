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

// NewPatchConversationsEmailParticipantCommunicationParams creates a new PatchConversationsEmailParticipantCommunicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchConversationsEmailParticipantCommunicationParams() *PatchConversationsEmailParticipantCommunicationParams {
	return &PatchConversationsEmailParticipantCommunicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConversationsEmailParticipantCommunicationParamsWithTimeout creates a new PatchConversationsEmailParticipantCommunicationParams object
// with the ability to set a timeout on a request.
func NewPatchConversationsEmailParticipantCommunicationParamsWithTimeout(timeout time.Duration) *PatchConversationsEmailParticipantCommunicationParams {
	return &PatchConversationsEmailParticipantCommunicationParams{
		timeout: timeout,
	}
}

// NewPatchConversationsEmailParticipantCommunicationParamsWithContext creates a new PatchConversationsEmailParticipantCommunicationParams object
// with the ability to set a context for a request.
func NewPatchConversationsEmailParticipantCommunicationParamsWithContext(ctx context.Context) *PatchConversationsEmailParticipantCommunicationParams {
	return &PatchConversationsEmailParticipantCommunicationParams{
		Context: ctx,
	}
}

// NewPatchConversationsEmailParticipantCommunicationParamsWithHTTPClient creates a new PatchConversationsEmailParticipantCommunicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchConversationsEmailParticipantCommunicationParamsWithHTTPClient(client *http.Client) *PatchConversationsEmailParticipantCommunicationParams {
	return &PatchConversationsEmailParticipantCommunicationParams{
		HTTPClient: client,
	}
}

/*
PatchConversationsEmailParticipantCommunicationParams contains all the parameters to send to the API endpoint

	for the patch conversations email participant communication operation.

	Typically these are written to a http.Request.
*/
type PatchConversationsEmailParticipantCommunicationParams struct {

	/* Body.

	   Participant
	*/
	Body *models.MediaParticipantRequest

	/* CommunicationID.

	   communicationId
	*/
	CommunicationID string

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

// WithDefaults hydrates default values in the patch conversations email participant communication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchConversationsEmailParticipantCommunicationParams) WithDefaults() *PatchConversationsEmailParticipantCommunicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch conversations email participant communication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchConversationsEmailParticipantCommunicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) WithTimeout(timeout time.Duration) *PatchConversationsEmailParticipantCommunicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) WithContext(ctx context.Context) *PatchConversationsEmailParticipantCommunicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) WithHTTPClient(client *http.Client) *PatchConversationsEmailParticipantCommunicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) WithBody(body *models.MediaParticipantRequest) *PatchConversationsEmailParticipantCommunicationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) SetBody(body *models.MediaParticipantRequest) {
	o.Body = body
}

// WithCommunicationID adds the communicationID to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) WithCommunicationID(communicationID string) *PatchConversationsEmailParticipantCommunicationParams {
	o.SetCommunicationID(communicationID)
	return o
}

// SetCommunicationID adds the communicationId to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) SetCommunicationID(communicationID string) {
	o.CommunicationID = communicationID
}

// WithConversationID adds the conversationID to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) WithConversationID(conversationID string) *PatchConversationsEmailParticipantCommunicationParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) WithParticipantID(participantID string) *PatchConversationsEmailParticipantCommunicationParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the patch conversations email participant communication params
func (o *PatchConversationsEmailParticipantCommunicationParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConversationsEmailParticipantCommunicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param communicationId
	if err := r.SetPathParam("communicationId", o.CommunicationID); err != nil {
		return err
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
