// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetIntegrationsEventlogEventIDParams creates a new GetIntegrationsEventlogEventIDParams object
// with the default values initialized.
func NewGetIntegrationsEventlogEventIDParams() *GetIntegrationsEventlogEventIDParams {
	var ()
	return &GetIntegrationsEventlogEventIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationsEventlogEventIDParamsWithTimeout creates a new GetIntegrationsEventlogEventIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIntegrationsEventlogEventIDParamsWithTimeout(timeout time.Duration) *GetIntegrationsEventlogEventIDParams {
	var ()
	return &GetIntegrationsEventlogEventIDParams{

		timeout: timeout,
	}
}

// NewGetIntegrationsEventlogEventIDParamsWithContext creates a new GetIntegrationsEventlogEventIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIntegrationsEventlogEventIDParamsWithContext(ctx context.Context) *GetIntegrationsEventlogEventIDParams {
	var ()
	return &GetIntegrationsEventlogEventIDParams{

		Context: ctx,
	}
}

// NewGetIntegrationsEventlogEventIDParamsWithHTTPClient creates a new GetIntegrationsEventlogEventIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIntegrationsEventlogEventIDParamsWithHTTPClient(client *http.Client) *GetIntegrationsEventlogEventIDParams {
	var ()
	return &GetIntegrationsEventlogEventIDParams{
		HTTPClient: client,
	}
}

/*GetIntegrationsEventlogEventIDParams contains all the parameters to send to the API endpoint
for the get integrations eventlog event Id operation typically these are written to a http.Request
*/
type GetIntegrationsEventlogEventIDParams struct {

	/*EventID
	  Event Id

	*/
	EventID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get integrations eventlog event Id params
func (o *GetIntegrationsEventlogEventIDParams) WithTimeout(timeout time.Duration) *GetIntegrationsEventlogEventIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrations eventlog event Id params
func (o *GetIntegrationsEventlogEventIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrations eventlog event Id params
func (o *GetIntegrationsEventlogEventIDParams) WithContext(ctx context.Context) *GetIntegrationsEventlogEventIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrations eventlog event Id params
func (o *GetIntegrationsEventlogEventIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrations eventlog event Id params
func (o *GetIntegrationsEventlogEventIDParams) WithHTTPClient(client *http.Client) *GetIntegrationsEventlogEventIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrations eventlog event Id params
func (o *GetIntegrationsEventlogEventIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEventID adds the eventID to the get integrations eventlog event Id params
func (o *GetIntegrationsEventlogEventIDParams) WithEventID(eventID string) *GetIntegrationsEventlogEventIDParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the get integrations eventlog event Id params
func (o *GetIntegrationsEventlogEventIDParams) SetEventID(eventID string) {
	o.EventID = eventID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationsEventlogEventIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
