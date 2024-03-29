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

// NewGetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams creates a new GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams() *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams {
	return &GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParamsWithTimeout creates a new GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams object
// with the ability to set a timeout on a request.
func NewGetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams {
	return &GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams{
		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParamsWithContext creates a new GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams object
// with the ability to set a context for a request.
func NewGetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams {
	return &GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams{
		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams {
	return &GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams{
		HTTPClient: client,
	}
}

/*
GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams contains all the parameters to send to the API endpoint

	for the get telephony providers edges edgegroup edgetrunkbase operation.

	Typically these are written to a http.Request.
*/
type GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams struct {

	/* EdgegroupID.

	   Edge Group ID
	*/
	EdgegroupID string

	/* EdgetrunkbaseID.

	   Edge Trunk Base ID
	*/
	EdgetrunkbaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get telephony providers edges edgegroup edgetrunkbase params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) WithDefaults() *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get telephony providers edges edgegroup edgetrunkbase params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get telephony providers edges edgegroup edgetrunkbase params
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges edgegroup edgetrunkbase params
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges edgegroup edgetrunkbase params
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges edgegroup edgetrunkbase params
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges edgegroup edgetrunkbase params
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges edgegroup edgetrunkbase params
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEdgegroupID adds the edgegroupID to the get telephony providers edges edgegroup edgetrunkbase params
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) WithEdgegroupID(edgegroupID string) *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams {
	o.SetEdgegroupID(edgegroupID)
	return o
}

// SetEdgegroupID adds the edgegroupId to the get telephony providers edges edgegroup edgetrunkbase params
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) SetEdgegroupID(edgegroupID string) {
	o.EdgegroupID = edgegroupID
}

// WithEdgetrunkbaseID adds the edgetrunkbaseID to the get telephony providers edges edgegroup edgetrunkbase params
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) WithEdgetrunkbaseID(edgetrunkbaseID string) *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams {
	o.SetEdgetrunkbaseID(edgetrunkbaseID)
	return o
}

// SetEdgetrunkbaseID adds the edgetrunkbaseId to the get telephony providers edges edgegroup edgetrunkbase params
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) SetEdgetrunkbaseID(edgetrunkbaseID string) {
	o.EdgetrunkbaseID = edgetrunkbaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesEdgegroupEdgetrunkbaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param edgegroupId
	if err := r.SetPathParam("edgegroupId", o.EdgegroupID); err != nil {
		return err
	}

	// path param edgetrunkbaseId
	if err := r.SetPathParam("edgetrunkbaseId", o.EdgetrunkbaseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
