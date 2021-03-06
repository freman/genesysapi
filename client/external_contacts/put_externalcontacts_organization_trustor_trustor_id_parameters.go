// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

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

// NewPutExternalcontactsOrganizationTrustorTrustorIDParams creates a new PutExternalcontactsOrganizationTrustorTrustorIDParams object
// with the default values initialized.
func NewPutExternalcontactsOrganizationTrustorTrustorIDParams() *PutExternalcontactsOrganizationTrustorTrustorIDParams {
	var ()
	return &PutExternalcontactsOrganizationTrustorTrustorIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutExternalcontactsOrganizationTrustorTrustorIDParamsWithTimeout creates a new PutExternalcontactsOrganizationTrustorTrustorIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutExternalcontactsOrganizationTrustorTrustorIDParamsWithTimeout(timeout time.Duration) *PutExternalcontactsOrganizationTrustorTrustorIDParams {
	var ()
	return &PutExternalcontactsOrganizationTrustorTrustorIDParams{

		timeout: timeout,
	}
}

// NewPutExternalcontactsOrganizationTrustorTrustorIDParamsWithContext creates a new PutExternalcontactsOrganizationTrustorTrustorIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutExternalcontactsOrganizationTrustorTrustorIDParamsWithContext(ctx context.Context) *PutExternalcontactsOrganizationTrustorTrustorIDParams {
	var ()
	return &PutExternalcontactsOrganizationTrustorTrustorIDParams{

		Context: ctx,
	}
}

// NewPutExternalcontactsOrganizationTrustorTrustorIDParamsWithHTTPClient creates a new PutExternalcontactsOrganizationTrustorTrustorIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutExternalcontactsOrganizationTrustorTrustorIDParamsWithHTTPClient(client *http.Client) *PutExternalcontactsOrganizationTrustorTrustorIDParams {
	var ()
	return &PutExternalcontactsOrganizationTrustorTrustorIDParams{
		HTTPClient: client,
	}
}

/*PutExternalcontactsOrganizationTrustorTrustorIDParams contains all the parameters to send to the API endpoint
for the put externalcontacts organization trustor trustor Id operation typically these are written to a http.Request
*/
type PutExternalcontactsOrganizationTrustorTrustorIDParams struct {

	/*ExternalOrganizationID
	  External Organization ID

	*/
	ExternalOrganizationID string
	/*TrustorID
	  Trustor ID

	*/
	TrustorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put externalcontacts organization trustor trustor Id params
func (o *PutExternalcontactsOrganizationTrustorTrustorIDParams) WithTimeout(timeout time.Duration) *PutExternalcontactsOrganizationTrustorTrustorIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put externalcontacts organization trustor trustor Id params
func (o *PutExternalcontactsOrganizationTrustorTrustorIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put externalcontacts organization trustor trustor Id params
func (o *PutExternalcontactsOrganizationTrustorTrustorIDParams) WithContext(ctx context.Context) *PutExternalcontactsOrganizationTrustorTrustorIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put externalcontacts organization trustor trustor Id params
func (o *PutExternalcontactsOrganizationTrustorTrustorIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put externalcontacts organization trustor trustor Id params
func (o *PutExternalcontactsOrganizationTrustorTrustorIDParams) WithHTTPClient(client *http.Client) *PutExternalcontactsOrganizationTrustorTrustorIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put externalcontacts organization trustor trustor Id params
func (o *PutExternalcontactsOrganizationTrustorTrustorIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExternalOrganizationID adds the externalOrganizationID to the put externalcontacts organization trustor trustor Id params
func (o *PutExternalcontactsOrganizationTrustorTrustorIDParams) WithExternalOrganizationID(externalOrganizationID string) *PutExternalcontactsOrganizationTrustorTrustorIDParams {
	o.SetExternalOrganizationID(externalOrganizationID)
	return o
}

// SetExternalOrganizationID adds the externalOrganizationId to the put externalcontacts organization trustor trustor Id params
func (o *PutExternalcontactsOrganizationTrustorTrustorIDParams) SetExternalOrganizationID(externalOrganizationID string) {
	o.ExternalOrganizationID = externalOrganizationID
}

// WithTrustorID adds the trustorID to the put externalcontacts organization trustor trustor Id params
func (o *PutExternalcontactsOrganizationTrustorTrustorIDParams) WithTrustorID(trustorID string) *PutExternalcontactsOrganizationTrustorTrustorIDParams {
	o.SetTrustorID(trustorID)
	return o
}

// SetTrustorID adds the trustorId to the put externalcontacts organization trustor trustor Id params
func (o *PutExternalcontactsOrganizationTrustorTrustorIDParams) SetTrustorID(trustorID string) {
	o.TrustorID = trustorID
}

// WriteToRequest writes these params to a swagger request
func (o *PutExternalcontactsOrganizationTrustorTrustorIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param externalOrganizationId
	if err := r.SetPathParam("externalOrganizationId", o.ExternalOrganizationID); err != nil {
		return err
	}

	// path param trustorId
	if err := r.SetPathParam("trustorId", o.TrustorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
