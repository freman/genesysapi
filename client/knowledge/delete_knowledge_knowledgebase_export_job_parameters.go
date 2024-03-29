// Code generated by go-swagger; DO NOT EDIT.

package knowledge

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

// NewDeleteKnowledgeKnowledgebaseExportJobParams creates a new DeleteKnowledgeKnowledgebaseExportJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteKnowledgeKnowledgebaseExportJobParams() *DeleteKnowledgeKnowledgebaseExportJobParams {
	return &DeleteKnowledgeKnowledgebaseExportJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKnowledgeKnowledgebaseExportJobParamsWithTimeout creates a new DeleteKnowledgeKnowledgebaseExportJobParams object
// with the ability to set a timeout on a request.
func NewDeleteKnowledgeKnowledgebaseExportJobParamsWithTimeout(timeout time.Duration) *DeleteKnowledgeKnowledgebaseExportJobParams {
	return &DeleteKnowledgeKnowledgebaseExportJobParams{
		timeout: timeout,
	}
}

// NewDeleteKnowledgeKnowledgebaseExportJobParamsWithContext creates a new DeleteKnowledgeKnowledgebaseExportJobParams object
// with the ability to set a context for a request.
func NewDeleteKnowledgeKnowledgebaseExportJobParamsWithContext(ctx context.Context) *DeleteKnowledgeKnowledgebaseExportJobParams {
	return &DeleteKnowledgeKnowledgebaseExportJobParams{
		Context: ctx,
	}
}

// NewDeleteKnowledgeKnowledgebaseExportJobParamsWithHTTPClient creates a new DeleteKnowledgeKnowledgebaseExportJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteKnowledgeKnowledgebaseExportJobParamsWithHTTPClient(client *http.Client) *DeleteKnowledgeKnowledgebaseExportJobParams {
	return &DeleteKnowledgeKnowledgebaseExportJobParams{
		HTTPClient: client,
	}
}

/*
DeleteKnowledgeKnowledgebaseExportJobParams contains all the parameters to send to the API endpoint

	for the delete knowledge knowledgebase export job operation.

	Typically these are written to a http.Request.
*/
type DeleteKnowledgeKnowledgebaseExportJobParams struct {

	/* ExportJobID.

	   Export job ID
	*/
	ExportJobID string

	/* KnowledgeBaseID.

	   Knowledge base ID
	*/
	KnowledgeBaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete knowledge knowledgebase export job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) WithDefaults() *DeleteKnowledgeKnowledgebaseExportJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete knowledge knowledgebase export job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete knowledge knowledgebase export job params
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) WithTimeout(timeout time.Duration) *DeleteKnowledgeKnowledgebaseExportJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete knowledge knowledgebase export job params
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete knowledge knowledgebase export job params
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) WithContext(ctx context.Context) *DeleteKnowledgeKnowledgebaseExportJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete knowledge knowledgebase export job params
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete knowledge knowledgebase export job params
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) WithHTTPClient(client *http.Client) *DeleteKnowledgeKnowledgebaseExportJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete knowledge knowledgebase export job params
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExportJobID adds the exportJobID to the delete knowledge knowledgebase export job params
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) WithExportJobID(exportJobID string) *DeleteKnowledgeKnowledgebaseExportJobParams {
	o.SetExportJobID(exportJobID)
	return o
}

// SetExportJobID adds the exportJobId to the delete knowledge knowledgebase export job params
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) SetExportJobID(exportJobID string) {
	o.ExportJobID = exportJobID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the delete knowledge knowledgebase export job params
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) WithKnowledgeBaseID(knowledgeBaseID string) *DeleteKnowledgeKnowledgebaseExportJobParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the delete knowledge knowledgebase export job params
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKnowledgeKnowledgebaseExportJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param exportJobId
	if err := r.SetPathParam("exportJobId", o.ExportJobID); err != nil {
		return err
	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
