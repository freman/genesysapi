// Code generated by go-swagger; DO NOT EDIT.

package gamification

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

// NewPatchEmployeeperformanceExternalmetricsDefinitionParams creates a new PatchEmployeeperformanceExternalmetricsDefinitionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchEmployeeperformanceExternalmetricsDefinitionParams() *PatchEmployeeperformanceExternalmetricsDefinitionParams {
	return &PatchEmployeeperformanceExternalmetricsDefinitionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEmployeeperformanceExternalmetricsDefinitionParamsWithTimeout creates a new PatchEmployeeperformanceExternalmetricsDefinitionParams object
// with the ability to set a timeout on a request.
func NewPatchEmployeeperformanceExternalmetricsDefinitionParamsWithTimeout(timeout time.Duration) *PatchEmployeeperformanceExternalmetricsDefinitionParams {
	return &PatchEmployeeperformanceExternalmetricsDefinitionParams{
		timeout: timeout,
	}
}

// NewPatchEmployeeperformanceExternalmetricsDefinitionParamsWithContext creates a new PatchEmployeeperformanceExternalmetricsDefinitionParams object
// with the ability to set a context for a request.
func NewPatchEmployeeperformanceExternalmetricsDefinitionParamsWithContext(ctx context.Context) *PatchEmployeeperformanceExternalmetricsDefinitionParams {
	return &PatchEmployeeperformanceExternalmetricsDefinitionParams{
		Context: ctx,
	}
}

// NewPatchEmployeeperformanceExternalmetricsDefinitionParamsWithHTTPClient creates a new PatchEmployeeperformanceExternalmetricsDefinitionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchEmployeeperformanceExternalmetricsDefinitionParamsWithHTTPClient(client *http.Client) *PatchEmployeeperformanceExternalmetricsDefinitionParams {
	return &PatchEmployeeperformanceExternalmetricsDefinitionParams{
		HTTPClient: client,
	}
}

/*
PatchEmployeeperformanceExternalmetricsDefinitionParams contains all the parameters to send to the API endpoint

	for the patch employeeperformance externalmetrics definition operation.

	Typically these are written to a http.Request.
*/
type PatchEmployeeperformanceExternalmetricsDefinitionParams struct {

	/* Body.

	   The External Metric Definition parameters to be updated
	*/
	Body *models.ExternalMetricDefinitionUpdateRequest

	/* MetricID.

	   Specifies the metric definition ID
	*/
	MetricID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch employeeperformance externalmetrics definition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) WithDefaults() *PatchEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch employeeperformance externalmetrics definition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch employeeperformance externalmetrics definition params
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) WithTimeout(timeout time.Duration) *PatchEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch employeeperformance externalmetrics definition params
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch employeeperformance externalmetrics definition params
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) WithContext(ctx context.Context) *PatchEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch employeeperformance externalmetrics definition params
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch employeeperformance externalmetrics definition params
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) WithHTTPClient(client *http.Client) *PatchEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch employeeperformance externalmetrics definition params
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch employeeperformance externalmetrics definition params
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) WithBody(body *models.ExternalMetricDefinitionUpdateRequest) *PatchEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch employeeperformance externalmetrics definition params
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) SetBody(body *models.ExternalMetricDefinitionUpdateRequest) {
	o.Body = body
}

// WithMetricID adds the metricID to the patch employeeperformance externalmetrics definition params
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) WithMetricID(metricID string) *PatchEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetMetricID(metricID)
	return o
}

// SetMetricID adds the metricId to the patch employeeperformance externalmetrics definition params
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) SetMetricID(metricID string) {
	o.MetricID = metricID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEmployeeperformanceExternalmetricsDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param metricId
	if err := r.SetPathParam("metricId", o.MetricID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
