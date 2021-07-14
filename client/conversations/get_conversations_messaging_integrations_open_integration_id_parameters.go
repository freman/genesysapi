// Code generated by go-swagger; DO NOT EDIT.

package conversations

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

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDParams creates a new GetConversationsMessagingIntegrationsOpenIntegrationIDParams object
// with the default values initialized.
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDParams() *GetConversationsMessagingIntegrationsOpenIntegrationIDParams {
	var ()
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDParamsWithTimeout creates a new GetConversationsMessagingIntegrationsOpenIntegrationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDParamsWithTimeout(timeout time.Duration) *GetConversationsMessagingIntegrationsOpenIntegrationIDParams {
	var ()
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDParams{

		timeout: timeout,
	}
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDParamsWithContext creates a new GetConversationsMessagingIntegrationsOpenIntegrationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDParamsWithContext(ctx context.Context) *GetConversationsMessagingIntegrationsOpenIntegrationIDParams {
	var ()
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDParams{

		Context: ctx,
	}
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDParamsWithHTTPClient creates a new GetConversationsMessagingIntegrationsOpenIntegrationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDParamsWithHTTPClient(client *http.Client) *GetConversationsMessagingIntegrationsOpenIntegrationIDParams {
	var ()
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDParams{
		HTTPClient: client,
	}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDParams contains all the parameters to send to the API endpoint
for the get conversations messaging integrations open integration Id operation typically these are written to a http.Request
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDParams struct {

	/*Expand
	  Expand instructions for the return value.

	*/
	Expand *string
	/*IntegrationID
	  Integration ID

	*/
	IntegrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get conversations messaging integrations open integration Id params
func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDParams) WithTimeout(timeout time.Duration) *GetConversationsMessagingIntegrationsOpenIntegrationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations messaging integrations open integration Id params
func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations messaging integrations open integration Id params
func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDParams) WithContext(ctx context.Context) *GetConversationsMessagingIntegrationsOpenIntegrationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations messaging integrations open integration Id params
func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations messaging integrations open integration Id params
func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDParams) WithHTTPClient(client *http.Client) *GetConversationsMessagingIntegrationsOpenIntegrationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations messaging integrations open integration Id params
func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get conversations messaging integrations open integration Id params
func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDParams) WithExpand(expand *string) *GetConversationsMessagingIntegrationsOpenIntegrationIDParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get conversations messaging integrations open integration Id params
func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithIntegrationID adds the integrationID to the get conversations messaging integrations open integration Id params
func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDParams) WithIntegrationID(integrationID string) *GetConversationsMessagingIntegrationsOpenIntegrationIDParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the get conversations messaging integrations open integration Id params
func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// query param expand
		var qrExpand string
		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {
			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}

	}

	// path param integrationId
	if err := r.SetPathParam("integrationId", o.IntegrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}