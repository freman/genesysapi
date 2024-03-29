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

// NewPutTelephonyProvidersEdgesSiteOutboundrouteParams creates a new PutTelephonyProvidersEdgesSiteOutboundrouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutTelephonyProvidersEdgesSiteOutboundrouteParams() *PutTelephonyProvidersEdgesSiteOutboundrouteParams {
	return &PutTelephonyProvidersEdgesSiteOutboundrouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutTelephonyProvidersEdgesSiteOutboundrouteParamsWithTimeout creates a new PutTelephonyProvidersEdgesSiteOutboundrouteParams object
// with the ability to set a timeout on a request.
func NewPutTelephonyProvidersEdgesSiteOutboundrouteParamsWithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesSiteOutboundrouteParams {
	return &PutTelephonyProvidersEdgesSiteOutboundrouteParams{
		timeout: timeout,
	}
}

// NewPutTelephonyProvidersEdgesSiteOutboundrouteParamsWithContext creates a new PutTelephonyProvidersEdgesSiteOutboundrouteParams object
// with the ability to set a context for a request.
func NewPutTelephonyProvidersEdgesSiteOutboundrouteParamsWithContext(ctx context.Context) *PutTelephonyProvidersEdgesSiteOutboundrouteParams {
	return &PutTelephonyProvidersEdgesSiteOutboundrouteParams{
		Context: ctx,
	}
}

// NewPutTelephonyProvidersEdgesSiteOutboundrouteParamsWithHTTPClient creates a new PutTelephonyProvidersEdgesSiteOutboundrouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutTelephonyProvidersEdgesSiteOutboundrouteParamsWithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesSiteOutboundrouteParams {
	return &PutTelephonyProvidersEdgesSiteOutboundrouteParams{
		HTTPClient: client,
	}
}

/*
PutTelephonyProvidersEdgesSiteOutboundrouteParams contains all the parameters to send to the API endpoint

	for the put telephony providers edges site outboundroute operation.

	Typically these are written to a http.Request.
*/
type PutTelephonyProvidersEdgesSiteOutboundrouteParams struct {

	/* Body.

	   OutboundRoute
	*/
	Body *models.OutboundRouteBase

	/* OutboundRouteID.

	   Outbound route ID
	*/
	OutboundRouteID string

	/* SiteID.

	   Site ID
	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put telephony providers edges site outboundroute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) WithDefaults() *PutTelephonyProvidersEdgesSiteOutboundrouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put telephony providers edges site outboundroute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put telephony providers edges site outboundroute params
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) WithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesSiteOutboundrouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put telephony providers edges site outboundroute params
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put telephony providers edges site outboundroute params
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) WithContext(ctx context.Context) *PutTelephonyProvidersEdgesSiteOutboundrouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put telephony providers edges site outboundroute params
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put telephony providers edges site outboundroute params
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) WithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesSiteOutboundrouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put telephony providers edges site outboundroute params
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put telephony providers edges site outboundroute params
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) WithBody(body *models.OutboundRouteBase) *PutTelephonyProvidersEdgesSiteOutboundrouteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put telephony providers edges site outboundroute params
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) SetBody(body *models.OutboundRouteBase) {
	o.Body = body
}

// WithOutboundRouteID adds the outboundRouteID to the put telephony providers edges site outboundroute params
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) WithOutboundRouteID(outboundRouteID string) *PutTelephonyProvidersEdgesSiteOutboundrouteParams {
	o.SetOutboundRouteID(outboundRouteID)
	return o
}

// SetOutboundRouteID adds the outboundRouteId to the put telephony providers edges site outboundroute params
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) SetOutboundRouteID(outboundRouteID string) {
	o.OutboundRouteID = outboundRouteID
}

// WithSiteID adds the siteID to the put telephony providers edges site outboundroute params
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) WithSiteID(siteID string) *PutTelephonyProvidersEdgesSiteOutboundrouteParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the put telephony providers edges site outboundroute params
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PutTelephonyProvidersEdgesSiteOutboundrouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param outboundRouteId
	if err := r.SetPathParam("outboundRouteId", o.OutboundRouteID); err != nil {
		return err
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
