// Code generated by go-swagger; DO NOT EDIT.

package uploads

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

// NewPostUploadsRecordingsParams creates a new PostUploadsRecordingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUploadsRecordingsParams() *PostUploadsRecordingsParams {
	return &PostUploadsRecordingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUploadsRecordingsParamsWithTimeout creates a new PostUploadsRecordingsParams object
// with the ability to set a timeout on a request.
func NewPostUploadsRecordingsParamsWithTimeout(timeout time.Duration) *PostUploadsRecordingsParams {
	return &PostUploadsRecordingsParams{
		timeout: timeout,
	}
}

// NewPostUploadsRecordingsParamsWithContext creates a new PostUploadsRecordingsParams object
// with the ability to set a context for a request.
func NewPostUploadsRecordingsParamsWithContext(ctx context.Context) *PostUploadsRecordingsParams {
	return &PostUploadsRecordingsParams{
		Context: ctx,
	}
}

// NewPostUploadsRecordingsParamsWithHTTPClient creates a new PostUploadsRecordingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUploadsRecordingsParamsWithHTTPClient(client *http.Client) *PostUploadsRecordingsParams {
	return &PostUploadsRecordingsParams{
		HTTPClient: client,
	}
}

/*
PostUploadsRecordingsParams contains all the parameters to send to the API endpoint

	for the post uploads recordings operation.

	Typically these are written to a http.Request.
*/
type PostUploadsRecordingsParams struct {

	/* Body.

	   query
	*/
	Body *models.UploadURLRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post uploads recordings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUploadsRecordingsParams) WithDefaults() *PostUploadsRecordingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post uploads recordings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUploadsRecordingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post uploads recordings params
func (o *PostUploadsRecordingsParams) WithTimeout(timeout time.Duration) *PostUploadsRecordingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post uploads recordings params
func (o *PostUploadsRecordingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post uploads recordings params
func (o *PostUploadsRecordingsParams) WithContext(ctx context.Context) *PostUploadsRecordingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post uploads recordings params
func (o *PostUploadsRecordingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post uploads recordings params
func (o *PostUploadsRecordingsParams) WithHTTPClient(client *http.Client) *PostUploadsRecordingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post uploads recordings params
func (o *PostUploadsRecordingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post uploads recordings params
func (o *PostUploadsRecordingsParams) WithBody(body *models.UploadURLRequest) *PostUploadsRecordingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post uploads recordings params
func (o *PostUploadsRecordingsParams) SetBody(body *models.UploadURLRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUploadsRecordingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
