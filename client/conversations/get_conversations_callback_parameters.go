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

// NewGetConversationsCallbackParams creates a new GetConversationsCallbackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConversationsCallbackParams() *GetConversationsCallbackParams {
	return &GetConversationsCallbackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsCallbackParamsWithTimeout creates a new GetConversationsCallbackParams object
// with the ability to set a timeout on a request.
func NewGetConversationsCallbackParamsWithTimeout(timeout time.Duration) *GetConversationsCallbackParams {
	return &GetConversationsCallbackParams{
		timeout: timeout,
	}
}

// NewGetConversationsCallbackParamsWithContext creates a new GetConversationsCallbackParams object
// with the ability to set a context for a request.
func NewGetConversationsCallbackParamsWithContext(ctx context.Context) *GetConversationsCallbackParams {
	return &GetConversationsCallbackParams{
		Context: ctx,
	}
}

// NewGetConversationsCallbackParamsWithHTTPClient creates a new GetConversationsCallbackParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConversationsCallbackParamsWithHTTPClient(client *http.Client) *GetConversationsCallbackParams {
	return &GetConversationsCallbackParams{
		HTTPClient: client,
	}
}

/*
GetConversationsCallbackParams contains all the parameters to send to the API endpoint

	for the get conversations callback operation.

	Typically these are written to a http.Request.
*/
type GetConversationsCallbackParams struct {

	/* ConversationID.

	   conversationId
	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get conversations callback params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsCallbackParams) WithDefaults() *GetConversationsCallbackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get conversations callback params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsCallbackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get conversations callback params
func (o *GetConversationsCallbackParams) WithTimeout(timeout time.Duration) *GetConversationsCallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations callback params
func (o *GetConversationsCallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations callback params
func (o *GetConversationsCallbackParams) WithContext(ctx context.Context) *GetConversationsCallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations callback params
func (o *GetConversationsCallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations callback params
func (o *GetConversationsCallbackParams) WithHTTPClient(client *http.Client) *GetConversationsCallbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations callback params
func (o *GetConversationsCallbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the get conversations callback params
func (o *GetConversationsCallbackParams) WithConversationID(conversationID string) *GetConversationsCallbackParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get conversations callback params
func (o *GetConversationsCallbackParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsCallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param conversationId
	if err := r.SetPathParam("conversationId", o.ConversationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
