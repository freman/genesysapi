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

// NewPostWorkforcemanagementBusinessunitServicegoaltemplatesParams creates a new PostWorkforcemanagementBusinessunitServicegoaltemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementBusinessunitServicegoaltemplatesParams() *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams {
	return &PostWorkforcemanagementBusinessunitServicegoaltemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementBusinessunitServicegoaltemplatesParamsWithTimeout creates a new PostWorkforcemanagementBusinessunitServicegoaltemplatesParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementBusinessunitServicegoaltemplatesParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams {
	return &PostWorkforcemanagementBusinessunitServicegoaltemplatesParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementBusinessunitServicegoaltemplatesParamsWithContext creates a new PostWorkforcemanagementBusinessunitServicegoaltemplatesParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementBusinessunitServicegoaltemplatesParamsWithContext(ctx context.Context) *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams {
	return &PostWorkforcemanagementBusinessunitServicegoaltemplatesParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementBusinessunitServicegoaltemplatesParamsWithHTTPClient creates a new PostWorkforcemanagementBusinessunitServicegoaltemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementBusinessunitServicegoaltemplatesParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams {
	return &PostWorkforcemanagementBusinessunitServicegoaltemplatesParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementBusinessunitServicegoaltemplatesParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement businessunit servicegoaltemplates operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementBusinessunitServicegoaltemplatesParams struct {

	/* Body.

	   body
	*/
	Body *models.CreateServiceGoalTemplate

	/* BusinessUnitID.

	   The ID of the business unit.
	*/
	BusinessUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workforcemanagement businessunit servicegoaltemplates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) WithDefaults() *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement businessunit servicegoaltemplates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workforcemanagement businessunit servicegoaltemplates params
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement businessunit servicegoaltemplates params
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement businessunit servicegoaltemplates params
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) WithContext(ctx context.Context) *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement businessunit servicegoaltemplates params
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement businessunit servicegoaltemplates params
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement businessunit servicegoaltemplates params
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement businessunit servicegoaltemplates params
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) WithBody(body *models.CreateServiceGoalTemplate) *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement businessunit servicegoaltemplates params
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) SetBody(body *models.CreateServiceGoalTemplate) {
	o.Body = body
}

// WithBusinessUnitID adds the businessUnitID to the post workforcemanagement businessunit servicegoaltemplates params
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) WithBusinessUnitID(businessUnitID string) *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the post workforcemanagement businessunit servicegoaltemplates params
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementBusinessunitServicegoaltemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
