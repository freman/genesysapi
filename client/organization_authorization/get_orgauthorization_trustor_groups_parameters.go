// Code generated by go-swagger; DO NOT EDIT.

package organization_authorization

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

// NewGetOrgauthorizationTrustorGroupsParams creates a new GetOrgauthorizationTrustorGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrgauthorizationTrustorGroupsParams() *GetOrgauthorizationTrustorGroupsParams {
	return &GetOrgauthorizationTrustorGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgauthorizationTrustorGroupsParamsWithTimeout creates a new GetOrgauthorizationTrustorGroupsParams object
// with the ability to set a timeout on a request.
func NewGetOrgauthorizationTrustorGroupsParamsWithTimeout(timeout time.Duration) *GetOrgauthorizationTrustorGroupsParams {
	return &GetOrgauthorizationTrustorGroupsParams{
		timeout: timeout,
	}
}

// NewGetOrgauthorizationTrustorGroupsParamsWithContext creates a new GetOrgauthorizationTrustorGroupsParams object
// with the ability to set a context for a request.
func NewGetOrgauthorizationTrustorGroupsParamsWithContext(ctx context.Context) *GetOrgauthorizationTrustorGroupsParams {
	return &GetOrgauthorizationTrustorGroupsParams{
		Context: ctx,
	}
}

// NewGetOrgauthorizationTrustorGroupsParamsWithHTTPClient creates a new GetOrgauthorizationTrustorGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrgauthorizationTrustorGroupsParamsWithHTTPClient(client *http.Client) *GetOrgauthorizationTrustorGroupsParams {
	return &GetOrgauthorizationTrustorGroupsParams{
		HTTPClient: client,
	}
}

/*
GetOrgauthorizationTrustorGroupsParams contains all the parameters to send to the API endpoint

	for the get orgauthorization trustor groups operation.

	Typically these are written to a http.Request.
*/
type GetOrgauthorizationTrustorGroupsParams struct {

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

	/* TrustorOrgID.

	   Trustee Organization Id
	*/
	TrustorOrgID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get orgauthorization trustor groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgauthorizationTrustorGroupsParams) WithDefaults() *GetOrgauthorizationTrustorGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get orgauthorization trustor groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgauthorizationTrustorGroupsParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetOrgauthorizationTrustorGroupsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get orgauthorization trustor groups params
func (o *GetOrgauthorizationTrustorGroupsParams) WithTimeout(timeout time.Duration) *GetOrgauthorizationTrustorGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orgauthorization trustor groups params
func (o *GetOrgauthorizationTrustorGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orgauthorization trustor groups params
func (o *GetOrgauthorizationTrustorGroupsParams) WithContext(ctx context.Context) *GetOrgauthorizationTrustorGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orgauthorization trustor groups params
func (o *GetOrgauthorizationTrustorGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orgauthorization trustor groups params
func (o *GetOrgauthorizationTrustorGroupsParams) WithHTTPClient(client *http.Client) *GetOrgauthorizationTrustorGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orgauthorization trustor groups params
func (o *GetOrgauthorizationTrustorGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get orgauthorization trustor groups params
func (o *GetOrgauthorizationTrustorGroupsParams) WithPageNumber(pageNumber *int32) *GetOrgauthorizationTrustorGroupsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get orgauthorization trustor groups params
func (o *GetOrgauthorizationTrustorGroupsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get orgauthorization trustor groups params
func (o *GetOrgauthorizationTrustorGroupsParams) WithPageSize(pageSize *int32) *GetOrgauthorizationTrustorGroupsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get orgauthorization trustor groups params
func (o *GetOrgauthorizationTrustorGroupsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithTrustorOrgID adds the trustorOrgID to the get orgauthorization trustor groups params
func (o *GetOrgauthorizationTrustorGroupsParams) WithTrustorOrgID(trustorOrgID string) *GetOrgauthorizationTrustorGroupsParams {
	o.SetTrustorOrgID(trustorOrgID)
	return o
}

// SetTrustorOrgID adds the trustorOrgId to the get orgauthorization trustor groups params
func (o *GetOrgauthorizationTrustorGroupsParams) SetTrustorOrgID(trustorOrgID string) {
	o.TrustorOrgID = trustorOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgauthorizationTrustorGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param trustorOrgId
	if err := r.SetPathParam("trustorOrgId", o.TrustorOrgID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}