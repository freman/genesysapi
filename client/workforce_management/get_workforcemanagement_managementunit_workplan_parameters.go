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

// NewGetWorkforcemanagementManagementunitWorkplanParams creates a new GetWorkforcemanagementManagementunitWorkplanParams object
// with the default values initialized.
func NewGetWorkforcemanagementManagementunitWorkplanParams() *GetWorkforcemanagementManagementunitWorkplanParams {
	var ()
	return &GetWorkforcemanagementManagementunitWorkplanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementManagementunitWorkplanParamsWithTimeout creates a new GetWorkforcemanagementManagementunitWorkplanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementManagementunitWorkplanParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitWorkplanParams {
	var ()
	return &GetWorkforcemanagementManagementunitWorkplanParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementManagementunitWorkplanParamsWithContext creates a new GetWorkforcemanagementManagementunitWorkplanParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementManagementunitWorkplanParamsWithContext(ctx context.Context) *GetWorkforcemanagementManagementunitWorkplanParams {
	var ()
	return &GetWorkforcemanagementManagementunitWorkplanParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementManagementunitWorkplanParamsWithHTTPClient creates a new GetWorkforcemanagementManagementunitWorkplanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementManagementunitWorkplanParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitWorkplanParams {
	var ()
	return &GetWorkforcemanagementManagementunitWorkplanParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementManagementunitWorkplanParams contains all the parameters to send to the API endpoint
for the get workforcemanagement managementunit workplan operation typically these are written to a http.Request
*/
type GetWorkforcemanagementManagementunitWorkplanParams struct {

	/*ManagementUnitID
	  The ID of the management unit, or 'mine' for the management unit of the logged-in user.

	*/
	ManagementUnitID string
	/*WorkPlanID
	  The ID of the work plan to fetch

	*/
	WorkPlanID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement managementunit workplan params
func (o *GetWorkforcemanagementManagementunitWorkplanParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitWorkplanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement managementunit workplan params
func (o *GetWorkforcemanagementManagementunitWorkplanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement managementunit workplan params
func (o *GetWorkforcemanagementManagementunitWorkplanParams) WithContext(ctx context.Context) *GetWorkforcemanagementManagementunitWorkplanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement managementunit workplan params
func (o *GetWorkforcemanagementManagementunitWorkplanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement managementunit workplan params
func (o *GetWorkforcemanagementManagementunitWorkplanParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitWorkplanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement managementunit workplan params
func (o *GetWorkforcemanagementManagementunitWorkplanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManagementUnitID adds the managementUnitID to the get workforcemanagement managementunit workplan params
func (o *GetWorkforcemanagementManagementunitWorkplanParams) WithManagementUnitID(managementUnitID string) *GetWorkforcemanagementManagementunitWorkplanParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the get workforcemanagement managementunit workplan params
func (o *GetWorkforcemanagementManagementunitWorkplanParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WithWorkPlanID adds the workPlanID to the get workforcemanagement managementunit workplan params
func (o *GetWorkforcemanagementManagementunitWorkplanParams) WithWorkPlanID(workPlanID string) *GetWorkforcemanagementManagementunitWorkplanParams {
	o.SetWorkPlanID(workPlanID)
	return o
}

// SetWorkPlanID adds the workPlanId to the get workforcemanagement managementunit workplan params
func (o *GetWorkforcemanagementManagementunitWorkplanParams) SetWorkPlanID(workPlanID string) {
	o.WorkPlanID = workPlanID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementManagementunitWorkplanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param managementUnitId
	if err := r.SetPathParam("managementUnitId", o.ManagementUnitID); err != nil {
		return err
	}

	// path param workPlanId
	if err := r.SetPathParam("workPlanId", o.WorkPlanID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
