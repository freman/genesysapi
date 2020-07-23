// Code generated by go-swagger; DO NOT EDIT.

package license

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

// NewPostLicenseInferParams creates a new PostLicenseInferParams object
// with the default values initialized.
func NewPostLicenseInferParams() *PostLicenseInferParams {
	var ()
	return &PostLicenseInferParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLicenseInferParamsWithTimeout creates a new PostLicenseInferParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLicenseInferParamsWithTimeout(timeout time.Duration) *PostLicenseInferParams {
	var ()
	return &PostLicenseInferParams{

		timeout: timeout,
	}
}

// NewPostLicenseInferParamsWithContext creates a new PostLicenseInferParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLicenseInferParamsWithContext(ctx context.Context) *PostLicenseInferParams {
	var ()
	return &PostLicenseInferParams{

		Context: ctx,
	}
}

// NewPostLicenseInferParamsWithHTTPClient creates a new PostLicenseInferParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLicenseInferParamsWithHTTPClient(client *http.Client) *PostLicenseInferParams {
	var ()
	return &PostLicenseInferParams{
		HTTPClient: client,
	}
}

/*PostLicenseInferParams contains all the parameters to send to the API endpoint
for the post license infer operation typically these are written to a http.Request
*/
type PostLicenseInferParams struct {

	/*Body
	  The roleIds to use while inferring licenses

	*/
	Body []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post license infer params
func (o *PostLicenseInferParams) WithTimeout(timeout time.Duration) *PostLicenseInferParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post license infer params
func (o *PostLicenseInferParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post license infer params
func (o *PostLicenseInferParams) WithContext(ctx context.Context) *PostLicenseInferParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post license infer params
func (o *PostLicenseInferParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post license infer params
func (o *PostLicenseInferParams) WithHTTPClient(client *http.Client) *PostLicenseInferParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post license infer params
func (o *PostLicenseInferParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post license infer params
func (o *PostLicenseInferParams) WithBody(body []string) *PostLicenseInferParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post license infer params
func (o *PostLicenseInferParams) SetBody(body []string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostLicenseInferParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
