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

// NewDeleteConversationRecordingAnnotationParams creates a new DeleteConversationRecordingAnnotationParams object
// with the default values initialized.
func NewDeleteConversationRecordingAnnotationParams() *DeleteConversationRecordingAnnotationParams {
	var ()
	return &DeleteConversationRecordingAnnotationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConversationRecordingAnnotationParamsWithTimeout creates a new DeleteConversationRecordingAnnotationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteConversationRecordingAnnotationParamsWithTimeout(timeout time.Duration) *DeleteConversationRecordingAnnotationParams {
	var ()
	return &DeleteConversationRecordingAnnotationParams{

		timeout: timeout,
	}
}

// NewDeleteConversationRecordingAnnotationParamsWithContext creates a new DeleteConversationRecordingAnnotationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteConversationRecordingAnnotationParamsWithContext(ctx context.Context) *DeleteConversationRecordingAnnotationParams {
	var ()
	return &DeleteConversationRecordingAnnotationParams{

		Context: ctx,
	}
}

// NewDeleteConversationRecordingAnnotationParamsWithHTTPClient creates a new DeleteConversationRecordingAnnotationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteConversationRecordingAnnotationParamsWithHTTPClient(client *http.Client) *DeleteConversationRecordingAnnotationParams {
	var ()
	return &DeleteConversationRecordingAnnotationParams{
		HTTPClient: client,
	}
}

/*DeleteConversationRecordingAnnotationParams contains all the parameters to send to the API endpoint
for the delete conversation recording annotation operation typically these are written to a http.Request
*/
type DeleteConversationRecordingAnnotationParams struct {

	/*AnnotationID
	  Annotation ID

	*/
	AnnotationID string
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

// WithTimeout adds the timeout to the delete conversation recording annotation params
func (o *DeleteConversationRecordingAnnotationParams) WithTimeout(timeout time.Duration) *DeleteConversationRecordingAnnotationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete conversation recording annotation params
func (o *DeleteConversationRecordingAnnotationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete conversation recording annotation params
func (o *DeleteConversationRecordingAnnotationParams) WithContext(ctx context.Context) *DeleteConversationRecordingAnnotationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete conversation recording annotation params
func (o *DeleteConversationRecordingAnnotationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete conversation recording annotation params
func (o *DeleteConversationRecordingAnnotationParams) WithHTTPClient(client *http.Client) *DeleteConversationRecordingAnnotationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete conversation recording annotation params
func (o *DeleteConversationRecordingAnnotationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAnnotationID adds the annotationID to the delete conversation recording annotation params
func (o *DeleteConversationRecordingAnnotationParams) WithAnnotationID(annotationID string) *DeleteConversationRecordingAnnotationParams {
	o.SetAnnotationID(annotationID)
	return o
}

// SetAnnotationID adds the annotationId to the delete conversation recording annotation params
func (o *DeleteConversationRecordingAnnotationParams) SetAnnotationID(annotationID string) {
	o.AnnotationID = annotationID
}

// WithConversationID adds the conversationID to the delete conversation recording annotation params
func (o *DeleteConversationRecordingAnnotationParams) WithConversationID(conversationID string) *DeleteConversationRecordingAnnotationParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the delete conversation recording annotation params
func (o *DeleteConversationRecordingAnnotationParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithRecordingID adds the recordingID to the delete conversation recording annotation params
func (o *DeleteConversationRecordingAnnotationParams) WithRecordingID(recordingID string) *DeleteConversationRecordingAnnotationParams {
	o.SetRecordingID(recordingID)
	return o
}

// SetRecordingID adds the recordingId to the delete conversation recording annotation params
func (o *DeleteConversationRecordingAnnotationParams) SetRecordingID(recordingID string) {
	o.RecordingID = recordingID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConversationRecordingAnnotationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param annotationId
	if err := r.SetPathParam("annotationId", o.AnnotationID); err != nil {
		return err
	}

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
