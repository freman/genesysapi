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

// NewDeleteWorkforcemanagementManagementunitTimeofflimitParams creates a new DeleteWorkforcemanagementManagementunitTimeofflimitParams object
// with the default values initialized.
func NewDeleteWorkforcemanagementManagementunitTimeofflimitParams() *DeleteWorkforcemanagementManagementunitTimeofflimitParams {
	var ()
	return &DeleteWorkforcemanagementManagementunitTimeofflimitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkforcemanagementManagementunitTimeofflimitParamsWithTimeout creates a new DeleteWorkforcemanagementManagementunitTimeofflimitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteWorkforcemanagementManagementunitTimeofflimitParamsWithTimeout(timeout time.Duration) *DeleteWorkforcemanagementManagementunitTimeofflimitParams {
	var ()
	return &DeleteWorkforcemanagementManagementunitTimeofflimitParams{

		timeout: timeout,
	}
}

// NewDeleteWorkforcemanagementManagementunitTimeofflimitParamsWithContext creates a new DeleteWorkforcemanagementManagementunitTimeofflimitParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteWorkforcemanagementManagementunitTimeofflimitParamsWithContext(ctx context.Context) *DeleteWorkforcemanagementManagementunitTimeofflimitParams {
	var ()
	return &DeleteWorkforcemanagementManagementunitTimeofflimitParams{

		Context: ctx,
	}
}

// NewDeleteWorkforcemanagementManagementunitTimeofflimitParamsWithHTTPClient creates a new DeleteWorkforcemanagementManagementunitTimeofflimitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteWorkforcemanagementManagementunitTimeofflimitParamsWithHTTPClient(client *http.Client) *DeleteWorkforcemanagementManagementunitTimeofflimitParams {
	var ()
	return &DeleteWorkforcemanagementManagementunitTimeofflimitParams{
		HTTPClient: client,
	}
}

/*DeleteWorkforcemanagementManagementunitTimeofflimitParams contains all the parameters to send to the API endpoint
for the delete workforcemanagement managementunit timeofflimit operation typically these are written to a http.Request
*/
type DeleteWorkforcemanagementManagementunitTimeofflimitParams struct {

	/*ManagementUnitID
	  The management unit ID of the management unit.

	*/
	ManagementUnitID string
	/*TimeOffLimitID
	  The ID of the time off limit object to delete

	*/
	TimeOffLimitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete workforcemanagement managementunit timeofflimit params
func (o *DeleteWorkforcemanagementManagementunitTimeofflimitParams) WithTimeout(timeout time.Duration) *DeleteWorkforcemanagementManagementunitTimeofflimitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete workforcemanagement managementunit timeofflimit params
func (o *DeleteWorkforcemanagementManagementunitTimeofflimitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete workforcemanagement managementunit timeofflimit params
func (o *DeleteWorkforcemanagementManagementunitTimeofflimitParams) WithContext(ctx context.Context) *DeleteWorkforcemanagementManagementunitTimeofflimitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete workforcemanagement managementunit timeofflimit params
func (o *DeleteWorkforcemanagementManagementunitTimeofflimitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete workforcemanagement managementunit timeofflimit params
func (o *DeleteWorkforcemanagementManagementunitTimeofflimitParams) WithHTTPClient(client *http.Client) *DeleteWorkforcemanagementManagementunitTimeofflimitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete workforcemanagement managementunit timeofflimit params
func (o *DeleteWorkforcemanagementManagementunitTimeofflimitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManagementUnitID adds the managementUnitID to the delete workforcemanagement managementunit timeofflimit params
func (o *DeleteWorkforcemanagementManagementunitTimeofflimitParams) WithManagementUnitID(managementUnitID string) *DeleteWorkforcemanagementManagementunitTimeofflimitParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the delete workforcemanagement managementunit timeofflimit params
func (o *DeleteWorkforcemanagementManagementunitTimeofflimitParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WithTimeOffLimitID adds the timeOffLimitID to the delete workforcemanagement managementunit timeofflimit params
func (o *DeleteWorkforcemanagementManagementunitTimeofflimitParams) WithTimeOffLimitID(timeOffLimitID string) *DeleteWorkforcemanagementManagementunitTimeofflimitParams {
	o.SetTimeOffLimitID(timeOffLimitID)
	return o
}

// SetTimeOffLimitID adds the timeOffLimitId to the delete workforcemanagement managementunit timeofflimit params
func (o *DeleteWorkforcemanagementManagementunitTimeofflimitParams) SetTimeOffLimitID(timeOffLimitID string) {
	o.TimeOffLimitID = timeOffLimitID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkforcemanagementManagementunitTimeofflimitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param managementUnitId
	if err := r.SetPathParam("managementUnitId", o.ManagementUnitID); err != nil {
		return err
	}

	// path param timeOffLimitId
	if err := r.SetPathParam("timeOffLimitId", o.TimeOffLimitID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
