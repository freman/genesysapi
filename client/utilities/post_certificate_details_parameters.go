// Code generated by go-swagger; DO NOT EDIT.

package utilities

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

// NewPostCertificateDetailsParams creates a new PostCertificateDetailsParams object
// with the default values initialized.
func NewPostCertificateDetailsParams() *PostCertificateDetailsParams {
	var ()
	return &PostCertificateDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCertificateDetailsParamsWithTimeout creates a new PostCertificateDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCertificateDetailsParamsWithTimeout(timeout time.Duration) *PostCertificateDetailsParams {
	var ()
	return &PostCertificateDetailsParams{

		timeout: timeout,
	}
}

// NewPostCertificateDetailsParamsWithContext creates a new PostCertificateDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCertificateDetailsParamsWithContext(ctx context.Context) *PostCertificateDetailsParams {
	var ()
	return &PostCertificateDetailsParams{

		Context: ctx,
	}
}

// NewPostCertificateDetailsParamsWithHTTPClient creates a new PostCertificateDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCertificateDetailsParamsWithHTTPClient(client *http.Client) *PostCertificateDetailsParams {
	var ()
	return &PostCertificateDetailsParams{
		HTTPClient: client,
	}
}

/*PostCertificateDetailsParams contains all the parameters to send to the API endpoint
for the post certificate details operation typically these are written to a http.Request
*/
type PostCertificateDetailsParams struct {

	/*Body
	  Certificate

	*/
	Body *models.Certificate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post certificate details params
func (o *PostCertificateDetailsParams) WithTimeout(timeout time.Duration) *PostCertificateDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post certificate details params
func (o *PostCertificateDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post certificate details params
func (o *PostCertificateDetailsParams) WithContext(ctx context.Context) *PostCertificateDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post certificate details params
func (o *PostCertificateDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post certificate details params
func (o *PostCertificateDetailsParams) WithHTTPClient(client *http.Client) *PostCertificateDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post certificate details params
func (o *PostCertificateDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post certificate details params
func (o *PostCertificateDetailsParams) WithBody(body *models.Certificate) *PostCertificateDetailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post certificate details params
func (o *PostCertificateDetailsParams) SetBody(body *models.Certificate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostCertificateDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
