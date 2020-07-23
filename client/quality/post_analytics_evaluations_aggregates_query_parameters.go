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

// NewPostAnalyticsEvaluationsAggregatesQueryParams creates a new PostAnalyticsEvaluationsAggregatesQueryParams object
// with the default values initialized.
func NewPostAnalyticsEvaluationsAggregatesQueryParams() *PostAnalyticsEvaluationsAggregatesQueryParams {
	var ()
	return &PostAnalyticsEvaluationsAggregatesQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAnalyticsEvaluationsAggregatesQueryParamsWithTimeout creates a new PostAnalyticsEvaluationsAggregatesQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAnalyticsEvaluationsAggregatesQueryParamsWithTimeout(timeout time.Duration) *PostAnalyticsEvaluationsAggregatesQueryParams {
	var ()
	return &PostAnalyticsEvaluationsAggregatesQueryParams{

		timeout: timeout,
	}
}

// NewPostAnalyticsEvaluationsAggregatesQueryParamsWithContext creates a new PostAnalyticsEvaluationsAggregatesQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAnalyticsEvaluationsAggregatesQueryParamsWithContext(ctx context.Context) *PostAnalyticsEvaluationsAggregatesQueryParams {
	var ()
	return &PostAnalyticsEvaluationsAggregatesQueryParams{

		Context: ctx,
	}
}

// NewPostAnalyticsEvaluationsAggregatesQueryParamsWithHTTPClient creates a new PostAnalyticsEvaluationsAggregatesQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAnalyticsEvaluationsAggregatesQueryParamsWithHTTPClient(client *http.Client) *PostAnalyticsEvaluationsAggregatesQueryParams {
	var ()
	return &PostAnalyticsEvaluationsAggregatesQueryParams{
		HTTPClient: client,
	}
}

/*PostAnalyticsEvaluationsAggregatesQueryParams contains all the parameters to send to the API endpoint
for the post analytics evaluations aggregates query operation typically these are written to a http.Request
*/
type PostAnalyticsEvaluationsAggregatesQueryParams struct {

	/*Body
	  query

	*/
	Body *models.EvaluationAggregationQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post analytics evaluations aggregates query params
func (o *PostAnalyticsEvaluationsAggregatesQueryParams) WithTimeout(timeout time.Duration) *PostAnalyticsEvaluationsAggregatesQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post analytics evaluations aggregates query params
func (o *PostAnalyticsEvaluationsAggregatesQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post analytics evaluations aggregates query params
func (o *PostAnalyticsEvaluationsAggregatesQueryParams) WithContext(ctx context.Context) *PostAnalyticsEvaluationsAggregatesQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post analytics evaluations aggregates query params
func (o *PostAnalyticsEvaluationsAggregatesQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post analytics evaluations aggregates query params
func (o *PostAnalyticsEvaluationsAggregatesQueryParams) WithHTTPClient(client *http.Client) *PostAnalyticsEvaluationsAggregatesQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post analytics evaluations aggregates query params
func (o *PostAnalyticsEvaluationsAggregatesQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post analytics evaluations aggregates query params
func (o *PostAnalyticsEvaluationsAggregatesQueryParams) WithBody(body *models.EvaluationAggregationQuery) *PostAnalyticsEvaluationsAggregatesQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post analytics evaluations aggregates query params
func (o *PostAnalyticsEvaluationsAggregatesQueryParams) SetBody(body *models.EvaluationAggregationQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAnalyticsEvaluationsAggregatesQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
