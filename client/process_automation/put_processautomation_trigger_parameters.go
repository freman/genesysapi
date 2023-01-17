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

	"github.com/freman/genesysapi/models"
)

// NewPutProcessautomationTriggerParams creates a new PutProcessautomationTriggerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutProcessautomationTriggerParams() *PutProcessautomationTriggerParams {
	return &PutProcessautomationTriggerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutProcessautomationTriggerParamsWithTimeout creates a new PutProcessautomationTriggerParams object
// with the ability to set a timeout on a request.
func NewPutProcessautomationTriggerParamsWithTimeout(timeout time.Duration) *PutProcessautomationTriggerParams {
	return &PutProcessautomationTriggerParams{
		timeout: timeout,
	}
}

// NewPutProcessautomationTriggerParamsWithContext creates a new PutProcessautomationTriggerParams object
// with the ability to set a context for a request.
func NewPutProcessautomationTriggerParamsWithContext(ctx context.Context) *PutProcessautomationTriggerParams {
	return &PutProcessautomationTriggerParams{
		Context: ctx,
	}
}

// NewPutProcessautomationTriggerParamsWithHTTPClient creates a new PutProcessautomationTriggerParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutProcessautomationTriggerParamsWithHTTPClient(client *http.Client) *PutProcessautomationTriggerParams {
	return &PutProcessautomationTriggerParams{
		HTTPClient: client,
	}
}

/*
PutProcessautomationTriggerParams contains all the parameters to send to the API endpoint

	for the put processautomation trigger operation.

	Typically these are written to a http.Request.
*/
type PutProcessautomationTriggerParams struct {

	/* Body.

	   Input to update Trigger. (topicName cannot be updated, a new trigger must be created to use a new topicName)
	*/
	Body *models.UpdateTriggerRequest

	/* TriggerID.

	   triggerId
	*/
	TriggerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put processautomation trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProcessautomationTriggerParams) WithDefaults() *PutProcessautomationTriggerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put processautomation trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutProcessautomationTriggerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put processautomation trigger params
func (o *PutProcessautomationTriggerParams) WithTimeout(timeout time.Duration) *PutProcessautomationTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put processautomation trigger params
func (o *PutProcessautomationTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put processautomation trigger params
func (o *PutProcessautomationTriggerParams) WithContext(ctx context.Context) *PutProcessautomationTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put processautomation trigger params
func (o *PutProcessautomationTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put processautomation trigger params
func (o *PutProcessautomationTriggerParams) WithHTTPClient(client *http.Client) *PutProcessautomationTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put processautomation trigger params
func (o *PutProcessautomationTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put processautomation trigger params
func (o *PutProcessautomationTriggerParams) WithBody(body *models.UpdateTriggerRequest) *PutProcessautomationTriggerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put processautomation trigger params
func (o *PutProcessautomationTriggerParams) SetBody(body *models.UpdateTriggerRequest) {
	o.Body = body
}

// WithTriggerID adds the triggerID to the put processautomation trigger params
func (o *PutProcessautomationTriggerParams) WithTriggerID(triggerID string) *PutProcessautomationTriggerParams {
	o.SetTriggerID(triggerID)
	return o
}

// SetTriggerID adds the triggerId to the put processautomation trigger params
func (o *PutProcessautomationTriggerParams) SetTriggerID(triggerID string) {
	o.TriggerID = triggerID
}

// WriteToRequest writes these params to a swagger request
func (o *PutProcessautomationTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
