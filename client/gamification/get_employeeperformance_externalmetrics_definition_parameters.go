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
)

// NewGetEmployeeperformanceExternalmetricsDefinitionParams creates a new GetEmployeeperformanceExternalmetricsDefinitionParams object
// with the default values initialized.
func NewGetEmployeeperformanceExternalmetricsDefinitionParams() *GetEmployeeperformanceExternalmetricsDefinitionParams {
	var ()
	return &GetEmployeeperformanceExternalmetricsDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmployeeperformanceExternalmetricsDefinitionParamsWithTimeout creates a new GetEmployeeperformanceExternalmetricsDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmployeeperformanceExternalmetricsDefinitionParamsWithTimeout(timeout time.Duration) *GetEmployeeperformanceExternalmetricsDefinitionParams {
	var ()
	return &GetEmployeeperformanceExternalmetricsDefinitionParams{

		timeout: timeout,
	}
}

// NewGetEmployeeperformanceExternalmetricsDefinitionParamsWithContext creates a new GetEmployeeperformanceExternalmetricsDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmployeeperformanceExternalmetricsDefinitionParamsWithContext(ctx context.Context) *GetEmployeeperformanceExternalmetricsDefinitionParams {
	var ()
	return &GetEmployeeperformanceExternalmetricsDefinitionParams{

		Context: ctx,
	}
}

// NewGetEmployeeperformanceExternalmetricsDefinitionParamsWithHTTPClient creates a new GetEmployeeperformanceExternalmetricsDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmployeeperformanceExternalmetricsDefinitionParamsWithHTTPClient(client *http.Client) *GetEmployeeperformanceExternalmetricsDefinitionParams {
	var ()
	return &GetEmployeeperformanceExternalmetricsDefinitionParams{
		HTTPClient: client,
	}
}

/*GetEmployeeperformanceExternalmetricsDefinitionParams contains all the parameters to send to the API endpoint
for the get employeeperformance externalmetrics definition operation typically these are written to a http.Request
*/
type GetEmployeeperformanceExternalmetricsDefinitionParams struct {

	/*MetricID
	  Specifies the External Metric Definition ID

	*/
	MetricID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get employeeperformance externalmetrics definition params
func (o *GetEmployeeperformanceExternalmetricsDefinitionParams) WithTimeout(timeout time.Duration) *GetEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get employeeperformance externalmetrics definition params
func (o *GetEmployeeperformanceExternalmetricsDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get employeeperformance externalmetrics definition params
func (o *GetEmployeeperformanceExternalmetricsDefinitionParams) WithContext(ctx context.Context) *GetEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get employeeperformance externalmetrics definition params
func (o *GetEmployeeperformanceExternalmetricsDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get employeeperformance externalmetrics definition params
func (o *GetEmployeeperformanceExternalmetricsDefinitionParams) WithHTTPClient(client *http.Client) *GetEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get employeeperformance externalmetrics definition params
func (o *GetEmployeeperformanceExternalmetricsDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetricID adds the metricID to the get employeeperformance externalmetrics definition params
func (o *GetEmployeeperformanceExternalmetricsDefinitionParams) WithMetricID(metricID string) *GetEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetMetricID(metricID)
	return o
}

// SetMetricID adds the metricId to the get employeeperformance externalmetrics definition params
func (o *GetEmployeeperformanceExternalmetricsDefinitionParams) SetMetricID(metricID string) {
	o.MetricID = metricID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmployeeperformanceExternalmetricsDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param metricId
	if err := r.SetPathParam("metricId", o.MetricID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}