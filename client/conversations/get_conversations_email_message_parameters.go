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

// NewGetConversationsEmailMessageParams creates a new GetConversationsEmailMessageParams object
// with the default values initialized.
func NewGetConversationsEmailMessageParams() *GetConversationsEmailMessageParams {
	var ()
	return &GetConversationsEmailMessageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsEmailMessageParamsWithTimeout creates a new GetConversationsEmailMessageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationsEmailMessageParamsWithTimeout(timeout time.Duration) *GetConversationsEmailMessageParams {
	var ()
	return &GetConversationsEmailMessageParams{

		timeout: timeout,
	}
}

// NewGetConversationsEmailMessageParamsWithContext creates a new GetConversationsEmailMessageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationsEmailMessageParamsWithContext(ctx context.Context) *GetConversationsEmailMessageParams {
	var ()
	return &GetConversationsEmailMessageParams{

		Context: ctx,
	}
}

// NewGetConversationsEmailMessageParamsWithHTTPClient creates a new GetConversationsEmailMessageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationsEmailMessageParamsWithHTTPClient(client *http.Client) *GetConversationsEmailMessageParams {
	var ()
	return &GetConversationsEmailMessageParams{
		HTTPClient: client,
	}
}

/*GetConversationsEmailMessageParams contains all the parameters to send to the API endpoint
for the get conversations email message operation typically these are written to a http.Request
*/
type GetConversationsEmailMessageParams struct {

	/*ConversationID
	  conversationId

	*/
	ConversationID string
	/*MessageID
	  messageId

	*/
	MessageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get conversations email message params
func (o *GetConversationsEmailMessageParams) WithTimeout(timeout time.Duration) *GetConversationsEmailMessageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations email message params
func (o *GetConversationsEmailMessageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations email message params
func (o *GetConversationsEmailMessageParams) WithContext(ctx context.Context) *GetConversationsEmailMessageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations email message params
func (o *GetConversationsEmailMessageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations email message params
func (o *GetConversationsEmailMessageParams) WithHTTPClient(client *http.Client) *GetConversationsEmailMessageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations email message params
func (o *GetConversationsEmailMessageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the get conversations email message params
func (o *GetConversationsEmailMessageParams) WithConversationID(conversationID string) *GetConversationsEmailMessageParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get conversations email message params
func (o *GetConversationsEmailMessageParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithMessageID adds the messageID to the get conversations email message params
func (o *GetConversationsEmailMessageParams) WithMessageID(messageID string) *GetConversationsEmailMessageParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the get conversations email message params
func (o *GetConversationsEmailMessageParams) SetMessageID(messageID string) {
	o.MessageID = messageID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsEmailMessageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param conversationId
	if err := r.SetPathParam("conversationId", o.ConversationID); err != nil {
		return err
	}

	// path param messageId
	if err := r.SetPathParam("messageId", o.MessageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}