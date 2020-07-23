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
	"github.com/go-openapi/swag"
)

// NewGetConversationsCallParticipantWrapupParams creates a new GetConversationsCallParticipantWrapupParams object
// with the default values initialized.
func NewGetConversationsCallParticipantWrapupParams() *GetConversationsCallParticipantWrapupParams {
	var (
		provisionalDefault = bool(false)
	)
	return &GetConversationsCallParticipantWrapupParams{
		Provisional: &provisionalDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsCallParticipantWrapupParamsWithTimeout creates a new GetConversationsCallParticipantWrapupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationsCallParticipantWrapupParamsWithTimeout(timeout time.Duration) *GetConversationsCallParticipantWrapupParams {
	var (
		provisionalDefault = bool(false)
	)
	return &GetConversationsCallParticipantWrapupParams{
		Provisional: &provisionalDefault,

		timeout: timeout,
	}
}

// NewGetConversationsCallParticipantWrapupParamsWithContext creates a new GetConversationsCallParticipantWrapupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationsCallParticipantWrapupParamsWithContext(ctx context.Context) *GetConversationsCallParticipantWrapupParams {
	var (
		provisionalDefault = bool(false)
	)
	return &GetConversationsCallParticipantWrapupParams{
		Provisional: &provisionalDefault,

		Context: ctx,
	}
}

// NewGetConversationsCallParticipantWrapupParamsWithHTTPClient creates a new GetConversationsCallParticipantWrapupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationsCallParticipantWrapupParamsWithHTTPClient(client *http.Client) *GetConversationsCallParticipantWrapupParams {
	var (
		provisionalDefault = bool(false)
	)
	return &GetConversationsCallParticipantWrapupParams{
		Provisional: &provisionalDefault,
		HTTPClient:  client,
	}
}

/*GetConversationsCallParticipantWrapupParams contains all the parameters to send to the API endpoint
for the get conversations call participant wrapup operation typically these are written to a http.Request
*/
type GetConversationsCallParticipantWrapupParams struct {

	/*ConversationID
	  conversationId

	*/
	ConversationID string
	/*ParticipantID
	  participantId

	*/
	ParticipantID string
	/*Provisional
	  Indicates if the wrap-up code is provisional.

	*/
	Provisional *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get conversations call participant wrapup params
func (o *GetConversationsCallParticipantWrapupParams) WithTimeout(timeout time.Duration) *GetConversationsCallParticipantWrapupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations call participant wrapup params
func (o *GetConversationsCallParticipantWrapupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations call participant wrapup params
func (o *GetConversationsCallParticipantWrapupParams) WithContext(ctx context.Context) *GetConversationsCallParticipantWrapupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations call participant wrapup params
func (o *GetConversationsCallParticipantWrapupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations call participant wrapup params
func (o *GetConversationsCallParticipantWrapupParams) WithHTTPClient(client *http.Client) *GetConversationsCallParticipantWrapupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations call participant wrapup params
func (o *GetConversationsCallParticipantWrapupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the get conversations call participant wrapup params
func (o *GetConversationsCallParticipantWrapupParams) WithConversationID(conversationID string) *GetConversationsCallParticipantWrapupParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get conversations call participant wrapup params
func (o *GetConversationsCallParticipantWrapupParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the get conversations call participant wrapup params
func (o *GetConversationsCallParticipantWrapupParams) WithParticipantID(participantID string) *GetConversationsCallParticipantWrapupParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the get conversations call participant wrapup params
func (o *GetConversationsCallParticipantWrapupParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WithProvisional adds the provisional to the get conversations call participant wrapup params
func (o *GetConversationsCallParticipantWrapupParams) WithProvisional(provisional *bool) *GetConversationsCallParticipantWrapupParams {
	o.SetProvisional(provisional)
	return o
}

// SetProvisional adds the provisional to the get conversations call participant wrapup params
func (o *GetConversationsCallParticipantWrapupParams) SetProvisional(provisional *bool) {
	o.Provisional = provisional
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsCallParticipantWrapupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Provisional != nil {

		// query param provisional
		var qrProvisional bool
		if o.Provisional != nil {
			qrProvisional = *o.Provisional
		}
		qProvisional := swag.FormatBool(qrProvisional)
		if qProvisional != "" {
			if err := r.SetQueryParam("provisional", qProvisional); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
