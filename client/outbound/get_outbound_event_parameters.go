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

// NewGetOutboundEventParams creates a new GetOutboundEventParams object
// with the default values initialized.
func NewGetOutboundEventParams() *GetOutboundEventParams {
	var ()
	return &GetOutboundEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundEventParamsWithTimeout creates a new GetOutboundEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutboundEventParamsWithTimeout(timeout time.Duration) *GetOutboundEventParams {
	var ()
	return &GetOutboundEventParams{

		timeout: timeout,
	}
}

// NewGetOutboundEventParamsWithContext creates a new GetOutboundEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutboundEventParamsWithContext(ctx context.Context) *GetOutboundEventParams {
	var ()
	return &GetOutboundEventParams{

		Context: ctx,
	}
}

// NewGetOutboundEventParamsWithHTTPClient creates a new GetOutboundEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutboundEventParamsWithHTTPClient(client *http.Client) *GetOutboundEventParams {
	var ()
	return &GetOutboundEventParams{
		HTTPClient: client,
	}
}

/*GetOutboundEventParams contains all the parameters to send to the API endpoint
for the get outbound event operation typically these are written to a http.Request
*/
type GetOutboundEventParams struct {

	/*EventID
	  Event Log ID

	*/
	EventID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get outbound event params
func (o *GetOutboundEventParams) WithTimeout(timeout time.Duration) *GetOutboundEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound event params
func (o *GetOutboundEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound event params
func (o *GetOutboundEventParams) WithContext(ctx context.Context) *GetOutboundEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound event params
func (o *GetOutboundEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound event params
func (o *GetOutboundEventParams) WithHTTPClient(client *http.Client) *GetOutboundEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound event params
func (o *GetOutboundEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEventID adds the eventID to the get outbound event params
func (o *GetOutboundEventParams) WithEventID(eventID string) *GetOutboundEventParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the get outbound event params
func (o *GetOutboundEventParams) SetEventID(eventID string) {
	o.EventID = eventID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param eventId
	if err := r.SetPathParam("eventId", o.EventID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}