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
)

// NewGetConversationsEmailParticipantWrapupcodesParams creates a new GetConversationsEmailParticipantWrapupcodesParams object
// with the default values initialized.
func NewGetConversationsEmailParticipantWrapupcodesParams() *GetConversationsEmailParticipantWrapupcodesParams {
	var ()
	return &GetConversationsEmailParticipantWrapupcodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsEmailParticipantWrapupcodesParamsWithTimeout creates a new GetConversationsEmailParticipantWrapupcodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationsEmailParticipantWrapupcodesParamsWithTimeout(timeout time.Duration) *GetConversationsEmailParticipantWrapupcodesParams {
	var ()
	return &GetConversationsEmailParticipantWrapupcodesParams{

		timeout: timeout,
	}
}

// NewGetConversationsEmailParticipantWrapupcodesParamsWithContext creates a new GetConversationsEmailParticipantWrapupcodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationsEmailParticipantWrapupcodesParamsWithContext(ctx context.Context) *GetConversationsEmailParticipantWrapupcodesParams {
	var ()
	return &GetConversationsEmailParticipantWrapupcodesParams{

		Context: ctx,
	}
}

// NewGetConversationsEmailParticipantWrapupcodesParamsWithHTTPClient creates a new GetConversationsEmailParticipantWrapupcodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationsEmailParticipantWrapupcodesParamsWithHTTPClient(client *http.Client) *GetConversationsEmailParticipantWrapupcodesParams {
	var ()
	return &GetConversationsEmailParticipantWrapupcodesParams{
		HTTPClient: client,
	}
}

/*GetConversationsEmailParticipantWrapupcodesParams contains all the parameters to send to the API endpoint
for the get conversations email participant wrapupcodes operation typically these are written to a http.Request
*/
type GetConversationsEmailParticipantWrapupcodesParams struct {

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

// WithTimeout adds the timeout to the get conversations email participant wrapupcodes params
func (o *GetConversationsEmailParticipantWrapupcodesParams) WithTimeout(timeout time.Duration) *GetConversationsEmailParticipantWrapupcodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations email participant wrapupcodes params
func (o *GetConversationsEmailParticipantWrapupcodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations email participant wrapupcodes params
func (o *GetConversationsEmailParticipantWrapupcodesParams) WithContext(ctx context.Context) *GetConversationsEmailParticipantWrapupcodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations email participant wrapupcodes params
func (o *GetConversationsEmailParticipantWrapupcodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations email participant wrapupcodes params
func (o *GetConversationsEmailParticipantWrapupcodesParams) WithHTTPClient(client *http.Client) *GetConversationsEmailParticipantWrapupcodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations email participant wrapupcodes params
func (o *GetConversationsEmailParticipantWrapupcodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the get conversations email participant wrapupcodes params
func (o *GetConversationsEmailParticipantWrapupcodesParams) WithConversationID(conversationID string) *GetConversationsEmailParticipantWrapupcodesParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get conversations email participant wrapupcodes params
func (o *GetConversationsEmailParticipantWrapupcodesParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the get conversations email participant wrapupcodes params
func (o *GetConversationsEmailParticipantWrapupcodesParams) WithParticipantID(participantID string) *GetConversationsEmailParticipantWrapupcodesParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the get conversations email participant wrapupcodes params
func (o *GetConversationsEmailParticipantWrapupcodesParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsEmailParticipantWrapupcodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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