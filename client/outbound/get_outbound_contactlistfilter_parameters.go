// Code generated by go-swagger; DO NOT EDIT.

package outbound

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

// NewGetOutboundContactlistfilterParams creates a new GetOutboundContactlistfilterParams object
// with the default values initialized.
func NewGetOutboundContactlistfilterParams() *GetOutboundContactlistfilterParams {
	var ()
	return &GetOutboundContactlistfilterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundContactlistfilterParamsWithTimeout creates a new GetOutboundContactlistfilterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutboundContactlistfilterParamsWithTimeout(timeout time.Duration) *GetOutboundContactlistfilterParams {
	var ()
	return &GetOutboundContactlistfilterParams{

		timeout: timeout,
	}
}

// NewGetOutboundContactlistfilterParamsWithContext creates a new GetOutboundContactlistfilterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutboundContactlistfilterParamsWithContext(ctx context.Context) *GetOutboundContactlistfilterParams {
	var ()
	return &GetOutboundContactlistfilterParams{

		Context: ctx,
	}
}

// NewGetOutboundContactlistfilterParamsWithHTTPClient creates a new GetOutboundContactlistfilterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutboundContactlistfilterParamsWithHTTPClient(client *http.Client) *GetOutboundContactlistfilterParams {
	var ()
	return &GetOutboundContactlistfilterParams{
		HTTPClient: client,
	}
}

/*GetOutboundContactlistfilterParams contains all the parameters to send to the API endpoint
for the get outbound contactlistfilter operation typically these are written to a http.Request
*/
type GetOutboundContactlistfilterParams struct {

	/*ContactListFilterID
	  Contact List Filter ID

	*/
	ContactListFilterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get outbound contactlistfilter params
func (o *GetOutboundContactlistfilterParams) WithTimeout(timeout time.Duration) *GetOutboundContactlistfilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound contactlistfilter params
func (o *GetOutboundContactlistfilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound contactlistfilter params
func (o *GetOutboundContactlistfilterParams) WithContext(ctx context.Context) *GetOutboundContactlistfilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound contactlistfilter params
func (o *GetOutboundContactlistfilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound contactlistfilter params
func (o *GetOutboundContactlistfilterParams) WithHTTPClient(client *http.Client) *GetOutboundContactlistfilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound contactlistfilter params
func (o *GetOutboundContactlistfilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactListFilterID adds the contactListFilterID to the get outbound contactlistfilter params
func (o *GetOutboundContactlistfilterParams) WithContactListFilterID(contactListFilterID string) *GetOutboundContactlistfilterParams {
	o.SetContactListFilterID(contactListFilterID)
	return o
}

// SetContactListFilterID adds the contactListFilterId to the get outbound contactlistfilter params
func (o *GetOutboundContactlistfilterParams) SetContactListFilterID(contactListFilterID string) {
	o.ContactListFilterID = contactListFilterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundContactlistfilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contactListFilterId
	if err := r.SetPathParam("contactListFilterId", o.ContactListFilterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
