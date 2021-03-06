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

// NewPutTelephonyProvidersEdgeLineParams creates a new PutTelephonyProvidersEdgeLineParams object
// with the default values initialized.
func NewPutTelephonyProvidersEdgeLineParams() *PutTelephonyProvidersEdgeLineParams {
	var ()
	return &PutTelephonyProvidersEdgeLineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTelephonyProvidersEdgeLineParamsWithTimeout creates a new PutTelephonyProvidersEdgeLineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTelephonyProvidersEdgeLineParamsWithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgeLineParams {
	var ()
	return &PutTelephonyProvidersEdgeLineParams{

		timeout: timeout,
	}
}

// NewPutTelephonyProvidersEdgeLineParamsWithContext creates a new PutTelephonyProvidersEdgeLineParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTelephonyProvidersEdgeLineParamsWithContext(ctx context.Context) *PutTelephonyProvidersEdgeLineParams {
	var ()
	return &PutTelephonyProvidersEdgeLineParams{

		Context: ctx,
	}
}

// NewPutTelephonyProvidersEdgeLineParamsWithHTTPClient creates a new PutTelephonyProvidersEdgeLineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTelephonyProvidersEdgeLineParamsWithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgeLineParams {
	var ()
	return &PutTelephonyProvidersEdgeLineParams{
		HTTPClient: client,
	}
}

/*PutTelephonyProvidersEdgeLineParams contains all the parameters to send to the API endpoint
for the put telephony providers edge line operation typically these are written to a http.Request
*/
type PutTelephonyProvidersEdgeLineParams struct {

	/*Body
	  Line

	*/
	Body *models.EdgeLine
	/*EdgeID
	  Edge ID

	*/
	EdgeID string
	/*LineID
	  Line ID

	*/
	LineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put telephony providers edge line params
func (o *PutTelephonyProvidersEdgeLineParams) WithTimeout(timeout time.Duration) *PutTelephonyProvidersEdgeLineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put telephony providers edge line params
func (o *PutTelephonyProvidersEdgeLineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put telephony providers edge line params
func (o *PutTelephonyProvidersEdgeLineParams) WithContext(ctx context.Context) *PutTelephonyProvidersEdgeLineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put telephony providers edge line params
func (o *PutTelephonyProvidersEdgeLineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put telephony providers edge line params
func (o *PutTelephonyProvidersEdgeLineParams) WithHTTPClient(client *http.Client) *PutTelephonyProvidersEdgeLineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put telephony providers edge line params
func (o *PutTelephonyProvidersEdgeLineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put telephony providers edge line params
func (o *PutTelephonyProvidersEdgeLineParams) WithBody(body *models.EdgeLine) *PutTelephonyProvidersEdgeLineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put telephony providers edge line params
func (o *PutTelephonyProvidersEdgeLineParams) SetBody(body *models.EdgeLine) {
	o.Body = body
}

// WithEdgeID adds the edgeID to the put telephony providers edge line params
func (o *PutTelephonyProvidersEdgeLineParams) WithEdgeID(edgeID string) *PutTelephonyProvidersEdgeLineParams {
	o.SetEdgeID(edgeID)
	return o
}

// SetEdgeID adds the edgeId to the put telephony providers edge line params
func (o *PutTelephonyProvidersEdgeLineParams) SetEdgeID(edgeID string) {
	o.EdgeID = edgeID
}

// WithLineID adds the lineID to the put telephony providers edge line params
func (o *PutTelephonyProvidersEdgeLineParams) WithLineID(lineID string) *PutTelephonyProvidersEdgeLineParams {
	o.SetLineID(lineID)
	return o
}

// SetLineID adds the lineId to the put telephony providers edge line params
func (o *PutTelephonyProvidersEdgeLineParams) SetLineID(lineID string) {
	o.LineID = lineID
}

// WriteToRequest writes these params to a swagger request
func (o *PutTelephonyProvidersEdgeLineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param lineId
	if err := r.SetPathParam("lineId", o.LineID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
