// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

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

// NewPostTelephonyProvidersEdgesAddressvalidationParams creates a new PostTelephonyProvidersEdgesAddressvalidationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostTelephonyProvidersEdgesAddressvalidationParams() *PostTelephonyProvidersEdgesAddressvalidationParams {
	return &PostTelephonyProvidersEdgesAddressvalidationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostTelephonyProvidersEdgesAddressvalidationParamsWithTimeout creates a new PostTelephonyProvidersEdgesAddressvalidationParams object
// with the ability to set a timeout on a request.
func NewPostTelephonyProvidersEdgesAddressvalidationParamsWithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgesAddressvalidationParams {
	return &PostTelephonyProvidersEdgesAddressvalidationParams{
		timeout: timeout,
	}
}

// NewPostTelephonyProvidersEdgesAddressvalidationParamsWithContext creates a new PostTelephonyProvidersEdgesAddressvalidationParams object
// with the ability to set a context for a request.
func NewPostTelephonyProvidersEdgesAddressvalidationParamsWithContext(ctx context.Context) *PostTelephonyProvidersEdgesAddressvalidationParams {
	return &PostTelephonyProvidersEdgesAddressvalidationParams{
		Context: ctx,
	}
}

// NewPostTelephonyProvidersEdgesAddressvalidationParamsWithHTTPClient creates a new PostTelephonyProvidersEdgesAddressvalidationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostTelephonyProvidersEdgesAddressvalidationParamsWithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgesAddressvalidationParams {
	return &PostTelephonyProvidersEdgesAddressvalidationParams{
		HTTPClient: client,
	}
}

/*
PostTelephonyProvidersEdgesAddressvalidationParams contains all the parameters to send to the API endpoint

	for the post telephony providers edges addressvalidation operation.

	Typically these are written to a http.Request.
*/
type PostTelephonyProvidersEdgesAddressvalidationParams struct {

	/* Body.

	   Address
	*/
	Body *models.ValidateAddressRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post telephony providers edges addressvalidation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostTelephonyProvidersEdgesAddressvalidationParams) WithDefaults() *PostTelephonyProvidersEdgesAddressvalidationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post telephony providers edges addressvalidation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostTelephonyProvidersEdgesAddressvalidationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post telephony providers edges addressvalidation params
func (o *PostTelephonyProvidersEdgesAddressvalidationParams) WithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgesAddressvalidationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post telephony providers edges addressvalidation params
func (o *PostTelephonyProvidersEdgesAddressvalidationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post telephony providers edges addressvalidation params
func (o *PostTelephonyProvidersEdgesAddressvalidationParams) WithContext(ctx context.Context) *PostTelephonyProvidersEdgesAddressvalidationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post telephony providers edges addressvalidation params
func (o *PostTelephonyProvidersEdgesAddressvalidationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post telephony providers edges addressvalidation params
func (o *PostTelephonyProvidersEdgesAddressvalidationParams) WithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgesAddressvalidationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post telephony providers edges addressvalidation params
func (o *PostTelephonyProvidersEdgesAddressvalidationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post telephony providers edges addressvalidation params
func (o *PostTelephonyProvidersEdgesAddressvalidationParams) WithBody(body *models.ValidateAddressRequest) *PostTelephonyProvidersEdgesAddressvalidationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post telephony providers edges addressvalidation params
func (o *PostTelephonyProvidersEdgesAddressvalidationParams) SetBody(body *models.ValidateAddressRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTelephonyProvidersEdgesAddressvalidationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
