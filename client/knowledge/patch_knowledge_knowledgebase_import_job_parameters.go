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

	"github.com/freman/genesysapi/models"
)

// NewPatchKnowledgeKnowledgebaseImportJobParams creates a new PatchKnowledgeKnowledgebaseImportJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchKnowledgeKnowledgebaseImportJobParams() *PatchKnowledgeKnowledgebaseImportJobParams {
	return &PatchKnowledgeKnowledgebaseImportJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchKnowledgeKnowledgebaseImportJobParamsWithTimeout creates a new PatchKnowledgeKnowledgebaseImportJobParams object
// with the ability to set a timeout on a request.
func NewPatchKnowledgeKnowledgebaseImportJobParamsWithTimeout(timeout time.Duration) *PatchKnowledgeKnowledgebaseImportJobParams {
	return &PatchKnowledgeKnowledgebaseImportJobParams{
		timeout: timeout,
	}
}

// NewPatchKnowledgeKnowledgebaseImportJobParamsWithContext creates a new PatchKnowledgeKnowledgebaseImportJobParams object
// with the ability to set a context for a request.
func NewPatchKnowledgeKnowledgebaseImportJobParamsWithContext(ctx context.Context) *PatchKnowledgeKnowledgebaseImportJobParams {
	return &PatchKnowledgeKnowledgebaseImportJobParams{
		Context: ctx,
	}
}

// NewPatchKnowledgeKnowledgebaseImportJobParamsWithHTTPClient creates a new PatchKnowledgeKnowledgebaseImportJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchKnowledgeKnowledgebaseImportJobParamsWithHTTPClient(client *http.Client) *PatchKnowledgeKnowledgebaseImportJobParams {
	return &PatchKnowledgeKnowledgebaseImportJobParams{
		HTTPClient: client,
	}
}

/*
PatchKnowledgeKnowledgebaseImportJobParams contains all the parameters to send to the API endpoint

	for the patch knowledge knowledgebase import job operation.

	Typically these are written to a http.Request.
*/
type PatchKnowledgeKnowledgebaseImportJobParams struct {

	// Body.
	Body *models.ImportStatusRequest

	/* ImportJobID.

	   Import job ID
	*/
	ImportJobID string

	/* KnowledgeBaseID.

	   Knowledge base ID
	*/
	KnowledgeBaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch knowledge knowledgebase import job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchKnowledgeKnowledgebaseImportJobParams) WithDefaults() *PatchKnowledgeKnowledgebaseImportJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch knowledge knowledgebase import job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchKnowledgeKnowledgebaseImportJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch knowledge knowledgebase import job params
func (o *PatchKnowledgeKnowledgebaseImportJobParams) WithTimeout(timeout time.Duration) *PatchKnowledgeKnowledgebaseImportJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch knowledge knowledgebase import job params
func (o *PatchKnowledgeKnowledgebaseImportJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch knowledge knowledgebase import job params
func (o *PatchKnowledgeKnowledgebaseImportJobParams) WithContext(ctx context.Context) *PatchKnowledgeKnowledgebaseImportJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch knowledge knowledgebase import job params
func (o *PatchKnowledgeKnowledgebaseImportJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch knowledge knowledgebase import job params
func (o *PatchKnowledgeKnowledgebaseImportJobParams) WithHTTPClient(client *http.Client) *PatchKnowledgeKnowledgebaseImportJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch knowledge knowledgebase import job params
func (o *PatchKnowledgeKnowledgebaseImportJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch knowledge knowledgebase import job params
func (o *PatchKnowledgeKnowledgebaseImportJobParams) WithBody(body *models.ImportStatusRequest) *PatchKnowledgeKnowledgebaseImportJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch knowledge knowledgebase import job params
func (o *PatchKnowledgeKnowledgebaseImportJobParams) SetBody(body *models.ImportStatusRequest) {
	o.Body = body
}

// WithImportJobID adds the importJobID to the patch knowledge knowledgebase import job params
func (o *PatchKnowledgeKnowledgebaseImportJobParams) WithImportJobID(importJobID string) *PatchKnowledgeKnowledgebaseImportJobParams {
	o.SetImportJobID(importJobID)
	return o
}

// SetImportJobID adds the importJobId to the patch knowledge knowledgebase import job params
func (o *PatchKnowledgeKnowledgebaseImportJobParams) SetImportJobID(importJobID string) {
	o.ImportJobID = importJobID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the patch knowledge knowledgebase import job params
func (o *PatchKnowledgeKnowledgebaseImportJobParams) WithKnowledgeBaseID(knowledgeBaseID string) *PatchKnowledgeKnowledgebaseImportJobParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the patch knowledge knowledgebase import job params
func (o *PatchKnowledgeKnowledgebaseImportJobParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchKnowledgeKnowledgebaseImportJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param importJobId
	if err := r.SetPathParam("importJobId", o.ImportJobID); err != nil {
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
