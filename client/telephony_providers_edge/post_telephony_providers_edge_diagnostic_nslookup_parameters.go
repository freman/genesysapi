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

// NewPostTelephonyProvidersEdgeDiagnosticNslookupParams creates a new PostTelephonyProvidersEdgeDiagnosticNslookupParams object
// with the default values initialized.
func NewPostTelephonyProvidersEdgeDiagnosticNslookupParams() *PostTelephonyProvidersEdgeDiagnosticNslookupParams {
	var ()
	return &PostTelephonyProvidersEdgeDiagnosticNslookupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTelephonyProvidersEdgeDiagnosticNslookupParamsWithTimeout creates a new PostTelephonyProvidersEdgeDiagnosticNslookupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTelephonyProvidersEdgeDiagnosticNslookupParamsWithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgeDiagnosticNslookupParams {
	var ()
	return &PostTelephonyProvidersEdgeDiagnosticNslookupParams{

		timeout: timeout,
	}
}

// NewPostTelephonyProvidersEdgeDiagnosticNslookupParamsWithContext creates a new PostTelephonyProvidersEdgeDiagnosticNslookupParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTelephonyProvidersEdgeDiagnosticNslookupParamsWithContext(ctx context.Context) *PostTelephonyProvidersEdgeDiagnosticNslookupParams {
	var ()
	return &PostTelephonyProvidersEdgeDiagnosticNslookupParams{

		Context: ctx,
	}
}

// NewPostTelephonyProvidersEdgeDiagnosticNslookupParamsWithHTTPClient creates a new PostTelephonyProvidersEdgeDiagnosticNslookupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTelephonyProvidersEdgeDiagnosticNslookupParamsWithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgeDiagnosticNslookupParams {
	var ()
	return &PostTelephonyProvidersEdgeDiagnosticNslookupParams{
		HTTPClient: client,
	}
}

/*PostTelephonyProvidersEdgeDiagnosticNslookupParams contains all the parameters to send to the API endpoint
for the post telephony providers edge diagnostic nslookup operation typically these are written to a http.Request
*/
type PostTelephonyProvidersEdgeDiagnosticNslookupParams struct {

	/*Body
	  request payload to get network diagnostic

	*/
	Body *models.EdgeNetworkDiagnosticRequest
	/*EdgeID
	  Edge Id

	*/
	EdgeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post telephony providers edge diagnostic nslookup params
func (o *PostTelephonyProvidersEdgeDiagnosticNslookupParams) WithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgeDiagnosticNslookupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post telephony providers edge diagnostic nslookup params
func (o *PostTelephonyProvidersEdgeDiagnosticNslookupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post telephony providers edge diagnostic nslookup params
func (o *PostTelephonyProvidersEdgeDiagnosticNslookupParams) WithContext(ctx context.Context) *PostTelephonyProvidersEdgeDiagnosticNslookupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post telephony providers edge diagnostic nslookup params
func (o *PostTelephonyProvidersEdgeDiagnosticNslookupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post telephony providers edge diagnostic nslookup params
func (o *PostTelephonyProvidersEdgeDiagnosticNslookupParams) WithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgeDiagnosticNslookupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post telephony providers edge diagnostic nslookup params
func (o *PostTelephonyProvidersEdgeDiagnosticNslookupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post telephony providers edge diagnostic nslookup params
func (o *PostTelephonyProvidersEdgeDiagnosticNslookupParams) WithBody(body *models.EdgeNetworkDiagnosticRequest) *PostTelephonyProvidersEdgeDiagnosticNslookupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post telephony providers edge diagnostic nslookup params
func (o *PostTelephonyProvidersEdgeDiagnosticNslookupParams) SetBody(body *models.EdgeNetworkDiagnosticRequest) {
	o.Body = body
}

// WithEdgeID adds the edgeID to the post telephony providers edge diagnostic nslookup params
func (o *PostTelephonyProvidersEdgeDiagnosticNslookupParams) WithEdgeID(edgeID string) *PostTelephonyProvidersEdgeDiagnosticNslookupParams {
	o.SetEdgeID(edgeID)
	return o
}

// SetEdgeID adds the edgeId to the post telephony providers edge diagnostic nslookup params
func (o *PostTelephonyProvidersEdgeDiagnosticNslookupParams) SetEdgeID(edgeID string) {
	o.EdgeID = edgeID
}

// WriteToRequest writes these params to a swagger request
func (o *PostTelephonyProvidersEdgeDiagnosticNslookupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param edgeId
	if err := r.SetPathParam("edgeId", o.EdgeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
