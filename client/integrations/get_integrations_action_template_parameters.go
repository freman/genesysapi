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
)

// NewGetIntegrationsActionTemplateParams creates a new GetIntegrationsActionTemplateParams object
// with the default values initialized.
func NewGetIntegrationsActionTemplateParams() *GetIntegrationsActionTemplateParams {
	var ()
	return &GetIntegrationsActionTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationsActionTemplateParamsWithTimeout creates a new GetIntegrationsActionTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIntegrationsActionTemplateParamsWithTimeout(timeout time.Duration) *GetIntegrationsActionTemplateParams {
	var ()
	return &GetIntegrationsActionTemplateParams{

		timeout: timeout,
	}
}

// NewGetIntegrationsActionTemplateParamsWithContext creates a new GetIntegrationsActionTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIntegrationsActionTemplateParamsWithContext(ctx context.Context) *GetIntegrationsActionTemplateParams {
	var ()
	return &GetIntegrationsActionTemplateParams{

		Context: ctx,
	}
}

// NewGetIntegrationsActionTemplateParamsWithHTTPClient creates a new GetIntegrationsActionTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIntegrationsActionTemplateParamsWithHTTPClient(client *http.Client) *GetIntegrationsActionTemplateParams {
	var ()
	return &GetIntegrationsActionTemplateParams{
		HTTPClient: client,
	}
}

/*GetIntegrationsActionTemplateParams contains all the parameters to send to the API endpoint
for the get integrations action template operation typically these are written to a http.Request
*/
type GetIntegrationsActionTemplateParams struct {

	/*ActionID
	  actionId

	*/
	ActionID string
	/*FileName
	  Name of template file to be retrieved for this action.

	*/
	FileName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get integrations action template params
func (o *GetIntegrationsActionTemplateParams) WithTimeout(timeout time.Duration) *GetIntegrationsActionTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrations action template params
func (o *GetIntegrationsActionTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrations action template params
func (o *GetIntegrationsActionTemplateParams) WithContext(ctx context.Context) *GetIntegrationsActionTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrations action template params
func (o *GetIntegrationsActionTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrations action template params
func (o *GetIntegrationsActionTemplateParams) WithHTTPClient(client *http.Client) *GetIntegrationsActionTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrations action template params
func (o *GetIntegrationsActionTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the get integrations action template params
func (o *GetIntegrationsActionTemplateParams) WithActionID(actionID string) *GetIntegrationsActionTemplateParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the get integrations action template params
func (o *GetIntegrationsActionTemplateParams) SetActionID(actionID string) {
	o.ActionID = actionID
}

// WithFileName adds the fileName to the get integrations action template params
func (o *GetIntegrationsActionTemplateParams) WithFileName(fileName string) *GetIntegrationsActionTemplateParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the get integrations action template params
func (o *GetIntegrationsActionTemplateParams) SetFileName(fileName string) {
	o.FileName = fileName
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationsActionTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID); err != nil {
		return err
	}

	// path param fileName
	if err := r.SetPathParam("fileName", o.FileName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
