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

// NewGetTelephonyProvidersEdgesTrunkMetricsParams creates a new GetTelephonyProvidersEdgesTrunkMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTelephonyProvidersEdgesTrunkMetricsParams() *GetTelephonyProvidersEdgesTrunkMetricsParams {
	return &GetTelephonyProvidersEdgesTrunkMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesTrunkMetricsParamsWithTimeout creates a new GetTelephonyProvidersEdgesTrunkMetricsParams object
// with the ability to set a timeout on a request.
func NewGetTelephonyProvidersEdgesTrunkMetricsParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesTrunkMetricsParams {
	return &GetTelephonyProvidersEdgesTrunkMetricsParams{
		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesTrunkMetricsParamsWithContext creates a new GetTelephonyProvidersEdgesTrunkMetricsParams object
// with the ability to set a context for a request.
func NewGetTelephonyProvidersEdgesTrunkMetricsParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesTrunkMetricsParams {
	return &GetTelephonyProvidersEdgesTrunkMetricsParams{
		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesTrunkMetricsParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesTrunkMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTelephonyProvidersEdgesTrunkMetricsParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesTrunkMetricsParams {
	return &GetTelephonyProvidersEdgesTrunkMetricsParams{
		HTTPClient: client,
	}
}

/*
GetTelephonyProvidersEdgesTrunkMetricsParams contains all the parameters to send to the API endpoint

	for the get telephony providers edges trunk metrics operation.

	Typically these are written to a http.Request.
*/
type GetTelephonyProvidersEdgesTrunkMetricsParams struct {

	/* TrunkID.

	   Trunk Id
	*/
	TrunkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get telephony providers edges trunk metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesTrunkMetricsParams) WithDefaults() *GetTelephonyProvidersEdgesTrunkMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get telephony providers edges trunk metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesTrunkMetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get telephony providers edges trunk metrics params
func (o *GetTelephonyProvidersEdgesTrunkMetricsParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesTrunkMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges trunk metrics params
func (o *GetTelephonyProvidersEdgesTrunkMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges trunk metrics params
func (o *GetTelephonyProvidersEdgesTrunkMetricsParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesTrunkMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges trunk metrics params
func (o *GetTelephonyProvidersEdgesTrunkMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges trunk metrics params
func (o *GetTelephonyProvidersEdgesTrunkMetricsParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesTrunkMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges trunk metrics params
func (o *GetTelephonyProvidersEdgesTrunkMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrunkID adds the trunkID to the get telephony providers edges trunk metrics params
func (o *GetTelephonyProvidersEdgesTrunkMetricsParams) WithTrunkID(trunkID string) *GetTelephonyProvidersEdgesTrunkMetricsParams {
	o.SetTrunkID(trunkID)
	return o
}

// SetTrunkID adds the trunkId to the get telephony providers edges trunk metrics params
func (o *GetTelephonyProvidersEdgesTrunkMetricsParams) SetTrunkID(trunkID string) {
	o.TrunkID = trunkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesTrunkMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param trunkId
	if err := r.SetPathParam("trunkId", o.TrunkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
