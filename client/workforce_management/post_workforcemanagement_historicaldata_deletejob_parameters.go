// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

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

// NewPostWorkforcemanagementHistoricaldataDeletejobParams creates a new PostWorkforcemanagementHistoricaldataDeletejobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementHistoricaldataDeletejobParams() *PostWorkforcemanagementHistoricaldataDeletejobParams {
	return &PostWorkforcemanagementHistoricaldataDeletejobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementHistoricaldataDeletejobParamsWithTimeout creates a new PostWorkforcemanagementHistoricaldataDeletejobParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementHistoricaldataDeletejobParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementHistoricaldataDeletejobParams {
	return &PostWorkforcemanagementHistoricaldataDeletejobParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementHistoricaldataDeletejobParamsWithContext creates a new PostWorkforcemanagementHistoricaldataDeletejobParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementHistoricaldataDeletejobParamsWithContext(ctx context.Context) *PostWorkforcemanagementHistoricaldataDeletejobParams {
	return &PostWorkforcemanagementHistoricaldataDeletejobParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementHistoricaldataDeletejobParamsWithHTTPClient creates a new PostWorkforcemanagementHistoricaldataDeletejobParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementHistoricaldataDeletejobParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementHistoricaldataDeletejobParams {
	return &PostWorkforcemanagementHistoricaldataDeletejobParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementHistoricaldataDeletejobParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement historicaldata deletejob operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementHistoricaldataDeletejobParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workforcemanagement historicaldata deletejob params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementHistoricaldataDeletejobParams) WithDefaults() *PostWorkforcemanagementHistoricaldataDeletejobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement historicaldata deletejob params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementHistoricaldataDeletejobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workforcemanagement historicaldata deletejob params
func (o *PostWorkforcemanagementHistoricaldataDeletejobParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementHistoricaldataDeletejobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement historicaldata deletejob params
func (o *PostWorkforcemanagementHistoricaldataDeletejobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement historicaldata deletejob params
func (o *PostWorkforcemanagementHistoricaldataDeletejobParams) WithContext(ctx context.Context) *PostWorkforcemanagementHistoricaldataDeletejobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement historicaldata deletejob params
func (o *PostWorkforcemanagementHistoricaldataDeletejobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement historicaldata deletejob params
func (o *PostWorkforcemanagementHistoricaldataDeletejobParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementHistoricaldataDeletejobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement historicaldata deletejob params
func (o *PostWorkforcemanagementHistoricaldataDeletejobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementHistoricaldataDeletejobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
