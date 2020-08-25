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

// NewGetRoutingEmailDomainParams creates a new GetRoutingEmailDomainParams object
// with the default values initialized.
func NewGetRoutingEmailDomainParams() *GetRoutingEmailDomainParams {
	var ()
	return &GetRoutingEmailDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingEmailDomainParamsWithTimeout creates a new GetRoutingEmailDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoutingEmailDomainParamsWithTimeout(timeout time.Duration) *GetRoutingEmailDomainParams {
	var ()
	return &GetRoutingEmailDomainParams{

		timeout: timeout,
	}
}

// NewGetRoutingEmailDomainParamsWithContext creates a new GetRoutingEmailDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoutingEmailDomainParamsWithContext(ctx context.Context) *GetRoutingEmailDomainParams {
	var ()
	return &GetRoutingEmailDomainParams{

		Context: ctx,
	}
}

// NewGetRoutingEmailDomainParamsWithHTTPClient creates a new GetRoutingEmailDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoutingEmailDomainParamsWithHTTPClient(client *http.Client) *GetRoutingEmailDomainParams {
	var ()
	return &GetRoutingEmailDomainParams{
		HTTPClient: client,
	}
}

/*GetRoutingEmailDomainParams contains all the parameters to send to the API endpoint
for the get routing email domain operation typically these are written to a http.Request
*/
type GetRoutingEmailDomainParams struct {

	/*DomainID
	  domain ID

	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routing email domain params
func (o *GetRoutingEmailDomainParams) WithTimeout(timeout time.Duration) *GetRoutingEmailDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing email domain params
func (o *GetRoutingEmailDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing email domain params
func (o *GetRoutingEmailDomainParams) WithContext(ctx context.Context) *GetRoutingEmailDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing email domain params
func (o *GetRoutingEmailDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing email domain params
func (o *GetRoutingEmailDomainParams) WithHTTPClient(client *http.Client) *GetRoutingEmailDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing email domain params
func (o *GetRoutingEmailDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get routing email domain params
func (o *GetRoutingEmailDomainParams) WithDomainID(domainID string) *GetRoutingEmailDomainParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get routing email domain params
func (o *GetRoutingEmailDomainParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingEmailDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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