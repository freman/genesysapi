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

// NewPostLearningModuleJobsParams creates a new PostLearningModuleJobsParams object
// with the default values initialized.
func NewPostLearningModuleJobsParams() *PostLearningModuleJobsParams {
	var ()
	return &PostLearningModuleJobsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLearningModuleJobsParamsWithTimeout creates a new PostLearningModuleJobsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLearningModuleJobsParamsWithTimeout(timeout time.Duration) *PostLearningModuleJobsParams {
	var ()
	return &PostLearningModuleJobsParams{

		timeout: timeout,
	}
}

// NewPostLearningModuleJobsParamsWithContext creates a new PostLearningModuleJobsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLearningModuleJobsParamsWithContext(ctx context.Context) *PostLearningModuleJobsParams {
	var ()
	return &PostLearningModuleJobsParams{

		Context: ctx,
	}
}

// NewPostLearningModuleJobsParamsWithHTTPClient creates a new PostLearningModuleJobsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLearningModuleJobsParamsWithHTTPClient(client *http.Client) *PostLearningModuleJobsParams {
	var ()
	return &PostLearningModuleJobsParams{
		HTTPClient: client,
	}
}

/*PostLearningModuleJobsParams contains all the parameters to send to the API endpoint
for the post learning module jobs operation typically these are written to a http.Request
*/
type PostLearningModuleJobsParams struct {

	/*Body
	  The learning module job request

	*/
	Body *models.LearningModuleJobRequest
	/*ModuleID
	  The ID of the learning module

	*/
	ModuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post learning module jobs params
func (o *PostLearningModuleJobsParams) WithTimeout(timeout time.Duration) *PostLearningModuleJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post learning module jobs params
func (o *PostLearningModuleJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post learning module jobs params
func (o *PostLearningModuleJobsParams) WithContext(ctx context.Context) *PostLearningModuleJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post learning module jobs params
func (o *PostLearningModuleJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post learning module jobs params
func (o *PostLearningModuleJobsParams) WithHTTPClient(client *http.Client) *PostLearningModuleJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post learning module jobs params
func (o *PostLearningModuleJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post learning module jobs params
func (o *PostLearningModuleJobsParams) WithBody(body *models.LearningModuleJobRequest) *PostLearningModuleJobsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post learning module jobs params
func (o *PostLearningModuleJobsParams) SetBody(body *models.LearningModuleJobRequest) {
	o.Body = body
}

// WithModuleID adds the moduleID to the post learning module jobs params
func (o *PostLearningModuleJobsParams) WithModuleID(moduleID string) *PostLearningModuleJobsParams {
	o.SetModuleID(moduleID)
	return o
}

// SetModuleID adds the moduleId to the post learning module jobs params
func (o *PostLearningModuleJobsParams) SetModuleID(moduleID string) {
	o.ModuleID = moduleID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLearningModuleJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
