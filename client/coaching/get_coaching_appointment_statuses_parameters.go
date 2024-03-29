// Code generated by go-swagger; DO NOT EDIT.

package coaching

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
	"github.com/go-openapi/swag"
)

// NewGetCoachingAppointmentStatusesParams creates a new GetCoachingAppointmentStatusesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCoachingAppointmentStatusesParams() *GetCoachingAppointmentStatusesParams {
	return &GetCoachingAppointmentStatusesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCoachingAppointmentStatusesParamsWithTimeout creates a new GetCoachingAppointmentStatusesParams object
// with the ability to set a timeout on a request.
func NewGetCoachingAppointmentStatusesParamsWithTimeout(timeout time.Duration) *GetCoachingAppointmentStatusesParams {
	return &GetCoachingAppointmentStatusesParams{
		timeout: timeout,
	}
}

// NewGetCoachingAppointmentStatusesParamsWithContext creates a new GetCoachingAppointmentStatusesParams object
// with the ability to set a context for a request.
func NewGetCoachingAppointmentStatusesParamsWithContext(ctx context.Context) *GetCoachingAppointmentStatusesParams {
	return &GetCoachingAppointmentStatusesParams{
		Context: ctx,
	}
}

// NewGetCoachingAppointmentStatusesParamsWithHTTPClient creates a new GetCoachingAppointmentStatusesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCoachingAppointmentStatusesParamsWithHTTPClient(client *http.Client) *GetCoachingAppointmentStatusesParams {
	return &GetCoachingAppointmentStatusesParams{
		HTTPClient: client,
	}
}

/*
GetCoachingAppointmentStatusesParams contains all the parameters to send to the API endpoint

	for the get coaching appointment statuses operation.

	Typically these are written to a http.Request.
*/
type GetCoachingAppointmentStatusesParams struct {

	/* AppointmentID.

	   The ID of the coaching appointment.
	*/
	AppointmentID string

	/* PageNumber.

	   Page number

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   Page size

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get coaching appointment statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoachingAppointmentStatusesParams) WithDefaults() *GetCoachingAppointmentStatusesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get coaching appointment statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoachingAppointmentStatusesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetCoachingAppointmentStatusesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get coaching appointment statuses params
func (o *GetCoachingAppointmentStatusesParams) WithTimeout(timeout time.Duration) *GetCoachingAppointmentStatusesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get coaching appointment statuses params
func (o *GetCoachingAppointmentStatusesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get coaching appointment statuses params
func (o *GetCoachingAppointmentStatusesParams) WithContext(ctx context.Context) *GetCoachingAppointmentStatusesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get coaching appointment statuses params
func (o *GetCoachingAppointmentStatusesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get coaching appointment statuses params
func (o *GetCoachingAppointmentStatusesParams) WithHTTPClient(client *http.Client) *GetCoachingAppointmentStatusesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get coaching appointment statuses params
func (o *GetCoachingAppointmentStatusesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppointmentID adds the appointmentID to the get coaching appointment statuses params
func (o *GetCoachingAppointmentStatusesParams) WithAppointmentID(appointmentID string) *GetCoachingAppointmentStatusesParams {
	o.SetAppointmentID(appointmentID)
	return o
}

// SetAppointmentID adds the appointmentId to the get coaching appointment statuses params
func (o *GetCoachingAppointmentStatusesParams) SetAppointmentID(appointmentID string) {
	o.AppointmentID = appointmentID
}

// WithPageNumber adds the pageNumber to the get coaching appointment statuses params
func (o *GetCoachingAppointmentStatusesParams) WithPageNumber(pageNumber *int32) *GetCoachingAppointmentStatusesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get coaching appointment statuses params
func (o *GetCoachingAppointmentStatusesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get coaching appointment statuses params
func (o *GetCoachingAppointmentStatusesParams) WithPageSize(pageSize *int32) *GetCoachingAppointmentStatusesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get coaching appointment statuses params
func (o *GetCoachingAppointmentStatusesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetCoachingAppointmentStatusesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appointmentId
	if err := r.SetPathParam("appointmentId", o.AppointmentID); err != nil {
		return err
	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
