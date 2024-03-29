// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

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

// NewGetSpeechandtextanalyticsTopicParams creates a new GetSpeechandtextanalyticsTopicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSpeechandtextanalyticsTopicParams() *GetSpeechandtextanalyticsTopicParams {
	return &GetSpeechandtextanalyticsTopicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpeechandtextanalyticsTopicParamsWithTimeout creates a new GetSpeechandtextanalyticsTopicParams object
// with the ability to set a timeout on a request.
func NewGetSpeechandtextanalyticsTopicParamsWithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsTopicParams {
	return &GetSpeechandtextanalyticsTopicParams{
		timeout: timeout,
	}
}

// NewGetSpeechandtextanalyticsTopicParamsWithContext creates a new GetSpeechandtextanalyticsTopicParams object
// with the ability to set a context for a request.
func NewGetSpeechandtextanalyticsTopicParamsWithContext(ctx context.Context) *GetSpeechandtextanalyticsTopicParams {
	return &GetSpeechandtextanalyticsTopicParams{
		Context: ctx,
	}
}

// NewGetSpeechandtextanalyticsTopicParamsWithHTTPClient creates a new GetSpeechandtextanalyticsTopicParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSpeechandtextanalyticsTopicParamsWithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsTopicParams {
	return &GetSpeechandtextanalyticsTopicParams{
		HTTPClient: client,
	}
}

/*
GetSpeechandtextanalyticsTopicParams contains all the parameters to send to the API endpoint

	for the get speechandtextanalytics topic operation.

	Typically these are written to a http.Request.
*/
type GetSpeechandtextanalyticsTopicParams struct {

	/* TopicID.

	   The id of the topic
	*/
	TopicID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get speechandtextanalytics topic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpeechandtextanalyticsTopicParams) WithDefaults() *GetSpeechandtextanalyticsTopicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get speechandtextanalytics topic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpeechandtextanalyticsTopicParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get speechandtextanalytics topic params
func (o *GetSpeechandtextanalyticsTopicParams) WithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsTopicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get speechandtextanalytics topic params
func (o *GetSpeechandtextanalyticsTopicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get speechandtextanalytics topic params
func (o *GetSpeechandtextanalyticsTopicParams) WithContext(ctx context.Context) *GetSpeechandtextanalyticsTopicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get speechandtextanalytics topic params
func (o *GetSpeechandtextanalyticsTopicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get speechandtextanalytics topic params
func (o *GetSpeechandtextanalyticsTopicParams) WithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsTopicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get speechandtextanalytics topic params
func (o *GetSpeechandtextanalyticsTopicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTopicID adds the topicID to the get speechandtextanalytics topic params
func (o *GetSpeechandtextanalyticsTopicParams) WithTopicID(topicID string) *GetSpeechandtextanalyticsTopicParams {
	o.SetTopicID(topicID)
	return o
}

// SetTopicID adds the topicId to the get speechandtextanalytics topic params
func (o *GetSpeechandtextanalyticsTopicParams) SetTopicID(topicID string) {
	o.TopicID = topicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpeechandtextanalyticsTopicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param topicId
	if err := r.SetPathParam("topicId", o.TopicID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
