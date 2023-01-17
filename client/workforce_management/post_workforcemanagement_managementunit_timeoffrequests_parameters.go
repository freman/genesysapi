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

// NewPostWorkforcemanagementManagementunitTimeoffrequestsParams creates a new PostWorkforcemanagementManagementunitTimeoffrequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkforcemanagementManagementunitTimeoffrequestsParams() *PostWorkforcemanagementManagementunitTimeoffrequestsParams {
	return &PostWorkforcemanagementManagementunitTimeoffrequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementManagementunitTimeoffrequestsParamsWithTimeout creates a new PostWorkforcemanagementManagementunitTimeoffrequestsParams object
// with the ability to set a timeout on a request.
func NewPostWorkforcemanagementManagementunitTimeoffrequestsParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementManagementunitTimeoffrequestsParams {
	return &PostWorkforcemanagementManagementunitTimeoffrequestsParams{
		timeout: timeout,
	}
}

// NewPostWorkforcemanagementManagementunitTimeoffrequestsParamsWithContext creates a new PostWorkforcemanagementManagementunitTimeoffrequestsParams object
// with the ability to set a context for a request.
func NewPostWorkforcemanagementManagementunitTimeoffrequestsParamsWithContext(ctx context.Context) *PostWorkforcemanagementManagementunitTimeoffrequestsParams {
	return &PostWorkforcemanagementManagementunitTimeoffrequestsParams{
		Context: ctx,
	}
}

// NewPostWorkforcemanagementManagementunitTimeoffrequestsParamsWithHTTPClient creates a new PostWorkforcemanagementManagementunitTimeoffrequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkforcemanagementManagementunitTimeoffrequestsParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementManagementunitTimeoffrequestsParams {
	return &PostWorkforcemanagementManagementunitTimeoffrequestsParams{
		HTTPClient: client,
	}
}

/*
PostWorkforcemanagementManagementunitTimeoffrequestsParams contains all the parameters to send to the API endpoint

	for the post workforcemanagement managementunit timeoffrequests operation.

	Typically these are written to a http.Request.
*/
type PostWorkforcemanagementManagementunitTimeoffrequestsParams struct {

	/* Body.

	   body
	*/
	Body *models.CreateAdminTimeOffRequest

	/* ManagementUnitID.

	   The ID of the management unit, or 'mine' for the management unit of the logged-in user.
	*/
	ManagementUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workforcemanagement managementunit timeoffrequests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) WithDefaults() *PostWorkforcemanagementManagementunitTimeoffrequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workforcemanagement managementunit timeoffrequests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workforcemanagement managementunit timeoffrequests params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementManagementunitTimeoffrequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement managementunit timeoffrequests params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement managementunit timeoffrequests params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) WithContext(ctx context.Context) *PostWorkforcemanagementManagementunitTimeoffrequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement managementunit timeoffrequests params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement managementunit timeoffrequests params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementManagementunitTimeoffrequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement managementunit timeoffrequests params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement managementunit timeoffrequests params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) WithBody(body *models.CreateAdminTimeOffRequest) *PostWorkforcemanagementManagementunitTimeoffrequestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement managementunit timeoffrequests params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) SetBody(body *models.CreateAdminTimeOffRequest) {
	o.Body = body
}

// WithManagementUnitID adds the managementUnitID to the post workforcemanagement managementunit timeoffrequests params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) WithManagementUnitID(managementUnitID string) *PostWorkforcemanagementManagementunitTimeoffrequestsParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the post workforcemanagement managementunit timeoffrequests params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
