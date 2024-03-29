// Code generated by go-swagger; DO NOT EDIT.

package greetings

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

// NewGetGreetingMediaParams creates a new GetGreetingMediaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGreetingMediaParams() *GetGreetingMediaParams {
	return &GetGreetingMediaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGreetingMediaParamsWithTimeout creates a new GetGreetingMediaParams object
// with the ability to set a timeout on a request.
func NewGetGreetingMediaParamsWithTimeout(timeout time.Duration) *GetGreetingMediaParams {
	return &GetGreetingMediaParams{
		timeout: timeout,
	}
}

// NewGetGreetingMediaParamsWithContext creates a new GetGreetingMediaParams object
// with the ability to set a context for a request.
func NewGetGreetingMediaParamsWithContext(ctx context.Context) *GetGreetingMediaParams {
	return &GetGreetingMediaParams{
		Context: ctx,
	}
}

// NewGetGreetingMediaParamsWithHTTPClient creates a new GetGreetingMediaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGreetingMediaParamsWithHTTPClient(client *http.Client) *GetGreetingMediaParams {
	return &GetGreetingMediaParams{
		HTTPClient: client,
	}
}

/*
GetGreetingMediaParams contains all the parameters to send to the API endpoint

	for the get greeting media operation.

	Typically these are written to a http.Request.
*/
type GetGreetingMediaParams struct {

	/* FormatID.

	   The desired media format.

	   Default: "WAV"
	*/
	FormatID *string

	/* GreetingID.

	   Greeting ID
	*/
	GreetingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get greeting media params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGreetingMediaParams) WithDefaults() *GetGreetingMediaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get greeting media params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGreetingMediaParams) SetDefaults() {
	var (
		formatIDDefault = string("WAV")
	)

	val := GetGreetingMediaParams{
		FormatID: &formatIDDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get greeting media params
func (o *GetGreetingMediaParams) WithTimeout(timeout time.Duration) *GetGreetingMediaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get greeting media params
func (o *GetGreetingMediaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get greeting media params
func (o *GetGreetingMediaParams) WithContext(ctx context.Context) *GetGreetingMediaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get greeting media params
func (o *GetGreetingMediaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get greeting media params
func (o *GetGreetingMediaParams) WithHTTPClient(client *http.Client) *GetGreetingMediaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get greeting media params
func (o *GetGreetingMediaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormatID adds the formatID to the get greeting media params
func (o *GetGreetingMediaParams) WithFormatID(formatID *string) *GetGreetingMediaParams {
	o.SetFormatID(formatID)
	return o
}

// SetFormatID adds the formatId to the get greeting media params
func (o *GetGreetingMediaParams) SetFormatID(formatID *string) {
	o.FormatID = formatID
}

// WithGreetingID adds the greetingID to the get greeting media params
func (o *GetGreetingMediaParams) WithGreetingID(greetingID string) *GetGreetingMediaParams {
	o.SetGreetingID(greetingID)
	return o
}

// SetGreetingID adds the greetingId to the get greeting media params
func (o *GetGreetingMediaParams) SetGreetingID(greetingID string) {
	o.GreetingID = greetingID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGreetingMediaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FormatID != nil {

		// query param formatId
		var qrFormatID string

		if o.FormatID != nil {
			qrFormatID = *o.FormatID
		}
		qFormatID := qrFormatID
		if qFormatID != "" {

			if err := r.SetQueryParam("formatId", qFormatID); err != nil {
				return err
			}
		}
	}

	// path param greetingId
	if err := r.SetPathParam("greetingId", o.GreetingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
