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

// NewPostConversationParticipantReplaceParams creates a new PostConversationParticipantReplaceParams object
// with the default values initialized.
func NewPostConversationParticipantReplaceParams() *PostConversationParticipantReplaceParams {
	var ()
	return &PostConversationParticipantReplaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationParticipantReplaceParamsWithTimeout creates a new PostConversationParticipantReplaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostConversationParticipantReplaceParamsWithTimeout(timeout time.Duration) *PostConversationParticipantReplaceParams {
	var ()
	return &PostConversationParticipantReplaceParams{

		timeout: timeout,
	}
}

// NewPostConversationParticipantReplaceParamsWithContext creates a new PostConversationParticipantReplaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostConversationParticipantReplaceParamsWithContext(ctx context.Context) *PostConversationParticipantReplaceParams {
	var ()
	return &PostConversationParticipantReplaceParams{

		Context: ctx,
	}
}

// NewPostConversationParticipantReplaceParamsWithHTTPClient creates a new PostConversationParticipantReplaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostConversationParticipantReplaceParamsWithHTTPClient(client *http.Client) *PostConversationParticipantReplaceParams {
	var ()
	return &PostConversationParticipantReplaceParams{
		HTTPClient: client,
	}
}

/*PostConversationParticipantReplaceParams contains all the parameters to send to the API endpoint
for the post conversation participant replace operation typically these are written to a http.Request
*/
type PostConversationParticipantReplaceParams struct {

	/*Body
	  Transfer request

	*/
	Body *models.TransferRequest
	/*ConversationID
	  conversation ID

	*/
	ConversationID string
	/*ParticipantID
	  participant ID

	*/
	ParticipantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post conversation participant replace params
func (o *PostConversationParticipantReplaceParams) WithTimeout(timeout time.Duration) *PostConversationParticipantReplaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversation participant replace params
func (o *PostConversationParticipantReplaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversation participant replace params
func (o *PostConversationParticipantReplaceParams) WithContext(ctx context.Context) *PostConversationParticipantReplaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversation participant replace params
func (o *PostConversationParticipantReplaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversation participant replace params
func (o *PostConversationParticipantReplaceParams) WithHTTPClient(client *http.Client) *PostConversationParticipantReplaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversation participant replace params
func (o *PostConversationParticipantReplaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversation participant replace params
func (o *PostConversationParticipantReplaceParams) WithBody(body *models.TransferRequest) *PostConversationParticipantReplaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversation participant replace params
func (o *PostConversationParticipantReplaceParams) SetBody(body *models.TransferRequest) {
	o.Body = body
}

// WithConversationID adds the conversationID to the post conversation participant replace params
func (o *PostConversationParticipantReplaceParams) WithConversationID(conversationID string) *PostConversationParticipantReplaceParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the post conversation participant replace params
func (o *PostConversationParticipantReplaceParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the post conversation participant replace params
func (o *PostConversationParticipantReplaceParams) WithParticipantID(participantID string) *PostConversationParticipantReplaceParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the post conversation participant replace params
func (o *PostConversationParticipantReplaceParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationParticipantReplaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param participantId
	if err := r.SetPathParam("participantId", o.ParticipantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
