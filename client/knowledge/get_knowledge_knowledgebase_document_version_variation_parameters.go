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

// NewGetKnowledgeKnowledgebaseDocumentVersionVariationParams creates a new GetKnowledgeKnowledgebaseDocumentVersionVariationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKnowledgeKnowledgebaseDocumentVersionVariationParams() *GetKnowledgeKnowledgebaseDocumentVersionVariationParams {
	return &GetKnowledgeKnowledgebaseDocumentVersionVariationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentVersionVariationParamsWithTimeout creates a new GetKnowledgeKnowledgebaseDocumentVersionVariationParams object
// with the ability to set a timeout on a request.
func NewGetKnowledgeKnowledgebaseDocumentVersionVariationParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseDocumentVersionVariationParams {
	return &GetKnowledgeKnowledgebaseDocumentVersionVariationParams{
		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentVersionVariationParamsWithContext creates a new GetKnowledgeKnowledgebaseDocumentVersionVariationParams object
// with the ability to set a context for a request.
func NewGetKnowledgeKnowledgebaseDocumentVersionVariationParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseDocumentVersionVariationParams {
	return &GetKnowledgeKnowledgebaseDocumentVersionVariationParams{
		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentVersionVariationParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseDocumentVersionVariationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKnowledgeKnowledgebaseDocumentVersionVariationParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseDocumentVersionVariationParams {
	return &GetKnowledgeKnowledgebaseDocumentVersionVariationParams{
		HTTPClient: client,
	}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionVariationParams contains all the parameters to send to the API endpoint

	for the get knowledge knowledgebase document version variation operation.

	Typically these are written to a http.Request.
*/
type GetKnowledgeKnowledgebaseDocumentVersionVariationParams struct {

	/* DocumentID.

	   Globally unique identifier for the document.
	*/
	DocumentID string

	/* KnowledgeBaseID.

	   Globally unique identifier for the knowledge base.
	*/
	KnowledgeBaseID string

	/* VariationID.

	   Globally unique identifier for the document version variation.
	*/
	VariationID string

	/* VersionID.

	   Globally unique identifier for the document version.
	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get knowledge knowledgebase document version variation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) WithDefaults() *GetKnowledgeKnowledgebaseDocumentVersionVariationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get knowledge knowledgebase document version variation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseDocumentVersionVariationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseDocumentVersionVariationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseDocumentVersionVariationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) WithDocumentID(documentID string) *GetKnowledgeKnowledgebaseDocumentVersionVariationParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseDocumentVersionVariationParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithVariationID adds the variationID to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) WithVariationID(variationID string) *GetKnowledgeKnowledgebaseDocumentVersionVariationParams {
	o.SetVariationID(variationID)
	return o
}

// SetVariationID adds the variationId to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) SetVariationID(variationID string) {
	o.VariationID = variationID
}

// WithVersionID adds the versionID to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) WithVersionID(versionID string) *GetKnowledgeKnowledgebaseDocumentVersionVariationParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get knowledge knowledgebase document version variation params
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseDocumentVersionVariationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param documentId
	if err := r.SetPathParam("documentId", o.DocumentID); err != nil {
		return err
	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	// path param variationId
	if err := r.SetPathParam("variationId", o.VariationID); err != nil {
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
