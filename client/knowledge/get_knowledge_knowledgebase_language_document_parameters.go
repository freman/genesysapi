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

// NewGetKnowledgeKnowledgebaseLanguageDocumentParams creates a new GetKnowledgeKnowledgebaseLanguageDocumentParams object
// with the default values initialized.
func NewGetKnowledgeKnowledgebaseLanguageDocumentParams() *GetKnowledgeKnowledgebaseLanguageDocumentParams {
	var ()
	return &GetKnowledgeKnowledgebaseLanguageDocumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseLanguageDocumentParamsWithTimeout creates a new GetKnowledgeKnowledgebaseLanguageDocumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKnowledgeKnowledgebaseLanguageDocumentParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseLanguageDocumentParams {
	var ()
	return &GetKnowledgeKnowledgebaseLanguageDocumentParams{

		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseLanguageDocumentParamsWithContext creates a new GetKnowledgeKnowledgebaseLanguageDocumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKnowledgeKnowledgebaseLanguageDocumentParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseLanguageDocumentParams {
	var ()
	return &GetKnowledgeKnowledgebaseLanguageDocumentParams{

		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseLanguageDocumentParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseLanguageDocumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKnowledgeKnowledgebaseLanguageDocumentParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseLanguageDocumentParams {
	var ()
	return &GetKnowledgeKnowledgebaseLanguageDocumentParams{
		HTTPClient: client,
	}
}

/*GetKnowledgeKnowledgebaseLanguageDocumentParams contains all the parameters to send to the API endpoint
for the get knowledge knowledgebase language document operation typically these are written to a http.Request
*/
type GetKnowledgeKnowledgebaseLanguageDocumentParams struct {

	/*DocumentID
	  Document ID

	*/
	DocumentID string
	/*KnowledgeBaseID
	  Knowledge base ID

	*/
	KnowledgeBaseID string
	/*LanguageCode
	  Language code, format: iso2-LOCALE

	*/
	LanguageCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get knowledge knowledgebase language document params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseLanguageDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase language document params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase language document params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseLanguageDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase language document params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase language document params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseLanguageDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase language document params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the get knowledge knowledgebase language document params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) WithDocumentID(documentID string) *GetKnowledgeKnowledgebaseLanguageDocumentParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get knowledge knowledgebase language document params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase language document params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseLanguageDocumentParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase language document params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithLanguageCode adds the languageCode to the get knowledge knowledgebase language document params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) WithLanguageCode(languageCode string) *GetKnowledgeKnowledgebaseLanguageDocumentParams {
	o.SetLanguageCode(languageCode)
	return o
}

// SetLanguageCode adds the languageCode to the get knowledge knowledgebase language document params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) SetLanguageCode(languageCode string) {
	o.LanguageCode = languageCode
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseLanguageDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param languageCode
	if err := r.SetPathParam("languageCode", o.LanguageCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
