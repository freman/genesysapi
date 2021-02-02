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
	"github.com/go-openapi/swag"
)

// NewGetWorkforcemanagementManagementunitWorkplanrotationsParams creates a new GetWorkforcemanagementManagementunitWorkplanrotationsParams object
// with the default values initialized.
func NewGetWorkforcemanagementManagementunitWorkplanrotationsParams() *GetWorkforcemanagementManagementunitWorkplanrotationsParams {
	var ()
	return &GetWorkforcemanagementManagementunitWorkplanrotationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsParamsWithTimeout creates a new GetWorkforcemanagementManagementunitWorkplanrotationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementManagementunitWorkplanrotationsParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitWorkplanrotationsParams {
	var ()
	return &GetWorkforcemanagementManagementunitWorkplanrotationsParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsParamsWithContext creates a new GetWorkforcemanagementManagementunitWorkplanrotationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementManagementunitWorkplanrotationsParamsWithContext(ctx context.Context) *GetWorkforcemanagementManagementunitWorkplanrotationsParams {
	var ()
	return &GetWorkforcemanagementManagementunitWorkplanrotationsParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsParamsWithHTTPClient creates a new GetWorkforcemanagementManagementunitWorkplanrotationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementManagementunitWorkplanrotationsParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitWorkplanrotationsParams {
	var ()
	return &GetWorkforcemanagementManagementunitWorkplanrotationsParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsParams contains all the parameters to send to the API endpoint
for the get workforcemanagement managementunit workplanrotations operation typically these are written to a http.Request
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsParams struct {

	/*Expand*/
	Expand []string
	/*ManagementUnitID
	  The ID of the management unit, or 'mine' for the management unit of the logged-in user.

	*/
	ManagementUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement managementunit workplanrotations params
func (o *GetWorkforcemanagementManagementunitWorkplanrotationsParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitWorkplanrotationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement managementunit workplanrotations params
func (o *GetWorkforcemanagementManagementunitWorkplanrotationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement managementunit workplanrotations params
func (o *GetWorkforcemanagementManagementunitWorkplanrotationsParams) WithContext(ctx context.Context) *GetWorkforcemanagementManagementunitWorkplanrotationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement managementunit workplanrotations params
func (o *GetWorkforcemanagementManagementunitWorkplanrotationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement managementunit workplanrotations params
func (o *GetWorkforcemanagementManagementunitWorkplanrotationsParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitWorkplanrotationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement managementunit workplanrotations params
func (o *GetWorkforcemanagementManagementunitWorkplanrotationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get workforcemanagement managementunit workplanrotations params
func (o *GetWorkforcemanagementManagementunitWorkplanrotationsParams) WithExpand(expand []string) *GetWorkforcemanagementManagementunitWorkplanrotationsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get workforcemanagement managementunit workplanrotations params
func (o *GetWorkforcemanagementManagementunitWorkplanrotationsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithManagementUnitID adds the managementUnitID to the get workforcemanagement managementunit workplanrotations params
func (o *GetWorkforcemanagementManagementunitWorkplanrotationsParams) WithManagementUnitID(managementUnitID string) *GetWorkforcemanagementManagementunitWorkplanrotationsParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the get workforcemanagement managementunit workplanrotations params
func (o *GetWorkforcemanagementManagementunitWorkplanrotationsParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementManagementunitWorkplanrotationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
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