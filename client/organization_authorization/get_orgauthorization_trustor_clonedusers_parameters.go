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

// NewGetOrgauthorizationTrustorClonedusersParams creates a new GetOrgauthorizationTrustorClonedusersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrgauthorizationTrustorClonedusersParams() *GetOrgauthorizationTrustorClonedusersParams {
	return &GetOrgauthorizationTrustorClonedusersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgauthorizationTrustorClonedusersParamsWithTimeout creates a new GetOrgauthorizationTrustorClonedusersParams object
// with the ability to set a timeout on a request.
func NewGetOrgauthorizationTrustorClonedusersParamsWithTimeout(timeout time.Duration) *GetOrgauthorizationTrustorClonedusersParams {
	return &GetOrgauthorizationTrustorClonedusersParams{
		timeout: timeout,
	}
}

// NewGetOrgauthorizationTrustorClonedusersParamsWithContext creates a new GetOrgauthorizationTrustorClonedusersParams object
// with the ability to set a context for a request.
func NewGetOrgauthorizationTrustorClonedusersParamsWithContext(ctx context.Context) *GetOrgauthorizationTrustorClonedusersParams {
	return &GetOrgauthorizationTrustorClonedusersParams{
		Context: ctx,
	}
}

// NewGetOrgauthorizationTrustorClonedusersParamsWithHTTPClient creates a new GetOrgauthorizationTrustorClonedusersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrgauthorizationTrustorClonedusersParamsWithHTTPClient(client *http.Client) *GetOrgauthorizationTrustorClonedusersParams {
	return &GetOrgauthorizationTrustorClonedusersParams{
		HTTPClient: client,
	}
}

/*
GetOrgauthorizationTrustorClonedusersParams contains all the parameters to send to the API endpoint

	for the get orgauthorization trustor clonedusers operation.

	Typically these are written to a http.Request.
*/
type GetOrgauthorizationTrustorClonedusersParams struct {

	/* TrustorOrgID.

	   Trustor Organization Id
	*/
	TrustorOrgID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get orgauthorization trustor clonedusers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgauthorizationTrustorClonedusersParams) WithDefaults() *GetOrgauthorizationTrustorClonedusersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get orgauthorization trustor clonedusers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgauthorizationTrustorClonedusersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get orgauthorization trustor clonedusers params
func (o *GetOrgauthorizationTrustorClonedusersParams) WithTimeout(timeout time.Duration) *GetOrgauthorizationTrustorClonedusersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orgauthorization trustor clonedusers params
func (o *GetOrgauthorizationTrustorClonedusersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orgauthorization trustor clonedusers params
func (o *GetOrgauthorizationTrustorClonedusersParams) WithContext(ctx context.Context) *GetOrgauthorizationTrustorClonedusersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orgauthorization trustor clonedusers params
func (o *GetOrgauthorizationTrustorClonedusersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orgauthorization trustor clonedusers params
func (o *GetOrgauthorizationTrustorClonedusersParams) WithHTTPClient(client *http.Client) *GetOrgauthorizationTrustorClonedusersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orgauthorization trustor clonedusers params
func (o *GetOrgauthorizationTrustorClonedusersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrustorOrgID adds the trustorOrgID to the get orgauthorization trustor clonedusers params
func (o *GetOrgauthorizationTrustorClonedusersParams) WithTrustorOrgID(trustorOrgID string) *GetOrgauthorizationTrustorClonedusersParams {
	o.SetTrustorOrgID(trustorOrgID)
	return o
}

// SetTrustorOrgID adds the trustorOrgId to the get orgauthorization trustor clonedusers params
func (o *GetOrgauthorizationTrustorClonedusersParams) SetTrustorOrgID(trustorOrgID string) {
	o.TrustorOrgID = trustorOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgauthorizationTrustorClonedusersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param trustorOrgId
	if err := r.SetPathParam("trustorOrgId", o.TrustorOrgID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
