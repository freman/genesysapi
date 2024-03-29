// Code generated by go-swagger; DO NOT EDIT.

package outbound

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

// NewGetOutboundCampaignDiagnosticsParams creates a new GetOutboundCampaignDiagnosticsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOutboundCampaignDiagnosticsParams() *GetOutboundCampaignDiagnosticsParams {
	return &GetOutboundCampaignDiagnosticsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundCampaignDiagnosticsParamsWithTimeout creates a new GetOutboundCampaignDiagnosticsParams object
// with the ability to set a timeout on a request.
func NewGetOutboundCampaignDiagnosticsParamsWithTimeout(timeout time.Duration) *GetOutboundCampaignDiagnosticsParams {
	return &GetOutboundCampaignDiagnosticsParams{
		timeout: timeout,
	}
}

// NewGetOutboundCampaignDiagnosticsParamsWithContext creates a new GetOutboundCampaignDiagnosticsParams object
// with the ability to set a context for a request.
func NewGetOutboundCampaignDiagnosticsParamsWithContext(ctx context.Context) *GetOutboundCampaignDiagnosticsParams {
	return &GetOutboundCampaignDiagnosticsParams{
		Context: ctx,
	}
}

// NewGetOutboundCampaignDiagnosticsParamsWithHTTPClient creates a new GetOutboundCampaignDiagnosticsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOutboundCampaignDiagnosticsParamsWithHTTPClient(client *http.Client) *GetOutboundCampaignDiagnosticsParams {
	return &GetOutboundCampaignDiagnosticsParams{
		HTTPClient: client,
	}
}

/*
GetOutboundCampaignDiagnosticsParams contains all the parameters to send to the API endpoint

	for the get outbound campaign diagnostics operation.

	Typically these are written to a http.Request.
*/
type GetOutboundCampaignDiagnosticsParams struct {

	/* CampaignID.

	   Campaign ID
	*/
	CampaignID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get outbound campaign diagnostics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundCampaignDiagnosticsParams) WithDefaults() *GetOutboundCampaignDiagnosticsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get outbound campaign diagnostics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundCampaignDiagnosticsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get outbound campaign diagnostics params
func (o *GetOutboundCampaignDiagnosticsParams) WithTimeout(timeout time.Duration) *GetOutboundCampaignDiagnosticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound campaign diagnostics params
func (o *GetOutboundCampaignDiagnosticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound campaign diagnostics params
func (o *GetOutboundCampaignDiagnosticsParams) WithContext(ctx context.Context) *GetOutboundCampaignDiagnosticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound campaign diagnostics params
func (o *GetOutboundCampaignDiagnosticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound campaign diagnostics params
func (o *GetOutboundCampaignDiagnosticsParams) WithHTTPClient(client *http.Client) *GetOutboundCampaignDiagnosticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound campaign diagnostics params
func (o *GetOutboundCampaignDiagnosticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the get outbound campaign diagnostics params
func (o *GetOutboundCampaignDiagnosticsParams) WithCampaignID(campaignID string) *GetOutboundCampaignDiagnosticsParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the get outbound campaign diagnostics params
func (o *GetOutboundCampaignDiagnosticsParams) SetCampaignID(campaignID string) {
	o.CampaignID = campaignID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundCampaignDiagnosticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignId
	if err := r.SetPathParam("campaignId", o.CampaignID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
