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

// NewPostAnalyticsReportingSchedulesParams creates a new PostAnalyticsReportingSchedulesParams object
// with the default values initialized.
func NewPostAnalyticsReportingSchedulesParams() *PostAnalyticsReportingSchedulesParams {
	var ()
	return &PostAnalyticsReportingSchedulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAnalyticsReportingSchedulesParamsWithTimeout creates a new PostAnalyticsReportingSchedulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAnalyticsReportingSchedulesParamsWithTimeout(timeout time.Duration) *PostAnalyticsReportingSchedulesParams {
	var ()
	return &PostAnalyticsReportingSchedulesParams{

		timeout: timeout,
	}
}

// NewPostAnalyticsReportingSchedulesParamsWithContext creates a new PostAnalyticsReportingSchedulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAnalyticsReportingSchedulesParamsWithContext(ctx context.Context) *PostAnalyticsReportingSchedulesParams {
	var ()
	return &PostAnalyticsReportingSchedulesParams{

		Context: ctx,
	}
}

// NewPostAnalyticsReportingSchedulesParamsWithHTTPClient creates a new PostAnalyticsReportingSchedulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAnalyticsReportingSchedulesParamsWithHTTPClient(client *http.Client) *PostAnalyticsReportingSchedulesParams {
	var ()
	return &PostAnalyticsReportingSchedulesParams{
		HTTPClient: client,
	}
}

/*PostAnalyticsReportingSchedulesParams contains all the parameters to send to the API endpoint
for the post analytics reporting schedules operation typically these are written to a http.Request
*/
type PostAnalyticsReportingSchedulesParams struct {

	/*Body
	  ReportSchedule

	*/
	Body *models.ReportSchedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post analytics reporting schedules params
func (o *PostAnalyticsReportingSchedulesParams) WithTimeout(timeout time.Duration) *PostAnalyticsReportingSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post analytics reporting schedules params
func (o *PostAnalyticsReportingSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post analytics reporting schedules params
func (o *PostAnalyticsReportingSchedulesParams) WithContext(ctx context.Context) *PostAnalyticsReportingSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post analytics reporting schedules params
func (o *PostAnalyticsReportingSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post analytics reporting schedules params
func (o *PostAnalyticsReportingSchedulesParams) WithHTTPClient(client *http.Client) *PostAnalyticsReportingSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post analytics reporting schedules params
func (o *PostAnalyticsReportingSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post analytics reporting schedules params
func (o *PostAnalyticsReportingSchedulesParams) WithBody(body *models.ReportSchedule) *PostAnalyticsReportingSchedulesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post analytics reporting schedules params
func (o *PostAnalyticsReportingSchedulesParams) SetBody(body *models.ReportSchedule) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAnalyticsReportingSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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