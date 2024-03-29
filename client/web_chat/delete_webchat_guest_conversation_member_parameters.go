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
)

// NewDeleteWebchatGuestConversationMemberParams creates a new DeleteWebchatGuestConversationMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteWebchatGuestConversationMemberParams() *DeleteWebchatGuestConversationMemberParams {
	return &DeleteWebchatGuestConversationMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWebchatGuestConversationMemberParamsWithTimeout creates a new DeleteWebchatGuestConversationMemberParams object
// with the ability to set a timeout on a request.
func NewDeleteWebchatGuestConversationMemberParamsWithTimeout(timeout time.Duration) *DeleteWebchatGuestConversationMemberParams {
	return &DeleteWebchatGuestConversationMemberParams{
		timeout: timeout,
	}
}

// NewDeleteWebchatGuestConversationMemberParamsWithContext creates a new DeleteWebchatGuestConversationMemberParams object
// with the ability to set a context for a request.
func NewDeleteWebchatGuestConversationMemberParamsWithContext(ctx context.Context) *DeleteWebchatGuestConversationMemberParams {
	return &DeleteWebchatGuestConversationMemberParams{
		Context: ctx,
	}
}

// NewDeleteWebchatGuestConversationMemberParamsWithHTTPClient creates a new DeleteWebchatGuestConversationMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteWebchatGuestConversationMemberParamsWithHTTPClient(client *http.Client) *DeleteWebchatGuestConversationMemberParams {
	return &DeleteWebchatGuestConversationMemberParams{
		HTTPClient: client,
	}
}

/*
DeleteWebchatGuestConversationMemberParams contains all the parameters to send to the API endpoint

	for the delete webchat guest conversation member operation.

	Typically these are written to a http.Request.
*/
type DeleteWebchatGuestConversationMemberParams struct {

	/* ConversationID.

	   conversationId
	*/
	ConversationID string

	/* MemberID.

	   memberId
	*/
	MemberID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete webchat guest conversation member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWebchatGuestConversationMemberParams) WithDefaults() *DeleteWebchatGuestConversationMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete webchat guest conversation member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWebchatGuestConversationMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete webchat guest conversation member params
func (o *DeleteWebchatGuestConversationMemberParams) WithTimeout(timeout time.Duration) *DeleteWebchatGuestConversationMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete webchat guest conversation member params
func (o *DeleteWebchatGuestConversationMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete webchat guest conversation member params
func (o *DeleteWebchatGuestConversationMemberParams) WithContext(ctx context.Context) *DeleteWebchatGuestConversationMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete webchat guest conversation member params
func (o *DeleteWebchatGuestConversationMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete webchat guest conversation member params
func (o *DeleteWebchatGuestConversationMemberParams) WithHTTPClient(client *http.Client) *DeleteWebchatGuestConversationMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete webchat guest conversation member params
func (o *DeleteWebchatGuestConversationMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the delete webchat guest conversation member params
func (o *DeleteWebchatGuestConversationMemberParams) WithConversationID(conversationID string) *DeleteWebchatGuestConversationMemberParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the delete webchat guest conversation member params
func (o *DeleteWebchatGuestConversationMemberParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithMemberID adds the memberID to the delete webchat guest conversation member params
func (o *DeleteWebchatGuestConversationMemberParams) WithMemberID(memberID string) *DeleteWebchatGuestConversationMemberParams {
	o.SetMemberID(memberID)
	return o
}

// SetMemberID adds the memberId to the delete webchat guest conversation member params
func (o *DeleteWebchatGuestConversationMemberParams) SetMemberID(memberID string) {
	o.MemberID = memberID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWebchatGuestConversationMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param conversationId
	if err := r.SetPathParam("conversationId", o.ConversationID); err != nil {
		return err
	}

	// path param memberId
	if err := r.SetPathParam("memberId", o.MemberID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
