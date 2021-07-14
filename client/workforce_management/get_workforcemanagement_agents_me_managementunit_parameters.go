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

// NewGetWorkforcemanagementAgentsMeManagementunitParams creates a new GetWorkforcemanagementAgentsMeManagementunitParams object
// with the default values initialized.
func NewGetWorkforcemanagementAgentsMeManagementunitParams() *GetWorkforcemanagementAgentsMeManagementunitParams {

	return &GetWorkforcemanagementAgentsMeManagementunitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementAgentsMeManagementunitParamsWithTimeout creates a new GetWorkforcemanagementAgentsMeManagementunitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementAgentsMeManagementunitParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementAgentsMeManagementunitParams {

	return &GetWorkforcemanagementAgentsMeManagementunitParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementAgentsMeManagementunitParamsWithContext creates a new GetWorkforcemanagementAgentsMeManagementunitParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementAgentsMeManagementunitParamsWithContext(ctx context.Context) *GetWorkforcemanagementAgentsMeManagementunitParams {

	return &GetWorkforcemanagementAgentsMeManagementunitParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementAgentsMeManagementunitParamsWithHTTPClient creates a new GetWorkforcemanagementAgentsMeManagementunitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementAgentsMeManagementunitParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementAgentsMeManagementunitParams {

	return &GetWorkforcemanagementAgentsMeManagementunitParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementAgentsMeManagementunitParams contains all the parameters to send to the API endpoint
for the get workforcemanagement agents me managementunit operation typically these are written to a http.Request
*/
type GetWorkforcemanagementAgentsMeManagementunitParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement agents me managementunit params
func (o *GetWorkforcemanagementAgentsMeManagementunitParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementAgentsMeManagementunitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement agents me managementunit params
func (o *GetWorkforcemanagementAgentsMeManagementunitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement agents me managementunit params
func (o *GetWorkforcemanagementAgentsMeManagementunitParams) WithContext(ctx context.Context) *GetWorkforcemanagementAgentsMeManagementunitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement agents me managementunit params
func (o *GetWorkforcemanagementAgentsMeManagementunitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement agents me managementunit params
func (o *GetWorkforcemanagementAgentsMeManagementunitParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementAgentsMeManagementunitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement agents me managementunit params
func (o *GetWorkforcemanagementAgentsMeManagementunitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementAgentsMeManagementunitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}