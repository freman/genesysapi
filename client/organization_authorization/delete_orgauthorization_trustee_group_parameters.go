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

// NewDeleteOrgauthorizationTrusteeGroupParams creates a new DeleteOrgauthorizationTrusteeGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrgauthorizationTrusteeGroupParams() *DeleteOrgauthorizationTrusteeGroupParams {
	return &DeleteOrgauthorizationTrusteeGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrgauthorizationTrusteeGroupParamsWithTimeout creates a new DeleteOrgauthorizationTrusteeGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteOrgauthorizationTrusteeGroupParamsWithTimeout(timeout time.Duration) *DeleteOrgauthorizationTrusteeGroupParams {
	return &DeleteOrgauthorizationTrusteeGroupParams{
		timeout: timeout,
	}
}

// NewDeleteOrgauthorizationTrusteeGroupParamsWithContext creates a new DeleteOrgauthorizationTrusteeGroupParams object
// with the ability to set a context for a request.
func NewDeleteOrgauthorizationTrusteeGroupParamsWithContext(ctx context.Context) *DeleteOrgauthorizationTrusteeGroupParams {
	return &DeleteOrgauthorizationTrusteeGroupParams{
		Context: ctx,
	}
}

// NewDeleteOrgauthorizationTrusteeGroupParamsWithHTTPClient creates a new DeleteOrgauthorizationTrusteeGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrgauthorizationTrusteeGroupParamsWithHTTPClient(client *http.Client) *DeleteOrgauthorizationTrusteeGroupParams {
	return &DeleteOrgauthorizationTrusteeGroupParams{
		HTTPClient: client,
	}
}

/*
DeleteOrgauthorizationTrusteeGroupParams contains all the parameters to send to the API endpoint

	for the delete orgauthorization trustee group operation.

	Typically these are written to a http.Request.
*/
type DeleteOrgauthorizationTrusteeGroupParams struct {

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

// WithDefaults hydrates default values in the delete orgauthorization trustee group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgauthorizationTrusteeGroupParams) WithDefaults() *DeleteOrgauthorizationTrusteeGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete orgauthorization trustee group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgauthorizationTrusteeGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete orgauthorization trustee group params
func (o *DeleteOrgauthorizationTrusteeGroupParams) WithTimeout(timeout time.Duration) *DeleteOrgauthorizationTrusteeGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete orgauthorization trustee group params
func (o *DeleteOrgauthorizationTrusteeGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete orgauthorization trustee group params
func (o *DeleteOrgauthorizationTrusteeGroupParams) WithContext(ctx context.Context) *DeleteOrgauthorizationTrusteeGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete orgauthorization trustee group params
func (o *DeleteOrgauthorizationTrusteeGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete orgauthorization trustee group params
func (o *DeleteOrgauthorizationTrusteeGroupParams) WithHTTPClient(client *http.Client) *DeleteOrgauthorizationTrusteeGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete orgauthorization trustee group params
func (o *DeleteOrgauthorizationTrusteeGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrusteeGroupID adds the trusteeGroupID to the delete orgauthorization trustee group params
func (o *DeleteOrgauthorizationTrusteeGroupParams) WithTrusteeGroupID(trusteeGroupID string) *DeleteOrgauthorizationTrusteeGroupParams {
	o.SetTrusteeGroupID(trusteeGroupID)
	return o
}

// SetTrusteeGroupID adds the trusteeGroupId to the delete orgauthorization trustee group params
func (o *DeleteOrgauthorizationTrusteeGroupParams) SetTrusteeGroupID(trusteeGroupID string) {
	o.TrusteeGroupID = trusteeGroupID
}

// WithTrusteeOrgID adds the trusteeOrgID to the delete orgauthorization trustee group params
func (o *DeleteOrgauthorizationTrusteeGroupParams) WithTrusteeOrgID(trusteeOrgID string) *DeleteOrgauthorizationTrusteeGroupParams {
	o.SetTrusteeOrgID(trusteeOrgID)
	return o
}

// SetTrusteeOrgID adds the trusteeOrgId to the delete orgauthorization trustee group params
func (o *DeleteOrgauthorizationTrusteeGroupParams) SetTrusteeOrgID(trusteeOrgID string) {
	o.TrusteeOrgID = trusteeOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrgauthorizationTrusteeGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
