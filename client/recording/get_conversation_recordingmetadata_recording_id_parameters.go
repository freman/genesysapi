// Code generated by go-swagger; DO NOT EDIT.

package recording

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
)

// NewGetConversationRecordingmetadataRecordingIDParams creates a new GetConversationRecordingmetadataRecordingIDParams object
// with the default values initialized.
func NewGetConversationRecordingmetadataRecordingIDParams() *GetConversationRecordingmetadataRecordingIDParams {
	var ()
	return &GetConversationRecordingmetadataRecordingIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationRecordingmetadataRecordingIDParamsWithTimeout creates a new GetConversationRecordingmetadataRecordingIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationRecordingmetadataRecordingIDParamsWithTimeout(timeout time.Duration) *GetConversationRecordingmetadataRecordingIDParams {
	var ()
	return &GetConversationRecordingmetadataRecordingIDParams{

		timeout: timeout,
	}
}

// NewGetConversationRecordingmetadataRecordingIDParamsWithContext creates a new GetConversationRecordingmetadataRecordingIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationRecordingmetadataRecordingIDParamsWithContext(ctx context.Context) *GetConversationRecordingmetadataRecordingIDParams {
	var ()
	return &GetConversationRecordingmetadataRecordingIDParams{

		Context: ctx,
	}
}

// NewGetConversationRecordingmetadataRecordingIDParamsWithHTTPClient creates a new GetConversationRecordingmetadataRecordingIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationRecordingmetadataRecordingIDParamsWithHTTPClient(client *http.Client) *GetConversationRecordingmetadataRecordingIDParams {
	var ()
	return &GetConversationRecordingmetadataRecordingIDParams{
		HTTPClient: client,
	}
}

/*GetConversationRecordingmetadataRecordingIDParams contains all the parameters to send to the API endpoint
for the get conversation recordingmetadata recording Id operation typically these are written to a http.Request
*/
type GetConversationRecordingmetadataRecordingIDParams struct {

	/*ConversationID
	  Conversation ID

	*/
	ConversationID string
	/*RecordingID
	  Recording ID

	*/
	RecordingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get conversation recordingmetadata recording Id params
func (o *GetConversationRecordingmetadataRecordingIDParams) WithTimeout(timeout time.Duration) *GetConversationRecordingmetadataRecordingIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversation recordingmetadata recording Id params
func (o *GetConversationRecordingmetadataRecordingIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversation recordingmetadata recording Id params
func (o *GetConversationRecordingmetadataRecordingIDParams) WithContext(ctx context.Context) *GetConversationRecordingmetadataRecordingIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversation recordingmetadata recording Id params
func (o *GetConversationRecordingmetadataRecordingIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversation recordingmetadata recording Id params
func (o *GetConversationRecordingmetadataRecordingIDParams) WithHTTPClient(client *http.Client) *GetConversationRecordingmetadataRecordingIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversation recordingmetadata recording Id params
func (o *GetConversationRecordingmetadataRecordingIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the get conversation recordingmetadata recording Id params
func (o *GetConversationRecordingmetadataRecordingIDParams) WithConversationID(conversationID string) *GetConversationRecordingmetadataRecordingIDParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get conversation recordingmetadata recording Id params
func (o *GetConversationRecordingmetadataRecordingIDParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithRecordingID adds the recordingID to the get conversation recordingmetadata recording Id params
func (o *GetConversationRecordingmetadataRecordingIDParams) WithRecordingID(recordingID string) *GetConversationRecordingmetadataRecordingIDParams {
	o.SetRecordingID(recordingID)
	return o
}

// SetRecordingID adds the recordingId to the get conversation recordingmetadata recording Id params
func (o *GetConversationRecordingmetadataRecordingIDParams) SetRecordingID(recordingID string) {
	o.RecordingID = recordingID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationRecordingmetadataRecordingIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param conversationId
	if err := r.SetPathParam("conversationId", o.ConversationID); err != nil {
		return err
	}

	// path param recordingId
	if err := r.SetPathParam("recordingId", o.RecordingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
