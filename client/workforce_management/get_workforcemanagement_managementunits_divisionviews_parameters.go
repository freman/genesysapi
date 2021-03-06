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

// NewGetWorkforcemanagementManagementunitsDivisionviewsParams creates a new GetWorkforcemanagementManagementunitsDivisionviewsParams object
// with the default values initialized.
func NewGetWorkforcemanagementManagementunitsDivisionviewsParams() *GetWorkforcemanagementManagementunitsDivisionviewsParams {
	var ()
	return &GetWorkforcemanagementManagementunitsDivisionviewsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementManagementunitsDivisionviewsParamsWithTimeout creates a new GetWorkforcemanagementManagementunitsDivisionviewsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementManagementunitsDivisionviewsParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitsDivisionviewsParams {
	var ()
	return &GetWorkforcemanagementManagementunitsDivisionviewsParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementManagementunitsDivisionviewsParamsWithContext creates a new GetWorkforcemanagementManagementunitsDivisionviewsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementManagementunitsDivisionviewsParamsWithContext(ctx context.Context) *GetWorkforcemanagementManagementunitsDivisionviewsParams {
	var ()
	return &GetWorkforcemanagementManagementunitsDivisionviewsParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementManagementunitsDivisionviewsParamsWithHTTPClient creates a new GetWorkforcemanagementManagementunitsDivisionviewsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementManagementunitsDivisionviewsParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitsDivisionviewsParams {
	var ()
	return &GetWorkforcemanagementManagementunitsDivisionviewsParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementManagementunitsDivisionviewsParams contains all the parameters to send to the API endpoint
for the get workforcemanagement managementunits divisionviews operation typically these are written to a http.Request
*/
type GetWorkforcemanagementManagementunitsDivisionviewsParams struct {

	/*DivisionID
	  The divisionIds to filter by. If omitted, will return all divisions

	*/
	DivisionID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement managementunits divisionviews params
func (o *GetWorkforcemanagementManagementunitsDivisionviewsParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementManagementunitsDivisionviewsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement managementunits divisionviews params
func (o *GetWorkforcemanagementManagementunitsDivisionviewsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement managementunits divisionviews params
func (o *GetWorkforcemanagementManagementunitsDivisionviewsParams) WithContext(ctx context.Context) *GetWorkforcemanagementManagementunitsDivisionviewsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement managementunits divisionviews params
func (o *GetWorkforcemanagementManagementunitsDivisionviewsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement managementunits divisionviews params
func (o *GetWorkforcemanagementManagementunitsDivisionviewsParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementManagementunitsDivisionviewsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement managementunits divisionviews params
func (o *GetWorkforcemanagementManagementunitsDivisionviewsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDivisionID adds the divisionID to the get workforcemanagement managementunits divisionviews params
func (o *GetWorkforcemanagementManagementunitsDivisionviewsParams) WithDivisionID(divisionID []string) *GetWorkforcemanagementManagementunitsDivisionviewsParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the get workforcemanagement managementunits divisionviews params
func (o *GetWorkforcemanagementManagementunitsDivisionviewsParams) SetDivisionID(divisionID []string) {
	o.DivisionID = divisionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementManagementunitsDivisionviewsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDivisionID := o.DivisionID

	joinedDivisionID := swag.JoinByFormat(valuesDivisionID, "multi")
	// query array param divisionId
	if err := r.SetQueryParam("divisionId", joinedDivisionID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
