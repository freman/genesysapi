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

// NewGetOutboundMessagingcampaignsDivisionviewParams creates a new GetOutboundMessagingcampaignsDivisionviewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOutboundMessagingcampaignsDivisionviewParams() *GetOutboundMessagingcampaignsDivisionviewParams {
	return &GetOutboundMessagingcampaignsDivisionviewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundMessagingcampaignsDivisionviewParamsWithTimeout creates a new GetOutboundMessagingcampaignsDivisionviewParams object
// with the ability to set a timeout on a request.
func NewGetOutboundMessagingcampaignsDivisionviewParamsWithTimeout(timeout time.Duration) *GetOutboundMessagingcampaignsDivisionviewParams {
	return &GetOutboundMessagingcampaignsDivisionviewParams{
		timeout: timeout,
	}
}

// NewGetOutboundMessagingcampaignsDivisionviewParamsWithContext creates a new GetOutboundMessagingcampaignsDivisionviewParams object
// with the ability to set a context for a request.
func NewGetOutboundMessagingcampaignsDivisionviewParamsWithContext(ctx context.Context) *GetOutboundMessagingcampaignsDivisionviewParams {
	return &GetOutboundMessagingcampaignsDivisionviewParams{
		Context: ctx,
	}
}

// NewGetOutboundMessagingcampaignsDivisionviewParamsWithHTTPClient creates a new GetOutboundMessagingcampaignsDivisionviewParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOutboundMessagingcampaignsDivisionviewParamsWithHTTPClient(client *http.Client) *GetOutboundMessagingcampaignsDivisionviewParams {
	return &GetOutboundMessagingcampaignsDivisionviewParams{
		HTTPClient: client,
	}
}

/*
GetOutboundMessagingcampaignsDivisionviewParams contains all the parameters to send to the API endpoint

	for the get outbound messagingcampaigns divisionview operation.

	Typically these are written to a http.Request.
*/
type GetOutboundMessagingcampaignsDivisionviewParams struct {

	/* MessagingCampaignID.

	   The Messaging Campaign ID
	*/
	MessagingCampaignID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get outbound messagingcampaigns divisionview params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundMessagingcampaignsDivisionviewParams) WithDefaults() *GetOutboundMessagingcampaignsDivisionviewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get outbound messagingcampaigns divisionview params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundMessagingcampaignsDivisionviewParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get outbound messagingcampaigns divisionview params
func (o *GetOutboundMessagingcampaignsDivisionviewParams) WithTimeout(timeout time.Duration) *GetOutboundMessagingcampaignsDivisionviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound messagingcampaigns divisionview params
func (o *GetOutboundMessagingcampaignsDivisionviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound messagingcampaigns divisionview params
func (o *GetOutboundMessagingcampaignsDivisionviewParams) WithContext(ctx context.Context) *GetOutboundMessagingcampaignsDivisionviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound messagingcampaigns divisionview params
func (o *GetOutboundMessagingcampaignsDivisionviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound messagingcampaigns divisionview params
func (o *GetOutboundMessagingcampaignsDivisionviewParams) WithHTTPClient(client *http.Client) *GetOutboundMessagingcampaignsDivisionviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound messagingcampaigns divisionview params
func (o *GetOutboundMessagingcampaignsDivisionviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessagingCampaignID adds the messagingCampaignID to the get outbound messagingcampaigns divisionview params
func (o *GetOutboundMessagingcampaignsDivisionviewParams) WithMessagingCampaignID(messagingCampaignID string) *GetOutboundMessagingcampaignsDivisionviewParams {
	o.SetMessagingCampaignID(messagingCampaignID)
	return o
}

// SetMessagingCampaignID adds the messagingCampaignId to the get outbound messagingcampaigns divisionview params
func (o *GetOutboundMessagingcampaignsDivisionviewParams) SetMessagingCampaignID(messagingCampaignID string) {
	o.MessagingCampaignID = messagingCampaignID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundMessagingcampaignsDivisionviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param messagingCampaignId
	if err := r.SetPathParam("messagingCampaignId", o.MessagingCampaignID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
