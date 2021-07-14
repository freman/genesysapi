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

	"github.com/freman/genesysapi/models"
)

// NewPutOrgauthorizationTrusteeParams creates a new PutOrgauthorizationTrusteeParams object
// with the default values initialized.
func NewPutOrgauthorizationTrusteeParams() *PutOrgauthorizationTrusteeParams {
	var ()
	return &PutOrgauthorizationTrusteeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrgauthorizationTrusteeParamsWithTimeout creates a new PutOrgauthorizationTrusteeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutOrgauthorizationTrusteeParamsWithTimeout(timeout time.Duration) *PutOrgauthorizationTrusteeParams {
	var ()
	return &PutOrgauthorizationTrusteeParams{

		timeout: timeout,
	}
}

// NewPutOrgauthorizationTrusteeParamsWithContext creates a new PutOrgauthorizationTrusteeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutOrgauthorizationTrusteeParamsWithContext(ctx context.Context) *PutOrgauthorizationTrusteeParams {
	var ()
	return &PutOrgauthorizationTrusteeParams{

		Context: ctx,
	}
}

// NewPutOrgauthorizationTrusteeParamsWithHTTPClient creates a new PutOrgauthorizationTrusteeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutOrgauthorizationTrusteeParamsWithHTTPClient(client *http.Client) *PutOrgauthorizationTrusteeParams {
	var ()
	return &PutOrgauthorizationTrusteeParams{
		HTTPClient: client,
	}
}

/*PutOrgauthorizationTrusteeParams contains all the parameters to send to the API endpoint
for the put orgauthorization trustee operation typically these are written to a http.Request
*/
type PutOrgauthorizationTrusteeParams struct {

	/*Body
	  Client

	*/
	Body *models.TrustUpdate
	/*TrusteeOrgID
	  Trustee Organization Id

	*/
	TrusteeOrgID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put orgauthorization trustee params
func (o *PutOrgauthorizationTrusteeParams) WithTimeout(timeout time.Duration) *PutOrgauthorizationTrusteeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put orgauthorization trustee params
func (o *PutOrgauthorizationTrusteeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put orgauthorization trustee params
func (o *PutOrgauthorizationTrusteeParams) WithContext(ctx context.Context) *PutOrgauthorizationTrusteeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put orgauthorization trustee params
func (o *PutOrgauthorizationTrusteeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put orgauthorization trustee params
func (o *PutOrgauthorizationTrusteeParams) WithHTTPClient(client *http.Client) *PutOrgauthorizationTrusteeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put orgauthorization trustee params
func (o *PutOrgauthorizationTrusteeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put orgauthorization trustee params
func (o *PutOrgauthorizationTrusteeParams) WithBody(body *models.TrustUpdate) *PutOrgauthorizationTrusteeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put orgauthorization trustee params
func (o *PutOrgauthorizationTrusteeParams) SetBody(body *models.TrustUpdate) {
	o.Body = body
}

// WithTrusteeOrgID adds the trusteeOrgID to the put orgauthorization trustee params
func (o *PutOrgauthorizationTrusteeParams) WithTrusteeOrgID(trusteeOrgID string) *PutOrgauthorizationTrusteeParams {
	o.SetTrusteeOrgID(trusteeOrgID)
	return o
}

// SetTrusteeOrgID adds the trusteeOrgId to the put orgauthorization trustee params
func (o *PutOrgauthorizationTrusteeParams) SetTrusteeOrgID(trusteeOrgID string) {
	o.TrusteeOrgID = trusteeOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrgauthorizationTrusteeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
