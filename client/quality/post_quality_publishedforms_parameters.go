// Code generated by go-swagger; DO NOT EDIT.

package quality

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

// NewPostQualityPublishedformsParams creates a new PostQualityPublishedformsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostQualityPublishedformsParams() *PostQualityPublishedformsParams {
	return &PostQualityPublishedformsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostQualityPublishedformsParamsWithTimeout creates a new PostQualityPublishedformsParams object
// with the ability to set a timeout on a request.
func NewPostQualityPublishedformsParamsWithTimeout(timeout time.Duration) *PostQualityPublishedformsParams {
	return &PostQualityPublishedformsParams{
		timeout: timeout,
	}
}

// NewPostQualityPublishedformsParamsWithContext creates a new PostQualityPublishedformsParams object
// with the ability to set a context for a request.
func NewPostQualityPublishedformsParamsWithContext(ctx context.Context) *PostQualityPublishedformsParams {
	return &PostQualityPublishedformsParams{
		Context: ctx,
	}
}

// NewPostQualityPublishedformsParamsWithHTTPClient creates a new PostQualityPublishedformsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostQualityPublishedformsParamsWithHTTPClient(client *http.Client) *PostQualityPublishedformsParams {
	return &PostQualityPublishedformsParams{
		HTTPClient: client,
	}
}

/*
PostQualityPublishedformsParams contains all the parameters to send to the API endpoint

	for the post quality publishedforms operation.

	Typically these are written to a http.Request.
*/
type PostQualityPublishedformsParams struct {

	/* Body.

	   Publish request containing id of form to publish
	*/
	Body *models.PublishForm

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post quality publishedforms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQualityPublishedformsParams) WithDefaults() *PostQualityPublishedformsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post quality publishedforms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQualityPublishedformsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post quality publishedforms params
func (o *PostQualityPublishedformsParams) WithTimeout(timeout time.Duration) *PostQualityPublishedformsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post quality publishedforms params
func (o *PostQualityPublishedformsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post quality publishedforms params
func (o *PostQualityPublishedformsParams) WithContext(ctx context.Context) *PostQualityPublishedformsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post quality publishedforms params
func (o *PostQualityPublishedformsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post quality publishedforms params
func (o *PostQualityPublishedformsParams) WithHTTPClient(client *http.Client) *PostQualityPublishedformsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post quality publishedforms params
func (o *PostQualityPublishedformsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post quality publishedforms params
func (o *PostQualityPublishedformsParams) WithBody(body *models.PublishForm) *PostQualityPublishedformsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post quality publishedforms params
func (o *PostQualityPublishedformsParams) SetBody(body *models.PublishForm) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostQualityPublishedformsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
