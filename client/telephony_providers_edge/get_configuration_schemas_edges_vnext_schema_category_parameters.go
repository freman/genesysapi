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

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryParams creates a new GetConfigurationSchemasEdgesVnextSchemaCategoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryParams() *GetConfigurationSchemasEdgesVnextSchemaCategoryParams {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryParamsWithTimeout creates a new GetConfigurationSchemasEdgesVnextSchemaCategoryParams object
// with the ability to set a timeout on a request.
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryParamsWithTimeout(timeout time.Duration) *GetConfigurationSchemasEdgesVnextSchemaCategoryParams {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryParams{
		timeout: timeout,
	}
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryParamsWithContext creates a new GetConfigurationSchemasEdgesVnextSchemaCategoryParams object
// with the ability to set a context for a request.
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryParamsWithContext(ctx context.Context) *GetConfigurationSchemasEdgesVnextSchemaCategoryParams {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryParams{
		Context: ctx,
	}
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryParamsWithHTTPClient creates a new GetConfigurationSchemasEdgesVnextSchemaCategoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryParamsWithHTTPClient(client *http.Client) *GetConfigurationSchemasEdgesVnextSchemaCategoryParams {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryParams{
		HTTPClient: client,
	}
}

/*
GetConfigurationSchemasEdgesVnextSchemaCategoryParams contains all the parameters to send to the API endpoint

	for the get configuration schemas edges vnext schema category operation.

	Typically these are written to a http.Request.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryParams struct {

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

	/* SchemaCategory.

	   Schema category
	*/
	SchemaCategory string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get configuration schemas edges vnext schema category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) WithDefaults() *GetConfigurationSchemasEdgesVnextSchemaCategoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get configuration schemas edges vnext schema category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetConfigurationSchemasEdgesVnextSchemaCategoryParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get configuration schemas edges vnext schema category params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) WithTimeout(timeout time.Duration) *GetConfigurationSchemasEdgesVnextSchemaCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get configuration schemas edges vnext schema category params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get configuration schemas edges vnext schema category params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) WithContext(ctx context.Context) *GetConfigurationSchemasEdgesVnextSchemaCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get configuration schemas edges vnext schema category params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get configuration schemas edges vnext schema category params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) WithHTTPClient(client *http.Client) *GetConfigurationSchemasEdgesVnextSchemaCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get configuration schemas edges vnext schema category params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get configuration schemas edges vnext schema category params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) WithPageNumber(pageNumber *int32) *GetConfigurationSchemasEdgesVnextSchemaCategoryParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get configuration schemas edges vnext schema category params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get configuration schemas edges vnext schema category params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) WithPageSize(pageSize *int32) *GetConfigurationSchemasEdgesVnextSchemaCategoryParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get configuration schemas edges vnext schema category params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSchemaCategory adds the schemaCategory to the get configuration schemas edges vnext schema category params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) WithSchemaCategory(schemaCategory string) *GetConfigurationSchemasEdgesVnextSchemaCategoryParams {
	o.SetSchemaCategory(schemaCategory)
	return o
}

// SetSchemaCategory adds the schemaCategory to the get configuration schemas edges vnext schema category params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) SetSchemaCategory(schemaCategory string) {
	o.SchemaCategory = schemaCategory
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param schemaCategory
	if err := r.SetPathParam("schemaCategory", o.SchemaCategory); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
