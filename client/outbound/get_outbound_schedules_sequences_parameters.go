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

// NewGetOutboundSchedulesSequencesParams creates a new GetOutboundSchedulesSequencesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOutboundSchedulesSequencesParams() *GetOutboundSchedulesSequencesParams {
	return &GetOutboundSchedulesSequencesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundSchedulesSequencesParamsWithTimeout creates a new GetOutboundSchedulesSequencesParams object
// with the ability to set a timeout on a request.
func NewGetOutboundSchedulesSequencesParamsWithTimeout(timeout time.Duration) *GetOutboundSchedulesSequencesParams {
	return &GetOutboundSchedulesSequencesParams{
		timeout: timeout,
	}
}

// NewGetOutboundSchedulesSequencesParamsWithContext creates a new GetOutboundSchedulesSequencesParams object
// with the ability to set a context for a request.
func NewGetOutboundSchedulesSequencesParamsWithContext(ctx context.Context) *GetOutboundSchedulesSequencesParams {
	return &GetOutboundSchedulesSequencesParams{
		Context: ctx,
	}
}

// NewGetOutboundSchedulesSequencesParamsWithHTTPClient creates a new GetOutboundSchedulesSequencesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOutboundSchedulesSequencesParamsWithHTTPClient(client *http.Client) *GetOutboundSchedulesSequencesParams {
	return &GetOutboundSchedulesSequencesParams{
		HTTPClient: client,
	}
}

/*
GetOutboundSchedulesSequencesParams contains all the parameters to send to the API endpoint

	for the get outbound schedules sequences operation.

	Typically these are written to a http.Request.
*/
type GetOutboundSchedulesSequencesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get outbound schedules sequences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundSchedulesSequencesParams) WithDefaults() *GetOutboundSchedulesSequencesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get outbound schedules sequences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundSchedulesSequencesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get outbound schedules sequences params
func (o *GetOutboundSchedulesSequencesParams) WithTimeout(timeout time.Duration) *GetOutboundSchedulesSequencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound schedules sequences params
func (o *GetOutboundSchedulesSequencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound schedules sequences params
func (o *GetOutboundSchedulesSequencesParams) WithContext(ctx context.Context) *GetOutboundSchedulesSequencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound schedules sequences params
func (o *GetOutboundSchedulesSequencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound schedules sequences params
func (o *GetOutboundSchedulesSequencesParams) WithHTTPClient(client *http.Client) *GetOutboundSchedulesSequencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound schedules sequences params
func (o *GetOutboundSchedulesSequencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundSchedulesSequencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
