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

// NewGetTelephonyProvidersEdgesSiteNumberplanParams creates a new GetTelephonyProvidersEdgesSiteNumberplanParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesSiteNumberplanParams() *GetTelephonyProvidersEdgesSiteNumberplanParams {
	var ()
	return &GetTelephonyProvidersEdgesSiteNumberplanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesSiteNumberplanParamsWithTimeout creates a new GetTelephonyProvidersEdgesSiteNumberplanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesSiteNumberplanParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesSiteNumberplanParams {
	var ()
	return &GetTelephonyProvidersEdgesSiteNumberplanParams{

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesSiteNumberplanParamsWithContext creates a new GetTelephonyProvidersEdgesSiteNumberplanParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesSiteNumberplanParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesSiteNumberplanParams {
	var ()
	return &GetTelephonyProvidersEdgesSiteNumberplanParams{

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesSiteNumberplanParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesSiteNumberplanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesSiteNumberplanParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesSiteNumberplanParams {
	var ()
	return &GetTelephonyProvidersEdgesSiteNumberplanParams{
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesSiteNumberplanParams contains all the parameters to send to the API endpoint
for the get telephony providers edges site numberplan operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesSiteNumberplanParams struct {

	/*NumberPlanID
	  Number Plan ID

	*/
	NumberPlanID string
	/*SiteID
	  Site ID

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edges site numberplan params
func (o *GetTelephonyProvidersEdgesSiteNumberplanParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesSiteNumberplanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges site numberplan params
func (o *GetTelephonyProvidersEdgesSiteNumberplanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges site numberplan params
func (o *GetTelephonyProvidersEdgesSiteNumberplanParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesSiteNumberplanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges site numberplan params
func (o *GetTelephonyProvidersEdgesSiteNumberplanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges site numberplan params
func (o *GetTelephonyProvidersEdgesSiteNumberplanParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesSiteNumberplanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges site numberplan params
func (o *GetTelephonyProvidersEdgesSiteNumberplanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNumberPlanID adds the numberPlanID to the get telephony providers edges site numberplan params
func (o *GetTelephonyProvidersEdgesSiteNumberplanParams) WithNumberPlanID(numberPlanID string) *GetTelephonyProvidersEdgesSiteNumberplanParams {
	o.SetNumberPlanID(numberPlanID)
	return o
}

// SetNumberPlanID adds the numberPlanId to the get telephony providers edges site numberplan params
func (o *GetTelephonyProvidersEdgesSiteNumberplanParams) SetNumberPlanID(numberPlanID string) {
	o.NumberPlanID = numberPlanID
}

// WithSiteID adds the siteID to the get telephony providers edges site numberplan params
func (o *GetTelephonyProvidersEdgesSiteNumberplanParams) WithSiteID(siteID string) *GetTelephonyProvidersEdgesSiteNumberplanParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get telephony providers edges site numberplan params
func (o *GetTelephonyProvidersEdgesSiteNumberplanParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesSiteNumberplanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param numberPlanId
	if err := r.SetPathParam("numberPlanId", o.NumberPlanID); err != nil {
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