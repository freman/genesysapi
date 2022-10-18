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

// NewDeleteRoutingEmailOutboundDomainParams creates a new DeleteRoutingEmailOutboundDomainParams object
// with the default values initialized.
func NewDeleteRoutingEmailOutboundDomainParams() *DeleteRoutingEmailOutboundDomainParams {
	var ()
	return &DeleteRoutingEmailOutboundDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoutingEmailOutboundDomainParamsWithTimeout creates a new DeleteRoutingEmailOutboundDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRoutingEmailOutboundDomainParamsWithTimeout(timeout time.Duration) *DeleteRoutingEmailOutboundDomainParams {
	var ()
	return &DeleteRoutingEmailOutboundDomainParams{

		timeout: timeout,
	}
}

// NewDeleteRoutingEmailOutboundDomainParamsWithContext creates a new DeleteRoutingEmailOutboundDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRoutingEmailOutboundDomainParamsWithContext(ctx context.Context) *DeleteRoutingEmailOutboundDomainParams {
	var ()
	return &DeleteRoutingEmailOutboundDomainParams{

		Context: ctx,
	}
}

// NewDeleteRoutingEmailOutboundDomainParamsWithHTTPClient creates a new DeleteRoutingEmailOutboundDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRoutingEmailOutboundDomainParamsWithHTTPClient(client *http.Client) *DeleteRoutingEmailOutboundDomainParams {
	var ()
	return &DeleteRoutingEmailOutboundDomainParams{
		HTTPClient: client,
	}
}

/*DeleteRoutingEmailOutboundDomainParams contains all the parameters to send to the API endpoint
for the delete routing email outbound domain operation typically these are written to a http.Request
*/
type DeleteRoutingEmailOutboundDomainParams struct {

	/*DomainID
	  domain ID

	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete routing email outbound domain params
func (o *DeleteRoutingEmailOutboundDomainParams) WithTimeout(timeout time.Duration) *DeleteRoutingEmailOutboundDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete routing email outbound domain params
func (o *DeleteRoutingEmailOutboundDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete routing email outbound domain params
func (o *DeleteRoutingEmailOutboundDomainParams) WithContext(ctx context.Context) *DeleteRoutingEmailOutboundDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete routing email outbound domain params
func (o *DeleteRoutingEmailOutboundDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete routing email outbound domain params
func (o *DeleteRoutingEmailOutboundDomainParams) WithHTTPClient(client *http.Client) *DeleteRoutingEmailOutboundDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete routing email outbound domain params
func (o *DeleteRoutingEmailOutboundDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the delete routing email outbound domain params
func (o *DeleteRoutingEmailOutboundDomainParams) WithDomainID(domainID string) *DeleteRoutingEmailOutboundDomainParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the delete routing email outbound domain params
func (o *DeleteRoutingEmailOutboundDomainParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoutingEmailOutboundDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
