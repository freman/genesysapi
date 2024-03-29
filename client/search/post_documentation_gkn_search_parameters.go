// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewPostDocumentationGknSearchParams creates a new PostDocumentationGknSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDocumentationGknSearchParams() *PostDocumentationGknSearchParams {
	return &PostDocumentationGknSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDocumentationGknSearchParamsWithTimeout creates a new PostDocumentationGknSearchParams object
// with the ability to set a timeout on a request.
func NewPostDocumentationGknSearchParamsWithTimeout(timeout time.Duration) *PostDocumentationGknSearchParams {
	return &PostDocumentationGknSearchParams{
		timeout: timeout,
	}
}

// NewPostDocumentationGknSearchParamsWithContext creates a new PostDocumentationGknSearchParams object
// with the ability to set a context for a request.
func NewPostDocumentationGknSearchParamsWithContext(ctx context.Context) *PostDocumentationGknSearchParams {
	return &PostDocumentationGknSearchParams{
		Context: ctx,
	}
}

// NewPostDocumentationGknSearchParamsWithHTTPClient creates a new PostDocumentationGknSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDocumentationGknSearchParamsWithHTTPClient(client *http.Client) *PostDocumentationGknSearchParams {
	return &PostDocumentationGknSearchParams{
		HTTPClient: client,
	}
}

/*
PostDocumentationGknSearchParams contains all the parameters to send to the API endpoint

	for the post documentation gkn search operation.

	Typically these are written to a http.Request.
*/
type PostDocumentationGknSearchParams struct {

	/* Body.

	   Search request options
	*/
	Body *models.GKNDocumentationSearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post documentation gkn search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDocumentationGknSearchParams) WithDefaults() *PostDocumentationGknSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post documentation gkn search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDocumentationGknSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post documentation gkn search params
func (o *PostDocumentationGknSearchParams) WithTimeout(timeout time.Duration) *PostDocumentationGknSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post documentation gkn search params
func (o *PostDocumentationGknSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post documentation gkn search params
func (o *PostDocumentationGknSearchParams) WithContext(ctx context.Context) *PostDocumentationGknSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post documentation gkn search params
func (o *PostDocumentationGknSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post documentation gkn search params
func (o *PostDocumentationGknSearchParams) WithHTTPClient(client *http.Client) *PostDocumentationGknSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post documentation gkn search params
func (o *PostDocumentationGknSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post documentation gkn search params
func (o *PostDocumentationGknSearchParams) WithBody(body *models.GKNDocumentationSearchRequest) *PostDocumentationGknSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post documentation gkn search params
func (o *PostDocumentationGknSearchParams) SetBody(body *models.GKNDocumentationSearchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDocumentationGknSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
