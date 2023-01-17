// Code generated by go-swagger; DO NOT EDIT.

package scripts

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

// NewGetScriptsDivisionviewsParams creates a new GetScriptsDivisionviewsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScriptsDivisionviewsParams() *GetScriptsDivisionviewsParams {
	return &GetScriptsDivisionviewsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScriptsDivisionviewsParamsWithTimeout creates a new GetScriptsDivisionviewsParams object
// with the ability to set a timeout on a request.
func NewGetScriptsDivisionviewsParamsWithTimeout(timeout time.Duration) *GetScriptsDivisionviewsParams {
	return &GetScriptsDivisionviewsParams{
		timeout: timeout,
	}
}

// NewGetScriptsDivisionviewsParamsWithContext creates a new GetScriptsDivisionviewsParams object
// with the ability to set a context for a request.
func NewGetScriptsDivisionviewsParamsWithContext(ctx context.Context) *GetScriptsDivisionviewsParams {
	return &GetScriptsDivisionviewsParams{
		Context: ctx,
	}
}

// NewGetScriptsDivisionviewsParamsWithHTTPClient creates a new GetScriptsDivisionviewsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScriptsDivisionviewsParamsWithHTTPClient(client *http.Client) *GetScriptsDivisionviewsParams {
	return &GetScriptsDivisionviewsParams{
		HTTPClient: client,
	}
}

/*
GetScriptsDivisionviewsParams contains all the parameters to send to the API endpoint

	for the get scripts divisionviews operation.

	Typically these are written to a http.Request.
*/
type GetScriptsDivisionviewsParams struct {

	/* DivisionIds.

	   Filters scripts to requested divisionIds
	*/
	DivisionIds *string

	/* Expand.

	   Expand
	*/
	Expand *string

	/* Feature.

	   Feature filter
	*/
	Feature *string

	/* FlowID.

	   Secure flow id filter
	*/
	FlowID *string

	/* Name.

	   Name filter
	*/
	Name *string

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

	/* ScriptDataVersion.

	   Advanced usage - controls the data version of the script
	*/
	ScriptDataVersion *string

	/* SortBy.

	   SortBy
	*/
	SortBy *string

	/* SortOrder.

	   SortOrder
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get scripts divisionviews params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScriptsDivisionviewsParams) WithDefaults() *GetScriptsDivisionviewsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scripts divisionviews params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScriptsDivisionviewsParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetScriptsDivisionviewsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithTimeout(timeout time.Duration) *GetScriptsDivisionviewsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithContext(ctx context.Context) *GetScriptsDivisionviewsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithHTTPClient(client *http.Client) *GetScriptsDivisionviewsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDivisionIds adds the divisionIds to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithDivisionIds(divisionIds *string) *GetScriptsDivisionviewsParams {
	o.SetDivisionIds(divisionIds)
	return o
}

// SetDivisionIds adds the divisionIds to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetDivisionIds(divisionIds *string) {
	o.DivisionIds = divisionIds
}

// WithExpand adds the expand to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithExpand(expand *string) *GetScriptsDivisionviewsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithFeature adds the feature to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithFeature(feature *string) *GetScriptsDivisionviewsParams {
	o.SetFeature(feature)
	return o
}

// SetFeature adds the feature to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetFeature(feature *string) {
	o.Feature = feature
}

// WithFlowID adds the flowID to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithFlowID(flowID *string) *GetScriptsDivisionviewsParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetFlowID(flowID *string) {
	o.FlowID = flowID
}

// WithName adds the name to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithName(name *string) *GetScriptsDivisionviewsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithPageNumber(pageNumber *int32) *GetScriptsDivisionviewsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithPageSize(pageSize *int32) *GetScriptsDivisionviewsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithScriptDataVersion adds the scriptDataVersion to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithScriptDataVersion(scriptDataVersion *string) *GetScriptsDivisionviewsParams {
	o.SetScriptDataVersion(scriptDataVersion)
	return o
}

// SetScriptDataVersion adds the scriptDataVersion to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetScriptDataVersion(scriptDataVersion *string) {
	o.ScriptDataVersion = scriptDataVersion
}

// WithSortBy adds the sortBy to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithSortBy(sortBy *string) *GetScriptsDivisionviewsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) WithSortOrder(sortOrder *string) *GetScriptsDivisionviewsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get scripts divisionviews params
func (o *GetScriptsDivisionviewsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetScriptsDivisionviewsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DivisionIds != nil {

		// query param divisionIds
		var qrDivisionIds string

		if o.DivisionIds != nil {
			qrDivisionIds = *o.DivisionIds
		}
		qDivisionIds := qrDivisionIds
		if qDivisionIds != "" {

			if err := r.SetQueryParam("divisionIds", qDivisionIds); err != nil {
				return err
			}
		}
	}

	if o.Expand != nil {

		// query param expand
		var qrExpand string

		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {

			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}
	}

	if o.Feature != nil {

		// query param feature
		var qrFeature string

		if o.Feature != nil {
			qrFeature = *o.Feature
		}
		qFeature := qrFeature
		if qFeature != "" {

			if err := r.SetQueryParam("feature", qFeature); err != nil {
				return err
			}
		}
	}

	if o.FlowID != nil {

		// query param flowId
		var qrFlowID string

		if o.FlowID != nil {
			qrFlowID = *o.FlowID
		}
		qFlowID := qrFlowID
		if qFlowID != "" {

			if err := r.SetQueryParam("flowId", qFlowID); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
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

	if o.ScriptDataVersion != nil {

		// query param scriptDataVersion
		var qrScriptDataVersion string

		if o.ScriptDataVersion != nil {
			qrScriptDataVersion = *o.ScriptDataVersion
		}
		qScriptDataVersion := qrScriptDataVersion
		if qScriptDataVersion != "" {

			if err := r.SetQueryParam("scriptDataVersion", qScriptDataVersion); err != nil {
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
