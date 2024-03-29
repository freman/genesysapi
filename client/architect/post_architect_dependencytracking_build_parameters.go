// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewPostArchitectDependencytrackingBuildParams creates a new PostArchitectDependencytrackingBuildParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostArchitectDependencytrackingBuildParams() *PostArchitectDependencytrackingBuildParams {
	return &PostArchitectDependencytrackingBuildParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostArchitectDependencytrackingBuildParamsWithTimeout creates a new PostArchitectDependencytrackingBuildParams object
// with the ability to set a timeout on a request.
func NewPostArchitectDependencytrackingBuildParamsWithTimeout(timeout time.Duration) *PostArchitectDependencytrackingBuildParams {
	return &PostArchitectDependencytrackingBuildParams{
		timeout: timeout,
	}
}

// NewPostArchitectDependencytrackingBuildParamsWithContext creates a new PostArchitectDependencytrackingBuildParams object
// with the ability to set a context for a request.
func NewPostArchitectDependencytrackingBuildParamsWithContext(ctx context.Context) *PostArchitectDependencytrackingBuildParams {
	return &PostArchitectDependencytrackingBuildParams{
		Context: ctx,
	}
}

// NewPostArchitectDependencytrackingBuildParamsWithHTTPClient creates a new PostArchitectDependencytrackingBuildParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostArchitectDependencytrackingBuildParamsWithHTTPClient(client *http.Client) *PostArchitectDependencytrackingBuildParams {
	return &PostArchitectDependencytrackingBuildParams{
		HTTPClient: client,
	}
}

/*
PostArchitectDependencytrackingBuildParams contains all the parameters to send to the API endpoint

	for the post architect dependencytracking build operation.

	Typically these are written to a http.Request.
*/
type PostArchitectDependencytrackingBuildParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post architect dependencytracking build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostArchitectDependencytrackingBuildParams) WithDefaults() *PostArchitectDependencytrackingBuildParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post architect dependencytracking build params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostArchitectDependencytrackingBuildParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post architect dependencytracking build params
func (o *PostArchitectDependencytrackingBuildParams) WithTimeout(timeout time.Duration) *PostArchitectDependencytrackingBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post architect dependencytracking build params
func (o *PostArchitectDependencytrackingBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post architect dependencytracking build params
func (o *PostArchitectDependencytrackingBuildParams) WithContext(ctx context.Context) *PostArchitectDependencytrackingBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post architect dependencytracking build params
func (o *PostArchitectDependencytrackingBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post architect dependencytracking build params
func (o *PostArchitectDependencytrackingBuildParams) WithHTTPClient(client *http.Client) *PostArchitectDependencytrackingBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post architect dependencytracking build params
func (o *PostArchitectDependencytrackingBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostArchitectDependencytrackingBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
