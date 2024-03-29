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

// NewPostConversationsCallbackParticipantReplaceParams creates a new PostConversationsCallbackParticipantReplaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostConversationsCallbackParticipantReplaceParams() *PostConversationsCallbackParticipantReplaceParams {
	return &PostConversationsCallbackParticipantReplaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationsCallbackParticipantReplaceParamsWithTimeout creates a new PostConversationsCallbackParticipantReplaceParams object
// with the ability to set a timeout on a request.
func NewPostConversationsCallbackParticipantReplaceParamsWithTimeout(timeout time.Duration) *PostConversationsCallbackParticipantReplaceParams {
	return &PostConversationsCallbackParticipantReplaceParams{
		timeout: timeout,
	}
}

// NewPostConversationsCallbackParticipantReplaceParamsWithContext creates a new PostConversationsCallbackParticipantReplaceParams object
// with the ability to set a context for a request.
func NewPostConversationsCallbackParticipantReplaceParamsWithContext(ctx context.Context) *PostConversationsCallbackParticipantReplaceParams {
	return &PostConversationsCallbackParticipantReplaceParams{
		Context: ctx,
	}
}

// NewPostConversationsCallbackParticipantReplaceParamsWithHTTPClient creates a new PostConversationsCallbackParticipantReplaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostConversationsCallbackParticipantReplaceParamsWithHTTPClient(client *http.Client) *PostConversationsCallbackParticipantReplaceParams {
	return &PostConversationsCallbackParticipantReplaceParams{
		HTTPClient: client,
	}
}

/*
PostConversationsCallbackParticipantReplaceParams contains all the parameters to send to the API endpoint

	for the post conversations callback participant replace operation.

	Typically these are written to a http.Request.
*/
type PostConversationsCallbackParticipantReplaceParams struct {

	/* Body.

	   Transfer request
	*/
	Body *models.TransferRequest

	/* ConversationID.

	   conversationId
	*/
	ConversationID string

	/* ParticipantID.

	   participantId
	*/
	ParticipantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post conversations callback participant replace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationsCallbackParticipantReplaceParams) WithDefaults() *PostConversationsCallbackParticipantReplaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post conversations callback participant replace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationsCallbackParticipantReplaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post conversations callback participant replace params
func (o *PostConversationsCallbackParticipantReplaceParams) WithTimeout(timeout time.Duration) *PostConversationsCallbackParticipantReplaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversations callback participant replace params
func (o *PostConversationsCallbackParticipantReplaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversations callback participant replace params
func (o *PostConversationsCallbackParticipantReplaceParams) WithContext(ctx context.Context) *PostConversationsCallbackParticipantReplaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversations callback participant replace params
func (o *PostConversationsCallbackParticipantReplaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversations callback participant replace params
func (o *PostConversationsCallbackParticipantReplaceParams) WithHTTPClient(client *http.Client) *PostConversationsCallbackParticipantReplaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversations callback participant replace params
func (o *PostConversationsCallbackParticipantReplaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversations callback participant replace params
func (o *PostConversationsCallbackParticipantReplaceParams) WithBody(body *models.TransferRequest) *PostConversationsCallbackParticipantReplaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversations callback participant replace params
func (o *PostConversationsCallbackParticipantReplaceParams) SetBody(body *models.TransferRequest) {
	o.Body = body
}

// WithConversationID adds the conversationID to the post conversations callback participant replace params
func (o *PostConversationsCallbackParticipantReplaceParams) WithConversationID(conversationID string) *PostConversationsCallbackParticipantReplaceParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the post conversations callback participant replace params
func (o *PostConversationsCallbackParticipantReplaceParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the post conversations callback participant replace params
func (o *PostConversationsCallbackParticipantReplaceParams) WithParticipantID(participantID string) *PostConversationsCallbackParticipantReplaceParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the post conversations callback participant replace params
func (o *PostConversationsCallbackParticipantReplaceParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationsCallbackParticipantReplaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
