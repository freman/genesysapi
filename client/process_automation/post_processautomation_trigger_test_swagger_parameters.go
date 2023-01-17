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

// NewPostProcessautomationTriggerTestParams creates a new PostProcessautomationTriggerTestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostProcessautomationTriggerTestParams() *PostProcessautomationTriggerTestParams {
	return &PostProcessautomationTriggerTestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostProcessautomationTriggerTestParamsWithTimeout creates a new PostProcessautomationTriggerTestParams object
// with the ability to set a timeout on a request.
func NewPostProcessautomationTriggerTestParamsWithTimeout(timeout time.Duration) *PostProcessautomationTriggerTestParams {
	return &PostProcessautomationTriggerTestParams{
		timeout: timeout,
	}
}

// NewPostProcessautomationTriggerTestParamsWithContext creates a new PostProcessautomationTriggerTestParams object
// with the ability to set a context for a request.
func NewPostProcessautomationTriggerTestParamsWithContext(ctx context.Context) *PostProcessautomationTriggerTestParams {
	return &PostProcessautomationTriggerTestParams{
		Context: ctx,
	}
}

// NewPostProcessautomationTriggerTestParamsWithHTTPClient creates a new PostProcessautomationTriggerTestParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostProcessautomationTriggerTestParamsWithHTTPClient(client *http.Client) *PostProcessautomationTriggerTestParams {
	return &PostProcessautomationTriggerTestParams{
		HTTPClient: client,
	}
}

/*
PostProcessautomationTriggerTestParams contains all the parameters to send to the API endpoint

	for the post processautomation trigger test operation.

	Typically these are written to a http.Request.
*/
type PostProcessautomationTriggerTestParams struct {

	/* Body.

	   eventBody
	*/
	Body string

	/* TriggerID.

	   triggerId
	*/
	TriggerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post processautomation trigger test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProcessautomationTriggerTestParams) WithDefaults() *PostProcessautomationTriggerTestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post processautomation trigger test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProcessautomationTriggerTestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post processautomation trigger test params
func (o *PostProcessautomationTriggerTestParams) WithTimeout(timeout time.Duration) *PostProcessautomationTriggerTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post processautomation trigger test params
func (o *PostProcessautomationTriggerTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post processautomation trigger test params
func (o *PostProcessautomationTriggerTestParams) WithContext(ctx context.Context) *PostProcessautomationTriggerTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post processautomation trigger test params
func (o *PostProcessautomationTriggerTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post processautomation trigger test params
func (o *PostProcessautomationTriggerTestParams) WithHTTPClient(client *http.Client) *PostProcessautomationTriggerTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post processautomation trigger test params
func (o *PostProcessautomationTriggerTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post processautomation trigger test params
func (o *PostProcessautomationTriggerTestParams) WithBody(body string) *PostProcessautomationTriggerTestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post processautomation trigger test params
func (o *PostProcessautomationTriggerTestParams) SetBody(body string) {
	o.Body = body
}

// WithTriggerID adds the triggerID to the post processautomation trigger test params
func (o *PostProcessautomationTriggerTestParams) WithTriggerID(triggerID string) *PostProcessautomationTriggerTestParams {
	o.SetTriggerID(triggerID)
	return o
}

// SetTriggerID adds the triggerId to the post processautomation trigger test params
func (o *PostProcessautomationTriggerTestParams) SetTriggerID(triggerID string) {
	o.TriggerID = triggerID
}

// WriteToRequest writes these params to a swagger request
func (o *PostProcessautomationTriggerTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param triggerId
	if err := r.SetPathParam("triggerId", o.TriggerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
