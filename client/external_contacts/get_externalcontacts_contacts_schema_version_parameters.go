// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

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

// NewGetExternalcontactsContactsSchemaVersionParams creates a new GetExternalcontactsContactsSchemaVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExternalcontactsContactsSchemaVersionParams() *GetExternalcontactsContactsSchemaVersionParams {
	return &GetExternalcontactsContactsSchemaVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalcontactsContactsSchemaVersionParamsWithTimeout creates a new GetExternalcontactsContactsSchemaVersionParams object
// with the ability to set a timeout on a request.
func NewGetExternalcontactsContactsSchemaVersionParamsWithTimeout(timeout time.Duration) *GetExternalcontactsContactsSchemaVersionParams {
	return &GetExternalcontactsContactsSchemaVersionParams{
		timeout: timeout,
	}
}

// NewGetExternalcontactsContactsSchemaVersionParamsWithContext creates a new GetExternalcontactsContactsSchemaVersionParams object
// with the ability to set a context for a request.
func NewGetExternalcontactsContactsSchemaVersionParamsWithContext(ctx context.Context) *GetExternalcontactsContactsSchemaVersionParams {
	return &GetExternalcontactsContactsSchemaVersionParams{
		Context: ctx,
	}
}

// NewGetExternalcontactsContactsSchemaVersionParamsWithHTTPClient creates a new GetExternalcontactsContactsSchemaVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExternalcontactsContactsSchemaVersionParamsWithHTTPClient(client *http.Client) *GetExternalcontactsContactsSchemaVersionParams {
	return &GetExternalcontactsContactsSchemaVersionParams{
		HTTPClient: client,
	}
}

/*
GetExternalcontactsContactsSchemaVersionParams contains all the parameters to send to the API endpoint

	for the get externalcontacts contacts schema version operation.

	Typically these are written to a http.Request.
*/
type GetExternalcontactsContactsSchemaVersionParams struct {

	/* SchemaID.

	   Schema ID
	*/
	SchemaID string

	/* VersionID.

	   Schema version
	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get externalcontacts contacts schema version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalcontactsContactsSchemaVersionParams) WithDefaults() *GetExternalcontactsContactsSchemaVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get externalcontacts contacts schema version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalcontactsContactsSchemaVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get externalcontacts contacts schema version params
func (o *GetExternalcontactsContactsSchemaVersionParams) WithTimeout(timeout time.Duration) *GetExternalcontactsContactsSchemaVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get externalcontacts contacts schema version params
func (o *GetExternalcontactsContactsSchemaVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get externalcontacts contacts schema version params
func (o *GetExternalcontactsContactsSchemaVersionParams) WithContext(ctx context.Context) *GetExternalcontactsContactsSchemaVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get externalcontacts contacts schema version params
func (o *GetExternalcontactsContactsSchemaVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get externalcontacts contacts schema version params
func (o *GetExternalcontactsContactsSchemaVersionParams) WithHTTPClient(client *http.Client) *GetExternalcontactsContactsSchemaVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get externalcontacts contacts schema version params
func (o *GetExternalcontactsContactsSchemaVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSchemaID adds the schemaID to the get externalcontacts contacts schema version params
func (o *GetExternalcontactsContactsSchemaVersionParams) WithSchemaID(schemaID string) *GetExternalcontactsContactsSchemaVersionParams {
	o.SetSchemaID(schemaID)
	return o
}

// SetSchemaID adds the schemaId to the get externalcontacts contacts schema version params
func (o *GetExternalcontactsContactsSchemaVersionParams) SetSchemaID(schemaID string) {
	o.SchemaID = schemaID
}

// WithVersionID adds the versionID to the get externalcontacts contacts schema version params
func (o *GetExternalcontactsContactsSchemaVersionParams) WithVersionID(versionID string) *GetExternalcontactsContactsSchemaVersionParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get externalcontacts contacts schema version params
func (o *GetExternalcontactsContactsSchemaVersionParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalcontactsContactsSchemaVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param schemaId
	if err := r.SetPathParam("schemaId", o.SchemaID); err != nil {
		return err
	}

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
