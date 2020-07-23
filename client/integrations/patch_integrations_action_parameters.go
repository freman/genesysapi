// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewPatchIntegrationsActionParams creates a new PatchIntegrationsActionParams object
// with the default values initialized.
func NewPatchIntegrationsActionParams() *PatchIntegrationsActionParams {
	var ()
	return &PatchIntegrationsActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchIntegrationsActionParamsWithTimeout creates a new PatchIntegrationsActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchIntegrationsActionParamsWithTimeout(timeout time.Duration) *PatchIntegrationsActionParams {
	var ()
	return &PatchIntegrationsActionParams{

		timeout: timeout,
	}
}

// NewPatchIntegrationsActionParamsWithContext creates a new PatchIntegrationsActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchIntegrationsActionParamsWithContext(ctx context.Context) *PatchIntegrationsActionParams {
	var ()
	return &PatchIntegrationsActionParams{

		Context: ctx,
	}
}

// NewPatchIntegrationsActionParamsWithHTTPClient creates a new PatchIntegrationsActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchIntegrationsActionParamsWithHTTPClient(client *http.Client) *PatchIntegrationsActionParams {
	var ()
	return &PatchIntegrationsActionParams{
		HTTPClient: client,
	}
}

/*PatchIntegrationsActionParams contains all the parameters to send to the API endpoint
for the patch integrations action operation typically these are written to a http.Request
*/
type PatchIntegrationsActionParams struct {

	/*ActionID
	  actionId

	*/
	ActionID string
	/*Body
	  Input used to patch the Action.

	*/
	Body *models.UpdateActionInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch integrations action params
func (o *PatchIntegrationsActionParams) WithTimeout(timeout time.Duration) *PatchIntegrationsActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch integrations action params
func (o *PatchIntegrationsActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch integrations action params
func (o *PatchIntegrationsActionParams) WithContext(ctx context.Context) *PatchIntegrationsActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch integrations action params
func (o *PatchIntegrationsActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch integrations action params
func (o *PatchIntegrationsActionParams) WithHTTPClient(client *http.Client) *PatchIntegrationsActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch integrations action params
func (o *PatchIntegrationsActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the patch integrations action params
func (o *PatchIntegrationsActionParams) WithActionID(actionID string) *PatchIntegrationsActionParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the patch integrations action params
func (o *PatchIntegrationsActionParams) SetActionID(actionID string) {
	o.ActionID = actionID
}

// WithBody adds the body to the patch integrations action params
func (o *PatchIntegrationsActionParams) WithBody(body *models.UpdateActionInput) *PatchIntegrationsActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch integrations action params
func (o *PatchIntegrationsActionParams) SetBody(body *models.UpdateActionInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchIntegrationsActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
