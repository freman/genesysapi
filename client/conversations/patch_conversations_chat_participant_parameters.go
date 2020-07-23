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

// NewPatchConversationsChatParticipantParams creates a new PatchConversationsChatParticipantParams object
// with the default values initialized.
func NewPatchConversationsChatParticipantParams() *PatchConversationsChatParticipantParams {
	var ()
	return &PatchConversationsChatParticipantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConversationsChatParticipantParamsWithTimeout creates a new PatchConversationsChatParticipantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchConversationsChatParticipantParamsWithTimeout(timeout time.Duration) *PatchConversationsChatParticipantParams {
	var ()
	return &PatchConversationsChatParticipantParams{

		timeout: timeout,
	}
}

// NewPatchConversationsChatParticipantParamsWithContext creates a new PatchConversationsChatParticipantParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchConversationsChatParticipantParamsWithContext(ctx context.Context) *PatchConversationsChatParticipantParams {
	var ()
	return &PatchConversationsChatParticipantParams{

		Context: ctx,
	}
}

// NewPatchConversationsChatParticipantParamsWithHTTPClient creates a new PatchConversationsChatParticipantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchConversationsChatParticipantParamsWithHTTPClient(client *http.Client) *PatchConversationsChatParticipantParams {
	var ()
	return &PatchConversationsChatParticipantParams{
		HTTPClient: client,
	}
}

/*PatchConversationsChatParticipantParams contains all the parameters to send to the API endpoint
for the patch conversations chat participant operation typically these are written to a http.Request
*/
type PatchConversationsChatParticipantParams struct {

	/*Body
	  Update request

	*/
	Body *models.MediaParticipantRequest
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

// WithTimeout adds the timeout to the patch conversations chat participant params
func (o *PatchConversationsChatParticipantParams) WithTimeout(timeout time.Duration) *PatchConversationsChatParticipantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch conversations chat participant params
func (o *PatchConversationsChatParticipantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch conversations chat participant params
func (o *PatchConversationsChatParticipantParams) WithContext(ctx context.Context) *PatchConversationsChatParticipantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch conversations chat participant params
func (o *PatchConversationsChatParticipantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch conversations chat participant params
func (o *PatchConversationsChatParticipantParams) WithHTTPClient(client *http.Client) *PatchConversationsChatParticipantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch conversations chat participant params
func (o *PatchConversationsChatParticipantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch conversations chat participant params
func (o *PatchConversationsChatParticipantParams) WithBody(body *models.MediaParticipantRequest) *PatchConversationsChatParticipantParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch conversations chat participant params
func (o *PatchConversationsChatParticipantParams) SetBody(body *models.MediaParticipantRequest) {
	o.Body = body
}

// WithConversationID adds the conversationID to the patch conversations chat participant params
func (o *PatchConversationsChatParticipantParams) WithConversationID(conversationID string) *PatchConversationsChatParticipantParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch conversations chat participant params
func (o *PatchConversationsChatParticipantParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the patch conversations chat participant params
func (o *PatchConversationsChatParticipantParams) WithParticipantID(participantID string) *PatchConversationsChatParticipantParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the patch conversations chat participant params
func (o *PatchConversationsChatParticipantParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConversationsChatParticipantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
