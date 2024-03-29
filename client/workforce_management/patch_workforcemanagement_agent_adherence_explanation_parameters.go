// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

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

// NewPatchWorkforcemanagementAgentAdherenceExplanationParams creates a new PatchWorkforcemanagementAgentAdherenceExplanationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchWorkforcemanagementAgentAdherenceExplanationParams() *PatchWorkforcemanagementAgentAdherenceExplanationParams {
	return &PatchWorkforcemanagementAgentAdherenceExplanationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWorkforcemanagementAgentAdherenceExplanationParamsWithTimeout creates a new PatchWorkforcemanagementAgentAdherenceExplanationParams object
// with the ability to set a timeout on a request.
func NewPatchWorkforcemanagementAgentAdherenceExplanationParamsWithTimeout(timeout time.Duration) *PatchWorkforcemanagementAgentAdherenceExplanationParams {
	return &PatchWorkforcemanagementAgentAdherenceExplanationParams{
		timeout: timeout,
	}
}

// NewPatchWorkforcemanagementAgentAdherenceExplanationParamsWithContext creates a new PatchWorkforcemanagementAgentAdherenceExplanationParams object
// with the ability to set a context for a request.
func NewPatchWorkforcemanagementAgentAdherenceExplanationParamsWithContext(ctx context.Context) *PatchWorkforcemanagementAgentAdherenceExplanationParams {
	return &PatchWorkforcemanagementAgentAdherenceExplanationParams{
		Context: ctx,
	}
}

// NewPatchWorkforcemanagementAgentAdherenceExplanationParamsWithHTTPClient creates a new PatchWorkforcemanagementAgentAdherenceExplanationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchWorkforcemanagementAgentAdherenceExplanationParamsWithHTTPClient(client *http.Client) *PatchWorkforcemanagementAgentAdherenceExplanationParams {
	return &PatchWorkforcemanagementAgentAdherenceExplanationParams{
		HTTPClient: client,
	}
}

/*
PatchWorkforcemanagementAgentAdherenceExplanationParams contains all the parameters to send to the API endpoint

	for the patch workforcemanagement agent adherence explanation operation.

	Typically these are written to a http.Request.
*/
type PatchWorkforcemanagementAgentAdherenceExplanationParams struct {

	/* AgentID.

	   The ID of the agent to query
	*/
	AgentID string

	/* Body.

	   The request body
	*/
	Body *models.UpdateAdherenceExplanationStatusRequest

	/* ExplanationID.

	   The ID of the explanation to update
	*/
	ExplanationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch workforcemanagement agent adherence explanation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) WithDefaults() *PatchWorkforcemanagementAgentAdherenceExplanationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch workforcemanagement agent adherence explanation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch workforcemanagement agent adherence explanation params
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) WithTimeout(timeout time.Duration) *PatchWorkforcemanagementAgentAdherenceExplanationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch workforcemanagement agent adherence explanation params
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch workforcemanagement agent adherence explanation params
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) WithContext(ctx context.Context) *PatchWorkforcemanagementAgentAdherenceExplanationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch workforcemanagement agent adherence explanation params
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch workforcemanagement agent adherence explanation params
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) WithHTTPClient(client *http.Client) *PatchWorkforcemanagementAgentAdherenceExplanationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch workforcemanagement agent adherence explanation params
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the patch workforcemanagement agent adherence explanation params
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) WithAgentID(agentID string) *PatchWorkforcemanagementAgentAdherenceExplanationParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the patch workforcemanagement agent adherence explanation params
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) SetAgentID(agentID string) {
	o.AgentID = agentID
}

// WithBody adds the body to the patch workforcemanagement agent adherence explanation params
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) WithBody(body *models.UpdateAdherenceExplanationStatusRequest) *PatchWorkforcemanagementAgentAdherenceExplanationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch workforcemanagement agent adherence explanation params
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) SetBody(body *models.UpdateAdherenceExplanationStatusRequest) {
	o.Body = body
}

// WithExplanationID adds the explanationID to the patch workforcemanagement agent adherence explanation params
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) WithExplanationID(explanationID string) *PatchWorkforcemanagementAgentAdherenceExplanationParams {
	o.SetExplanationID(explanationID)
	return o
}

// SetExplanationID adds the explanationId to the patch workforcemanagement agent adherence explanation params
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) SetExplanationID(explanationID string) {
	o.ExplanationID = explanationID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWorkforcemanagementAgentAdherenceExplanationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentId
	if err := r.SetPathParam("agentId", o.AgentID); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param explanationId
	if err := r.SetPathParam("explanationId", o.ExplanationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
