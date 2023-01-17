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

	"github.com/freman/genesysapi/models"
)

// NewPostWorkforcemanagementManagementunitSchedulesSearchParams creates a new PostWorkforcemanagementManagementunitSchedulesSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementManagementunitSchedulesSearchParams() *PostWorkforcemanagementManagementunitSchedulesSearchParams {
	return &PostWorkforcemanagementManagementunitSchedulesSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementManagementunitSchedulesSearchParamsWithTimeout creates a new PostWorkforcemanagementManagementunitSchedulesSearchParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementManagementunitSchedulesSearchParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementManagementunitSchedulesSearchParams {
	return &PostWorkforcemanagementManagementunitSchedulesSearchParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementManagementunitSchedulesSearchParamsWithContext creates a new PostWorkforcemanagementManagementunitSchedulesSearchParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementManagementunitSchedulesSearchParamsWithContext(ctx context.Context) *PostWorkforcemanagementManagementunitSchedulesSearchParams {
	return &PostWorkforcemanagementManagementunitSchedulesSearchParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementManagementunitSchedulesSearchParamsWithHTTPClient creates a new PostWorkforcemanagementManagementunitSchedulesSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementManagementunitSchedulesSearchParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementManagementunitSchedulesSearchParams {
	return &PostWorkforcemanagementManagementunitSchedulesSearchParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementManagementunitSchedulesSearchParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement managementunit schedules search operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementManagementunitSchedulesSearchParams struct {

	/* Body.

	   body
	*/
	Body *models.UserListScheduleRequestBody

	/* ManagementUnitID.

	   The ID of the management unit, or 'mine' for the management unit of the logged-in user.
	*/
	ManagementUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workforcemanagement managementunit schedules search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) WithDefaults() *PostWorkforcemanagementManagementunitSchedulesSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement managementunit schedules search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workforcemanagement managementunit schedules search params
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementManagementunitSchedulesSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement managementunit schedules search params
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement managementunit schedules search params
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) WithContext(ctx context.Context) *PostWorkforcemanagementManagementunitSchedulesSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement managementunit schedules search params
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement managementunit schedules search params
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementManagementunitSchedulesSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement managementunit schedules search params
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement managementunit schedules search params
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) WithBody(body *models.UserListScheduleRequestBody) *PostWorkforcemanagementManagementunitSchedulesSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement managementunit schedules search params
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) SetBody(body *models.UserListScheduleRequestBody) {
	o.Body = body
}

// WithManagementUnitID adds the managementUnitID to the post workforcemanagement managementunit schedules search params
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) WithManagementUnitID(managementUnitID string) *PostWorkforcemanagementManagementunitSchedulesSearchParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the post workforcemanagement managementunit schedules search params
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementManagementunitSchedulesSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param managementUnitId
	if err := r.SetPathParam("managementUnitId", o.ManagementUnitID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
