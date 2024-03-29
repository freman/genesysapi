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

// NewGetConversationRecordingAnnotationParams creates a new GetConversationRecordingAnnotationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConversationRecordingAnnotationParams() *GetConversationRecordingAnnotationParams {
	return &GetConversationRecordingAnnotationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationRecordingAnnotationParamsWithTimeout creates a new GetConversationRecordingAnnotationParams object
// with the ability to set a timeout on a request.
func NewGetConversationRecordingAnnotationParamsWithTimeout(timeout time.Duration) *GetConversationRecordingAnnotationParams {
	return &GetConversationRecordingAnnotationParams{
		timeout: timeout,
	}
}

// NewGetConversationRecordingAnnotationParamsWithContext creates a new GetConversationRecordingAnnotationParams object
// with the ability to set a context for a request.
func NewGetConversationRecordingAnnotationParamsWithContext(ctx context.Context) *GetConversationRecordingAnnotationParams {
	return &GetConversationRecordingAnnotationParams{
		Context: ctx,
	}
}

// NewGetConversationRecordingAnnotationParamsWithHTTPClient creates a new GetConversationRecordingAnnotationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConversationRecordingAnnotationParamsWithHTTPClient(client *http.Client) *GetConversationRecordingAnnotationParams {
	return &GetConversationRecordingAnnotationParams{
		HTTPClient: client,
	}
}

/*
GetConversationRecordingAnnotationParams contains all the parameters to send to the API endpoint

	for the get conversation recording annotation operation.

	Typically these are written to a http.Request.
*/
type GetConversationRecordingAnnotationParams struct {

	/* AnnotationID.

	   Annotation ID
	*/
	AnnotationID string

	/* ConversationID.

	   Conversation ID
	*/
	ConversationID string

	/* RecordingID.

	   Recording ID
	*/
	RecordingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get conversation recording annotation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationRecordingAnnotationParams) WithDefaults() *GetConversationRecordingAnnotationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get conversation recording annotation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationRecordingAnnotationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get conversation recording annotation params
func (o *GetConversationRecordingAnnotationParams) WithTimeout(timeout time.Duration) *GetConversationRecordingAnnotationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversation recording annotation params
func (o *GetConversationRecordingAnnotationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversation recording annotation params
func (o *GetConversationRecordingAnnotationParams) WithContext(ctx context.Context) *GetConversationRecordingAnnotationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversation recording annotation params
func (o *GetConversationRecordingAnnotationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversation recording annotation params
func (o *GetConversationRecordingAnnotationParams) WithHTTPClient(client *http.Client) *GetConversationRecordingAnnotationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversation recording annotation params
func (o *GetConversationRecordingAnnotationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAnnotationID adds the annotationID to the get conversation recording annotation params
func (o *GetConversationRecordingAnnotationParams) WithAnnotationID(annotationID string) *GetConversationRecordingAnnotationParams {
	o.SetAnnotationID(annotationID)
	return o
}

// SetAnnotationID adds the annotationId to the get conversation recording annotation params
func (o *GetConversationRecordingAnnotationParams) SetAnnotationID(annotationID string) {
	o.AnnotationID = annotationID
}

// WithConversationID adds the conversationID to the get conversation recording annotation params
func (o *GetConversationRecordingAnnotationParams) WithConversationID(conversationID string) *GetConversationRecordingAnnotationParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get conversation recording annotation params
func (o *GetConversationRecordingAnnotationParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithRecordingID adds the recordingID to the get conversation recording annotation params
func (o *GetConversationRecordingAnnotationParams) WithRecordingID(recordingID string) *GetConversationRecordingAnnotationParams {
	o.SetRecordingID(recordingID)
	return o
}

// SetRecordingID adds the recordingId to the get conversation recording annotation params
func (o *GetConversationRecordingAnnotationParams) SetRecordingID(recordingID string) {
	o.RecordingID = recordingID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationRecordingAnnotationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
