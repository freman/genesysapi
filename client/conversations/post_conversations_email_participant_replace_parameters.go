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

// NewPostConversationsEmailParticipantReplaceParams creates a new PostConversationsEmailParticipantReplaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostConversationsEmailParticipantReplaceParams() *PostConversationsEmailParticipantReplaceParams {
	return &PostConversationsEmailParticipantReplaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationsEmailParticipantReplaceParamsWithTimeout creates a new PostConversationsEmailParticipantReplaceParams object
// with the ability to set a timeout on a request.
func NewPostConversationsEmailParticipantReplaceParamsWithTimeout(timeout time.Duration) *PostConversationsEmailParticipantReplaceParams {
	return &PostConversationsEmailParticipantReplaceParams{
		timeout: timeout,
	}
}

// NewPostConversationsEmailParticipantReplaceParamsWithContext creates a new PostConversationsEmailParticipantReplaceParams object
// with the ability to set a context for a request.
func NewPostConversationsEmailParticipantReplaceParamsWithContext(ctx context.Context) *PostConversationsEmailParticipantReplaceParams {
	return &PostConversationsEmailParticipantReplaceParams{
		Context: ctx,
	}
}

// NewPostConversationsEmailParticipantReplaceParamsWithHTTPClient creates a new PostConversationsEmailParticipantReplaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostConversationsEmailParticipantReplaceParamsWithHTTPClient(client *http.Client) *PostConversationsEmailParticipantReplaceParams {
	return &PostConversationsEmailParticipantReplaceParams{
		HTTPClient: client,
	}
}

/*
PostConversationsEmailParticipantReplaceParams contains all the parameters to send to the API endpoint

	for the post conversations email participant replace operation.

	Typically these are written to a http.Request.
*/
type PostConversationsEmailParticipantReplaceParams struct {

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

// WithDefaults hydrates default values in the post conversations email participant replace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationsEmailParticipantReplaceParams) WithDefaults() *PostConversationsEmailParticipantReplaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post conversations email participant replace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationsEmailParticipantReplaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post conversations email participant replace params
func (o *PostConversationsEmailParticipantReplaceParams) WithTimeout(timeout time.Duration) *PostConversationsEmailParticipantReplaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversations email participant replace params
func (o *PostConversationsEmailParticipantReplaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversations email participant replace params
func (o *PostConversationsEmailParticipantReplaceParams) WithContext(ctx context.Context) *PostConversationsEmailParticipantReplaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversations email participant replace params
func (o *PostConversationsEmailParticipantReplaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversations email participant replace params
func (o *PostConversationsEmailParticipantReplaceParams) WithHTTPClient(client *http.Client) *PostConversationsEmailParticipantReplaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversations email participant replace params
func (o *PostConversationsEmailParticipantReplaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversations email participant replace params
func (o *PostConversationsEmailParticipantReplaceParams) WithBody(body *models.TransferRequest) *PostConversationsEmailParticipantReplaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversations email participant replace params
func (o *PostConversationsEmailParticipantReplaceParams) SetBody(body *models.TransferRequest) {
	o.Body = body
}

// WithConversationID adds the conversationID to the post conversations email participant replace params
func (o *PostConversationsEmailParticipantReplaceParams) WithConversationID(conversationID string) *PostConversationsEmailParticipantReplaceParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the post conversations email participant replace params
func (o *PostConversationsEmailParticipantReplaceParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the post conversations email participant replace params
func (o *PostConversationsEmailParticipantReplaceParams) WithParticipantID(participantID string) *PostConversationsEmailParticipantReplaceParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the post conversations email participant replace params
func (o *PostConversationsEmailParticipantReplaceParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationsEmailParticipantReplaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
