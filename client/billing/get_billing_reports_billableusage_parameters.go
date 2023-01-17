// Code generated by go-swagger; DO NOT EDIT.

package billing

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

// NewGetBillingReportsBillableusageParams creates a new GetBillingReportsBillableusageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBillingReportsBillableusageParams() *GetBillingReportsBillableusageParams {
	return &GetBillingReportsBillableusageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBillingReportsBillableusageParamsWithTimeout creates a new GetBillingReportsBillableusageParams object
// with the ability to set a timeout on a request.
func NewGetBillingReportsBillableusageParamsWithTimeout(timeout time.Duration) *GetBillingReportsBillableusageParams {
	return &GetBillingReportsBillableusageParams{
		timeout: timeout,
	}
}

// NewGetBillingReportsBillableusageParamsWithContext creates a new GetBillingReportsBillableusageParams object
// with the ability to set a context for a request.
func NewGetBillingReportsBillableusageParamsWithContext(ctx context.Context) *GetBillingReportsBillableusageParams {
	return &GetBillingReportsBillableusageParams{
		Context: ctx,
	}
}

// NewGetBillingReportsBillableusageParamsWithHTTPClient creates a new GetBillingReportsBillableusageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBillingReportsBillableusageParamsWithHTTPClient(client *http.Client) *GetBillingReportsBillableusageParams {
	return &GetBillingReportsBillableusageParams{
		HTTPClient: client,
	}
}

/*
GetBillingReportsBillableusageParams contains all the parameters to send to the API endpoint

	for the get billing reports billableusage operation.

	Typically these are written to a http.Request.
*/
type GetBillingReportsBillableusageParams struct {

	/* EndDate.

	   The period end date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z

	   Format: date-time
	*/
	EndDate strfmt.DateTime

	/* StartDate.

	   The period start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z

	   Format: date-time
	*/
	StartDate strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get billing reports billableusage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBillingReportsBillableusageParams) WithDefaults() *GetBillingReportsBillableusageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get billing reports billableusage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBillingReportsBillableusageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get billing reports billableusage params
func (o *GetBillingReportsBillableusageParams) WithTimeout(timeout time.Duration) *GetBillingReportsBillableusageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get billing reports billableusage params
func (o *GetBillingReportsBillableusageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get billing reports billableusage params
func (o *GetBillingReportsBillableusageParams) WithContext(ctx context.Context) *GetBillingReportsBillableusageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get billing reports billableusage params
func (o *GetBillingReportsBillableusageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get billing reports billableusage params
func (o *GetBillingReportsBillableusageParams) WithHTTPClient(client *http.Client) *GetBillingReportsBillableusageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get billing reports billableusage params
func (o *GetBillingReportsBillableusageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the get billing reports billableusage params
func (o *GetBillingReportsBillableusageParams) WithEndDate(endDate strfmt.DateTime) *GetBillingReportsBillableusageParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get billing reports billableusage params
func (o *GetBillingReportsBillableusageParams) SetEndDate(endDate strfmt.DateTime) {
	o.EndDate = endDate
}

// WithStartDate adds the startDate to the get billing reports billableusage params
func (o *GetBillingReportsBillableusageParams) WithStartDate(startDate strfmt.DateTime) *GetBillingReportsBillableusageParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get billing reports billableusage params
func (o *GetBillingReportsBillableusageParams) SetStartDate(startDate strfmt.DateTime) {
	o.StartDate = startDate
}

// WriteToRequest writes these params to a swagger request
func (o *GetBillingReportsBillableusageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param endDate
	qrEndDate := o.EndDate
	qEndDate := qrEndDate.String()
	if qEndDate != "" {

		if err := r.SetQueryParam("endDate", qEndDate); err != nil {
			return err
		}
	}

	// query param startDate
	qrStartDate := o.StartDate
	qStartDate := qrStartDate.String()
	if qStartDate != "" {

		if err := r.SetQueryParam("startDate", qStartDate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
