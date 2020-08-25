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
)

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateParams creates a new GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateParams() *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams {
	var ()
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateParamsWithTimeout creates a new GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams {
	var ()
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams{

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateParamsWithContext creates a new GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams {
	var ()
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams{

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesPhonebasesettingsTemplateParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams {
	var ()
	return &GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams{
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams contains all the parameters to send to the API endpoint
for the get telephony providers edges phonebasesettings template operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams struct {

	/*PhoneMetabaseID
	  The id of a metabase object upon which to base this Phone Base Settings

	*/
	PhoneMetabaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edges phonebasesettings template params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges phonebasesettings template params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges phonebasesettings template params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges phonebasesettings template params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges phonebasesettings template params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges phonebasesettings template params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPhoneMetabaseID adds the phoneMetabaseID to the get telephony providers edges phonebasesettings template params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams) WithPhoneMetabaseID(phoneMetabaseID string) *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams {
	o.SetPhoneMetabaseID(phoneMetabaseID)
	return o
}

// SetPhoneMetabaseID adds the phoneMetabaseId to the get telephony providers edges phonebasesettings template params
func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams) SetPhoneMetabaseID(phoneMetabaseID string) {
	o.PhoneMetabaseID = phoneMetabaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesPhonebasesettingsTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param phoneMetabaseId
	qrPhoneMetabaseID := o.PhoneMetabaseID
	qPhoneMetabaseID := qrPhoneMetabaseID
	if qPhoneMetabaseID != "" {
		if err := r.SetQueryParam("phoneMetabaseId", qPhoneMetabaseID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}