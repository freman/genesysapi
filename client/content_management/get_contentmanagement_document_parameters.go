// Code generated by go-swagger; DO NOT EDIT.

package content_management

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

// NewGetContentmanagementDocumentParams creates a new GetContentmanagementDocumentParams object
// with the default values initialized.
func NewGetContentmanagementDocumentParams() *GetContentmanagementDocumentParams {
	var ()
	return &GetContentmanagementDocumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentmanagementDocumentParamsWithTimeout creates a new GetContentmanagementDocumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContentmanagementDocumentParamsWithTimeout(timeout time.Duration) *GetContentmanagementDocumentParams {
	var ()
	return &GetContentmanagementDocumentParams{

		timeout: timeout,
	}
}

// NewGetContentmanagementDocumentParamsWithContext creates a new GetContentmanagementDocumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContentmanagementDocumentParamsWithContext(ctx context.Context) *GetContentmanagementDocumentParams {
	var ()
	return &GetContentmanagementDocumentParams{

		Context: ctx,
	}
}

// NewGetContentmanagementDocumentParamsWithHTTPClient creates a new GetContentmanagementDocumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContentmanagementDocumentParamsWithHTTPClient(client *http.Client) *GetContentmanagementDocumentParams {
	var ()
	return &GetContentmanagementDocumentParams{
		HTTPClient: client,
	}
}

/*GetContentmanagementDocumentParams contains all the parameters to send to the API endpoint
for the get contentmanagement document operation typically these are written to a http.Request
*/
type GetContentmanagementDocumentParams struct {

	/*DocumentID
	  Document ID

	*/
	DocumentID string
	/*Expand
	  Which fields, if any, to expand.

	*/
	Expand []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get contentmanagement document params
func (o *GetContentmanagementDocumentParams) WithTimeout(timeout time.Duration) *GetContentmanagementDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contentmanagement document params
func (o *GetContentmanagementDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contentmanagement document params
func (o *GetContentmanagementDocumentParams) WithContext(ctx context.Context) *GetContentmanagementDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contentmanagement document params
func (o *GetContentmanagementDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contentmanagement document params
func (o *GetContentmanagementDocumentParams) WithHTTPClient(client *http.Client) *GetContentmanagementDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contentmanagement document params
func (o *GetContentmanagementDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the get contentmanagement document params
func (o *GetContentmanagementDocumentParams) WithDocumentID(documentID string) *GetContentmanagementDocumentParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get contentmanagement document params
func (o *GetContentmanagementDocumentParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithExpand adds the expand to the get contentmanagement document params
func (o *GetContentmanagementDocumentParams) WithExpand(expand []string) *GetContentmanagementDocumentParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get contentmanagement document params
func (o *GetContentmanagementDocumentParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentmanagementDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param documentId
	if err := r.SetPathParam("documentId", o.DocumentID); err != nil {
		return err
	}

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
