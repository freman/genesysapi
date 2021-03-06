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

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams creates a new DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams object
// with the default values initialized.
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams() *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams {
	var ()
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDParamsWithTimeout creates a new DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDParamsWithTimeout(timeout time.Duration) *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams {
	var ()
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams{

		timeout: timeout,
	}
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDParamsWithContext creates a new DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDParamsWithContext(ctx context.Context) *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams {
	var ()
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams{

		Context: ctx,
	}
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDParamsWithHTTPClient creates a new DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDParamsWithHTTPClient(client *http.Client) *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams {
	var ()
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams{
		HTTPClient: client,
	}
}

/*DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams contains all the parameters to send to the API endpoint
for the delete conversations messaging integrations twitter integration Id operation typically these are written to a http.Request
*/
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams struct {

	/*IntegrationID
	  Integration ID

	*/
	IntegrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete conversations messaging integrations twitter integration Id params
func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams) WithTimeout(timeout time.Duration) *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete conversations messaging integrations twitter integration Id params
func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete conversations messaging integrations twitter integration Id params
func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams) WithContext(ctx context.Context) *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete conversations messaging integrations twitter integration Id params
func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete conversations messaging integrations twitter integration Id params
func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams) WithHTTPClient(client *http.Client) *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete conversations messaging integrations twitter integration Id params
func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntegrationID adds the integrationID to the delete conversations messaging integrations twitter integration Id params
func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams) WithIntegrationID(integrationID string) *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the delete conversations messaging integrations twitter integration Id params
func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param integrationId
	if err := r.SetPathParam("integrationId", o.IntegrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
