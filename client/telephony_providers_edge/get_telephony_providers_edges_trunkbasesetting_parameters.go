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
	"github.com/go-openapi/swag"
)

// NewGetTelephonyProvidersEdgesTrunkbasesettingParams creates a new GetTelephonyProvidersEdgesTrunkbasesettingParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesTrunkbasesettingParams() *GetTelephonyProvidersEdgesTrunkbasesettingParams {
	var ()
	return &GetTelephonyProvidersEdgesTrunkbasesettingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesTrunkbasesettingParamsWithTimeout creates a new GetTelephonyProvidersEdgesTrunkbasesettingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesTrunkbasesettingParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesTrunkbasesettingParams {
	var ()
	return &GetTelephonyProvidersEdgesTrunkbasesettingParams{

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesTrunkbasesettingParamsWithContext creates a new GetTelephonyProvidersEdgesTrunkbasesettingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesTrunkbasesettingParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesTrunkbasesettingParams {
	var ()
	return &GetTelephonyProvidersEdgesTrunkbasesettingParams{

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesTrunkbasesettingParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesTrunkbasesettingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesTrunkbasesettingParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesTrunkbasesettingParams {
	var ()
	return &GetTelephonyProvidersEdgesTrunkbasesettingParams{
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesTrunkbasesettingParams contains all the parameters to send to the API endpoint
for the get telephony providers edges trunkbasesetting operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesTrunkbasesettingParams struct {

	/*IgnoreHidden
	  Set this to true to not receive trunk properties that are meant to be hidden or for internal system usage only.

	*/
	IgnoreHidden *bool
	/*TrunkBaseSettingsID
	  Trunk Base ID

	*/
	TrunkBaseSettingsID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edges trunkbasesetting params
func (o *GetTelephonyProvidersEdgesTrunkbasesettingParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesTrunkbasesettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges trunkbasesetting params
func (o *GetTelephonyProvidersEdgesTrunkbasesettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges trunkbasesetting params
func (o *GetTelephonyProvidersEdgesTrunkbasesettingParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesTrunkbasesettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges trunkbasesetting params
func (o *GetTelephonyProvidersEdgesTrunkbasesettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges trunkbasesetting params
func (o *GetTelephonyProvidersEdgesTrunkbasesettingParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesTrunkbasesettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges trunkbasesetting params
func (o *GetTelephonyProvidersEdgesTrunkbasesettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIgnoreHidden adds the ignoreHidden to the get telephony providers edges trunkbasesetting params
func (o *GetTelephonyProvidersEdgesTrunkbasesettingParams) WithIgnoreHidden(ignoreHidden *bool) *GetTelephonyProvidersEdgesTrunkbasesettingParams {
	o.SetIgnoreHidden(ignoreHidden)
	return o
}

// SetIgnoreHidden adds the ignoreHidden to the get telephony providers edges trunkbasesetting params
func (o *GetTelephonyProvidersEdgesTrunkbasesettingParams) SetIgnoreHidden(ignoreHidden *bool) {
	o.IgnoreHidden = ignoreHidden
}

// WithTrunkBaseSettingsID adds the trunkBaseSettingsID to the get telephony providers edges trunkbasesetting params
func (o *GetTelephonyProvidersEdgesTrunkbasesettingParams) WithTrunkBaseSettingsID(trunkBaseSettingsID string) *GetTelephonyProvidersEdgesTrunkbasesettingParams {
	o.SetTrunkBaseSettingsID(trunkBaseSettingsID)
	return o
}

// SetTrunkBaseSettingsID adds the trunkBaseSettingsId to the get telephony providers edges trunkbasesetting params
func (o *GetTelephonyProvidersEdgesTrunkbasesettingParams) SetTrunkBaseSettingsID(trunkBaseSettingsID string) {
	o.TrunkBaseSettingsID = trunkBaseSettingsID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesTrunkbasesettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IgnoreHidden != nil {

		// query param ignoreHidden
		var qrIgnoreHidden bool
		if o.IgnoreHidden != nil {
			qrIgnoreHidden = *o.IgnoreHidden
		}
		qIgnoreHidden := swag.FormatBool(qrIgnoreHidden)
		if qIgnoreHidden != "" {
			if err := r.SetQueryParam("ignoreHidden", qIgnoreHidden); err != nil {
				return err
			}
		}

	}

	// path param trunkBaseSettingsId
	if err := r.SetPathParam("trunkBaseSettingsId", o.TrunkBaseSettingsID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
