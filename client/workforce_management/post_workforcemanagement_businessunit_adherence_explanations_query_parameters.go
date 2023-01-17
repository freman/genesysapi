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

// NewPostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams creates a new PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams() *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams {
	return &PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParamsWithTimeout creates a new PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams {
	return &PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParamsWithContext creates a new PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParamsWithContext(ctx context.Context) *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams {
	return &PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParamsWithHTTPClient creates a new PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams {
	return &PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement businessunit adherence explanations query operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams struct {

	/* Body.

	   The request body
	*/
	Body *models.BuQueryAdherenceExplanationsRequest

	/* BusinessUnitID.

	   The ID of the business unit
	*/
	BusinessUnitID string

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

// WithDefaults hydrates default values in the post workforcemanagement businessunit adherence explanations query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) WithDefaults() *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement businessunit adherence explanations query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) WithContext(ctx context.Context) *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) WithBody(body *models.BuQueryAdherenceExplanationsRequest) *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) SetBody(body *models.BuQueryAdherenceExplanationsRequest) {
	o.Body = body
}

// WithBusinessUnitID adds the businessUnitID to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) WithBusinessUnitID(businessUnitID string) *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithForceAsync adds the forceAsync to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) WithForceAsync(forceAsync *bool) *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams {
	o.SetForceAsync(forceAsync)
	return o
}

// SetForceAsync adds the forceAsync to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) SetForceAsync(forceAsync *bool) {
	o.ForceAsync = forceAsync
}

// WithForceDownloadService adds the forceDownloadService to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) WithForceDownloadService(forceDownloadService *bool) *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams {
	o.SetForceDownloadService(forceDownloadService)
	return o
}

// SetForceDownloadService adds the forceDownloadService to the post workforcemanagement businessunit adherence explanations query params
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) SetForceDownloadService(forceDownloadService *bool) {
	o.ForceDownloadService = forceDownloadService
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementBusinessunitAdherenceExplanationsQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param businessUnitId
	if err := r.SetPathParam("businessUnitId", o.BusinessUnitID); err != nil {
		return err
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
