// Code generated by go-swagger; DO NOT EDIT.

package architect

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
)

// NewPostFlowsJobsParams creates a new PostFlowsJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostFlowsJobsParams() *PostFlowsJobsParams {
	return &PostFlowsJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostFlowsJobsParamsWithTimeout creates a new PostFlowsJobsParams object
// with the ability to set a timeout on a request.
func NewPostFlowsJobsParamsWithTimeout(timeout time.Duration) *PostFlowsJobsParams {
	return &PostFlowsJobsParams{
		timeout: timeout,
	}
}

// NewPostFlowsJobsParamsWithContext creates a new PostFlowsJobsParams object
// with the ability to set a context for a request.
func NewPostFlowsJobsParamsWithContext(ctx context.Context) *PostFlowsJobsParams {
	return &PostFlowsJobsParams{
		Context: ctx,
	}
}

// NewPostFlowsJobsParamsWithHTTPClient creates a new PostFlowsJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostFlowsJobsParamsWithHTTPClient(client *http.Client) *PostFlowsJobsParams {
	return &PostFlowsJobsParams{
		HTTPClient: client,
	}
}

/*
PostFlowsJobsParams contains all the parameters to send to the API endpoint

	for the post flows jobs operation.

	Typically these are written to a http.Request.
*/
type PostFlowsJobsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post flows jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostFlowsJobsParams) WithDefaults() *PostFlowsJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post flows jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostFlowsJobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post flows jobs params
func (o *PostFlowsJobsParams) WithTimeout(timeout time.Duration) *PostFlowsJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post flows jobs params
func (o *PostFlowsJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post flows jobs params
func (o *PostFlowsJobsParams) WithContext(ctx context.Context) *PostFlowsJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post flows jobs params
func (o *PostFlowsJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post flows jobs params
func (o *PostFlowsJobsParams) WithHTTPClient(client *http.Client) *PostFlowsJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post flows jobs params
func (o *PostFlowsJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostFlowsJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
