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

// NewPutTelephonyProvidersEdgesSiteParams creates a new PutTelephonyProvidersEdgesSiteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutTelephonyProvidersEdgesSiteParams() *PutTelephonyProvidersEdgesSiteParams {
	return &PutTelephonyProvidersEdgesSiteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutTelephonyProvidersEdgesSiteParamsWithTimeout creates a new PutTelephonyProvidersEdgesSiteParams object
// with the ability to set a timeout on a request.
func NewPutTelephonyProvidersEdgesSiteParamsWithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesSiteParams {
	return &PutTelephonyProvidersEdgesSiteParams{
		timeout: timeout,
	}
}

// NewPutTelephonyProvidersEdgesSiteParamsWithContext creates a new PutTelephonyProvidersEdgesSiteParams object
// with the ability to set a context for a request.
func NewPutTelephonyProvidersEdgesSiteParamsWithContext(ctx context.Context) *PutTelephonyProvidersEdgesSiteParams {
	return &PutTelephonyProvidersEdgesSiteParams{
		Context: ctx,
	}
}

// NewPutTelephonyProvidersEdgesSiteParamsWithHTTPClient creates a new PutTelephonyProvidersEdgesSiteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutTelephonyProvidersEdgesSiteParamsWithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesSiteParams {
	return &PutTelephonyProvidersEdgesSiteParams{
		HTTPClient: client,
	}
}

/*
PutTelephonyProvidersEdgesSiteParams contains all the parameters to send to the API endpoint

	for the put telephony providers edges site operation.

	Typically these are written to a http.Request.
*/
type PutTelephonyProvidersEdgesSiteParams struct {

	/* Body.

	   Site
	*/
	Body *models.Site

	/* SiteID.

	   Site ID
	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put telephony providers edges site params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTelephonyProvidersEdgesSiteParams) WithDefaults() *PutTelephonyProvidersEdgesSiteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put telephony providers edges site params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTelephonyProvidersEdgesSiteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put telephony providers edges site params
func (o *PutTelephonyProvidersEdgesSiteParams) WithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesSiteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put telephony providers edges site params
func (o *PutTelephonyProvidersEdgesSiteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put telephony providers edges site params
func (o *PutTelephonyProvidersEdgesSiteParams) WithContext(ctx context.Context) *PutTelephonyProvidersEdgesSiteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put telephony providers edges site params
func (o *PutTelephonyProvidersEdgesSiteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put telephony providers edges site params
func (o *PutTelephonyProvidersEdgesSiteParams) WithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesSiteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put telephony providers edges site params
func (o *PutTelephonyProvidersEdgesSiteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put telephony providers edges site params
func (o *PutTelephonyProvidersEdgesSiteParams) WithBody(body *models.Site) *PutTelephonyProvidersEdgesSiteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put telephony providers edges site params
func (o *PutTelephonyProvidersEdgesSiteParams) SetBody(body *models.Site) {
	o.Body = body
}

// WithSiteID adds the siteID to the put telephony providers edges site params
func (o *PutTelephonyProvidersEdgesSiteParams) WithSiteID(siteID string) *PutTelephonyProvidersEdgesSiteParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the put telephony providers edges site params
func (o *PutTelephonyProvidersEdgesSiteParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PutTelephonyProvidersEdgesSiteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param siteId
	if err := r.SetPathParam("siteId", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
