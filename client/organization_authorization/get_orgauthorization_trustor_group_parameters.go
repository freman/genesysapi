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

// NewGetOrgauthorizationTrustorGroupParams creates a new GetOrgauthorizationTrustorGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrgauthorizationTrustorGroupParams() *GetOrgauthorizationTrustorGroupParams {
	return &GetOrgauthorizationTrustorGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgauthorizationTrustorGroupParamsWithTimeout creates a new GetOrgauthorizationTrustorGroupParams object
// with the ability to set a timeout on a request.
func NewGetOrgauthorizationTrustorGroupParamsWithTimeout(timeout time.Duration) *GetOrgauthorizationTrustorGroupParams {
	return &GetOrgauthorizationTrustorGroupParams{
		timeout: timeout,
	}
}

// NewGetOrgauthorizationTrustorGroupParamsWithContext creates a new GetOrgauthorizationTrustorGroupParams object
// with the ability to set a context for a request.
func NewGetOrgauthorizationTrustorGroupParamsWithContext(ctx context.Context) *GetOrgauthorizationTrustorGroupParams {
	return &GetOrgauthorizationTrustorGroupParams{
		Context: ctx,
	}
}

// NewGetOrgauthorizationTrustorGroupParamsWithHTTPClient creates a new GetOrgauthorizationTrustorGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrgauthorizationTrustorGroupParamsWithHTTPClient(client *http.Client) *GetOrgauthorizationTrustorGroupParams {
	return &GetOrgauthorizationTrustorGroupParams{
		HTTPClient: client,
	}
}

/*
GetOrgauthorizationTrustorGroupParams contains all the parameters to send to the API endpoint

	for the get orgauthorization trustor group operation.

	Typically these are written to a http.Request.
*/
type GetOrgauthorizationTrustorGroupParams struct {

	/* TrustorGroupID.

	   Trustor Group Id
	*/
	TrustorGroupID string

	/* TrustorOrgID.

	   Trustor Organization Id
	*/
	TrustorOrgID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get orgauthorization trustor group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgauthorizationTrustorGroupParams) WithDefaults() *GetOrgauthorizationTrustorGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get orgauthorization trustor group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgauthorizationTrustorGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get orgauthorization trustor group params
func (o *GetOrgauthorizationTrustorGroupParams) WithTimeout(timeout time.Duration) *GetOrgauthorizationTrustorGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orgauthorization trustor group params
func (o *GetOrgauthorizationTrustorGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orgauthorization trustor group params
func (o *GetOrgauthorizationTrustorGroupParams) WithContext(ctx context.Context) *GetOrgauthorizationTrustorGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orgauthorization trustor group params
func (o *GetOrgauthorizationTrustorGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orgauthorization trustor group params
func (o *GetOrgauthorizationTrustorGroupParams) WithHTTPClient(client *http.Client) *GetOrgauthorizationTrustorGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orgauthorization trustor group params
func (o *GetOrgauthorizationTrustorGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrustorGroupID adds the trustorGroupID to the get orgauthorization trustor group params
func (o *GetOrgauthorizationTrustorGroupParams) WithTrustorGroupID(trustorGroupID string) *GetOrgauthorizationTrustorGroupParams {
	o.SetTrustorGroupID(trustorGroupID)
	return o
}

// SetTrustorGroupID adds the trustorGroupId to the get orgauthorization trustor group params
func (o *GetOrgauthorizationTrustorGroupParams) SetTrustorGroupID(trustorGroupID string) {
	o.TrustorGroupID = trustorGroupID
}

// WithTrustorOrgID adds the trustorOrgID to the get orgauthorization trustor group params
func (o *GetOrgauthorizationTrustorGroupParams) WithTrustorOrgID(trustorOrgID string) *GetOrgauthorizationTrustorGroupParams {
	o.SetTrustorOrgID(trustorOrgID)
	return o
}

// SetTrustorOrgID adds the trustorOrgId to the get orgauthorization trustor group params
func (o *GetOrgauthorizationTrustorGroupParams) SetTrustorOrgID(trustorOrgID string) {
	o.TrustorOrgID = trustorOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgauthorizationTrustorGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param trustorGroupId
	if err := r.SetPathParam("trustorGroupId", o.TrustorGroupID); err != nil {
		return err
	}

	// path param trustorOrgId
	if err := r.SetPathParam("trustorOrgId", o.TrustorOrgID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}