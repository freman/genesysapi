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

// NewGetWebdeploymentsDeploymentsParams creates a new GetWebdeploymentsDeploymentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWebdeploymentsDeploymentsParams() *GetWebdeploymentsDeploymentsParams {
	return &GetWebdeploymentsDeploymentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebdeploymentsDeploymentsParamsWithTimeout creates a new GetWebdeploymentsDeploymentsParams object
// with the ability to set a timeout on a request.
func NewGetWebdeploymentsDeploymentsParamsWithTimeout(timeout time.Duration) *GetWebdeploymentsDeploymentsParams {
	return &GetWebdeploymentsDeploymentsParams{
		timeout: timeout,
	}
}

// NewGetWebdeploymentsDeploymentsParamsWithContext creates a new GetWebdeploymentsDeploymentsParams object
// with the ability to set a context for a request.
func NewGetWebdeploymentsDeploymentsParamsWithContext(ctx context.Context) *GetWebdeploymentsDeploymentsParams {
	return &GetWebdeploymentsDeploymentsParams{
		Context: ctx,
	}
}

// NewGetWebdeploymentsDeploymentsParamsWithHTTPClient creates a new GetWebdeploymentsDeploymentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWebdeploymentsDeploymentsParamsWithHTTPClient(client *http.Client) *GetWebdeploymentsDeploymentsParams {
	return &GetWebdeploymentsDeploymentsParams{
		HTTPClient: client,
	}
}

/*
GetWebdeploymentsDeploymentsParams contains all the parameters to send to the API endpoint

	for the get webdeployments deployments operation.

	Typically these are written to a http.Request.
*/
type GetWebdeploymentsDeploymentsParams struct {

	/* Expand.

	   The specified entity attributes will be filled. Comma separated values expected. Valid values:
	*/
	Expand []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get webdeployments deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebdeploymentsDeploymentsParams) WithDefaults() *GetWebdeploymentsDeploymentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get webdeployments deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebdeploymentsDeploymentsParams) SetDefaults() {
	// no default values defined for this parameter
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

// WithExpand adds the expand to the get webdeployments deployments params
func (o *GetWebdeploymentsDeploymentsParams) WithExpand(expand []string) *GetWebdeploymentsDeploymentsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get webdeployments deployments params
func (o *GetWebdeploymentsDeploymentsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebdeploymentsDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetWebdeploymentsDeployments binds the parameter expand
func (o *GetWebdeploymentsDeploymentsParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}
