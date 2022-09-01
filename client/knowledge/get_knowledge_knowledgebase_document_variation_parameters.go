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

// NewGetKnowledgeKnowledgebaseDocumentVariationParams creates a new GetKnowledgeKnowledgebaseDocumentVariationParams object
// with the default values initialized.
func NewGetKnowledgeKnowledgebaseDocumentVariationParams() *GetKnowledgeKnowledgebaseDocumentVariationParams {
	var ()
	return &GetKnowledgeKnowledgebaseDocumentVariationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentVariationParamsWithTimeout creates a new GetKnowledgeKnowledgebaseDocumentVariationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKnowledgeKnowledgebaseDocumentVariationParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseDocumentVariationParams {
	var ()
	return &GetKnowledgeKnowledgebaseDocumentVariationParams{

		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentVariationParamsWithContext creates a new GetKnowledgeKnowledgebaseDocumentVariationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKnowledgeKnowledgebaseDocumentVariationParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseDocumentVariationParams {
	var ()
	return &GetKnowledgeKnowledgebaseDocumentVariationParams{

		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentVariationParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseDocumentVariationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKnowledgeKnowledgebaseDocumentVariationParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseDocumentVariationParams {
	var ()
	return &GetKnowledgeKnowledgebaseDocumentVariationParams{
		HTTPClient: client,
	}
}

/*GetKnowledgeKnowledgebaseDocumentVariationParams contains all the parameters to send to the API endpoint
for the get knowledge knowledgebase document variation operation typically these are written to a http.Request
*/
type GetKnowledgeKnowledgebaseDocumentVariationParams struct {

	/*DocumentID
	  Globally unique identifier for a document.

	*/
	DocumentID string
	/*DocumentState
	  The state of the document.

	*/
	DocumentState *string
	/*DocumentVariationID
	  Globally unique identifier for a document variation.

	*/
	DocumentVariationID string
	/*KnowledgeBaseID
	  Globally unique identifier for a knowledge base.

	*/
	KnowledgeBaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) WithDocumentID(documentID string) *GetKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithDocumentState adds the documentState to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) WithDocumentState(documentState *string) *GetKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetDocumentState(documentState)
	return o
}

// SetDocumentState adds the documentState to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) SetDocumentState(documentState *string) {
	o.DocumentState = documentState
}

// WithDocumentVariationID adds the documentVariationID to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) WithDocumentVariationID(documentVariationID string) *GetKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetDocumentVariationID(documentVariationID)
	return o
}

// SetDocumentVariationID adds the documentVariationId to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) SetDocumentVariationID(documentVariationID string) {
	o.DocumentVariationID = documentVariationID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase document variation params
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseDocumentVariationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param documentId
	if err := r.SetPathParam("documentId", o.DocumentID); err != nil {
		return err
	}

	if o.DocumentState != nil {

		// query param documentState
		var qrDocumentState string
		if o.DocumentState != nil {
			qrDocumentState = *o.DocumentState
		}
		qDocumentState := qrDocumentState
		if qDocumentState != "" {
			if err := r.SetQueryParam("documentState", qDocumentState); err != nil {
				return err
			}
		}

	}

	// path param documentVariationId
	if err := r.SetPathParam("documentVariationId", o.DocumentVariationID); err != nil {
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
