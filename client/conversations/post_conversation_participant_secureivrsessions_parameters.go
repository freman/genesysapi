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

// NewPostConversationParticipantSecureivrsessionsParams creates a new PostConversationParticipantSecureivrsessionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostConversationParticipantSecureivrsessionsParams() *PostConversationParticipantSecureivrsessionsParams {
	return &PostConversationParticipantSecureivrsessionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationParticipantSecureivrsessionsParamsWithTimeout creates a new PostConversationParticipantSecureivrsessionsParams object
// with the ability to set a timeout on a request.
func NewPostConversationParticipantSecureivrsessionsParamsWithTimeout(timeout time.Duration) *PostConversationParticipantSecureivrsessionsParams {
	return &PostConversationParticipantSecureivrsessionsParams{
		timeout: timeout,
	}
}

// NewPostConversationParticipantSecureivrsessionsParamsWithContext creates a new PostConversationParticipantSecureivrsessionsParams object
// with the ability to set a context for a request.
func NewPostConversationParticipantSecureivrsessionsParamsWithContext(ctx context.Context) *PostConversationParticipantSecureivrsessionsParams {
	return &PostConversationParticipantSecureivrsessionsParams{
		Context: ctx,
	}
}

// NewPostConversationParticipantSecureivrsessionsParamsWithHTTPClient creates a new PostConversationParticipantSecureivrsessionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostConversationParticipantSecureivrsessionsParamsWithHTTPClient(client *http.Client) *PostConversationParticipantSecureivrsessionsParams {
	return &PostConversationParticipantSecureivrsessionsParams{
		HTTPClient: client,
	}
}

/*
PostConversationParticipantSecureivrsessionsParams contains all the parameters to send to the API endpoint

	for the post conversation participant secureivrsessions operation.

	Typically these are written to a http.Request.
*/
type PostConversationParticipantSecureivrsessionsParams struct {

	// Body.
	Body *models.CreateSecureSession

	/* ConversationID.

	   conversation ID
	*/
	ConversationID string

	/* ParticipantID.

	   participant ID
	*/
	ParticipantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post conversation participant secureivrsessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationParticipantSecureivrsessionsParams) WithDefaults() *PostConversationParticipantSecureivrsessionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post conversation participant secureivrsessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationParticipantSecureivrsessionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post conversation participant secureivrsessions params
func (o *PostConversationParticipantSecureivrsessionsParams) WithTimeout(timeout time.Duration) *PostConversationParticipantSecureivrsessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversation participant secureivrsessions params
func (o *PostConversationParticipantSecureivrsessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversation participant secureivrsessions params
func (o *PostConversationParticipantSecureivrsessionsParams) WithContext(ctx context.Context) *PostConversationParticipantSecureivrsessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversation participant secureivrsessions params
func (o *PostConversationParticipantSecureivrsessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversation participant secureivrsessions params
func (o *PostConversationParticipantSecureivrsessionsParams) WithHTTPClient(client *http.Client) *PostConversationParticipantSecureivrsessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversation participant secureivrsessions params
func (o *PostConversationParticipantSecureivrsessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversation participant secureivrsessions params
func (o *PostConversationParticipantSecureivrsessionsParams) WithBody(body *models.CreateSecureSession) *PostConversationParticipantSecureivrsessionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversation participant secureivrsessions params
func (o *PostConversationParticipantSecureivrsessionsParams) SetBody(body *models.CreateSecureSession) {
	o.Body = body
}

// WithConversationID adds the conversationID to the post conversation participant secureivrsessions params
func (o *PostConversationParticipantSecureivrsessionsParams) WithConversationID(conversationID string) *PostConversationParticipantSecureivrsessionsParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the post conversation participant secureivrsessions params
func (o *PostConversationParticipantSecureivrsessionsParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the post conversation participant secureivrsessions params
func (o *PostConversationParticipantSecureivrsessionsParams) WithParticipantID(participantID string) *PostConversationParticipantSecureivrsessionsParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the post conversation participant secureivrsessions params
func (o *PostConversationParticipantSecureivrsessionsParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationParticipantSecureivrsessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
