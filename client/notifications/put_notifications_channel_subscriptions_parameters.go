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
	"github.com/go-openapi/swag"

	"github.com/freman/genesysapi/models"
)

// NewPutNotificationsChannelSubscriptionsParams creates a new PutNotificationsChannelSubscriptionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutNotificationsChannelSubscriptionsParams() *PutNotificationsChannelSubscriptionsParams {
	return &PutNotificationsChannelSubscriptionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutNotificationsChannelSubscriptionsParamsWithTimeout creates a new PutNotificationsChannelSubscriptionsParams object
// with the ability to set a timeout on a request.
func NewPutNotificationsChannelSubscriptionsParamsWithTimeout(timeout time.Duration) *PutNotificationsChannelSubscriptionsParams {
	return &PutNotificationsChannelSubscriptionsParams{
		timeout: timeout,
	}
}

// NewPutNotificationsChannelSubscriptionsParamsWithContext creates a new PutNotificationsChannelSubscriptionsParams object
// with the ability to set a context for a request.
func NewPutNotificationsChannelSubscriptionsParamsWithContext(ctx context.Context) *PutNotificationsChannelSubscriptionsParams {
	return &PutNotificationsChannelSubscriptionsParams{
		Context: ctx,
	}
}

// NewPutNotificationsChannelSubscriptionsParamsWithHTTPClient creates a new PutNotificationsChannelSubscriptionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutNotificationsChannelSubscriptionsParamsWithHTTPClient(client *http.Client) *PutNotificationsChannelSubscriptionsParams {
	return &PutNotificationsChannelSubscriptionsParams{
		HTTPClient: client,
	}
}

/*
PutNotificationsChannelSubscriptionsParams contains all the parameters to send to the API endpoint

	for the put notifications channel subscriptions operation.

	Typically these are written to a http.Request.
*/
type PutNotificationsChannelSubscriptionsParams struct {

	/* Body.

	   Body
	*/
	Body []*models.ChannelTopic

	/* ChannelID.

	   Channel ID
	*/
	ChannelID string

	/* IgnoreErrors.

	   Optionally prevent throwing of errors for failed permissions checks.
	*/
	IgnoreErrors *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put notifications channel subscriptions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNotificationsChannelSubscriptionsParams) WithDefaults() *PutNotificationsChannelSubscriptionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put notifications channel subscriptions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutNotificationsChannelSubscriptionsParams) SetDefaults() {
	var (
		ignoreErrorsDefault = bool(false)
	)

	val := PutNotificationsChannelSubscriptionsParams{
		IgnoreErrors: &ignoreErrorsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the put notifications channel subscriptions params
func (o *PutNotificationsChannelSubscriptionsParams) WithTimeout(timeout time.Duration) *PutNotificationsChannelSubscriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put notifications channel subscriptions params
func (o *PutNotificationsChannelSubscriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put notifications channel subscriptions params
func (o *PutNotificationsChannelSubscriptionsParams) WithContext(ctx context.Context) *PutNotificationsChannelSubscriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put notifications channel subscriptions params
func (o *PutNotificationsChannelSubscriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put notifications channel subscriptions params
func (o *PutNotificationsChannelSubscriptionsParams) WithHTTPClient(client *http.Client) *PutNotificationsChannelSubscriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put notifications channel subscriptions params
func (o *PutNotificationsChannelSubscriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put notifications channel subscriptions params
func (o *PutNotificationsChannelSubscriptionsParams) WithBody(body []*models.ChannelTopic) *PutNotificationsChannelSubscriptionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put notifications channel subscriptions params
func (o *PutNotificationsChannelSubscriptionsParams) SetBody(body []*models.ChannelTopic) {
	o.Body = body
}

// WithChannelID adds the channelID to the put notifications channel subscriptions params
func (o *PutNotificationsChannelSubscriptionsParams) WithChannelID(channelID string) *PutNotificationsChannelSubscriptionsParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the put notifications channel subscriptions params
func (o *PutNotificationsChannelSubscriptionsParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WithIgnoreErrors adds the ignoreErrors to the put notifications channel subscriptions params
func (o *PutNotificationsChannelSubscriptionsParams) WithIgnoreErrors(ignoreErrors *bool) *PutNotificationsChannelSubscriptionsParams {
	o.SetIgnoreErrors(ignoreErrors)
	return o
}

// SetIgnoreErrors adds the ignoreErrors to the put notifications channel subscriptions params
func (o *PutNotificationsChannelSubscriptionsParams) SetIgnoreErrors(ignoreErrors *bool) {
	o.IgnoreErrors = ignoreErrors
}

// WriteToRequest writes these params to a swagger request
func (o *PutNotificationsChannelSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param channelId
	if err := r.SetPathParam("channelId", o.ChannelID); err != nil {
		return err
	}

	if o.IgnoreErrors != nil {

		// query param ignoreErrors
		var qrIgnoreErrors bool

		if o.IgnoreErrors != nil {
			qrIgnoreErrors = *o.IgnoreErrors
		}
		qIgnoreErrors := swag.FormatBool(qrIgnoreErrors)
		if qIgnoreErrors != "" {

			if err := r.SetQueryParam("ignoreErrors", qIgnoreErrors); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
