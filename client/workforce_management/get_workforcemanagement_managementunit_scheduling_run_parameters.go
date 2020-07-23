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
)

// NewGetWorkforcemanagementManagementunitSchedulingRunParams creates a new GetWorkforcemanagementManagementunitSchedulingRunParams object
// with the default values initialized.
func NewGetWorkforcemanagementManagementunitSchedulingRunParams() *GetWorkforcemanagementManagementunitSchedulingRunParams {
	var ()
	return &GetWorkforcemanagementManagementunitSchedulingRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementManagementunitSchedulingRunParamsWithTimeout creates a new GetWorkforcemanagementManagementunitSchedulingRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementManagementunitSchedulingRunParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitSchedulingRunParams {
	var ()
	return &GetWorkforcemanagementManagementunitSchedulingRunParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementManagementunitSchedulingRunParamsWithContext creates a new GetWorkforcemanagementManagementunitSchedulingRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementManagementunitSchedulingRunParamsWithContext(ctx context.Context) *GetWorkforcemanagementManagementunitSchedulingRunParams {
	var ()
	return &GetWorkforcemanagementManagementunitSchedulingRunParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementManagementunitSchedulingRunParamsWithHTTPClient creates a new GetWorkforcemanagementManagementunitSchedulingRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementManagementunitSchedulingRunParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitSchedulingRunParams {
	var ()
	return &GetWorkforcemanagementManagementunitSchedulingRunParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementManagementunitSchedulingRunParams contains all the parameters to send to the API endpoint
for the get workforcemanagement managementunit scheduling run operation typically these are written to a http.Request
*/
type GetWorkforcemanagementManagementunitSchedulingRunParams struct {

	/*ManagementUnitID
	  (Deprecated/gone): The ID of the management unit.

	*/
	ManagementUnitID string
	/*RunID
	  The ID of the schedule run

	*/
	RunID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement managementunit scheduling run params
func (o *GetWorkforcemanagementManagementunitSchedulingRunParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitSchedulingRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement managementunit scheduling run params
func (o *GetWorkforcemanagementManagementunitSchedulingRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement managementunit scheduling run params
func (o *GetWorkforcemanagementManagementunitSchedulingRunParams) WithContext(ctx context.Context) *GetWorkforcemanagementManagementunitSchedulingRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement managementunit scheduling run params
func (o *GetWorkforcemanagementManagementunitSchedulingRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement managementunit scheduling run params
func (o *GetWorkforcemanagementManagementunitSchedulingRunParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitSchedulingRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement managementunit scheduling run params
func (o *GetWorkforcemanagementManagementunitSchedulingRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManagementUnitID adds the managementUnitID to the get workforcemanagement managementunit scheduling run params
func (o *GetWorkforcemanagementManagementunitSchedulingRunParams) WithManagementUnitID(managementUnitID string) *GetWorkforcemanagementManagementunitSchedulingRunParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the get workforcemanagement managementunit scheduling run params
func (o *GetWorkforcemanagementManagementunitSchedulingRunParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WithRunID adds the runID to the get workforcemanagement managementunit scheduling run params
func (o *GetWorkforcemanagementManagementunitSchedulingRunParams) WithRunID(runID string) *GetWorkforcemanagementManagementunitSchedulingRunParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the get workforcemanagement managementunit scheduling run params
func (o *GetWorkforcemanagementManagementunitSchedulingRunParams) SetRunID(runID string) {
	o.RunID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementManagementunitSchedulingRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param managementUnitId
	if err := r.SetPathParam("managementUnitId", o.ManagementUnitID); err != nil {
		return err
	}

	// path param runId
	if err := r.SetPathParam("runId", o.RunID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
