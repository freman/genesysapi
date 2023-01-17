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

// NewPostWorkforcemanagementAdherenceExplanationsParams creates a new PostWorkforcemanagementAdherenceExplanationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementAdherenceExplanationsParams() *PostWorkforcemanagementAdherenceExplanationsParams {
	return &PostWorkforcemanagementAdherenceExplanationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementAdherenceExplanationsParamsWithTimeout creates a new PostWorkforcemanagementAdherenceExplanationsParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementAdherenceExplanationsParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementAdherenceExplanationsParams {
	return &PostWorkforcemanagementAdherenceExplanationsParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementAdherenceExplanationsParamsWithContext creates a new PostWorkforcemanagementAdherenceExplanationsParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementAdherenceExplanationsParamsWithContext(ctx context.Context) *PostWorkforcemanagementAdherenceExplanationsParams {
	return &PostWorkforcemanagementAdherenceExplanationsParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementAdherenceExplanationsParamsWithHTTPClient creates a new PostWorkforcemanagementAdherenceExplanationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementAdherenceExplanationsParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementAdherenceExplanationsParams {
	return &PostWorkforcemanagementAdherenceExplanationsParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementAdherenceExplanationsParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement adherence explanations operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementAdherenceExplanationsParams struct {

	/* Body.

	   The request body
	*/
	Body *models.AddAdherenceExplanationAgentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workforcemanagement adherence explanations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementAdherenceExplanationsParams) WithDefaults() *PostWorkforcemanagementAdherenceExplanationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement adherence explanations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementAdherenceExplanationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workforcemanagement adherence explanations params
func (o *PostWorkforcemanagementAdherenceExplanationsParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementAdherenceExplanationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement adherence explanations params
func (o *PostWorkforcemanagementAdherenceExplanationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement adherence explanations params
func (o *PostWorkforcemanagementAdherenceExplanationsParams) WithContext(ctx context.Context) *PostWorkforcemanagementAdherenceExplanationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement adherence explanations params
func (o *PostWorkforcemanagementAdherenceExplanationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement adherence explanations params
func (o *PostWorkforcemanagementAdherenceExplanationsParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementAdherenceExplanationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement adherence explanations params
func (o *PostWorkforcemanagementAdherenceExplanationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement adherence explanations params
func (o *PostWorkforcemanagementAdherenceExplanationsParams) WithBody(body *models.AddAdherenceExplanationAgentRequest) *PostWorkforcemanagementAdherenceExplanationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement adherence explanations params
func (o *PostWorkforcemanagementAdherenceExplanationsParams) SetBody(body *models.AddAdherenceExplanationAgentRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementAdherenceExplanationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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