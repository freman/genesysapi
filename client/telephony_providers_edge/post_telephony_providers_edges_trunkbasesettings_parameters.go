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

// NewPostTelephonyProvidersEdgesTrunkbasesettingsParams creates a new PostTelephonyProvidersEdgesTrunkbasesettingsParams object
// with the default values initialized.
func NewPostTelephonyProvidersEdgesTrunkbasesettingsParams() *PostTelephonyProvidersEdgesTrunkbasesettingsParams {
	var ()
	return &PostTelephonyProvidersEdgesTrunkbasesettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTelephonyProvidersEdgesTrunkbasesettingsParamsWithTimeout creates a new PostTelephonyProvidersEdgesTrunkbasesettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTelephonyProvidersEdgesTrunkbasesettingsParamsWithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgesTrunkbasesettingsParams {
	var ()
	return &PostTelephonyProvidersEdgesTrunkbasesettingsParams{

		timeout: timeout,
	}
}

// NewPostTelephonyProvidersEdgesTrunkbasesettingsParamsWithContext creates a new PostTelephonyProvidersEdgesTrunkbasesettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTelephonyProvidersEdgesTrunkbasesettingsParamsWithContext(ctx context.Context) *PostTelephonyProvidersEdgesTrunkbasesettingsParams {
	var ()
	return &PostTelephonyProvidersEdgesTrunkbasesettingsParams{

		Context: ctx,
	}
}

// NewPostTelephonyProvidersEdgesTrunkbasesettingsParamsWithHTTPClient creates a new PostTelephonyProvidersEdgesTrunkbasesettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTelephonyProvidersEdgesTrunkbasesettingsParamsWithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgesTrunkbasesettingsParams {
	var ()
	return &PostTelephonyProvidersEdgesTrunkbasesettingsParams{
		HTTPClient: client,
	}
}

/*PostTelephonyProvidersEdgesTrunkbasesettingsParams contains all the parameters to send to the API endpoint
for the post telephony providers edges trunkbasesettings operation typically these are written to a http.Request
*/
type PostTelephonyProvidersEdgesTrunkbasesettingsParams struct {

	/*Body
	  Trunk base settings

	*/
	Body *models.TrunkBase

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post telephony providers edges trunkbasesettings params
func (o *PostTelephonyProvidersEdgesTrunkbasesettingsParams) WithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgesTrunkbasesettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post telephony providers edges trunkbasesettings params
func (o *PostTelephonyProvidersEdgesTrunkbasesettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post telephony providers edges trunkbasesettings params
func (o *PostTelephonyProvidersEdgesTrunkbasesettingsParams) WithContext(ctx context.Context) *PostTelephonyProvidersEdgesTrunkbasesettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post telephony providers edges trunkbasesettings params
func (o *PostTelephonyProvidersEdgesTrunkbasesettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post telephony providers edges trunkbasesettings params
func (o *PostTelephonyProvidersEdgesTrunkbasesettingsParams) WithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgesTrunkbasesettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post telephony providers edges trunkbasesettings params
func (o *PostTelephonyProvidersEdgesTrunkbasesettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post telephony providers edges trunkbasesettings params
func (o *PostTelephonyProvidersEdgesTrunkbasesettingsParams) WithBody(body *models.TrunkBase) *PostTelephonyProvidersEdgesTrunkbasesettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post telephony providers edges trunkbasesettings params
func (o *PostTelephonyProvidersEdgesTrunkbasesettingsParams) SetBody(body *models.TrunkBase) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTelephonyProvidersEdgesTrunkbasesettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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