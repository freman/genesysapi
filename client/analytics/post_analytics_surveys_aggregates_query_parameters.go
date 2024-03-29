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

// NewPostAnalyticsSurveysAggregatesQueryParams creates a new PostAnalyticsSurveysAggregatesQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAnalyticsSurveysAggregatesQueryParams() *PostAnalyticsSurveysAggregatesQueryParams {
	return &PostAnalyticsSurveysAggregatesQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAnalyticsSurveysAggregatesQueryParamsWithTimeout creates a new PostAnalyticsSurveysAggregatesQueryParams object
// with the ability to set a timeout on a request.
func NewPostAnalyticsSurveysAggregatesQueryParamsWithTimeout(timeout time.Duration) *PostAnalyticsSurveysAggregatesQueryParams {
	return &PostAnalyticsSurveysAggregatesQueryParams{
		timeout: timeout,
	}
}

// NewPostAnalyticsSurveysAggregatesQueryParamsWithContext creates a new PostAnalyticsSurveysAggregatesQueryParams object
// with the ability to set a context for a request.
func NewPostAnalyticsSurveysAggregatesQueryParamsWithContext(ctx context.Context) *PostAnalyticsSurveysAggregatesQueryParams {
	return &PostAnalyticsSurveysAggregatesQueryParams{
		Context: ctx,
	}
}

// NewPostAnalyticsSurveysAggregatesQueryParamsWithHTTPClient creates a new PostAnalyticsSurveysAggregatesQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAnalyticsSurveysAggregatesQueryParamsWithHTTPClient(client *http.Client) *PostAnalyticsSurveysAggregatesQueryParams {
	return &PostAnalyticsSurveysAggregatesQueryParams{
		HTTPClient: client,
	}
}

/*
PostAnalyticsSurveysAggregatesQueryParams contains all the parameters to send to the API endpoint

	for the post analytics surveys aggregates query operation.

	Typically these are written to a http.Request.
*/
type PostAnalyticsSurveysAggregatesQueryParams struct {

	/* Body.

	   query
	*/
	Body *models.SurveyAggregationQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post analytics surveys aggregates query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAnalyticsSurveysAggregatesQueryParams) WithDefaults() *PostAnalyticsSurveysAggregatesQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post analytics surveys aggregates query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAnalyticsSurveysAggregatesQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post analytics surveys aggregates query params
func (o *PostAnalyticsSurveysAggregatesQueryParams) WithTimeout(timeout time.Duration) *PostAnalyticsSurveysAggregatesQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post analytics surveys aggregates query params
func (o *PostAnalyticsSurveysAggregatesQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post analytics surveys aggregates query params
func (o *PostAnalyticsSurveysAggregatesQueryParams) WithContext(ctx context.Context) *PostAnalyticsSurveysAggregatesQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post analytics surveys aggregates query params
func (o *PostAnalyticsSurveysAggregatesQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post analytics surveys aggregates query params
func (o *PostAnalyticsSurveysAggregatesQueryParams) WithHTTPClient(client *http.Client) *PostAnalyticsSurveysAggregatesQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post analytics surveys aggregates query params
func (o *PostAnalyticsSurveysAggregatesQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post analytics surveys aggregates query params
func (o *PostAnalyticsSurveysAggregatesQueryParams) WithBody(body *models.SurveyAggregationQuery) *PostAnalyticsSurveysAggregatesQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post analytics surveys aggregates query params
func (o *PostAnalyticsSurveysAggregatesQueryParams) SetBody(body *models.SurveyAggregationQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAnalyticsSurveysAggregatesQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
