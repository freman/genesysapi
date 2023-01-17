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
)

// NewGetTelephonyProvidersEdgesCertificateauthorityParams creates a new GetTelephonyProvidersEdgesCertificateauthorityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTelephonyProvidersEdgesCertificateauthorityParams() *GetTelephonyProvidersEdgesCertificateauthorityParams {
	return &GetTelephonyProvidersEdgesCertificateauthorityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesCertificateauthorityParamsWithTimeout creates a new GetTelephonyProvidersEdgesCertificateauthorityParams object
// with the ability to set a timeout on a request.
func NewGetTelephonyProvidersEdgesCertificateauthorityParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesCertificateauthorityParams {
	return &GetTelephonyProvidersEdgesCertificateauthorityParams{
		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesCertificateauthorityParamsWithContext creates a new GetTelephonyProvidersEdgesCertificateauthorityParams object
// with the ability to set a context for a request.
func NewGetTelephonyProvidersEdgesCertificateauthorityParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesCertificateauthorityParams {
	return &GetTelephonyProvidersEdgesCertificateauthorityParams{
		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesCertificateauthorityParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesCertificateauthorityParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTelephonyProvidersEdgesCertificateauthorityParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesCertificateauthorityParams {
	return &GetTelephonyProvidersEdgesCertificateauthorityParams{
		HTTPClient: client,
	}
}

/*
GetTelephonyProvidersEdgesCertificateauthorityParams contains all the parameters to send to the API endpoint

	for the get telephony providers edges certificateauthority operation.

	Typically these are written to a http.Request.
*/
type GetTelephonyProvidersEdgesCertificateauthorityParams struct {

	/* CertificateID.

	   Certificate ID
	*/
	CertificateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get telephony providers edges certificateauthority params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesCertificateauthorityParams) WithDefaults() *GetTelephonyProvidersEdgesCertificateauthorityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get telephony providers edges certificateauthority params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesCertificateauthorityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get telephony providers edges certificateauthority params
func (o *GetTelephonyProvidersEdgesCertificateauthorityParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesCertificateauthorityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges certificateauthority params
func (o *GetTelephonyProvidersEdgesCertificateauthorityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges certificateauthority params
func (o *GetTelephonyProvidersEdgesCertificateauthorityParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesCertificateauthorityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges certificateauthority params
func (o *GetTelephonyProvidersEdgesCertificateauthorityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges certificateauthority params
func (o *GetTelephonyProvidersEdgesCertificateauthorityParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesCertificateauthorityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges certificateauthority params
func (o *GetTelephonyProvidersEdgesCertificateauthorityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificateID adds the certificateID to the get telephony providers edges certificateauthority params
func (o *GetTelephonyProvidersEdgesCertificateauthorityParams) WithCertificateID(certificateID string) *GetTelephonyProvidersEdgesCertificateauthorityParams {
	o.SetCertificateID(certificateID)
	return o
}

// SetCertificateID adds the certificateId to the get telephony providers edges certificateauthority params
func (o *GetTelephonyProvidersEdgesCertificateauthorityParams) SetCertificateID(certificateID string) {
	o.CertificateID = certificateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesCertificateauthorityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param certificateId
	if err := r.SetPathParam("certificateId", o.CertificateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
