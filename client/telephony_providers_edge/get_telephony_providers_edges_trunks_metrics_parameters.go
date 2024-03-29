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

// NewGetTelephonyProvidersEdgesTrunksMetricsParams creates a new GetTelephonyProvidersEdgesTrunksMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTelephonyProvidersEdgesTrunksMetricsParams() *GetTelephonyProvidersEdgesTrunksMetricsParams {
	return &GetTelephonyProvidersEdgesTrunksMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesTrunksMetricsParamsWithTimeout creates a new GetTelephonyProvidersEdgesTrunksMetricsParams object
// with the ability to set a timeout on a request.
func NewGetTelephonyProvidersEdgesTrunksMetricsParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesTrunksMetricsParams {
	return &GetTelephonyProvidersEdgesTrunksMetricsParams{
		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesTrunksMetricsParamsWithContext creates a new GetTelephonyProvidersEdgesTrunksMetricsParams object
// with the ability to set a context for a request.
func NewGetTelephonyProvidersEdgesTrunksMetricsParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesTrunksMetricsParams {
	return &GetTelephonyProvidersEdgesTrunksMetricsParams{
		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesTrunksMetricsParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesTrunksMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTelephonyProvidersEdgesTrunksMetricsParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesTrunksMetricsParams {
	return &GetTelephonyProvidersEdgesTrunksMetricsParams{
		HTTPClient: client,
	}
}

/*
GetTelephonyProvidersEdgesTrunksMetricsParams contains all the parameters to send to the API endpoint

	for the get telephony providers edges trunks metrics operation.

	Typically these are written to a http.Request.
*/
type GetTelephonyProvidersEdgesTrunksMetricsParams struct {

	/* TrunkIds.

	   Comma separated list of Trunk Id's
	*/
	TrunkIds string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get telephony providers edges trunks metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesTrunksMetricsParams) WithDefaults() *GetTelephonyProvidersEdgesTrunksMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get telephony providers edges trunks metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesTrunksMetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get telephony providers edges trunks metrics params
func (o *GetTelephonyProvidersEdgesTrunksMetricsParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesTrunksMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges trunks metrics params
func (o *GetTelephonyProvidersEdgesTrunksMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges trunks metrics params
func (o *GetTelephonyProvidersEdgesTrunksMetricsParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesTrunksMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges trunks metrics params
func (o *GetTelephonyProvidersEdgesTrunksMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges trunks metrics params
func (o *GetTelephonyProvidersEdgesTrunksMetricsParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesTrunksMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges trunks metrics params
func (o *GetTelephonyProvidersEdgesTrunksMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrunkIds adds the trunkIds to the get telephony providers edges trunks metrics params
func (o *GetTelephonyProvidersEdgesTrunksMetricsParams) WithTrunkIds(trunkIds string) *GetTelephonyProvidersEdgesTrunksMetricsParams {
	o.SetTrunkIds(trunkIds)
	return o
}

// SetTrunkIds adds the trunkIds to the get telephony providers edges trunks metrics params
func (o *GetTelephonyProvidersEdgesTrunksMetricsParams) SetTrunkIds(trunkIds string) {
	o.TrunkIds = trunkIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesTrunksMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param trunkIds
	qrTrunkIds := o.TrunkIds
	qTrunkIds := qrTrunkIds
	if qTrunkIds != "" {

		if err := r.SetQueryParam("trunkIds", qTrunkIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
