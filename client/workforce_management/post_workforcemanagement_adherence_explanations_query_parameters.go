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
	"github.com/go-openapi/swag"

	"github.com/freman/genesysapi/models"
)

// NewPostWorkforcemanagementAdherenceExplanationsQueryParams creates a new PostWorkforcemanagementAdherenceExplanationsQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementAdherenceExplanationsQueryParams() *PostWorkforcemanagementAdherenceExplanationsQueryParams {
	return &PostWorkforcemanagementAdherenceExplanationsQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementAdherenceExplanationsQueryParamsWithTimeout creates a new PostWorkforcemanagementAdherenceExplanationsQueryParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementAdherenceExplanationsQueryParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementAdherenceExplanationsQueryParams {
	return &PostWorkforcemanagementAdherenceExplanationsQueryParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementAdherenceExplanationsQueryParamsWithContext creates a new PostWorkforcemanagementAdherenceExplanationsQueryParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementAdherenceExplanationsQueryParamsWithContext(ctx context.Context) *PostWorkforcemanagementAdherenceExplanationsQueryParams {
	return &PostWorkforcemanagementAdherenceExplanationsQueryParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementAdherenceExplanationsQueryParamsWithHTTPClient creates a new PostWorkforcemanagementAdherenceExplanationsQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementAdherenceExplanationsQueryParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementAdherenceExplanationsQueryParams {
	return &PostWorkforcemanagementAdherenceExplanationsQueryParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementAdherenceExplanationsQueryParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement adherence explanations query operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementAdherenceExplanationsQueryParams struct {

	/* Body.

	   The request body
	*/
	Body *models.AgentQueryAdherenceExplanationsRequest

	/* ForceAsync.

	   Force the result of this operation to be sent asynchronously via notification. For testing/app development purposes
	*/
	ForceAsync *bool

	/* ForceDownloadService.

	   Force the result of this operation to be sent via download service. For testing/app development purposes
	*/
	ForceDownloadService *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workforcemanagement adherence explanations query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) WithDefaults() *PostWorkforcemanagementAdherenceExplanationsQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement adherence explanations query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workforcemanagement adherence explanations query params
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementAdherenceExplanationsQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement adherence explanations query params
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement adherence explanations query params
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) WithContext(ctx context.Context) *PostWorkforcemanagementAdherenceExplanationsQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement adherence explanations query params
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement adherence explanations query params
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementAdherenceExplanationsQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement adherence explanations query params
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement adherence explanations query params
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) WithBody(body *models.AgentQueryAdherenceExplanationsRequest) *PostWorkforcemanagementAdherenceExplanationsQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement adherence explanations query params
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) SetBody(body *models.AgentQueryAdherenceExplanationsRequest) {
	o.Body = body
}

// WithForceAsync adds the forceAsync to the post workforcemanagement adherence explanations query params
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) WithForceAsync(forceAsync *bool) *PostWorkforcemanagementAdherenceExplanationsQueryParams {
	o.SetForceAsync(forceAsync)
	return o
}

// SetForceAsync adds the forceAsync to the post workforcemanagement adherence explanations query params
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) SetForceAsync(forceAsync *bool) {
	o.ForceAsync = forceAsync
}

// WithForceDownloadService adds the forceDownloadService to the post workforcemanagement adherence explanations query params
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) WithForceDownloadService(forceDownloadService *bool) *PostWorkforcemanagementAdherenceExplanationsQueryParams {
	o.SetForceDownloadService(forceDownloadService)
	return o
}

// SetForceDownloadService adds the forceDownloadService to the post workforcemanagement adherence explanations query params
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) SetForceDownloadService(forceDownloadService *bool) {
	o.ForceDownloadService = forceDownloadService
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementAdherenceExplanationsQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ForceAsync != nil {

		// query param forceAsync
		var qrForceAsync bool

		if o.ForceAsync != nil {
			qrForceAsync = *o.ForceAsync
		}
		qForceAsync := swag.FormatBool(qrForceAsync)
		if qForceAsync != "" {

			if err := r.SetQueryParam("forceAsync", qForceAsync); err != nil {
				return err
			}
		}
	}

	if o.ForceDownloadService != nil {

		// query param forceDownloadService
		var qrForceDownloadService bool

		if o.ForceDownloadService != nil {
			qrForceDownloadService = *o.ForceDownloadService
		}
		qForceDownloadService := swag.FormatBool(qrForceDownloadService)
		if qForceDownloadService != "" {

			if err := r.SetQueryParam("forceDownloadService", qForceDownloadService); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
