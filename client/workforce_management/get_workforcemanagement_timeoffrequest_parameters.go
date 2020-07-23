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

// NewGetWorkforcemanagementTimeoffrequestParams creates a new GetWorkforcemanagementTimeoffrequestParams object
// with the default values initialized.
func NewGetWorkforcemanagementTimeoffrequestParams() *GetWorkforcemanagementTimeoffrequestParams {
	var ()
	return &GetWorkforcemanagementTimeoffrequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementTimeoffrequestParamsWithTimeout creates a new GetWorkforcemanagementTimeoffrequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementTimeoffrequestParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementTimeoffrequestParams {
	var ()
	return &GetWorkforcemanagementTimeoffrequestParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementTimeoffrequestParamsWithContext creates a new GetWorkforcemanagementTimeoffrequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementTimeoffrequestParamsWithContext(ctx context.Context) *GetWorkforcemanagementTimeoffrequestParams {
	var ()
	return &GetWorkforcemanagementTimeoffrequestParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementTimeoffrequestParamsWithHTTPClient creates a new GetWorkforcemanagementTimeoffrequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementTimeoffrequestParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementTimeoffrequestParams {
	var ()
	return &GetWorkforcemanagementTimeoffrequestParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementTimeoffrequestParams contains all the parameters to send to the API endpoint
for the get workforcemanagement timeoffrequest operation typically these are written to a http.Request
*/
type GetWorkforcemanagementTimeoffrequestParams struct {

	/*TimeOffRequestID
	  Time Off Request Id

	*/
	TimeOffRequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement timeoffrequest params
func (o *GetWorkforcemanagementTimeoffrequestParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementTimeoffrequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement timeoffrequest params
func (o *GetWorkforcemanagementTimeoffrequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement timeoffrequest params
func (o *GetWorkforcemanagementTimeoffrequestParams) WithContext(ctx context.Context) *GetWorkforcemanagementTimeoffrequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement timeoffrequest params
func (o *GetWorkforcemanagementTimeoffrequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement timeoffrequest params
func (o *GetWorkforcemanagementTimeoffrequestParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementTimeoffrequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement timeoffrequest params
func (o *GetWorkforcemanagementTimeoffrequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimeOffRequestID adds the timeOffRequestID to the get workforcemanagement timeoffrequest params
func (o *GetWorkforcemanagementTimeoffrequestParams) WithTimeOffRequestID(timeOffRequestID string) *GetWorkforcemanagementTimeoffrequestParams {
	o.SetTimeOffRequestID(timeOffRequestID)
	return o
}

// SetTimeOffRequestID adds the timeOffRequestId to the get workforcemanagement timeoffrequest params
func (o *GetWorkforcemanagementTimeoffrequestParams) SetTimeOffRequestID(timeOffRequestID string) {
	o.TimeOffRequestID = timeOffRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementTimeoffrequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param timeOffRequestId
	if err := r.SetPathParam("timeOffRequestId", o.TimeOffRequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
