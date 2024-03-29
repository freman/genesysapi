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
)

// NewGetArchitectSchedulegroupParams creates a new GetArchitectSchedulegroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetArchitectSchedulegroupParams() *GetArchitectSchedulegroupParams {
	return &GetArchitectSchedulegroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectSchedulegroupParamsWithTimeout creates a new GetArchitectSchedulegroupParams object
// with the ability to set a timeout on a request.
func NewGetArchitectSchedulegroupParamsWithTimeout(timeout time.Duration) *GetArchitectSchedulegroupParams {
	return &GetArchitectSchedulegroupParams{
		timeout: timeout,
	}
}

// NewGetArchitectSchedulegroupParamsWithContext creates a new GetArchitectSchedulegroupParams object
// with the ability to set a context for a request.
func NewGetArchitectSchedulegroupParamsWithContext(ctx context.Context) *GetArchitectSchedulegroupParams {
	return &GetArchitectSchedulegroupParams{
		Context: ctx,
	}
}

// NewGetArchitectSchedulegroupParamsWithHTTPClient creates a new GetArchitectSchedulegroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetArchitectSchedulegroupParamsWithHTTPClient(client *http.Client) *GetArchitectSchedulegroupParams {
	return &GetArchitectSchedulegroupParams{
		HTTPClient: client,
	}
}

/*
GetArchitectSchedulegroupParams contains all the parameters to send to the API endpoint

	for the get architect schedulegroup operation.

	Typically these are written to a http.Request.
*/
type GetArchitectSchedulegroupParams struct {

	/* ScheduleGroupID.

	   Schedule group ID
	*/
	ScheduleGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get architect schedulegroup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectSchedulegroupParams) WithDefaults() *GetArchitectSchedulegroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get architect schedulegroup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectSchedulegroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get architect schedulegroup params
func (o *GetArchitectSchedulegroupParams) WithTimeout(timeout time.Duration) *GetArchitectSchedulegroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect schedulegroup params
func (o *GetArchitectSchedulegroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect schedulegroup params
func (o *GetArchitectSchedulegroupParams) WithContext(ctx context.Context) *GetArchitectSchedulegroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect schedulegroup params
func (o *GetArchitectSchedulegroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect schedulegroup params
func (o *GetArchitectSchedulegroupParams) WithHTTPClient(client *http.Client) *GetArchitectSchedulegroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect schedulegroup params
func (o *GetArchitectSchedulegroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScheduleGroupID adds the scheduleGroupID to the get architect schedulegroup params
func (o *GetArchitectSchedulegroupParams) WithScheduleGroupID(scheduleGroupID string) *GetArchitectSchedulegroupParams {
	o.SetScheduleGroupID(scheduleGroupID)
	return o
}

// SetScheduleGroupID adds the scheduleGroupId to the get architect schedulegroup params
func (o *GetArchitectSchedulegroupParams) SetScheduleGroupID(scheduleGroupID string) {
	o.ScheduleGroupID = scheduleGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectSchedulegroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param scheduleGroupId
	if err := r.SetPathParam("scheduleGroupId", o.ScheduleGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
