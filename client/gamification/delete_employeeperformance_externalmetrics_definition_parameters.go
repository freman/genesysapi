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

// NewDeleteEmployeeperformanceExternalmetricsDefinitionParams creates a new DeleteEmployeeperformanceExternalmetricsDefinitionParams object
// with the default values initialized.
func NewDeleteEmployeeperformanceExternalmetricsDefinitionParams() *DeleteEmployeeperformanceExternalmetricsDefinitionParams {
	var ()
	return &DeleteEmployeeperformanceExternalmetricsDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEmployeeperformanceExternalmetricsDefinitionParamsWithTimeout creates a new DeleteEmployeeperformanceExternalmetricsDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEmployeeperformanceExternalmetricsDefinitionParamsWithTimeout(timeout time.Duration) *DeleteEmployeeperformanceExternalmetricsDefinitionParams {
	var ()
	return &DeleteEmployeeperformanceExternalmetricsDefinitionParams{

		timeout: timeout,
	}
}

// NewDeleteEmployeeperformanceExternalmetricsDefinitionParamsWithContext creates a new DeleteEmployeeperformanceExternalmetricsDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEmployeeperformanceExternalmetricsDefinitionParamsWithContext(ctx context.Context) *DeleteEmployeeperformanceExternalmetricsDefinitionParams {
	var ()
	return &DeleteEmployeeperformanceExternalmetricsDefinitionParams{

		Context: ctx,
	}
}

// NewDeleteEmployeeperformanceExternalmetricsDefinitionParamsWithHTTPClient creates a new DeleteEmployeeperformanceExternalmetricsDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEmployeeperformanceExternalmetricsDefinitionParamsWithHTTPClient(client *http.Client) *DeleteEmployeeperformanceExternalmetricsDefinitionParams {
	var ()
	return &DeleteEmployeeperformanceExternalmetricsDefinitionParams{
		HTTPClient: client,
	}
}

/*DeleteEmployeeperformanceExternalmetricsDefinitionParams contains all the parameters to send to the API endpoint
for the delete employeeperformance externalmetrics definition operation typically these are written to a http.Request
*/
type DeleteEmployeeperformanceExternalmetricsDefinitionParams struct {

	/*MetricID
	  Specifies the External Metric Definition ID

	*/
	MetricID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete employeeperformance externalmetrics definition params
func (o *DeleteEmployeeperformanceExternalmetricsDefinitionParams) WithTimeout(timeout time.Duration) *DeleteEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete employeeperformance externalmetrics definition params
func (o *DeleteEmployeeperformanceExternalmetricsDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete employeeperformance externalmetrics definition params
func (o *DeleteEmployeeperformanceExternalmetricsDefinitionParams) WithContext(ctx context.Context) *DeleteEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete employeeperformance externalmetrics definition params
func (o *DeleteEmployeeperformanceExternalmetricsDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete employeeperformance externalmetrics definition params
func (o *DeleteEmployeeperformanceExternalmetricsDefinitionParams) WithHTTPClient(client *http.Client) *DeleteEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete employeeperformance externalmetrics definition params
func (o *DeleteEmployeeperformanceExternalmetricsDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetricID adds the metricID to the delete employeeperformance externalmetrics definition params
func (o *DeleteEmployeeperformanceExternalmetricsDefinitionParams) WithMetricID(metricID string) *DeleteEmployeeperformanceExternalmetricsDefinitionParams {
	o.SetMetricID(metricID)
	return o
}

// SetMetricID adds the metricId to the delete employeeperformance externalmetrics definition params
func (o *DeleteEmployeeperformanceExternalmetricsDefinitionParams) SetMetricID(metricID string) {
	o.MetricID = metricID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEmployeeperformanceExternalmetricsDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
