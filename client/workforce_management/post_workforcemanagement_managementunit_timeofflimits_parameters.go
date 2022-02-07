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

// NewPostWorkforcemanagementManagementunitTimeofflimitsParams creates a new PostWorkforcemanagementManagementunitTimeofflimitsParams object
// with the default values initialized.
func NewPostWorkforcemanagementManagementunitTimeofflimitsParams() *PostWorkforcemanagementManagementunitTimeofflimitsParams {
	var ()
	return &PostWorkforcemanagementManagementunitTimeofflimitsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsParamsWithTimeout creates a new PostWorkforcemanagementManagementunitTimeofflimitsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWorkforcemanagementManagementunitTimeofflimitsParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementManagementunitTimeofflimitsParams {
	var ()
	return &PostWorkforcemanagementManagementunitTimeofflimitsParams{

		timeout: timeout,
	}
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsParamsWithContext creates a new PostWorkforcemanagementManagementunitTimeofflimitsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWorkforcemanagementManagementunitTimeofflimitsParamsWithContext(ctx context.Context) *PostWorkforcemanagementManagementunitTimeofflimitsParams {
	var ()
	return &PostWorkforcemanagementManagementunitTimeofflimitsParams{

		Context: ctx,
	}
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsParamsWithHTTPClient creates a new PostWorkforcemanagementManagementunitTimeofflimitsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWorkforcemanagementManagementunitTimeofflimitsParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementManagementunitTimeofflimitsParams {
	var ()
	return &PostWorkforcemanagementManagementunitTimeofflimitsParams{
		HTTPClient: client,
	}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsParams contains all the parameters to send to the API endpoint
for the post workforcemanagement managementunit timeofflimits operation typically these are written to a http.Request
*/
type PostWorkforcemanagementManagementunitTimeofflimitsParams struct {

	/*Body
	  body

	*/
	Body *models.CreateTimeOffLimitRequest
	/*ManagementUnitID
	  The management unit ID of the management unit.

	*/
	ManagementUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post workforcemanagement managementunit timeofflimits params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementManagementunitTimeofflimitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement managementunit timeofflimits params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement managementunit timeofflimits params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsParams) WithContext(ctx context.Context) *PostWorkforcemanagementManagementunitTimeofflimitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement managementunit timeofflimits params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement managementunit timeofflimits params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementManagementunitTimeofflimitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement managementunit timeofflimits params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement managementunit timeofflimits params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsParams) WithBody(body *models.CreateTimeOffLimitRequest) *PostWorkforcemanagementManagementunitTimeofflimitsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement managementunit timeofflimits params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsParams) SetBody(body *models.CreateTimeOffLimitRequest) {
	o.Body = body
}

// WithManagementUnitID adds the managementUnitID to the post workforcemanagement managementunit timeofflimits params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsParams) WithManagementUnitID(managementUnitID string) *PostWorkforcemanagementManagementunitTimeofflimitsParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the post workforcemanagement managementunit timeofflimits params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementManagementunitTimeofflimitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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