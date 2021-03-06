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

// NewGetOrgauthorizationTrusteeParams creates a new GetOrgauthorizationTrusteeParams object
// with the default values initialized.
func NewGetOrgauthorizationTrusteeParams() *GetOrgauthorizationTrusteeParams {
	var ()
	return &GetOrgauthorizationTrusteeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgauthorizationTrusteeParamsWithTimeout creates a new GetOrgauthorizationTrusteeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrgauthorizationTrusteeParamsWithTimeout(timeout time.Duration) *GetOrgauthorizationTrusteeParams {
	var ()
	return &GetOrgauthorizationTrusteeParams{

		timeout: timeout,
	}
}

// NewGetOrgauthorizationTrusteeParamsWithContext creates a new GetOrgauthorizationTrusteeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrgauthorizationTrusteeParamsWithContext(ctx context.Context) *GetOrgauthorizationTrusteeParams {
	var ()
	return &GetOrgauthorizationTrusteeParams{

		Context: ctx,
	}
}

// NewGetOrgauthorizationTrusteeParamsWithHTTPClient creates a new GetOrgauthorizationTrusteeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrgauthorizationTrusteeParamsWithHTTPClient(client *http.Client) *GetOrgauthorizationTrusteeParams {
	var ()
	return &GetOrgauthorizationTrusteeParams{
		HTTPClient: client,
	}
}

/*GetOrgauthorizationTrusteeParams contains all the parameters to send to the API endpoint
for the get orgauthorization trustee operation typically these are written to a http.Request
*/
type GetOrgauthorizationTrusteeParams struct {

	/*TrusteeOrgID
	  Trustee Organization Id

	*/
	TrusteeOrgID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get orgauthorization trustee params
func (o *GetOrgauthorizationTrusteeParams) WithTimeout(timeout time.Duration) *GetOrgauthorizationTrusteeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orgauthorization trustee params
func (o *GetOrgauthorizationTrusteeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orgauthorization trustee params
func (o *GetOrgauthorizationTrusteeParams) WithContext(ctx context.Context) *GetOrgauthorizationTrusteeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orgauthorization trustee params
func (o *GetOrgauthorizationTrusteeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orgauthorization trustee params
func (o *GetOrgauthorizationTrusteeParams) WithHTTPClient(client *http.Client) *GetOrgauthorizationTrusteeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orgauthorization trustee params
func (o *GetOrgauthorizationTrusteeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrusteeOrgID adds the trusteeOrgID to the get orgauthorization trustee params
func (o *GetOrgauthorizationTrusteeParams) WithTrusteeOrgID(trusteeOrgID string) *GetOrgauthorizationTrusteeParams {
	o.SetTrusteeOrgID(trusteeOrgID)
	return o
}

// SetTrusteeOrgID adds the trusteeOrgId to the get orgauthorization trustee params
func (o *GetOrgauthorizationTrusteeParams) SetTrusteeOrgID(trusteeOrgID string) {
	o.TrusteeOrgID = trusteeOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgauthorizationTrusteeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
