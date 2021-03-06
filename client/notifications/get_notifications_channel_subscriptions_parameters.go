// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NewGetNotificationsChannelSubscriptionsParams creates a new GetNotificationsChannelSubscriptionsParams object
// with the default values initialized.
func NewGetNotificationsChannelSubscriptionsParams() *GetNotificationsChannelSubscriptionsParams {
	var ()
	return &GetNotificationsChannelSubscriptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNotificationsChannelSubscriptionsParamsWithTimeout creates a new GetNotificationsChannelSubscriptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNotificationsChannelSubscriptionsParamsWithTimeout(timeout time.Duration) *GetNotificationsChannelSubscriptionsParams {
	var ()
	return &GetNotificationsChannelSubscriptionsParams{

		timeout: timeout,
	}
}

// NewGetNotificationsChannelSubscriptionsParamsWithContext creates a new GetNotificationsChannelSubscriptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNotificationsChannelSubscriptionsParamsWithContext(ctx context.Context) *GetNotificationsChannelSubscriptionsParams {
	var ()
	return &GetNotificationsChannelSubscriptionsParams{

		Context: ctx,
	}
}

// NewGetNotificationsChannelSubscriptionsParamsWithHTTPClient creates a new GetNotificationsChannelSubscriptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNotificationsChannelSubscriptionsParamsWithHTTPClient(client *http.Client) *GetNotificationsChannelSubscriptionsParams {
	var ()
	return &GetNotificationsChannelSubscriptionsParams{
		HTTPClient: client,
	}
}

/*GetNotificationsChannelSubscriptionsParams contains all the parameters to send to the API endpoint
for the get notifications channel subscriptions operation typically these are written to a http.Request
*/
type GetNotificationsChannelSubscriptionsParams struct {

	/*ChannelID
	  Channel ID

	*/
	ChannelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get notifications channel subscriptions params
func (o *GetNotificationsChannelSubscriptionsParams) WithTimeout(timeout time.Duration) *GetNotificationsChannelSubscriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get notifications channel subscriptions params
func (o *GetNotificationsChannelSubscriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get notifications channel subscriptions params
func (o *GetNotificationsChannelSubscriptionsParams) WithContext(ctx context.Context) *GetNotificationsChannelSubscriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get notifications channel subscriptions params
func (o *GetNotificationsChannelSubscriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get notifications channel subscriptions params
func (o *GetNotificationsChannelSubscriptionsParams) WithHTTPClient(client *http.Client) *GetNotificationsChannelSubscriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get notifications channel subscriptions params
func (o *GetNotificationsChannelSubscriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelID adds the channelID to the get notifications channel subscriptions params
func (o *GetNotificationsChannelSubscriptionsParams) WithChannelID(channelID string) *GetNotificationsChannelSubscriptionsParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the get notifications channel subscriptions params
func (o *GetNotificationsChannelSubscriptionsParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNotificationsChannelSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channelId
	if err := r.SetPathParam("channelId", o.ChannelID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
