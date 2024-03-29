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

// NewPatchWorkforcemanagementManagementunitWeekShifttradeParams creates a new PatchWorkforcemanagementManagementunitWeekShifttradeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchWorkforcemanagementManagementunitWeekShifttradeParams() *PatchWorkforcemanagementManagementunitWeekShifttradeParams {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeParamsWithTimeout creates a new PatchWorkforcemanagementManagementunitWeekShifttradeParams object
// with the ability to set a timeout on a request.
func NewPatchWorkforcemanagementManagementunitWeekShifttradeParamsWithTimeout(timeout time.Duration) *PatchWorkforcemanagementManagementunitWeekShifttradeParams {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeParams{
		timeout: timeout,
	}
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeParamsWithContext creates a new PatchWorkforcemanagementManagementunitWeekShifttradeParams object
// with the ability to set a context for a request.
func NewPatchWorkforcemanagementManagementunitWeekShifttradeParamsWithContext(ctx context.Context) *PatchWorkforcemanagementManagementunitWeekShifttradeParams {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeParams{
		Context: ctx,
	}
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeParamsWithHTTPClient creates a new PatchWorkforcemanagementManagementunitWeekShifttradeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchWorkforcemanagementManagementunitWeekShifttradeParamsWithHTTPClient(client *http.Client) *PatchWorkforcemanagementManagementunitWeekShifttradeParams {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeParams{
		HTTPClient: client,
	}
}

/*
PatchWorkforcemanagementManagementunitWeekShifttradeParams contains all the parameters to send to the API endpoint

	for the patch workforcemanagement managementunit week shifttrade operation.

	Typically these are written to a http.Request.
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeParams struct {

	/* Body.

	   body
	*/
	Body *models.PatchShiftTradeRequest

	/* ManagementUnitID.

	   The ID of the management unit, or 'mine' for the management unit of the logged-in user.
	*/
	ManagementUnitID string

	/* TradeID.

	   The ID of the shift trade to update
	*/
	TradeID string

	/* WeekDateID.

	   The start date of the week schedule in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd

	   Format: date
	*/
	WeekDateID strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch workforcemanagement managementunit week shifttrade params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) WithDefaults() *PatchWorkforcemanagementManagementunitWeekShifttradeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch workforcemanagement managementunit week shifttrade params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) WithTimeout(timeout time.Duration) *PatchWorkforcemanagementManagementunitWeekShifttradeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) WithContext(ctx context.Context) *PatchWorkforcemanagementManagementunitWeekShifttradeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) WithHTTPClient(client *http.Client) *PatchWorkforcemanagementManagementunitWeekShifttradeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) WithBody(body *models.PatchShiftTradeRequest) *PatchWorkforcemanagementManagementunitWeekShifttradeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) SetBody(body *models.PatchShiftTradeRequest) {
	o.Body = body
}

// WithManagementUnitID adds the managementUnitID to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) WithManagementUnitID(managementUnitID string) *PatchWorkforcemanagementManagementunitWeekShifttradeParams {
	o.SetManagementUnitID(managementUnitID)
	return o
}

// SetManagementUnitID adds the managementUnitId to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) SetManagementUnitID(managementUnitID string) {
	o.ManagementUnitID = managementUnitID
}

// WithTradeID adds the tradeID to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) WithTradeID(tradeID string) *PatchWorkforcemanagementManagementunitWeekShifttradeParams {
	o.SetTradeID(tradeID)
	return o
}

// SetTradeID adds the tradeId to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) SetTradeID(tradeID string) {
	o.TradeID = tradeID
}

// WithWeekDateID adds the weekDateID to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) WithWeekDateID(weekDateID strfmt.Date) *PatchWorkforcemanagementManagementunitWeekShifttradeParams {
	o.SetWeekDateID(weekDateID)
	return o
}

// SetWeekDateID adds the weekDateId to the patch workforcemanagement managementunit week shifttrade params
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) SetWeekDateID(weekDateID strfmt.Date) {
	o.WeekDateID = weekDateID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tradeId
	if err := r.SetPathParam("tradeId", o.TradeID); err != nil {
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
