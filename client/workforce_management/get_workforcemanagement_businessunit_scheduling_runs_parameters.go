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

// NewGetWorkforcemanagementBusinessunitSchedulingRunsParams creates a new GetWorkforcemanagementBusinessunitSchedulingRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkforcemanagementBusinessunitSchedulingRunsParams() *GetWorkforcemanagementBusinessunitSchedulingRunsParams {
	return &GetWorkforcemanagementBusinessunitSchedulingRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunsParamsWithTimeout creates a new GetWorkforcemanagementBusinessunitSchedulingRunsParams object
// with the ability to set a timeout on a request.
func NewGetWorkforcemanagementBusinessunitSchedulingRunsParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitSchedulingRunsParams {
	return &GetWorkforcemanagementBusinessunitSchedulingRunsParams{
		timeout: timeout,
	}
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunsParamsWithContext creates a new GetWorkforcemanagementBusinessunitSchedulingRunsParams object
// with the ability to set a context for a request.
func NewGetWorkforcemanagementBusinessunitSchedulingRunsParamsWithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitSchedulingRunsParams {
	return &GetWorkforcemanagementBusinessunitSchedulingRunsParams{
		Context: ctx,
	}
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunsParamsWithHTTPClient creates a new GetWorkforcemanagementBusinessunitSchedulingRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkforcemanagementBusinessunitSchedulingRunsParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitSchedulingRunsParams {
	return &GetWorkforcemanagementBusinessunitSchedulingRunsParams{
		HTTPClient: client,
	}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunsParams contains all the parameters to send to the API endpoint

	for the get workforcemanagement businessunit scheduling runs operation.

	Typically these are written to a http.Request.
*/
type GetWorkforcemanagementBusinessunitSchedulingRunsParams struct {

	/* BusinessUnitID.

	   The ID of the business unit
	*/
	BusinessUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workforcemanagement businessunit scheduling runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementBusinessunitSchedulingRunsParams) WithDefaults() *GetWorkforcemanagementBusinessunitSchedulingRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workforcemanagement businessunit scheduling runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementBusinessunitSchedulingRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workforcemanagement businessunit scheduling runs params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunsParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitSchedulingRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement businessunit scheduling runs params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement businessunit scheduling runs params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunsParams) WithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitSchedulingRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement businessunit scheduling runs params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement businessunit scheduling runs params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunsParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitSchedulingRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement businessunit scheduling runs params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBusinessUnitID adds the businessUnitID to the get workforcemanagement businessunit scheduling runs params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunsParams) WithBusinessUnitID(businessUnitID string) *GetWorkforcemanagementBusinessunitSchedulingRunsParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the get workforcemanagement businessunit scheduling runs params
func (o *GetWorkforcemanagementBusinessunitSchedulingRunsParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementBusinessunitSchedulingRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param businessUnitId
	if err := r.SetPathParam("businessUnitId", o.BusinessUnitID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
