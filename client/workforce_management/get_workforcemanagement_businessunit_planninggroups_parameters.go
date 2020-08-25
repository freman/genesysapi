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

// NewGetWorkforcemanagementBusinessunitPlanninggroupsParams creates a new GetWorkforcemanagementBusinessunitPlanninggroupsParams object
// with the default values initialized.
func NewGetWorkforcemanagementBusinessunitPlanninggroupsParams() *GetWorkforcemanagementBusinessunitPlanninggroupsParams {
	var ()
	return &GetWorkforcemanagementBusinessunitPlanninggroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsParamsWithTimeout creates a new GetWorkforcemanagementBusinessunitPlanninggroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementBusinessunitPlanninggroupsParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitPlanninggroupsParams {
	var ()
	return &GetWorkforcemanagementBusinessunitPlanninggroupsParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsParamsWithContext creates a new GetWorkforcemanagementBusinessunitPlanninggroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementBusinessunitPlanninggroupsParamsWithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitPlanninggroupsParams {
	var ()
	return &GetWorkforcemanagementBusinessunitPlanninggroupsParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsParamsWithHTTPClient creates a new GetWorkforcemanagementBusinessunitPlanninggroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementBusinessunitPlanninggroupsParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitPlanninggroupsParams {
	var ()
	return &GetWorkforcemanagementBusinessunitPlanninggroupsParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsParams contains all the parameters to send to the API endpoint
for the get workforcemanagement businessunit planninggroups operation typically these are written to a http.Request
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsParams struct {

	/*BusinessUnitID
	  The ID of the business unit.

	*/
	BusinessUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement businessunit planninggroups params
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementBusinessunitPlanninggroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement businessunit planninggroups params
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement businessunit planninggroups params
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsParams) WithContext(ctx context.Context) *GetWorkforcemanagementBusinessunitPlanninggroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement businessunit planninggroups params
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement businessunit planninggroups params
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementBusinessunitPlanninggroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement businessunit planninggroups params
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBusinessUnitID adds the businessUnitID to the get workforcemanagement businessunit planninggroups params
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsParams) WithBusinessUnitID(businessUnitID string) *GetWorkforcemanagementBusinessunitPlanninggroupsParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the get workforcemanagement businessunit planninggroups params
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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