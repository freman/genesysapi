// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

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

// NewPutExternalcontactsConversationParams creates a new PutExternalcontactsConversationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutExternalcontactsConversationParams() *PutExternalcontactsConversationParams {
	return &PutExternalcontactsConversationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutExternalcontactsConversationParamsWithTimeout creates a new PutExternalcontactsConversationParams object
// with the ability to set a timeout on a request.
func NewPutExternalcontactsConversationParamsWithTimeout(timeout time.Duration) *PutExternalcontactsConversationParams {
	return &PutExternalcontactsConversationParams{
		timeout: timeout,
	}
}

// NewPutExternalcontactsConversationParamsWithContext creates a new PutExternalcontactsConversationParams object
// with the ability to set a context for a request.
func NewPutExternalcontactsConversationParamsWithContext(ctx context.Context) *PutExternalcontactsConversationParams {
	return &PutExternalcontactsConversationParams{
		Context: ctx,
	}
}

// NewPutExternalcontactsConversationParamsWithHTTPClient creates a new PutExternalcontactsConversationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutExternalcontactsConversationParamsWithHTTPClient(client *http.Client) *PutExternalcontactsConversationParams {
	return &PutExternalcontactsConversationParams{
		HTTPClient: client,
	}
}

/*
PutExternalcontactsConversationParams contains all the parameters to send to the API endpoint

	for the put externalcontacts conversation operation.

	Typically these are written to a http.Request.
*/
type PutExternalcontactsConversationParams struct {

	/* Body.

	   ConversationAssociation
	*/
	Body *models.ConversationAssociation

	/* ConversationID.

	   Conversation ID
	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put externalcontacts conversation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutExternalcontactsConversationParams) WithDefaults() *PutExternalcontactsConversationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put externalcontacts conversation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutExternalcontactsConversationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put externalcontacts conversation params
func (o *PutExternalcontactsConversationParams) WithTimeout(timeout time.Duration) *PutExternalcontactsConversationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put externalcontacts conversation params
func (o *PutExternalcontactsConversationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put externalcontacts conversation params
func (o *PutExternalcontactsConversationParams) WithContext(ctx context.Context) *PutExternalcontactsConversationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put externalcontacts conversation params
func (o *PutExternalcontactsConversationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put externalcontacts conversation params
func (o *PutExternalcontactsConversationParams) WithHTTPClient(client *http.Client) *PutExternalcontactsConversationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put externalcontacts conversation params
func (o *PutExternalcontactsConversationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put externalcontacts conversation params
func (o *PutExternalcontactsConversationParams) WithBody(body *models.ConversationAssociation) *PutExternalcontactsConversationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put externalcontacts conversation params
func (o *PutExternalcontactsConversationParams) SetBody(body *models.ConversationAssociation) {
	o.Body = body
}

// WithConversationID adds the conversationID to the put externalcontacts conversation params
func (o *PutExternalcontactsConversationParams) WithConversationID(conversationID string) *PutExternalcontactsConversationParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the put externalcontacts conversation params
func (o *PutExternalcontactsConversationParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PutExternalcontactsConversationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
