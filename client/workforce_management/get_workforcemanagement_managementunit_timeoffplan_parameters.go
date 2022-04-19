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

// NewGetWorkforcemanagementManagementunitTimeoffplanParams creates a new GetWorkforcemanagementManagementunitTimeoffplanParams object
// with the default values initialized.
func NewGetWorkforcemanagementManagementunitTimeoffplanParams() *GetWorkforcemanagementManagementunitTimeoffplanParams {
	var ()
	return &GetWorkforcemanagementManagementunitTimeoffplanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementManagementunitTimeoffplanParamsWithTimeout creates a new GetWorkforcemanagementManagementunitTimeoffplanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementManagementunitTimeoffplanParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitTimeoffplanParams {
	var ()
	return &GetWorkforcemanagementManagementunitTimeoffplanParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementManagementunitTimeoffplanParamsWithContext creates a new GetWorkforcemanagementManagementunitTimeoffplanParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementManagementunitTimeoffplanParamsWithContext(ctx context.Context) *GetWorkforcemanagementManagementunitTimeoffplanParams {
	var ()
	return &GetWorkforcemanagementManagementunitTimeoffplanParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementManagementunitTimeoffplanParamsWithHTTPClient creates a new GetWorkforcemanagementManagementunitTimeoffplanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementManagementunitTimeoffplanParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitTimeoffplanParams {
	var ()
	return &GetWorkforcemanagementManagementunitTimeoffplanParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementManagementunitTimeoffplanParams contains all the parameters to send to the API endpoint
for the get workforcemanagement managementunit timeoffplan operation typically these are written to a http.Request
*/
type GetWorkforcemanagementManagementunitTimeoffplanParams struct {

	/*ManagementUnitID
	  The ID of the management unit

	*/
	ManagementUnitID string
	/*TimeOffPlanID
	  The ID of the time off plan to fetch

	*/
	TimeOffPlanID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement managementunit timeoffplan params
func (o *GetWorkforcemanagementManagementunitTimeoffplanParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitTimeoffplanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement managementunit timeoffplan params
func (o *GetWorkforcemanagementManagementunitTimeoffplanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement managementunit timeoffplan params
func (o *GetWorkforcemanagementManagementunitTimeoffplanParams) WithContext(ctx context.Context) *GetWorkforcemanagementManagementunitTimeoffplanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement managementunit timeoffplan params
func (o *GetWorkforcemanagementManagementunitTimeoffplanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement managementunit timeoffplan params
func (o *GetWorkforcemanagementManagementunitTimeoffplanParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitTimeoffplanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement managementunit timeoffplan params
func (o *GetWorkforcemanagementManagementunitTimeoffplanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManagementUnitID adds the managementUnitID to the get workforcemanagement managementunit timeoffplan params
func (o *GetWorkforcemanagementManagementunitTimeoffplanParams) WithManagementUnitID(managementUnitID string) *GetWorkforcemanagementManagementunitTimeoffplanParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the get workforcemanagement managementunit timeoffplan params
func (o *GetWorkforcemanagementManagementunitTimeoffplanParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WithTimeOffPlanID adds the timeOffPlanID to the get workforcemanagement managementunit timeoffplan params
func (o *GetWorkforcemanagementManagementunitTimeoffplanParams) WithTimeOffPlanID(timeOffPlanID string) *GetWorkforcemanagementManagementunitTimeoffplanParams {
	o.SetTimeOffPlanID(timeOffPlanID)
	return o
}

// SetTimeOffPlanID adds the timeOffPlanId to the get workforcemanagement managementunit timeoffplan params
func (o *GetWorkforcemanagementManagementunitTimeoffplanParams) SetTimeOffPlanID(timeOffPlanID string) {
	o.TimeOffPlanID = timeOffPlanID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementManagementunitTimeoffplanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param managementUnitId
	if err := r.SetPathParam("managementUnitId", o.ManagementUnitID); err != nil {
		return err
	}

	// path param timeOffPlanId
	if err := r.SetPathParam("timeOffPlanId", o.TimeOffPlanID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
