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

// NewGetWorkforcemanagementManagementunitAgentShifttradesParams creates a new GetWorkforcemanagementManagementunitAgentShifttradesParams object
// with the default values initialized.
func NewGetWorkforcemanagementManagementunitAgentShifttradesParams() *GetWorkforcemanagementManagementunitAgentShifttradesParams {
	var ()
	return &GetWorkforcemanagementManagementunitAgentShifttradesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesParamsWithTimeout creates a new GetWorkforcemanagementManagementunitAgentShifttradesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementManagementunitAgentShifttradesParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitAgentShifttradesParams {
	var ()
	return &GetWorkforcemanagementManagementunitAgentShifttradesParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesParamsWithContext creates a new GetWorkforcemanagementManagementunitAgentShifttradesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementManagementunitAgentShifttradesParamsWithContext(ctx context.Context) *GetWorkforcemanagementManagementunitAgentShifttradesParams {
	var ()
	return &GetWorkforcemanagementManagementunitAgentShifttradesParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesParamsWithHTTPClient creates a new GetWorkforcemanagementManagementunitAgentShifttradesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementManagementunitAgentShifttradesParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitAgentShifttradesParams {
	var ()
	return &GetWorkforcemanagementManagementunitAgentShifttradesParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementManagementunitAgentShifttradesParams contains all the parameters to send to the API endpoint
for the get workforcemanagement managementunit agent shifttrades operation typically these are written to a http.Request
*/
type GetWorkforcemanagementManagementunitAgentShifttradesParams struct {

	/*AgentID
	  The agent id

	*/
	AgentID string
	/*ManagementUnitID
	  The id of the management unit, or 'mine' for the management unit of the logged-in user.

	*/
	ManagementUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement managementunit agent shifttrades params
func (o *GetWorkforcemanagementManagementunitAgentShifttradesParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitAgentShifttradesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement managementunit agent shifttrades params
func (o *GetWorkforcemanagementManagementunitAgentShifttradesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement managementunit agent shifttrades params
func (o *GetWorkforcemanagementManagementunitAgentShifttradesParams) WithContext(ctx context.Context) *GetWorkforcemanagementManagementunitAgentShifttradesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement managementunit agent shifttrades params
func (o *GetWorkforcemanagementManagementunitAgentShifttradesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement managementunit agent shifttrades params
func (o *GetWorkforcemanagementManagementunitAgentShifttradesParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitAgentShifttradesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement managementunit agent shifttrades params
func (o *GetWorkforcemanagementManagementunitAgentShifttradesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the get workforcemanagement managementunit agent shifttrades params
func (o *GetWorkforcemanagementManagementunitAgentShifttradesParams) WithAgentID(agentID string) *GetWorkforcemanagementManagementunitAgentShifttradesParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the get workforcemanagement managementunit agent shifttrades params
func (o *GetWorkforcemanagementManagementunitAgentShifttradesParams) SetAgentID(agentID string) {
	o.AgentID = agentID
}

// WithManagementUnitID adds the managementUnitID to the get workforcemanagement managementunit agent shifttrades params
func (o *GetWorkforcemanagementManagementunitAgentShifttradesParams) WithManagementUnitID(managementUnitID string) *GetWorkforcemanagementManagementunitAgentShifttradesParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the get workforcemanagement managementunit agent shifttrades params
func (o *GetWorkforcemanagementManagementunitAgentShifttradesParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementManagementunitAgentShifttradesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentId
	if err := r.SetPathParam("agentId", o.AgentID); err != nil {
		return err
	}

	// path param managementUnitId
	if err := r.SetPathParam("managementUnitId", o.ManagementUnitID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
