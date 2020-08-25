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

// NewGetWorkforcemanagementManagementunitShifttradesMatchedParams creates a new GetWorkforcemanagementManagementunitShifttradesMatchedParams object
// with the default values initialized.
func NewGetWorkforcemanagementManagementunitShifttradesMatchedParams() *GetWorkforcemanagementManagementunitShifttradesMatchedParams {
	var ()
	return &GetWorkforcemanagementManagementunitShifttradesMatchedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedParamsWithTimeout creates a new GetWorkforcemanagementManagementunitShifttradesMatchedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementManagementunitShifttradesMatchedParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitShifttradesMatchedParams {
	var ()
	return &GetWorkforcemanagementManagementunitShifttradesMatchedParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedParamsWithContext creates a new GetWorkforcemanagementManagementunitShifttradesMatchedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementManagementunitShifttradesMatchedParamsWithContext(ctx context.Context) *GetWorkforcemanagementManagementunitShifttradesMatchedParams {
	var ()
	return &GetWorkforcemanagementManagementunitShifttradesMatchedParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedParamsWithHTTPClient creates a new GetWorkforcemanagementManagementunitShifttradesMatchedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementManagementunitShifttradesMatchedParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitShifttradesMatchedParams {
	var ()
	return &GetWorkforcemanagementManagementunitShifttradesMatchedParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementManagementunitShifttradesMatchedParams contains all the parameters to send to the API endpoint
for the get workforcemanagement managementunit shifttrades matched operation typically these are written to a http.Request
*/
type GetWorkforcemanagementManagementunitShifttradesMatchedParams struct {

	/*MuID
	  The management unit ID of the management unit, or 'mine' for the management unit of the logged-in user.

	*/
	MuID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement managementunit shifttrades matched params
func (o *GetWorkforcemanagementManagementunitShifttradesMatchedParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitShifttradesMatchedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement managementunit shifttrades matched params
func (o *GetWorkforcemanagementManagementunitShifttradesMatchedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement managementunit shifttrades matched params
func (o *GetWorkforcemanagementManagementunitShifttradesMatchedParams) WithContext(ctx context.Context) *GetWorkforcemanagementManagementunitShifttradesMatchedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement managementunit shifttrades matched params
func (o *GetWorkforcemanagementManagementunitShifttradesMatchedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement managementunit shifttrades matched params
func (o *GetWorkforcemanagementManagementunitShifttradesMatchedParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitShifttradesMatchedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement managementunit shifttrades matched params
func (o *GetWorkforcemanagementManagementunitShifttradesMatchedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMuID adds the muID to the get workforcemanagement managementunit shifttrades matched params
func (o *GetWorkforcemanagementManagementunitShifttradesMatchedParams) WithMuID(muID string) *GetWorkforcemanagementManagementunitShifttradesMatchedParams {
	o.SetMuID(muID)
	return o
}

// SetMuID adds the muId to the get workforcemanagement managementunit shifttrades matched params
func (o *GetWorkforcemanagementManagementunitShifttradesMatchedParams) SetMuID(muID string) {
	o.MuID = muID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementManagementunitShifttradesMatchedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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