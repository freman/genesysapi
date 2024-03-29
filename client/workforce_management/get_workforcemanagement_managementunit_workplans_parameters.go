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

// NewGetWorkforcemanagementManagementunitWorkplansParams creates a new GetWorkforcemanagementManagementunitWorkplansParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkforcemanagementManagementunitWorkplansParams() *GetWorkforcemanagementManagementunitWorkplansParams {
	return &GetWorkforcemanagementManagementunitWorkplansParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementManagementunitWorkplansParamsWithTimeout creates a new GetWorkforcemanagementManagementunitWorkplansParams object
// with the ability to set a timeout on a request.
func NewGetWorkforcemanagementManagementunitWorkplansParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitWorkplansParams {
	return &GetWorkforcemanagementManagementunitWorkplansParams{
		timeout: timeout,
	}
}

// NewGetWorkforcemanagementManagementunitWorkplansParamsWithContext creates a new GetWorkforcemanagementManagementunitWorkplansParams object
// with the ability to set a context for a request.
func NewGetWorkforcemanagementManagementunitWorkplansParamsWithContext(ctx context.Context) *GetWorkforcemanagementManagementunitWorkplansParams {
	return &GetWorkforcemanagementManagementunitWorkplansParams{
		Context: ctx,
	}
}

// NewGetWorkforcemanagementManagementunitWorkplansParamsWithHTTPClient creates a new GetWorkforcemanagementManagementunitWorkplansParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkforcemanagementManagementunitWorkplansParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitWorkplansParams {
	return &GetWorkforcemanagementManagementunitWorkplansParams{
		HTTPClient: client,
	}
}

/*
GetWorkforcemanagementManagementunitWorkplansParams contains all the parameters to send to the API endpoint

	for the get workforcemanagement managementunit workplans operation.

	Typically these are written to a http.Request.
*/
type GetWorkforcemanagementManagementunitWorkplansParams struct {

	// Expand.
	Expand []string

	/* ManagementUnitID.

	   The ID of the management unit, or 'mine' for the management unit of the logged-in user.
	*/
	ManagementUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workforcemanagement managementunit workplans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementManagementunitWorkplansParams) WithDefaults() *GetWorkforcemanagementManagementunitWorkplansParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workforcemanagement managementunit workplans params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkforcemanagementManagementunitWorkplansParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workforcemanagement managementunit workplans params
func (o *GetWorkforcemanagementManagementunitWorkplansParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitWorkplansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement managementunit workplans params
func (o *GetWorkforcemanagementManagementunitWorkplansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement managementunit workplans params
func (o *GetWorkforcemanagementManagementunitWorkplansParams) WithContext(ctx context.Context) *GetWorkforcemanagementManagementunitWorkplansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement managementunit workplans params
func (o *GetWorkforcemanagementManagementunitWorkplansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement managementunit workplans params
func (o *GetWorkforcemanagementManagementunitWorkplansParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitWorkplansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement managementunit workplans params
func (o *GetWorkforcemanagementManagementunitWorkplansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get workforcemanagement managementunit workplans params
func (o *GetWorkforcemanagementManagementunitWorkplansParams) WithExpand(expand []string) *GetWorkforcemanagementManagementunitWorkplansParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get workforcemanagement managementunit workplans params
func (o *GetWorkforcemanagementManagementunitWorkplansParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithManagementUnitID adds the managementUnitID to the get workforcemanagement managementunit workplans params
func (o *GetWorkforcemanagementManagementunitWorkplansParams) WithManagementUnitID(managementUnitID string) *GetWorkforcemanagementManagementunitWorkplansParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the get workforcemanagement managementunit workplans params
func (o *GetWorkforcemanagementManagementunitWorkplansParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementManagementunitWorkplansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
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

// bindParamGetWorkforcemanagementManagementunitWorkplans binds the parameter expand
func (o *GetWorkforcemanagementManagementunitWorkplansParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}
