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

// NewGetWorkforcemanagementBusinessunitServicegoaltemplateParams creates a new GetWorkforcemanagementBusinessunitServicegoaltemplateParams object
// with the default values initialized.
func NewGetWorkforcemanagementBusinessunitServicegoaltemplateParams() *GetWorkforcemanagementBusinessunitServicegoaltemplateParams {
	var ()
	return &GetWorkforcemanagementBusinessunitServicegoaltemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementBusinessunitServicegoaltemplateParamsWithTimeout creates a new GetWorkforcemanagementBusinessunitServicegoaltemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementBusinessunitServicegoaltemplateParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitServicegoaltemplateParams {
	var ()
	return &GetWorkforcemanagementBusinessunitServicegoaltemplateParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementBusinessunitServicegoaltemplateParamsWithContext creates a new GetWorkforcemanagementBusinessunitServicegoaltemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementBusinessunitServicegoaltemplateParamsWithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitServicegoaltemplateParams {
	var ()
	return &GetWorkforcemanagementBusinessunitServicegoaltemplateParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementBusinessunitServicegoaltemplateParamsWithHTTPClient creates a new GetWorkforcemanagementBusinessunitServicegoaltemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementBusinessunitServicegoaltemplateParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitServicegoaltemplateParams {
	var ()
	return &GetWorkforcemanagementBusinessunitServicegoaltemplateParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementBusinessunitServicegoaltemplateParams contains all the parameters to send to the API endpoint
for the get workforcemanagement businessunit servicegoaltemplate operation typically these are written to a http.Request
*/
type GetWorkforcemanagementBusinessunitServicegoaltemplateParams struct {

	/*BusinessUnitID
	  The ID of the business unit.

	*/
	BusinessUnitID string
	/*ServiceGoalTemplateID
	  The ID of a service goal template to fetch

	*/
	ServiceGoalTemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement businessunit servicegoaltemplate params
func (o *GetWorkforcemanagementBusinessunitServicegoaltemplateParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitServicegoaltemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement businessunit servicegoaltemplate params
func (o *GetWorkforcemanagementBusinessunitServicegoaltemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement businessunit servicegoaltemplate params
func (o *GetWorkforcemanagementBusinessunitServicegoaltemplateParams) WithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitServicegoaltemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement businessunit servicegoaltemplate params
func (o *GetWorkforcemanagementBusinessunitServicegoaltemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement businessunit servicegoaltemplate params
func (o *GetWorkforcemanagementBusinessunitServicegoaltemplateParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitServicegoaltemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement businessunit servicegoaltemplate params
func (o *GetWorkforcemanagementBusinessunitServicegoaltemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBusinessUnitID adds the businessUnitID to the get workforcemanagement businessunit servicegoaltemplate params
func (o *GetWorkforcemanagementBusinessunitServicegoaltemplateParams) WithBusinessUnitID(businessUnitID string) *GetWorkforcemanagementBusinessunitServicegoaltemplateParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the get workforcemanagement businessunit servicegoaltemplate params
func (o *GetWorkforcemanagementBusinessunitServicegoaltemplateParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithServiceGoalTemplateID adds the serviceGoalTemplateID to the get workforcemanagement businessunit servicegoaltemplate params
func (o *GetWorkforcemanagementBusinessunitServicegoaltemplateParams) WithServiceGoalTemplateID(serviceGoalTemplateID string) *GetWorkforcemanagementBusinessunitServicegoaltemplateParams {
	o.SetServiceGoalTemplateID(serviceGoalTemplateID)
	return o
}

// SetServiceGoalTemplateID adds the serviceGoalTemplateId to the get workforcemanagement businessunit servicegoaltemplate params
func (o *GetWorkforcemanagementBusinessunitServicegoaltemplateParams) SetServiceGoalTemplateID(serviceGoalTemplateID string) {
	o.ServiceGoalTemplateID = serviceGoalTemplateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementBusinessunitServicegoaltemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param businessUnitId
	if err := r.SetPathParam("businessUnitId", o.BusinessUnitID); err != nil {
		return err
	}

	// path param serviceGoalTemplateId
	if err := r.SetPathParam("serviceGoalTemplateId", o.ServiceGoalTemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}