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

// NewDeleteFaxDocumentParams creates a new DeleteFaxDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteFaxDocumentParams() *DeleteFaxDocumentParams {
	return &DeleteFaxDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFaxDocumentParamsWithTimeout creates a new DeleteFaxDocumentParams object
// with the ability to set a timeout on a request.
func NewDeleteFaxDocumentParamsWithTimeout(timeout time.Duration) *DeleteFaxDocumentParams {
	return &DeleteFaxDocumentParams{
		timeout: timeout,
	}
}

// NewDeleteFaxDocumentParamsWithContext creates a new DeleteFaxDocumentParams object
// with the ability to set a context for a request.
func NewDeleteFaxDocumentParamsWithContext(ctx context.Context) *DeleteFaxDocumentParams {
	return &DeleteFaxDocumentParams{
		Context: ctx,
	}
}

// NewDeleteFaxDocumentParamsWithHTTPClient creates a new DeleteFaxDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteFaxDocumentParamsWithHTTPClient(client *http.Client) *DeleteFaxDocumentParams {
	return &DeleteFaxDocumentParams{
		HTTPClient: client,
	}
}

/*
DeleteFaxDocumentParams contains all the parameters to send to the API endpoint

	for the delete fax document operation.

	Typically these are written to a http.Request.
*/
type DeleteFaxDocumentParams struct {

	/* DocumentID.

	   Document ID
	*/
	DocumentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete fax document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFaxDocumentParams) WithDefaults() *DeleteFaxDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete fax document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFaxDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete fax document params
func (o *DeleteFaxDocumentParams) WithTimeout(timeout time.Duration) *DeleteFaxDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete fax document params
func (o *DeleteFaxDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete fax document params
func (o *DeleteFaxDocumentParams) WithContext(ctx context.Context) *DeleteFaxDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete fax document params
func (o *DeleteFaxDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete fax document params
func (o *DeleteFaxDocumentParams) WithHTTPClient(client *http.Client) *DeleteFaxDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete fax document params
func (o *DeleteFaxDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the delete fax document params
func (o *DeleteFaxDocumentParams) WithDocumentID(documentID string) *DeleteFaxDocumentParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the delete fax document params
func (o *DeleteFaxDocumentParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFaxDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
