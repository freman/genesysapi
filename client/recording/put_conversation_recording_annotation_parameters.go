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

	"github.com/freman/genesysapi/models"
)

// NewPutConversationRecordingAnnotationParams creates a new PutConversationRecordingAnnotationParams object
// with the default values initialized.
func NewPutConversationRecordingAnnotationParams() *PutConversationRecordingAnnotationParams {
	var ()
	return &PutConversationRecordingAnnotationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutConversationRecordingAnnotationParamsWithTimeout creates a new PutConversationRecordingAnnotationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutConversationRecordingAnnotationParamsWithTimeout(timeout time.Duration) *PutConversationRecordingAnnotationParams {
	var ()
	return &PutConversationRecordingAnnotationParams{

		timeout: timeout,
	}
}

// NewPutConversationRecordingAnnotationParamsWithContext creates a new PutConversationRecordingAnnotationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutConversationRecordingAnnotationParamsWithContext(ctx context.Context) *PutConversationRecordingAnnotationParams {
	var ()
	return &PutConversationRecordingAnnotationParams{

		Context: ctx,
	}
}

// NewPutConversationRecordingAnnotationParamsWithHTTPClient creates a new PutConversationRecordingAnnotationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutConversationRecordingAnnotationParamsWithHTTPClient(client *http.Client) *PutConversationRecordingAnnotationParams {
	var ()
	return &PutConversationRecordingAnnotationParams{
		HTTPClient: client,
	}
}

/*PutConversationRecordingAnnotationParams contains all the parameters to send to the API endpoint
for the put conversation recording annotation operation typically these are written to a http.Request
*/
type PutConversationRecordingAnnotationParams struct {

	/*AnnotationID
	  Annotation ID

	*/
	AnnotationID string
	/*Body
	  annotation

	*/
	Body *models.Annotation
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

// WithTimeout adds the timeout to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) WithTimeout(timeout time.Duration) *PutConversationRecordingAnnotationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) WithContext(ctx context.Context) *PutConversationRecordingAnnotationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) WithHTTPClient(client *http.Client) *PutConversationRecordingAnnotationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAnnotationID adds the annotationID to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) WithAnnotationID(annotationID string) *PutConversationRecordingAnnotationParams {
	o.SetAnnotationID(annotationID)
	return o
}

// SetAnnotationID adds the annotationId to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) SetAnnotationID(annotationID string) {
	o.AnnotationID = annotationID
}

// WithBody adds the body to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) WithBody(body *models.Annotation) *PutConversationRecordingAnnotationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) SetBody(body *models.Annotation) {
	o.Body = body
}

// WithConversationID adds the conversationID to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) WithConversationID(conversationID string) *PutConversationRecordingAnnotationParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithRecordingID adds the recordingID to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) WithRecordingID(recordingID string) *PutConversationRecordingAnnotationParams {
	o.SetRecordingID(recordingID)
	return o
}

// SetRecordingID adds the recordingId to the put conversation recording annotation params
func (o *PutConversationRecordingAnnotationParams) SetRecordingID(recordingID string) {
	o.RecordingID = recordingID
}

// WriteToRequest writes these params to a swagger request
func (o *PutConversationRecordingAnnotationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param annotationId
	if err := r.SetPathParam("annotationId", o.AnnotationID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
