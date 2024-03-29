// Code generated by go-swagger; DO NOT EDIT.

package analytics

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

// NewPostAnalyticsConversationsDetailsJobsParams creates a new PostAnalyticsConversationsDetailsJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAnalyticsConversationsDetailsJobsParams() *PostAnalyticsConversationsDetailsJobsParams {
	return &PostAnalyticsConversationsDetailsJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAnalyticsConversationsDetailsJobsParamsWithTimeout creates a new PostAnalyticsConversationsDetailsJobsParams object
// with the ability to set a timeout on a request.
func NewPostAnalyticsConversationsDetailsJobsParamsWithTimeout(timeout time.Duration) *PostAnalyticsConversationsDetailsJobsParams {
	return &PostAnalyticsConversationsDetailsJobsParams{
		timeout: timeout,
	}
}

// NewPostAnalyticsConversationsDetailsJobsParamsWithContext creates a new PostAnalyticsConversationsDetailsJobsParams object
// with the ability to set a context for a request.
func NewPostAnalyticsConversationsDetailsJobsParamsWithContext(ctx context.Context) *PostAnalyticsConversationsDetailsJobsParams {
	return &PostAnalyticsConversationsDetailsJobsParams{
		Context: ctx,
	}
}

// NewPostAnalyticsConversationsDetailsJobsParamsWithHTTPClient creates a new PostAnalyticsConversationsDetailsJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAnalyticsConversationsDetailsJobsParamsWithHTTPClient(client *http.Client) *PostAnalyticsConversationsDetailsJobsParams {
	return &PostAnalyticsConversationsDetailsJobsParams{
		HTTPClient: client,
	}
}

/*
PostAnalyticsConversationsDetailsJobsParams contains all the parameters to send to the API endpoint

	for the post analytics conversations details jobs operation.

	Typically these are written to a http.Request.
*/
type PostAnalyticsConversationsDetailsJobsParams struct {

	/* Body.

	   query
	*/
	Body *models.AsyncConversationQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post analytics conversations details jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAnalyticsConversationsDetailsJobsParams) WithDefaults() *PostAnalyticsConversationsDetailsJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post analytics conversations details jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAnalyticsConversationsDetailsJobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post analytics conversations details jobs params
func (o *PostAnalyticsConversationsDetailsJobsParams) WithTimeout(timeout time.Duration) *PostAnalyticsConversationsDetailsJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post analytics conversations details jobs params
func (o *PostAnalyticsConversationsDetailsJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post analytics conversations details jobs params
func (o *PostAnalyticsConversationsDetailsJobsParams) WithContext(ctx context.Context) *PostAnalyticsConversationsDetailsJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post analytics conversations details jobs params
func (o *PostAnalyticsConversationsDetailsJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post analytics conversations details jobs params
func (o *PostAnalyticsConversationsDetailsJobsParams) WithHTTPClient(client *http.Client) *PostAnalyticsConversationsDetailsJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post analytics conversations details jobs params
func (o *PostAnalyticsConversationsDetailsJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post analytics conversations details jobs params
func (o *PostAnalyticsConversationsDetailsJobsParams) WithBody(body *models.AsyncConversationQuery) *PostAnalyticsConversationsDetailsJobsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post analytics conversations details jobs params
func (o *PostAnalyticsConversationsDetailsJobsParams) SetBody(body *models.AsyncConversationQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAnalyticsConversationsDetailsJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
