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

// NewDeleteRoutingEmailDomainParams creates a new DeleteRoutingEmailDomainParams object
// with the default values initialized.
func NewDeleteRoutingEmailDomainParams() *DeleteRoutingEmailDomainParams {
	var ()
	return &DeleteRoutingEmailDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoutingEmailDomainParamsWithTimeout creates a new DeleteRoutingEmailDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRoutingEmailDomainParamsWithTimeout(timeout time.Duration) *DeleteRoutingEmailDomainParams {
	var ()
	return &DeleteRoutingEmailDomainParams{

		timeout: timeout,
	}
}

// NewDeleteRoutingEmailDomainParamsWithContext creates a new DeleteRoutingEmailDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRoutingEmailDomainParamsWithContext(ctx context.Context) *DeleteRoutingEmailDomainParams {
	var ()
	return &DeleteRoutingEmailDomainParams{

		Context: ctx,
	}
}

// NewDeleteRoutingEmailDomainParamsWithHTTPClient creates a new DeleteRoutingEmailDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRoutingEmailDomainParamsWithHTTPClient(client *http.Client) *DeleteRoutingEmailDomainParams {
	var ()
	return &DeleteRoutingEmailDomainParams{
		HTTPClient: client,
	}
}

/*DeleteRoutingEmailDomainParams contains all the parameters to send to the API endpoint
for the delete routing email domain operation typically these are written to a http.Request
*/
type DeleteRoutingEmailDomainParams struct {

	/*DomainID
	  domain ID

	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete routing email domain params
func (o *DeleteRoutingEmailDomainParams) WithTimeout(timeout time.Duration) *DeleteRoutingEmailDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete routing email domain params
func (o *DeleteRoutingEmailDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete routing email domain params
func (o *DeleteRoutingEmailDomainParams) WithContext(ctx context.Context) *DeleteRoutingEmailDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete routing email domain params
func (o *DeleteRoutingEmailDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete routing email domain params
func (o *DeleteRoutingEmailDomainParams) WithHTTPClient(client *http.Client) *DeleteRoutingEmailDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete routing email domain params
func (o *DeleteRoutingEmailDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the delete routing email domain params
func (o *DeleteRoutingEmailDomainParams) WithDomainID(domainID string) *DeleteRoutingEmailDomainParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the delete routing email domain params
func (o *DeleteRoutingEmailDomainParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoutingEmailDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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