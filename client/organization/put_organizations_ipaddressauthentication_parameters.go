// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// NewPutOrganizationsIpaddressauthenticationParams creates a new PutOrganizationsIpaddressauthenticationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutOrganizationsIpaddressauthenticationParams() *PutOrganizationsIpaddressauthenticationParams {
	return &PutOrganizationsIpaddressauthenticationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrganizationsIpaddressauthenticationParamsWithTimeout creates a new PutOrganizationsIpaddressauthenticationParams object
// with the ability to set a timeout on a request.
func NewPutOrganizationsIpaddressauthenticationParamsWithTimeout(timeout time.Duration) *PutOrganizationsIpaddressauthenticationParams {
	return &PutOrganizationsIpaddressauthenticationParams{
		timeout: timeout,
	}
}

// NewPutOrganizationsIpaddressauthenticationParamsWithContext creates a new PutOrganizationsIpaddressauthenticationParams object
// with the ability to set a context for a request.
func NewPutOrganizationsIpaddressauthenticationParamsWithContext(ctx context.Context) *PutOrganizationsIpaddressauthenticationParams {
	return &PutOrganizationsIpaddressauthenticationParams{
		Context: ctx,
	}
}

// NewPutOrganizationsIpaddressauthenticationParamsWithHTTPClient creates a new PutOrganizationsIpaddressauthenticationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutOrganizationsIpaddressauthenticationParamsWithHTTPClient(client *http.Client) *PutOrganizationsIpaddressauthenticationParams {
	return &PutOrganizationsIpaddressauthenticationParams{
		HTTPClient: client,
	}
}

/*
PutOrganizationsIpaddressauthenticationParams contains all the parameters to send to the API endpoint

	for the put organizations ipaddressauthentication operation.

	Typically these are written to a http.Request.
*/
type PutOrganizationsIpaddressauthenticationParams struct {

	/* Body.

	   IP address Whitelist settings
	*/
	Body *models.IPAddressAuthentication

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put organizations ipaddressauthentication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrganizationsIpaddressauthenticationParams) WithDefaults() *PutOrganizationsIpaddressauthenticationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put organizations ipaddressauthentication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOrganizationsIpaddressauthenticationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put organizations ipaddressauthentication params
func (o *PutOrganizationsIpaddressauthenticationParams) WithTimeout(timeout time.Duration) *PutOrganizationsIpaddressauthenticationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put organizations ipaddressauthentication params
func (o *PutOrganizationsIpaddressauthenticationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put organizations ipaddressauthentication params
func (o *PutOrganizationsIpaddressauthenticationParams) WithContext(ctx context.Context) *PutOrganizationsIpaddressauthenticationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put organizations ipaddressauthentication params
func (o *PutOrganizationsIpaddressauthenticationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put organizations ipaddressauthentication params
func (o *PutOrganizationsIpaddressauthenticationParams) WithHTTPClient(client *http.Client) *PutOrganizationsIpaddressauthenticationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put organizations ipaddressauthentication params
func (o *PutOrganizationsIpaddressauthenticationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put organizations ipaddressauthentication params
func (o *PutOrganizationsIpaddressauthenticationParams) WithBody(body *models.IPAddressAuthentication) *PutOrganizationsIpaddressauthenticationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put organizations ipaddressauthentication params
func (o *PutOrganizationsIpaddressauthenticationParams) SetBody(body *models.IPAddressAuthentication) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrganizationsIpaddressauthenticationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
