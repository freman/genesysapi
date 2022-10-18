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

	"github.com/freman/genesysapi/models"
)

// NewPostConversationsMessageCommunicationMessagesParams creates a new PostConversationsMessageCommunicationMessagesParams object
// with the default values initialized.
func NewPostConversationsMessageCommunicationMessagesParams() *PostConversationsMessageCommunicationMessagesParams {
	var (
		useNormalizedMessageDefault = bool(false)
	)
	return &PostConversationsMessageCommunicationMessagesParams{
		UseNormalizedMessage: &useNormalizedMessageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationsMessageCommunicationMessagesParamsWithTimeout creates a new PostConversationsMessageCommunicationMessagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostConversationsMessageCommunicationMessagesParamsWithTimeout(timeout time.Duration) *PostConversationsMessageCommunicationMessagesParams {
	var (
		useNormalizedMessageDefault = bool(false)
	)
	return &PostConversationsMessageCommunicationMessagesParams{
		UseNormalizedMessage: &useNormalizedMessageDefault,

		timeout: timeout,
	}
}

// NewPostConversationsMessageCommunicationMessagesParamsWithContext creates a new PostConversationsMessageCommunicationMessagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostConversationsMessageCommunicationMessagesParamsWithContext(ctx context.Context) *PostConversationsMessageCommunicationMessagesParams {
	var (
		useNormalizedMessageDefault = bool(false)
	)
	return &PostConversationsMessageCommunicationMessagesParams{
		UseNormalizedMessage: &useNormalizedMessageDefault,

		Context: ctx,
	}
}

// NewPostConversationsMessageCommunicationMessagesParamsWithHTTPClient creates a new PostConversationsMessageCommunicationMessagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostConversationsMessageCommunicationMessagesParamsWithHTTPClient(client *http.Client) *PostConversationsMessageCommunicationMessagesParams {
	var (
		useNormalizedMessageDefault = bool(false)
	)
	return &PostConversationsMessageCommunicationMessagesParams{
		UseNormalizedMessage: &useNormalizedMessageDefault,
		HTTPClient:           client,
	}
}

/*PostConversationsMessageCommunicationMessagesParams contains all the parameters to send to the API endpoint
for the post conversations message communication messages operation typically these are written to a http.Request
*/
type PostConversationsMessageCommunicationMessagesParams struct {

	/*Body
	  Message

	*/
	Body *models.AdditionalMessage
	/*CommunicationID
	  communicationId

	*/
	CommunicationID string
	/*ConversationID
	  conversationId

	*/
	ConversationID string
	/*UseNormalizedMessage
	  If true, response removes deprecated fields (textBody, media, stickers)

	*/
	UseNormalizedMessage *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) WithTimeout(timeout time.Duration) *PostConversationsMessageCommunicationMessagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) WithContext(ctx context.Context) *PostConversationsMessageCommunicationMessagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) WithHTTPClient(client *http.Client) *PostConversationsMessageCommunicationMessagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) WithBody(body *models.AdditionalMessage) *PostConversationsMessageCommunicationMessagesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) SetBody(body *models.AdditionalMessage) {
	o.Body = body
}

// WithCommunicationID adds the communicationID to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) WithCommunicationID(communicationID string) *PostConversationsMessageCommunicationMessagesParams {
	o.SetCommunicationID(communicationID)
	return o
}

// SetCommunicationID adds the communicationId to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) SetCommunicationID(communicationID string) {
	o.CommunicationID = communicationID
}

// WithConversationID adds the conversationID to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) WithConversationID(conversationID string) *PostConversationsMessageCommunicationMessagesParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithUseNormalizedMessage adds the useNormalizedMessage to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) WithUseNormalizedMessage(useNormalizedMessage *bool) *PostConversationsMessageCommunicationMessagesParams {
	o.SetUseNormalizedMessage(useNormalizedMessage)
	return o
}

// SetUseNormalizedMessage adds the useNormalizedMessage to the post conversations message communication messages params
func (o *PostConversationsMessageCommunicationMessagesParams) SetUseNormalizedMessage(useNormalizedMessage *bool) {
	o.UseNormalizedMessage = useNormalizedMessage
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationsMessageCommunicationMessagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.UseNormalizedMessage != nil {

		// query param useNormalizedMessage
		var qrUseNormalizedMessage bool
		if o.UseNormalizedMessage != nil {
			qrUseNormalizedMessage = *o.UseNormalizedMessage
		}
		qUseNormalizedMessage := swag.FormatBool(qrUseNormalizedMessage)
		if qUseNormalizedMessage != "" {
			if err := r.SetQueryParam("useNormalizedMessage", qUseNormalizedMessage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
