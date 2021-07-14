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

// NewGetWorkforcemanagementAgentManagementunitParams creates a new GetWorkforcemanagementAgentManagementunitParams object
// with the default values initialized.
func NewGetWorkforcemanagementAgentManagementunitParams() *GetWorkforcemanagementAgentManagementunitParams {
	var ()
	return &GetWorkforcemanagementAgentManagementunitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementAgentManagementunitParamsWithTimeout creates a new GetWorkforcemanagementAgentManagementunitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementAgentManagementunitParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementAgentManagementunitParams {
	var ()
	return &GetWorkforcemanagementAgentManagementunitParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementAgentManagementunitParamsWithContext creates a new GetWorkforcemanagementAgentManagementunitParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementAgentManagementunitParamsWithContext(ctx context.Context) *GetWorkforcemanagementAgentManagementunitParams {
	var ()
	return &GetWorkforcemanagementAgentManagementunitParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementAgentManagementunitParamsWithHTTPClient creates a new GetWorkforcemanagementAgentManagementunitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementAgentManagementunitParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementAgentManagementunitParams {
	var ()
	return &GetWorkforcemanagementAgentManagementunitParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementAgentManagementunitParams contains all the parameters to send to the API endpoint
for the get workforcemanagement agent managementunit operation typically these are written to a http.Request
*/
type GetWorkforcemanagementAgentManagementunitParams struct {

	/*AgentID
	  The ID of the agent to look up

	*/
	AgentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement agent managementunit params
func (o *GetWorkforcemanagementAgentManagementunitParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementAgentManagementunitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement agent managementunit params
func (o *GetWorkforcemanagementAgentManagementunitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement agent managementunit params
func (o *GetWorkforcemanagementAgentManagementunitParams) WithContext(ctx context.Context) *GetWorkforcemanagementAgentManagementunitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement agent managementunit params
func (o *GetWorkforcemanagementAgentManagementunitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement agent managementunit params
func (o *GetWorkforcemanagementAgentManagementunitParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementAgentManagementunitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement agent managementunit params
func (o *GetWorkforcemanagementAgentManagementunitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the get workforcemanagement agent managementunit params
func (o *GetWorkforcemanagementAgentManagementunitParams) WithAgentID(agentID string) *GetWorkforcemanagementAgentManagementunitParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the get workforcemanagement agent managementunit params
func (o *GetWorkforcemanagementAgentManagementunitParams) SetAgentID(agentID string) {
	o.AgentID = agentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementAgentManagementunitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentId
	if err := r.SetPathParam("agentId", o.AgentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}