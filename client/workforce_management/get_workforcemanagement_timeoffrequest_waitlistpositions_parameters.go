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

// NewGetWorkforcemanagementTimeoffrequestWaitlistpositionsParams creates a new GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams object
// with the default values initialized.
func NewGetWorkforcemanagementTimeoffrequestWaitlistpositionsParams() *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams {
	var ()
	return &GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementTimeoffrequestWaitlistpositionsParamsWithTimeout creates a new GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementTimeoffrequestWaitlistpositionsParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams {
	var ()
	return &GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementTimeoffrequestWaitlistpositionsParamsWithContext creates a new GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementTimeoffrequestWaitlistpositionsParamsWithContext(ctx context.Context) *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams {
	var ()
	return &GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementTimeoffrequestWaitlistpositionsParamsWithHTTPClient creates a new GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementTimeoffrequestWaitlistpositionsParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams {
	var ()
	return &GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams contains all the parameters to send to the API endpoint
for the get workforcemanagement timeoffrequest waitlistpositions operation typically these are written to a http.Request
*/
type GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams struct {

	/*TimeOffRequestID
	  The ID of the time off request

	*/
	TimeOffRequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement timeoffrequest waitlistpositions params
func (o *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement timeoffrequest waitlistpositions params
func (o *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement timeoffrequest waitlistpositions params
func (o *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams) WithContext(ctx context.Context) *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement timeoffrequest waitlistpositions params
func (o *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement timeoffrequest waitlistpositions params
func (o *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement timeoffrequest waitlistpositions params
func (o *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimeOffRequestID adds the timeOffRequestID to the get workforcemanagement timeoffrequest waitlistpositions params
func (o *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams) WithTimeOffRequestID(timeOffRequestID string) *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams {
	o.SetTimeOffRequestID(timeOffRequestID)
	return o
}

// SetTimeOffRequestID adds the timeOffRequestId to the get workforcemanagement timeoffrequest waitlistpositions params
func (o *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams) SetTimeOffRequestID(timeOffRequestID string) {
	o.TimeOffRequestID = timeOffRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementTimeoffrequestWaitlistpositionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
