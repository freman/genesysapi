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

// NewGetWorkforcemanagementManagementunitTimeofflimitParams creates a new GetWorkforcemanagementManagementunitTimeofflimitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkforcemanagementManagementunitTimeofflimitParams() *GetWorkforcemanagementManagementunitTimeofflimitParams {
	return &GetWorkforcemanagementManagementunitTimeofflimitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementManagementunitTimeofflimitParamsWithTimeout creates a new GetWorkforcemanagementManagementunitTimeofflimitParams object
// with the ability to set a timeout on a request.
func NewGetWorkforcemanagementManagementunitTimeofflimitParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitTimeofflimitParams {
	return &GetWorkforcemanagementManagementunitTimeofflimitParams{
		timeout: timeout,
	}
}

// NewGetWorkforcemanagementManagementunitTimeofflimitParamsWithContext creates a new GetWorkforcemanagementManagementunitTimeofflimitParams object
// with the ability to set a context for a request.
func NewGetWorkforcemanagementManagementunitTimeofflimitParamsWithContext(ctx context.Context) *GetWorkforcemanagementManagementunitTimeofflimitParams {
	return &GetWorkforcemanagementManagementunitTimeofflimitParams{
		Context: ctx,
	}
}

// NewGetWorkforcemanagementManagementunitTimeofflimitParamsWithHTTPClient creates a new GetWorkforcemanagementManagementunitTimeofflimitParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkforcemanagementManagementunitTimeofflimitParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitTimeofflimitParams {
	return &GetWorkforcemanagementManagementunitTimeofflimitParams{
		HTTPClient: client,
	}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitParams contains all the parameters to send to the API endpoint

	for the get workforcemanagement managementunit timeofflimit operation.

	Typically these are written to a http.Request.
*/
type GetWorkforcemanagementManagementunitTimeofflimitParams struct {

	/* ManagementUnitID.

	   The ID of the management unit.
	*/
	ManagementUnitID string

	/* TimeOffLimitID.

	   The ID of the time off limit to fetch
	*/
	TimeOffLimitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workforcemanagement managementunit timeofflimit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) WithDefaults() *GetWorkforcemanagementManagementunitTimeofflimitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workforcemanagement managementunit timeofflimit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workforcemanagement managementunit timeofflimit params
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitTimeofflimitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement managementunit timeofflimit params
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement managementunit timeofflimit params
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) WithContext(ctx context.Context) *GetWorkforcemanagementManagementunitTimeofflimitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement managementunit timeofflimit params
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement managementunit timeofflimit params
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitTimeofflimitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement managementunit timeofflimit params
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManagementUnitID adds the managementUnitID to the get workforcemanagement managementunit timeofflimit params
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) WithManagementUnitID(managementUnitID string) *GetWorkforcemanagementManagementunitTimeofflimitParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the get workforcemanagement managementunit timeofflimit params
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WithTimeOffLimitID adds the timeOffLimitID to the get workforcemanagement managementunit timeofflimit params
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) WithTimeOffLimitID(timeOffLimitID string) *GetWorkforcemanagementManagementunitTimeofflimitParams {
	o.SetTimeOffLimitID(timeOffLimitID)
	return o
}

// SetTimeOffLimitID adds the timeOffLimitId to the get workforcemanagement managementunit timeofflimit params
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) SetTimeOffLimitID(timeOffLimitID string) {
	o.TimeOffLimitID = timeOffLimitID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementManagementunitTimeofflimitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
