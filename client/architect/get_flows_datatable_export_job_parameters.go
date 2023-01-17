// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewGetFlowsDatatableExportJobParams creates a new GetFlowsDatatableExportJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFlowsDatatableExportJobParams() *GetFlowsDatatableExportJobParams {
	return &GetFlowsDatatableExportJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowsDatatableExportJobParamsWithTimeout creates a new GetFlowsDatatableExportJobParams object
// with the ability to set a timeout on a request.
func NewGetFlowsDatatableExportJobParamsWithTimeout(timeout time.Duration) *GetFlowsDatatableExportJobParams {
	return &GetFlowsDatatableExportJobParams{
		timeout: timeout,
	}
}

// NewGetFlowsDatatableExportJobParamsWithContext creates a new GetFlowsDatatableExportJobParams object
// with the ability to set a context for a request.
func NewGetFlowsDatatableExportJobParamsWithContext(ctx context.Context) *GetFlowsDatatableExportJobParams {
	return &GetFlowsDatatableExportJobParams{
		Context: ctx,
	}
}

// NewGetFlowsDatatableExportJobParamsWithHTTPClient creates a new GetFlowsDatatableExportJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFlowsDatatableExportJobParamsWithHTTPClient(client *http.Client) *GetFlowsDatatableExportJobParams {
	return &GetFlowsDatatableExportJobParams{
		HTTPClient: client,
	}
}

/*
GetFlowsDatatableExportJobParams contains all the parameters to send to the API endpoint

	for the get flows datatable export job operation.

	Typically these are written to a http.Request.
*/
type GetFlowsDatatableExportJobParams struct {

	/* DatatableID.

	   id of datatable
	*/
	DatatableID string

	/* ExportJobID.

	   id of export job
	*/
	ExportJobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get flows datatable export job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowsDatatableExportJobParams) WithDefaults() *GetFlowsDatatableExportJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get flows datatable export job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFlowsDatatableExportJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get flows datatable export job params
func (o *GetFlowsDatatableExportJobParams) WithTimeout(timeout time.Duration) *GetFlowsDatatableExportJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flows datatable export job params
func (o *GetFlowsDatatableExportJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flows datatable export job params
func (o *GetFlowsDatatableExportJobParams) WithContext(ctx context.Context) *GetFlowsDatatableExportJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flows datatable export job params
func (o *GetFlowsDatatableExportJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flows datatable export job params
func (o *GetFlowsDatatableExportJobParams) WithHTTPClient(client *http.Client) *GetFlowsDatatableExportJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flows datatable export job params
func (o *GetFlowsDatatableExportJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatatableID adds the datatableID to the get flows datatable export job params
func (o *GetFlowsDatatableExportJobParams) WithDatatableID(datatableID string) *GetFlowsDatatableExportJobParams {
	o.SetDatatableID(datatableID)
	return o
}

// SetDatatableID adds the datatableId to the get flows datatable export job params
func (o *GetFlowsDatatableExportJobParams) SetDatatableID(datatableID string) {
	o.DatatableID = datatableID
}

// WithExportJobID adds the exportJobID to the get flows datatable export job params
func (o *GetFlowsDatatableExportJobParams) WithExportJobID(exportJobID string) *GetFlowsDatatableExportJobParams {
	o.SetExportJobID(exportJobID)
	return o
}

// SetExportJobID adds the exportJobId to the get flows datatable export job params
func (o *GetFlowsDatatableExportJobParams) SetExportJobID(exportJobID string) {
	o.ExportJobID = exportJobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowsDatatableExportJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datatableId
	if err := r.SetPathParam("datatableId", o.DatatableID); err != nil {
		return err
	}

	// path param exportJobId
	if err := r.SetPathParam("exportJobId", o.ExportJobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
