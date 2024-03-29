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

// NewGetConversationParticipantWrapupcodesParams creates a new GetConversationParticipantWrapupcodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConversationParticipantWrapupcodesParams() *GetConversationParticipantWrapupcodesParams {
	return &GetConversationParticipantWrapupcodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationParticipantWrapupcodesParamsWithTimeout creates a new GetConversationParticipantWrapupcodesParams object
// with the ability to set a timeout on a request.
func NewGetConversationParticipantWrapupcodesParamsWithTimeout(timeout time.Duration) *GetConversationParticipantWrapupcodesParams {
	return &GetConversationParticipantWrapupcodesParams{
		timeout: timeout,
	}
}

// NewGetConversationParticipantWrapupcodesParamsWithContext creates a new GetConversationParticipantWrapupcodesParams object
// with the ability to set a context for a request.
func NewGetConversationParticipantWrapupcodesParamsWithContext(ctx context.Context) *GetConversationParticipantWrapupcodesParams {
	return &GetConversationParticipantWrapupcodesParams{
		Context: ctx,
	}
}

// NewGetConversationParticipantWrapupcodesParamsWithHTTPClient creates a new GetConversationParticipantWrapupcodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConversationParticipantWrapupcodesParamsWithHTTPClient(client *http.Client) *GetConversationParticipantWrapupcodesParams {
	return &GetConversationParticipantWrapupcodesParams{
		HTTPClient: client,
	}
}

/*
GetConversationParticipantWrapupcodesParams contains all the parameters to send to the API endpoint

	for the get conversation participant wrapupcodes operation.

	Typically these are written to a http.Request.
*/
type GetConversationParticipantWrapupcodesParams struct {

	/* ConversationID.

	   conversation ID
	*/
	ConversationID string

	/* ParticipantID.

	   participant ID
	*/
	ParticipantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get conversation participant wrapupcodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationParticipantWrapupcodesParams) WithDefaults() *GetConversationParticipantWrapupcodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get conversation participant wrapupcodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationParticipantWrapupcodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get conversation participant wrapupcodes params
func (o *GetConversationParticipantWrapupcodesParams) WithTimeout(timeout time.Duration) *GetConversationParticipantWrapupcodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversation participant wrapupcodes params
func (o *GetConversationParticipantWrapupcodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversation participant wrapupcodes params
func (o *GetConversationParticipantWrapupcodesParams) WithContext(ctx context.Context) *GetConversationParticipantWrapupcodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversation participant wrapupcodes params
func (o *GetConversationParticipantWrapupcodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversation participant wrapupcodes params
func (o *GetConversationParticipantWrapupcodesParams) WithHTTPClient(client *http.Client) *GetConversationParticipantWrapupcodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversation participant wrapupcodes params
func (o *GetConversationParticipantWrapupcodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the get conversation participant wrapupcodes params
func (o *GetConversationParticipantWrapupcodesParams) WithConversationID(conversationID string) *GetConversationParticipantWrapupcodesParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get conversation participant wrapupcodes params
func (o *GetConversationParticipantWrapupcodesParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the get conversation participant wrapupcodes params
func (o *GetConversationParticipantWrapupcodesParams) WithParticipantID(participantID string) *GetConversationParticipantWrapupcodesParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the get conversation participant wrapupcodes params
func (o *GetConversationParticipantWrapupcodesParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationParticipantWrapupcodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
