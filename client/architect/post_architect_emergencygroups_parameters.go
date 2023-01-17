// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewPostArchitectEmergencygroupsParams creates a new PostArchitectEmergencygroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostArchitectEmergencygroupsParams() *PostArchitectEmergencygroupsParams {
	return &PostArchitectEmergencygroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostArchitectEmergencygroupsParamsWithTimeout creates a new PostArchitectEmergencygroupsParams object
// with the ability to set a timeout on a request.
func NewPostArchitectEmergencygroupsParamsWithTimeout(timeout time.Duration) *PostArchitectEmergencygroupsParams {
	return &PostArchitectEmergencygroupsParams{
		timeout: timeout,
	}
}

// NewPostArchitectEmergencygroupsParamsWithContext creates a new PostArchitectEmergencygroupsParams object
// with the ability to set a context for a request.
func NewPostArchitectEmergencygroupsParamsWithContext(ctx context.Context) *PostArchitectEmergencygroupsParams {
	return &PostArchitectEmergencygroupsParams{
		Context: ctx,
	}
}

// NewPostArchitectEmergencygroupsParamsWithHTTPClient creates a new PostArchitectEmergencygroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostArchitectEmergencygroupsParamsWithHTTPClient(client *http.Client) *PostArchitectEmergencygroupsParams {
	return &PostArchitectEmergencygroupsParams{
		HTTPClient: client,
	}
}

/*
PostArchitectEmergencygroupsParams contains all the parameters to send to the API endpoint

	for the post architect emergencygroups operation.

	Typically these are written to a http.Request.
*/
type PostArchitectEmergencygroupsParams struct {

	// Body.
	Body *models.EmergencyGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post architect emergencygroups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostArchitectEmergencygroupsParams) WithDefaults() *PostArchitectEmergencygroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post architect emergencygroups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostArchitectEmergencygroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post architect emergencygroups params
func (o *PostArchitectEmergencygroupsParams) WithTimeout(timeout time.Duration) *PostArchitectEmergencygroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post architect emergencygroups params
func (o *PostArchitectEmergencygroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post architect emergencygroups params
func (o *PostArchitectEmergencygroupsParams) WithContext(ctx context.Context) *PostArchitectEmergencygroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post architect emergencygroups params
func (o *PostArchitectEmergencygroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post architect emergencygroups params
func (o *PostArchitectEmergencygroupsParams) WithHTTPClient(client *http.Client) *PostArchitectEmergencygroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post architect emergencygroups params
func (o *PostArchitectEmergencygroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post architect emergencygroups params
func (o *PostArchitectEmergencygroupsParams) WithBody(body *models.EmergencyGroup) *PostArchitectEmergencygroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post architect emergencygroups params
func (o *PostArchitectEmergencygroupsParams) SetBody(body *models.EmergencyGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostArchitectEmergencygroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
