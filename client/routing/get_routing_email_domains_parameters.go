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
	"github.com/go-openapi/swag"
)

// NewGetRoutingEmailDomainsParams creates a new GetRoutingEmailDomainsParams object
// with the default values initialized.
func NewGetRoutingEmailDomainsParams() *GetRoutingEmailDomainsParams {
	var (
		excludeStatusDefault = bool(false)
	)
	return &GetRoutingEmailDomainsParams{
		ExcludeStatus: &excludeStatusDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingEmailDomainsParamsWithTimeout creates a new GetRoutingEmailDomainsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoutingEmailDomainsParamsWithTimeout(timeout time.Duration) *GetRoutingEmailDomainsParams {
	var (
		excludeStatusDefault = bool(false)
	)
	return &GetRoutingEmailDomainsParams{
		ExcludeStatus: &excludeStatusDefault,

		timeout: timeout,
	}
}

// NewGetRoutingEmailDomainsParamsWithContext creates a new GetRoutingEmailDomainsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoutingEmailDomainsParamsWithContext(ctx context.Context) *GetRoutingEmailDomainsParams {
	var (
		excludeStatusDefault = bool(false)
	)
	return &GetRoutingEmailDomainsParams{
		ExcludeStatus: &excludeStatusDefault,

		Context: ctx,
	}
}

// NewGetRoutingEmailDomainsParamsWithHTTPClient creates a new GetRoutingEmailDomainsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoutingEmailDomainsParamsWithHTTPClient(client *http.Client) *GetRoutingEmailDomainsParams {
	var (
		excludeStatusDefault = bool(false)
	)
	return &GetRoutingEmailDomainsParams{
		ExcludeStatus: &excludeStatusDefault,
		HTTPClient:    client,
	}
}

/*GetRoutingEmailDomainsParams contains all the parameters to send to the API endpoint
for the get routing email domains operation typically these are written to a http.Request
*/
type GetRoutingEmailDomainsParams struct {

	/*ExcludeStatus
	  Exclude MX record data

	*/
	ExcludeStatus *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routing email domains params
func (o *GetRoutingEmailDomainsParams) WithTimeout(timeout time.Duration) *GetRoutingEmailDomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing email domains params
func (o *GetRoutingEmailDomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing email domains params
func (o *GetRoutingEmailDomainsParams) WithContext(ctx context.Context) *GetRoutingEmailDomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing email domains params
func (o *GetRoutingEmailDomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing email domains params
func (o *GetRoutingEmailDomainsParams) WithHTTPClient(client *http.Client) *GetRoutingEmailDomainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing email domains params
func (o *GetRoutingEmailDomainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExcludeStatus adds the excludeStatus to the get routing email domains params
func (o *GetRoutingEmailDomainsParams) WithExcludeStatus(excludeStatus *bool) *GetRoutingEmailDomainsParams {
	o.SetExcludeStatus(excludeStatus)
	return o
}

// SetExcludeStatus adds the excludeStatus to the get routing email domains params
func (o *GetRoutingEmailDomainsParams) SetExcludeStatus(excludeStatus *bool) {
	o.ExcludeStatus = excludeStatus
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingEmailDomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExcludeStatus != nil {

		// query param excludeStatus
		var qrExcludeStatus bool
		if o.ExcludeStatus != nil {
			qrExcludeStatus = *o.ExcludeStatus
		}
		qExcludeStatus := swag.FormatBool(qrExcludeStatus)
		if qExcludeStatus != "" {
			if err := r.SetQueryParam("excludeStatus", qExcludeStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
