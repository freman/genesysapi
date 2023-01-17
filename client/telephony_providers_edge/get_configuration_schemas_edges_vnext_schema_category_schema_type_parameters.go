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

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams creates a new GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParamsWithTimeout creates a new GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams object
// with the ability to set a timeout on a request.
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParamsWithTimeout(timeout time.Duration) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams{
		timeout: timeout,
	}
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParamsWithContext creates a new GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams object
// with the ability to set a context for a request.
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParamsWithContext(ctx context.Context) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams{
		Context: ctx,
	}
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParamsWithHTTPClient creates a new GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParamsWithHTTPClient(client *http.Client) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams{
		HTTPClient: client,
	}
}

/*
GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams contains all the parameters to send to the API endpoint

	for the get configuration schemas edges vnext schema category schema type operation.

	Typically these are written to a http.Request.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams struct {

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

	/* SchemaType.

	   Schema type
	*/
	SchemaType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get configuration schemas edges vnext schema category schema type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) WithDefaults() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get configuration schemas edges vnext schema category schema type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) WithTimeout(timeout time.Duration) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) WithContext(ctx context.Context) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) WithHTTPClient(client *http.Client) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) WithPageNumber(pageNumber *int32) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) WithPageSize(pageSize *int32) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSchemaCategory adds the schemaCategory to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) WithSchemaCategory(schemaCategory string) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams {
	o.SetSchemaCategory(schemaCategory)
	return o
}

// SetSchemaCategory adds the schemaCategory to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) SetSchemaCategory(schemaCategory string) {
	o.SchemaCategory = schemaCategory
}

// WithSchemaType adds the schemaType to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) WithSchemaType(schemaType string) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams {
	o.SetSchemaType(schemaType)
	return o
}

// SetSchemaType adds the schemaType to the get configuration schemas edges vnext schema category schema type params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) SetSchemaType(schemaType string) {
	o.SchemaType = schemaType
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param schemaType
	if err := r.SetPathParam("schemaType", o.SchemaType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
