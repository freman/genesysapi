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

// NewGetOutboundSchedulesMessagingcampaignsParams creates a new GetOutboundSchedulesMessagingcampaignsParams object
// with the default values initialized.
func NewGetOutboundSchedulesMessagingcampaignsParams() *GetOutboundSchedulesMessagingcampaignsParams {

	return &GetOutboundSchedulesMessagingcampaignsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundSchedulesMessagingcampaignsParamsWithTimeout creates a new GetOutboundSchedulesMessagingcampaignsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutboundSchedulesMessagingcampaignsParamsWithTimeout(timeout time.Duration) *GetOutboundSchedulesMessagingcampaignsParams {

	return &GetOutboundSchedulesMessagingcampaignsParams{

		timeout: timeout,
	}
}

// NewGetOutboundSchedulesMessagingcampaignsParamsWithContext creates a new GetOutboundSchedulesMessagingcampaignsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutboundSchedulesMessagingcampaignsParamsWithContext(ctx context.Context) *GetOutboundSchedulesMessagingcampaignsParams {

	return &GetOutboundSchedulesMessagingcampaignsParams{

		Context: ctx,
	}
}

// NewGetOutboundSchedulesMessagingcampaignsParamsWithHTTPClient creates a new GetOutboundSchedulesMessagingcampaignsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutboundSchedulesMessagingcampaignsParamsWithHTTPClient(client *http.Client) *GetOutboundSchedulesMessagingcampaignsParams {

	return &GetOutboundSchedulesMessagingcampaignsParams{
		HTTPClient: client,
	}
}

/*GetOutboundSchedulesMessagingcampaignsParams contains all the parameters to send to the API endpoint
for the get outbound schedules messagingcampaigns operation typically these are written to a http.Request
*/
type GetOutboundSchedulesMessagingcampaignsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get outbound schedules messagingcampaigns params
func (o *GetOutboundSchedulesMessagingcampaignsParams) WithTimeout(timeout time.Duration) *GetOutboundSchedulesMessagingcampaignsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound schedules messagingcampaigns params
func (o *GetOutboundSchedulesMessagingcampaignsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound schedules messagingcampaigns params
func (o *GetOutboundSchedulesMessagingcampaignsParams) WithContext(ctx context.Context) *GetOutboundSchedulesMessagingcampaignsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound schedules messagingcampaigns params
func (o *GetOutboundSchedulesMessagingcampaignsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound schedules messagingcampaigns params
func (o *GetOutboundSchedulesMessagingcampaignsParams) WithHTTPClient(client *http.Client) *GetOutboundSchedulesMessagingcampaignsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound schedules messagingcampaigns params
func (o *GetOutboundSchedulesMessagingcampaignsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundSchedulesMessagingcampaignsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}