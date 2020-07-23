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
)

// NewDeleteAuthorizationRoleParams creates a new DeleteAuthorizationRoleParams object
// with the default values initialized.
func NewDeleteAuthorizationRoleParams() *DeleteAuthorizationRoleParams {
	var ()
	return &DeleteAuthorizationRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAuthorizationRoleParamsWithTimeout creates a new DeleteAuthorizationRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAuthorizationRoleParamsWithTimeout(timeout time.Duration) *DeleteAuthorizationRoleParams {
	var ()
	return &DeleteAuthorizationRoleParams{

		timeout: timeout,
	}
}

// NewDeleteAuthorizationRoleParamsWithContext creates a new DeleteAuthorizationRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAuthorizationRoleParamsWithContext(ctx context.Context) *DeleteAuthorizationRoleParams {
	var ()
	return &DeleteAuthorizationRoleParams{

		Context: ctx,
	}
}

// NewDeleteAuthorizationRoleParamsWithHTTPClient creates a new DeleteAuthorizationRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAuthorizationRoleParamsWithHTTPClient(client *http.Client) *DeleteAuthorizationRoleParams {
	var ()
	return &DeleteAuthorizationRoleParams{
		HTTPClient: client,
	}
}

/*DeleteAuthorizationRoleParams contains all the parameters to send to the API endpoint
for the delete authorization role operation typically these are written to a http.Request
*/
type DeleteAuthorizationRoleParams struct {

	/*RoleID
	  Role ID

	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete authorization role params
func (o *DeleteAuthorizationRoleParams) WithTimeout(timeout time.Duration) *DeleteAuthorizationRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete authorization role params
func (o *DeleteAuthorizationRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete authorization role params
func (o *DeleteAuthorizationRoleParams) WithContext(ctx context.Context) *DeleteAuthorizationRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete authorization role params
func (o *DeleteAuthorizationRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete authorization role params
func (o *DeleteAuthorizationRoleParams) WithHTTPClient(client *http.Client) *DeleteAuthorizationRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete authorization role params
func (o *DeleteAuthorizationRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleID adds the roleID to the delete authorization role params
func (o *DeleteAuthorizationRoleParams) WithRoleID(roleID string) *DeleteAuthorizationRoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the delete authorization role params
func (o *DeleteAuthorizationRoleParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAuthorizationRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param roleId
	if err := r.SetPathParam("roleId", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
