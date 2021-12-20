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

// NewPostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams creates a new PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams object
// with the default values initialized.
func NewPostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams() *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams {
	var ()
	return &PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParamsWithTimeout creates a new PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams {
	var ()
	return &PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams{

		timeout: timeout,
	}
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParamsWithContext creates a new PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParamsWithContext(ctx context.Context) *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams {
	var ()
	return &PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams{

		Context: ctx,
	}
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParamsWithHTTPClient creates a new PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams {
	var ()
	return &PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams{
		HTTPClient: client,
	}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams contains all the parameters to send to the API endpoint
for the post workforcemanagement managementunit timeofflimits values query operation typically these are written to a http.Request
*/
type PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams struct {

	/*Body
	  body

	*/
	Body *models.QueryTimeOffLimitValuesRequest
	/*ManagementUnitID
	  The management unit ID of the management unit.

	*/
	ManagementUnitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post workforcemanagement managementunit timeofflimits values query params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement managementunit timeofflimits values query params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement managementunit timeofflimits values query params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams) WithContext(ctx context.Context) *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement managementunit timeofflimits values query params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement managementunit timeofflimits values query params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement managementunit timeofflimits values query params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement managementunit timeofflimits values query params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams) WithBody(body *models.QueryTimeOffLimitValuesRequest) *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement managementunit timeofflimits values query params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams) SetBody(body *models.QueryTimeOffLimitValuesRequest) {
	o.Body = body
}

// WithManagementUnitID adds the managementUnitID to the post workforcemanagement managementunit timeofflimits values query params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams) WithManagementUnitID(managementUnitID string) *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the post workforcemanagement managementunit timeofflimits values query params
func (o *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementManagementunitTimeofflimitsValuesQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
