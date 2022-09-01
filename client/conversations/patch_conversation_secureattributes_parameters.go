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

// NewPatchConversationSecureattributesParams creates a new PatchConversationSecureattributesParams object
// with the default values initialized.
func NewPatchConversationSecureattributesParams() *PatchConversationSecureattributesParams {
	var ()
	return &PatchConversationSecureattributesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConversationSecureattributesParamsWithTimeout creates a new PatchConversationSecureattributesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchConversationSecureattributesParamsWithTimeout(timeout time.Duration) *PatchConversationSecureattributesParams {
	var ()
	return &PatchConversationSecureattributesParams{

		timeout: timeout,
	}
}

// NewPatchConversationSecureattributesParamsWithContext creates a new PatchConversationSecureattributesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchConversationSecureattributesParamsWithContext(ctx context.Context) *PatchConversationSecureattributesParams {
	var ()
	return &PatchConversationSecureattributesParams{

		Context: ctx,
	}
}

// NewPatchConversationSecureattributesParamsWithHTTPClient creates a new PatchConversationSecureattributesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchConversationSecureattributesParamsWithHTTPClient(client *http.Client) *PatchConversationSecureattributesParams {
	var ()
	return &PatchConversationSecureattributesParams{
		HTTPClient: client,
	}
}

/*PatchConversationSecureattributesParams contains all the parameters to send to the API endpoint
for the patch conversation secureattributes operation typically these are written to a http.Request
*/
type PatchConversationSecureattributesParams struct {

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

// WithTimeout adds the timeout to the patch conversation secureattributes params
func (o *PatchConversationSecureattributesParams) WithTimeout(timeout time.Duration) *PatchConversationSecureattributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch conversation secureattributes params
func (o *PatchConversationSecureattributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch conversation secureattributes params
func (o *PatchConversationSecureattributesParams) WithContext(ctx context.Context) *PatchConversationSecureattributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch conversation secureattributes params
func (o *PatchConversationSecureattributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch conversation secureattributes params
func (o *PatchConversationSecureattributesParams) WithHTTPClient(client *http.Client) *PatchConversationSecureattributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch conversation secureattributes params
func (o *PatchConversationSecureattributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch conversation secureattributes params
func (o *PatchConversationSecureattributesParams) WithBody(body *models.ConversationSecureAttributes) *PatchConversationSecureattributesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch conversation secureattributes params
func (o *PatchConversationSecureattributesParams) SetBody(body *models.ConversationSecureAttributes) {
	o.Body = body
}

// WithConversationID adds the conversationID to the patch conversation secureattributes params
func (o *PatchConversationSecureattributesParams) WithConversationID(conversationID string) *PatchConversationSecureattributesParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch conversation secureattributes params
func (o *PatchConversationSecureattributesParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConversationSecureattributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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