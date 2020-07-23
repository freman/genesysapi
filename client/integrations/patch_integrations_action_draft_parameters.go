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

// NewPatchIntegrationsActionDraftParams creates a new PatchIntegrationsActionDraftParams object
// with the default values initialized.
func NewPatchIntegrationsActionDraftParams() *PatchIntegrationsActionDraftParams {
	var ()
	return &PatchIntegrationsActionDraftParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchIntegrationsActionDraftParamsWithTimeout creates a new PatchIntegrationsActionDraftParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchIntegrationsActionDraftParamsWithTimeout(timeout time.Duration) *PatchIntegrationsActionDraftParams {
	var ()
	return &PatchIntegrationsActionDraftParams{

		timeout: timeout,
	}
}

// NewPatchIntegrationsActionDraftParamsWithContext creates a new PatchIntegrationsActionDraftParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchIntegrationsActionDraftParamsWithContext(ctx context.Context) *PatchIntegrationsActionDraftParams {
	var ()
	return &PatchIntegrationsActionDraftParams{

		Context: ctx,
	}
}

// NewPatchIntegrationsActionDraftParamsWithHTTPClient creates a new PatchIntegrationsActionDraftParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchIntegrationsActionDraftParamsWithHTTPClient(client *http.Client) *PatchIntegrationsActionDraftParams {
	var ()
	return &PatchIntegrationsActionDraftParams{
		HTTPClient: client,
	}
}

/*PatchIntegrationsActionDraftParams contains all the parameters to send to the API endpoint
for the patch integrations action draft operation typically these are written to a http.Request
*/
type PatchIntegrationsActionDraftParams struct {

	/*ActionID
	  actionId

	*/
	ActionID string
	/*Body
	  Input used to patch the Action Draft.

	*/
	Body *models.UpdateDraftInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch integrations action draft params
func (o *PatchIntegrationsActionDraftParams) WithTimeout(timeout time.Duration) *PatchIntegrationsActionDraftParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch integrations action draft params
func (o *PatchIntegrationsActionDraftParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch integrations action draft params
func (o *PatchIntegrationsActionDraftParams) WithContext(ctx context.Context) *PatchIntegrationsActionDraftParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch integrations action draft params
func (o *PatchIntegrationsActionDraftParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch integrations action draft params
func (o *PatchIntegrationsActionDraftParams) WithHTTPClient(client *http.Client) *PatchIntegrationsActionDraftParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch integrations action draft params
func (o *PatchIntegrationsActionDraftParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the patch integrations action draft params
func (o *PatchIntegrationsActionDraftParams) WithActionID(actionID string) *PatchIntegrationsActionDraftParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the patch integrations action draft params
func (o *PatchIntegrationsActionDraftParams) SetActionID(actionID string) {
	o.ActionID = actionID
}

// WithBody adds the body to the patch integrations action draft params
func (o *PatchIntegrationsActionDraftParams) WithBody(body *models.UpdateDraftInput) *PatchIntegrationsActionDraftParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch integrations action draft params
func (o *PatchIntegrationsActionDraftParams) SetBody(body *models.UpdateDraftInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchIntegrationsActionDraftParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
