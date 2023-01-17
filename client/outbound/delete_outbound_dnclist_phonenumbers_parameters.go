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
	"github.com/go-openapi/swag"
)

// NewDeleteOutboundDnclistPhonenumbersParams creates a new DeleteOutboundDnclistPhonenumbersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOutboundDnclistPhonenumbersParams() *DeleteOutboundDnclistPhonenumbersParams {
	return &DeleteOutboundDnclistPhonenumbersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOutboundDnclistPhonenumbersParamsWithTimeout creates a new DeleteOutboundDnclistPhonenumbersParams object
// with the ability to set a timeout on a request.
func NewDeleteOutboundDnclistPhonenumbersParamsWithTimeout(timeout time.Duration) *DeleteOutboundDnclistPhonenumbersParams {
	return &DeleteOutboundDnclistPhonenumbersParams{
		timeout: timeout,
	}
}

// NewDeleteOutboundDnclistPhonenumbersParamsWithContext creates a new DeleteOutboundDnclistPhonenumbersParams object
// with the ability to set a context for a request.
func NewDeleteOutboundDnclistPhonenumbersParamsWithContext(ctx context.Context) *DeleteOutboundDnclistPhonenumbersParams {
	return &DeleteOutboundDnclistPhonenumbersParams{
		Context: ctx,
	}
}

// NewDeleteOutboundDnclistPhonenumbersParamsWithHTTPClient creates a new DeleteOutboundDnclistPhonenumbersParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOutboundDnclistPhonenumbersParamsWithHTTPClient(client *http.Client) *DeleteOutboundDnclistPhonenumbersParams {
	return &DeleteOutboundDnclistPhonenumbersParams{
		HTTPClient: client,
	}
}

/*
DeleteOutboundDnclistPhonenumbersParams contains all the parameters to send to the API endpoint

	for the delete outbound dnclist phonenumbers operation.

	Typically these are written to a http.Request.
*/
type DeleteOutboundDnclistPhonenumbersParams struct {

	/* DncListID.

	   DncList ID
	*/
	DncListID string

	/* ExpiredOnly.

	   Set to true to only remove DNC entries that are expired
	*/
	ExpiredOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete outbound dnclist phonenumbers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOutboundDnclistPhonenumbersParams) WithDefaults() *DeleteOutboundDnclistPhonenumbersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete outbound dnclist phonenumbers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOutboundDnclistPhonenumbersParams) SetDefaults() {
	var (
		expiredOnlyDefault = bool(false)
	)

	val := DeleteOutboundDnclistPhonenumbersParams{
		ExpiredOnly: &expiredOnlyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete outbound dnclist phonenumbers params
func (o *DeleteOutboundDnclistPhonenumbersParams) WithTimeout(timeout time.Duration) *DeleteOutboundDnclistPhonenumbersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete outbound dnclist phonenumbers params
func (o *DeleteOutboundDnclistPhonenumbersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete outbound dnclist phonenumbers params
func (o *DeleteOutboundDnclistPhonenumbersParams) WithContext(ctx context.Context) *DeleteOutboundDnclistPhonenumbersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete outbound dnclist phonenumbers params
func (o *DeleteOutboundDnclistPhonenumbersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete outbound dnclist phonenumbers params
func (o *DeleteOutboundDnclistPhonenumbersParams) WithHTTPClient(client *http.Client) *DeleteOutboundDnclistPhonenumbersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete outbound dnclist phonenumbers params
func (o *DeleteOutboundDnclistPhonenumbersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDncListID adds the dncListID to the delete outbound dnclist phonenumbers params
func (o *DeleteOutboundDnclistPhonenumbersParams) WithDncListID(dncListID string) *DeleteOutboundDnclistPhonenumbersParams {
	o.SetDncListID(dncListID)
	return o
}

// SetDncListID adds the dncListId to the delete outbound dnclist phonenumbers params
func (o *DeleteOutboundDnclistPhonenumbersParams) SetDncListID(dncListID string) {
	o.DncListID = dncListID
}

// WithExpiredOnly adds the expiredOnly to the delete outbound dnclist phonenumbers params
func (o *DeleteOutboundDnclistPhonenumbersParams) WithExpiredOnly(expiredOnly *bool) *DeleteOutboundDnclistPhonenumbersParams {
	o.SetExpiredOnly(expiredOnly)
	return o
}

// SetExpiredOnly adds the expiredOnly to the delete outbound dnclist phonenumbers params
func (o *DeleteOutboundDnclistPhonenumbersParams) SetExpiredOnly(expiredOnly *bool) {
	o.ExpiredOnly = expiredOnly
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOutboundDnclistPhonenumbersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dncListId
	if err := r.SetPathParam("dncListId", o.DncListID); err != nil {
		return err
	}

	if o.ExpiredOnly != nil {

		// query param expiredOnly
		var qrExpiredOnly bool

		if o.ExpiredOnly != nil {
			qrExpiredOnly = *o.ExpiredOnly
		}
		qExpiredOnly := swag.FormatBool(qrExpiredOnly)
		if qExpiredOnly != "" {

			if err := r.SetQueryParam("expiredOnly", qExpiredOnly); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
