// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewPatchRoutingEmailDomainValidateParams creates a new PatchRoutingEmailDomainValidateParams object
// with the default values initialized.
func NewPatchRoutingEmailDomainValidateParams() *PatchRoutingEmailDomainValidateParams {
	var ()
	return &PatchRoutingEmailDomainValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchRoutingEmailDomainValidateParamsWithTimeout creates a new PatchRoutingEmailDomainValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchRoutingEmailDomainValidateParamsWithTimeout(timeout time.Duration) *PatchRoutingEmailDomainValidateParams {
	var ()
	return &PatchRoutingEmailDomainValidateParams{

		timeout: timeout,
	}
}

// NewPatchRoutingEmailDomainValidateParamsWithContext creates a new PatchRoutingEmailDomainValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchRoutingEmailDomainValidateParamsWithContext(ctx context.Context) *PatchRoutingEmailDomainValidateParams {
	var ()
	return &PatchRoutingEmailDomainValidateParams{

		Context: ctx,
	}
}

// NewPatchRoutingEmailDomainValidateParamsWithHTTPClient creates a new PatchRoutingEmailDomainValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchRoutingEmailDomainValidateParamsWithHTTPClient(client *http.Client) *PatchRoutingEmailDomainValidateParams {
	var ()
	return &PatchRoutingEmailDomainValidateParams{
		HTTPClient: client,
	}
}

/*PatchRoutingEmailDomainValidateParams contains all the parameters to send to the API endpoint
for the patch routing email domain validate operation typically these are written to a http.Request
*/
type PatchRoutingEmailDomainValidateParams struct {

	/*Body
	  Domain settings

	*/
	Body *models.InboundDomainPatchRequest
	/*DomainID
	  domain ID

	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch routing email domain validate params
func (o *PatchRoutingEmailDomainValidateParams) WithTimeout(timeout time.Duration) *PatchRoutingEmailDomainValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch routing email domain validate params
func (o *PatchRoutingEmailDomainValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch routing email domain validate params
func (o *PatchRoutingEmailDomainValidateParams) WithContext(ctx context.Context) *PatchRoutingEmailDomainValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch routing email domain validate params
func (o *PatchRoutingEmailDomainValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch routing email domain validate params
func (o *PatchRoutingEmailDomainValidateParams) WithHTTPClient(client *http.Client) *PatchRoutingEmailDomainValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch routing email domain validate params
func (o *PatchRoutingEmailDomainValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch routing email domain validate params
func (o *PatchRoutingEmailDomainValidateParams) WithBody(body *models.InboundDomainPatchRequest) *PatchRoutingEmailDomainValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch routing email domain validate params
func (o *PatchRoutingEmailDomainValidateParams) SetBody(body *models.InboundDomainPatchRequest) {
	o.Body = body
}

// WithDomainID adds the domainID to the patch routing email domain validate params
func (o *PatchRoutingEmailDomainValidateParams) WithDomainID(domainID string) *PatchRoutingEmailDomainValidateParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the patch routing email domain validate params
func (o *PatchRoutingEmailDomainValidateParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchRoutingEmailDomainValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}