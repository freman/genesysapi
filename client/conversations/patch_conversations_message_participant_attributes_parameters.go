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

// NewPatchConversationsMessageParticipantAttributesParams creates a new PatchConversationsMessageParticipantAttributesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchConversationsMessageParticipantAttributesParams() *PatchConversationsMessageParticipantAttributesParams {
	return &PatchConversationsMessageParticipantAttributesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConversationsMessageParticipantAttributesParamsWithTimeout creates a new PatchConversationsMessageParticipantAttributesParams object
// with the ability to set a timeout on a request.
func NewPatchConversationsMessageParticipantAttributesParamsWithTimeout(timeout time.Duration) *PatchConversationsMessageParticipantAttributesParams {
	return &PatchConversationsMessageParticipantAttributesParams{
		timeout: timeout,
	}
}

// NewPatchConversationsMessageParticipantAttributesParamsWithContext creates a new PatchConversationsMessageParticipantAttributesParams object
// with the ability to set a context for a request.
func NewPatchConversationsMessageParticipantAttributesParamsWithContext(ctx context.Context) *PatchConversationsMessageParticipantAttributesParams {
	return &PatchConversationsMessageParticipantAttributesParams{
		Context: ctx,
	}
}

// NewPatchConversationsMessageParticipantAttributesParamsWithHTTPClient creates a new PatchConversationsMessageParticipantAttributesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchConversationsMessageParticipantAttributesParamsWithHTTPClient(client *http.Client) *PatchConversationsMessageParticipantAttributesParams {
	return &PatchConversationsMessageParticipantAttributesParams{
		HTTPClient: client,
	}
}

/*
PatchConversationsMessageParticipantAttributesParams contains all the parameters to send to the API endpoint

	for the patch conversations message participant attributes operation.

	Typically these are written to a http.Request.
*/
type PatchConversationsMessageParticipantAttributesParams struct {

	// Body.
	Body *models.ParticipantAttributes

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

// WithDefaults hydrates default values in the patch conversations message participant attributes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchConversationsMessageParticipantAttributesParams) WithDefaults() *PatchConversationsMessageParticipantAttributesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch conversations message participant attributes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchConversationsMessageParticipantAttributesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch conversations message participant attributes params
func (o *PatchConversationsMessageParticipantAttributesParams) WithTimeout(timeout time.Duration) *PatchConversationsMessageParticipantAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch conversations message participant attributes params
func (o *PatchConversationsMessageParticipantAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch conversations message participant attributes params
func (o *PatchConversationsMessageParticipantAttributesParams) WithContext(ctx context.Context) *PatchConversationsMessageParticipantAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch conversations message participant attributes params
func (o *PatchConversationsMessageParticipantAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch conversations message participant attributes params
func (o *PatchConversationsMessageParticipantAttributesParams) WithHTTPClient(client *http.Client) *PatchConversationsMessageParticipantAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch conversations message participant attributes params
func (o *PatchConversationsMessageParticipantAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch conversations message participant attributes params
func (o *PatchConversationsMessageParticipantAttributesParams) WithBody(body *models.ParticipantAttributes) *PatchConversationsMessageParticipantAttributesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch conversations message participant attributes params
func (o *PatchConversationsMessageParticipantAttributesParams) SetBody(body *models.ParticipantAttributes) {
	o.Body = body
}

// WithConversationID adds the conversationID to the patch conversations message participant attributes params
func (o *PatchConversationsMessageParticipantAttributesParams) WithConversationID(conversationID string) *PatchConversationsMessageParticipantAttributesParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch conversations message participant attributes params
func (o *PatchConversationsMessageParticipantAttributesParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the patch conversations message participant attributes params
func (o *PatchConversationsMessageParticipantAttributesParams) WithParticipantID(participantID string) *PatchConversationsMessageParticipantAttributesParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the patch conversations message participant attributes params
func (o *PatchConversationsMessageParticipantAttributesParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConversationsMessageParticipantAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
