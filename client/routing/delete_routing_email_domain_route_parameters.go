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

// NewDeleteRoutingEmailDomainRouteParams creates a new DeleteRoutingEmailDomainRouteParams object
// with the default values initialized.
func NewDeleteRoutingEmailDomainRouteParams() *DeleteRoutingEmailDomainRouteParams {
	var ()
	return &DeleteRoutingEmailDomainRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoutingEmailDomainRouteParamsWithTimeout creates a new DeleteRoutingEmailDomainRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRoutingEmailDomainRouteParamsWithTimeout(timeout time.Duration) *DeleteRoutingEmailDomainRouteParams {
	var ()
	return &DeleteRoutingEmailDomainRouteParams{

		timeout: timeout,
	}
}

// NewDeleteRoutingEmailDomainRouteParamsWithContext creates a new DeleteRoutingEmailDomainRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRoutingEmailDomainRouteParamsWithContext(ctx context.Context) *DeleteRoutingEmailDomainRouteParams {
	var ()
	return &DeleteRoutingEmailDomainRouteParams{

		Context: ctx,
	}
}

// NewDeleteRoutingEmailDomainRouteParamsWithHTTPClient creates a new DeleteRoutingEmailDomainRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRoutingEmailDomainRouteParamsWithHTTPClient(client *http.Client) *DeleteRoutingEmailDomainRouteParams {
	var ()
	return &DeleteRoutingEmailDomainRouteParams{
		HTTPClient: client,
	}
}

/*DeleteRoutingEmailDomainRouteParams contains all the parameters to send to the API endpoint
for the delete routing email domain route operation typically these are written to a http.Request
*/
type DeleteRoutingEmailDomainRouteParams struct {

	/*DomainName
	  email domain

	*/
	DomainName string
	/*RouteID
	  route ID

	*/
	RouteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete routing email domain route params
func (o *DeleteRoutingEmailDomainRouteParams) WithTimeout(timeout time.Duration) *DeleteRoutingEmailDomainRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete routing email domain route params
func (o *DeleteRoutingEmailDomainRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete routing email domain route params
func (o *DeleteRoutingEmailDomainRouteParams) WithContext(ctx context.Context) *DeleteRoutingEmailDomainRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete routing email domain route params
func (o *DeleteRoutingEmailDomainRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete routing email domain route params
func (o *DeleteRoutingEmailDomainRouteParams) WithHTTPClient(client *http.Client) *DeleteRoutingEmailDomainRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete routing email domain route params
func (o *DeleteRoutingEmailDomainRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainName adds the domainName to the delete routing email domain route params
func (o *DeleteRoutingEmailDomainRouteParams) WithDomainName(domainName string) *DeleteRoutingEmailDomainRouteParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the delete routing email domain route params
func (o *DeleteRoutingEmailDomainRouteParams) SetDomainName(domainName string) {
	o.DomainName = domainName
}

// WithRouteID adds the routeID to the delete routing email domain route params
func (o *DeleteRoutingEmailDomainRouteParams) WithRouteID(routeID string) *DeleteRoutingEmailDomainRouteParams {
	o.SetRouteID(routeID)
	return o
}

// SetRouteID adds the routeId to the delete routing email domain route params
func (o *DeleteRoutingEmailDomainRouteParams) SetRouteID(routeID string) {
	o.RouteID = routeID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoutingEmailDomainRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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