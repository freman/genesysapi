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

// NewGetAuthorizationRoleParams creates a new GetAuthorizationRoleParams object
// with the default values initialized.
func NewGetAuthorizationRoleParams() *GetAuthorizationRoleParams {
	var ()
	return &GetAuthorizationRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationRoleParamsWithTimeout creates a new GetAuthorizationRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthorizationRoleParamsWithTimeout(timeout time.Duration) *GetAuthorizationRoleParams {
	var ()
	return &GetAuthorizationRoleParams{

		timeout: timeout,
	}
}

// NewGetAuthorizationRoleParamsWithContext creates a new GetAuthorizationRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthorizationRoleParamsWithContext(ctx context.Context) *GetAuthorizationRoleParams {
	var ()
	return &GetAuthorizationRoleParams{

		Context: ctx,
	}
}

// NewGetAuthorizationRoleParamsWithHTTPClient creates a new GetAuthorizationRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthorizationRoleParamsWithHTTPClient(client *http.Client) *GetAuthorizationRoleParams {
	var ()
	return &GetAuthorizationRoleParams{
		HTTPClient: client,
	}
}

/*GetAuthorizationRoleParams contains all the parameters to send to the API endpoint
for the get authorization role operation typically these are written to a http.Request
*/
type GetAuthorizationRoleParams struct {

	/*Expand
	  Which fields, if any, to expand. "unusedPermissions" returns the permissions not used for the role

	*/
	Expand []string
	/*RoleID
	  Role ID

	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get authorization role params
func (o *GetAuthorizationRoleParams) WithTimeout(timeout time.Duration) *GetAuthorizationRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization role params
func (o *GetAuthorizationRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization role params
func (o *GetAuthorizationRoleParams) WithContext(ctx context.Context) *GetAuthorizationRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization role params
func (o *GetAuthorizationRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization role params
func (o *GetAuthorizationRoleParams) WithHTTPClient(client *http.Client) *GetAuthorizationRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization role params
func (o *GetAuthorizationRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get authorization role params
func (o *GetAuthorizationRoleParams) WithExpand(expand []string) *GetAuthorizationRoleParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get authorization role params
func (o *GetAuthorizationRoleParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithRoleID adds the roleID to the get authorization role params
func (o *GetAuthorizationRoleParams) WithRoleID(roleID string) *GetAuthorizationRoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the get authorization role params
func (o *GetAuthorizationRoleParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	// path param roleId
	if err := r.SetPathParam("roleId", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}