// Code generated by go-swagger; DO NOT EDIT.

package fax

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

// NewGetFaxDocumentContentParams creates a new GetFaxDocumentContentParams object
// with the default values initialized.
func NewGetFaxDocumentContentParams() *GetFaxDocumentContentParams {
	var ()
	return &GetFaxDocumentContentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFaxDocumentContentParamsWithTimeout creates a new GetFaxDocumentContentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFaxDocumentContentParamsWithTimeout(timeout time.Duration) *GetFaxDocumentContentParams {
	var ()
	return &GetFaxDocumentContentParams{

		timeout: timeout,
	}
}

// NewGetFaxDocumentContentParamsWithContext creates a new GetFaxDocumentContentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFaxDocumentContentParamsWithContext(ctx context.Context) *GetFaxDocumentContentParams {
	var ()
	return &GetFaxDocumentContentParams{

		Context: ctx,
	}
}

// NewGetFaxDocumentContentParamsWithHTTPClient creates a new GetFaxDocumentContentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFaxDocumentContentParamsWithHTTPClient(client *http.Client) *GetFaxDocumentContentParams {
	var ()
	return &GetFaxDocumentContentParams{
		HTTPClient: client,
	}
}

/*GetFaxDocumentContentParams contains all the parameters to send to the API endpoint
for the get fax document content operation typically these are written to a http.Request
*/
type GetFaxDocumentContentParams struct {

	/*DocumentID
	  Document ID

	*/
	DocumentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fax document content params
func (o *GetFaxDocumentContentParams) WithTimeout(timeout time.Duration) *GetFaxDocumentContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fax document content params
func (o *GetFaxDocumentContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fax document content params
func (o *GetFaxDocumentContentParams) WithContext(ctx context.Context) *GetFaxDocumentContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fax document content params
func (o *GetFaxDocumentContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fax document content params
func (o *GetFaxDocumentContentParams) WithHTTPClient(client *http.Client) *GetFaxDocumentContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fax document content params
func (o *GetFaxDocumentContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the get fax document content params
func (o *GetFaxDocumentContentParams) WithDocumentID(documentID string) *GetFaxDocumentContentParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get fax document content params
func (o *GetFaxDocumentContentParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFaxDocumentContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param documentId
	if err := r.SetPathParam("documentId", o.DocumentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
