// Code generated by go-swagger; DO NOT EDIT.

package organization_authorization

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

// NewDeleteOrgauthorizationTrusteeGroupRolesParams creates a new DeleteOrgauthorizationTrusteeGroupRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrgauthorizationTrusteeGroupRolesParams() *DeleteOrgauthorizationTrusteeGroupRolesParams {
	return &DeleteOrgauthorizationTrusteeGroupRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrgauthorizationTrusteeGroupRolesParamsWithTimeout creates a new DeleteOrgauthorizationTrusteeGroupRolesParams object
// with the ability to set a timeout on a request.
func NewDeleteOrgauthorizationTrusteeGroupRolesParamsWithTimeout(timeout time.Duration) *DeleteOrgauthorizationTrusteeGroupRolesParams {
	return &DeleteOrgauthorizationTrusteeGroupRolesParams{
		timeout: timeout,
	}
}

// NewDeleteOrgauthorizationTrusteeGroupRolesParamsWithContext creates a new DeleteOrgauthorizationTrusteeGroupRolesParams object
// with the ability to set a context for a request.
func NewDeleteOrgauthorizationTrusteeGroupRolesParamsWithContext(ctx context.Context) *DeleteOrgauthorizationTrusteeGroupRolesParams {
	return &DeleteOrgauthorizationTrusteeGroupRolesParams{
		Context: ctx,
	}
}

// NewDeleteOrgauthorizationTrusteeGroupRolesParamsWithHTTPClient creates a new DeleteOrgauthorizationTrusteeGroupRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrgauthorizationTrusteeGroupRolesParamsWithHTTPClient(client *http.Client) *DeleteOrgauthorizationTrusteeGroupRolesParams {
	return &DeleteOrgauthorizationTrusteeGroupRolesParams{
		HTTPClient: client,
	}
}

/*
DeleteOrgauthorizationTrusteeGroupRolesParams contains all the parameters to send to the API endpoint

	for the delete orgauthorization trustee group roles operation.

	Typically these are written to a http.Request.
*/
type DeleteOrgauthorizationTrusteeGroupRolesParams struct {

	/* TrusteeGroupID.

	   Trustee Group Id
	*/
	TrusteeGroupID string

	/* TrusteeOrgID.

	   Trustee Organization Id
	*/
	TrusteeOrgID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete orgauthorization trustee group roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) WithDefaults() *DeleteOrgauthorizationTrusteeGroupRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete orgauthorization trustee group roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete orgauthorization trustee group roles params
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) WithTimeout(timeout time.Duration) *DeleteOrgauthorizationTrusteeGroupRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete orgauthorization trustee group roles params
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete orgauthorization trustee group roles params
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) WithContext(ctx context.Context) *DeleteOrgauthorizationTrusteeGroupRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete orgauthorization trustee group roles params
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete orgauthorization trustee group roles params
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) WithHTTPClient(client *http.Client) *DeleteOrgauthorizationTrusteeGroupRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete orgauthorization trustee group roles params
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrusteeGroupID adds the trusteeGroupID to the delete orgauthorization trustee group roles params
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) WithTrusteeGroupID(trusteeGroupID string) *DeleteOrgauthorizationTrusteeGroupRolesParams {
	o.SetTrusteeGroupID(trusteeGroupID)
	return o
}

// SetTrusteeGroupID adds the trusteeGroupId to the delete orgauthorization trustee group roles params
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) SetTrusteeGroupID(trusteeGroupID string) {
	o.TrusteeGroupID = trusteeGroupID
}

// WithTrusteeOrgID adds the trusteeOrgID to the delete orgauthorization trustee group roles params
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) WithTrusteeOrgID(trusteeOrgID string) *DeleteOrgauthorizationTrusteeGroupRolesParams {
	o.SetTrusteeOrgID(trusteeOrgID)
	return o
}

// SetTrusteeOrgID adds the trusteeOrgId to the delete orgauthorization trustee group roles params
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) SetTrusteeOrgID(trusteeOrgID string) {
	o.TrusteeOrgID = trusteeOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrgauthorizationTrusteeGroupRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param trusteeGroupId
	if err := r.SetPathParam("trusteeGroupId", o.TrusteeGroupID); err != nil {
		return err
	}

	// path param trusteeOrgId
	if err := r.SetPathParam("trusteeOrgId", o.TrusteeOrgID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}