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

// NewPatchWorkforcemanagementBusinessunitActivitycodeParams creates a new PatchWorkforcemanagementBusinessunitActivitycodeParams object
// with the default values initialized.
func NewPatchWorkforcemanagementBusinessunitActivitycodeParams() *PatchWorkforcemanagementBusinessunitActivitycodeParams {
	var ()
	return &PatchWorkforcemanagementBusinessunitActivitycodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeParamsWithTimeout creates a new PatchWorkforcemanagementBusinessunitActivitycodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchWorkforcemanagementBusinessunitActivitycodeParamsWithTimeout(timeout time.Duration) *PatchWorkforcemanagementBusinessunitActivitycodeParams {
	var ()
	return &PatchWorkforcemanagementBusinessunitActivitycodeParams{

		timeout: timeout,
	}
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeParamsWithContext creates a new PatchWorkforcemanagementBusinessunitActivitycodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchWorkforcemanagementBusinessunitActivitycodeParamsWithContext(ctx context.Context) *PatchWorkforcemanagementBusinessunitActivitycodeParams {
	var ()
	return &PatchWorkforcemanagementBusinessunitActivitycodeParams{

		Context: ctx,
	}
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeParamsWithHTTPClient creates a new PatchWorkforcemanagementBusinessunitActivitycodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchWorkforcemanagementBusinessunitActivitycodeParamsWithHTTPClient(client *http.Client) *PatchWorkforcemanagementBusinessunitActivitycodeParams {
	var ()
	return &PatchWorkforcemanagementBusinessunitActivitycodeParams{
		HTTPClient: client,
	}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeParams contains all the parameters to send to the API endpoint
for the patch workforcemanagement businessunit activitycode operation typically these are written to a http.Request
*/
type PatchWorkforcemanagementBusinessunitActivitycodeParams struct {

	/*ActivityCodeID
	  The ID of the activity code to update

	*/
	ActivityCodeID string
	/*Body
	  body

	*/
	Body *models.UpdateActivityCodeRequest
	/*BusinessUnitID
	  The ID of the business unit, or 'mine' for the business unit of the logged-in user.

	*/
	BusinessUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch workforcemanagement businessunit activitycode params
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) WithTimeout(timeout time.Duration) *PatchWorkforcemanagementBusinessunitActivitycodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch workforcemanagement businessunit activitycode params
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch workforcemanagement businessunit activitycode params
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) WithContext(ctx context.Context) *PatchWorkforcemanagementBusinessunitActivitycodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch workforcemanagement businessunit activitycode params
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch workforcemanagement businessunit activitycode params
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) WithHTTPClient(client *http.Client) *PatchWorkforcemanagementBusinessunitActivitycodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch workforcemanagement businessunit activitycode params
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActivityCodeID adds the activityCodeID to the patch workforcemanagement businessunit activitycode params
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) WithActivityCodeID(activityCodeID string) *PatchWorkforcemanagementBusinessunitActivitycodeParams {
	o.SetActivityCodeID(activityCodeID)
	return o
}

// SetActivityCodeID adds the activityCodeId to the patch workforcemanagement businessunit activitycode params
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) SetActivityCodeID(activityCodeID string) {
	o.ActivityCodeID = activityCodeID
}

// WithBody adds the body to the patch workforcemanagement businessunit activitycode params
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) WithBody(body *models.UpdateActivityCodeRequest) *PatchWorkforcemanagementBusinessunitActivitycodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch workforcemanagement businessunit activitycode params
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) SetBody(body *models.UpdateActivityCodeRequest) {
	o.Body = body
}

// WithBusinessUnitID adds the businessUnitID to the patch workforcemanagement businessunit activitycode params
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) WithBusinessUnitID(businessUnitID string) *PatchWorkforcemanagementBusinessunitActivitycodeParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the patch workforcemanagement businessunit activitycode params
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWorkforcemanagementBusinessunitActivitycodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param activityCodeId
	if err := r.SetPathParam("activityCodeId", o.ActivityCodeID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param businessUnitId
	if err := r.SetPathParam("businessUnitId", o.BusinessUnitID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
