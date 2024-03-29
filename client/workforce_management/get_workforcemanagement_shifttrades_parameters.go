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

// NewGetWorkforcemanagementShifttradesParams creates a new GetWorkforcemanagementShifttradesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkforcemanagementShifttradesParams() *GetWorkforcemanagementShifttradesParams {
	return &GetWorkforcemanagementShifttradesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementShifttradesParamsWithTimeout creates a new GetWorkforcemanagementShifttradesParams object
// with the ability to set a timeout on a request.
func NewGetWorkforcemanagementShifttradesParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementShifttradesParams {
	return &GetWorkforcemanagementShifttradesParams{
		timeout: timeout,
	}
}

// NewGetWorkforcemanagementShifttradesParamsWithContext creates a new GetWorkforcemanagementShifttradesParams object
// with the ability to set a context for a request.
func NewGetWorkforcemanagementShifttradesParamsWithContext(ctx context.Context) *GetWorkforcemanagementShifttradesParams {
	return &GetWorkforcemanagementShifttradesParams{
		Context: ctx,
	}
}

// NewGetWorkforcemanagementShifttradesParamsWithHTTPClient creates a new GetWorkforcemanagementShifttradesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkforcemanagementShifttradesParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementShifttradesParams {
	return &GetWorkforcemanagementShifttradesParams{
		HTTPClient: client,
	}
}

/*
GetWorkforcemanagementShifttradesParams contains all the parameters to send to the API endpoint

	for the get workforcemanagement shifttrades operation.

	Typically these are written to a http.Request.
*/
type GetWorkforcemanagementShifttradesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workforcemanagement shifttrades params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementShifttradesParams) WithDefaults() *GetWorkforcemanagementShifttradesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workforcemanagement shifttrades params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementShifttradesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workforcemanagement shifttrades params
func (o *GetWorkforcemanagementShifttradesParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementShifttradesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement shifttrades params
func (o *GetWorkforcemanagementShifttradesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement shifttrades params
func (o *GetWorkforcemanagementShifttradesParams) WithContext(ctx context.Context) *GetWorkforcemanagementShifttradesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement shifttrades params
func (o *GetWorkforcemanagementShifttradesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement shifttrades params
func (o *GetWorkforcemanagementShifttradesParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementShifttradesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement shifttrades params
func (o *GetWorkforcemanagementShifttradesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementShifttradesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
