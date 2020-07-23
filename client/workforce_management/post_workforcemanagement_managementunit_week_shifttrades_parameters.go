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

// NewPostWorkforcemanagementManagementunitWeekShifttradesParams creates a new PostWorkforcemanagementManagementunitWeekShifttradesParams object
// with the default values initialized.
func NewPostWorkforcemanagementManagementunitWeekShifttradesParams() *PostWorkforcemanagementManagementunitWeekShifttradesParams {
	var ()
	return &PostWorkforcemanagementManagementunitWeekShifttradesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementManagementunitWeekShifttradesParamsWithTimeout creates a new PostWorkforcemanagementManagementunitWeekShifttradesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWorkforcemanagementManagementunitWeekShifttradesParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementManagementunitWeekShifttradesParams {
	var ()
	return &PostWorkforcemanagementManagementunitWeekShifttradesParams{

		timeout: timeout,
	}
}

// NewPostWorkforcemanagementManagementunitWeekShifttradesParamsWithContext creates a new PostWorkforcemanagementManagementunitWeekShifttradesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWorkforcemanagementManagementunitWeekShifttradesParamsWithContext(ctx context.Context) *PostWorkforcemanagementManagementunitWeekShifttradesParams {
	var ()
	return &PostWorkforcemanagementManagementunitWeekShifttradesParams{

		Context: ctx,
	}
}

// NewPostWorkforcemanagementManagementunitWeekShifttradesParamsWithHTTPClient creates a new PostWorkforcemanagementManagementunitWeekShifttradesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWorkforcemanagementManagementunitWeekShifttradesParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementManagementunitWeekShifttradesParams {
	var ()
	return &PostWorkforcemanagementManagementunitWeekShifttradesParams{
		HTTPClient: client,
	}
}

/*PostWorkforcemanagementManagementunitWeekShifttradesParams contains all the parameters to send to the API endpoint
for the post workforcemanagement managementunit week shifttrades operation typically these are written to a http.Request
*/
type PostWorkforcemanagementManagementunitWeekShifttradesParams struct {

	/*Body
	  body

	*/
	Body *models.AddShiftTradeRequest
	/*ManagementUnitID
	  The management unit ID of the management unit, or 'mine' for the management unit of the logged-in user.

	*/
	ManagementUnitID string
	/*WeekDateID
	  The start date of the week schedule in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	*/
	WeekDateID strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post workforcemanagement managementunit week shifttrades params
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementManagementunitWeekShifttradesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement managementunit week shifttrades params
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement managementunit week shifttrades params
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) WithContext(ctx context.Context) *PostWorkforcemanagementManagementunitWeekShifttradesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement managementunit week shifttrades params
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement managementunit week shifttrades params
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementManagementunitWeekShifttradesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement managementunit week shifttrades params
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement managementunit week shifttrades params
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) WithBody(body *models.AddShiftTradeRequest) *PostWorkforcemanagementManagementunitWeekShifttradesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement managementunit week shifttrades params
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) SetBody(body *models.AddShiftTradeRequest) {
	o.Body = body
}

// WithManagementUnitID adds the managementUnitID to the post workforcemanagement managementunit week shifttrades params
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) WithManagementUnitID(managementUnitID string) *PostWorkforcemanagementManagementunitWeekShifttradesParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the post workforcemanagement managementunit week shifttrades params
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WithWeekDateID adds the weekDateID to the post workforcemanagement managementunit week shifttrades params
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) WithWeekDateID(weekDateID strfmt.Date) *PostWorkforcemanagementManagementunitWeekShifttradesParams {
	o.SetWeekDateID(weekDateID)
	return o
}

// SetWeekDateID adds the weekDateId to the post workforcemanagement managementunit week shifttrades params
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) SetWeekDateID(weekDateID strfmt.Date) {
	o.WeekDateID = weekDateID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementManagementunitWeekShifttradesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param weekDateId
	if err := r.SetPathParam("weekDateId", o.WeekDateID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
