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

// NewPatchKnowledgeKnowledgebaseDocumentVariationParams creates a new PatchKnowledgeKnowledgebaseDocumentVariationParams object
// with the default values initialized.
func NewPatchKnowledgeKnowledgebaseDocumentVariationParams() *PatchKnowledgeKnowledgebaseDocumentVariationParams {
	var ()
	return &PatchKnowledgeKnowledgebaseDocumentVariationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchKnowledgeKnowledgebaseDocumentVariationParamsWithTimeout creates a new PatchKnowledgeKnowledgebaseDocumentVariationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchKnowledgeKnowledgebaseDocumentVariationParamsWithTimeout(timeout time.Duration) *PatchKnowledgeKnowledgebaseDocumentVariationParams {
	var ()
	return &PatchKnowledgeKnowledgebaseDocumentVariationParams{

		timeout: timeout,
	}
}

// NewPatchKnowledgeKnowledgebaseDocumentVariationParamsWithContext creates a new PatchKnowledgeKnowledgebaseDocumentVariationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchKnowledgeKnowledgebaseDocumentVariationParamsWithContext(ctx context.Context) *PatchKnowledgeKnowledgebaseDocumentVariationParams {
	var ()
	return &PatchKnowledgeKnowledgebaseDocumentVariationParams{

		Context: ctx,
	}
}

// NewPatchKnowledgeKnowledgebaseDocumentVariationParamsWithHTTPClient creates a new PatchKnowledgeKnowledgebaseDocumentVariationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchKnowledgeKnowledgebaseDocumentVariationParamsWithHTTPClient(client *http.Client) *PatchKnowledgeKnowledgebaseDocumentVariationParams {
	var ()
	return &PatchKnowledgeKnowledgebaseDocumentVariationParams{
		HTTPClient: client,
	}
}

/*PatchKnowledgeKnowledgebaseDocumentVariationParams contains all the parameters to send to the API endpoint
for the patch knowledge knowledgebase document variation operation typically these are written to a http.Request
*/
type PatchKnowledgeKnowledgebaseDocumentVariationParams struct {

	/*Body*/
	Body *models.DocumentVariation
	/*DocumentID
	  Globally unique identifier for a document.

	*/
	DocumentID string
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

// WithTimeout adds the timeout to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) WithTimeout(timeout time.Duration) *PatchKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) WithContext(ctx context.Context) *PatchKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) WithHTTPClient(client *http.Client) *PatchKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) WithBody(body *models.DocumentVariation) *PatchKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) SetBody(body *models.DocumentVariation) {
	o.Body = body
}

// WithDocumentID adds the documentID to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) WithDocumentID(documentID string) *PatchKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithDocumentVariationID adds the documentVariationID to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) WithDocumentVariationID(documentVariationID string) *PatchKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetDocumentVariationID(documentVariationID)
	return o
}

// SetDocumentVariationID adds the documentVariationId to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) SetDocumentVariationID(documentVariationID string) {
	o.DocumentVariationID = documentVariationID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) WithKnowledgeBaseID(knowledgeBaseID string) *PatchKnowledgeKnowledgebaseDocumentVariationParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the patch knowledge knowledgebase document variation params
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchKnowledgeKnowledgebaseDocumentVariationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param documentId
	if err := r.SetPathParam("documentId", o.DocumentID); err != nil {
		return err
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
