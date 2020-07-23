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

// NewDeleteWorkforcemanagementManagementunitActivitycodeParams creates a new DeleteWorkforcemanagementManagementunitActivitycodeParams object
// with the default values initialized.
func NewDeleteWorkforcemanagementManagementunitActivitycodeParams() *DeleteWorkforcemanagementManagementunitActivitycodeParams {
	var ()
	return &DeleteWorkforcemanagementManagementunitActivitycodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkforcemanagementManagementunitActivitycodeParamsWithTimeout creates a new DeleteWorkforcemanagementManagementunitActivitycodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteWorkforcemanagementManagementunitActivitycodeParamsWithTimeout(timeout time.Duration) *DeleteWorkforcemanagementManagementunitActivitycodeParams {
	var ()
	return &DeleteWorkforcemanagementManagementunitActivitycodeParams{

		timeout: timeout,
	}
}

// NewDeleteWorkforcemanagementManagementunitActivitycodeParamsWithContext creates a new DeleteWorkforcemanagementManagementunitActivitycodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteWorkforcemanagementManagementunitActivitycodeParamsWithContext(ctx context.Context) *DeleteWorkforcemanagementManagementunitActivitycodeParams {
	var ()
	return &DeleteWorkforcemanagementManagementunitActivitycodeParams{

		Context: ctx,
	}
}

// NewDeleteWorkforcemanagementManagementunitActivitycodeParamsWithHTTPClient creates a new DeleteWorkforcemanagementManagementunitActivitycodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteWorkforcemanagementManagementunitActivitycodeParamsWithHTTPClient(client *http.Client) *DeleteWorkforcemanagementManagementunitActivitycodeParams {
	var ()
	return &DeleteWorkforcemanagementManagementunitActivitycodeParams{
		HTTPClient: client,
	}
}

/*DeleteWorkforcemanagementManagementunitActivitycodeParams contains all the parameters to send to the API endpoint
for the delete workforcemanagement managementunit activitycode operation typically these are written to a http.Request
*/
type DeleteWorkforcemanagementManagementunitActivitycodeParams struct {

	/*AcID
	  The ID of the activity code to delete

	*/
	AcID string
	/*MuID
	  The ID of the management unit, or 'mine' for the management unit of the logged-in user.

	*/
	MuID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete workforcemanagement managementunit activitycode params
func (o *DeleteWorkforcemanagementManagementunitActivitycodeParams) WithTimeout(timeout time.Duration) *DeleteWorkforcemanagementManagementunitActivitycodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete workforcemanagement managementunit activitycode params
func (o *DeleteWorkforcemanagementManagementunitActivitycodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete workforcemanagement managementunit activitycode params
func (o *DeleteWorkforcemanagementManagementunitActivitycodeParams) WithContext(ctx context.Context) *DeleteWorkforcemanagementManagementunitActivitycodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete workforcemanagement managementunit activitycode params
func (o *DeleteWorkforcemanagementManagementunitActivitycodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete workforcemanagement managementunit activitycode params
func (o *DeleteWorkforcemanagementManagementunitActivitycodeParams) WithHTTPClient(client *http.Client) *DeleteWorkforcemanagementManagementunitActivitycodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete workforcemanagement managementunit activitycode params
func (o *DeleteWorkforcemanagementManagementunitActivitycodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcID adds the acID to the delete workforcemanagement managementunit activitycode params
func (o *DeleteWorkforcemanagementManagementunitActivitycodeParams) WithAcID(acID string) *DeleteWorkforcemanagementManagementunitActivitycodeParams {
	o.SetAcID(acID)
	return o
}

// SetAcID adds the acId to the delete workforcemanagement managementunit activitycode params
func (o *DeleteWorkforcemanagementManagementunitActivitycodeParams) SetAcID(acID string) {
	o.AcID = acID
}

// WithMuID adds the muID to the delete workforcemanagement managementunit activitycode params
func (o *DeleteWorkforcemanagementManagementunitActivitycodeParams) WithMuID(muID string) *DeleteWorkforcemanagementManagementunitActivitycodeParams {
	o.SetMuID(muID)
	return o
}

// SetMuID adds the muId to the delete workforcemanagement managementunit activitycode params
func (o *DeleteWorkforcemanagementManagementunitActivitycodeParams) SetMuID(muID string) {
	o.MuID = muID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkforcemanagementManagementunitActivitycodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param acId
	if err := r.SetPathParam("acId", o.AcID); err != nil {
		return err
	}

	// path param muId
	if err := r.SetPathParam("muId", o.MuID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
