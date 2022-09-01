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

// NewPatchKnowledgeKnowledgebaseDocumentParams creates a new PatchKnowledgeKnowledgebaseDocumentParams object
// with the default values initialized.
func NewPatchKnowledgeKnowledgebaseDocumentParams() *PatchKnowledgeKnowledgebaseDocumentParams {
	var ()
	return &PatchKnowledgeKnowledgebaseDocumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchKnowledgeKnowledgebaseDocumentParamsWithTimeout creates a new PatchKnowledgeKnowledgebaseDocumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchKnowledgeKnowledgebaseDocumentParamsWithTimeout(timeout time.Duration) *PatchKnowledgeKnowledgebaseDocumentParams {
	var ()
	return &PatchKnowledgeKnowledgebaseDocumentParams{

		timeout: timeout,
	}
}

// NewPatchKnowledgeKnowledgebaseDocumentParamsWithContext creates a new PatchKnowledgeKnowledgebaseDocumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchKnowledgeKnowledgebaseDocumentParamsWithContext(ctx context.Context) *PatchKnowledgeKnowledgebaseDocumentParams {
	var ()
	return &PatchKnowledgeKnowledgebaseDocumentParams{

		Context: ctx,
	}
}

// NewPatchKnowledgeKnowledgebaseDocumentParamsWithHTTPClient creates a new PatchKnowledgeKnowledgebaseDocumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchKnowledgeKnowledgebaseDocumentParamsWithHTTPClient(client *http.Client) *PatchKnowledgeKnowledgebaseDocumentParams {
	var ()
	return &PatchKnowledgeKnowledgebaseDocumentParams{
		HTTPClient: client,
	}
}

/*PatchKnowledgeKnowledgebaseDocumentParams contains all the parameters to send to the API endpoint
for the patch knowledge knowledgebase document operation typically these are written to a http.Request
*/
type PatchKnowledgeKnowledgebaseDocumentParams struct {

	/*Body*/
	Body *models.KnowledgeDocumentReq
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

// WithTimeout adds the timeout to the patch knowledge knowledgebase document params
func (o *PatchKnowledgeKnowledgebaseDocumentParams) WithTimeout(timeout time.Duration) *PatchKnowledgeKnowledgebaseDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch knowledge knowledgebase document params
func (o *PatchKnowledgeKnowledgebaseDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch knowledge knowledgebase document params
func (o *PatchKnowledgeKnowledgebaseDocumentParams) WithContext(ctx context.Context) *PatchKnowledgeKnowledgebaseDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch knowledge knowledgebase document params
func (o *PatchKnowledgeKnowledgebaseDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch knowledge knowledgebase document params
func (o *PatchKnowledgeKnowledgebaseDocumentParams) WithHTTPClient(client *http.Client) *PatchKnowledgeKnowledgebaseDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch knowledge knowledgebase document params
func (o *PatchKnowledgeKnowledgebaseDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch knowledge knowledgebase document params
func (o *PatchKnowledgeKnowledgebaseDocumentParams) WithBody(body *models.KnowledgeDocumentReq) *PatchKnowledgeKnowledgebaseDocumentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch knowledge knowledgebase document params
func (o *PatchKnowledgeKnowledgebaseDocumentParams) SetBody(body *models.KnowledgeDocumentReq) {
	o.Body = body
}

// WithDocumentID adds the documentID to the patch knowledge knowledgebase document params
func (o *PatchKnowledgeKnowledgebaseDocumentParams) WithDocumentID(documentID string) *PatchKnowledgeKnowledgebaseDocumentParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the patch knowledge knowledgebase document params
func (o *PatchKnowledgeKnowledgebaseDocumentParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the patch knowledge knowledgebase document params
func (o *PatchKnowledgeKnowledgebaseDocumentParams) WithKnowledgeBaseID(knowledgeBaseID string) *PatchKnowledgeKnowledgebaseDocumentParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the patch knowledge knowledgebase document params
func (o *PatchKnowledgeKnowledgebaseDocumentParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchKnowledgeKnowledgebaseDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}