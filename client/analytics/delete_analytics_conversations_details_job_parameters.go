// Code generated by go-swagger; DO NOT EDIT.

package analytics

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

// NewDeleteAnalyticsConversationsDetailsJobParams creates a new DeleteAnalyticsConversationsDetailsJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAnalyticsConversationsDetailsJobParams() *DeleteAnalyticsConversationsDetailsJobParams {
	return &DeleteAnalyticsConversationsDetailsJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAnalyticsConversationsDetailsJobParamsWithTimeout creates a new DeleteAnalyticsConversationsDetailsJobParams object
// with the ability to set a timeout on a request.
func NewDeleteAnalyticsConversationsDetailsJobParamsWithTimeout(timeout time.Duration) *DeleteAnalyticsConversationsDetailsJobParams {
	return &DeleteAnalyticsConversationsDetailsJobParams{
		timeout: timeout,
	}
}

// NewDeleteAnalyticsConversationsDetailsJobParamsWithContext creates a new DeleteAnalyticsConversationsDetailsJobParams object
// with the ability to set a context for a request.
func NewDeleteAnalyticsConversationsDetailsJobParamsWithContext(ctx context.Context) *DeleteAnalyticsConversationsDetailsJobParams {
	return &DeleteAnalyticsConversationsDetailsJobParams{
		Context: ctx,
	}
}

// NewDeleteAnalyticsConversationsDetailsJobParamsWithHTTPClient creates a new DeleteAnalyticsConversationsDetailsJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAnalyticsConversationsDetailsJobParamsWithHTTPClient(client *http.Client) *DeleteAnalyticsConversationsDetailsJobParams {
	return &DeleteAnalyticsConversationsDetailsJobParams{
		HTTPClient: client,
	}
}

/*
DeleteAnalyticsConversationsDetailsJobParams contains all the parameters to send to the API endpoint

	for the delete analytics conversations details job operation.

	Typically these are written to a http.Request.
*/
type DeleteAnalyticsConversationsDetailsJobParams struct {

	/* JobID.

	   jobId
	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete analytics conversations details job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAnalyticsConversationsDetailsJobParams) WithDefaults() *DeleteAnalyticsConversationsDetailsJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete analytics conversations details job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAnalyticsConversationsDetailsJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete analytics conversations details job params
func (o *DeleteAnalyticsConversationsDetailsJobParams) WithTimeout(timeout time.Duration) *DeleteAnalyticsConversationsDetailsJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete analytics conversations details job params
func (o *DeleteAnalyticsConversationsDetailsJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete analytics conversations details job params
func (o *DeleteAnalyticsConversationsDetailsJobParams) WithContext(ctx context.Context) *DeleteAnalyticsConversationsDetailsJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete analytics conversations details job params
func (o *DeleteAnalyticsConversationsDetailsJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete analytics conversations details job params
func (o *DeleteAnalyticsConversationsDetailsJobParams) WithHTTPClient(client *http.Client) *DeleteAnalyticsConversationsDetailsJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete analytics conversations details job params
func (o *DeleteAnalyticsConversationsDetailsJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the delete analytics conversations details job params
func (o *DeleteAnalyticsConversationsDetailsJobParams) WithJobID(jobID string) *DeleteAnalyticsConversationsDetailsJobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the delete analytics conversations details job params
func (o *DeleteAnalyticsConversationsDetailsJobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAnalyticsConversationsDetailsJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param jobId
	if err := r.SetPathParam("jobId", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
