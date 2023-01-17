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

	"github.com/freman/genesysapi/models"
)

// NewPostOutboundRulesetsParams creates a new PostOutboundRulesetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOutboundRulesetsParams() *PostOutboundRulesetsParams {
	return &PostOutboundRulesetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOutboundRulesetsParamsWithTimeout creates a new PostOutboundRulesetsParams object
// with the ability to set a timeout on a request.
func NewPostOutboundRulesetsParamsWithTimeout(timeout time.Duration) *PostOutboundRulesetsParams {
	return &PostOutboundRulesetsParams{
		timeout: timeout,
	}
}

// NewPostOutboundRulesetsParamsWithContext creates a new PostOutboundRulesetsParams object
// with the ability to set a context for a request.
func NewPostOutboundRulesetsParamsWithContext(ctx context.Context) *PostOutboundRulesetsParams {
	return &PostOutboundRulesetsParams{
		Context: ctx,
	}
}

// NewPostOutboundRulesetsParamsWithHTTPClient creates a new PostOutboundRulesetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOutboundRulesetsParamsWithHTTPClient(client *http.Client) *PostOutboundRulesetsParams {
	return &PostOutboundRulesetsParams{
		HTTPClient: client,
	}
}

/*
PostOutboundRulesetsParams contains all the parameters to send to the API endpoint

	for the post outbound rulesets operation.

	Typically these are written to a http.Request.
*/
type PostOutboundRulesetsParams struct {

	/* Body.

	   RuleSet
	*/
	Body *models.RuleSet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post outbound rulesets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOutboundRulesetsParams) WithDefaults() *PostOutboundRulesetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post outbound rulesets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOutboundRulesetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post outbound rulesets params
func (o *PostOutboundRulesetsParams) WithTimeout(timeout time.Duration) *PostOutboundRulesetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post outbound rulesets params
func (o *PostOutboundRulesetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post outbound rulesets params
func (o *PostOutboundRulesetsParams) WithContext(ctx context.Context) *PostOutboundRulesetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post outbound rulesets params
func (o *PostOutboundRulesetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post outbound rulesets params
func (o *PostOutboundRulesetsParams) WithHTTPClient(client *http.Client) *PostOutboundRulesetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post outbound rulesets params
func (o *PostOutboundRulesetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post outbound rulesets params
func (o *PostOutboundRulesetsParams) WithBody(body *models.RuleSet) *PostOutboundRulesetsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post outbound rulesets params
func (o *PostOutboundRulesetsParams) SetBody(body *models.RuleSet) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostOutboundRulesetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
