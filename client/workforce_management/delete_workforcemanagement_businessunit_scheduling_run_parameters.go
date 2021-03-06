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

// NewDeleteWorkforcemanagementBusinessunitSchedulingRunParams creates a new DeleteWorkforcemanagementBusinessunitSchedulingRunParams object
// with the default values initialized.
func NewDeleteWorkforcemanagementBusinessunitSchedulingRunParams() *DeleteWorkforcemanagementBusinessunitSchedulingRunParams {
	var ()
	return &DeleteWorkforcemanagementBusinessunitSchedulingRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkforcemanagementBusinessunitSchedulingRunParamsWithTimeout creates a new DeleteWorkforcemanagementBusinessunitSchedulingRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteWorkforcemanagementBusinessunitSchedulingRunParamsWithTimeout(timeout time.Duration) *DeleteWorkforcemanagementBusinessunitSchedulingRunParams {
	var ()
	return &DeleteWorkforcemanagementBusinessunitSchedulingRunParams{

		timeout: timeout,
	}
}

// NewDeleteWorkforcemanagementBusinessunitSchedulingRunParamsWithContext creates a new DeleteWorkforcemanagementBusinessunitSchedulingRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteWorkforcemanagementBusinessunitSchedulingRunParamsWithContext(ctx context.Context) *DeleteWorkforcemanagementBusinessunitSchedulingRunParams {
	var ()
	return &DeleteWorkforcemanagementBusinessunitSchedulingRunParams{

		Context: ctx,
	}
}

// NewDeleteWorkforcemanagementBusinessunitSchedulingRunParamsWithHTTPClient creates a new DeleteWorkforcemanagementBusinessunitSchedulingRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteWorkforcemanagementBusinessunitSchedulingRunParamsWithHTTPClient(client *http.Client) *DeleteWorkforcemanagementBusinessunitSchedulingRunParams {
	var ()
	return &DeleteWorkforcemanagementBusinessunitSchedulingRunParams{
		HTTPClient: client,
	}
}

/*DeleteWorkforcemanagementBusinessunitSchedulingRunParams contains all the parameters to send to the API endpoint
for the delete workforcemanagement businessunit scheduling run operation typically these are written to a http.Request
*/
type DeleteWorkforcemanagementBusinessunitSchedulingRunParams struct {

	/*BusinessUnitID
	  The ID of the business unit

	*/
	BusinessUnitID string
	/*RunID
	  The ID of the schedule run

	*/
	RunID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete workforcemanagement businessunit scheduling run params
func (o *DeleteWorkforcemanagementBusinessunitSchedulingRunParams) WithTimeout(timeout time.Duration) *DeleteWorkforcemanagementBusinessunitSchedulingRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete workforcemanagement businessunit scheduling run params
func (o *DeleteWorkforcemanagementBusinessunitSchedulingRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete workforcemanagement businessunit scheduling run params
func (o *DeleteWorkforcemanagementBusinessunitSchedulingRunParams) WithContext(ctx context.Context) *DeleteWorkforcemanagementBusinessunitSchedulingRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete workforcemanagement businessunit scheduling run params
func (o *DeleteWorkforcemanagementBusinessunitSchedulingRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete workforcemanagement businessunit scheduling run params
func (o *DeleteWorkforcemanagementBusinessunitSchedulingRunParams) WithHTTPClient(client *http.Client) *DeleteWorkforcemanagementBusinessunitSchedulingRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete workforcemanagement businessunit scheduling run params
func (o *DeleteWorkforcemanagementBusinessunitSchedulingRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBusinessUnitID adds the businessUnitID to the delete workforcemanagement businessunit scheduling run params
func (o *DeleteWorkforcemanagementBusinessunitSchedulingRunParams) WithBusinessUnitID(businessUnitID string) *DeleteWorkforcemanagementBusinessunitSchedulingRunParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the delete workforcemanagement businessunit scheduling run params
func (o *DeleteWorkforcemanagementBusinessunitSchedulingRunParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithRunID adds the runID to the delete workforcemanagement businessunit scheduling run params
func (o *DeleteWorkforcemanagementBusinessunitSchedulingRunParams) WithRunID(runID string) *DeleteWorkforcemanagementBusinessunitSchedulingRunParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the delete workforcemanagement businessunit scheduling run params
func (o *DeleteWorkforcemanagementBusinessunitSchedulingRunParams) SetRunID(runID string) {
	o.RunID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkforcemanagementBusinessunitSchedulingRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param businessUnitId
	if err := r.SetPathParam("businessUnitId", o.BusinessUnitID); err != nil {
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
