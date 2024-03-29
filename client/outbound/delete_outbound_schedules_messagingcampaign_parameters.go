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

// NewDeleteOutboundSchedulesMessagingcampaignParams creates a new DeleteOutboundSchedulesMessagingcampaignParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOutboundSchedulesMessagingcampaignParams() *DeleteOutboundSchedulesMessagingcampaignParams {
	return &DeleteOutboundSchedulesMessagingcampaignParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOutboundSchedulesMessagingcampaignParamsWithTimeout creates a new DeleteOutboundSchedulesMessagingcampaignParams object
// with the ability to set a timeout on a request.
func NewDeleteOutboundSchedulesMessagingcampaignParamsWithTimeout(timeout time.Duration) *DeleteOutboundSchedulesMessagingcampaignParams {
	return &DeleteOutboundSchedulesMessagingcampaignParams{
		timeout: timeout,
	}
}

// NewDeleteOutboundSchedulesMessagingcampaignParamsWithContext creates a new DeleteOutboundSchedulesMessagingcampaignParams object
// with the ability to set a context for a request.
func NewDeleteOutboundSchedulesMessagingcampaignParamsWithContext(ctx context.Context) *DeleteOutboundSchedulesMessagingcampaignParams {
	return &DeleteOutboundSchedulesMessagingcampaignParams{
		Context: ctx,
	}
}

// NewDeleteOutboundSchedulesMessagingcampaignParamsWithHTTPClient creates a new DeleteOutboundSchedulesMessagingcampaignParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOutboundSchedulesMessagingcampaignParamsWithHTTPClient(client *http.Client) *DeleteOutboundSchedulesMessagingcampaignParams {
	return &DeleteOutboundSchedulesMessagingcampaignParams{
		HTTPClient: client,
	}
}

/*
DeleteOutboundSchedulesMessagingcampaignParams contains all the parameters to send to the API endpoint

	for the delete outbound schedules messagingcampaign operation.

	Typically these are written to a http.Request.
*/
type DeleteOutboundSchedulesMessagingcampaignParams struct {

	/* MessagingCampaignID.

	   Messaging Campaign ID
	*/
	MessagingCampaignID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete outbound schedules messagingcampaign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOutboundSchedulesMessagingcampaignParams) WithDefaults() *DeleteOutboundSchedulesMessagingcampaignParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete outbound schedules messagingcampaign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOutboundSchedulesMessagingcampaignParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete outbound schedules messagingcampaign params
func (o *DeleteOutboundSchedulesMessagingcampaignParams) WithTimeout(timeout time.Duration) *DeleteOutboundSchedulesMessagingcampaignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete outbound schedules messagingcampaign params
func (o *DeleteOutboundSchedulesMessagingcampaignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete outbound schedules messagingcampaign params
func (o *DeleteOutboundSchedulesMessagingcampaignParams) WithContext(ctx context.Context) *DeleteOutboundSchedulesMessagingcampaignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete outbound schedules messagingcampaign params
func (o *DeleteOutboundSchedulesMessagingcampaignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete outbound schedules messagingcampaign params
func (o *DeleteOutboundSchedulesMessagingcampaignParams) WithHTTPClient(client *http.Client) *DeleteOutboundSchedulesMessagingcampaignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete outbound schedules messagingcampaign params
func (o *DeleteOutboundSchedulesMessagingcampaignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessagingCampaignID adds the messagingCampaignID to the delete outbound schedules messagingcampaign params
func (o *DeleteOutboundSchedulesMessagingcampaignParams) WithMessagingCampaignID(messagingCampaignID string) *DeleteOutboundSchedulesMessagingcampaignParams {
	o.SetMessagingCampaignID(messagingCampaignID)
	return o
}

// SetMessagingCampaignID adds the messagingCampaignId to the delete outbound schedules messagingcampaign params
func (o *DeleteOutboundSchedulesMessagingcampaignParams) SetMessagingCampaignID(messagingCampaignID string) {
	o.MessagingCampaignID = messagingCampaignID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOutboundSchedulesMessagingcampaignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
