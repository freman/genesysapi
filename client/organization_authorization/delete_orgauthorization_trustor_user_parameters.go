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

// NewDeleteOrgauthorizationTrustorUserParams creates a new DeleteOrgauthorizationTrustorUserParams object
// with the default values initialized.
func NewDeleteOrgauthorizationTrustorUserParams() *DeleteOrgauthorizationTrustorUserParams {
	var ()
	return &DeleteOrgauthorizationTrustorUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrgauthorizationTrustorUserParamsWithTimeout creates a new DeleteOrgauthorizationTrustorUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOrgauthorizationTrustorUserParamsWithTimeout(timeout time.Duration) *DeleteOrgauthorizationTrustorUserParams {
	var ()
	return &DeleteOrgauthorizationTrustorUserParams{

		timeout: timeout,
	}
}

// NewDeleteOrgauthorizationTrustorUserParamsWithContext creates a new DeleteOrgauthorizationTrustorUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOrgauthorizationTrustorUserParamsWithContext(ctx context.Context) *DeleteOrgauthorizationTrustorUserParams {
	var ()
	return &DeleteOrgauthorizationTrustorUserParams{

		Context: ctx,
	}
}

// NewDeleteOrgauthorizationTrustorUserParamsWithHTTPClient creates a new DeleteOrgauthorizationTrustorUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOrgauthorizationTrustorUserParamsWithHTTPClient(client *http.Client) *DeleteOrgauthorizationTrustorUserParams {
	var ()
	return &DeleteOrgauthorizationTrustorUserParams{
		HTTPClient: client,
	}
}

/*DeleteOrgauthorizationTrustorUserParams contains all the parameters to send to the API endpoint
for the delete orgauthorization trustor user operation typically these are written to a http.Request
*/
type DeleteOrgauthorizationTrustorUserParams struct {

	/*TrusteeUserID
	  Trustee User Id

	*/
	TrusteeUserID string
	/*TrustorOrgID
	  Trustor Organization Id

	*/
	TrustorOrgID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete orgauthorization trustor user params
func (o *DeleteOrgauthorizationTrustorUserParams) WithTimeout(timeout time.Duration) *DeleteOrgauthorizationTrustorUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete orgauthorization trustor user params
func (o *DeleteOrgauthorizationTrustorUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete orgauthorization trustor user params
func (o *DeleteOrgauthorizationTrustorUserParams) WithContext(ctx context.Context) *DeleteOrgauthorizationTrustorUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete orgauthorization trustor user params
func (o *DeleteOrgauthorizationTrustorUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete orgauthorization trustor user params
func (o *DeleteOrgauthorizationTrustorUserParams) WithHTTPClient(client *http.Client) *DeleteOrgauthorizationTrustorUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete orgauthorization trustor user params
func (o *DeleteOrgauthorizationTrustorUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrusteeUserID adds the trusteeUserID to the delete orgauthorization trustor user params
func (o *DeleteOrgauthorizationTrustorUserParams) WithTrusteeUserID(trusteeUserID string) *DeleteOrgauthorizationTrustorUserParams {
	o.SetTrusteeUserID(trusteeUserID)
	return o
}

// SetTrusteeUserID adds the trusteeUserId to the delete orgauthorization trustor user params
func (o *DeleteOrgauthorizationTrustorUserParams) SetTrusteeUserID(trusteeUserID string) {
	o.TrusteeUserID = trusteeUserID
}

// WithTrustorOrgID adds the trustorOrgID to the delete orgauthorization trustor user params
func (o *DeleteOrgauthorizationTrustorUserParams) WithTrustorOrgID(trustorOrgID string) *DeleteOrgauthorizationTrustorUserParams {
	o.SetTrustorOrgID(trustorOrgID)
	return o
}

// SetTrustorOrgID adds the trustorOrgId to the delete orgauthorization trustor user params
func (o *DeleteOrgauthorizationTrustorUserParams) SetTrustorOrgID(trustorOrgID string) {
	o.TrustorOrgID = trustorOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrgauthorizationTrustorUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param trusteeUserId
	if err := r.SetPathParam("trusteeUserId", o.TrusteeUserID); err != nil {
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
