// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewPostIntegrationsActionsParams creates a new PostIntegrationsActionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostIntegrationsActionsParams() *PostIntegrationsActionsParams {
	return &PostIntegrationsActionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostIntegrationsActionsParamsWithTimeout creates a new PostIntegrationsActionsParams object
// with the ability to set a timeout on a request.
func NewPostIntegrationsActionsParamsWithTimeout(timeout time.Duration) *PostIntegrationsActionsParams {
	return &PostIntegrationsActionsParams{
		timeout: timeout,
	}
}

// NewPostIntegrationsActionsParamsWithContext creates a new PostIntegrationsActionsParams object
// with the ability to set a context for a request.
func NewPostIntegrationsActionsParamsWithContext(ctx context.Context) *PostIntegrationsActionsParams {
	return &PostIntegrationsActionsParams{
		Context: ctx,
	}
}

// NewPostIntegrationsActionsParamsWithHTTPClient creates a new PostIntegrationsActionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostIntegrationsActionsParamsWithHTTPClient(client *http.Client) *PostIntegrationsActionsParams {
	return &PostIntegrationsActionsParams{
		HTTPClient: client,
	}
}

/*
PostIntegrationsActionsParams contains all the parameters to send to the API endpoint

	for the post integrations actions operation.

	Typically these are written to a http.Request.
*/
type PostIntegrationsActionsParams struct {

	/* Body.

	   Input used to create Action.
	*/
	Body *models.PostActionInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post integrations actions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostIntegrationsActionsParams) WithDefaults() *PostIntegrationsActionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post integrations actions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostIntegrationsActionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post integrations actions params
func (o *PostIntegrationsActionsParams) WithTimeout(timeout time.Duration) *PostIntegrationsActionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post integrations actions params
func (o *PostIntegrationsActionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post integrations actions params
func (o *PostIntegrationsActionsParams) WithContext(ctx context.Context) *PostIntegrationsActionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post integrations actions params
func (o *PostIntegrationsActionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post integrations actions params
func (o *PostIntegrationsActionsParams) WithHTTPClient(client *http.Client) *PostIntegrationsActionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post integrations actions params
func (o *PostIntegrationsActionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post integrations actions params
func (o *PostIntegrationsActionsParams) WithBody(body *models.PostActionInput) *PostIntegrationsActionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post integrations actions params
func (o *PostIntegrationsActionsParams) SetBody(body *models.PostActionInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostIntegrationsActionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
