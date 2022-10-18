// Code generated by go-swagger; DO NOT EDIT.

package analytics

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

// NewGetAnalyticsDataretentionSettingsParams creates a new GetAnalyticsDataretentionSettingsParams object
// with the default values initialized.
func NewGetAnalyticsDataretentionSettingsParams() *GetAnalyticsDataretentionSettingsParams {

	return &GetAnalyticsDataretentionSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnalyticsDataretentionSettingsParamsWithTimeout creates a new GetAnalyticsDataretentionSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAnalyticsDataretentionSettingsParamsWithTimeout(timeout time.Duration) *GetAnalyticsDataretentionSettingsParams {

	return &GetAnalyticsDataretentionSettingsParams{

		timeout: timeout,
	}
}

// NewGetAnalyticsDataretentionSettingsParamsWithContext creates a new GetAnalyticsDataretentionSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAnalyticsDataretentionSettingsParamsWithContext(ctx context.Context) *GetAnalyticsDataretentionSettingsParams {

	return &GetAnalyticsDataretentionSettingsParams{

		Context: ctx,
	}
}

// NewGetAnalyticsDataretentionSettingsParamsWithHTTPClient creates a new GetAnalyticsDataretentionSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAnalyticsDataretentionSettingsParamsWithHTTPClient(client *http.Client) *GetAnalyticsDataretentionSettingsParams {

	return &GetAnalyticsDataretentionSettingsParams{
		HTTPClient: client,
	}
}

/*GetAnalyticsDataretentionSettingsParams contains all the parameters to send to the API endpoint
for the get analytics dataretention settings operation typically these are written to a http.Request
*/
type GetAnalyticsDataretentionSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get analytics dataretention settings params
func (o *GetAnalyticsDataretentionSettingsParams) WithTimeout(timeout time.Duration) *GetAnalyticsDataretentionSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get analytics dataretention settings params
func (o *GetAnalyticsDataretentionSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get analytics dataretention settings params
func (o *GetAnalyticsDataretentionSettingsParams) WithContext(ctx context.Context) *GetAnalyticsDataretentionSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get analytics dataretention settings params
func (o *GetAnalyticsDataretentionSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get analytics dataretention settings params
func (o *GetAnalyticsDataretentionSettingsParams) WithHTTPClient(client *http.Client) *GetAnalyticsDataretentionSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get analytics dataretention settings params
func (o *GetAnalyticsDataretentionSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnalyticsDataretentionSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
