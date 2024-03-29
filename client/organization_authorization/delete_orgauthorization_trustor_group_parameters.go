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

// NewDeleteOrgauthorizationTrustorGroupParams creates a new DeleteOrgauthorizationTrustorGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrgauthorizationTrustorGroupParams() *DeleteOrgauthorizationTrustorGroupParams {
	return &DeleteOrgauthorizationTrustorGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrgauthorizationTrustorGroupParamsWithTimeout creates a new DeleteOrgauthorizationTrustorGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteOrgauthorizationTrustorGroupParamsWithTimeout(timeout time.Duration) *DeleteOrgauthorizationTrustorGroupParams {
	return &DeleteOrgauthorizationTrustorGroupParams{
		timeout: timeout,
	}
}

// NewDeleteOrgauthorizationTrustorGroupParamsWithContext creates a new DeleteOrgauthorizationTrustorGroupParams object
// with the ability to set a context for a request.
func NewDeleteOrgauthorizationTrustorGroupParamsWithContext(ctx context.Context) *DeleteOrgauthorizationTrustorGroupParams {
	return &DeleteOrgauthorizationTrustorGroupParams{
		Context: ctx,
	}
}

// NewDeleteOrgauthorizationTrustorGroupParamsWithHTTPClient creates a new DeleteOrgauthorizationTrustorGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrgauthorizationTrustorGroupParamsWithHTTPClient(client *http.Client) *DeleteOrgauthorizationTrustorGroupParams {
	return &DeleteOrgauthorizationTrustorGroupParams{
		HTTPClient: client,
	}
}

/*
DeleteOrgauthorizationTrustorGroupParams contains all the parameters to send to the API endpoint

	for the delete orgauthorization trustor group operation.

	Typically these are written to a http.Request.
*/
type DeleteOrgauthorizationTrustorGroupParams struct {

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

// WithDefaults hydrates default values in the delete orgauthorization trustor group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgauthorizationTrustorGroupParams) WithDefaults() *DeleteOrgauthorizationTrustorGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete orgauthorization trustor group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgauthorizationTrustorGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete orgauthorization trustor group params
func (o *DeleteOrgauthorizationTrustorGroupParams) WithTimeout(timeout time.Duration) *DeleteOrgauthorizationTrustorGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete orgauthorization trustor group params
func (o *DeleteOrgauthorizationTrustorGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete orgauthorization trustor group params
func (o *DeleteOrgauthorizationTrustorGroupParams) WithContext(ctx context.Context) *DeleteOrgauthorizationTrustorGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete orgauthorization trustor group params
func (o *DeleteOrgauthorizationTrustorGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete orgauthorization trustor group params
func (o *DeleteOrgauthorizationTrustorGroupParams) WithHTTPClient(client *http.Client) *DeleteOrgauthorizationTrustorGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete orgauthorization trustor group params
func (o *DeleteOrgauthorizationTrustorGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrustorGroupID adds the trustorGroupID to the delete orgauthorization trustor group params
func (o *DeleteOrgauthorizationTrustorGroupParams) WithTrustorGroupID(trustorGroupID string) *DeleteOrgauthorizationTrustorGroupParams {
	o.SetTrustorGroupID(trustorGroupID)
	return o
}

// SetTrustorGroupID adds the trustorGroupId to the delete orgauthorization trustor group params
func (o *DeleteOrgauthorizationTrustorGroupParams) SetTrustorGroupID(trustorGroupID string) {
	o.TrustorGroupID = trustorGroupID
}

// WithTrustorOrgID adds the trustorOrgID to the delete orgauthorization trustor group params
func (o *DeleteOrgauthorizationTrustorGroupParams) WithTrustorOrgID(trustorOrgID string) *DeleteOrgauthorizationTrustorGroupParams {
	o.SetTrustorOrgID(trustorOrgID)
	return o
}

// SetTrustorOrgID adds the trustorOrgId to the delete orgauthorization trustor group params
func (o *DeleteOrgauthorizationTrustorGroupParams) SetTrustorOrgID(trustorOrgID string) {
	o.TrustorOrgID = trustorOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrgauthorizationTrustorGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
