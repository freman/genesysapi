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

// NewGetConversationsMessagingIntegrationsLineIntegrationIDParams creates a new GetConversationsMessagingIntegrationsLineIntegrationIDParams object
// with the default values initialized.
func NewGetConversationsMessagingIntegrationsLineIntegrationIDParams() *GetConversationsMessagingIntegrationsLineIntegrationIDParams {
	var ()
	return &GetConversationsMessagingIntegrationsLineIntegrationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsMessagingIntegrationsLineIntegrationIDParamsWithTimeout creates a new GetConversationsMessagingIntegrationsLineIntegrationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationsMessagingIntegrationsLineIntegrationIDParamsWithTimeout(timeout time.Duration) *GetConversationsMessagingIntegrationsLineIntegrationIDParams {
	var ()
	return &GetConversationsMessagingIntegrationsLineIntegrationIDParams{

		timeout: timeout,
	}
}

// NewGetConversationsMessagingIntegrationsLineIntegrationIDParamsWithContext creates a new GetConversationsMessagingIntegrationsLineIntegrationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationsMessagingIntegrationsLineIntegrationIDParamsWithContext(ctx context.Context) *GetConversationsMessagingIntegrationsLineIntegrationIDParams {
	var ()
	return &GetConversationsMessagingIntegrationsLineIntegrationIDParams{

		Context: ctx,
	}
}

// NewGetConversationsMessagingIntegrationsLineIntegrationIDParamsWithHTTPClient creates a new GetConversationsMessagingIntegrationsLineIntegrationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationsMessagingIntegrationsLineIntegrationIDParamsWithHTTPClient(client *http.Client) *GetConversationsMessagingIntegrationsLineIntegrationIDParams {
	var ()
	return &GetConversationsMessagingIntegrationsLineIntegrationIDParams{
		HTTPClient: client,
	}
}

/*GetConversationsMessagingIntegrationsLineIntegrationIDParams contains all the parameters to send to the API endpoint
for the get conversations messaging integrations line integration Id operation typically these are written to a http.Request
*/
type GetConversationsMessagingIntegrationsLineIntegrationIDParams struct {

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

// WithTimeout adds the timeout to the get conversations messaging integrations line integration Id params
func (o *GetConversationsMessagingIntegrationsLineIntegrationIDParams) WithTimeout(timeout time.Duration) *GetConversationsMessagingIntegrationsLineIntegrationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations messaging integrations line integration Id params
func (o *GetConversationsMessagingIntegrationsLineIntegrationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations messaging integrations line integration Id params
func (o *GetConversationsMessagingIntegrationsLineIntegrationIDParams) WithContext(ctx context.Context) *GetConversationsMessagingIntegrationsLineIntegrationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations messaging integrations line integration Id params
func (o *GetConversationsMessagingIntegrationsLineIntegrationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations messaging integrations line integration Id params
func (o *GetConversationsMessagingIntegrationsLineIntegrationIDParams) WithHTTPClient(client *http.Client) *GetConversationsMessagingIntegrationsLineIntegrationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations messaging integrations line integration Id params
func (o *GetConversationsMessagingIntegrationsLineIntegrationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get conversations messaging integrations line integration Id params
func (o *GetConversationsMessagingIntegrationsLineIntegrationIDParams) WithExpand(expand *string) *GetConversationsMessagingIntegrationsLineIntegrationIDParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get conversations messaging integrations line integration Id params
func (o *GetConversationsMessagingIntegrationsLineIntegrationIDParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithIntegrationID adds the integrationID to the get conversations messaging integrations line integration Id params
func (o *GetConversationsMessagingIntegrationsLineIntegrationIDParams) WithIntegrationID(integrationID string) *GetConversationsMessagingIntegrationsLineIntegrationIDParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the get conversations messaging integrations line integration Id params
func (o *GetConversationsMessagingIntegrationsLineIntegrationIDParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsMessagingIntegrationsLineIntegrationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
