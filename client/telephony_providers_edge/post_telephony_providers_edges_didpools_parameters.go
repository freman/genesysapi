// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

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

// NewPostTelephonyProvidersEdgesDidpoolsParams creates a new PostTelephonyProvidersEdgesDidpoolsParams object
// with the default values initialized.
func NewPostTelephonyProvidersEdgesDidpoolsParams() *PostTelephonyProvidersEdgesDidpoolsParams {
	var ()
	return &PostTelephonyProvidersEdgesDidpoolsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTelephonyProvidersEdgesDidpoolsParamsWithTimeout creates a new PostTelephonyProvidersEdgesDidpoolsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTelephonyProvidersEdgesDidpoolsParamsWithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgesDidpoolsParams {
	var ()
	return &PostTelephonyProvidersEdgesDidpoolsParams{

		timeout: timeout,
	}
}

// NewPostTelephonyProvidersEdgesDidpoolsParamsWithContext creates a new PostTelephonyProvidersEdgesDidpoolsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTelephonyProvidersEdgesDidpoolsParamsWithContext(ctx context.Context) *PostTelephonyProvidersEdgesDidpoolsParams {
	var ()
	return &PostTelephonyProvidersEdgesDidpoolsParams{

		Context: ctx,
	}
}

// NewPostTelephonyProvidersEdgesDidpoolsParamsWithHTTPClient creates a new PostTelephonyProvidersEdgesDidpoolsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTelephonyProvidersEdgesDidpoolsParamsWithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgesDidpoolsParams {
	var ()
	return &PostTelephonyProvidersEdgesDidpoolsParams{
		HTTPClient: client,
	}
}

/*PostTelephonyProvidersEdgesDidpoolsParams contains all the parameters to send to the API endpoint
for the post telephony providers edges didpools operation typically these are written to a http.Request
*/
type PostTelephonyProvidersEdgesDidpoolsParams struct {

	/*Body
	  DID pool

	*/
	Body *models.DIDPool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post telephony providers edges didpools params
func (o *PostTelephonyProvidersEdgesDidpoolsParams) WithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgesDidpoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post telephony providers edges didpools params
func (o *PostTelephonyProvidersEdgesDidpoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post telephony providers edges didpools params
func (o *PostTelephonyProvidersEdgesDidpoolsParams) WithContext(ctx context.Context) *PostTelephonyProvidersEdgesDidpoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post telephony providers edges didpools params
func (o *PostTelephonyProvidersEdgesDidpoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post telephony providers edges didpools params
func (o *PostTelephonyProvidersEdgesDidpoolsParams) WithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgesDidpoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post telephony providers edges didpools params
func (o *PostTelephonyProvidersEdgesDidpoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post telephony providers edges didpools params
func (o *PostTelephonyProvidersEdgesDidpoolsParams) WithBody(body *models.DIDPool) *PostTelephonyProvidersEdgesDidpoolsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post telephony providers edges didpools params
func (o *PostTelephonyProvidersEdgesDidpoolsParams) SetBody(body *models.DIDPool) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTelephonyProvidersEdgesDidpoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
