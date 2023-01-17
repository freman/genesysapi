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

	"github.com/freman/genesysapi/models"
)

// NewPutWebdeploymentsConfigurationVersionsDraftParams creates a new PutWebdeploymentsConfigurationVersionsDraftParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutWebdeploymentsConfigurationVersionsDraftParams() *PutWebdeploymentsConfigurationVersionsDraftParams {
	return &PutWebdeploymentsConfigurationVersionsDraftParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutWebdeploymentsConfigurationVersionsDraftParamsWithTimeout creates a new PutWebdeploymentsConfigurationVersionsDraftParams object
// with the ability to set a timeout on a request.
func NewPutWebdeploymentsConfigurationVersionsDraftParamsWithTimeout(timeout time.Duration) *PutWebdeploymentsConfigurationVersionsDraftParams {
	return &PutWebdeploymentsConfigurationVersionsDraftParams{
		timeout: timeout,
	}
}

// NewPutWebdeploymentsConfigurationVersionsDraftParamsWithContext creates a new PutWebdeploymentsConfigurationVersionsDraftParams object
// with the ability to set a context for a request.
func NewPutWebdeploymentsConfigurationVersionsDraftParamsWithContext(ctx context.Context) *PutWebdeploymentsConfigurationVersionsDraftParams {
	return &PutWebdeploymentsConfigurationVersionsDraftParams{
		Context: ctx,
	}
}

// NewPutWebdeploymentsConfigurationVersionsDraftParamsWithHTTPClient creates a new PutWebdeploymentsConfigurationVersionsDraftParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutWebdeploymentsConfigurationVersionsDraftParamsWithHTTPClient(client *http.Client) *PutWebdeploymentsConfigurationVersionsDraftParams {
	return &PutWebdeploymentsConfigurationVersionsDraftParams{
		HTTPClient: client,
	}
}

/*
PutWebdeploymentsConfigurationVersionsDraftParams contains all the parameters to send to the API endpoint

	for the put webdeployments configuration versions draft operation.

	Typically these are written to a http.Request.
*/
type PutWebdeploymentsConfigurationVersionsDraftParams struct {

	/* ConfigurationID.

	   The configuration version ID
	*/
	ConfigurationID string

	// ConfigurationVersion.
	ConfigurationVersion *models.WebDeploymentConfigurationVersion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put webdeployments configuration versions draft params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) WithDefaults() *PutWebdeploymentsConfigurationVersionsDraftParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put webdeployments configuration versions draft params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put webdeployments configuration versions draft params
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) WithTimeout(timeout time.Duration) *PutWebdeploymentsConfigurationVersionsDraftParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put webdeployments configuration versions draft params
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put webdeployments configuration versions draft params
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) WithContext(ctx context.Context) *PutWebdeploymentsConfigurationVersionsDraftParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put webdeployments configuration versions draft params
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put webdeployments configuration versions draft params
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) WithHTTPClient(client *http.Client) *PutWebdeploymentsConfigurationVersionsDraftParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put webdeployments configuration versions draft params
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigurationID adds the configurationID to the put webdeployments configuration versions draft params
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) WithConfigurationID(configurationID string) *PutWebdeploymentsConfigurationVersionsDraftParams {
	o.SetConfigurationID(configurationID)
	return o
}

// SetConfigurationID adds the configurationId to the put webdeployments configuration versions draft params
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) SetConfigurationID(configurationID string) {
	o.ConfigurationID = configurationID
}

// WithConfigurationVersion adds the configurationVersion to the put webdeployments configuration versions draft params
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) WithConfigurationVersion(configurationVersion *models.WebDeploymentConfigurationVersion) *PutWebdeploymentsConfigurationVersionsDraftParams {
	o.SetConfigurationVersion(configurationVersion)
	return o
}

// SetConfigurationVersion adds the configurationVersion to the put webdeployments configuration versions draft params
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) SetConfigurationVersion(configurationVersion *models.WebDeploymentConfigurationVersion) {
	o.ConfigurationVersion = configurationVersion
}

// WriteToRequest writes these params to a swagger request
func (o *PutWebdeploymentsConfigurationVersionsDraftParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param configurationId
	if err := r.SetPathParam("configurationId", o.ConfigurationID); err != nil {
		return err
	}
	if o.ConfigurationVersion != nil {
		if err := r.SetBodyParam(o.ConfigurationVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
