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

// NewPatchWorkforcemanagementTimeoffrequestParams creates a new PatchWorkforcemanagementTimeoffrequestParams object
// with the default values initialized.
func NewPatchWorkforcemanagementTimeoffrequestParams() *PatchWorkforcemanagementTimeoffrequestParams {
	var ()
	return &PatchWorkforcemanagementTimeoffrequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWorkforcemanagementTimeoffrequestParamsWithTimeout creates a new PatchWorkforcemanagementTimeoffrequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchWorkforcemanagementTimeoffrequestParamsWithTimeout(timeout time.Duration) *PatchWorkforcemanagementTimeoffrequestParams {
	var ()
	return &PatchWorkforcemanagementTimeoffrequestParams{

		timeout: timeout,
	}
}

// NewPatchWorkforcemanagementTimeoffrequestParamsWithContext creates a new PatchWorkforcemanagementTimeoffrequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchWorkforcemanagementTimeoffrequestParamsWithContext(ctx context.Context) *PatchWorkforcemanagementTimeoffrequestParams {
	var ()
	return &PatchWorkforcemanagementTimeoffrequestParams{

		Context: ctx,
	}
}

// NewPatchWorkforcemanagementTimeoffrequestParamsWithHTTPClient creates a new PatchWorkforcemanagementTimeoffrequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchWorkforcemanagementTimeoffrequestParamsWithHTTPClient(client *http.Client) *PatchWorkforcemanagementTimeoffrequestParams {
	var ()
	return &PatchWorkforcemanagementTimeoffrequestParams{
		HTTPClient: client,
	}
}

/*PatchWorkforcemanagementTimeoffrequestParams contains all the parameters to send to the API endpoint
for the patch workforcemanagement timeoffrequest operation typically these are written to a http.Request
*/
type PatchWorkforcemanagementTimeoffrequestParams struct {

	/*Body
	  body

	*/
	Body *models.AgentTimeOffRequestPatch
	/*TimeOffRequestID
	  Time Off Request Id

	*/
	TimeOffRequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch workforcemanagement timeoffrequest params
func (o *PatchWorkforcemanagementTimeoffrequestParams) WithTimeout(timeout time.Duration) *PatchWorkforcemanagementTimeoffrequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch workforcemanagement timeoffrequest params
func (o *PatchWorkforcemanagementTimeoffrequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch workforcemanagement timeoffrequest params
func (o *PatchWorkforcemanagementTimeoffrequestParams) WithContext(ctx context.Context) *PatchWorkforcemanagementTimeoffrequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch workforcemanagement timeoffrequest params
func (o *PatchWorkforcemanagementTimeoffrequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch workforcemanagement timeoffrequest params
func (o *PatchWorkforcemanagementTimeoffrequestParams) WithHTTPClient(client *http.Client) *PatchWorkforcemanagementTimeoffrequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch workforcemanagement timeoffrequest params
func (o *PatchWorkforcemanagementTimeoffrequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch workforcemanagement timeoffrequest params
func (o *PatchWorkforcemanagementTimeoffrequestParams) WithBody(body *models.AgentTimeOffRequestPatch) *PatchWorkforcemanagementTimeoffrequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch workforcemanagement timeoffrequest params
func (o *PatchWorkforcemanagementTimeoffrequestParams) SetBody(body *models.AgentTimeOffRequestPatch) {
	o.Body = body
}

// WithTimeOffRequestID adds the timeOffRequestID to the patch workforcemanagement timeoffrequest params
func (o *PatchWorkforcemanagementTimeoffrequestParams) WithTimeOffRequestID(timeOffRequestID string) *PatchWorkforcemanagementTimeoffrequestParams {
	o.SetTimeOffRequestID(timeOffRequestID)
	return o
}

// SetTimeOffRequestID adds the timeOffRequestId to the patch workforcemanagement timeoffrequest params
func (o *PatchWorkforcemanagementTimeoffrequestParams) SetTimeOffRequestID(timeOffRequestID string) {
	o.TimeOffRequestID = timeOffRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWorkforcemanagementTimeoffrequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param timeOffRequestId
	if err := r.SetPathParam("timeOffRequestId", o.TimeOffRequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}