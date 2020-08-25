// Code generated by go-swagger; DO NOT EDIT.

package outbound

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

// NewPostOutboundConversationDncParams creates a new PostOutboundConversationDncParams object
// with the default values initialized.
func NewPostOutboundConversationDncParams() *PostOutboundConversationDncParams {
	var ()
	return &PostOutboundConversationDncParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOutboundConversationDncParamsWithTimeout creates a new PostOutboundConversationDncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOutboundConversationDncParamsWithTimeout(timeout time.Duration) *PostOutboundConversationDncParams {
	var ()
	return &PostOutboundConversationDncParams{

		timeout: timeout,
	}
}

// NewPostOutboundConversationDncParamsWithContext creates a new PostOutboundConversationDncParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOutboundConversationDncParamsWithContext(ctx context.Context) *PostOutboundConversationDncParams {
	var ()
	return &PostOutboundConversationDncParams{

		Context: ctx,
	}
}

// NewPostOutboundConversationDncParamsWithHTTPClient creates a new PostOutboundConversationDncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOutboundConversationDncParamsWithHTTPClient(client *http.Client) *PostOutboundConversationDncParams {
	var ()
	return &PostOutboundConversationDncParams{
		HTTPClient: client,
	}
}

/*PostOutboundConversationDncParams contains all the parameters to send to the API endpoint
for the post outbound conversation dnc operation typically these are written to a http.Request
*/
type PostOutboundConversationDncParams struct {

	/*ConversationID
	  Conversation ID

	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post outbound conversation dnc params
func (o *PostOutboundConversationDncParams) WithTimeout(timeout time.Duration) *PostOutboundConversationDncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post outbound conversation dnc params
func (o *PostOutboundConversationDncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post outbound conversation dnc params
func (o *PostOutboundConversationDncParams) WithContext(ctx context.Context) *PostOutboundConversationDncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post outbound conversation dnc params
func (o *PostOutboundConversationDncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post outbound conversation dnc params
func (o *PostOutboundConversationDncParams) WithHTTPClient(client *http.Client) *PostOutboundConversationDncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post outbound conversation dnc params
func (o *PostOutboundConversationDncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the post outbound conversation dnc params
func (o *PostOutboundConversationDncParams) WithConversationID(conversationID string) *PostOutboundConversationDncParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the post outbound conversation dnc params
func (o *PostOutboundConversationDncParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PostOutboundConversationDncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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