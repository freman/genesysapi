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

// NewGetTelephonyProvidersEdgeParams creates a new GetTelephonyProvidersEdgeParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgeParams() *GetTelephonyProvidersEdgeParams {
	var ()
	return &GetTelephonyProvidersEdgeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgeParamsWithTimeout creates a new GetTelephonyProvidersEdgeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgeParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgeParams {
	var ()
	return &GetTelephonyProvidersEdgeParams{

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgeParamsWithContext creates a new GetTelephonyProvidersEdgeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgeParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgeParams {
	var ()
	return &GetTelephonyProvidersEdgeParams{

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgeParamsWithHTTPClient creates a new GetTelephonyProvidersEdgeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgeParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgeParams {
	var ()
	return &GetTelephonyProvidersEdgeParams{
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgeParams contains all the parameters to send to the API endpoint
for the get telephony providers edge operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgeParams struct {

	/*EdgeID
	  Edge ID

	*/
	EdgeID string
	/*Expand
	  Fields to expand in the response, comma-separated

	*/
	Expand []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edge params
func (o *GetTelephonyProvidersEdgeParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edge params
func (o *GetTelephonyProvidersEdgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edge params
func (o *GetTelephonyProvidersEdgeParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edge params
func (o *GetTelephonyProvidersEdgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edge params
func (o *GetTelephonyProvidersEdgeParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edge params
func (o *GetTelephonyProvidersEdgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEdgeID adds the edgeID to the get telephony providers edge params
func (o *GetTelephonyProvidersEdgeParams) WithEdgeID(edgeID string) *GetTelephonyProvidersEdgeParams {
	o.SetEdgeID(edgeID)
	return o
}

// SetEdgeID adds the edgeId to the get telephony providers edge params
func (o *GetTelephonyProvidersEdgeParams) SetEdgeID(edgeID string) {
	o.EdgeID = edgeID
}

// WithExpand adds the expand to the get telephony providers edge params
func (o *GetTelephonyProvidersEdgeParams) WithExpand(expand []string) *GetTelephonyProvidersEdgeParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get telephony providers edge params
func (o *GetTelephonyProvidersEdgeParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param edgeId
	if err := r.SetPathParam("edgeId", o.EdgeID); err != nil {
		return err
	}

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
