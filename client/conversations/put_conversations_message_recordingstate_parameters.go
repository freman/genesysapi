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

// NewPutConversationsMessageRecordingstateParams creates a new PutConversationsMessageRecordingstateParams object
// with the default values initialized.
func NewPutConversationsMessageRecordingstateParams() *PutConversationsMessageRecordingstateParams {
	var ()
	return &PutConversationsMessageRecordingstateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutConversationsMessageRecordingstateParamsWithTimeout creates a new PutConversationsMessageRecordingstateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutConversationsMessageRecordingstateParamsWithTimeout(timeout time.Duration) *PutConversationsMessageRecordingstateParams {
	var ()
	return &PutConversationsMessageRecordingstateParams{

		timeout: timeout,
	}
}

// NewPutConversationsMessageRecordingstateParamsWithContext creates a new PutConversationsMessageRecordingstateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutConversationsMessageRecordingstateParamsWithContext(ctx context.Context) *PutConversationsMessageRecordingstateParams {
	var ()
	return &PutConversationsMessageRecordingstateParams{

		Context: ctx,
	}
}

// NewPutConversationsMessageRecordingstateParamsWithHTTPClient creates a new PutConversationsMessageRecordingstateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutConversationsMessageRecordingstateParamsWithHTTPClient(client *http.Client) *PutConversationsMessageRecordingstateParams {
	var ()
	return &PutConversationsMessageRecordingstateParams{
		HTTPClient: client,
	}
}

/*PutConversationsMessageRecordingstateParams contains all the parameters to send to the API endpoint
for the put conversations message recordingstate operation typically these are written to a http.Request
*/
type PutConversationsMessageRecordingstateParams struct {

	/*Body
	  SetRecordingState

	*/
	Body *models.SetRecordingState
	/*ConversationID
	  conversationId

	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put conversations message recordingstate params
func (o *PutConversationsMessageRecordingstateParams) WithTimeout(timeout time.Duration) *PutConversationsMessageRecordingstateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put conversations message recordingstate params
func (o *PutConversationsMessageRecordingstateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put conversations message recordingstate params
func (o *PutConversationsMessageRecordingstateParams) WithContext(ctx context.Context) *PutConversationsMessageRecordingstateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put conversations message recordingstate params
func (o *PutConversationsMessageRecordingstateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put conversations message recordingstate params
func (o *PutConversationsMessageRecordingstateParams) WithHTTPClient(client *http.Client) *PutConversationsMessageRecordingstateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put conversations message recordingstate params
func (o *PutConversationsMessageRecordingstateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put conversations message recordingstate params
func (o *PutConversationsMessageRecordingstateParams) WithBody(body *models.SetRecordingState) *PutConversationsMessageRecordingstateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put conversations message recordingstate params
func (o *PutConversationsMessageRecordingstateParams) SetBody(body *models.SetRecordingState) {
	o.Body = body
}

// WithConversationID adds the conversationID to the put conversations message recordingstate params
func (o *PutConversationsMessageRecordingstateParams) WithConversationID(conversationID string) *PutConversationsMessageRecordingstateParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the put conversations message recordingstate params
func (o *PutConversationsMessageRecordingstateParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PutConversationsMessageRecordingstateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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