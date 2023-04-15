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

// NewPutOrgauthorizationTrustorGroupParams creates a new PutOrgauthorizationTrustorGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutOrgauthorizationTrustorGroupParams() *PutOrgauthorizationTrustorGroupParams {
	return &PutOrgauthorizationTrustorGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrgauthorizationTrustorGroupParamsWithTimeout creates a new PutOrgauthorizationTrustorGroupParams object
// with the ability to set a timeout on a request.
func NewPutOrgauthorizationTrustorGroupParamsWithTimeout(timeout time.Duration) *PutOrgauthorizationTrustorGroupParams {
	return &PutOrgauthorizationTrustorGroupParams{
		timeout: timeout,
	}
}

// NewPutOrgauthorizationTrustorGroupParamsWithContext creates a new PutOrgauthorizationTrustorGroupParams object
// with the ability to set a context for a request.
func NewPutOrgauthorizationTrustorGroupParamsWithContext(ctx context.Context) *PutOrgauthorizationTrustorGroupParams {
	return &PutOrgauthorizationTrustorGroupParams{
		Context: ctx,
	}
}

// NewPutOrgauthorizationTrustorGroupParamsWithHTTPClient creates a new PutOrgauthorizationTrustorGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutOrgauthorizationTrustorGroupParamsWithHTTPClient(client *http.Client) *PutOrgauthorizationTrustorGroupParams {
	return &PutOrgauthorizationTrustorGroupParams{
		HTTPClient: client,
	}
}

/*
PutOrgauthorizationTrustorGroupParams contains all the parameters to send to the API endpoint

	for the put orgauthorization trustor group operation.

	Typically these are written to a http.Request.
*/
type PutOrgauthorizationTrustorGroupParams struct {

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

// WithDefaults hydrates default values in the put orgauthorization trustor group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrgauthorizationTrustorGroupParams) WithDefaults() *PutOrgauthorizationTrustorGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put orgauthorization trustor group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrgauthorizationTrustorGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put orgauthorization trustor group params
func (o *PutOrgauthorizationTrustorGroupParams) WithTimeout(timeout time.Duration) *PutOrgauthorizationTrustorGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put orgauthorization trustor group params
func (o *PutOrgauthorizationTrustorGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put orgauthorization trustor group params
func (o *PutOrgauthorizationTrustorGroupParams) WithContext(ctx context.Context) *PutOrgauthorizationTrustorGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put orgauthorization trustor group params
func (o *PutOrgauthorizationTrustorGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put orgauthorization trustor group params
func (o *PutOrgauthorizationTrustorGroupParams) WithHTTPClient(client *http.Client) *PutOrgauthorizationTrustorGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put orgauthorization trustor group params
func (o *PutOrgauthorizationTrustorGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrustorGroupID adds the trustorGroupID to the put orgauthorization trustor group params
func (o *PutOrgauthorizationTrustorGroupParams) WithTrustorGroupID(trustorGroupID string) *PutOrgauthorizationTrustorGroupParams {
	o.SetTrustorGroupID(trustorGroupID)
	return o
}

// SetTrustorGroupID adds the trustorGroupId to the put orgauthorization trustor group params
func (o *PutOrgauthorizationTrustorGroupParams) SetTrustorGroupID(trustorGroupID string) {
	o.TrustorGroupID = trustorGroupID
}

// WithTrustorOrgID adds the trustorOrgID to the put orgauthorization trustor group params
func (o *PutOrgauthorizationTrustorGroupParams) WithTrustorOrgID(trustorOrgID string) *PutOrgauthorizationTrustorGroupParams {
	o.SetTrustorOrgID(trustorOrgID)
	return o
}

// SetTrustorOrgID adds the trustorOrgId to the put orgauthorization trustor group params
func (o *PutOrgauthorizationTrustorGroupParams) SetTrustorOrgID(trustorOrgID string) {
	o.TrustorOrgID = trustorOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrgauthorizationTrustorGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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