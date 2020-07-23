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

// NewPatchWorkforcemanagementBusinessunitParams creates a new PatchWorkforcemanagementBusinessunitParams object
// with the default values initialized.
func NewPatchWorkforcemanagementBusinessunitParams() *PatchWorkforcemanagementBusinessunitParams {
	var ()
	return &PatchWorkforcemanagementBusinessunitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWorkforcemanagementBusinessunitParamsWithTimeout creates a new PatchWorkforcemanagementBusinessunitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchWorkforcemanagementBusinessunitParamsWithTimeout(timeout time.Duration) *PatchWorkforcemanagementBusinessunitParams {
	var ()
	return &PatchWorkforcemanagementBusinessunitParams{

		timeout: timeout,
	}
}

// NewPatchWorkforcemanagementBusinessunitParamsWithContext creates a new PatchWorkforcemanagementBusinessunitParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchWorkforcemanagementBusinessunitParamsWithContext(ctx context.Context) *PatchWorkforcemanagementBusinessunitParams {
	var ()
	return &PatchWorkforcemanagementBusinessunitParams{

		Context: ctx,
	}
}

// NewPatchWorkforcemanagementBusinessunitParamsWithHTTPClient creates a new PatchWorkforcemanagementBusinessunitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchWorkforcemanagementBusinessunitParamsWithHTTPClient(client *http.Client) *PatchWorkforcemanagementBusinessunitParams {
	var ()
	return &PatchWorkforcemanagementBusinessunitParams{
		HTTPClient: client,
	}
}

/*PatchWorkforcemanagementBusinessunitParams contains all the parameters to send to the API endpoint
for the patch workforcemanagement businessunit operation typically these are written to a http.Request
*/
type PatchWorkforcemanagementBusinessunitParams struct {

	/*Body
	  body

	*/
	Body *models.UpdateBusinessUnitRequest
	/*BusinessUnitID
	  The ID of the business unit, or 'mine' for the business unit of the logged-in user.

	*/
	BusinessUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch workforcemanagement businessunit params
func (o *PatchWorkforcemanagementBusinessunitParams) WithTimeout(timeout time.Duration) *PatchWorkforcemanagementBusinessunitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch workforcemanagement businessunit params
func (o *PatchWorkforcemanagementBusinessunitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch workforcemanagement businessunit params
func (o *PatchWorkforcemanagementBusinessunitParams) WithContext(ctx context.Context) *PatchWorkforcemanagementBusinessunitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch workforcemanagement businessunit params
func (o *PatchWorkforcemanagementBusinessunitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch workforcemanagement businessunit params
func (o *PatchWorkforcemanagementBusinessunitParams) WithHTTPClient(client *http.Client) *PatchWorkforcemanagementBusinessunitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch workforcemanagement businessunit params
func (o *PatchWorkforcemanagementBusinessunitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch workforcemanagement businessunit params
func (o *PatchWorkforcemanagementBusinessunitParams) WithBody(body *models.UpdateBusinessUnitRequest) *PatchWorkforcemanagementBusinessunitParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch workforcemanagement businessunit params
func (o *PatchWorkforcemanagementBusinessunitParams) SetBody(body *models.UpdateBusinessUnitRequest) {
	o.Body = body
}

// WithBusinessUnitID adds the businessUnitID to the patch workforcemanagement businessunit params
func (o *PatchWorkforcemanagementBusinessunitParams) WithBusinessUnitID(businessUnitID string) *PatchWorkforcemanagementBusinessunitParams {
	o.SetBusinessUnitID(businessUnitID)
	return o
}

// SetBusinessUnitID adds the businessUnitId to the patch workforcemanagement businessunit params
func (o *PatchWorkforcemanagementBusinessunitParams) SetBusinessUnitID(businessUnitID string) {
	o.BusinessUnitID = businessUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWorkforcemanagementBusinessunitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
