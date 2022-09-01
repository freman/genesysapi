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

// NewGetKnowledgeKnowledgebaseExportJobParams creates a new GetKnowledgeKnowledgebaseExportJobParams object
// with the default values initialized.
func NewGetKnowledgeKnowledgebaseExportJobParams() *GetKnowledgeKnowledgebaseExportJobParams {
	var ()
	return &GetKnowledgeKnowledgebaseExportJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseExportJobParamsWithTimeout creates a new GetKnowledgeKnowledgebaseExportJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKnowledgeKnowledgebaseExportJobParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseExportJobParams {
	var ()
	return &GetKnowledgeKnowledgebaseExportJobParams{

		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseExportJobParamsWithContext creates a new GetKnowledgeKnowledgebaseExportJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKnowledgeKnowledgebaseExportJobParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseExportJobParams {
	var ()
	return &GetKnowledgeKnowledgebaseExportJobParams{

		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseExportJobParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseExportJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKnowledgeKnowledgebaseExportJobParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseExportJobParams {
	var ()
	return &GetKnowledgeKnowledgebaseExportJobParams{
		HTTPClient: client,
	}
}

/*GetKnowledgeKnowledgebaseExportJobParams contains all the parameters to send to the API endpoint
for the get knowledge knowledgebase export job operation typically these are written to a http.Request
*/
type GetKnowledgeKnowledgebaseExportJobParams struct {

	/*ExportJobID
	  Export job ID

	*/
	ExportJobID string
	/*KnowledgeBaseID
	  Knowledge base ID

	*/
	KnowledgeBaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get knowledge knowledgebase export job params
func (o *GetKnowledgeKnowledgebaseExportJobParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseExportJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase export job params
func (o *GetKnowledgeKnowledgebaseExportJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase export job params
func (o *GetKnowledgeKnowledgebaseExportJobParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseExportJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase export job params
func (o *GetKnowledgeKnowledgebaseExportJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase export job params
func (o *GetKnowledgeKnowledgebaseExportJobParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseExportJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase export job params
func (o *GetKnowledgeKnowledgebaseExportJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExportJobID adds the exportJobID to the get knowledge knowledgebase export job params
func (o *GetKnowledgeKnowledgebaseExportJobParams) WithExportJobID(exportJobID string) *GetKnowledgeKnowledgebaseExportJobParams {
	o.SetExportJobID(exportJobID)
	return o
}

// SetExportJobID adds the exportJobId to the get knowledge knowledgebase export job params
func (o *GetKnowledgeKnowledgebaseExportJobParams) SetExportJobID(exportJobID string) {
	o.ExportJobID = exportJobID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase export job params
func (o *GetKnowledgeKnowledgebaseExportJobParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseExportJobParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase export job params
func (o *GetKnowledgeKnowledgebaseExportJobParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseExportJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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