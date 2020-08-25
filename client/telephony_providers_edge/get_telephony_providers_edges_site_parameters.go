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

// NewGetTelephonyProvidersEdgesSiteParams creates a new GetTelephonyProvidersEdgesSiteParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesSiteParams() *GetTelephonyProvidersEdgesSiteParams {
	var ()
	return &GetTelephonyProvidersEdgesSiteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesSiteParamsWithTimeout creates a new GetTelephonyProvidersEdgesSiteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesSiteParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesSiteParams {
	var ()
	return &GetTelephonyProvidersEdgesSiteParams{

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesSiteParamsWithContext creates a new GetTelephonyProvidersEdgesSiteParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesSiteParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesSiteParams {
	var ()
	return &GetTelephonyProvidersEdgesSiteParams{

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesSiteParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesSiteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesSiteParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesSiteParams {
	var ()
	return &GetTelephonyProvidersEdgesSiteParams{
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesSiteParams contains all the parameters to send to the API endpoint
for the get telephony providers edges site operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesSiteParams struct {

	/*SiteID
	  Site ID

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edges site params
func (o *GetTelephonyProvidersEdgesSiteParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesSiteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges site params
func (o *GetTelephonyProvidersEdgesSiteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges site params
func (o *GetTelephonyProvidersEdgesSiteParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesSiteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges site params
func (o *GetTelephonyProvidersEdgesSiteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges site params
func (o *GetTelephonyProvidersEdgesSiteParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesSiteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges site params
func (o *GetTelephonyProvidersEdgesSiteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the get telephony providers edges site params
func (o *GetTelephonyProvidersEdgesSiteParams) WithSiteID(siteID string) *GetTelephonyProvidersEdgesSiteParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get telephony providers edges site params
func (o *GetTelephonyProvidersEdgesSiteParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesSiteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param siteId
	if err := r.SetPathParam("siteId", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}