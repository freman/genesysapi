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

// NewGetTelephonyProvidersEdgesLogicalinterfacesParams creates a new GetTelephonyProvidersEdgesLogicalinterfacesParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgesLogicalinterfacesParams() *GetTelephonyProvidersEdgesLogicalinterfacesParams {
	var ()
	return &GetTelephonyProvidersEdgesLogicalinterfacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesParamsWithTimeout creates a new GetTelephonyProvidersEdgesLogicalinterfacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgesLogicalinterfacesParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesLogicalinterfacesParams {
	var ()
	return &GetTelephonyProvidersEdgesLogicalinterfacesParams{

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesParamsWithContext creates a new GetTelephonyProvidersEdgesLogicalinterfacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgesLogicalinterfacesParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesLogicalinterfacesParams {
	var ()
	return &GetTelephonyProvidersEdgesLogicalinterfacesParams{

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesLogicalinterfacesParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesLogicalinterfacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgesLogicalinterfacesParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesLogicalinterfacesParams {
	var ()
	return &GetTelephonyProvidersEdgesLogicalinterfacesParams{
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgesLogicalinterfacesParams contains all the parameters to send to the API endpoint
for the get telephony providers edges logicalinterfaces operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgesLogicalinterfacesParams struct {

	/*EdgeIds
	  Comma separated list of Edge Id's

	*/
	EdgeIds string
	/*Expand
	  Field to expand in the response

	*/
	Expand []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edges logicalinterfaces params
func (o *GetTelephonyProvidersEdgesLogicalinterfacesParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesLogicalinterfacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges logicalinterfaces params
func (o *GetTelephonyProvidersEdgesLogicalinterfacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges logicalinterfaces params
func (o *GetTelephonyProvidersEdgesLogicalinterfacesParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesLogicalinterfacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges logicalinterfaces params
func (o *GetTelephonyProvidersEdgesLogicalinterfacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges logicalinterfaces params
func (o *GetTelephonyProvidersEdgesLogicalinterfacesParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesLogicalinterfacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges logicalinterfaces params
func (o *GetTelephonyProvidersEdgesLogicalinterfacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEdgeIds adds the edgeIds to the get telephony providers edges logicalinterfaces params
func (o *GetTelephonyProvidersEdgesLogicalinterfacesParams) WithEdgeIds(edgeIds string) *GetTelephonyProvidersEdgesLogicalinterfacesParams {
	o.SetEdgeIds(edgeIds)
	return o
}

// SetEdgeIds adds the edgeIds to the get telephony providers edges logicalinterfaces params
func (o *GetTelephonyProvidersEdgesLogicalinterfacesParams) SetEdgeIds(edgeIds string) {
	o.EdgeIds = edgeIds
}

// WithExpand adds the expand to the get telephony providers edges logicalinterfaces params
func (o *GetTelephonyProvidersEdgesLogicalinterfacesParams) WithExpand(expand []string) *GetTelephonyProvidersEdgesLogicalinterfacesParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get telephony providers edges logicalinterfaces params
func (o *GetTelephonyProvidersEdgesLogicalinterfacesParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesLogicalinterfacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param edgeIds
	qrEdgeIds := o.EdgeIds
	qEdgeIds := qrEdgeIds
	if qEdgeIds != "" {
		if err := r.SetQueryParam("edgeIds", qEdgeIds); err != nil {
			return err
		}
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
