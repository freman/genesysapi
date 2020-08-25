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

// NewPatchConversationsCobrowsesessionParams creates a new PatchConversationsCobrowsesessionParams object
// with the default values initialized.
func NewPatchConversationsCobrowsesessionParams() *PatchConversationsCobrowsesessionParams {
	var ()
	return &PatchConversationsCobrowsesessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConversationsCobrowsesessionParamsWithTimeout creates a new PatchConversationsCobrowsesessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchConversationsCobrowsesessionParamsWithTimeout(timeout time.Duration) *PatchConversationsCobrowsesessionParams {
	var ()
	return &PatchConversationsCobrowsesessionParams{

		timeout: timeout,
	}
}

// NewPatchConversationsCobrowsesessionParamsWithContext creates a new PatchConversationsCobrowsesessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchConversationsCobrowsesessionParamsWithContext(ctx context.Context) *PatchConversationsCobrowsesessionParams {
	var ()
	return &PatchConversationsCobrowsesessionParams{

		Context: ctx,
	}
}

// NewPatchConversationsCobrowsesessionParamsWithHTTPClient creates a new PatchConversationsCobrowsesessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchConversationsCobrowsesessionParamsWithHTTPClient(client *http.Client) *PatchConversationsCobrowsesessionParams {
	var ()
	return &PatchConversationsCobrowsesessionParams{
		HTTPClient: client,
	}
}

/*PatchConversationsCobrowsesessionParams contains all the parameters to send to the API endpoint
for the patch conversations cobrowsesession operation typically these are written to a http.Request
*/
type PatchConversationsCobrowsesessionParams struct {

	/*Body
	  Conversation

	*/
	Body *models.Conversation
	/*ConversationID
	  conversationId

	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch conversations cobrowsesession params
func (o *PatchConversationsCobrowsesessionParams) WithTimeout(timeout time.Duration) *PatchConversationsCobrowsesessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch conversations cobrowsesession params
func (o *PatchConversationsCobrowsesessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch conversations cobrowsesession params
func (o *PatchConversationsCobrowsesessionParams) WithContext(ctx context.Context) *PatchConversationsCobrowsesessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch conversations cobrowsesession params
func (o *PatchConversationsCobrowsesessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch conversations cobrowsesession params
func (o *PatchConversationsCobrowsesessionParams) WithHTTPClient(client *http.Client) *PatchConversationsCobrowsesessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch conversations cobrowsesession params
func (o *PatchConversationsCobrowsesessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch conversations cobrowsesession params
func (o *PatchConversationsCobrowsesessionParams) WithBody(body *models.Conversation) *PatchConversationsCobrowsesessionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch conversations cobrowsesession params
func (o *PatchConversationsCobrowsesessionParams) SetBody(body *models.Conversation) {
	o.Body = body
}

// WithConversationID adds the conversationID to the patch conversations cobrowsesession params
func (o *PatchConversationsCobrowsesessionParams) WithConversationID(conversationID string) *PatchConversationsCobrowsesessionParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch conversations cobrowsesession params
func (o *PatchConversationsCobrowsesessionParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConversationsCobrowsesessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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