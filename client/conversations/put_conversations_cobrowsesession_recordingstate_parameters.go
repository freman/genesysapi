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

// NewPutConversationsCobrowsesessionRecordingstateParams creates a new PutConversationsCobrowsesessionRecordingstateParams object
// with the default values initialized.
func NewPutConversationsCobrowsesessionRecordingstateParams() *PutConversationsCobrowsesessionRecordingstateParams {
	var ()
	return &PutConversationsCobrowsesessionRecordingstateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutConversationsCobrowsesessionRecordingstateParamsWithTimeout creates a new PutConversationsCobrowsesessionRecordingstateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutConversationsCobrowsesessionRecordingstateParamsWithTimeout(timeout time.Duration) *PutConversationsCobrowsesessionRecordingstateParams {
	var ()
	return &PutConversationsCobrowsesessionRecordingstateParams{

		timeout: timeout,
	}
}

// NewPutConversationsCobrowsesessionRecordingstateParamsWithContext creates a new PutConversationsCobrowsesessionRecordingstateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutConversationsCobrowsesessionRecordingstateParamsWithContext(ctx context.Context) *PutConversationsCobrowsesessionRecordingstateParams {
	var ()
	return &PutConversationsCobrowsesessionRecordingstateParams{

		Context: ctx,
	}
}

// NewPutConversationsCobrowsesessionRecordingstateParamsWithHTTPClient creates a new PutConversationsCobrowsesessionRecordingstateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutConversationsCobrowsesessionRecordingstateParamsWithHTTPClient(client *http.Client) *PutConversationsCobrowsesessionRecordingstateParams {
	var ()
	return &PutConversationsCobrowsesessionRecordingstateParams{
		HTTPClient: client,
	}
}

/*PutConversationsCobrowsesessionRecordingstateParams contains all the parameters to send to the API endpoint
for the put conversations cobrowsesession recordingstate operation typically these are written to a http.Request
*/
type PutConversationsCobrowsesessionRecordingstateParams struct {

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

// WithTimeout adds the timeout to the put conversations cobrowsesession recordingstate params
func (o *PutConversationsCobrowsesessionRecordingstateParams) WithTimeout(timeout time.Duration) *PutConversationsCobrowsesessionRecordingstateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put conversations cobrowsesession recordingstate params
func (o *PutConversationsCobrowsesessionRecordingstateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put conversations cobrowsesession recordingstate params
func (o *PutConversationsCobrowsesessionRecordingstateParams) WithContext(ctx context.Context) *PutConversationsCobrowsesessionRecordingstateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put conversations cobrowsesession recordingstate params
func (o *PutConversationsCobrowsesessionRecordingstateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put conversations cobrowsesession recordingstate params
func (o *PutConversationsCobrowsesessionRecordingstateParams) WithHTTPClient(client *http.Client) *PutConversationsCobrowsesessionRecordingstateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put conversations cobrowsesession recordingstate params
func (o *PutConversationsCobrowsesessionRecordingstateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put conversations cobrowsesession recordingstate params
func (o *PutConversationsCobrowsesessionRecordingstateParams) WithBody(body *models.SetRecordingState) *PutConversationsCobrowsesessionRecordingstateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put conversations cobrowsesession recordingstate params
func (o *PutConversationsCobrowsesessionRecordingstateParams) SetBody(body *models.SetRecordingState) {
	o.Body = body
}

// WithConversationID adds the conversationID to the put conversations cobrowsesession recordingstate params
func (o *PutConversationsCobrowsesessionRecordingstateParams) WithConversationID(conversationID string) *PutConversationsCobrowsesessionRecordingstateParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the put conversations cobrowsesession recordingstate params
func (o *PutConversationsCobrowsesessionRecordingstateParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PutConversationsCobrowsesessionRecordingstateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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