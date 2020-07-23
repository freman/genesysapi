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

// NewPostConversationsCallParticipantReplaceParams creates a new PostConversationsCallParticipantReplaceParams object
// with the default values initialized.
func NewPostConversationsCallParticipantReplaceParams() *PostConversationsCallParticipantReplaceParams {
	var ()
	return &PostConversationsCallParticipantReplaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationsCallParticipantReplaceParamsWithTimeout creates a new PostConversationsCallParticipantReplaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostConversationsCallParticipantReplaceParamsWithTimeout(timeout time.Duration) *PostConversationsCallParticipantReplaceParams {
	var ()
	return &PostConversationsCallParticipantReplaceParams{

		timeout: timeout,
	}
}

// NewPostConversationsCallParticipantReplaceParamsWithContext creates a new PostConversationsCallParticipantReplaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostConversationsCallParticipantReplaceParamsWithContext(ctx context.Context) *PostConversationsCallParticipantReplaceParams {
	var ()
	return &PostConversationsCallParticipantReplaceParams{

		Context: ctx,
	}
}

// NewPostConversationsCallParticipantReplaceParamsWithHTTPClient creates a new PostConversationsCallParticipantReplaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostConversationsCallParticipantReplaceParamsWithHTTPClient(client *http.Client) *PostConversationsCallParticipantReplaceParams {
	var ()
	return &PostConversationsCallParticipantReplaceParams{
		HTTPClient: client,
	}
}

/*PostConversationsCallParticipantReplaceParams contains all the parameters to send to the API endpoint
for the post conversations call participant replace operation typically these are written to a http.Request
*/
type PostConversationsCallParticipantReplaceParams struct {

	/*Body
	  Transfer request

	*/
	Body *models.TransferRequest
	/*ConversationID
	  conversationId

	*/
	ConversationID string
	/*ParticipantID
	  participantId

	*/
	ParticipantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post conversations call participant replace params
func (o *PostConversationsCallParticipantReplaceParams) WithTimeout(timeout time.Duration) *PostConversationsCallParticipantReplaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversations call participant replace params
func (o *PostConversationsCallParticipantReplaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversations call participant replace params
func (o *PostConversationsCallParticipantReplaceParams) WithContext(ctx context.Context) *PostConversationsCallParticipantReplaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversations call participant replace params
func (o *PostConversationsCallParticipantReplaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversations call participant replace params
func (o *PostConversationsCallParticipantReplaceParams) WithHTTPClient(client *http.Client) *PostConversationsCallParticipantReplaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversations call participant replace params
func (o *PostConversationsCallParticipantReplaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversations call participant replace params
func (o *PostConversationsCallParticipantReplaceParams) WithBody(body *models.TransferRequest) *PostConversationsCallParticipantReplaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversations call participant replace params
func (o *PostConversationsCallParticipantReplaceParams) SetBody(body *models.TransferRequest) {
	o.Body = body
}

// WithConversationID adds the conversationID to the post conversations call participant replace params
func (o *PostConversationsCallParticipantReplaceParams) WithConversationID(conversationID string) *PostConversationsCallParticipantReplaceParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the post conversations call participant replace params
func (o *PostConversationsCallParticipantReplaceParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the post conversations call participant replace params
func (o *PostConversationsCallParticipantReplaceParams) WithParticipantID(participantID string) *PostConversationsCallParticipantReplaceParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the post conversations call participant replace params
func (o *PostConversationsCallParticipantReplaceParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationsCallParticipantReplaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
