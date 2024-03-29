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

	"github.com/freman/genesysapi/models"
)

// NewPostAnalyticsActionsAggregatesQueryParams creates a new PostAnalyticsActionsAggregatesQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAnalyticsActionsAggregatesQueryParams() *PostAnalyticsActionsAggregatesQueryParams {
	return &PostAnalyticsActionsAggregatesQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAnalyticsActionsAggregatesQueryParamsWithTimeout creates a new PostAnalyticsActionsAggregatesQueryParams object
// with the ability to set a timeout on a request.
func NewPostAnalyticsActionsAggregatesQueryParamsWithTimeout(timeout time.Duration) *PostAnalyticsActionsAggregatesQueryParams {
	return &PostAnalyticsActionsAggregatesQueryParams{
		timeout: timeout,
	}
}

// NewPostAnalyticsActionsAggregatesQueryParamsWithContext creates a new PostAnalyticsActionsAggregatesQueryParams object
// with the ability to set a context for a request.
func NewPostAnalyticsActionsAggregatesQueryParamsWithContext(ctx context.Context) *PostAnalyticsActionsAggregatesQueryParams {
	return &PostAnalyticsActionsAggregatesQueryParams{
		Context: ctx,
	}
}

// NewPostAnalyticsActionsAggregatesQueryParamsWithHTTPClient creates a new PostAnalyticsActionsAggregatesQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAnalyticsActionsAggregatesQueryParamsWithHTTPClient(client *http.Client) *PostAnalyticsActionsAggregatesQueryParams {
	return &PostAnalyticsActionsAggregatesQueryParams{
		HTTPClient: client,
	}
}

/*
PostAnalyticsActionsAggregatesQueryParams contains all the parameters to send to the API endpoint

	for the post analytics actions aggregates query operation.

	Typically these are written to a http.Request.
*/
type PostAnalyticsActionsAggregatesQueryParams struct {

	/* Body.

	   query
	*/
	Body *models.ActionAggregationQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post analytics actions aggregates query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAnalyticsActionsAggregatesQueryParams) WithDefaults() *PostAnalyticsActionsAggregatesQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post analytics actions aggregates query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAnalyticsActionsAggregatesQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post analytics actions aggregates query params
func (o *PostAnalyticsActionsAggregatesQueryParams) WithTimeout(timeout time.Duration) *PostAnalyticsActionsAggregatesQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post analytics actions aggregates query params
func (o *PostAnalyticsActionsAggregatesQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post analytics actions aggregates query params
func (o *PostAnalyticsActionsAggregatesQueryParams) WithContext(ctx context.Context) *PostAnalyticsActionsAggregatesQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post analytics actions aggregates query params
func (o *PostAnalyticsActionsAggregatesQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post analytics actions aggregates query params
func (o *PostAnalyticsActionsAggregatesQueryParams) WithHTTPClient(client *http.Client) *PostAnalyticsActionsAggregatesQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post analytics actions aggregates query params
func (o *PostAnalyticsActionsAggregatesQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post analytics actions aggregates query params
func (o *PostAnalyticsActionsAggregatesQueryParams) WithBody(body *models.ActionAggregationQuery) *PostAnalyticsActionsAggregatesQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post analytics actions aggregates query params
func (o *PostAnalyticsActionsAggregatesQueryParams) SetBody(body *models.ActionAggregationQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAnalyticsActionsAggregatesQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
