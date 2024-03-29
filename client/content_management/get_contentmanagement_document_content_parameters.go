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
)

// NewGetContentmanagementDocumentContentParams creates a new GetContentmanagementDocumentContentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContentmanagementDocumentContentParams() *GetContentmanagementDocumentContentParams {
	return &GetContentmanagementDocumentContentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentmanagementDocumentContentParamsWithTimeout creates a new GetContentmanagementDocumentContentParams object
// with the ability to set a timeout on a request.
func NewGetContentmanagementDocumentContentParamsWithTimeout(timeout time.Duration) *GetContentmanagementDocumentContentParams {
	return &GetContentmanagementDocumentContentParams{
		timeout: timeout,
	}
}

// NewGetContentmanagementDocumentContentParamsWithContext creates a new GetContentmanagementDocumentContentParams object
// with the ability to set a context for a request.
func NewGetContentmanagementDocumentContentParamsWithContext(ctx context.Context) *GetContentmanagementDocumentContentParams {
	return &GetContentmanagementDocumentContentParams{
		Context: ctx,
	}
}

// NewGetContentmanagementDocumentContentParamsWithHTTPClient creates a new GetContentmanagementDocumentContentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContentmanagementDocumentContentParamsWithHTTPClient(client *http.Client) *GetContentmanagementDocumentContentParams {
	return &GetContentmanagementDocumentContentParams{
		HTTPClient: client,
	}
}

/*
GetContentmanagementDocumentContentParams contains all the parameters to send to the API endpoint

	for the get contentmanagement document content operation.

	Typically these are written to a http.Request.
*/
type GetContentmanagementDocumentContentParams struct {

	/* ContentType.

	   The requested format for the specified document. If supported, the document will be returned in that format. Example contentType=audio/wav
	*/
	ContentType *string

	/* Disposition.

	   Request how the content will be downloaded: a file attachment or inline. Default is attachment.
	*/
	Disposition *string

	/* DocumentID.

	   Document ID
	*/
	DocumentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contentmanagement document content params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentmanagementDocumentContentParams) WithDefaults() *GetContentmanagementDocumentContentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contentmanagement document content params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentmanagementDocumentContentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get contentmanagement document content params
func (o *GetContentmanagementDocumentContentParams) WithTimeout(timeout time.Duration) *GetContentmanagementDocumentContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contentmanagement document content params
func (o *GetContentmanagementDocumentContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contentmanagement document content params
func (o *GetContentmanagementDocumentContentParams) WithContext(ctx context.Context) *GetContentmanagementDocumentContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contentmanagement document content params
func (o *GetContentmanagementDocumentContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contentmanagement document content params
func (o *GetContentmanagementDocumentContentParams) WithHTTPClient(client *http.Client) *GetContentmanagementDocumentContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contentmanagement document content params
func (o *GetContentmanagementDocumentContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the get contentmanagement document content params
func (o *GetContentmanagementDocumentContentParams) WithContentType(contentType *string) *GetContentmanagementDocumentContentParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the get contentmanagement document content params
func (o *GetContentmanagementDocumentContentParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithDisposition adds the disposition to the get contentmanagement document content params
func (o *GetContentmanagementDocumentContentParams) WithDisposition(disposition *string) *GetContentmanagementDocumentContentParams {
	o.SetDisposition(disposition)
	return o
}

// SetDisposition adds the disposition to the get contentmanagement document content params
func (o *GetContentmanagementDocumentContentParams) SetDisposition(disposition *string) {
	o.Disposition = disposition
}

// WithDocumentID adds the documentID to the get contentmanagement document content params
func (o *GetContentmanagementDocumentContentParams) WithDocumentID(documentID string) *GetContentmanagementDocumentContentParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get contentmanagement document content params
func (o *GetContentmanagementDocumentContentParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentmanagementDocumentContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// query param contentType
		var qrContentType string

		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {

			if err := r.SetQueryParam("contentType", qContentType); err != nil {
				return err
			}
		}
	}

	if o.Disposition != nil {

		// query param disposition
		var qrDisposition string

		if o.Disposition != nil {
			qrDisposition = *o.Disposition
		}
		qDisposition := qrDisposition
		if qDisposition != "" {

			if err := r.SetQueryParam("disposition", qDisposition); err != nil {
				return err
			}
		}
	}

	// path param documentId
	if err := r.SetPathParam("documentId", o.DocumentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
