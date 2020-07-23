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

// NewGetArchitectDependencytrackingBuildParams creates a new GetArchitectDependencytrackingBuildParams object
// with the default values initialized.
func NewGetArchitectDependencytrackingBuildParams() *GetArchitectDependencytrackingBuildParams {

	return &GetArchitectDependencytrackingBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectDependencytrackingBuildParamsWithTimeout creates a new GetArchitectDependencytrackingBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArchitectDependencytrackingBuildParamsWithTimeout(timeout time.Duration) *GetArchitectDependencytrackingBuildParams {

	return &GetArchitectDependencytrackingBuildParams{

		timeout: timeout,
	}
}

// NewGetArchitectDependencytrackingBuildParamsWithContext creates a new GetArchitectDependencytrackingBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArchitectDependencytrackingBuildParamsWithContext(ctx context.Context) *GetArchitectDependencytrackingBuildParams {

	return &GetArchitectDependencytrackingBuildParams{

		Context: ctx,
	}
}

// NewGetArchitectDependencytrackingBuildParamsWithHTTPClient creates a new GetArchitectDependencytrackingBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArchitectDependencytrackingBuildParamsWithHTTPClient(client *http.Client) *GetArchitectDependencytrackingBuildParams {

	return &GetArchitectDependencytrackingBuildParams{
		HTTPClient: client,
	}
}

/*GetArchitectDependencytrackingBuildParams contains all the parameters to send to the API endpoint
for the get architect dependencytracking build operation typically these are written to a http.Request
*/
type GetArchitectDependencytrackingBuildParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get architect dependencytracking build params
func (o *GetArchitectDependencytrackingBuildParams) WithTimeout(timeout time.Duration) *GetArchitectDependencytrackingBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect dependencytracking build params
func (o *GetArchitectDependencytrackingBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect dependencytracking build params
func (o *GetArchitectDependencytrackingBuildParams) WithContext(ctx context.Context) *GetArchitectDependencytrackingBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect dependencytracking build params
func (o *GetArchitectDependencytrackingBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect dependencytracking build params
func (o *GetArchitectDependencytrackingBuildParams) WithHTTPClient(client *http.Client) *GetArchitectDependencytrackingBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect dependencytracking build params
func (o *GetArchitectDependencytrackingBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectDependencytrackingBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
