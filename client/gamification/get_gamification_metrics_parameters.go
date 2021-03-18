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
)

// NewGetGamificationMetricsParams creates a new GetGamificationMetricsParams object
// with the default values initialized.
func NewGetGamificationMetricsParams() *GetGamificationMetricsParams {
	var ()
	return &GetGamificationMetricsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGamificationMetricsParamsWithTimeout creates a new GetGamificationMetricsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGamificationMetricsParamsWithTimeout(timeout time.Duration) *GetGamificationMetricsParams {
	var ()
	return &GetGamificationMetricsParams{

		timeout: timeout,
	}
}

// NewGetGamificationMetricsParamsWithContext creates a new GetGamificationMetricsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGamificationMetricsParamsWithContext(ctx context.Context) *GetGamificationMetricsParams {
	var ()
	return &GetGamificationMetricsParams{

		Context: ctx,
	}
}

// NewGetGamificationMetricsParamsWithHTTPClient creates a new GetGamificationMetricsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGamificationMetricsParamsWithHTTPClient(client *http.Client) *GetGamificationMetricsParams {
	var ()
	return &GetGamificationMetricsParams{
		HTTPClient: client,
	}
}

/*GetGamificationMetricsParams contains all the parameters to send to the API endpoint
for the get gamification metrics operation typically these are written to a http.Request
*/
type GetGamificationMetricsParams struct {

	/*PerformanceProfileID
	  The profile id of the metrics you are trying to retrieve. The DEFAULT profile is used if nothing is given.

	*/
	PerformanceProfileID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gamification metrics params
func (o *GetGamificationMetricsParams) WithTimeout(timeout time.Duration) *GetGamificationMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gamification metrics params
func (o *GetGamificationMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gamification metrics params
func (o *GetGamificationMetricsParams) WithContext(ctx context.Context) *GetGamificationMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gamification metrics params
func (o *GetGamificationMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gamification metrics params
func (o *GetGamificationMetricsParams) WithHTTPClient(client *http.Client) *GetGamificationMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gamification metrics params
func (o *GetGamificationMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPerformanceProfileID adds the performanceProfileID to the get gamification metrics params
func (o *GetGamificationMetricsParams) WithPerformanceProfileID(performanceProfileID *string) *GetGamificationMetricsParams {
	o.SetPerformanceProfileID(performanceProfileID)
	return o
}

// SetPerformanceProfileID adds the performanceProfileId to the get gamification metrics params
func (o *GetGamificationMetricsParams) SetPerformanceProfileID(performanceProfileID *string) {
	o.PerformanceProfileID = performanceProfileID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGamificationMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PerformanceProfileID != nil {

		// query param performance profile id
		var qrPerformanceProfileID string
		if o.PerformanceProfileID != nil {
			qrPerformanceProfileID = *o.PerformanceProfileID
		}
		qPerformanceProfileID := qrPerformanceProfileID
		if qPerformanceProfileID != "" {
			if err := r.SetQueryParam("performance profile id", qPerformanceProfileID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}