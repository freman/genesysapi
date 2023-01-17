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

// NewGetConversationsMessageParticipantWrapupcodesParams creates a new GetConversationsMessageParticipantWrapupcodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConversationsMessageParticipantWrapupcodesParams() *GetConversationsMessageParticipantWrapupcodesParams {
	return &GetConversationsMessageParticipantWrapupcodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsMessageParticipantWrapupcodesParamsWithTimeout creates a new GetConversationsMessageParticipantWrapupcodesParams object
// with the ability to set a timeout on a request.
func NewGetConversationsMessageParticipantWrapupcodesParamsWithTimeout(timeout time.Duration) *GetConversationsMessageParticipantWrapupcodesParams {
	return &GetConversationsMessageParticipantWrapupcodesParams{
		timeout: timeout,
	}
}

// NewGetConversationsMessageParticipantWrapupcodesParamsWithContext creates a new GetConversationsMessageParticipantWrapupcodesParams object
// with the ability to set a context for a request.
func NewGetConversationsMessageParticipantWrapupcodesParamsWithContext(ctx context.Context) *GetConversationsMessageParticipantWrapupcodesParams {
	return &GetConversationsMessageParticipantWrapupcodesParams{
		Context: ctx,
	}
}

// NewGetConversationsMessageParticipantWrapupcodesParamsWithHTTPClient creates a new GetConversationsMessageParticipantWrapupcodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConversationsMessageParticipantWrapupcodesParamsWithHTTPClient(client *http.Client) *GetConversationsMessageParticipantWrapupcodesParams {
	return &GetConversationsMessageParticipantWrapupcodesParams{
		HTTPClient: client,
	}
}

/*
GetConversationsMessageParticipantWrapupcodesParams contains all the parameters to send to the API endpoint

	for the get conversations message participant wrapupcodes operation.

	Typically these are written to a http.Request.
*/
type GetConversationsMessageParticipantWrapupcodesParams struct {

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

// WithDefaults hydrates default values in the get conversations message participant wrapupcodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsMessageParticipantWrapupcodesParams) WithDefaults() *GetConversationsMessageParticipantWrapupcodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get conversations message participant wrapupcodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsMessageParticipantWrapupcodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get conversations message participant wrapupcodes params
func (o *GetConversationsMessageParticipantWrapupcodesParams) WithTimeout(timeout time.Duration) *GetConversationsMessageParticipantWrapupcodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations message participant wrapupcodes params
func (o *GetConversationsMessageParticipantWrapupcodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations message participant wrapupcodes params
func (o *GetConversationsMessageParticipantWrapupcodesParams) WithContext(ctx context.Context) *GetConversationsMessageParticipantWrapupcodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations message participant wrapupcodes params
func (o *GetConversationsMessageParticipantWrapupcodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations message participant wrapupcodes params
func (o *GetConversationsMessageParticipantWrapupcodesParams) WithHTTPClient(client *http.Client) *GetConversationsMessageParticipantWrapupcodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations message participant wrapupcodes params
func (o *GetConversationsMessageParticipantWrapupcodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the get conversations message participant wrapupcodes params
func (o *GetConversationsMessageParticipantWrapupcodesParams) WithConversationID(conversationID string) *GetConversationsMessageParticipantWrapupcodesParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get conversations message participant wrapupcodes params
func (o *GetConversationsMessageParticipantWrapupcodesParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the get conversations message participant wrapupcodes params
func (o *GetConversationsMessageParticipantWrapupcodesParams) WithParticipantID(participantID string) *GetConversationsMessageParticipantWrapupcodesParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the get conversations message participant wrapupcodes params
func (o *GetConversationsMessageParticipantWrapupcodesParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsMessageParticipantWrapupcodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
