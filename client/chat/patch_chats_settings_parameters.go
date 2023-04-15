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

// NewPatchChatsSettingsParams creates a new PatchChatsSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchChatsSettingsParams() *PatchChatsSettingsParams {
	return &PatchChatsSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchChatsSettingsParamsWithTimeout creates a new PatchChatsSettingsParams object
// with the ability to set a timeout on a request.
func NewPatchChatsSettingsParamsWithTimeout(timeout time.Duration) *PatchChatsSettingsParams {
	return &PatchChatsSettingsParams{
		timeout: timeout,
	}
}

// NewPatchChatsSettingsParamsWithContext creates a new PatchChatsSettingsParams object
// with the ability to set a context for a request.
func NewPatchChatsSettingsParamsWithContext(ctx context.Context) *PatchChatsSettingsParams {
	return &PatchChatsSettingsParams{
		Context: ctx,
	}
}

// NewPatchChatsSettingsParamsWithHTTPClient creates a new PatchChatsSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchChatsSettingsParamsWithHTTPClient(client *http.Client) *PatchChatsSettingsParams {
	return &PatchChatsSettingsParams{
		HTTPClient: client,
	}
}

/*
PatchChatsSettingsParams contains all the parameters to send to the API endpoint

	for the patch chats settings operation.

	Typically these are written to a http.Request.
*/
type PatchChatsSettingsParams struct {

	/* Body.

	   Chat
	*/
	Body *models.ChatSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch chats settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchChatsSettingsParams) WithDefaults() *PatchChatsSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch chats settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchChatsSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch chats settings params
func (o *PatchChatsSettingsParams) WithTimeout(timeout time.Duration) *PatchChatsSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch chats settings params
func (o *PatchChatsSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch chats settings params
func (o *PatchChatsSettingsParams) WithContext(ctx context.Context) *PatchChatsSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch chats settings params
func (o *PatchChatsSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch chats settings params
func (o *PatchChatsSettingsParams) WithHTTPClient(client *http.Client) *PatchChatsSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch chats settings params
func (o *PatchChatsSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch chats settings params
func (o *PatchChatsSettingsParams) WithBody(body *models.ChatSettings) *PatchChatsSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch chats settings params
func (o *PatchChatsSettingsParams) SetBody(body *models.ChatSettings) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchChatsSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
