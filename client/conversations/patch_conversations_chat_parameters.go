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

// NewPatchConversationsChatParams creates a new PatchConversationsChatParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchConversationsChatParams() *PatchConversationsChatParams {
	return &PatchConversationsChatParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConversationsChatParamsWithTimeout creates a new PatchConversationsChatParams object
// with the ability to set a timeout on a request.
func NewPatchConversationsChatParamsWithTimeout(timeout time.Duration) *PatchConversationsChatParams {
	return &PatchConversationsChatParams{
		timeout: timeout,
	}
}

// NewPatchConversationsChatParamsWithContext creates a new PatchConversationsChatParams object
// with the ability to set a context for a request.
func NewPatchConversationsChatParamsWithContext(ctx context.Context) *PatchConversationsChatParams {
	return &PatchConversationsChatParams{
		Context: ctx,
	}
}

// NewPatchConversationsChatParamsWithHTTPClient creates a new PatchConversationsChatParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchConversationsChatParamsWithHTTPClient(client *http.Client) *PatchConversationsChatParams {
	return &PatchConversationsChatParams{
		HTTPClient: client,
	}
}

/*
PatchConversationsChatParams contains all the parameters to send to the API endpoint

	for the patch conversations chat operation.

	Typically these are written to a http.Request.
*/
type PatchConversationsChatParams struct {

	/* Body.

	   Conversation
	*/
	Body *models.Conversation

	/* ConversationID.

	   conversationId
	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch conversations chat params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchConversationsChatParams) WithDefaults() *PatchConversationsChatParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch conversations chat params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchConversationsChatParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch conversations chat params
func (o *PatchConversationsChatParams) WithTimeout(timeout time.Duration) *PatchConversationsChatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch conversations chat params
func (o *PatchConversationsChatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch conversations chat params
func (o *PatchConversationsChatParams) WithContext(ctx context.Context) *PatchConversationsChatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch conversations chat params
func (o *PatchConversationsChatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch conversations chat params
func (o *PatchConversationsChatParams) WithHTTPClient(client *http.Client) *PatchConversationsChatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch conversations chat params
func (o *PatchConversationsChatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch conversations chat params
func (o *PatchConversationsChatParams) WithBody(body *models.Conversation) *PatchConversationsChatParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch conversations chat params
func (o *PatchConversationsChatParams) SetBody(body *models.Conversation) {
	o.Body = body
}

// WithConversationID adds the conversationID to the patch conversations chat params
func (o *PatchConversationsChatParams) WithConversationID(conversationID string) *PatchConversationsChatParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch conversations chat params
func (o *PatchConversationsChatParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConversationsChatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
