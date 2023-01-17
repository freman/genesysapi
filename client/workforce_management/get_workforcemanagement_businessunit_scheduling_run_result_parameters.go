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
	"github.com/go-openapi/swag"
)

// NewGetWorkforcemanagementBusinessunitSchedulingRunResultParams creates a new GetWorkforcemanagementBusinessunitSchedulingRunResultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkforcemanagementBusinessunitSchedulingRunResultParams() *GetWorkforcemanagementBusinessunitSchedulingRunResultParams {
	return &GetWorkforcemanagementBusinessunitSchedulingRunResultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunResultParamsWithTimeout creates a new GetWorkforcemanagementBusinessunitSchedulingRunResultParams object
// with the ability to set a timeout on a request.
func NewGetWorkforcemanagementBusinessunitSchedulingRunResultParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitSchedulingRunResultParams {
	return &GetWorkforcemanagementBusinessunitSchedulingRunResultParams{
		timeout: timeout,
	}
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunResultParamsWithContext creates a new GetWorkforcemanagementBusinessunitSchedulingRunResultParams object
// with the ability to set a context for a request.
func NewGetWorkforcemanagementBusinessunitSchedulingRunResultParamsWithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitSchedulingRunResultParams {
	return &GetWorkforcemanagementBusinessunitSchedulingRunResultParams{
		Context: ctx,
	}
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunResultParamsWithHTTPClient creates a new GetWorkforcemanagementBusinessunitSchedulingRunResultParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkforcemanagementBusinessunitSchedulingRunResultParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitSchedulingRunResultParams {
	return &GetWorkforcemanagementBusinessunitSchedulingRunResultParams{
		HTTPClient: client,
	}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunResultParams contains all the parameters to send to the API endpoint

	for the get workforcemanagement businessunit scheduling run result operation.

	Typically these are written to a http.Request.
*/
type GetWorkforcemanagementBusinessunitSchedulingRunResultParams struct {

	/* BusinessUnitID.

	   The ID of the business unit
	*/
	BusinessUnitID string

	/* Expand.

	   The fields to expand. Omitting will return an empty response
	*/
	Expand []string

	/* ManagementUnitIds.

	   The IDs of the management units for which to fetch the reschedule results
	*/
	ManagementUnitIds []string

	/* RunID.

	   The ID of the schedule run
	*/
	RunID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workforcemanagement businessunit scheduling run result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) WithDefaults() *GetWorkforcemanagementBusinessunitSchedulingRunResultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workforcemanagement businessunit scheduling run result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitSchedulingRunResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) WithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitSchedulingRunResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitSchedulingRunResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBusinessUnitID adds the businessUnitID to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) WithBusinessUnitID(businessUnitID string) *GetWorkforcemanagementBusinessunitSchedulingRunResultParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WithExpand adds the expand to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) WithExpand(expand []string) *GetWorkforcemanagementBusinessunitSchedulingRunResultParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithManagementUnitIds adds the managementUnitIds to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) WithManagementUnitIds(managementUnitIds []string) *GetWorkforcemanagementBusinessunitSchedulingRunResultParams {
	o.SetManagementUnitIds(managementUnitIds)
	return o
}

// SetManagementUnitIds adds the managementUnitIds to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) SetManagementUnitIds(managementUnitIds []string) {
	o.ManagementUnitIds = managementUnitIds
}

// WithRunID adds the runID to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) WithRunID(runID string) *GetWorkforcemanagementBusinessunitSchedulingRunResultParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the get workforcemanagement businessunit scheduling run result params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) SetRunID(runID string) {
	o.RunID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param businessUnitId
	if err := r.SetPathParam("businessUnitId", o.BusinessUnitID); err != nil {
		return err
	}

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	if o.ManagementUnitIds != nil {

		// binding items for managementUnitIds
		joinedManagementUnitIds := o.bindParamManagementUnitIds(reg)

		// query array param managementUnitIds
		if err := r.SetQueryParam("managementUnitIds", joinedManagementUnitIds...); err != nil {
			return err
		}
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

// bindParamGetWorkforcemanagementBusinessunitSchedulingRunResult binds the parameter expand
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}

// bindParamGetWorkforcemanagementBusinessunitSchedulingRunResult binds the parameter managementUnitIds
func (o *GetWorkforcemanagementBusinessunitSchedulingRunResultParams) bindParamManagementUnitIds(formats strfmt.Registry) []string {
	managementUnitIdsIR := o.ManagementUnitIds

	var managementUnitIdsIC []string
	for _, managementUnitIdsIIR := range managementUnitIdsIR { // explode []string

		managementUnitIdsIIV := managementUnitIdsIIR // string as string
		managementUnitIdsIC = append(managementUnitIdsIC, managementUnitIdsIIV)
	}

	// items.CollectionFormat: "multi"
	managementUnitIdsIS := swag.JoinByFormat(managementUnitIdsIC, "multi")

	return managementUnitIdsIS
}
