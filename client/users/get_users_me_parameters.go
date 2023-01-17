// Code generated by go-swagger; DO NOT EDIT.

package users

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
	"github.com/go-openapi/swag"
)

// NewGetUsersMeParams creates a new GetUsersMeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsersMeParams() *GetUsersMeParams {
	return &GetUsersMeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersMeParamsWithTimeout creates a new GetUsersMeParams object
// with the ability to set a timeout on a request.
func NewGetUsersMeParamsWithTimeout(timeout time.Duration) *GetUsersMeParams {
	return &GetUsersMeParams{
		timeout: timeout,
	}
}

// NewGetUsersMeParamsWithContext creates a new GetUsersMeParams object
// with the ability to set a context for a request.
func NewGetUsersMeParamsWithContext(ctx context.Context) *GetUsersMeParams {
	return &GetUsersMeParams{
		Context: ctx,
	}
}

// NewGetUsersMeParamsWithHTTPClient creates a new GetUsersMeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsersMeParamsWithHTTPClient(client *http.Client) *GetUsersMeParams {
	return &GetUsersMeParams{
		HTTPClient: client,
	}
}

/*
GetUsersMeParams contains all the parameters to send to the API endpoint

	for the get users me operation.

	Typically these are written to a http.Request.
*/
type GetUsersMeParams struct {

	/* Expand.

	   Which fields, if any, to expand.
	*/
	Expand []string

	/* IntegrationPresenceSource.

	   Get your presence for a given integration. This parameter will only be used when presence is provided as an "expand".
	*/
	IntegrationPresenceSource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get users me params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersMeParams) WithDefaults() *GetUsersMeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get users me params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersMeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get users me params
func (o *GetUsersMeParams) WithTimeout(timeout time.Duration) *GetUsersMeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users me params
func (o *GetUsersMeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users me params
func (o *GetUsersMeParams) WithContext(ctx context.Context) *GetUsersMeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users me params
func (o *GetUsersMeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users me params
func (o *GetUsersMeParams) WithHTTPClient(client *http.Client) *GetUsersMeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users me params
func (o *GetUsersMeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get users me params
func (o *GetUsersMeParams) WithExpand(expand []string) *GetUsersMeParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get users me params
func (o *GetUsersMeParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithIntegrationPresenceSource adds the integrationPresenceSource to the get users me params
func (o *GetUsersMeParams) WithIntegrationPresenceSource(integrationPresenceSource *string) *GetUsersMeParams {
	o.SetIntegrationPresenceSource(integrationPresenceSource)
	return o
}

// SetIntegrationPresenceSource adds the integrationPresenceSource to the get users me params
func (o *GetUsersMeParams) SetIntegrationPresenceSource(integrationPresenceSource *string) {
	o.IntegrationPresenceSource = integrationPresenceSource
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersMeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	if o.IntegrationPresenceSource != nil {

		// query param integrationPresenceSource
		var qrIntegrationPresenceSource string

		if o.IntegrationPresenceSource != nil {
			qrIntegrationPresenceSource = *o.IntegrationPresenceSource
		}
		qIntegrationPresenceSource := qrIntegrationPresenceSource
		if qIntegrationPresenceSource != "" {

			if err := r.SetQueryParam("integrationPresenceSource", qIntegrationPresenceSource); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetUsersMe binds the parameter expand
func (o *GetUsersMeParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}
