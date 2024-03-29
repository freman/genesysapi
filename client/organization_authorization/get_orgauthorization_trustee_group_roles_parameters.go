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

// NewGetOrgauthorizationTrusteeGroupRolesParams creates a new GetOrgauthorizationTrusteeGroupRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrgauthorizationTrusteeGroupRolesParams() *GetOrgauthorizationTrusteeGroupRolesParams {
	return &GetOrgauthorizationTrusteeGroupRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgauthorizationTrusteeGroupRolesParamsWithTimeout creates a new GetOrgauthorizationTrusteeGroupRolesParams object
// with the ability to set a timeout on a request.
func NewGetOrgauthorizationTrusteeGroupRolesParamsWithTimeout(timeout time.Duration) *GetOrgauthorizationTrusteeGroupRolesParams {
	return &GetOrgauthorizationTrusteeGroupRolesParams{
		timeout: timeout,
	}
}

// NewGetOrgauthorizationTrusteeGroupRolesParamsWithContext creates a new GetOrgauthorizationTrusteeGroupRolesParams object
// with the ability to set a context for a request.
func NewGetOrgauthorizationTrusteeGroupRolesParamsWithContext(ctx context.Context) *GetOrgauthorizationTrusteeGroupRolesParams {
	return &GetOrgauthorizationTrusteeGroupRolesParams{
		Context: ctx,
	}
}

// NewGetOrgauthorizationTrusteeGroupRolesParamsWithHTTPClient creates a new GetOrgauthorizationTrusteeGroupRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrgauthorizationTrusteeGroupRolesParamsWithHTTPClient(client *http.Client) *GetOrgauthorizationTrusteeGroupRolesParams {
	return &GetOrgauthorizationTrusteeGroupRolesParams{
		HTTPClient: client,
	}
}

/*
GetOrgauthorizationTrusteeGroupRolesParams contains all the parameters to send to the API endpoint

	for the get orgauthorization trustee group roles operation.

	Typically these are written to a http.Request.
*/
type GetOrgauthorizationTrusteeGroupRolesParams struct {

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

// WithDefaults hydrates default values in the get orgauthorization trustee group roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgauthorizationTrusteeGroupRolesParams) WithDefaults() *GetOrgauthorizationTrusteeGroupRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get orgauthorization trustee group roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgauthorizationTrusteeGroupRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get orgauthorization trustee group roles params
func (o *GetOrgauthorizationTrusteeGroupRolesParams) WithTimeout(timeout time.Duration) *GetOrgauthorizationTrusteeGroupRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orgauthorization trustee group roles params
func (o *GetOrgauthorizationTrusteeGroupRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orgauthorization trustee group roles params
func (o *GetOrgauthorizationTrusteeGroupRolesParams) WithContext(ctx context.Context) *GetOrgauthorizationTrusteeGroupRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orgauthorization trustee group roles params
func (o *GetOrgauthorizationTrusteeGroupRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orgauthorization trustee group roles params
func (o *GetOrgauthorizationTrusteeGroupRolesParams) WithHTTPClient(client *http.Client) *GetOrgauthorizationTrusteeGroupRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orgauthorization trustee group roles params
func (o *GetOrgauthorizationTrusteeGroupRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrusteeGroupID adds the trusteeGroupID to the get orgauthorization trustee group roles params
func (o *GetOrgauthorizationTrusteeGroupRolesParams) WithTrusteeGroupID(trusteeGroupID string) *GetOrgauthorizationTrusteeGroupRolesParams {
	o.SetTrusteeGroupID(trusteeGroupID)
	return o
}

// SetTrusteeGroupID adds the trusteeGroupId to the get orgauthorization trustee group roles params
func (o *GetOrgauthorizationTrusteeGroupRolesParams) SetTrusteeGroupID(trusteeGroupID string) {
	o.TrusteeGroupID = trusteeGroupID
}

// WithTrusteeOrgID adds the trusteeOrgID to the get orgauthorization trustee group roles params
func (o *GetOrgauthorizationTrusteeGroupRolesParams) WithTrusteeOrgID(trusteeOrgID string) *GetOrgauthorizationTrusteeGroupRolesParams {
	o.SetTrusteeOrgID(trusteeOrgID)
	return o
}

// SetTrusteeOrgID adds the trusteeOrgId to the get orgauthorization trustee group roles params
func (o *GetOrgauthorizationTrusteeGroupRolesParams) SetTrusteeOrgID(trusteeOrgID string) {
	o.TrusteeOrgID = trusteeOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgauthorizationTrusteeGroupRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
