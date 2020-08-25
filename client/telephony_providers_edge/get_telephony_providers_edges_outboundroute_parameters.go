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

// NewGetTelephonyProvidersEdgesOutboundrouteParams creates a new GetTelephonyProvidersEdgesOutboundrouteParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesOutboundrouteParams() *GetTelephonyProvidersEdgesOutboundrouteParams {
	var ()
	return &GetTelephonyProvidersEdgesOutboundrouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesOutboundrouteParamsWithTimeout creates a new GetTelephonyProvidersEdgesOutboundrouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesOutboundrouteParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesOutboundrouteParams {
	var ()
	return &GetTelephonyProvidersEdgesOutboundrouteParams{

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesOutboundrouteParamsWithContext creates a new GetTelephonyProvidersEdgesOutboundrouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesOutboundrouteParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesOutboundrouteParams {
	var ()
	return &GetTelephonyProvidersEdgesOutboundrouteParams{

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesOutboundrouteParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesOutboundrouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesOutboundrouteParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesOutboundrouteParams {
	var ()
	return &GetTelephonyProvidersEdgesOutboundrouteParams{
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesOutboundrouteParams contains all the parameters to send to the API endpoint
for the get telephony providers edges outboundroute operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesOutboundrouteParams struct {

	/*OutboundRouteID
	  Outbound route ID

	*/
	OutboundRouteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edges outboundroute params
func (o *GetTelephonyProvidersEdgesOutboundrouteParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesOutboundrouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges outboundroute params
func (o *GetTelephonyProvidersEdgesOutboundrouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges outboundroute params
func (o *GetTelephonyProvidersEdgesOutboundrouteParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesOutboundrouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges outboundroute params
func (o *GetTelephonyProvidersEdgesOutboundrouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges outboundroute params
func (o *GetTelephonyProvidersEdgesOutboundrouteParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesOutboundrouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges outboundroute params
func (o *GetTelephonyProvidersEdgesOutboundrouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOutboundRouteID adds the outboundRouteID to the get telephony providers edges outboundroute params
func (o *GetTelephonyProvidersEdgesOutboundrouteParams) WithOutboundRouteID(outboundRouteID string) *GetTelephonyProvidersEdgesOutboundrouteParams {
	o.SetOutboundRouteID(outboundRouteID)
	return o
}

// SetOutboundRouteID adds the outboundRouteId to the get telephony providers edges outboundroute params
func (o *GetTelephonyProvidersEdgesOutboundrouteParams) SetOutboundRouteID(outboundRouteID string) {
	o.OutboundRouteID = outboundRouteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesOutboundrouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param outboundRouteId
	if err := r.SetPathParam("outboundRouteId", o.OutboundRouteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}