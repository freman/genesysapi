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

	"github.com/freman/genesysapi/models"
)

// NewPatchAuthorizationRoleParams creates a new PatchAuthorizationRoleParams object
// with the default values initialized.
func NewPatchAuthorizationRoleParams() *PatchAuthorizationRoleParams {
	var ()
	return &PatchAuthorizationRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAuthorizationRoleParamsWithTimeout creates a new PatchAuthorizationRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAuthorizationRoleParamsWithTimeout(timeout time.Duration) *PatchAuthorizationRoleParams {
	var ()
	return &PatchAuthorizationRoleParams{

		timeout: timeout,
	}
}

// NewPatchAuthorizationRoleParamsWithContext creates a new PatchAuthorizationRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAuthorizationRoleParamsWithContext(ctx context.Context) *PatchAuthorizationRoleParams {
	var ()
	return &PatchAuthorizationRoleParams{

		Context: ctx,
	}
}

// NewPatchAuthorizationRoleParamsWithHTTPClient creates a new PatchAuthorizationRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAuthorizationRoleParamsWithHTTPClient(client *http.Client) *PatchAuthorizationRoleParams {
	var ()
	return &PatchAuthorizationRoleParams{
		HTTPClient: client,
	}
}

/*PatchAuthorizationRoleParams contains all the parameters to send to the API endpoint
for the patch authorization role operation typically these are written to a http.Request
*/
type PatchAuthorizationRoleParams struct {

	/*Body
	  Organization role

	*/
	Body *models.DomainOrganizationRole
	/*RoleID
	  Role ID

	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch authorization role params
func (o *PatchAuthorizationRoleParams) WithTimeout(timeout time.Duration) *PatchAuthorizationRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch authorization role params
func (o *PatchAuthorizationRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch authorization role params
func (o *PatchAuthorizationRoleParams) WithContext(ctx context.Context) *PatchAuthorizationRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch authorization role params
func (o *PatchAuthorizationRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch authorization role params
func (o *PatchAuthorizationRoleParams) WithHTTPClient(client *http.Client) *PatchAuthorizationRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch authorization role params
func (o *PatchAuthorizationRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch authorization role params
func (o *PatchAuthorizationRoleParams) WithBody(body *models.DomainOrganizationRole) *PatchAuthorizationRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch authorization role params
func (o *PatchAuthorizationRoleParams) SetBody(body *models.DomainOrganizationRole) {
	o.Body = body
}

// WithRoleID adds the roleID to the patch authorization role params
func (o *PatchAuthorizationRoleParams) WithRoleID(roleID string) *PatchAuthorizationRoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the patch authorization role params
func (o *PatchAuthorizationRoleParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAuthorizationRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
