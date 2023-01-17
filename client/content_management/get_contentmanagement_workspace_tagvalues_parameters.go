// Code generated by go-swagger; DO NOT EDIT.

package content_management

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

// NewGetContentmanagementWorkspaceTagvaluesParams creates a new GetContentmanagementWorkspaceTagvaluesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContentmanagementWorkspaceTagvaluesParams() *GetContentmanagementWorkspaceTagvaluesParams {
	return &GetContentmanagementWorkspaceTagvaluesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentmanagementWorkspaceTagvaluesParamsWithTimeout creates a new GetContentmanagementWorkspaceTagvaluesParams object
// with the ability to set a timeout on a request.
func NewGetContentmanagementWorkspaceTagvaluesParamsWithTimeout(timeout time.Duration) *GetContentmanagementWorkspaceTagvaluesParams {
	return &GetContentmanagementWorkspaceTagvaluesParams{
		timeout: timeout,
	}
}

// NewGetContentmanagementWorkspaceTagvaluesParamsWithContext creates a new GetContentmanagementWorkspaceTagvaluesParams object
// with the ability to set a context for a request.
func NewGetContentmanagementWorkspaceTagvaluesParamsWithContext(ctx context.Context) *GetContentmanagementWorkspaceTagvaluesParams {
	return &GetContentmanagementWorkspaceTagvaluesParams{
		Context: ctx,
	}
}

// NewGetContentmanagementWorkspaceTagvaluesParamsWithHTTPClient creates a new GetContentmanagementWorkspaceTagvaluesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContentmanagementWorkspaceTagvaluesParamsWithHTTPClient(client *http.Client) *GetContentmanagementWorkspaceTagvaluesParams {
	return &GetContentmanagementWorkspaceTagvaluesParams{
		HTTPClient: client,
	}
}

/*
GetContentmanagementWorkspaceTagvaluesParams contains all the parameters to send to the API endpoint

	for the get contentmanagement workspace tagvalues operation.

	Typically these are written to a http.Request.
*/
type GetContentmanagementWorkspaceTagvaluesParams struct {

	/* Expand.

	   Which fields, if any, to expand.
	*/
	Expand []string

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

	/* Value.

	   filter the list of tags returned
	*/
	Value *string

	/* WorkspaceID.

	   Workspace ID
	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contentmanagement workspace tagvalues params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentmanagementWorkspaceTagvaluesParams) WithDefaults() *GetContentmanagementWorkspaceTagvaluesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contentmanagement workspace tagvalues params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentmanagementWorkspaceTagvaluesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetContentmanagementWorkspaceTagvaluesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) WithTimeout(timeout time.Duration) *GetContentmanagementWorkspaceTagvaluesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) WithContext(ctx context.Context) *GetContentmanagementWorkspaceTagvaluesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) WithHTTPClient(client *http.Client) *GetContentmanagementWorkspaceTagvaluesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) WithExpand(expand []string) *GetContentmanagementWorkspaceTagvaluesParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithPageNumber adds the pageNumber to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) WithPageNumber(pageNumber *int32) *GetContentmanagementWorkspaceTagvaluesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) WithPageSize(pageSize *int32) *GetContentmanagementWorkspaceTagvaluesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithValue adds the value to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) WithValue(value *string) *GetContentmanagementWorkspaceTagvaluesParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) SetValue(value *string) {
	o.Value = value
}

// WithWorkspaceID adds the workspaceID to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) WithWorkspaceID(workspaceID string) *GetContentmanagementWorkspaceTagvaluesParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get contentmanagement workspace tagvalues params
func (o *GetContentmanagementWorkspaceTagvaluesParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentmanagementWorkspaceTagvaluesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
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

	if o.Value != nil {

		// query param value
		var qrValue string

		if o.Value != nil {
			qrValue = *o.Value
		}
		qValue := qrValue
		if qValue != "" {

			if err := r.SetQueryParam("value", qValue); err != nil {
				return err
			}
		}
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", o.WorkspaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetContentmanagementWorkspaceTagvalues binds the parameter expand
func (o *GetContentmanagementWorkspaceTagvaluesParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}
