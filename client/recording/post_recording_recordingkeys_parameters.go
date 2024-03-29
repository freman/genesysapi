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

// NewPostRecordingRecordingkeysParams creates a new PostRecordingRecordingkeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRecordingRecordingkeysParams() *PostRecordingRecordingkeysParams {
	return &PostRecordingRecordingkeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRecordingRecordingkeysParamsWithTimeout creates a new PostRecordingRecordingkeysParams object
// with the ability to set a timeout on a request.
func NewPostRecordingRecordingkeysParamsWithTimeout(timeout time.Duration) *PostRecordingRecordingkeysParams {
	return &PostRecordingRecordingkeysParams{
		timeout: timeout,
	}
}

// NewPostRecordingRecordingkeysParamsWithContext creates a new PostRecordingRecordingkeysParams object
// with the ability to set a context for a request.
func NewPostRecordingRecordingkeysParamsWithContext(ctx context.Context) *PostRecordingRecordingkeysParams {
	return &PostRecordingRecordingkeysParams{
		Context: ctx,
	}
}

// NewPostRecordingRecordingkeysParamsWithHTTPClient creates a new PostRecordingRecordingkeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRecordingRecordingkeysParamsWithHTTPClient(client *http.Client) *PostRecordingRecordingkeysParams {
	return &PostRecordingRecordingkeysParams{
		HTTPClient: client,
	}
}

/*
PostRecordingRecordingkeysParams contains all the parameters to send to the API endpoint

	for the post recording recordingkeys operation.

	Typically these are written to a http.Request.
*/
type PostRecordingRecordingkeysParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post recording recordingkeys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRecordingRecordingkeysParams) WithDefaults() *PostRecordingRecordingkeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post recording recordingkeys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRecordingRecordingkeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post recording recordingkeys params
func (o *PostRecordingRecordingkeysParams) WithTimeout(timeout time.Duration) *PostRecordingRecordingkeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post recording recordingkeys params
func (o *PostRecordingRecordingkeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post recording recordingkeys params
func (o *PostRecordingRecordingkeysParams) WithContext(ctx context.Context) *PostRecordingRecordingkeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post recording recordingkeys params
func (o *PostRecordingRecordingkeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post recording recordingkeys params
func (o *PostRecordingRecordingkeysParams) WithHTTPClient(client *http.Client) *PostRecordingRecordingkeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post recording recordingkeys params
func (o *PostRecordingRecordingkeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostRecordingRecordingkeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
