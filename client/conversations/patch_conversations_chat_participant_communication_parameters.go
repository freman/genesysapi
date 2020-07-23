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

// NewPatchConversationsChatParticipantCommunicationParams creates a new PatchConversationsChatParticipantCommunicationParams object
// with the default values initialized.
func NewPatchConversationsChatParticipantCommunicationParams() *PatchConversationsChatParticipantCommunicationParams {
	var ()
	return &PatchConversationsChatParticipantCommunicationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConversationsChatParticipantCommunicationParamsWithTimeout creates a new PatchConversationsChatParticipantCommunicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchConversationsChatParticipantCommunicationParamsWithTimeout(timeout time.Duration) *PatchConversationsChatParticipantCommunicationParams {
	var ()
	return &PatchConversationsChatParticipantCommunicationParams{

		timeout: timeout,
	}
}

// NewPatchConversationsChatParticipantCommunicationParamsWithContext creates a new PatchConversationsChatParticipantCommunicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchConversationsChatParticipantCommunicationParamsWithContext(ctx context.Context) *PatchConversationsChatParticipantCommunicationParams {
	var ()
	return &PatchConversationsChatParticipantCommunicationParams{

		Context: ctx,
	}
}

// NewPatchConversationsChatParticipantCommunicationParamsWithHTTPClient creates a new PatchConversationsChatParticipantCommunicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchConversationsChatParticipantCommunicationParamsWithHTTPClient(client *http.Client) *PatchConversationsChatParticipantCommunicationParams {
	var ()
	return &PatchConversationsChatParticipantCommunicationParams{
		HTTPClient: client,
	}
}

/*PatchConversationsChatParticipantCommunicationParams contains all the parameters to send to the API endpoint
for the patch conversations chat participant communication operation typically these are written to a http.Request
*/
type PatchConversationsChatParticipantCommunicationParams struct {

	/*Body
	  Participant

	*/
	Body *models.MediaParticipantRequest
	/*CommunicationID
	  communicationId

	*/
	CommunicationID string
	/*ConversationID
	  conversationId

	*/
	ConversationID string
	/*ParticipantID
	  participantId

	*/
	ParticipantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) WithTimeout(timeout time.Duration) *PatchConversationsChatParticipantCommunicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) WithContext(ctx context.Context) *PatchConversationsChatParticipantCommunicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) WithHTTPClient(client *http.Client) *PatchConversationsChatParticipantCommunicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) WithBody(body *models.MediaParticipantRequest) *PatchConversationsChatParticipantCommunicationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) SetBody(body *models.MediaParticipantRequest) {
	o.Body = body
}

// WithCommunicationID adds the communicationID to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) WithCommunicationID(communicationID string) *PatchConversationsChatParticipantCommunicationParams {
	o.SetCommunicationID(communicationID)
	return o
}

// SetCommunicationID adds the communicationId to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) SetCommunicationID(communicationID string) {
	o.CommunicationID = communicationID
}

// WithConversationID adds the conversationID to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) WithConversationID(conversationID string) *PatchConversationsChatParticipantCommunicationParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) WithParticipantID(participantID string) *PatchConversationsChatParticipantCommunicationParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the patch conversations chat participant communication params
func (o *PatchConversationsChatParticipantCommunicationParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConversationsChatParticipantCommunicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
