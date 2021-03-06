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

// NewPostAnalyticsReportingScheduleRunreportParams creates a new PostAnalyticsReportingScheduleRunreportParams object
// with the default values initialized.
func NewPostAnalyticsReportingScheduleRunreportParams() *PostAnalyticsReportingScheduleRunreportParams {
	var ()
	return &PostAnalyticsReportingScheduleRunreportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAnalyticsReportingScheduleRunreportParamsWithTimeout creates a new PostAnalyticsReportingScheduleRunreportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAnalyticsReportingScheduleRunreportParamsWithTimeout(timeout time.Duration) *PostAnalyticsReportingScheduleRunreportParams {
	var ()
	return &PostAnalyticsReportingScheduleRunreportParams{

		timeout: timeout,
	}
}

// NewPostAnalyticsReportingScheduleRunreportParamsWithContext creates a new PostAnalyticsReportingScheduleRunreportParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAnalyticsReportingScheduleRunreportParamsWithContext(ctx context.Context) *PostAnalyticsReportingScheduleRunreportParams {
	var ()
	return &PostAnalyticsReportingScheduleRunreportParams{

		Context: ctx,
	}
}

// NewPostAnalyticsReportingScheduleRunreportParamsWithHTTPClient creates a new PostAnalyticsReportingScheduleRunreportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAnalyticsReportingScheduleRunreportParamsWithHTTPClient(client *http.Client) *PostAnalyticsReportingScheduleRunreportParams {
	var ()
	return &PostAnalyticsReportingScheduleRunreportParams{
		HTTPClient: client,
	}
}

/*PostAnalyticsReportingScheduleRunreportParams contains all the parameters to send to the API endpoint
for the post analytics reporting schedule runreport operation typically these are written to a http.Request
*/
type PostAnalyticsReportingScheduleRunreportParams struct {

	/*ScheduleID
	  Schedule ID

	*/
	ScheduleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post analytics reporting schedule runreport params
func (o *PostAnalyticsReportingScheduleRunreportParams) WithTimeout(timeout time.Duration) *PostAnalyticsReportingScheduleRunreportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post analytics reporting schedule runreport params
func (o *PostAnalyticsReportingScheduleRunreportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post analytics reporting schedule runreport params
func (o *PostAnalyticsReportingScheduleRunreportParams) WithContext(ctx context.Context) *PostAnalyticsReportingScheduleRunreportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post analytics reporting schedule runreport params
func (o *PostAnalyticsReportingScheduleRunreportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post analytics reporting schedule runreport params
func (o *PostAnalyticsReportingScheduleRunreportParams) WithHTTPClient(client *http.Client) *PostAnalyticsReportingScheduleRunreportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post analytics reporting schedule runreport params
func (o *PostAnalyticsReportingScheduleRunreportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScheduleID adds the scheduleID to the post analytics reporting schedule runreport params
func (o *PostAnalyticsReportingScheduleRunreportParams) WithScheduleID(scheduleID string) *PostAnalyticsReportingScheduleRunreportParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the post analytics reporting schedule runreport params
func (o *PostAnalyticsReportingScheduleRunreportParams) SetScheduleID(scheduleID string) {
	o.ScheduleID = scheduleID
}

// WriteToRequest writes these params to a swagger request
func (o *PostAnalyticsReportingScheduleRunreportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param scheduleId
	if err := r.SetPathParam("scheduleId", o.ScheduleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
