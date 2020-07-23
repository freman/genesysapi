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

// NewPatchConversationsCobrowsesessionParticipantParams creates a new PatchConversationsCobrowsesessionParticipantParams object
// with the default values initialized.
func NewPatchConversationsCobrowsesessionParticipantParams() *PatchConversationsCobrowsesessionParticipantParams {
	var ()
	return &PatchConversationsCobrowsesessionParticipantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConversationsCobrowsesessionParticipantParamsWithTimeout creates a new PatchConversationsCobrowsesessionParticipantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchConversationsCobrowsesessionParticipantParamsWithTimeout(timeout time.Duration) *PatchConversationsCobrowsesessionParticipantParams {
	var ()
	return &PatchConversationsCobrowsesessionParticipantParams{

		timeout: timeout,
	}
}

// NewPatchConversationsCobrowsesessionParticipantParamsWithContext creates a new PatchConversationsCobrowsesessionParticipantParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchConversationsCobrowsesessionParticipantParamsWithContext(ctx context.Context) *PatchConversationsCobrowsesessionParticipantParams {
	var ()
	return &PatchConversationsCobrowsesessionParticipantParams{

		Context: ctx,
	}
}

// NewPatchConversationsCobrowsesessionParticipantParamsWithHTTPClient creates a new PatchConversationsCobrowsesessionParticipantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchConversationsCobrowsesessionParticipantParamsWithHTTPClient(client *http.Client) *PatchConversationsCobrowsesessionParticipantParams {
	var ()
	return &PatchConversationsCobrowsesessionParticipantParams{
		HTTPClient: client,
	}
}

/*PatchConversationsCobrowsesessionParticipantParams contains all the parameters to send to the API endpoint
for the patch conversations cobrowsesession participant operation typically these are written to a http.Request
*/
type PatchConversationsCobrowsesessionParticipantParams struct {

	/*Body*/
	Body *models.MediaParticipantRequest
	/*ConversationID
	  conversationId

	*/
	ConversationID string
	/*ParticipantID
	  participantId

	*/
	ParticipantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch conversations cobrowsesession participant params
func (o *PatchConversationsCobrowsesessionParticipantParams) WithTimeout(timeout time.Duration) *PatchConversationsCobrowsesessionParticipantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch conversations cobrowsesession participant params
func (o *PatchConversationsCobrowsesessionParticipantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch conversations cobrowsesession participant params
func (o *PatchConversationsCobrowsesessionParticipantParams) WithContext(ctx context.Context) *PatchConversationsCobrowsesessionParticipantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch conversations cobrowsesession participant params
func (o *PatchConversationsCobrowsesessionParticipantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch conversations cobrowsesession participant params
func (o *PatchConversationsCobrowsesessionParticipantParams) WithHTTPClient(client *http.Client) *PatchConversationsCobrowsesessionParticipantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch conversations cobrowsesession participant params
func (o *PatchConversationsCobrowsesessionParticipantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch conversations cobrowsesession participant params
func (o *PatchConversationsCobrowsesessionParticipantParams) WithBody(body *models.MediaParticipantRequest) *PatchConversationsCobrowsesessionParticipantParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch conversations cobrowsesession participant params
func (o *PatchConversationsCobrowsesessionParticipantParams) SetBody(body *models.MediaParticipantRequest) {
	o.Body = body
}

// WithConversationID adds the conversationID to the patch conversations cobrowsesession participant params
func (o *PatchConversationsCobrowsesessionParticipantParams) WithConversationID(conversationID string) *PatchConversationsCobrowsesessionParticipantParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch conversations cobrowsesession participant params
func (o *PatchConversationsCobrowsesessionParticipantParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithParticipantID adds the participantID to the patch conversations cobrowsesession participant params
func (o *PatchConversationsCobrowsesessionParticipantParams) WithParticipantID(participantID string) *PatchConversationsCobrowsesessionParticipantParams {
	o.SetParticipantID(participantID)
	return o
}

// SetParticipantID adds the participantId to the patch conversations cobrowsesession participant params
func (o *PatchConversationsCobrowsesessionParticipantParams) SetParticipantID(participantID string) {
	o.ParticipantID = participantID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConversationsCobrowsesessionParticipantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param participantId
	if err := r.SetPathParam("participantId", o.ParticipantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
