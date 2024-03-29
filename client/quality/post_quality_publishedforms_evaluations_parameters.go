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

// NewPostQualityPublishedformsEvaluationsParams creates a new PostQualityPublishedformsEvaluationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostQualityPublishedformsEvaluationsParams() *PostQualityPublishedformsEvaluationsParams {
	return &PostQualityPublishedformsEvaluationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostQualityPublishedformsEvaluationsParamsWithTimeout creates a new PostQualityPublishedformsEvaluationsParams object
// with the ability to set a timeout on a request.
func NewPostQualityPublishedformsEvaluationsParamsWithTimeout(timeout time.Duration) *PostQualityPublishedformsEvaluationsParams {
	return &PostQualityPublishedformsEvaluationsParams{
		timeout: timeout,
	}
}

// NewPostQualityPublishedformsEvaluationsParamsWithContext creates a new PostQualityPublishedformsEvaluationsParams object
// with the ability to set a context for a request.
func NewPostQualityPublishedformsEvaluationsParamsWithContext(ctx context.Context) *PostQualityPublishedformsEvaluationsParams {
	return &PostQualityPublishedformsEvaluationsParams{
		Context: ctx,
	}
}

// NewPostQualityPublishedformsEvaluationsParamsWithHTTPClient creates a new PostQualityPublishedformsEvaluationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostQualityPublishedformsEvaluationsParamsWithHTTPClient(client *http.Client) *PostQualityPublishedformsEvaluationsParams {
	return &PostQualityPublishedformsEvaluationsParams{
		HTTPClient: client,
	}
}

/*
PostQualityPublishedformsEvaluationsParams contains all the parameters to send to the API endpoint

	for the post quality publishedforms evaluations operation.

	Typically these are written to a http.Request.
*/
type PostQualityPublishedformsEvaluationsParams struct {

	/* Body.

	   Publish request containing id of form to publish
	*/
	Body *models.PublishForm

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post quality publishedforms evaluations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQualityPublishedformsEvaluationsParams) WithDefaults() *PostQualityPublishedformsEvaluationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post quality publishedforms evaluations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQualityPublishedformsEvaluationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post quality publishedforms evaluations params
func (o *PostQualityPublishedformsEvaluationsParams) WithTimeout(timeout time.Duration) *PostQualityPublishedformsEvaluationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post quality publishedforms evaluations params
func (o *PostQualityPublishedformsEvaluationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post quality publishedforms evaluations params
func (o *PostQualityPublishedformsEvaluationsParams) WithContext(ctx context.Context) *PostQualityPublishedformsEvaluationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post quality publishedforms evaluations params
func (o *PostQualityPublishedformsEvaluationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post quality publishedforms evaluations params
func (o *PostQualityPublishedformsEvaluationsParams) WithHTTPClient(client *http.Client) *PostQualityPublishedformsEvaluationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post quality publishedforms evaluations params
func (o *PostQualityPublishedformsEvaluationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post quality publishedforms evaluations params
func (o *PostQualityPublishedformsEvaluationsParams) WithBody(body *models.PublishForm) *PostQualityPublishedformsEvaluationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post quality publishedforms evaluations params
func (o *PostQualityPublishedformsEvaluationsParams) SetBody(body *models.PublishForm) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostQualityPublishedformsEvaluationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
