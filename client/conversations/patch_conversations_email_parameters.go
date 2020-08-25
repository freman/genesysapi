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

	"github.com/freman/genesysapi/models"
)

// NewPatchConversationsEmailParams creates a new PatchConversationsEmailParams object
// with the default values initialized.
func NewPatchConversationsEmailParams() *PatchConversationsEmailParams {
	var ()
	return &PatchConversationsEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConversationsEmailParamsWithTimeout creates a new PatchConversationsEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchConversationsEmailParamsWithTimeout(timeout time.Duration) *PatchConversationsEmailParams {
	var ()
	return &PatchConversationsEmailParams{

		timeout: timeout,
	}
}

// NewPatchConversationsEmailParamsWithContext creates a new PatchConversationsEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchConversationsEmailParamsWithContext(ctx context.Context) *PatchConversationsEmailParams {
	var ()
	return &PatchConversationsEmailParams{

		Context: ctx,
	}
}

// NewPatchConversationsEmailParamsWithHTTPClient creates a new PatchConversationsEmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchConversationsEmailParamsWithHTTPClient(client *http.Client) *PatchConversationsEmailParams {
	var ()
	return &PatchConversationsEmailParams{
		HTTPClient: client,
	}
}

/*PatchConversationsEmailParams contains all the parameters to send to the API endpoint
for the patch conversations email operation typically these are written to a http.Request
*/
type PatchConversationsEmailParams struct {

	/*Body
	  Conversation

	*/
	Body *models.Conversation
	/*ConversationID
	  conversationId

	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch conversations email params
func (o *PatchConversationsEmailParams) WithTimeout(timeout time.Duration) *PatchConversationsEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch conversations email params
func (o *PatchConversationsEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch conversations email params
func (o *PatchConversationsEmailParams) WithContext(ctx context.Context) *PatchConversationsEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch conversations email params
func (o *PatchConversationsEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch conversations email params
func (o *PatchConversationsEmailParams) WithHTTPClient(client *http.Client) *PatchConversationsEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch conversations email params
func (o *PatchConversationsEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch conversations email params
func (o *PatchConversationsEmailParams) WithBody(body *models.Conversation) *PatchConversationsEmailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch conversations email params
func (o *PatchConversationsEmailParams) SetBody(body *models.Conversation) {
	o.Body = body
}

// WithConversationID adds the conversationID to the patch conversations email params
func (o *PatchConversationsEmailParams) WithConversationID(conversationID string) *PatchConversationsEmailParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch conversations email params
func (o *PatchConversationsEmailParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConversationsEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param conversationId
	if err := r.SetPathParam("conversationId", o.ConversationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}