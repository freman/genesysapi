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

// NewPutConversationsChatRecordingstateParams creates a new PutConversationsChatRecordingstateParams object
// with the default values initialized.
func NewPutConversationsChatRecordingstateParams() *PutConversationsChatRecordingstateParams {
	var ()
	return &PutConversationsChatRecordingstateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutConversationsChatRecordingstateParamsWithTimeout creates a new PutConversationsChatRecordingstateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutConversationsChatRecordingstateParamsWithTimeout(timeout time.Duration) *PutConversationsChatRecordingstateParams {
	var ()
	return &PutConversationsChatRecordingstateParams{

		timeout: timeout,
	}
}

// NewPutConversationsChatRecordingstateParamsWithContext creates a new PutConversationsChatRecordingstateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutConversationsChatRecordingstateParamsWithContext(ctx context.Context) *PutConversationsChatRecordingstateParams {
	var ()
	return &PutConversationsChatRecordingstateParams{

		Context: ctx,
	}
}

// NewPutConversationsChatRecordingstateParamsWithHTTPClient creates a new PutConversationsChatRecordingstateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutConversationsChatRecordingstateParamsWithHTTPClient(client *http.Client) *PutConversationsChatRecordingstateParams {
	var ()
	return &PutConversationsChatRecordingstateParams{
		HTTPClient: client,
	}
}

/*PutConversationsChatRecordingstateParams contains all the parameters to send to the API endpoint
for the put conversations chat recordingstate operation typically these are written to a http.Request
*/
type PutConversationsChatRecordingstateParams struct {

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

// WithTimeout adds the timeout to the put conversations chat recordingstate params
func (o *PutConversationsChatRecordingstateParams) WithTimeout(timeout time.Duration) *PutConversationsChatRecordingstateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put conversations chat recordingstate params
func (o *PutConversationsChatRecordingstateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put conversations chat recordingstate params
func (o *PutConversationsChatRecordingstateParams) WithContext(ctx context.Context) *PutConversationsChatRecordingstateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put conversations chat recordingstate params
func (o *PutConversationsChatRecordingstateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put conversations chat recordingstate params
func (o *PutConversationsChatRecordingstateParams) WithHTTPClient(client *http.Client) *PutConversationsChatRecordingstateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put conversations chat recordingstate params
func (o *PutConversationsChatRecordingstateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put conversations chat recordingstate params
func (o *PutConversationsChatRecordingstateParams) WithBody(body *models.SetRecordingState) *PutConversationsChatRecordingstateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put conversations chat recordingstate params
func (o *PutConversationsChatRecordingstateParams) SetBody(body *models.SetRecordingState) {
	o.Body = body
}

// WithConversationID adds the conversationID to the put conversations chat recordingstate params
func (o *PutConversationsChatRecordingstateParams) WithConversationID(conversationID string) *PutConversationsChatRecordingstateParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the put conversations chat recordingstate params
func (o *PutConversationsChatRecordingstateParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PutConversationsChatRecordingstateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
