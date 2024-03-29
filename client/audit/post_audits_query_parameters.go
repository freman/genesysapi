// Code generated by go-swagger; DO NOT EDIT.

package audit

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

// NewPostAuditsQueryParams creates a new PostAuditsQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAuditsQueryParams() *PostAuditsQueryParams {
	return &PostAuditsQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuditsQueryParamsWithTimeout creates a new PostAuditsQueryParams object
// with the ability to set a timeout on a request.
func NewPostAuditsQueryParamsWithTimeout(timeout time.Duration) *PostAuditsQueryParams {
	return &PostAuditsQueryParams{
		timeout: timeout,
	}
}

// NewPostAuditsQueryParamsWithContext creates a new PostAuditsQueryParams object
// with the ability to set a context for a request.
func NewPostAuditsQueryParamsWithContext(ctx context.Context) *PostAuditsQueryParams {
	return &PostAuditsQueryParams{
		Context: ctx,
	}
}

// NewPostAuditsQueryParamsWithHTTPClient creates a new PostAuditsQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAuditsQueryParamsWithHTTPClient(client *http.Client) *PostAuditsQueryParams {
	return &PostAuditsQueryParams{
		HTTPClient: client,
	}
}

/*
PostAuditsQueryParams contains all the parameters to send to the API endpoint

	for the post audits query operation.

	Typically these are written to a http.Request.
*/
type PostAuditsQueryParams struct {

	/* Body.

	   query
	*/
	Body *models.AuditQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post audits query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuditsQueryParams) WithDefaults() *PostAuditsQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post audits query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuditsQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post audits query params
func (o *PostAuditsQueryParams) WithTimeout(timeout time.Duration) *PostAuditsQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post audits query params
func (o *PostAuditsQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post audits query params
func (o *PostAuditsQueryParams) WithContext(ctx context.Context) *PostAuditsQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post audits query params
func (o *PostAuditsQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post audits query params
func (o *PostAuditsQueryParams) WithHTTPClient(client *http.Client) *PostAuditsQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post audits query params
func (o *PostAuditsQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post audits query params
func (o *PostAuditsQueryParams) WithBody(body *models.AuditQueryRequest) *PostAuditsQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post audits query params
func (o *PostAuditsQueryParams) SetBody(body *models.AuditQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuditsQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
