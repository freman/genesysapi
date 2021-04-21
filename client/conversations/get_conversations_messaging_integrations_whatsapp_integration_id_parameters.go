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

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDParams creates a new GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams object
// with the default values initialized.
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDParams() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams {
	var ()
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDParamsWithTimeout creates a new GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDParamsWithTimeout(timeout time.Duration) *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams {
	var ()
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams{

		timeout: timeout,
	}
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDParamsWithContext creates a new GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDParamsWithContext(ctx context.Context) *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams {
	var ()
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams{

		Context: ctx,
	}
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDParamsWithHTTPClient creates a new GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDParamsWithHTTPClient(client *http.Client) *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams {
	var ()
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams{
		HTTPClient: client,
	}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams contains all the parameters to send to the API endpoint
for the get conversations messaging integrations whatsapp integration Id operation typically these are written to a http.Request
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams struct {

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

// WithTimeout adds the timeout to the get conversations messaging integrations whatsapp integration Id params
func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams) WithTimeout(timeout time.Duration) *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations messaging integrations whatsapp integration Id params
func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations messaging integrations whatsapp integration Id params
func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams) WithContext(ctx context.Context) *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations messaging integrations whatsapp integration Id params
func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations messaging integrations whatsapp integration Id params
func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams) WithHTTPClient(client *http.Client) *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations messaging integrations whatsapp integration Id params
func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get conversations messaging integrations whatsapp integration Id params
func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams) WithExpand(expand *string) *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get conversations messaging integrations whatsapp integration Id params
func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithIntegrationID adds the integrationID to the get conversations messaging integrations whatsapp integration Id params
func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams) WithIntegrationID(integrationID string) *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the get conversations messaging integrations whatsapp integration Id params
func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
