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

// NewDeleteContentmanagementDocumentParams creates a new DeleteContentmanagementDocumentParams object
// with the default values initialized.
func NewDeleteContentmanagementDocumentParams() *DeleteContentmanagementDocumentParams {
	var ()
	return &DeleteContentmanagementDocumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContentmanagementDocumentParamsWithTimeout creates a new DeleteContentmanagementDocumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteContentmanagementDocumentParamsWithTimeout(timeout time.Duration) *DeleteContentmanagementDocumentParams {
	var ()
	return &DeleteContentmanagementDocumentParams{

		timeout: timeout,
	}
}

// NewDeleteContentmanagementDocumentParamsWithContext creates a new DeleteContentmanagementDocumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteContentmanagementDocumentParamsWithContext(ctx context.Context) *DeleteContentmanagementDocumentParams {
	var ()
	return &DeleteContentmanagementDocumentParams{

		Context: ctx,
	}
}

// NewDeleteContentmanagementDocumentParamsWithHTTPClient creates a new DeleteContentmanagementDocumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteContentmanagementDocumentParamsWithHTTPClient(client *http.Client) *DeleteContentmanagementDocumentParams {
	var ()
	return &DeleteContentmanagementDocumentParams{
		HTTPClient: client,
	}
}

/*DeleteContentmanagementDocumentParams contains all the parameters to send to the API endpoint
for the delete contentmanagement document operation typically these are written to a http.Request
*/
type DeleteContentmanagementDocumentParams struct {

	/*DocumentID
	  Document ID

	*/
	DocumentID string
	/*Override
	  Override any lock on the document

	*/
	Override *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete contentmanagement document params
func (o *DeleteContentmanagementDocumentParams) WithTimeout(timeout time.Duration) *DeleteContentmanagementDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete contentmanagement document params
func (o *DeleteContentmanagementDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete contentmanagement document params
func (o *DeleteContentmanagementDocumentParams) WithContext(ctx context.Context) *DeleteContentmanagementDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete contentmanagement document params
func (o *DeleteContentmanagementDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete contentmanagement document params
func (o *DeleteContentmanagementDocumentParams) WithHTTPClient(client *http.Client) *DeleteContentmanagementDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete contentmanagement document params
func (o *DeleteContentmanagementDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the delete contentmanagement document params
func (o *DeleteContentmanagementDocumentParams) WithDocumentID(documentID string) *DeleteContentmanagementDocumentParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the delete contentmanagement document params
func (o *DeleteContentmanagementDocumentParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WithOverride adds the override to the delete contentmanagement document params
func (o *DeleteContentmanagementDocumentParams) WithOverride(override *bool) *DeleteContentmanagementDocumentParams {
	o.SetOverride(override)
	return o
}

// SetOverride adds the override to the delete contentmanagement document params
func (o *DeleteContentmanagementDocumentParams) SetOverride(override *bool) {
	o.Override = override
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContentmanagementDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param documentId
	if err := r.SetPathParam("documentId", o.DocumentID); err != nil {
		return err
	}

	if o.Override != nil {

		// query param override
		var qrOverride bool
		if o.Override != nil {
			qrOverride = *o.Override
		}
		qOverride := swag.FormatBool(qrOverride)
		if qOverride != "" {
			if err := r.SetQueryParam("override", qOverride); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}