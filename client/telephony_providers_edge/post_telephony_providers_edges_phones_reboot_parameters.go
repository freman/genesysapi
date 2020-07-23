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

// NewPostTelephonyProvidersEdgesPhonesRebootParams creates a new PostTelephonyProvidersEdgesPhonesRebootParams object
// with the default values initialized.
func NewPostTelephonyProvidersEdgesPhonesRebootParams() *PostTelephonyProvidersEdgesPhonesRebootParams {
	var ()
	return &PostTelephonyProvidersEdgesPhonesRebootParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTelephonyProvidersEdgesPhonesRebootParamsWithTimeout creates a new PostTelephonyProvidersEdgesPhonesRebootParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTelephonyProvidersEdgesPhonesRebootParamsWithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgesPhonesRebootParams {
	var ()
	return &PostTelephonyProvidersEdgesPhonesRebootParams{

		timeout: timeout,
	}
}

// NewPostTelephonyProvidersEdgesPhonesRebootParamsWithContext creates a new PostTelephonyProvidersEdgesPhonesRebootParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTelephonyProvidersEdgesPhonesRebootParamsWithContext(ctx context.Context) *PostTelephonyProvidersEdgesPhonesRebootParams {
	var ()
	return &PostTelephonyProvidersEdgesPhonesRebootParams{

		Context: ctx,
	}
}

// NewPostTelephonyProvidersEdgesPhonesRebootParamsWithHTTPClient creates a new PostTelephonyProvidersEdgesPhonesRebootParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTelephonyProvidersEdgesPhonesRebootParamsWithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgesPhonesRebootParams {
	var ()
	return &PostTelephonyProvidersEdgesPhonesRebootParams{
		HTTPClient: client,
	}
}

/*PostTelephonyProvidersEdgesPhonesRebootParams contains all the parameters to send to the API endpoint
for the post telephony providers edges phones reboot operation typically these are written to a http.Request
*/
type PostTelephonyProvidersEdgesPhonesRebootParams struct {

	/*Body
	  Phones

	*/
	Body *models.PhonesReboot

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post telephony providers edges phones reboot params
func (o *PostTelephonyProvidersEdgesPhonesRebootParams) WithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgesPhonesRebootParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post telephony providers edges phones reboot params
func (o *PostTelephonyProvidersEdgesPhonesRebootParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post telephony providers edges phones reboot params
func (o *PostTelephonyProvidersEdgesPhonesRebootParams) WithContext(ctx context.Context) *PostTelephonyProvidersEdgesPhonesRebootParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post telephony providers edges phones reboot params
func (o *PostTelephonyProvidersEdgesPhonesRebootParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post telephony providers edges phones reboot params
func (o *PostTelephonyProvidersEdgesPhonesRebootParams) WithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgesPhonesRebootParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post telephony providers edges phones reboot params
func (o *PostTelephonyProvidersEdgesPhonesRebootParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post telephony providers edges phones reboot params
func (o *PostTelephonyProvidersEdgesPhonesRebootParams) WithBody(body *models.PhonesReboot) *PostTelephonyProvidersEdgesPhonesRebootParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post telephony providers edges phones reboot params
func (o *PostTelephonyProvidersEdgesPhonesRebootParams) SetBody(body *models.PhonesReboot) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTelephonyProvidersEdgesPhonesRebootParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
