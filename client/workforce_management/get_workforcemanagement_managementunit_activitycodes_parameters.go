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

// NewGetWorkforcemanagementManagementunitActivitycodesParams creates a new GetWorkforcemanagementManagementunitActivitycodesParams object
// with the default values initialized.
func NewGetWorkforcemanagementManagementunitActivitycodesParams() *GetWorkforcemanagementManagementunitActivitycodesParams {
	var ()
	return &GetWorkforcemanagementManagementunitActivitycodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementManagementunitActivitycodesParamsWithTimeout creates a new GetWorkforcemanagementManagementunitActivitycodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementManagementunitActivitycodesParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitActivitycodesParams {
	var ()
	return &GetWorkforcemanagementManagementunitActivitycodesParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementManagementunitActivitycodesParamsWithContext creates a new GetWorkforcemanagementManagementunitActivitycodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementManagementunitActivitycodesParamsWithContext(ctx context.Context) *GetWorkforcemanagementManagementunitActivitycodesParams {
	var ()
	return &GetWorkforcemanagementManagementunitActivitycodesParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementManagementunitActivitycodesParamsWithHTTPClient creates a new GetWorkforcemanagementManagementunitActivitycodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementManagementunitActivitycodesParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitActivitycodesParams {
	var ()
	return &GetWorkforcemanagementManagementunitActivitycodesParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementManagementunitActivitycodesParams contains all the parameters to send to the API endpoint
for the get workforcemanagement managementunit activitycodes operation typically these are written to a http.Request
*/
type GetWorkforcemanagementManagementunitActivitycodesParams struct {

	/*MuID
	  The ID of the management unit, or 'mine' for the management unit of the logged-in user.

	*/
	MuID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement managementunit activitycodes params
func (o *GetWorkforcemanagementManagementunitActivitycodesParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitActivitycodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement managementunit activitycodes params
func (o *GetWorkforcemanagementManagementunitActivitycodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement managementunit activitycodes params
func (o *GetWorkforcemanagementManagementunitActivitycodesParams) WithContext(ctx context.Context) *GetWorkforcemanagementManagementunitActivitycodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement managementunit activitycodes params
func (o *GetWorkforcemanagementManagementunitActivitycodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement managementunit activitycodes params
func (o *GetWorkforcemanagementManagementunitActivitycodesParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitActivitycodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement managementunit activitycodes params
func (o *GetWorkforcemanagementManagementunitActivitycodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMuID adds the muID to the get workforcemanagement managementunit activitycodes params
func (o *GetWorkforcemanagementManagementunitActivitycodesParams) WithMuID(muID string) *GetWorkforcemanagementManagementunitActivitycodesParams {
	o.SetMuID(muID)
	return o
}

// SetMuID adds the muId to the get workforcemanagement managementunit activitycodes params
func (o *GetWorkforcemanagementManagementunitActivitycodesParams) SetMuID(muID string) {
	o.MuID = muID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementManagementunitActivitycodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param muId
	if err := r.SetPathParam("muId", o.MuID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}