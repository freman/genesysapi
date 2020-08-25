// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewGetDocumentationGknSearchParams creates a new GetDocumentationGknSearchParams object
// with the default values initialized.
func NewGetDocumentationGknSearchParams() *GetDocumentationGknSearchParams {
	var ()
	return &GetDocumentationGknSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDocumentationGknSearchParamsWithTimeout creates a new GetDocumentationGknSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDocumentationGknSearchParamsWithTimeout(timeout time.Duration) *GetDocumentationGknSearchParams {
	var ()
	return &GetDocumentationGknSearchParams{

		timeout: timeout,
	}
}

// NewGetDocumentationGknSearchParamsWithContext creates a new GetDocumentationGknSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDocumentationGknSearchParamsWithContext(ctx context.Context) *GetDocumentationGknSearchParams {
	var ()
	return &GetDocumentationGknSearchParams{

		Context: ctx,
	}
}

// NewGetDocumentationGknSearchParamsWithHTTPClient creates a new GetDocumentationGknSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDocumentationGknSearchParamsWithHTTPClient(client *http.Client) *GetDocumentationGknSearchParams {
	var ()
	return &GetDocumentationGknSearchParams{
		HTTPClient: client,
	}
}

/*GetDocumentationGknSearchParams contains all the parameters to send to the API endpoint
for the get documentation gkn search operation typically these are written to a http.Request
*/
type GetDocumentationGknSearchParams struct {

	/*Q64
	  q64

	*/
	Q64 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get documentation gkn search params
func (o *GetDocumentationGknSearchParams) WithTimeout(timeout time.Duration) *GetDocumentationGknSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get documentation gkn search params
func (o *GetDocumentationGknSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get documentation gkn search params
func (o *GetDocumentationGknSearchParams) WithContext(ctx context.Context) *GetDocumentationGknSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get documentation gkn search params
func (o *GetDocumentationGknSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get documentation gkn search params
func (o *GetDocumentationGknSearchParams) WithHTTPClient(client *http.Client) *GetDocumentationGknSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get documentation gkn search params
func (o *GetDocumentationGknSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQ64 adds the q64 to the get documentation gkn search params
func (o *GetDocumentationGknSearchParams) WithQ64(q64 string) *GetDocumentationGknSearchParams {
	o.SetQ64(q64)
	return o
}

// SetQ64 adds the q64 to the get documentation gkn search params
func (o *GetDocumentationGknSearchParams) SetQ64(q64 string) {
	o.Q64 = q64
}

// WriteToRequest writes these params to a swagger request
func (o *GetDocumentationGknSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param q64
	qrQ64 := o.Q64
	qQ64 := qrQ64
	if qQ64 != "" {
		if err := r.SetQueryParam("q64", qQ64); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}