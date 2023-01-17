// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewGetRoutingEmailDomainRouteParams creates a new GetRoutingEmailDomainRouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoutingEmailDomainRouteParams() *GetRoutingEmailDomainRouteParams {
	return &GetRoutingEmailDomainRouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingEmailDomainRouteParamsWithTimeout creates a new GetRoutingEmailDomainRouteParams object
// with the ability to set a timeout on a request.
func NewGetRoutingEmailDomainRouteParamsWithTimeout(timeout time.Duration) *GetRoutingEmailDomainRouteParams {
	return &GetRoutingEmailDomainRouteParams{
		timeout: timeout,
	}
}

// NewGetRoutingEmailDomainRouteParamsWithContext creates a new GetRoutingEmailDomainRouteParams object
// with the ability to set a context for a request.
func NewGetRoutingEmailDomainRouteParamsWithContext(ctx context.Context) *GetRoutingEmailDomainRouteParams {
	return &GetRoutingEmailDomainRouteParams{
		Context: ctx,
	}
}

// NewGetRoutingEmailDomainRouteParamsWithHTTPClient creates a new GetRoutingEmailDomainRouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoutingEmailDomainRouteParamsWithHTTPClient(client *http.Client) *GetRoutingEmailDomainRouteParams {
	return &GetRoutingEmailDomainRouteParams{
		HTTPClient: client,
	}
}

/*
GetRoutingEmailDomainRouteParams contains all the parameters to send to the API endpoint

	for the get routing email domain route operation.

	Typically these are written to a http.Request.
*/
type GetRoutingEmailDomainRouteParams struct {

	/* DomainName.

	   email domain
	*/
	DomainName string

	/* RouteID.

	   route ID
	*/
	RouteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get routing email domain route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingEmailDomainRouteParams) WithDefaults() *GetRoutingEmailDomainRouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get routing email domain route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingEmailDomainRouteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get routing email domain route params
func (o *GetRoutingEmailDomainRouteParams) WithTimeout(timeout time.Duration) *GetRoutingEmailDomainRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing email domain route params
func (o *GetRoutingEmailDomainRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing email domain route params
func (o *GetRoutingEmailDomainRouteParams) WithContext(ctx context.Context) *GetRoutingEmailDomainRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing email domain route params
func (o *GetRoutingEmailDomainRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing email domain route params
func (o *GetRoutingEmailDomainRouteParams) WithHTTPClient(client *http.Client) *GetRoutingEmailDomainRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing email domain route params
func (o *GetRoutingEmailDomainRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainName adds the domainName to the get routing email domain route params
func (o *GetRoutingEmailDomainRouteParams) WithDomainName(domainName string) *GetRoutingEmailDomainRouteParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the get routing email domain route params
func (o *GetRoutingEmailDomainRouteParams) SetDomainName(domainName string) {
	o.DomainName = domainName
}

// WithRouteID adds the routeID to the get routing email domain route params
func (o *GetRoutingEmailDomainRouteParams) WithRouteID(routeID string) *GetRoutingEmailDomainRouteParams {
	o.SetRouteID(routeID)
	return o
}

// SetRouteID adds the routeId to the get routing email domain route params
func (o *GetRoutingEmailDomainRouteParams) SetRouteID(routeID string) {
	o.RouteID = routeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingEmailDomainRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainName
	if err := r.SetPathParam("domainName", o.DomainName); err != nil {
		return err
	}

	// path param routeId
	if err := r.SetPathParam("routeId", o.RouteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
