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

// NewDeleteConversationsEmailMessagesDraftAttachmentParams creates a new DeleteConversationsEmailMessagesDraftAttachmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteConversationsEmailMessagesDraftAttachmentParams() *DeleteConversationsEmailMessagesDraftAttachmentParams {
	return &DeleteConversationsEmailMessagesDraftAttachmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConversationsEmailMessagesDraftAttachmentParamsWithTimeout creates a new DeleteConversationsEmailMessagesDraftAttachmentParams object
// with the ability to set a timeout on a request.
func NewDeleteConversationsEmailMessagesDraftAttachmentParamsWithTimeout(timeout time.Duration) *DeleteConversationsEmailMessagesDraftAttachmentParams {
	return &DeleteConversationsEmailMessagesDraftAttachmentParams{
		timeout: timeout,
	}
}

// NewDeleteConversationsEmailMessagesDraftAttachmentParamsWithContext creates a new DeleteConversationsEmailMessagesDraftAttachmentParams object
// with the ability to set a context for a request.
func NewDeleteConversationsEmailMessagesDraftAttachmentParamsWithContext(ctx context.Context) *DeleteConversationsEmailMessagesDraftAttachmentParams {
	return &DeleteConversationsEmailMessagesDraftAttachmentParams{
		Context: ctx,
	}
}

// NewDeleteConversationsEmailMessagesDraftAttachmentParamsWithHTTPClient creates a new DeleteConversationsEmailMessagesDraftAttachmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteConversationsEmailMessagesDraftAttachmentParamsWithHTTPClient(client *http.Client) *DeleteConversationsEmailMessagesDraftAttachmentParams {
	return &DeleteConversationsEmailMessagesDraftAttachmentParams{
		HTTPClient: client,
	}
}

/*
DeleteConversationsEmailMessagesDraftAttachmentParams contains all the parameters to send to the API endpoint

	for the delete conversations email messages draft attachment operation.

	Typically these are written to a http.Request.
*/
type DeleteConversationsEmailMessagesDraftAttachmentParams struct {

	/* AttachmentID.

	   attachmentId
	*/
	AttachmentID string

	/* ConversationID.

	   conversationId
	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete conversations email messages draft attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) WithDefaults() *DeleteConversationsEmailMessagesDraftAttachmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete conversations email messages draft attachment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete conversations email messages draft attachment params
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) WithTimeout(timeout time.Duration) *DeleteConversationsEmailMessagesDraftAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete conversations email messages draft attachment params
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete conversations email messages draft attachment params
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) WithContext(ctx context.Context) *DeleteConversationsEmailMessagesDraftAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete conversations email messages draft attachment params
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete conversations email messages draft attachment params
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) WithHTTPClient(client *http.Client) *DeleteConversationsEmailMessagesDraftAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete conversations email messages draft attachment params
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttachmentID adds the attachmentID to the delete conversations email messages draft attachment params
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) WithAttachmentID(attachmentID string) *DeleteConversationsEmailMessagesDraftAttachmentParams {
	o.SetAttachmentID(attachmentID)
	return o
}

// SetAttachmentID adds the attachmentId to the delete conversations email messages draft attachment params
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) SetAttachmentID(attachmentID string) {
	o.AttachmentID = attachmentID
}

// WithConversationID adds the conversationID to the delete conversations email messages draft attachment params
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) WithConversationID(conversationID string) *DeleteConversationsEmailMessagesDraftAttachmentParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the delete conversations email messages draft attachment params
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConversationsEmailMessagesDraftAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attachmentId
	if err := r.SetPathParam("attachmentId", o.AttachmentID); err != nil {
		return err
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
