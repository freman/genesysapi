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

// NewGetConversationsCallParams creates a new GetConversationsCallParams object
// with the default values initialized.
func NewGetConversationsCallParams() *GetConversationsCallParams {
	var ()
	return &GetConversationsCallParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsCallParamsWithTimeout creates a new GetConversationsCallParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationsCallParamsWithTimeout(timeout time.Duration) *GetConversationsCallParams {
	var ()
	return &GetConversationsCallParams{

		timeout: timeout,
	}
}

// NewGetConversationsCallParamsWithContext creates a new GetConversationsCallParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationsCallParamsWithContext(ctx context.Context) *GetConversationsCallParams {
	var ()
	return &GetConversationsCallParams{

		Context: ctx,
	}
}

// NewGetConversationsCallParamsWithHTTPClient creates a new GetConversationsCallParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationsCallParamsWithHTTPClient(client *http.Client) *GetConversationsCallParams {
	var ()
	return &GetConversationsCallParams{
		HTTPClient: client,
	}
}

/*GetConversationsCallParams contains all the parameters to send to the API endpoint
for the get conversations call operation typically these are written to a http.Request
*/
type GetConversationsCallParams struct {

	/*ConversationID
	  conversationId

	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get conversations call params
func (o *GetConversationsCallParams) WithTimeout(timeout time.Duration) *GetConversationsCallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations call params
func (o *GetConversationsCallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations call params
func (o *GetConversationsCallParams) WithContext(ctx context.Context) *GetConversationsCallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations call params
func (o *GetConversationsCallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations call params
func (o *GetConversationsCallParams) WithHTTPClient(client *http.Client) *GetConversationsCallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations call params
func (o *GetConversationsCallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the get conversations call params
func (o *GetConversationsCallParams) WithConversationID(conversationID string) *GetConversationsCallParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get conversations call params
func (o *GetConversationsCallParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsCallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
