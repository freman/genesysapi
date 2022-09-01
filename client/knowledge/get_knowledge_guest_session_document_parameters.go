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

// NewGetKnowledgeGuestSessionDocumentParams creates a new GetKnowledgeGuestSessionDocumentParams object
// with the default values initialized.
func NewGetKnowledgeGuestSessionDocumentParams() *GetKnowledgeGuestSessionDocumentParams {
	var ()
	return &GetKnowledgeGuestSessionDocumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeGuestSessionDocumentParamsWithTimeout creates a new GetKnowledgeGuestSessionDocumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKnowledgeGuestSessionDocumentParamsWithTimeout(timeout time.Duration) *GetKnowledgeGuestSessionDocumentParams {
	var ()
	return &GetKnowledgeGuestSessionDocumentParams{

		timeout: timeout,
	}
}

// NewGetKnowledgeGuestSessionDocumentParamsWithContext creates a new GetKnowledgeGuestSessionDocumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKnowledgeGuestSessionDocumentParamsWithContext(ctx context.Context) *GetKnowledgeGuestSessionDocumentParams {
	var ()
	return &GetKnowledgeGuestSessionDocumentParams{

		Context: ctx,
	}
}

// NewGetKnowledgeGuestSessionDocumentParamsWithHTTPClient creates a new GetKnowledgeGuestSessionDocumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKnowledgeGuestSessionDocumentParamsWithHTTPClient(client *http.Client) *GetKnowledgeGuestSessionDocumentParams {
	var ()
	return &GetKnowledgeGuestSessionDocumentParams{
		HTTPClient: client,
	}
}

/*GetKnowledgeGuestSessionDocumentParams contains all the parameters to send to the API endpoint
for the get knowledge guest session document operation typically these are written to a http.Request
*/
type GetKnowledgeGuestSessionDocumentParams struct {

	/*DocumentID
	  Document ID

	*/
	DocumentID string
	/*SessionID
	  Knowledge guest session ID.

	*/
	SessionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get knowledge guest session document params
func (o *GetKnowledgeGuestSessionDocumentParams) WithTimeout(timeout time.Duration) *GetKnowledgeGuestSessionDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge guest session document params
func (o *GetKnowledgeGuestSessionDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge guest session document params
func (o *GetKnowledgeGuestSessionDocumentParams) WithContext(ctx context.Context) *GetKnowledgeGuestSessionDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge guest session document params
func (o *GetKnowledgeGuestSessionDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge guest session document params
func (o *GetKnowledgeGuestSessionDocumentParams) WithHTTPClient(client *http.Client) *GetKnowledgeGuestSessionDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge guest session document params
func (o *GetKnowledgeGuestSessionDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the get knowledge guest session document params
func (o *GetKnowledgeGuestSessionDocumentParams) WithDocumentID(documentID string) *GetKnowledgeGuestSessionDocumentParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get knowledge guest session document params
func (o *GetKnowledgeGuestSessionDocumentParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithSessionID adds the sessionID to the get knowledge guest session document params
func (o *GetKnowledgeGuestSessionDocumentParams) WithSessionID(sessionID string) *GetKnowledgeGuestSessionDocumentParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the get knowledge guest session document params
func (o *GetKnowledgeGuestSessionDocumentParams) SetSessionID(sessionID string) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeGuestSessionDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param documentId
	if err := r.SetPathParam("documentId", o.DocumentID); err != nil {
		return err
	}

	// path param sessionId
	if err := r.SetPathParam("sessionId", o.SessionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
