// Code generated by go-swagger; DO NOT EDIT.

package learning

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

// NewPutLearningModuleParams creates a new PutLearningModuleParams object
// with the default values initialized.
func NewPutLearningModuleParams() *PutLearningModuleParams {
	var ()
	return &PutLearningModuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLearningModuleParamsWithTimeout creates a new PutLearningModuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLearningModuleParamsWithTimeout(timeout time.Duration) *PutLearningModuleParams {
	var ()
	return &PutLearningModuleParams{

		timeout: timeout,
	}
}

// NewPutLearningModuleParamsWithContext creates a new PutLearningModuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLearningModuleParamsWithContext(ctx context.Context) *PutLearningModuleParams {
	var ()
	return &PutLearningModuleParams{

		Context: ctx,
	}
}

// NewPutLearningModuleParamsWithHTTPClient creates a new PutLearningModuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLearningModuleParamsWithHTTPClient(client *http.Client) *PutLearningModuleParams {
	var ()
	return &PutLearningModuleParams{
		HTTPClient: client,
	}
}

/*PutLearningModuleParams contains all the parameters to send to the API endpoint
for the put learning module operation typically these are written to a http.Request
*/
type PutLearningModuleParams struct {

	/*Body
	  The learning module to be updated

	*/
	Body *models.LearningModuleRequest
	/*ModuleID
	  The ID of the learning module

	*/
	ModuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put learning module params
func (o *PutLearningModuleParams) WithTimeout(timeout time.Duration) *PutLearningModuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put learning module params
func (o *PutLearningModuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put learning module params
func (o *PutLearningModuleParams) WithContext(ctx context.Context) *PutLearningModuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put learning module params
func (o *PutLearningModuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put learning module params
func (o *PutLearningModuleParams) WithHTTPClient(client *http.Client) *PutLearningModuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put learning module params
func (o *PutLearningModuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put learning module params
func (o *PutLearningModuleParams) WithBody(body *models.LearningModuleRequest) *PutLearningModuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put learning module params
func (o *PutLearningModuleParams) SetBody(body *models.LearningModuleRequest) {
	o.Body = body
}

// WithModuleID adds the moduleID to the put learning module params
func (o *PutLearningModuleParams) WithModuleID(moduleID string) *PutLearningModuleParams {
	o.SetModuleID(moduleID)
	return o
}

// SetModuleID adds the moduleId to the put learning module params
func (o *PutLearningModuleParams) SetModuleID(moduleID string) {
	o.ModuleID = moduleID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLearningModuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param moduleId
	if err := r.SetPathParam("moduleId", o.ModuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
