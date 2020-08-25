// Code generated by go-swagger; DO NOT EDIT.

package user_recordings

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

// NewGetUserrecordingsSummaryParams creates a new GetUserrecordingsSummaryParams object
// with the default values initialized.
func NewGetUserrecordingsSummaryParams() *GetUserrecordingsSummaryParams {

	return &GetUserrecordingsSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserrecordingsSummaryParamsWithTimeout creates a new GetUserrecordingsSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserrecordingsSummaryParamsWithTimeout(timeout time.Duration) *GetUserrecordingsSummaryParams {

	return &GetUserrecordingsSummaryParams{

		timeout: timeout,
	}
}

// NewGetUserrecordingsSummaryParamsWithContext creates a new GetUserrecordingsSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserrecordingsSummaryParamsWithContext(ctx context.Context) *GetUserrecordingsSummaryParams {

	return &GetUserrecordingsSummaryParams{

		Context: ctx,
	}
}

// NewGetUserrecordingsSummaryParamsWithHTTPClient creates a new GetUserrecordingsSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserrecordingsSummaryParamsWithHTTPClient(client *http.Client) *GetUserrecordingsSummaryParams {

	return &GetUserrecordingsSummaryParams{
		HTTPClient: client,
	}
}

/*GetUserrecordingsSummaryParams contains all the parameters to send to the API endpoint
for the get userrecordings summary operation typically these are written to a http.Request
*/
type GetUserrecordingsSummaryParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get userrecordings summary params
func (o *GetUserrecordingsSummaryParams) WithTimeout(timeout time.Duration) *GetUserrecordingsSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get userrecordings summary params
func (o *GetUserrecordingsSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get userrecordings summary params
func (o *GetUserrecordingsSummaryParams) WithContext(ctx context.Context) *GetUserrecordingsSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get userrecordings summary params
func (o *GetUserrecordingsSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get userrecordings summary params
func (o *GetUserrecordingsSummaryParams) WithHTTPClient(client *http.Client) *GetUserrecordingsSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get userrecordings summary params
func (o *GetUserrecordingsSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserrecordingsSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}