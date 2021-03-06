// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

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

// NewGetSpeechandtextanalyticsDialectsParams creates a new GetSpeechandtextanalyticsDialectsParams object
// with the default values initialized.
func NewGetSpeechandtextanalyticsDialectsParams() *GetSpeechandtextanalyticsDialectsParams {

	return &GetSpeechandtextanalyticsDialectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpeechandtextanalyticsDialectsParamsWithTimeout creates a new GetSpeechandtextanalyticsDialectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSpeechandtextanalyticsDialectsParamsWithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsDialectsParams {

	return &GetSpeechandtextanalyticsDialectsParams{

		timeout: timeout,
	}
}

// NewGetSpeechandtextanalyticsDialectsParamsWithContext creates a new GetSpeechandtextanalyticsDialectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSpeechandtextanalyticsDialectsParamsWithContext(ctx context.Context) *GetSpeechandtextanalyticsDialectsParams {

	return &GetSpeechandtextanalyticsDialectsParams{

		Context: ctx,
	}
}

// NewGetSpeechandtextanalyticsDialectsParamsWithHTTPClient creates a new GetSpeechandtextanalyticsDialectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSpeechandtextanalyticsDialectsParamsWithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsDialectsParams {

	return &GetSpeechandtextanalyticsDialectsParams{
		HTTPClient: client,
	}
}

/*GetSpeechandtextanalyticsDialectsParams contains all the parameters to send to the API endpoint
for the get speechandtextanalytics dialects operation typically these are written to a http.Request
*/
type GetSpeechandtextanalyticsDialectsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get speechandtextanalytics dialects params
func (o *GetSpeechandtextanalyticsDialectsParams) WithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsDialectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get speechandtextanalytics dialects params
func (o *GetSpeechandtextanalyticsDialectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get speechandtextanalytics dialects params
func (o *GetSpeechandtextanalyticsDialectsParams) WithContext(ctx context.Context) *GetSpeechandtextanalyticsDialectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get speechandtextanalytics dialects params
func (o *GetSpeechandtextanalyticsDialectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get speechandtextanalytics dialects params
func (o *GetSpeechandtextanalyticsDialectsParams) WithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsDialectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get speechandtextanalytics dialects params
func (o *GetSpeechandtextanalyticsDialectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpeechandtextanalyticsDialectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
