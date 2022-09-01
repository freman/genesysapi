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

// NewDeleteKnowledgeKnowledgebaseDocumentParams creates a new DeleteKnowledgeKnowledgebaseDocumentParams object
// with the default values initialized.
func NewDeleteKnowledgeKnowledgebaseDocumentParams() *DeleteKnowledgeKnowledgebaseDocumentParams {
	var ()
	return &DeleteKnowledgeKnowledgebaseDocumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKnowledgeKnowledgebaseDocumentParamsWithTimeout creates a new DeleteKnowledgeKnowledgebaseDocumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteKnowledgeKnowledgebaseDocumentParamsWithTimeout(timeout time.Duration) *DeleteKnowledgeKnowledgebaseDocumentParams {
	var ()
	return &DeleteKnowledgeKnowledgebaseDocumentParams{

		timeout: timeout,
	}
}

// NewDeleteKnowledgeKnowledgebaseDocumentParamsWithContext creates a new DeleteKnowledgeKnowledgebaseDocumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteKnowledgeKnowledgebaseDocumentParamsWithContext(ctx context.Context) *DeleteKnowledgeKnowledgebaseDocumentParams {
	var ()
	return &DeleteKnowledgeKnowledgebaseDocumentParams{

		Context: ctx,
	}
}

// NewDeleteKnowledgeKnowledgebaseDocumentParamsWithHTTPClient creates a new DeleteKnowledgeKnowledgebaseDocumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteKnowledgeKnowledgebaseDocumentParamsWithHTTPClient(client *http.Client) *DeleteKnowledgeKnowledgebaseDocumentParams {
	var ()
	return &DeleteKnowledgeKnowledgebaseDocumentParams{
		HTTPClient: client,
	}
}

/*DeleteKnowledgeKnowledgebaseDocumentParams contains all the parameters to send to the API endpoint
for the delete knowledge knowledgebase document operation typically these are written to a http.Request
*/
type DeleteKnowledgeKnowledgebaseDocumentParams struct {

	/*DocumentID
	  Document ID.

	*/
	DocumentID string
	/*KnowledgeBaseID
	  Knowledge base ID.

	*/
	KnowledgeBaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete knowledge knowledgebase document params
func (o *DeleteKnowledgeKnowledgebaseDocumentParams) WithTimeout(timeout time.Duration) *DeleteKnowledgeKnowledgebaseDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete knowledge knowledgebase document params
func (o *DeleteKnowledgeKnowledgebaseDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete knowledge knowledgebase document params
func (o *DeleteKnowledgeKnowledgebaseDocumentParams) WithContext(ctx context.Context) *DeleteKnowledgeKnowledgebaseDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete knowledge knowledgebase document params
func (o *DeleteKnowledgeKnowledgebaseDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete knowledge knowledgebase document params
func (o *DeleteKnowledgeKnowledgebaseDocumentParams) WithHTTPClient(client *http.Client) *DeleteKnowledgeKnowledgebaseDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete knowledge knowledgebase document params
func (o *DeleteKnowledgeKnowledgebaseDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the delete knowledge knowledgebase document params
func (o *DeleteKnowledgeKnowledgebaseDocumentParams) WithDocumentID(documentID string) *DeleteKnowledgeKnowledgebaseDocumentParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the delete knowledge knowledgebase document params
func (o *DeleteKnowledgeKnowledgebaseDocumentParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the delete knowledge knowledgebase document params
func (o *DeleteKnowledgeKnowledgebaseDocumentParams) WithKnowledgeBaseID(knowledgeBaseID string) *DeleteKnowledgeKnowledgebaseDocumentParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the delete knowledge knowledgebase document params
func (o *DeleteKnowledgeKnowledgebaseDocumentParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKnowledgeKnowledgebaseDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}