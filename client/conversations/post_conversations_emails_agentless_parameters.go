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

// NewPostConversationsEmailsAgentlessParams creates a new PostConversationsEmailsAgentlessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostConversationsEmailsAgentlessParams() *PostConversationsEmailsAgentlessParams {
	return &PostConversationsEmailsAgentlessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationsEmailsAgentlessParamsWithTimeout creates a new PostConversationsEmailsAgentlessParams object
// with the ability to set a timeout on a request.
func NewPostConversationsEmailsAgentlessParamsWithTimeout(timeout time.Duration) *PostConversationsEmailsAgentlessParams {
	return &PostConversationsEmailsAgentlessParams{
		timeout: timeout,
	}
}

// NewPostConversationsEmailsAgentlessParamsWithContext creates a new PostConversationsEmailsAgentlessParams object
// with the ability to set a context for a request.
func NewPostConversationsEmailsAgentlessParamsWithContext(ctx context.Context) *PostConversationsEmailsAgentlessParams {
	return &PostConversationsEmailsAgentlessParams{
		Context: ctx,
	}
}

// NewPostConversationsEmailsAgentlessParamsWithHTTPClient creates a new PostConversationsEmailsAgentlessParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostConversationsEmailsAgentlessParamsWithHTTPClient(client *http.Client) *PostConversationsEmailsAgentlessParams {
	return &PostConversationsEmailsAgentlessParams{
		HTTPClient: client,
	}
}

/*
PostConversationsEmailsAgentlessParams contains all the parameters to send to the API endpoint

	for the post conversations emails agentless operation.

	Typically these are written to a http.Request.
*/
type PostConversationsEmailsAgentlessParams struct {

	/* Body.

	   Create agentless email request
	*/
	Body *models.AgentlessEmailSendRequestDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post conversations emails agentless params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationsEmailsAgentlessParams) WithDefaults() *PostConversationsEmailsAgentlessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post conversations emails agentless params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationsEmailsAgentlessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post conversations emails agentless params
func (o *PostConversationsEmailsAgentlessParams) WithTimeout(timeout time.Duration) *PostConversationsEmailsAgentlessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversations emails agentless params
func (o *PostConversationsEmailsAgentlessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversations emails agentless params
func (o *PostConversationsEmailsAgentlessParams) WithContext(ctx context.Context) *PostConversationsEmailsAgentlessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversations emails agentless params
func (o *PostConversationsEmailsAgentlessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversations emails agentless params
func (o *PostConversationsEmailsAgentlessParams) WithHTTPClient(client *http.Client) *PostConversationsEmailsAgentlessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversations emails agentless params
func (o *PostConversationsEmailsAgentlessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversations emails agentless params
func (o *PostConversationsEmailsAgentlessParams) WithBody(body *models.AgentlessEmailSendRequestDto) *PostConversationsEmailsAgentlessParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversations emails agentless params
func (o *PostConversationsEmailsAgentlessParams) SetBody(body *models.AgentlessEmailSendRequestDto) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationsEmailsAgentlessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
