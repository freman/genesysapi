// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// NewPatchAuthorizationSettingsParams creates a new PatchAuthorizationSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAuthorizationSettingsParams() *PatchAuthorizationSettingsParams {
	return &PatchAuthorizationSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAuthorizationSettingsParamsWithTimeout creates a new PatchAuthorizationSettingsParams object
// with the ability to set a timeout on a request.
func NewPatchAuthorizationSettingsParamsWithTimeout(timeout time.Duration) *PatchAuthorizationSettingsParams {
	return &PatchAuthorizationSettingsParams{
		timeout: timeout,
	}
}

// NewPatchAuthorizationSettingsParamsWithContext creates a new PatchAuthorizationSettingsParams object
// with the ability to set a context for a request.
func NewPatchAuthorizationSettingsParamsWithContext(ctx context.Context) *PatchAuthorizationSettingsParams {
	return &PatchAuthorizationSettingsParams{
		Context: ctx,
	}
}

// NewPatchAuthorizationSettingsParamsWithHTTPClient creates a new PatchAuthorizationSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAuthorizationSettingsParamsWithHTTPClient(client *http.Client) *PatchAuthorizationSettingsParams {
	return &PatchAuthorizationSettingsParams{
		HTTPClient: client,
	}
}

/*
PatchAuthorizationSettingsParams contains all the parameters to send to the API endpoint

	for the patch authorization settings operation.

	Typically these are written to a http.Request.
*/
type PatchAuthorizationSettingsParams struct {

	/* Body.

	   Authorization Settings
	*/
	Body *models.AuthorizationSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch authorization settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAuthorizationSettingsParams) WithDefaults() *PatchAuthorizationSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch authorization settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAuthorizationSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch authorization settings params
func (o *PatchAuthorizationSettingsParams) WithTimeout(timeout time.Duration) *PatchAuthorizationSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch authorization settings params
func (o *PatchAuthorizationSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch authorization settings params
func (o *PatchAuthorizationSettingsParams) WithContext(ctx context.Context) *PatchAuthorizationSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch authorization settings params
func (o *PatchAuthorizationSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch authorization settings params
func (o *PatchAuthorizationSettingsParams) WithHTTPClient(client *http.Client) *PatchAuthorizationSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch authorization settings params
func (o *PatchAuthorizationSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch authorization settings params
func (o *PatchAuthorizationSettingsParams) WithBody(body *models.AuthorizationSettings) *PatchAuthorizationSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch authorization settings params
func (o *PatchAuthorizationSettingsParams) SetBody(body *models.AuthorizationSettings) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAuthorizationSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
