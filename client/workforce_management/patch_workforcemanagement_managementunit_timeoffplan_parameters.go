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

	"github.com/freman/genesysapi/models"
)

// NewPatchWorkforcemanagementManagementunitTimeoffplanParams creates a new PatchWorkforcemanagementManagementunitTimeoffplanParams object
// with the default values initialized.
func NewPatchWorkforcemanagementManagementunitTimeoffplanParams() *PatchWorkforcemanagementManagementunitTimeoffplanParams {
	var ()
	return &PatchWorkforcemanagementManagementunitTimeoffplanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWorkforcemanagementManagementunitTimeoffplanParamsWithTimeout creates a new PatchWorkforcemanagementManagementunitTimeoffplanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchWorkforcemanagementManagementunitTimeoffplanParamsWithTimeout(timeout time.Duration) *PatchWorkforcemanagementManagementunitTimeoffplanParams {
	var ()
	return &PatchWorkforcemanagementManagementunitTimeoffplanParams{

		timeout: timeout,
	}
}

// NewPatchWorkforcemanagementManagementunitTimeoffplanParamsWithContext creates a new PatchWorkforcemanagementManagementunitTimeoffplanParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchWorkforcemanagementManagementunitTimeoffplanParamsWithContext(ctx context.Context) *PatchWorkforcemanagementManagementunitTimeoffplanParams {
	var ()
	return &PatchWorkforcemanagementManagementunitTimeoffplanParams{

		Context: ctx,
	}
}

// NewPatchWorkforcemanagementManagementunitTimeoffplanParamsWithHTTPClient creates a new PatchWorkforcemanagementManagementunitTimeoffplanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchWorkforcemanagementManagementunitTimeoffplanParamsWithHTTPClient(client *http.Client) *PatchWorkforcemanagementManagementunitTimeoffplanParams {
	var ()
	return &PatchWorkforcemanagementManagementunitTimeoffplanParams{
		HTTPClient: client,
	}
}

/*PatchWorkforcemanagementManagementunitTimeoffplanParams contains all the parameters to send to the API endpoint
for the patch workforcemanagement managementunit timeoffplan operation typically these are written to a http.Request
*/
type PatchWorkforcemanagementManagementunitTimeoffplanParams struct {

	/*Body
	  body

	*/
	Body *models.UpdateTimeOffPlanRequest
	/*ManagementUnitID
	  The management unit ID.

	*/
	ManagementUnitID string
	/*TimeOffPlanID
	  The ID of the time off plan to update

	*/
	TimeOffPlanID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch workforcemanagement managementunit timeoffplan params
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) WithTimeout(timeout time.Duration) *PatchWorkforcemanagementManagementunitTimeoffplanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch workforcemanagement managementunit timeoffplan params
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch workforcemanagement managementunit timeoffplan params
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) WithContext(ctx context.Context) *PatchWorkforcemanagementManagementunitTimeoffplanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch workforcemanagement managementunit timeoffplan params
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch workforcemanagement managementunit timeoffplan params
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) WithHTTPClient(client *http.Client) *PatchWorkforcemanagementManagementunitTimeoffplanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch workforcemanagement managementunit timeoffplan params
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch workforcemanagement managementunit timeoffplan params
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) WithBody(body *models.UpdateTimeOffPlanRequest) *PatchWorkforcemanagementManagementunitTimeoffplanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch workforcemanagement managementunit timeoffplan params
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) SetBody(body *models.UpdateTimeOffPlanRequest) {
	o.Body = body
}

// WithManagementUnitID adds the managementUnitID to the patch workforcemanagement managementunit timeoffplan params
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) WithManagementUnitID(managementUnitID string) *PatchWorkforcemanagementManagementunitTimeoffplanParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the patch workforcemanagement managementunit timeoffplan params
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WithTimeOffPlanID adds the timeOffPlanID to the patch workforcemanagement managementunit timeoffplan params
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) WithTimeOffPlanID(timeOffPlanID string) *PatchWorkforcemanagementManagementunitTimeoffplanParams {
	o.SetTimeOffPlanID(timeOffPlanID)
	return o
}

// SetTimeOffPlanID adds the timeOffPlanId to the patch workforcemanagement managementunit timeoffplan params
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) SetTimeOffPlanID(timeOffPlanID string) {
	o.TimeOffPlanID = timeOffPlanID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWorkforcemanagementManagementunitTimeoffplanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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