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
	"github.com/go-openapi/swag"
)

// NewGetKnowledgeKnowledgebaseDocumentVersionParams creates a new GetKnowledgeKnowledgebaseDocumentVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKnowledgeKnowledgebaseDocumentVersionParams() *GetKnowledgeKnowledgebaseDocumentVersionParams {
	return &GetKnowledgeKnowledgebaseDocumentVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentVersionParamsWithTimeout creates a new GetKnowledgeKnowledgebaseDocumentVersionParams object
// with the ability to set a timeout on a request.
func NewGetKnowledgeKnowledgebaseDocumentVersionParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseDocumentVersionParams {
	return &GetKnowledgeKnowledgebaseDocumentVersionParams{
		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentVersionParamsWithContext creates a new GetKnowledgeKnowledgebaseDocumentVersionParams object
// with the ability to set a context for a request.
func NewGetKnowledgeKnowledgebaseDocumentVersionParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseDocumentVersionParams {
	return &GetKnowledgeKnowledgebaseDocumentVersionParams{
		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentVersionParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseDocumentVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKnowledgeKnowledgebaseDocumentVersionParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseDocumentVersionParams {
	return &GetKnowledgeKnowledgebaseDocumentVersionParams{
		HTTPClient: client,
	}
}

/*
GetKnowledgeKnowledgebaseDocumentVersionParams contains all the parameters to send to the API endpoint

	for the get knowledge knowledgebase document version operation.

	Typically these are written to a http.Request.
*/
type GetKnowledgeKnowledgebaseDocumentVersionParams struct {

	/* DocumentID.

	   Globally unique identifier for the document.
	*/
	DocumentID string

	/* Expand.

	   The specified entity attributes will be filled. Comma separated values expected.
	*/
	Expand []string

	/* KnowledgeBaseID.

	   Globally unique identifier for the knowledge base.
	*/
	KnowledgeBaseID string

	/* VersionID.

	   Globally unique identifier for the document version.
	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get knowledge knowledgebase document version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) WithDefaults() *GetKnowledgeKnowledgebaseDocumentVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get knowledge knowledgebase document version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseDocumentVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseDocumentVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseDocumentVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) WithDocumentID(documentID string) *GetKnowledgeKnowledgebaseDocumentVersionParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithExpand adds the expand to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) WithExpand(expand []string) *GetKnowledgeKnowledgebaseDocumentVersionParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseDocumentVersionParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithVersionID adds the versionID to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) WithVersionID(versionID string) *GetKnowledgeKnowledgebaseDocumentVersionParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get knowledge knowledgebase document version params
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param documentId
	if err := r.SetPathParam("documentId", o.DocumentID); err != nil {
		return err
	}

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
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

// bindParamGetKnowledgeKnowledgebaseDocumentVersion binds the parameter expand
func (o *GetKnowledgeKnowledgebaseDocumentVersionParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}
