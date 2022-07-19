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

// NewPutConversationsSocialRecordingstateParams creates a new PutConversationsSocialRecordingstateParams object
// with the default values initialized.
func NewPutConversationsSocialRecordingstateParams() *PutConversationsSocialRecordingstateParams {
	var ()
	return &PutConversationsSocialRecordingstateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutConversationsSocialRecordingstateParamsWithTimeout creates a new PutConversationsSocialRecordingstateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutConversationsSocialRecordingstateParamsWithTimeout(timeout time.Duration) *PutConversationsSocialRecordingstateParams {
	var ()
	return &PutConversationsSocialRecordingstateParams{

		timeout: timeout,
	}
}

// NewPutConversationsSocialRecordingstateParamsWithContext creates a new PutConversationsSocialRecordingstateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutConversationsSocialRecordingstateParamsWithContext(ctx context.Context) *PutConversationsSocialRecordingstateParams {
	var ()
	return &PutConversationsSocialRecordingstateParams{

		Context: ctx,
	}
}

// NewPutConversationsSocialRecordingstateParamsWithHTTPClient creates a new PutConversationsSocialRecordingstateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutConversationsSocialRecordingstateParamsWithHTTPClient(client *http.Client) *PutConversationsSocialRecordingstateParams {
	var ()
	return &PutConversationsSocialRecordingstateParams{
		HTTPClient: client,
	}
}

/*PutConversationsSocialRecordingstateParams contains all the parameters to send to the API endpoint
for the put conversations social recordingstate operation typically these are written to a http.Request
*/
type PutConversationsSocialRecordingstateParams struct {

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

// WithTimeout adds the timeout to the put conversations social recordingstate params
func (o *PutConversationsSocialRecordingstateParams) WithTimeout(timeout time.Duration) *PutConversationsSocialRecordingstateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put conversations social recordingstate params
func (o *PutConversationsSocialRecordingstateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put conversations social recordingstate params
func (o *PutConversationsSocialRecordingstateParams) WithContext(ctx context.Context) *PutConversationsSocialRecordingstateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put conversations social recordingstate params
func (o *PutConversationsSocialRecordingstateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put conversations social recordingstate params
func (o *PutConversationsSocialRecordingstateParams) WithHTTPClient(client *http.Client) *PutConversationsSocialRecordingstateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put conversations social recordingstate params
func (o *PutConversationsSocialRecordingstateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put conversations social recordingstate params
func (o *PutConversationsSocialRecordingstateParams) WithBody(body *models.SetRecordingState) *PutConversationsSocialRecordingstateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put conversations social recordingstate params
func (o *PutConversationsSocialRecordingstateParams) SetBody(body *models.SetRecordingState) {
	o.Body = body
}

// WithConversationID adds the conversationID to the put conversations social recordingstate params
func (o *PutConversationsSocialRecordingstateParams) WithConversationID(conversationID string) *PutConversationsSocialRecordingstateParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the put conversations social recordingstate params
func (o *PutConversationsSocialRecordingstateParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PutConversationsSocialRecordingstateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
