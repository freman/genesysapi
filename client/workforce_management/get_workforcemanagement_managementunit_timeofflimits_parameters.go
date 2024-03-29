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

// NewGetWorkforcemanagementManagementunitTimeofflimitsParams creates a new GetWorkforcemanagementManagementunitTimeofflimitsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkforcemanagementManagementunitTimeofflimitsParams() *GetWorkforcemanagementManagementunitTimeofflimitsParams {
	return &GetWorkforcemanagementManagementunitTimeofflimitsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsParamsWithTimeout creates a new GetWorkforcemanagementManagementunitTimeofflimitsParams object
// with the ability to set a timeout on a request.
func NewGetWorkforcemanagementManagementunitTimeofflimitsParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitTimeofflimitsParams {
	return &GetWorkforcemanagementManagementunitTimeofflimitsParams{
		timeout: timeout,
	}
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsParamsWithContext creates a new GetWorkforcemanagementManagementunitTimeofflimitsParams object
// with the ability to set a context for a request.
func NewGetWorkforcemanagementManagementunitTimeofflimitsParamsWithContext(ctx context.Context) *GetWorkforcemanagementManagementunitTimeofflimitsParams {
	return &GetWorkforcemanagementManagementunitTimeofflimitsParams{
		Context: ctx,
	}
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsParamsWithHTTPClient creates a new GetWorkforcemanagementManagementunitTimeofflimitsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkforcemanagementManagementunitTimeofflimitsParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitTimeofflimitsParams {
	return &GetWorkforcemanagementManagementunitTimeofflimitsParams{
		HTTPClient: client,
	}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsParams contains all the parameters to send to the API endpoint

	for the get workforcemanagement managementunit timeofflimits operation.

	Typically these are written to a http.Request.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsParams struct {

	/* ManagementUnitID.

	   The ID of the management unit.
	*/
	ManagementUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workforcemanagement managementunit timeofflimits params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementManagementunitTimeofflimitsParams) WithDefaults() *GetWorkforcemanagementManagementunitTimeofflimitsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workforcemanagement managementunit timeofflimits params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementManagementunitTimeofflimitsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workforcemanagement managementunit timeofflimits params
func (o *GetWorkforcemanagementManagementunitTimeofflimitsParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitTimeofflimitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement managementunit timeofflimits params
func (o *GetWorkforcemanagementManagementunitTimeofflimitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement managementunit timeofflimits params
func (o *GetWorkforcemanagementManagementunitTimeofflimitsParams) WithContext(ctx context.Context) *GetWorkforcemanagementManagementunitTimeofflimitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement managementunit timeofflimits params
func (o *GetWorkforcemanagementManagementunitTimeofflimitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement managementunit timeofflimits params
func (o *GetWorkforcemanagementManagementunitTimeofflimitsParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitTimeofflimitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement managementunit timeofflimits params
func (o *GetWorkforcemanagementManagementunitTimeofflimitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManagementUnitID adds the managementUnitID to the get workforcemanagement managementunit timeofflimits params
func (o *GetWorkforcemanagementManagementunitTimeofflimitsParams) WithManagementUnitID(managementUnitID string) *GetWorkforcemanagementManagementunitTimeofflimitsParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the get workforcemanagement managementunit timeofflimits params
func (o *GetWorkforcemanagementManagementunitTimeofflimitsParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementManagementunitTimeofflimitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param managementUnitId
	if err := r.SetPathParam("managementUnitId", o.ManagementUnitID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
