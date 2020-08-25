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

// NewPutTelephonyProvidersEdgesEdgegroupParams creates a new PutTelephonyProvidersEdgesEdgegroupParams object
// with the default values initialized.
func NewPutTelephonyProvidersEdgesEdgegroupParams() *PutTelephonyProvidersEdgesEdgegroupParams {
	var ()
	return &PutTelephonyProvidersEdgesEdgegroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTelephonyProvidersEdgesEdgegroupParamsWithTimeout creates a new PutTelephonyProvidersEdgesEdgegroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTelephonyProvidersEdgesEdgegroupParamsWithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesEdgegroupParams {
	var ()
	return &PutTelephonyProvidersEdgesEdgegroupParams{

		timeout: timeout,
	}
}

// NewPutTelephonyProvidersEdgesEdgegroupParamsWithContext creates a new PutTelephonyProvidersEdgesEdgegroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTelephonyProvidersEdgesEdgegroupParamsWithContext(ctx context.Context) *PutTelephonyProvidersEdgesEdgegroupParams {
	var ()
	return &PutTelephonyProvidersEdgesEdgegroupParams{

		Context: ctx,
	}
}

// NewPutTelephonyProvidersEdgesEdgegroupParamsWithHTTPClient creates a new PutTelephonyProvidersEdgesEdgegroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTelephonyProvidersEdgesEdgegroupParamsWithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesEdgegroupParams {
	var ()
	return &PutTelephonyProvidersEdgesEdgegroupParams{
		HTTPClient: client,
	}
}

/*PutTelephonyProvidersEdgesEdgegroupParams contains all the parameters to send to the API endpoint
for the put telephony providers edges edgegroup operation typically these are written to a http.Request
*/
type PutTelephonyProvidersEdgesEdgegroupParams struct {

	/*Body
	  EdgeGroup

	*/
	Body *models.EdgeGroup
	/*EdgeGroupID
	  Edge group ID

	*/
	EdgeGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put telephony providers edges edgegroup params
func (o *PutTelephonyProvidersEdgesEdgegroupParams) WithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgesEdgegroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put telephony providers edges edgegroup params
func (o *PutTelephonyProvidersEdgesEdgegroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put telephony providers edges edgegroup params
func (o *PutTelephonyProvidersEdgesEdgegroupParams) WithContext(ctx context.Context) *PutTelephonyProvidersEdgesEdgegroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put telephony providers edges edgegroup params
func (o *PutTelephonyProvidersEdgesEdgegroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put telephony providers edges edgegroup params
func (o *PutTelephonyProvidersEdgesEdgegroupParams) WithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgesEdgegroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put telephony providers edges edgegroup params
func (o *PutTelephonyProvidersEdgesEdgegroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put telephony providers edges edgegroup params
func (o *PutTelephonyProvidersEdgesEdgegroupParams) WithBody(body *models.EdgeGroup) *PutTelephonyProvidersEdgesEdgegroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put telephony providers edges edgegroup params
func (o *PutTelephonyProvidersEdgesEdgegroupParams) SetBody(body *models.EdgeGroup) {
	o.Body = body
}

// WithEdgeGroupID adds the edgeGroupID to the put telephony providers edges edgegroup params
func (o *PutTelephonyProvidersEdgesEdgegroupParams) WithEdgeGroupID(edgeGroupID string) *PutTelephonyProvidersEdgesEdgegroupParams {
	o.SetEdgeGroupID(edgeGroupID)
	return o
}

// SetEdgeGroupID adds the edgeGroupId to the put telephony providers edges edgegroup params
func (o *PutTelephonyProvidersEdgesEdgegroupParams) SetEdgeGroupID(edgeGroupID string) {
	o.EdgeGroupID = edgeGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *PutTelephonyProvidersEdgesEdgegroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param edgeGroupId
	if err := r.SetPathParam("edgeGroupId", o.EdgeGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}