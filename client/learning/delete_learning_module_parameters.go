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
)

// NewDeleteLearningModuleParams creates a new DeleteLearningModuleParams object
// with the default values initialized.
func NewDeleteLearningModuleParams() *DeleteLearningModuleParams {
	var ()
	return &DeleteLearningModuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLearningModuleParamsWithTimeout creates a new DeleteLearningModuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLearningModuleParamsWithTimeout(timeout time.Duration) *DeleteLearningModuleParams {
	var ()
	return &DeleteLearningModuleParams{

		timeout: timeout,
	}
}

// NewDeleteLearningModuleParamsWithContext creates a new DeleteLearningModuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLearningModuleParamsWithContext(ctx context.Context) *DeleteLearningModuleParams {
	var ()
	return &DeleteLearningModuleParams{

		Context: ctx,
	}
}

// NewDeleteLearningModuleParamsWithHTTPClient creates a new DeleteLearningModuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLearningModuleParamsWithHTTPClient(client *http.Client) *DeleteLearningModuleParams {
	var ()
	return &DeleteLearningModuleParams{
		HTTPClient: client,
	}
}

/*DeleteLearningModuleParams contains all the parameters to send to the API endpoint
for the delete learning module operation typically these are written to a http.Request
*/
type DeleteLearningModuleParams struct {

	/*ModuleID
	  The ID of the learning module

	*/
	ModuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete learning module params
func (o *DeleteLearningModuleParams) WithTimeout(timeout time.Duration) *DeleteLearningModuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete learning module params
func (o *DeleteLearningModuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete learning module params
func (o *DeleteLearningModuleParams) WithContext(ctx context.Context) *DeleteLearningModuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete learning module params
func (o *DeleteLearningModuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete learning module params
func (o *DeleteLearningModuleParams) WithHTTPClient(client *http.Client) *DeleteLearningModuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete learning module params
func (o *DeleteLearningModuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModuleID adds the moduleID to the delete learning module params
func (o *DeleteLearningModuleParams) WithModuleID(moduleID string) *DeleteLearningModuleParams {
	o.SetModuleID(moduleID)
	return o
}

// SetModuleID adds the moduleId to the delete learning module params
func (o *DeleteLearningModuleParams) SetModuleID(moduleID string) {
	o.ModuleID = moduleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLearningModuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param moduleId
	if err := r.SetPathParam("moduleId", o.ModuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
