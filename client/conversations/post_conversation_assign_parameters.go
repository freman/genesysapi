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

// NewPostConversationAssignParams creates a new PostConversationAssignParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostConversationAssignParams() *PostConversationAssignParams {
	return &PostConversationAssignParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationAssignParamsWithTimeout creates a new PostConversationAssignParams object
// with the ability to set a timeout on a request.
func NewPostConversationAssignParamsWithTimeout(timeout time.Duration) *PostConversationAssignParams {
	return &PostConversationAssignParams{
		timeout: timeout,
	}
}

// NewPostConversationAssignParamsWithContext creates a new PostConversationAssignParams object
// with the ability to set a context for a request.
func NewPostConversationAssignParamsWithContext(ctx context.Context) *PostConversationAssignParams {
	return &PostConversationAssignParams{
		Context: ctx,
	}
}

// NewPostConversationAssignParamsWithHTTPClient creates a new PostConversationAssignParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostConversationAssignParamsWithHTTPClient(client *http.Client) *PostConversationAssignParams {
	return &PostConversationAssignParams{
		HTTPClient: client,
	}
}

/*
PostConversationAssignParams contains all the parameters to send to the API endpoint

	for the post conversation assign operation.

	Typically these are written to a http.Request.
*/
type PostConversationAssignParams struct {

	/* Body.

	   Targeted user
	*/
	Body *models.ConversationUser

	/* ConversationID.

	   conversation ID
	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post conversation assign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationAssignParams) WithDefaults() *PostConversationAssignParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post conversation assign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationAssignParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post conversation assign params
func (o *PostConversationAssignParams) WithTimeout(timeout time.Duration) *PostConversationAssignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversation assign params
func (o *PostConversationAssignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversation assign params
func (o *PostConversationAssignParams) WithContext(ctx context.Context) *PostConversationAssignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversation assign params
func (o *PostConversationAssignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversation assign params
func (o *PostConversationAssignParams) WithHTTPClient(client *http.Client) *PostConversationAssignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversation assign params
func (o *PostConversationAssignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversation assign params
func (o *PostConversationAssignParams) WithBody(body *models.ConversationUser) *PostConversationAssignParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversation assign params
func (o *PostConversationAssignParams) SetBody(body *models.ConversationUser) {
	o.Body = body
}

// WithConversationID adds the conversationID to the post conversation assign params
func (o *PostConversationAssignParams) WithConversationID(conversationID string) *PostConversationAssignParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the post conversation assign params
func (o *PostConversationAssignParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationAssignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
