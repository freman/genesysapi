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

// NewDeleteWorkforcemanagementManagementunitParams creates a new DeleteWorkforcemanagementManagementunitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteWorkforcemanagementManagementunitParams() *DeleteWorkforcemanagementManagementunitParams {
	return &DeleteWorkforcemanagementManagementunitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkforcemanagementManagementunitParamsWithTimeout creates a new DeleteWorkforcemanagementManagementunitParams object
// with the ability to set a timeout on a request.
func NewDeleteWorkforcemanagementManagementunitParamsWithTimeout(timeout time.Duration) *DeleteWorkforcemanagementManagementunitParams {
	return &DeleteWorkforcemanagementManagementunitParams{
		timeout: timeout,
	}
}

// NewDeleteWorkforcemanagementManagementunitParamsWithContext creates a new DeleteWorkforcemanagementManagementunitParams object
// with the ability to set a context for a request.
func NewDeleteWorkforcemanagementManagementunitParamsWithContext(ctx context.Context) *DeleteWorkforcemanagementManagementunitParams {
	return &DeleteWorkforcemanagementManagementunitParams{
		Context: ctx,
	}
}

// NewDeleteWorkforcemanagementManagementunitParamsWithHTTPClient creates a new DeleteWorkforcemanagementManagementunitParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteWorkforcemanagementManagementunitParamsWithHTTPClient(client *http.Client) *DeleteWorkforcemanagementManagementunitParams {
	return &DeleteWorkforcemanagementManagementunitParams{
		HTTPClient: client,
	}
}

/*
DeleteWorkforcemanagementManagementunitParams contains all the parameters to send to the API endpoint

	for the delete workforcemanagement managementunit operation.

	Typically these are written to a http.Request.
*/
type DeleteWorkforcemanagementManagementunitParams struct {

	/* ManagementUnitID.

	   The ID of the management unit, or 'mine' for the management unit of the logged-in user.
	*/
	ManagementUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete workforcemanagement managementunit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkforcemanagementManagementunitParams) WithDefaults() *DeleteWorkforcemanagementManagementunitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete workforcemanagement managementunit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkforcemanagementManagementunitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete workforcemanagement managementunit params
func (o *DeleteWorkforcemanagementManagementunitParams) WithTimeout(timeout time.Duration) *DeleteWorkforcemanagementManagementunitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete workforcemanagement managementunit params
func (o *DeleteWorkforcemanagementManagementunitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete workforcemanagement managementunit params
func (o *DeleteWorkforcemanagementManagementunitParams) WithContext(ctx context.Context) *DeleteWorkforcemanagementManagementunitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete workforcemanagement managementunit params
func (o *DeleteWorkforcemanagementManagementunitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete workforcemanagement managementunit params
func (o *DeleteWorkforcemanagementManagementunitParams) WithHTTPClient(client *http.Client) *DeleteWorkforcemanagementManagementunitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete workforcemanagement managementunit params
func (o *DeleteWorkforcemanagementManagementunitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManagementUnitID adds the managementUnitID to the delete workforcemanagement managementunit params
func (o *DeleteWorkforcemanagementManagementunitParams) WithManagementUnitID(managementUnitID string) *DeleteWorkforcemanagementManagementunitParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the delete workforcemanagement managementunit params
func (o *DeleteWorkforcemanagementManagementunitParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkforcemanagementManagementunitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param managementUnitId
	if err := r.SetPathParam("managementUnitId", o.ManagementUnitID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
