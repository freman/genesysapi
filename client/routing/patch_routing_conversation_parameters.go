// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewPatchRoutingConversationParams creates a new PatchRoutingConversationParams object
// with the default values initialized.
func NewPatchRoutingConversationParams() *PatchRoutingConversationParams {
	var ()
	return &PatchRoutingConversationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchRoutingConversationParamsWithTimeout creates a new PatchRoutingConversationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchRoutingConversationParamsWithTimeout(timeout time.Duration) *PatchRoutingConversationParams {
	var ()
	return &PatchRoutingConversationParams{

		timeout: timeout,
	}
}

// NewPatchRoutingConversationParamsWithContext creates a new PatchRoutingConversationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchRoutingConversationParamsWithContext(ctx context.Context) *PatchRoutingConversationParams {
	var ()
	return &PatchRoutingConversationParams{

		Context: ctx,
	}
}

// NewPatchRoutingConversationParamsWithHTTPClient creates a new PatchRoutingConversationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchRoutingConversationParamsWithHTTPClient(client *http.Client) *PatchRoutingConversationParams {
	var ()
	return &PatchRoutingConversationParams{
		HTTPClient: client,
	}
}

/*PatchRoutingConversationParams contains all the parameters to send to the API endpoint
for the patch routing conversation operation typically these are written to a http.Request
*/
type PatchRoutingConversationParams struct {

	/*Body
	  Conversation Attributes

	*/
	Body *models.RoutingConversationAttributesRequest
	/*ConversationID
	  Conversation ID

	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch routing conversation params
func (o *PatchRoutingConversationParams) WithTimeout(timeout time.Duration) *PatchRoutingConversationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch routing conversation params
func (o *PatchRoutingConversationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch routing conversation params
func (o *PatchRoutingConversationParams) WithContext(ctx context.Context) *PatchRoutingConversationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch routing conversation params
func (o *PatchRoutingConversationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch routing conversation params
func (o *PatchRoutingConversationParams) WithHTTPClient(client *http.Client) *PatchRoutingConversationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch routing conversation params
func (o *PatchRoutingConversationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch routing conversation params
func (o *PatchRoutingConversationParams) WithBody(body *models.RoutingConversationAttributesRequest) *PatchRoutingConversationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch routing conversation params
func (o *PatchRoutingConversationParams) SetBody(body *models.RoutingConversationAttributesRequest) {
	o.Body = body
}

// WithConversationID adds the conversationID to the patch routing conversation params
func (o *PatchRoutingConversationParams) WithConversationID(conversationID string) *PatchRoutingConversationParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch routing conversation params
func (o *PatchRoutingConversationParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchRoutingConversationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
