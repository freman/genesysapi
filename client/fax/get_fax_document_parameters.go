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

// NewGetFaxDocumentParams creates a new GetFaxDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFaxDocumentParams() *GetFaxDocumentParams {
	return &GetFaxDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFaxDocumentParamsWithTimeout creates a new GetFaxDocumentParams object
// with the ability to set a timeout on a request.
func NewGetFaxDocumentParamsWithTimeout(timeout time.Duration) *GetFaxDocumentParams {
	return &GetFaxDocumentParams{
		timeout: timeout,
	}
}

// NewGetFaxDocumentParamsWithContext creates a new GetFaxDocumentParams object
// with the ability to set a context for a request.
func NewGetFaxDocumentParamsWithContext(ctx context.Context) *GetFaxDocumentParams {
	return &GetFaxDocumentParams{
		Context: ctx,
	}
}

// NewGetFaxDocumentParamsWithHTTPClient creates a new GetFaxDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFaxDocumentParamsWithHTTPClient(client *http.Client) *GetFaxDocumentParams {
	return &GetFaxDocumentParams{
		HTTPClient: client,
	}
}

/*
GetFaxDocumentParams contains all the parameters to send to the API endpoint

	for the get fax document operation.

	Typically these are written to a http.Request.
*/
type GetFaxDocumentParams struct {

	/* DocumentID.

	   Document ID
	*/
	DocumentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get fax document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFaxDocumentParams) WithDefaults() *GetFaxDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get fax document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFaxDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get fax document params
func (o *GetFaxDocumentParams) WithTimeout(timeout time.Duration) *GetFaxDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fax document params
func (o *GetFaxDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fax document params
func (o *GetFaxDocumentParams) WithContext(ctx context.Context) *GetFaxDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fax document params
func (o *GetFaxDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fax document params
func (o *GetFaxDocumentParams) WithHTTPClient(client *http.Client) *GetFaxDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fax document params
func (o *GetFaxDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the get fax document params
func (o *GetFaxDocumentParams) WithDocumentID(documentID string) *GetFaxDocumentParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get fax document params
func (o *GetFaxDocumentParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFaxDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
