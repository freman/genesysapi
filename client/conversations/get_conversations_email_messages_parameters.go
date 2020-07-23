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

// NewGetConversationsEmailMessagesParams creates a new GetConversationsEmailMessagesParams object
// with the default values initialized.
func NewGetConversationsEmailMessagesParams() *GetConversationsEmailMessagesParams {
	var ()
	return &GetConversationsEmailMessagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsEmailMessagesParamsWithTimeout creates a new GetConversationsEmailMessagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationsEmailMessagesParamsWithTimeout(timeout time.Duration) *GetConversationsEmailMessagesParams {
	var ()
	return &GetConversationsEmailMessagesParams{

		timeout: timeout,
	}
}

// NewGetConversationsEmailMessagesParamsWithContext creates a new GetConversationsEmailMessagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationsEmailMessagesParamsWithContext(ctx context.Context) *GetConversationsEmailMessagesParams {
	var ()
	return &GetConversationsEmailMessagesParams{

		Context: ctx,
	}
}

// NewGetConversationsEmailMessagesParamsWithHTTPClient creates a new GetConversationsEmailMessagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationsEmailMessagesParamsWithHTTPClient(client *http.Client) *GetConversationsEmailMessagesParams {
	var ()
	return &GetConversationsEmailMessagesParams{
		HTTPClient: client,
	}
}

/*GetConversationsEmailMessagesParams contains all the parameters to send to the API endpoint
for the get conversations email messages operation typically these are written to a http.Request
*/
type GetConversationsEmailMessagesParams struct {

	/*ConversationID
	  conversationId

	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get conversations email messages params
func (o *GetConversationsEmailMessagesParams) WithTimeout(timeout time.Duration) *GetConversationsEmailMessagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations email messages params
func (o *GetConversationsEmailMessagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations email messages params
func (o *GetConversationsEmailMessagesParams) WithContext(ctx context.Context) *GetConversationsEmailMessagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations email messages params
func (o *GetConversationsEmailMessagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations email messages params
func (o *GetConversationsEmailMessagesParams) WithHTTPClient(client *http.Client) *GetConversationsEmailMessagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations email messages params
func (o *GetConversationsEmailMessagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the get conversations email messages params
func (o *GetConversationsEmailMessagesParams) WithConversationID(conversationID string) *GetConversationsEmailMessagesParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get conversations email messages params
func (o *GetConversationsEmailMessagesParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsEmailMessagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param conversationId
	if err := r.SetPathParam("conversationId", o.ConversationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
