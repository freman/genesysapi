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
)

// NewGetWebdeploymentsDeploymentsParams creates a new GetWebdeploymentsDeploymentsParams object
// with the default values initialized.
func NewGetWebdeploymentsDeploymentsParams() *GetWebdeploymentsDeploymentsParams {

	return &GetWebdeploymentsDeploymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebdeploymentsDeploymentsParamsWithTimeout creates a new GetWebdeploymentsDeploymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWebdeploymentsDeploymentsParamsWithTimeout(timeout time.Duration) *GetWebdeploymentsDeploymentsParams {

	return &GetWebdeploymentsDeploymentsParams{

		timeout: timeout,
	}
}

// NewGetWebdeploymentsDeploymentsParamsWithContext creates a new GetWebdeploymentsDeploymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWebdeploymentsDeploymentsParamsWithContext(ctx context.Context) *GetWebdeploymentsDeploymentsParams {

	return &GetWebdeploymentsDeploymentsParams{

		Context: ctx,
	}
}

// NewGetWebdeploymentsDeploymentsParamsWithHTTPClient creates a new GetWebdeploymentsDeploymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWebdeploymentsDeploymentsParamsWithHTTPClient(client *http.Client) *GetWebdeploymentsDeploymentsParams {

	return &GetWebdeploymentsDeploymentsParams{
		HTTPClient: client,
	}
}

/*GetWebdeploymentsDeploymentsParams contains all the parameters to send to the API endpoint
for the get webdeployments deployments operation typically these are written to a http.Request
*/
type GetWebdeploymentsDeploymentsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get webdeployments deployments params
func (o *GetWebdeploymentsDeploymentsParams) WithTimeout(timeout time.Duration) *GetWebdeploymentsDeploymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get webdeployments deployments params
func (o *GetWebdeploymentsDeploymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get webdeployments deployments params
func (o *GetWebdeploymentsDeploymentsParams) WithContext(ctx context.Context) *GetWebdeploymentsDeploymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get webdeployments deployments params
func (o *GetWebdeploymentsDeploymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get webdeployments deployments params
func (o *GetWebdeploymentsDeploymentsParams) WithHTTPClient(client *http.Client) *GetWebdeploymentsDeploymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get webdeployments deployments params
func (o *GetWebdeploymentsDeploymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebdeploymentsDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}