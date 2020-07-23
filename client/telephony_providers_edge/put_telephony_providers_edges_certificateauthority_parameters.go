// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

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

// NewPutTelephonyProvidersEdgesCertificateauthorityParams creates a new PutTelephonyProvidersEdgesCertificateauthorityParams object
// with the default values initialized.
func NewPutTelephonyProvidersEdgesCertificateauthorityParams() *PutTelephonyProvidersEdgesCertificateauthorityParams {
	var ()
	return &PutTelephonyProvidersEdgesCertificateauthorityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTelephonyProvidersEdgesCertificateauthorityParamsWithTimeout creates a new PutTelephonyProvidersEdgesCertificateauthorityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTelephonyProvidersEdgesCertificateauthorityParamsWithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesCertificateauthorityParams {
	var ()
	return &PutTelephonyProvidersEdgesCertificateauthorityParams{

		timeout: timeout,
	}
}

// NewPutTelephonyProvidersEdgesCertificateauthorityParamsWithContext creates a new PutTelephonyProvidersEdgesCertificateauthorityParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTelephonyProvidersEdgesCertificateauthorityParamsWithContext(ctx context.Context) *PutTelephonyProvidersEdgesCertificateauthorityParams {
	var ()
	return &PutTelephonyProvidersEdgesCertificateauthorityParams{

		Context: ctx,
	}
}

// NewPutTelephonyProvidersEdgesCertificateauthorityParamsWithHTTPClient creates a new PutTelephonyProvidersEdgesCertificateauthorityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTelephonyProvidersEdgesCertificateauthorityParamsWithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesCertificateauthorityParams {
	var ()
	return &PutTelephonyProvidersEdgesCertificateauthorityParams{
		HTTPClient: client,
	}
}

/*PutTelephonyProvidersEdgesCertificateauthorityParams contains all the parameters to send to the API endpoint
for the put telephony providers edges certificateauthority operation typically these are written to a http.Request
*/
type PutTelephonyProvidersEdgesCertificateauthorityParams struct {

	/*Body
	  Certificate authority

	*/
	Body *models.DomainCertificateAuthority
	/*CertificateID
	  Certificate ID

	*/
	CertificateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put telephony providers edges certificateauthority params
func (o *PutTelephonyProvidersEdgesCertificateauthorityParams) WithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesCertificateauthorityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put telephony providers edges certificateauthority params
func (o *PutTelephonyProvidersEdgesCertificateauthorityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put telephony providers edges certificateauthority params
func (o *PutTelephonyProvidersEdgesCertificateauthorityParams) WithContext(ctx context.Context) *PutTelephonyProvidersEdgesCertificateauthorityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put telephony providers edges certificateauthority params
func (o *PutTelephonyProvidersEdgesCertificateauthorityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put telephony providers edges certificateauthority params
func (o *PutTelephonyProvidersEdgesCertificateauthorityParams) WithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesCertificateauthorityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put telephony providers edges certificateauthority params
func (o *PutTelephonyProvidersEdgesCertificateauthorityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put telephony providers edges certificateauthority params
func (o *PutTelephonyProvidersEdgesCertificateauthorityParams) WithBody(body *models.DomainCertificateAuthority) *PutTelephonyProvidersEdgesCertificateauthorityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put telephony providers edges certificateauthority params
func (o *PutTelephonyProvidersEdgesCertificateauthorityParams) SetBody(body *models.DomainCertificateAuthority) {
	o.Body = body
}

// WithCertificateID adds the certificateID to the put telephony providers edges certificateauthority params
func (o *PutTelephonyProvidersEdgesCertificateauthorityParams) WithCertificateID(certificateID string) *PutTelephonyProvidersEdgesCertificateauthorityParams {
	o.SetCertificateID(certificateID)
	return o
}

// SetCertificateID adds the certificateId to the put telephony providers edges certificateauthority params
func (o *PutTelephonyProvidersEdgesCertificateauthorityParams) SetCertificateID(certificateID string) {
	o.CertificateID = certificateID
}

// WriteToRequest writes these params to a swagger request
func (o *PutTelephonyProvidersEdgesCertificateauthorityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param certificateId
	if err := r.SetPathParam("certificateId", o.CertificateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
