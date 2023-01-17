// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewPostAnalyticsConversationsTranscriptsQueryParams creates a new PostAnalyticsConversationsTranscriptsQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAnalyticsConversationsTranscriptsQueryParams() *PostAnalyticsConversationsTranscriptsQueryParams {
	return &PostAnalyticsConversationsTranscriptsQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAnalyticsConversationsTranscriptsQueryParamsWithTimeout creates a new PostAnalyticsConversationsTranscriptsQueryParams object
// with the ability to set a timeout on a request.
func NewPostAnalyticsConversationsTranscriptsQueryParamsWithTimeout(timeout time.Duration) *PostAnalyticsConversationsTranscriptsQueryParams {
	return &PostAnalyticsConversationsTranscriptsQueryParams{
		timeout: timeout,
	}
}

// NewPostAnalyticsConversationsTranscriptsQueryParamsWithContext creates a new PostAnalyticsConversationsTranscriptsQueryParams object
// with the ability to set a context for a request.
func NewPostAnalyticsConversationsTranscriptsQueryParamsWithContext(ctx context.Context) *PostAnalyticsConversationsTranscriptsQueryParams {
	return &PostAnalyticsConversationsTranscriptsQueryParams{
		Context: ctx,
	}
}

// NewPostAnalyticsConversationsTranscriptsQueryParamsWithHTTPClient creates a new PostAnalyticsConversationsTranscriptsQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAnalyticsConversationsTranscriptsQueryParamsWithHTTPClient(client *http.Client) *PostAnalyticsConversationsTranscriptsQueryParams {
	return &PostAnalyticsConversationsTranscriptsQueryParams{
		HTTPClient: client,
	}
}

/*
PostAnalyticsConversationsTranscriptsQueryParams contains all the parameters to send to the API endpoint

	for the post analytics conversations transcripts query operation.

	Typically these are written to a http.Request.
*/
type PostAnalyticsConversationsTranscriptsQueryParams struct {

	/* Body.

	   Search request options
	*/
	Body *models.TranscriptConversationDetailSearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post analytics conversations transcripts query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAnalyticsConversationsTranscriptsQueryParams) WithDefaults() *PostAnalyticsConversationsTranscriptsQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post analytics conversations transcripts query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAnalyticsConversationsTranscriptsQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post analytics conversations transcripts query params
func (o *PostAnalyticsConversationsTranscriptsQueryParams) WithTimeout(timeout time.Duration) *PostAnalyticsConversationsTranscriptsQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post analytics conversations transcripts query params
func (o *PostAnalyticsConversationsTranscriptsQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post analytics conversations transcripts query params
func (o *PostAnalyticsConversationsTranscriptsQueryParams) WithContext(ctx context.Context) *PostAnalyticsConversationsTranscriptsQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post analytics conversations transcripts query params
func (o *PostAnalyticsConversationsTranscriptsQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post analytics conversations transcripts query params
func (o *PostAnalyticsConversationsTranscriptsQueryParams) WithHTTPClient(client *http.Client) *PostAnalyticsConversationsTranscriptsQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post analytics conversations transcripts query params
func (o *PostAnalyticsConversationsTranscriptsQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post analytics conversations transcripts query params
func (o *PostAnalyticsConversationsTranscriptsQueryParams) WithBody(body *models.TranscriptConversationDetailSearchRequest) *PostAnalyticsConversationsTranscriptsQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post analytics conversations transcripts query params
func (o *PostAnalyticsConversationsTranscriptsQueryParams) SetBody(body *models.TranscriptConversationDetailSearchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAnalyticsConversationsTranscriptsQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
