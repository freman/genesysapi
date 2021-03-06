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

// NewGetTelephonyProvidersEdgesSiteOutboundrouteParams creates a new GetTelephonyProvidersEdgesSiteOutboundrouteParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesSiteOutboundrouteParams() *GetTelephonyProvidersEdgesSiteOutboundrouteParams {
	var ()
	return &GetTelephonyProvidersEdgesSiteOutboundrouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteParamsWithTimeout creates a new GetTelephonyProvidersEdgesSiteOutboundrouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesSiteOutboundrouteParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesSiteOutboundrouteParams {
	var ()
	return &GetTelephonyProvidersEdgesSiteOutboundrouteParams{

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteParamsWithContext creates a new GetTelephonyProvidersEdgesSiteOutboundrouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesSiteOutboundrouteParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesSiteOutboundrouteParams {
	var ()
	return &GetTelephonyProvidersEdgesSiteOutboundrouteParams{

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesSiteOutboundrouteParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesSiteOutboundrouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesSiteOutboundrouteParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesSiteOutboundrouteParams {
	var ()
	return &GetTelephonyProvidersEdgesSiteOutboundrouteParams{
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesSiteOutboundrouteParams contains all the parameters to send to the API endpoint
for the get telephony providers edges site outboundroute operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesSiteOutboundrouteParams struct {

	/*OutboundRouteID
	  Outbound route ID

	*/
	OutboundRouteID string
	/*SiteID
	  Site ID

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edges site outboundroute params
func (o *GetTelephonyProvidersEdgesSiteOutboundrouteParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesSiteOutboundrouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges site outboundroute params
func (o *GetTelephonyProvidersEdgesSiteOutboundrouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges site outboundroute params
func (o *GetTelephonyProvidersEdgesSiteOutboundrouteParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesSiteOutboundrouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges site outboundroute params
func (o *GetTelephonyProvidersEdgesSiteOutboundrouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges site outboundroute params
func (o *GetTelephonyProvidersEdgesSiteOutboundrouteParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesSiteOutboundrouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges site outboundroute params
func (o *GetTelephonyProvidersEdgesSiteOutboundrouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOutboundRouteID adds the outboundRouteID to the get telephony providers edges site outboundroute params
func (o *GetTelephonyProvidersEdgesSiteOutboundrouteParams) WithOutboundRouteID(outboundRouteID string) *GetTelephonyProvidersEdgesSiteOutboundrouteParams {
	o.SetOutboundRouteID(outboundRouteID)
	return o
}

// SetOutboundRouteID adds the outboundRouteId to the get telephony providers edges site outboundroute params
func (o *GetTelephonyProvidersEdgesSiteOutboundrouteParams) SetOutboundRouteID(outboundRouteID string) {
	o.OutboundRouteID = outboundRouteID
}

// WithSiteID adds the siteID to the get telephony providers edges site outboundroute params
func (o *GetTelephonyProvidersEdgesSiteOutboundrouteParams) WithSiteID(siteID string) *GetTelephonyProvidersEdgesSiteOutboundrouteParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get telephony providers edges site outboundroute params
func (o *GetTelephonyProvidersEdgesSiteOutboundrouteParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesSiteOutboundrouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
