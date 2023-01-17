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

// NewHeadNotificationsChannelParams creates a new HeadNotificationsChannelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadNotificationsChannelParams() *HeadNotificationsChannelParams {
	return &HeadNotificationsChannelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadNotificationsChannelParamsWithTimeout creates a new HeadNotificationsChannelParams object
// with the ability to set a timeout on a request.
func NewHeadNotificationsChannelParamsWithTimeout(timeout time.Duration) *HeadNotificationsChannelParams {
	return &HeadNotificationsChannelParams{
		timeout: timeout,
	}
}

// NewHeadNotificationsChannelParamsWithContext creates a new HeadNotificationsChannelParams object
// with the ability to set a context for a request.
func NewHeadNotificationsChannelParamsWithContext(ctx context.Context) *HeadNotificationsChannelParams {
	return &HeadNotificationsChannelParams{
		Context: ctx,
	}
}

// NewHeadNotificationsChannelParamsWithHTTPClient creates a new HeadNotificationsChannelParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadNotificationsChannelParamsWithHTTPClient(client *http.Client) *HeadNotificationsChannelParams {
	return &HeadNotificationsChannelParams{
		HTTPClient: client,
	}
}

/*
HeadNotificationsChannelParams contains all the parameters to send to the API endpoint

	for the head notifications channel operation.

	Typically these are written to a http.Request.
*/
type HeadNotificationsChannelParams struct {

	/* ChannelID.

	   Channel ID
	*/
	ChannelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the head notifications channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadNotificationsChannelParams) WithDefaults() *HeadNotificationsChannelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the head notifications channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadNotificationsChannelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the head notifications channel params
func (o *HeadNotificationsChannelParams) WithTimeout(timeout time.Duration) *HeadNotificationsChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head notifications channel params
func (o *HeadNotificationsChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head notifications channel params
func (o *HeadNotificationsChannelParams) WithContext(ctx context.Context) *HeadNotificationsChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head notifications channel params
func (o *HeadNotificationsChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head notifications channel params
func (o *HeadNotificationsChannelParams) WithHTTPClient(client *http.Client) *HeadNotificationsChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head notifications channel params
func (o *HeadNotificationsChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelID adds the channelID to the head notifications channel params
func (o *HeadNotificationsChannelParams) WithChannelID(channelID string) *HeadNotificationsChannelParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the head notifications channel params
func (o *HeadNotificationsChannelParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WriteToRequest writes these params to a swagger request
func (o *HeadNotificationsChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
