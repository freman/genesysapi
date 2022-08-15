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

// NewPutConversationSecureattributesParams creates a new PutConversationSecureattributesParams object
// with the default values initialized.
func NewPutConversationSecureattributesParams() *PutConversationSecureattributesParams {
	var ()
	return &PutConversationSecureattributesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutConversationSecureattributesParamsWithTimeout creates a new PutConversationSecureattributesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutConversationSecureattributesParamsWithTimeout(timeout time.Duration) *PutConversationSecureattributesParams {
	var ()
	return &PutConversationSecureattributesParams{

		timeout: timeout,
	}
}

// NewPutConversationSecureattributesParamsWithContext creates a new PutConversationSecureattributesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutConversationSecureattributesParamsWithContext(ctx context.Context) *PutConversationSecureattributesParams {
	var ()
	return &PutConversationSecureattributesParams{

		Context: ctx,
	}
}

// NewPutConversationSecureattributesParamsWithHTTPClient creates a new PutConversationSecureattributesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutConversationSecureattributesParamsWithHTTPClient(client *http.Client) *PutConversationSecureattributesParams {
	var ()
	return &PutConversationSecureattributesParams{
		HTTPClient: client,
	}
}

/*PutConversationSecureattributesParams contains all the parameters to send to the API endpoint
for the put conversation secureattributes operation typically these are written to a http.Request
*/
type PutConversationSecureattributesParams struct {

	/*Body
	  Conversation Secure Attributes

	*/
	Body *models.ConversationSecureAttributes
	/*ConversationID
	  conversation ID

	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put conversation secureattributes params
func (o *PutConversationSecureattributesParams) WithTimeout(timeout time.Duration) *PutConversationSecureattributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put conversation secureattributes params
func (o *PutConversationSecureattributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put conversation secureattributes params
func (o *PutConversationSecureattributesParams) WithContext(ctx context.Context) *PutConversationSecureattributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put conversation secureattributes params
func (o *PutConversationSecureattributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put conversation secureattributes params
func (o *PutConversationSecureattributesParams) WithHTTPClient(client *http.Client) *PutConversationSecureattributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put conversation secureattributes params
func (o *PutConversationSecureattributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put conversation secureattributes params
func (o *PutConversationSecureattributesParams) WithBody(body *models.ConversationSecureAttributes) *PutConversationSecureattributesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put conversation secureattributes params
func (o *PutConversationSecureattributesParams) SetBody(body *models.ConversationSecureAttributes) {
	o.Body = body
}

// WithConversationID adds the conversationID to the put conversation secureattributes params
func (o *PutConversationSecureattributesParams) WithConversationID(conversationID string) *PutConversationSecureattributesParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the put conversation secureattributes params
func (o *PutConversationSecureattributesParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PutConversationSecureattributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
