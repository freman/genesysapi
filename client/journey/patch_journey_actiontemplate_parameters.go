// Code generated by go-swagger; DO NOT EDIT.

package journey

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

// NewPatchJourneyActiontemplateParams creates a new PatchJourneyActiontemplateParams object
// with the default values initialized.
func NewPatchJourneyActiontemplateParams() *PatchJourneyActiontemplateParams {
	var ()
	return &PatchJourneyActiontemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchJourneyActiontemplateParamsWithTimeout creates a new PatchJourneyActiontemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchJourneyActiontemplateParamsWithTimeout(timeout time.Duration) *PatchJourneyActiontemplateParams {
	var ()
	return &PatchJourneyActiontemplateParams{

		timeout: timeout,
	}
}

// NewPatchJourneyActiontemplateParamsWithContext creates a new PatchJourneyActiontemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchJourneyActiontemplateParamsWithContext(ctx context.Context) *PatchJourneyActiontemplateParams {
	var ()
	return &PatchJourneyActiontemplateParams{

		Context: ctx,
	}
}

// NewPatchJourneyActiontemplateParamsWithHTTPClient creates a new PatchJourneyActiontemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchJourneyActiontemplateParamsWithHTTPClient(client *http.Client) *PatchJourneyActiontemplateParams {
	var ()
	return &PatchJourneyActiontemplateParams{
		HTTPClient: client,
	}
}

/*PatchJourneyActiontemplateParams contains all the parameters to send to the API endpoint
for the patch journey actiontemplate operation typically these are written to a http.Request
*/
type PatchJourneyActiontemplateParams struct {

	/*ActionTemplateID
	  ID of the action template.

	*/
	ActionTemplateID string
	/*Body*/
	Body *models.PatchActionTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch journey actiontemplate params
func (o *PatchJourneyActiontemplateParams) WithTimeout(timeout time.Duration) *PatchJourneyActiontemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch journey actiontemplate params
func (o *PatchJourneyActiontemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch journey actiontemplate params
func (o *PatchJourneyActiontemplateParams) WithContext(ctx context.Context) *PatchJourneyActiontemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch journey actiontemplate params
func (o *PatchJourneyActiontemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch journey actiontemplate params
func (o *PatchJourneyActiontemplateParams) WithHTTPClient(client *http.Client) *PatchJourneyActiontemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch journey actiontemplate params
func (o *PatchJourneyActiontemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionTemplateID adds the actionTemplateID to the patch journey actiontemplate params
func (o *PatchJourneyActiontemplateParams) WithActionTemplateID(actionTemplateID string) *PatchJourneyActiontemplateParams {
	o.SetActionTemplateID(actionTemplateID)
	return o
}

// SetActionTemplateID adds the actionTemplateId to the patch journey actiontemplate params
func (o *PatchJourneyActiontemplateParams) SetActionTemplateID(actionTemplateID string) {
	o.ActionTemplateID = actionTemplateID
}

// WithBody adds the body to the patch journey actiontemplate params
func (o *PatchJourneyActiontemplateParams) WithBody(body *models.PatchActionTemplate) *PatchJourneyActiontemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch journey actiontemplate params
func (o *PatchJourneyActiontemplateParams) SetBody(body *models.PatchActionTemplate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchJourneyActiontemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionTemplateId
	if err := r.SetPathParam("actionTemplateId", o.ActionTemplateID); err != nil {
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