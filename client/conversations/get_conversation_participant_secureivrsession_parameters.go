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

// NewGetConversationParticipantSecureivrsessionParams creates a new GetConversationParticipantSecureivrsessionParams object
// with the default values initialized.
func NewGetConversationParticipantSecureivrsessionParams() *GetConversationParticipantSecureivrsessionParams {
	var ()
	return &GetConversationParticipantSecureivrsessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationParticipantSecureivrsessionParamsWithTimeout creates a new GetConversationParticipantSecureivrsessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationParticipantSecureivrsessionParamsWithTimeout(timeout time.Duration) *GetConversationParticipantSecureivrsessionParams {
	var ()
	return &GetConversationParticipantSecureivrsessionParams{

		timeout: timeout,
	}
}

// NewGetConversationParticipantSecureivrsessionParamsWithContext creates a new GetConversationParticipantSecureivrsessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationParticipantSecureivrsessionParamsWithContext(ctx context.Context) *GetConversationParticipantSecureivrsessionParams {
	var ()
	return &GetConversationParticipantSecureivrsessionParams{

		Context: ctx,
	}
}

// NewGetConversationParticipantSecureivrsessionParamsWithHTTPClient creates a new GetConversationParticipantSecureivrsessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationParticipantSecureivrsessionParamsWithHTTPClient(client *http.Client) *GetConversationParticipantSecureivrsessionParams {
	var ()
	return &GetConversationParticipantSecureivrsessionParams{
		HTTPClient: client,
	}
}

/*GetConversationParticipantSecureivrsessionParams contains all the parameters to send to the API endpoint
for the get conversation participant secureivrsession operation typically these are written to a http.Request
*/
type GetConversationParticipantSecureivrsessionParams struct {

	/*ConversationID
	  conversation ID

	*/
	ConversationID string
	/*ParticipantID
	  participant ID

	*/
	ParticipantID string
	/*SecureSessionID
	  secure IVR session ID

	*/
	SecureSessionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get conversation participant secureivrsession params
func (o *GetConversationParticipantSecureivrsessionParams) WithTimeout(timeout time.Duration) *GetConversationParticipantSecureivrsessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversation participant secureivrsession params
func (o *GetConversationParticipantSecureivrsessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversation participant secureivrsession params
func (o *GetConversationParticipantSecureivrsessionParams) WithContext(ctx context.Context) *GetConversationParticipantSecureivrsessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversation participant secureivrsession params
func (o *GetConversationParticipantSecureivrsessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversation participant secureivrsession params
func (o *GetConversationParticipantSecureivrsessionParams) WithHTTPClient(client *http.Client) *GetConversationParticipantSecureivrsessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversation participant secureivrsession params
func (o *GetConversationParticipantSecureivrsessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the get conversation participant secureivrsession params
func (o *GetConversationParticipantSecureivrsessionParams) WithConversationID(conversationID string) *GetConversationParticipantSecureivrsessionParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get conversation participant secureivrsession params
func (o *GetConversationParticipantSecureivrsessionParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the get conversation participant secureivrsession params
func (o *GetConversationParticipantSecureivrsessionParams) WithParticipantID(participantID string) *GetConversationParticipantSecureivrsessionParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the get conversation participant secureivrsession params
func (o *GetConversationParticipantSecureivrsessionParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WithSecureSessionID adds the secureSessionID to the get conversation participant secureivrsession params
func (o *GetConversationParticipantSecureivrsessionParams) WithSecureSessionID(secureSessionID string) *GetConversationParticipantSecureivrsessionParams {
	o.SetSecureSessionID(secureSessionID)
	return o
}

// SetSecureSessionID adds the secureSessionId to the get conversation participant secureivrsession params
func (o *GetConversationParticipantSecureivrsessionParams) SetSecureSessionID(secureSessionID string) {
	o.SecureSessionID = secureSessionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationParticipantSecureivrsessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param secureSessionId
	if err := r.SetPathParam("secureSessionId", o.SecureSessionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}