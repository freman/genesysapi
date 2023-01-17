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

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryParams creates a new PostWorkforcemanagementTimeofflimitsAvailableQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryParams() *PostWorkforcemanagementTimeofflimitsAvailableQueryParams {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryParamsWithTimeout creates a new PostWorkforcemanagementTimeofflimitsAvailableQueryParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementTimeofflimitsAvailableQueryParams {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryParamsWithContext creates a new PostWorkforcemanagementTimeofflimitsAvailableQueryParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryParamsWithContext(ctx context.Context) *PostWorkforcemanagementTimeofflimitsAvailableQueryParams {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementTimeofflimitsAvailableQueryParamsWithHTTPClient creates a new PostWorkforcemanagementTimeofflimitsAvailableQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementTimeofflimitsAvailableQueryParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementTimeofflimitsAvailableQueryParams {
	return &PostWorkforcemanagementTimeofflimitsAvailableQueryParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementTimeofflimitsAvailableQueryParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement timeofflimits available query operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementTimeofflimitsAvailableQueryParams struct {

	/* Body.

	   body
	*/
	Body *models.AvailableTimeOffRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workforcemanagement timeofflimits available query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryParams) WithDefaults() *PostWorkforcemanagementTimeofflimitsAvailableQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement timeofflimits available query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workforcemanagement timeofflimits available query params
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementTimeofflimitsAvailableQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement timeofflimits available query params
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement timeofflimits available query params
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryParams) WithContext(ctx context.Context) *PostWorkforcemanagementTimeofflimitsAvailableQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement timeofflimits available query params
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement timeofflimits available query params
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementTimeofflimitsAvailableQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement timeofflimits available query params
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement timeofflimits available query params
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryParams) WithBody(body *models.AvailableTimeOffRequest) *PostWorkforcemanagementTimeofflimitsAvailableQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement timeofflimits available query params
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryParams) SetBody(body *models.AvailableTimeOffRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementTimeofflimitsAvailableQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
