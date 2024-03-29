// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetIntegrationsTypeConfigschemaParams creates a new GetIntegrationsTypeConfigschemaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIntegrationsTypeConfigschemaParams() *GetIntegrationsTypeConfigschemaParams {
	return &GetIntegrationsTypeConfigschemaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationsTypeConfigschemaParamsWithTimeout creates a new GetIntegrationsTypeConfigschemaParams object
// with the ability to set a timeout on a request.
func NewGetIntegrationsTypeConfigschemaParamsWithTimeout(timeout time.Duration) *GetIntegrationsTypeConfigschemaParams {
	return &GetIntegrationsTypeConfigschemaParams{
		timeout: timeout,
	}
}

// NewGetIntegrationsTypeConfigschemaParamsWithContext creates a new GetIntegrationsTypeConfigschemaParams object
// with the ability to set a context for a request.
func NewGetIntegrationsTypeConfigschemaParamsWithContext(ctx context.Context) *GetIntegrationsTypeConfigschemaParams {
	return &GetIntegrationsTypeConfigschemaParams{
		Context: ctx,
	}
}

// NewGetIntegrationsTypeConfigschemaParamsWithHTTPClient creates a new GetIntegrationsTypeConfigschemaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIntegrationsTypeConfigschemaParamsWithHTTPClient(client *http.Client) *GetIntegrationsTypeConfigschemaParams {
	return &GetIntegrationsTypeConfigschemaParams{
		HTTPClient: client,
	}
}

/*
GetIntegrationsTypeConfigschemaParams contains all the parameters to send to the API endpoint

	for the get integrations type configschema operation.

	Typically these are written to a http.Request.
*/
type GetIntegrationsTypeConfigschemaParams struct {

	/* ConfigType.

	   Config schema type
	*/
	ConfigType string

	/* TypeID.

	   Integration Type Id
	*/
	TypeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get integrations type configschema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationsTypeConfigschemaParams) WithDefaults() *GetIntegrationsTypeConfigschemaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get integrations type configschema params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationsTypeConfigschemaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get integrations type configschema params
func (o *GetIntegrationsTypeConfigschemaParams) WithTimeout(timeout time.Duration) *GetIntegrationsTypeConfigschemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrations type configschema params
func (o *GetIntegrationsTypeConfigschemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrations type configschema params
func (o *GetIntegrationsTypeConfigschemaParams) WithContext(ctx context.Context) *GetIntegrationsTypeConfigschemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrations type configschema params
func (o *GetIntegrationsTypeConfigschemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrations type configschema params
func (o *GetIntegrationsTypeConfigschemaParams) WithHTTPClient(client *http.Client) *GetIntegrationsTypeConfigschemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrations type configschema params
func (o *GetIntegrationsTypeConfigschemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigType adds the configType to the get integrations type configschema params
func (o *GetIntegrationsTypeConfigschemaParams) WithConfigType(configType string) *GetIntegrationsTypeConfigschemaParams {
	o.SetConfigType(configType)
	return o
}

// SetConfigType adds the configType to the get integrations type configschema params
func (o *GetIntegrationsTypeConfigschemaParams) SetConfigType(configType string) {
	o.ConfigType = configType
}

// WithTypeID adds the typeID to the get integrations type configschema params
func (o *GetIntegrationsTypeConfigschemaParams) WithTypeID(typeID string) *GetIntegrationsTypeConfigschemaParams {
	o.SetTypeID(typeID)
	return o
}

// SetTypeID adds the typeId to the get integrations type configschema params
func (o *GetIntegrationsTypeConfigschemaParams) SetTypeID(typeID string) {
	o.TypeID = typeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationsTypeConfigschemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param configType
	if err := r.SetPathParam("configType", o.ConfigType); err != nil {
		return err
	}

	// path param typeId
	if err := r.SetPathParam("typeId", o.TypeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
