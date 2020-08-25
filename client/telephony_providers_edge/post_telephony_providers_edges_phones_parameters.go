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

// NewPostTelephonyProvidersEdgesPhonesParams creates a new PostTelephonyProvidersEdgesPhonesParams object
// with the default values initialized.
func NewPostTelephonyProvidersEdgesPhonesParams() *PostTelephonyProvidersEdgesPhonesParams {
	var ()
	return &PostTelephonyProvidersEdgesPhonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTelephonyProvidersEdgesPhonesParamsWithTimeout creates a new PostTelephonyProvidersEdgesPhonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTelephonyProvidersEdgesPhonesParamsWithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgesPhonesParams {
	var ()
	return &PostTelephonyProvidersEdgesPhonesParams{

		timeout: timeout,
	}
}

// NewPostTelephonyProvidersEdgesPhonesParamsWithContext creates a new PostTelephonyProvidersEdgesPhonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTelephonyProvidersEdgesPhonesParamsWithContext(ctx context.Context) *PostTelephonyProvidersEdgesPhonesParams {
	var ()
	return &PostTelephonyProvidersEdgesPhonesParams{

		Context: ctx,
	}
}

// NewPostTelephonyProvidersEdgesPhonesParamsWithHTTPClient creates a new PostTelephonyProvidersEdgesPhonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTelephonyProvidersEdgesPhonesParamsWithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgesPhonesParams {
	var ()
	return &PostTelephonyProvidersEdgesPhonesParams{
		HTTPClient: client,
	}
}

/*PostTelephonyProvidersEdgesPhonesParams contains all the parameters to send to the API endpoint
for the post telephony providers edges phones operation typically these are written to a http.Request
*/
type PostTelephonyProvidersEdgesPhonesParams struct {

	/*Body
	  Phone

	*/
	Body *models.Phone

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post telephony providers edges phones params
func (o *PostTelephonyProvidersEdgesPhonesParams) WithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgesPhonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post telephony providers edges phones params
func (o *PostTelephonyProvidersEdgesPhonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post telephony providers edges phones params
func (o *PostTelephonyProvidersEdgesPhonesParams) WithContext(ctx context.Context) *PostTelephonyProvidersEdgesPhonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post telephony providers edges phones params
func (o *PostTelephonyProvidersEdgesPhonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post telephony providers edges phones params
func (o *PostTelephonyProvidersEdgesPhonesParams) WithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgesPhonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post telephony providers edges phones params
func (o *PostTelephonyProvidersEdgesPhonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post telephony providers edges phones params
func (o *PostTelephonyProvidersEdgesPhonesParams) WithBody(body *models.Phone) *PostTelephonyProvidersEdgesPhonesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post telephony providers edges phones params
func (o *PostTelephonyProvidersEdgesPhonesParams) SetBody(body *models.Phone) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTelephonyProvidersEdgesPhonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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