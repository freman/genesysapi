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

// NewGetConfigurationSchemasEdgesVnextParams creates a new GetConfigurationSchemasEdgesVnextParams object
// with the default values initialized.
func NewGetConfigurationSchemasEdgesVnextParams() *GetConfigurationSchemasEdgesVnextParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetConfigurationSchemasEdgesVnextParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigurationSchemasEdgesVnextParamsWithTimeout creates a new GetConfigurationSchemasEdgesVnextParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConfigurationSchemasEdgesVnextParamsWithTimeout(timeout time.Duration) *GetConfigurationSchemasEdgesVnextParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetConfigurationSchemasEdgesVnextParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetConfigurationSchemasEdgesVnextParamsWithContext creates a new GetConfigurationSchemasEdgesVnextParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConfigurationSchemasEdgesVnextParamsWithContext(ctx context.Context) *GetConfigurationSchemasEdgesVnextParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetConfigurationSchemasEdgesVnextParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetConfigurationSchemasEdgesVnextParamsWithHTTPClient creates a new GetConfigurationSchemasEdgesVnextParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConfigurationSchemasEdgesVnextParamsWithHTTPClient(client *http.Client) *GetConfigurationSchemasEdgesVnextParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetConfigurationSchemasEdgesVnextParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetConfigurationSchemasEdgesVnextParams contains all the parameters to send to the API endpoint
for the get configuration schemas edges vnext operation typically these are written to a http.Request
*/
type GetConfigurationSchemasEdgesVnextParams struct {

	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get configuration schemas edges vnext params
func (o *GetConfigurationSchemasEdgesVnextParams) WithTimeout(timeout time.Duration) *GetConfigurationSchemasEdgesVnextParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get configuration schemas edges vnext params
func (o *GetConfigurationSchemasEdgesVnextParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get configuration schemas edges vnext params
func (o *GetConfigurationSchemasEdgesVnextParams) WithContext(ctx context.Context) *GetConfigurationSchemasEdgesVnextParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get configuration schemas edges vnext params
func (o *GetConfigurationSchemasEdgesVnextParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get configuration schemas edges vnext params
func (o *GetConfigurationSchemasEdgesVnextParams) WithHTTPClient(client *http.Client) *GetConfigurationSchemasEdgesVnextParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get configuration schemas edges vnext params
func (o *GetConfigurationSchemasEdgesVnextParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get configuration schemas edges vnext params
func (o *GetConfigurationSchemasEdgesVnextParams) WithPageNumber(pageNumber *int32) *GetConfigurationSchemasEdgesVnextParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get configuration schemas edges vnext params
func (o *GetConfigurationSchemasEdgesVnextParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get configuration schemas edges vnext params
func (o *GetConfigurationSchemasEdgesVnextParams) WithPageSize(pageSize *int32) *GetConfigurationSchemasEdgesVnextParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get configuration schemas edges vnext params
func (o *GetConfigurationSchemasEdgesVnextParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigurationSchemasEdgesVnextParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
