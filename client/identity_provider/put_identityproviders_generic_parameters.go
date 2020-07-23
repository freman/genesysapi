// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

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

// NewPutIdentityprovidersGenericParams creates a new PutIdentityprovidersGenericParams object
// with the default values initialized.
func NewPutIdentityprovidersGenericParams() *PutIdentityprovidersGenericParams {
	var ()
	return &PutIdentityprovidersGenericParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutIdentityprovidersGenericParamsWithTimeout creates a new PutIdentityprovidersGenericParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutIdentityprovidersGenericParamsWithTimeout(timeout time.Duration) *PutIdentityprovidersGenericParams {
	var ()
	return &PutIdentityprovidersGenericParams{

		timeout: timeout,
	}
}

// NewPutIdentityprovidersGenericParamsWithContext creates a new PutIdentityprovidersGenericParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutIdentityprovidersGenericParamsWithContext(ctx context.Context) *PutIdentityprovidersGenericParams {
	var ()
	return &PutIdentityprovidersGenericParams{

		Context: ctx,
	}
}

// NewPutIdentityprovidersGenericParamsWithHTTPClient creates a new PutIdentityprovidersGenericParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutIdentityprovidersGenericParamsWithHTTPClient(client *http.Client) *PutIdentityprovidersGenericParams {
	var ()
	return &PutIdentityprovidersGenericParams{
		HTTPClient: client,
	}
}

/*PutIdentityprovidersGenericParams contains all the parameters to send to the API endpoint
for the put identityproviders generic operation typically these are written to a http.Request
*/
type PutIdentityprovidersGenericParams struct {

	/*Body
	  Provider

	*/
	Body *models.GenericSAML

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put identityproviders generic params
func (o *PutIdentityprovidersGenericParams) WithTimeout(timeout time.Duration) *PutIdentityprovidersGenericParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put identityproviders generic params
func (o *PutIdentityprovidersGenericParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put identityproviders generic params
func (o *PutIdentityprovidersGenericParams) WithContext(ctx context.Context) *PutIdentityprovidersGenericParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put identityproviders generic params
func (o *PutIdentityprovidersGenericParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put identityproviders generic params
func (o *PutIdentityprovidersGenericParams) WithHTTPClient(client *http.Client) *PutIdentityprovidersGenericParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put identityproviders generic params
func (o *PutIdentityprovidersGenericParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put identityproviders generic params
func (o *PutIdentityprovidersGenericParams) WithBody(body *models.GenericSAML) *PutIdentityprovidersGenericParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put identityproviders generic params
func (o *PutIdentityprovidersGenericParams) SetBody(body *models.GenericSAML) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutIdentityprovidersGenericParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
