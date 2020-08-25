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

// NewPostTelephonyProvidersEdgeDiagnosticTracepathParams creates a new PostTelephonyProvidersEdgeDiagnosticTracepathParams object
// with the default values initialized.
func NewPostTelephonyProvidersEdgeDiagnosticTracepathParams() *PostTelephonyProvidersEdgeDiagnosticTracepathParams {
	var ()
	return &PostTelephonyProvidersEdgeDiagnosticTracepathParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTelephonyProvidersEdgeDiagnosticTracepathParamsWithTimeout creates a new PostTelephonyProvidersEdgeDiagnosticTracepathParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTelephonyProvidersEdgeDiagnosticTracepathParamsWithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgeDiagnosticTracepathParams {
	var ()
	return &PostTelephonyProvidersEdgeDiagnosticTracepathParams{

		timeout: timeout,
	}
}

// NewPostTelephonyProvidersEdgeDiagnosticTracepathParamsWithContext creates a new PostTelephonyProvidersEdgeDiagnosticTracepathParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTelephonyProvidersEdgeDiagnosticTracepathParamsWithContext(ctx context.Context) *PostTelephonyProvidersEdgeDiagnosticTracepathParams {
	var ()
	return &PostTelephonyProvidersEdgeDiagnosticTracepathParams{

		Context: ctx,
	}
}

// NewPostTelephonyProvidersEdgeDiagnosticTracepathParamsWithHTTPClient creates a new PostTelephonyProvidersEdgeDiagnosticTracepathParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTelephonyProvidersEdgeDiagnosticTracepathParamsWithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgeDiagnosticTracepathParams {
	var ()
	return &PostTelephonyProvidersEdgeDiagnosticTracepathParams{
		HTTPClient: client,
	}
}

/*PostTelephonyProvidersEdgeDiagnosticTracepathParams contains all the parameters to send to the API endpoint
for the post telephony providers edge diagnostic tracepath operation typically these are written to a http.Request
*/
type PostTelephonyProvidersEdgeDiagnosticTracepathParams struct {

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

// WithTimeout adds the timeout to the post telephony providers edge diagnostic tracepath params
func (o *PostTelephonyProvidersEdgeDiagnosticTracepathParams) WithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgeDiagnosticTracepathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post telephony providers edge diagnostic tracepath params
func (o *PostTelephonyProvidersEdgeDiagnosticTracepathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post telephony providers edge diagnostic tracepath params
func (o *PostTelephonyProvidersEdgeDiagnosticTracepathParams) WithContext(ctx context.Context) *PostTelephonyProvidersEdgeDiagnosticTracepathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post telephony providers edge diagnostic tracepath params
func (o *PostTelephonyProvidersEdgeDiagnosticTracepathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post telephony providers edge diagnostic tracepath params
func (o *PostTelephonyProvidersEdgeDiagnosticTracepathParams) WithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgeDiagnosticTracepathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post telephony providers edge diagnostic tracepath params
func (o *PostTelephonyProvidersEdgeDiagnosticTracepathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post telephony providers edge diagnostic tracepath params
func (o *PostTelephonyProvidersEdgeDiagnosticTracepathParams) WithBody(body *models.EdgeNetworkDiagnosticRequest) *PostTelephonyProvidersEdgeDiagnosticTracepathParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post telephony providers edge diagnostic tracepath params
func (o *PostTelephonyProvidersEdgeDiagnosticTracepathParams) SetBody(body *models.EdgeNetworkDiagnosticRequest) {
	o.Body = body
}

// WithEdgeID adds the edgeID to the post telephony providers edge diagnostic tracepath params
func (o *PostTelephonyProvidersEdgeDiagnosticTracepathParams) WithEdgeID(edgeID string) *PostTelephonyProvidersEdgeDiagnosticTracepathParams {
	o.SetEdgeID(edgeID)
	return o
}

// SetEdgeID adds the edgeId to the post telephony providers edge diagnostic tracepath params
func (o *PostTelephonyProvidersEdgeDiagnosticTracepathParams) SetEdgeID(edgeID string) {
	o.EdgeID = edgeID
}

// WriteToRequest writes these params to a swagger request
func (o *PostTelephonyProvidersEdgeDiagnosticTracepathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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