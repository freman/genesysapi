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

// NewDeleteOrgauthorizationTrusteeCloneduserParams creates a new DeleteOrgauthorizationTrusteeCloneduserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrgauthorizationTrusteeCloneduserParams() *DeleteOrgauthorizationTrusteeCloneduserParams {
	return &DeleteOrgauthorizationTrusteeCloneduserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrgauthorizationTrusteeCloneduserParamsWithTimeout creates a new DeleteOrgauthorizationTrusteeCloneduserParams object
// with the ability to set a timeout on a request.
func NewDeleteOrgauthorizationTrusteeCloneduserParamsWithTimeout(timeout time.Duration) *DeleteOrgauthorizationTrusteeCloneduserParams {
	return &DeleteOrgauthorizationTrusteeCloneduserParams{
		timeout: timeout,
	}
}

// NewDeleteOrgauthorizationTrusteeCloneduserParamsWithContext creates a new DeleteOrgauthorizationTrusteeCloneduserParams object
// with the ability to set a context for a request.
func NewDeleteOrgauthorizationTrusteeCloneduserParamsWithContext(ctx context.Context) *DeleteOrgauthorizationTrusteeCloneduserParams {
	return &DeleteOrgauthorizationTrusteeCloneduserParams{
		Context: ctx,
	}
}

// NewDeleteOrgauthorizationTrusteeCloneduserParamsWithHTTPClient creates a new DeleteOrgauthorizationTrusteeCloneduserParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrgauthorizationTrusteeCloneduserParamsWithHTTPClient(client *http.Client) *DeleteOrgauthorizationTrusteeCloneduserParams {
	return &DeleteOrgauthorizationTrusteeCloneduserParams{
		HTTPClient: client,
	}
}

/*
DeleteOrgauthorizationTrusteeCloneduserParams contains all the parameters to send to the API endpoint

	for the delete orgauthorization trustee cloneduser operation.

	Typically these are written to a http.Request.
*/
type DeleteOrgauthorizationTrusteeCloneduserParams struct {

	/* TrusteeOrgID.

	   Trustee Organization Id
	*/
	TrusteeOrgID string

	/* TrusteeUserID.

	   Id of the cloned user to delete
	*/
	TrusteeUserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete orgauthorization trustee cloneduser params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) WithDefaults() *DeleteOrgauthorizationTrusteeCloneduserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete orgauthorization trustee cloneduser params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete orgauthorization trustee cloneduser params
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) WithTimeout(timeout time.Duration) *DeleteOrgauthorizationTrusteeCloneduserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete orgauthorization trustee cloneduser params
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete orgauthorization trustee cloneduser params
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) WithContext(ctx context.Context) *DeleteOrgauthorizationTrusteeCloneduserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete orgauthorization trustee cloneduser params
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete orgauthorization trustee cloneduser params
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) WithHTTPClient(client *http.Client) *DeleteOrgauthorizationTrusteeCloneduserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete orgauthorization trustee cloneduser params
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrusteeOrgID adds the trusteeOrgID to the delete orgauthorization trustee cloneduser params
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) WithTrusteeOrgID(trusteeOrgID string) *DeleteOrgauthorizationTrusteeCloneduserParams {
	o.SetTrusteeOrgID(trusteeOrgID)
	return o
}

// SetTrusteeOrgID adds the trusteeOrgId to the delete orgauthorization trustee cloneduser params
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) SetTrusteeOrgID(trusteeOrgID string) {
	o.TrusteeOrgID = trusteeOrgID
}

// WithTrusteeUserID adds the trusteeUserID to the delete orgauthorization trustee cloneduser params
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) WithTrusteeUserID(trusteeUserID string) *DeleteOrgauthorizationTrusteeCloneduserParams {
	o.SetTrusteeUserID(trusteeUserID)
	return o
}

// SetTrusteeUserID adds the trusteeUserId to the delete orgauthorization trustee cloneduser params
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) SetTrusteeUserID(trusteeUserID string) {
	o.TrusteeUserID = trusteeUserID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrgauthorizationTrusteeCloneduserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param trusteeOrgId
	if err := r.SetPathParam("trusteeOrgId", o.TrusteeOrgID); err != nil {
		return err
	}

	// path param trusteeUserId
	if err := r.SetPathParam("trusteeUserId", o.TrusteeUserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
