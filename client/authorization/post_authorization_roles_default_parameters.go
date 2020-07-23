// Code generated by go-swagger; DO NOT EDIT.

package authorization

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
	"github.com/go-openapi/swag"
)

// NewPostAuthorizationRolesDefaultParams creates a new PostAuthorizationRolesDefaultParams object
// with the default values initialized.
func NewPostAuthorizationRolesDefaultParams() *PostAuthorizationRolesDefaultParams {
	var (
		forceDefault = bool(false)
	)
	return &PostAuthorizationRolesDefaultParams{
		Force: &forceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuthorizationRolesDefaultParamsWithTimeout creates a new PostAuthorizationRolesDefaultParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAuthorizationRolesDefaultParamsWithTimeout(timeout time.Duration) *PostAuthorizationRolesDefaultParams {
	var (
		forceDefault = bool(false)
	)
	return &PostAuthorizationRolesDefaultParams{
		Force: &forceDefault,

		timeout: timeout,
	}
}

// NewPostAuthorizationRolesDefaultParamsWithContext creates a new PostAuthorizationRolesDefaultParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAuthorizationRolesDefaultParamsWithContext(ctx context.Context) *PostAuthorizationRolesDefaultParams {
	var (
		forceDefault = bool(false)
	)
	return &PostAuthorizationRolesDefaultParams{
		Force: &forceDefault,

		Context: ctx,
	}
}

// NewPostAuthorizationRolesDefaultParamsWithHTTPClient creates a new PostAuthorizationRolesDefaultParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAuthorizationRolesDefaultParamsWithHTTPClient(client *http.Client) *PostAuthorizationRolesDefaultParams {
	var (
		forceDefault = bool(false)
	)
	return &PostAuthorizationRolesDefaultParams{
		Force:      &forceDefault,
		HTTPClient: client,
	}
}

/*PostAuthorizationRolesDefaultParams contains all the parameters to send to the API endpoint
for the post authorization roles default operation typically these are written to a http.Request
*/
type PostAuthorizationRolesDefaultParams struct {

	/*Force
	  Restore default roles

	*/
	Force *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post authorization roles default params
func (o *PostAuthorizationRolesDefaultParams) WithTimeout(timeout time.Duration) *PostAuthorizationRolesDefaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post authorization roles default params
func (o *PostAuthorizationRolesDefaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post authorization roles default params
func (o *PostAuthorizationRolesDefaultParams) WithContext(ctx context.Context) *PostAuthorizationRolesDefaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post authorization roles default params
func (o *PostAuthorizationRolesDefaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post authorization roles default params
func (o *PostAuthorizationRolesDefaultParams) WithHTTPClient(client *http.Client) *PostAuthorizationRolesDefaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post authorization roles default params
func (o *PostAuthorizationRolesDefaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the post authorization roles default params
func (o *PostAuthorizationRolesDefaultParams) WithForce(force *bool) *PostAuthorizationRolesDefaultParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the post authorization roles default params
func (o *PostAuthorizationRolesDefaultParams) SetForce(force *bool) {
	o.Force = force
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuthorizationRolesDefaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
