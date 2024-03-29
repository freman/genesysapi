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

// NewGetConversationsCallbackParticipantWrapupcodesParams creates a new GetConversationsCallbackParticipantWrapupcodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConversationsCallbackParticipantWrapupcodesParams() *GetConversationsCallbackParticipantWrapupcodesParams {
	return &GetConversationsCallbackParticipantWrapupcodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsCallbackParticipantWrapupcodesParamsWithTimeout creates a new GetConversationsCallbackParticipantWrapupcodesParams object
// with the ability to set a timeout on a request.
func NewGetConversationsCallbackParticipantWrapupcodesParamsWithTimeout(timeout time.Duration) *GetConversationsCallbackParticipantWrapupcodesParams {
	return &GetConversationsCallbackParticipantWrapupcodesParams{
		timeout: timeout,
	}
}

// NewGetConversationsCallbackParticipantWrapupcodesParamsWithContext creates a new GetConversationsCallbackParticipantWrapupcodesParams object
// with the ability to set a context for a request.
func NewGetConversationsCallbackParticipantWrapupcodesParamsWithContext(ctx context.Context) *GetConversationsCallbackParticipantWrapupcodesParams {
	return &GetConversationsCallbackParticipantWrapupcodesParams{
		Context: ctx,
	}
}

// NewGetConversationsCallbackParticipantWrapupcodesParamsWithHTTPClient creates a new GetConversationsCallbackParticipantWrapupcodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConversationsCallbackParticipantWrapupcodesParamsWithHTTPClient(client *http.Client) *GetConversationsCallbackParticipantWrapupcodesParams {
	return &GetConversationsCallbackParticipantWrapupcodesParams{
		HTTPClient: client,
	}
}

/*
GetConversationsCallbackParticipantWrapupcodesParams contains all the parameters to send to the API endpoint

	for the get conversations callback participant wrapupcodes operation.

	Typically these are written to a http.Request.
*/
type GetConversationsCallbackParticipantWrapupcodesParams struct {

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

// WithDefaults hydrates default values in the get conversations callback participant wrapupcodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsCallbackParticipantWrapupcodesParams) WithDefaults() *GetConversationsCallbackParticipantWrapupcodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get conversations callback participant wrapupcodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsCallbackParticipantWrapupcodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get conversations callback participant wrapupcodes params
func (o *GetConversationsCallbackParticipantWrapupcodesParams) WithTimeout(timeout time.Duration) *GetConversationsCallbackParticipantWrapupcodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations callback participant wrapupcodes params
func (o *GetConversationsCallbackParticipantWrapupcodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations callback participant wrapupcodes params
func (o *GetConversationsCallbackParticipantWrapupcodesParams) WithContext(ctx context.Context) *GetConversationsCallbackParticipantWrapupcodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations callback participant wrapupcodes params
func (o *GetConversationsCallbackParticipantWrapupcodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations callback participant wrapupcodes params
func (o *GetConversationsCallbackParticipantWrapupcodesParams) WithHTTPClient(client *http.Client) *GetConversationsCallbackParticipantWrapupcodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations callback participant wrapupcodes params
func (o *GetConversationsCallbackParticipantWrapupcodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the get conversations callback participant wrapupcodes params
func (o *GetConversationsCallbackParticipantWrapupcodesParams) WithConversationID(conversationID string) *GetConversationsCallbackParticipantWrapupcodesParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get conversations callback participant wrapupcodes params
func (o *GetConversationsCallbackParticipantWrapupcodesParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the get conversations callback participant wrapupcodes params
func (o *GetConversationsCallbackParticipantWrapupcodesParams) WithParticipantID(participantID string) *GetConversationsCallbackParticipantWrapupcodesParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the get conversations callback participant wrapupcodes params
func (o *GetConversationsCallbackParticipantWrapupcodesParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsCallbackParticipantWrapupcodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
