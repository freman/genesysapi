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

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsParams creates a new PostWorkforcemanagementAgentsMePossibleworkshiftsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsParams() *PostWorkforcemanagementAgentsMePossibleworkshiftsParams {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsParamsWithTimeout creates a new PostWorkforcemanagementAgentsMePossibleworkshiftsParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementAgentsMePossibleworkshiftsParams {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsParamsWithContext creates a new PostWorkforcemanagementAgentsMePossibleworkshiftsParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsParamsWithContext(ctx context.Context) *PostWorkforcemanagementAgentsMePossibleworkshiftsParams {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementAgentsMePossibleworkshiftsParamsWithHTTPClient creates a new PostWorkforcemanagementAgentsMePossibleworkshiftsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementAgentsMePossibleworkshiftsParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementAgentsMePossibleworkshiftsParams {
	return &PostWorkforcemanagementAgentsMePossibleworkshiftsParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementAgentsMePossibleworkshiftsParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement agents me possibleworkshifts operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementAgentsMePossibleworkshiftsParams struct {

	/* Body.

	   body
	*/
	Body *models.AgentPossibleWorkShiftsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workforcemanagement agents me possibleworkshifts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsParams) WithDefaults() *PostWorkforcemanagementAgentsMePossibleworkshiftsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement agents me possibleworkshifts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workforcemanagement agents me possibleworkshifts params
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementAgentsMePossibleworkshiftsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement agents me possibleworkshifts params
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement agents me possibleworkshifts params
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsParams) WithContext(ctx context.Context) *PostWorkforcemanagementAgentsMePossibleworkshiftsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement agents me possibleworkshifts params
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement agents me possibleworkshifts params
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementAgentsMePossibleworkshiftsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement agents me possibleworkshifts params
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement agents me possibleworkshifts params
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsParams) WithBody(body *models.AgentPossibleWorkShiftsRequest) *PostWorkforcemanagementAgentsMePossibleworkshiftsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement agents me possibleworkshifts params
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsParams) SetBody(body *models.AgentPossibleWorkShiftsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementAgentsMePossibleworkshiftsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
