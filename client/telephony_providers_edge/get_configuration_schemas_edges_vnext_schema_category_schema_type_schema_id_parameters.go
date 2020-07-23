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
)

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams creates a new GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams object
// with the default values initialized.
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams {
	var ()
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParamsWithTimeout creates a new GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParamsWithTimeout(timeout time.Duration) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams {
	var ()
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams{

		timeout: timeout,
	}
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParamsWithContext creates a new GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParamsWithContext(ctx context.Context) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams {
	var ()
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams{

		Context: ctx,
	}
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParamsWithHTTPClient creates a new GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParamsWithHTTPClient(client *http.Client) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams {
	var ()
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams{
		HTTPClient: client,
	}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams contains all the parameters to send to the API endpoint
for the get configuration schemas edges vnext schema category schema type schema Id operation typically these are written to a http.Request
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams struct {

	/*SchemaCategory
	  Schema category

	*/
	SchemaCategory string
	/*SchemaID
	  Schema ID

	*/
	SchemaID string
	/*SchemaType
	  Schema type

	*/
	SchemaType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get configuration schemas edges vnext schema category schema type schema Id params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) WithTimeout(timeout time.Duration) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get configuration schemas edges vnext schema category schema type schema Id params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get configuration schemas edges vnext schema category schema type schema Id params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) WithContext(ctx context.Context) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get configuration schemas edges vnext schema category schema type schema Id params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get configuration schemas edges vnext schema category schema type schema Id params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) WithHTTPClient(client *http.Client) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get configuration schemas edges vnext schema category schema type schema Id params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSchemaCategory adds the schemaCategory to the get configuration schemas edges vnext schema category schema type schema Id params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) WithSchemaCategory(schemaCategory string) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams {
	o.SetSchemaCategory(schemaCategory)
	return o
}

// SetSchemaCategory adds the schemaCategory to the get configuration schemas edges vnext schema category schema type schema Id params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) SetSchemaCategory(schemaCategory string) {
	o.SchemaCategory = schemaCategory
}

// WithSchemaID adds the schemaID to the get configuration schemas edges vnext schema category schema type schema Id params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) WithSchemaID(schemaID string) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams {
	o.SetSchemaID(schemaID)
	return o
}

// SetSchemaID adds the schemaId to the get configuration schemas edges vnext schema category schema type schema Id params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) SetSchemaID(schemaID string) {
	o.SchemaID = schemaID
}

// WithSchemaType adds the schemaType to the get configuration schemas edges vnext schema category schema type schema Id params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) WithSchemaType(schemaType string) *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams {
	o.SetSchemaType(schemaType)
	return o
}

// SetSchemaType adds the schemaType to the get configuration schemas edges vnext schema category schema type schema Id params
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) SetSchemaType(schemaType string) {
	o.SchemaType = schemaType
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param schemaCategory
	if err := r.SetPathParam("schemaCategory", o.SchemaCategory); err != nil {
		return err
	}

	// path param schemaId
	if err := r.SetPathParam("schemaId", o.SchemaID); err != nil {
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
