// Code generated by go-swagger; DO NOT EDIT.

package web_deployments

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

// NewGetWebdeploymentsConfigurationsParams creates a new GetWebdeploymentsConfigurationsParams object
// with the default values initialized.
func NewGetWebdeploymentsConfigurationsParams() *GetWebdeploymentsConfigurationsParams {
	var (
		showOnlyPublishedDefault = bool(false)
	)
	return &GetWebdeploymentsConfigurationsParams{
		ShowOnlyPublished: &showOnlyPublishedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebdeploymentsConfigurationsParamsWithTimeout creates a new GetWebdeploymentsConfigurationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWebdeploymentsConfigurationsParamsWithTimeout(timeout time.Duration) *GetWebdeploymentsConfigurationsParams {
	var (
		showOnlyPublishedDefault = bool(false)
	)
	return &GetWebdeploymentsConfigurationsParams{
		ShowOnlyPublished: &showOnlyPublishedDefault,

		timeout: timeout,
	}
}

// NewGetWebdeploymentsConfigurationsParamsWithContext creates a new GetWebdeploymentsConfigurationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWebdeploymentsConfigurationsParamsWithContext(ctx context.Context) *GetWebdeploymentsConfigurationsParams {
	var (
		showOnlyPublishedDefault = bool(false)
	)
	return &GetWebdeploymentsConfigurationsParams{
		ShowOnlyPublished: &showOnlyPublishedDefault,

		Context: ctx,
	}
}

// NewGetWebdeploymentsConfigurationsParamsWithHTTPClient creates a new GetWebdeploymentsConfigurationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWebdeploymentsConfigurationsParamsWithHTTPClient(client *http.Client) *GetWebdeploymentsConfigurationsParams {
	var (
		showOnlyPublishedDefault = bool(false)
	)
	return &GetWebdeploymentsConfigurationsParams{
		ShowOnlyPublished: &showOnlyPublishedDefault,
		HTTPClient:        client,
	}
}

/*GetWebdeploymentsConfigurationsParams contains all the parameters to send to the API endpoint
for the get webdeployments configurations operation typically these are written to a http.Request
*/
type GetWebdeploymentsConfigurationsParams struct {

	/*ShowOnlyPublished
	  Get only configuration drafts with published versions

	*/
	ShowOnlyPublished *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get webdeployments configurations params
func (o *GetWebdeploymentsConfigurationsParams) WithTimeout(timeout time.Duration) *GetWebdeploymentsConfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get webdeployments configurations params
func (o *GetWebdeploymentsConfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get webdeployments configurations params
func (o *GetWebdeploymentsConfigurationsParams) WithContext(ctx context.Context) *GetWebdeploymentsConfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get webdeployments configurations params
func (o *GetWebdeploymentsConfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get webdeployments configurations params
func (o *GetWebdeploymentsConfigurationsParams) WithHTTPClient(client *http.Client) *GetWebdeploymentsConfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get webdeployments configurations params
func (o *GetWebdeploymentsConfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithShowOnlyPublished adds the showOnlyPublished to the get webdeployments configurations params
func (o *GetWebdeploymentsConfigurationsParams) WithShowOnlyPublished(showOnlyPublished *bool) *GetWebdeploymentsConfigurationsParams {
	o.SetShowOnlyPublished(showOnlyPublished)
	return o
}

// SetShowOnlyPublished adds the showOnlyPublished to the get webdeployments configurations params
func (o *GetWebdeploymentsConfigurationsParams) SetShowOnlyPublished(showOnlyPublished *bool) {
	o.ShowOnlyPublished = showOnlyPublished
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebdeploymentsConfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ShowOnlyPublished != nil {

		// query param showOnlyPublished
		var qrShowOnlyPublished bool
		if o.ShowOnlyPublished != nil {
			qrShowOnlyPublished = *o.ShowOnlyPublished
		}
		qShowOnlyPublished := swag.FormatBool(qrShowOnlyPublished)
		if qShowOnlyPublished != "" {
			if err := r.SetQueryParam("showOnlyPublished", qShowOnlyPublished); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
