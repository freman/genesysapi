// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewGetFlowHistoryHistoryIDParams creates a new GetFlowHistoryHistoryIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFlowHistoryHistoryIDParams() *GetFlowHistoryHistoryIDParams {
	return &GetFlowHistoryHistoryIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowHistoryHistoryIDParamsWithTimeout creates a new GetFlowHistoryHistoryIDParams object
// with the ability to set a timeout on a request.
func NewGetFlowHistoryHistoryIDParamsWithTimeout(timeout time.Duration) *GetFlowHistoryHistoryIDParams {
	return &GetFlowHistoryHistoryIDParams{
		timeout: timeout,
	}
}

// NewGetFlowHistoryHistoryIDParamsWithContext creates a new GetFlowHistoryHistoryIDParams object
// with the ability to set a context for a request.
func NewGetFlowHistoryHistoryIDParamsWithContext(ctx context.Context) *GetFlowHistoryHistoryIDParams {
	return &GetFlowHistoryHistoryIDParams{
		Context: ctx,
	}
}

// NewGetFlowHistoryHistoryIDParamsWithHTTPClient creates a new GetFlowHistoryHistoryIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFlowHistoryHistoryIDParamsWithHTTPClient(client *http.Client) *GetFlowHistoryHistoryIDParams {
	return &GetFlowHistoryHistoryIDParams{
		HTTPClient: client,
	}
}

/*
GetFlowHistoryHistoryIDParams contains all the parameters to send to the API endpoint

	for the get flow history history Id operation.

	Typically these are written to a http.Request.
*/
type GetFlowHistoryHistoryIDParams struct {

	/* Action.

	   Flow actions to include (omit to include all)
	*/
	Action []string

	/* FlowID.

	   Flow ID
	*/
	FlowID string

	/* HistoryID.

	   History request ID
	*/
	HistoryID string

	/* PageNumber.

	   Page number

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   Page size

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	/* SortBy.

	   Sort by

	   Default: "timestamp"
	*/
	SortBy *string

	/* SortOrder.

	   Sort order

	   Default: "desc"
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get flow history history Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowHistoryHistoryIDParams) WithDefaults() *GetFlowHistoryHistoryIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get flow history history Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowHistoryHistoryIDParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		sortByDefault = string("timestamp")

		sortOrderDefault = string("desc")
	)

	val := GetFlowHistoryHistoryIDParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) WithTimeout(timeout time.Duration) *GetFlowHistoryHistoryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) WithContext(ctx context.Context) *GetFlowHistoryHistoryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) WithHTTPClient(client *http.Client) *GetFlowHistoryHistoryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) WithAction(action []string) *GetFlowHistoryHistoryIDParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) SetAction(action []string) {
	o.Action = action
}

// WithFlowID adds the flowID to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) WithFlowID(flowID string) *GetFlowHistoryHistoryIDParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WithHistoryID adds the historyID to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) WithHistoryID(historyID string) *GetFlowHistoryHistoryIDParams {
	o.SetHistoryID(historyID)
	return o
}

// SetHistoryID adds the historyId to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) SetHistoryID(historyID string) {
	o.HistoryID = historyID
}

// WithPageNumber adds the pageNumber to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) WithPageNumber(pageNumber *int32) *GetFlowHistoryHistoryIDParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) WithPageSize(pageSize *int32) *GetFlowHistoryHistoryIDParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) WithSortBy(sortBy *string) *GetFlowHistoryHistoryIDParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) WithSortOrder(sortOrder *string) *GetFlowHistoryHistoryIDParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get flow history history Id params
func (o *GetFlowHistoryHistoryIDParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowHistoryHistoryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Action != nil {

		// binding items for action
		joinedAction := o.bindParamAction(reg)

		// query array param action
		if err := r.SetQueryParam("action", joinedAction...); err != nil {
			return err
		}
	}

	// path param flowId
	if err := r.SetPathParam("flowId", o.FlowID); err != nil {
		return err
	}

	// path param historyId
	if err := r.SetPathParam("historyId", o.HistoryID); err != nil {
		return err
	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string

		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {

			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}
	}

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string

		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {

			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetFlowHistoryHistoryID binds the parameter action
func (o *GetFlowHistoryHistoryIDParams) bindParamAction(formats strfmt.Registry) []string {
	actionIR := o.Action

	var actionIC []string
	for _, actionIIR := range actionIR { // explode []string

		actionIIV := actionIIR // string as string
		actionIC = append(actionIC, actionIIV)
	}

	// items.CollectionFormat: "multi"
	actionIS := swag.JoinByFormat(actionIC, "multi")

	return actionIS
}
