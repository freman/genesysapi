// Code generated by go-swagger; DO NOT EDIT.

package process_automation

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

// NewPostProcessautomationTriggersTopicTestParams creates a new PostProcessautomationTriggersTopicTestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostProcessautomationTriggersTopicTestParams() *PostProcessautomationTriggersTopicTestParams {
	return &PostProcessautomationTriggersTopicTestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostProcessautomationTriggersTopicTestParamsWithTimeout creates a new PostProcessautomationTriggersTopicTestParams object
// with the ability to set a timeout on a request.
func NewPostProcessautomationTriggersTopicTestParamsWithTimeout(timeout time.Duration) *PostProcessautomationTriggersTopicTestParams {
	return &PostProcessautomationTriggersTopicTestParams{
		timeout: timeout,
	}
}

// NewPostProcessautomationTriggersTopicTestParamsWithContext creates a new PostProcessautomationTriggersTopicTestParams object
// with the ability to set a context for a request.
func NewPostProcessautomationTriggersTopicTestParamsWithContext(ctx context.Context) *PostProcessautomationTriggersTopicTestParams {
	return &PostProcessautomationTriggersTopicTestParams{
		Context: ctx,
	}
}

// NewPostProcessautomationTriggersTopicTestParamsWithHTTPClient creates a new PostProcessautomationTriggersTopicTestParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostProcessautomationTriggersTopicTestParamsWithHTTPClient(client *http.Client) *PostProcessautomationTriggersTopicTestParams {
	return &PostProcessautomationTriggersTopicTestParams{
		HTTPClient: client,
	}
}

/*
PostProcessautomationTriggersTopicTestParams contains all the parameters to send to the API endpoint

	for the post processautomation triggers topic test operation.

	Typically these are written to a http.Request.
*/
type PostProcessautomationTriggersTopicTestParams struct {

	/* Body.

	   eventBody
	*/
	Body string

	/* TopicName.

	   topicName
	*/
	TopicName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post processautomation triggers topic test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProcessautomationTriggersTopicTestParams) WithDefaults() *PostProcessautomationTriggersTopicTestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post processautomation triggers topic test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProcessautomationTriggersTopicTestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post processautomation triggers topic test params
func (o *PostProcessautomationTriggersTopicTestParams) WithTimeout(timeout time.Duration) *PostProcessautomationTriggersTopicTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post processautomation triggers topic test params
func (o *PostProcessautomationTriggersTopicTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post processautomation triggers topic test params
func (o *PostProcessautomationTriggersTopicTestParams) WithContext(ctx context.Context) *PostProcessautomationTriggersTopicTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post processautomation triggers topic test params
func (o *PostProcessautomationTriggersTopicTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post processautomation triggers topic test params
func (o *PostProcessautomationTriggersTopicTestParams) WithHTTPClient(client *http.Client) *PostProcessautomationTriggersTopicTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post processautomation triggers topic test params
func (o *PostProcessautomationTriggersTopicTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post processautomation triggers topic test params
func (o *PostProcessautomationTriggersTopicTestParams) WithBody(body string) *PostProcessautomationTriggersTopicTestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post processautomation triggers topic test params
func (o *PostProcessautomationTriggersTopicTestParams) SetBody(body string) {
	o.Body = body
}

// WithTopicName adds the topicName to the post processautomation triggers topic test params
func (o *PostProcessautomationTriggersTopicTestParams) WithTopicName(topicName string) *PostProcessautomationTriggersTopicTestParams {
	o.SetTopicName(topicName)
	return o
}

// SetTopicName adds the topicName to the post processautomation triggers topic test params
func (o *PostProcessautomationTriggersTopicTestParams) SetTopicName(topicName string) {
	o.TopicName = topicName
}

// WriteToRequest writes these params to a swagger request
func (o *PostProcessautomationTriggersTopicTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param topicName
	if err := r.SetPathParam("topicName", o.TopicName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
