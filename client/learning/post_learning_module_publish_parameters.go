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

// NewPostLearningModulePublishParams creates a new PostLearningModulePublishParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLearningModulePublishParams() *PostLearningModulePublishParams {
	return &PostLearningModulePublishParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLearningModulePublishParamsWithTimeout creates a new PostLearningModulePublishParams object
// with the ability to set a timeout on a request.
func NewPostLearningModulePublishParamsWithTimeout(timeout time.Duration) *PostLearningModulePublishParams {
	return &PostLearningModulePublishParams{
		timeout: timeout,
	}
}

// NewPostLearningModulePublishParamsWithContext creates a new PostLearningModulePublishParams object
// with the ability to set a context for a request.
func NewPostLearningModulePublishParamsWithContext(ctx context.Context) *PostLearningModulePublishParams {
	return &PostLearningModulePublishParams{
		Context: ctx,
	}
}

// NewPostLearningModulePublishParamsWithHTTPClient creates a new PostLearningModulePublishParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLearningModulePublishParamsWithHTTPClient(client *http.Client) *PostLearningModulePublishParams {
	return &PostLearningModulePublishParams{
		HTTPClient: client,
	}
}

/*
PostLearningModulePublishParams contains all the parameters to send to the API endpoint

	for the post learning module publish operation.

	Typically these are written to a http.Request.
*/
type PostLearningModulePublishParams struct {

	/* ModuleID.

	   The ID of the learning module
	*/
	ModuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post learning module publish params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLearningModulePublishParams) WithDefaults() *PostLearningModulePublishParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post learning module publish params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLearningModulePublishParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post learning module publish params
func (o *PostLearningModulePublishParams) WithTimeout(timeout time.Duration) *PostLearningModulePublishParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post learning module publish params
func (o *PostLearningModulePublishParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post learning module publish params
func (o *PostLearningModulePublishParams) WithContext(ctx context.Context) *PostLearningModulePublishParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post learning module publish params
func (o *PostLearningModulePublishParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post learning module publish params
func (o *PostLearningModulePublishParams) WithHTTPClient(client *http.Client) *PostLearningModulePublishParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post learning module publish params
func (o *PostLearningModulePublishParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithModuleID adds the moduleID to the post learning module publish params
func (o *PostLearningModulePublishParams) WithModuleID(moduleID string) *PostLearningModulePublishParams {
	o.SetModuleID(moduleID)
	return o
}

// SetModuleID adds the moduleId to the post learning module publish params
func (o *PostLearningModulePublishParams) SetModuleID(moduleID string) {
	o.ModuleID = moduleID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLearningModulePublishParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
