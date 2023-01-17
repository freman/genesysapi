// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

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

// NewGetTelephonyProvidersEdgesTimezonesParams creates a new GetTelephonyProvidersEdgesTimezonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTelephonyProvidersEdgesTimezonesParams() *GetTelephonyProvidersEdgesTimezonesParams {
	return &GetTelephonyProvidersEdgesTimezonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesTimezonesParamsWithTimeout creates a new GetTelephonyProvidersEdgesTimezonesParams object
// with the ability to set a timeout on a request.
func NewGetTelephonyProvidersEdgesTimezonesParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesTimezonesParams {
	return &GetTelephonyProvidersEdgesTimezonesParams{
		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesTimezonesParamsWithContext creates a new GetTelephonyProvidersEdgesTimezonesParams object
// with the ability to set a context for a request.
func NewGetTelephonyProvidersEdgesTimezonesParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesTimezonesParams {
	return &GetTelephonyProvidersEdgesTimezonesParams{
		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesTimezonesParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesTimezonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTelephonyProvidersEdgesTimezonesParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesTimezonesParams {
	return &GetTelephonyProvidersEdgesTimezonesParams{
		HTTPClient: client,
	}
}

/*
GetTelephonyProvidersEdgesTimezonesParams contains all the parameters to send to the API endpoint

	for the get telephony providers edges timezones operation.

	Typically these are written to a http.Request.
*/
type GetTelephonyProvidersEdgesTimezonesParams struct {

	/* PageNumber.

	   Page number

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   Page size

	   Format: int32
	   Default: 1000
	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get telephony providers edges timezones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesTimezonesParams) WithDefaults() *GetTelephonyProvidersEdgesTimezonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get telephony providers edges timezones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesTimezonesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(1000)
	)

	val := GetTelephonyProvidersEdgesTimezonesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get telephony providers edges timezones params
func (o *GetTelephonyProvidersEdgesTimezonesParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesTimezonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges timezones params
func (o *GetTelephonyProvidersEdgesTimezonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges timezones params
func (o *GetTelephonyProvidersEdgesTimezonesParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesTimezonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges timezones params
func (o *GetTelephonyProvidersEdgesTimezonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges timezones params
func (o *GetTelephonyProvidersEdgesTimezonesParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesTimezonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges timezones params
func (o *GetTelephonyProvidersEdgesTimezonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get telephony providers edges timezones params
func (o *GetTelephonyProvidersEdgesTimezonesParams) WithPageNumber(pageNumber *int32) *GetTelephonyProvidersEdgesTimezonesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get telephony providers edges timezones params
func (o *GetTelephonyProvidersEdgesTimezonesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get telephony providers edges timezones params
func (o *GetTelephonyProvidersEdgesTimezonesParams) WithPageSize(pageSize *int32) *GetTelephonyProvidersEdgesTimezonesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get telephony providers edges timezones params
func (o *GetTelephonyProvidersEdgesTimezonesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesTimezonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
