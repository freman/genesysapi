// Code generated by go-swagger; DO NOT EDIT.

package chat

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

// NewPatchChatSettingsParams creates a new PatchChatSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchChatSettingsParams() *PatchChatSettingsParams {
	return &PatchChatSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchChatSettingsParamsWithTimeout creates a new PatchChatSettingsParams object
// with the ability to set a timeout on a request.
func NewPatchChatSettingsParamsWithTimeout(timeout time.Duration) *PatchChatSettingsParams {
	return &PatchChatSettingsParams{
		timeout: timeout,
	}
}

// NewPatchChatSettingsParamsWithContext creates a new PatchChatSettingsParams object
// with the ability to set a context for a request.
func NewPatchChatSettingsParamsWithContext(ctx context.Context) *PatchChatSettingsParams {
	return &PatchChatSettingsParams{
		Context: ctx,
	}
}

// NewPatchChatSettingsParamsWithHTTPClient creates a new PatchChatSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchChatSettingsParamsWithHTTPClient(client *http.Client) *PatchChatSettingsParams {
	return &PatchChatSettingsParams{
		HTTPClient: client,
	}
}

/*
PatchChatSettingsParams contains all the parameters to send to the API endpoint

	for the patch chat settings operation.

	Typically these are written to a http.Request.
*/
type PatchChatSettingsParams struct {

	/* Body.

	   Chat
	*/
	Body *models.ChatSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch chat settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchChatSettingsParams) WithDefaults() *PatchChatSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch chat settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchChatSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch chat settings params
func (o *PatchChatSettingsParams) WithTimeout(timeout time.Duration) *PatchChatSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch chat settings params
func (o *PatchChatSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch chat settings params
func (o *PatchChatSettingsParams) WithContext(ctx context.Context) *PatchChatSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch chat settings params
func (o *PatchChatSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch chat settings params
func (o *PatchChatSettingsParams) WithHTTPClient(client *http.Client) *PatchChatSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch chat settings params
func (o *PatchChatSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch chat settings params
func (o *PatchChatSettingsParams) WithBody(body *models.ChatSettings) *PatchChatSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch chat settings params
func (o *PatchChatSettingsParams) SetBody(body *models.ChatSettings) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchChatSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
