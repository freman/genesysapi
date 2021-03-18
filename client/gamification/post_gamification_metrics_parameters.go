// Code generated by go-swagger; DO NOT EDIT.

package gamification

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

// NewPostGamificationMetricsParams creates a new PostGamificationMetricsParams object
// with the default values initialized.
func NewPostGamificationMetricsParams() *PostGamificationMetricsParams {
	var ()
	return &PostGamificationMetricsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostGamificationMetricsParamsWithTimeout creates a new PostGamificationMetricsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostGamificationMetricsParamsWithTimeout(timeout time.Duration) *PostGamificationMetricsParams {
	var ()
	return &PostGamificationMetricsParams{

		timeout: timeout,
	}
}

// NewPostGamificationMetricsParamsWithContext creates a new PostGamificationMetricsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostGamificationMetricsParamsWithContext(ctx context.Context) *PostGamificationMetricsParams {
	var ()
	return &PostGamificationMetricsParams{

		Context: ctx,
	}
}

// NewPostGamificationMetricsParamsWithHTTPClient creates a new PostGamificationMetricsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostGamificationMetricsParamsWithHTTPClient(client *http.Client) *PostGamificationMetricsParams {
	var ()
	return &PostGamificationMetricsParams{
		HTTPClient: client,
	}
}

/*PostGamificationMetricsParams contains all the parameters to send to the API endpoint
for the post gamification metrics operation typically these are written to a http.Request
*/
type PostGamificationMetricsParams struct {

	/*Body
	  Metric

	*/
	Body *models.Metric

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post gamification metrics params
func (o *PostGamificationMetricsParams) WithTimeout(timeout time.Duration) *PostGamificationMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post gamification metrics params
func (o *PostGamificationMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post gamification metrics params
func (o *PostGamificationMetricsParams) WithContext(ctx context.Context) *PostGamificationMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post gamification metrics params
func (o *PostGamificationMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post gamification metrics params
func (o *PostGamificationMetricsParams) WithHTTPClient(client *http.Client) *PostGamificationMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post gamification metrics params
func (o *PostGamificationMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post gamification metrics params
func (o *PostGamificationMetricsParams) WithBody(body *models.Metric) *PostGamificationMetricsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post gamification metrics params
func (o *PostGamificationMetricsParams) SetBody(body *models.Metric) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostGamificationMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
