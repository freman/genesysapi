// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewPutArchitectScheduleParams creates a new PutArchitectScheduleParams object
// with the default values initialized.
func NewPutArchitectScheduleParams() *PutArchitectScheduleParams {
	var ()
	return &PutArchitectScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutArchitectScheduleParamsWithTimeout creates a new PutArchitectScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutArchitectScheduleParamsWithTimeout(timeout time.Duration) *PutArchitectScheduleParams {
	var ()
	return &PutArchitectScheduleParams{

		timeout: timeout,
	}
}

// NewPutArchitectScheduleParamsWithContext creates a new PutArchitectScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutArchitectScheduleParamsWithContext(ctx context.Context) *PutArchitectScheduleParams {
	var ()
	return &PutArchitectScheduleParams{

		Context: ctx,
	}
}

// NewPutArchitectScheduleParamsWithHTTPClient creates a new PutArchitectScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutArchitectScheduleParamsWithHTTPClient(client *http.Client) *PutArchitectScheduleParams {
	var ()
	return &PutArchitectScheduleParams{
		HTTPClient: client,
	}
}

/*PutArchitectScheduleParams contains all the parameters to send to the API endpoint
for the put architect schedule operation typically these are written to a http.Request
*/
type PutArchitectScheduleParams struct {

	/*Body*/
	Body *models.Schedule
	/*ScheduleID
	  Schedule ID

	*/
	ScheduleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put architect schedule params
func (o *PutArchitectScheduleParams) WithTimeout(timeout time.Duration) *PutArchitectScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put architect schedule params
func (o *PutArchitectScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put architect schedule params
func (o *PutArchitectScheduleParams) WithContext(ctx context.Context) *PutArchitectScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put architect schedule params
func (o *PutArchitectScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put architect schedule params
func (o *PutArchitectScheduleParams) WithHTTPClient(client *http.Client) *PutArchitectScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put architect schedule params
func (o *PutArchitectScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put architect schedule params
func (o *PutArchitectScheduleParams) WithBody(body *models.Schedule) *PutArchitectScheduleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put architect schedule params
func (o *PutArchitectScheduleParams) SetBody(body *models.Schedule) {
	o.Body = body
}

// WithScheduleID adds the scheduleID to the put architect schedule params
func (o *PutArchitectScheduleParams) WithScheduleID(scheduleID string) *PutArchitectScheduleParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the put architect schedule params
func (o *PutArchitectScheduleParams) SetScheduleID(scheduleID string) {
	o.ScheduleID = scheduleID
}

// WriteToRequest writes these params to a swagger request
func (o *PutArchitectScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param scheduleId
	if err := r.SetPathParam("scheduleId", o.ScheduleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
