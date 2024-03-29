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

// NewGetOrgauthorizationTrusteeClonedusersParams creates a new GetOrgauthorizationTrusteeClonedusersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrgauthorizationTrusteeClonedusersParams() *GetOrgauthorizationTrusteeClonedusersParams {
	return &GetOrgauthorizationTrusteeClonedusersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgauthorizationTrusteeClonedusersParamsWithTimeout creates a new GetOrgauthorizationTrusteeClonedusersParams object
// with the ability to set a timeout on a request.
func NewGetOrgauthorizationTrusteeClonedusersParamsWithTimeout(timeout time.Duration) *GetOrgauthorizationTrusteeClonedusersParams {
	return &GetOrgauthorizationTrusteeClonedusersParams{
		timeout: timeout,
	}
}

// NewGetOrgauthorizationTrusteeClonedusersParamsWithContext creates a new GetOrgauthorizationTrusteeClonedusersParams object
// with the ability to set a context for a request.
func NewGetOrgauthorizationTrusteeClonedusersParamsWithContext(ctx context.Context) *GetOrgauthorizationTrusteeClonedusersParams {
	return &GetOrgauthorizationTrusteeClonedusersParams{
		Context: ctx,
	}
}

// NewGetOrgauthorizationTrusteeClonedusersParamsWithHTTPClient creates a new GetOrgauthorizationTrusteeClonedusersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrgauthorizationTrusteeClonedusersParamsWithHTTPClient(client *http.Client) *GetOrgauthorizationTrusteeClonedusersParams {
	return &GetOrgauthorizationTrusteeClonedusersParams{
		HTTPClient: client,
	}
}

/*
GetOrgauthorizationTrusteeClonedusersParams contains all the parameters to send to the API endpoint

	for the get orgauthorization trustee clonedusers operation.

	Typically these are written to a http.Request.
*/
type GetOrgauthorizationTrusteeClonedusersParams struct {

	/* TrusteeOrgID.

	   Trustee Organization Id
	*/
	TrusteeOrgID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get orgauthorization trustee clonedusers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgauthorizationTrusteeClonedusersParams) WithDefaults() *GetOrgauthorizationTrusteeClonedusersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get orgauthorization trustee clonedusers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgauthorizationTrusteeClonedusersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get orgauthorization trustee clonedusers params
func (o *GetOrgauthorizationTrusteeClonedusersParams) WithTimeout(timeout time.Duration) *GetOrgauthorizationTrusteeClonedusersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orgauthorization trustee clonedusers params
func (o *GetOrgauthorizationTrusteeClonedusersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orgauthorization trustee clonedusers params
func (o *GetOrgauthorizationTrusteeClonedusersParams) WithContext(ctx context.Context) *GetOrgauthorizationTrusteeClonedusersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orgauthorization trustee clonedusers params
func (o *GetOrgauthorizationTrusteeClonedusersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orgauthorization trustee clonedusers params
func (o *GetOrgauthorizationTrusteeClonedusersParams) WithHTTPClient(client *http.Client) *GetOrgauthorizationTrusteeClonedusersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orgauthorization trustee clonedusers params
func (o *GetOrgauthorizationTrusteeClonedusersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrusteeOrgID adds the trusteeOrgID to the get orgauthorization trustee clonedusers params
func (o *GetOrgauthorizationTrusteeClonedusersParams) WithTrusteeOrgID(trusteeOrgID string) *GetOrgauthorizationTrusteeClonedusersParams {
	o.SetTrusteeOrgID(trusteeOrgID)
	return o
}

// SetTrusteeOrgID adds the trusteeOrgId to the get orgauthorization trustee clonedusers params
func (o *GetOrgauthorizationTrusteeClonedusersParams) SetTrusteeOrgID(trusteeOrgID string) {
	o.TrusteeOrgID = trusteeOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgauthorizationTrusteeClonedusersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param trusteeOrgId
	if err := r.SetPathParam("trusteeOrgId", o.TrusteeOrgID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
