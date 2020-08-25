// Code generated by go-swagger; DO NOT EDIT.

package web_chat

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

// NewPatchWebchatGuestConversationMediarequestParams creates a new PatchWebchatGuestConversationMediarequestParams object
// with the default values initialized.
func NewPatchWebchatGuestConversationMediarequestParams() *PatchWebchatGuestConversationMediarequestParams {
	var ()
	return &PatchWebchatGuestConversationMediarequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWebchatGuestConversationMediarequestParamsWithTimeout creates a new PatchWebchatGuestConversationMediarequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchWebchatGuestConversationMediarequestParamsWithTimeout(timeout time.Duration) *PatchWebchatGuestConversationMediarequestParams {
	var ()
	return &PatchWebchatGuestConversationMediarequestParams{

		timeout: timeout,
	}
}

// NewPatchWebchatGuestConversationMediarequestParamsWithContext creates a new PatchWebchatGuestConversationMediarequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchWebchatGuestConversationMediarequestParamsWithContext(ctx context.Context) *PatchWebchatGuestConversationMediarequestParams {
	var ()
	return &PatchWebchatGuestConversationMediarequestParams{

		Context: ctx,
	}
}

// NewPatchWebchatGuestConversationMediarequestParamsWithHTTPClient creates a new PatchWebchatGuestConversationMediarequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchWebchatGuestConversationMediarequestParamsWithHTTPClient(client *http.Client) *PatchWebchatGuestConversationMediarequestParams {
	var ()
	return &PatchWebchatGuestConversationMediarequestParams{
		HTTPClient: client,
	}
}

/*PatchWebchatGuestConversationMediarequestParams contains all the parameters to send to the API endpoint
for the patch webchat guest conversation mediarequest operation typically these are written to a http.Request
*/
type PatchWebchatGuestConversationMediarequestParams struct {

	/*Body
	  Request

	*/
	Body *models.WebChatGuestMediaRequest
	/*ConversationID
	  conversationId

	*/
	ConversationID string
	/*MediaRequestID
	  mediaRequestId

	*/
	MediaRequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch webchat guest conversation mediarequest params
func (o *PatchWebchatGuestConversationMediarequestParams) WithTimeout(timeout time.Duration) *PatchWebchatGuestConversationMediarequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch webchat guest conversation mediarequest params
func (o *PatchWebchatGuestConversationMediarequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch webchat guest conversation mediarequest params
func (o *PatchWebchatGuestConversationMediarequestParams) WithContext(ctx context.Context) *PatchWebchatGuestConversationMediarequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch webchat guest conversation mediarequest params
func (o *PatchWebchatGuestConversationMediarequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch webchat guest conversation mediarequest params
func (o *PatchWebchatGuestConversationMediarequestParams) WithHTTPClient(client *http.Client) *PatchWebchatGuestConversationMediarequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch webchat guest conversation mediarequest params
func (o *PatchWebchatGuestConversationMediarequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch webchat guest conversation mediarequest params
func (o *PatchWebchatGuestConversationMediarequestParams) WithBody(body *models.WebChatGuestMediaRequest) *PatchWebchatGuestConversationMediarequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch webchat guest conversation mediarequest params
func (o *PatchWebchatGuestConversationMediarequestParams) SetBody(body *models.WebChatGuestMediaRequest) {
	o.Body = body
}

// WithConversationID adds the conversationID to the patch webchat guest conversation mediarequest params
func (o *PatchWebchatGuestConversationMediarequestParams) WithConversationID(conversationID string) *PatchWebchatGuestConversationMediarequestParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch webchat guest conversation mediarequest params
func (o *PatchWebchatGuestConversationMediarequestParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithMediaRequestID adds the mediaRequestID to the patch webchat guest conversation mediarequest params
func (o *PatchWebchatGuestConversationMediarequestParams) WithMediaRequestID(mediaRequestID string) *PatchWebchatGuestConversationMediarequestParams {
	o.SetMediaRequestID(mediaRequestID)
	return o
}

// SetMediaRequestID adds the mediaRequestId to the patch webchat guest conversation mediarequest params
func (o *PatchWebchatGuestConversationMediarequestParams) SetMediaRequestID(mediaRequestID string) {
	o.MediaRequestID = mediaRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWebchatGuestConversationMediarequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param mediaRequestId
	if err := r.SetPathParam("mediaRequestId", o.MediaRequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}