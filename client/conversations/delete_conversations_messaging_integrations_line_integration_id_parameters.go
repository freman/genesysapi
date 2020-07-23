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

// NewDeleteConversationsMessagingIntegrationsLineIntegrationIDParams creates a new DeleteConversationsMessagingIntegrationsLineIntegrationIDParams object
// with the default values initialized.
func NewDeleteConversationsMessagingIntegrationsLineIntegrationIDParams() *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams {
	var ()
	return &DeleteConversationsMessagingIntegrationsLineIntegrationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConversationsMessagingIntegrationsLineIntegrationIDParamsWithTimeout creates a new DeleteConversationsMessagingIntegrationsLineIntegrationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteConversationsMessagingIntegrationsLineIntegrationIDParamsWithTimeout(timeout time.Duration) *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams {
	var ()
	return &DeleteConversationsMessagingIntegrationsLineIntegrationIDParams{

		timeout: timeout,
	}
}

// NewDeleteConversationsMessagingIntegrationsLineIntegrationIDParamsWithContext creates a new DeleteConversationsMessagingIntegrationsLineIntegrationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteConversationsMessagingIntegrationsLineIntegrationIDParamsWithContext(ctx context.Context) *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams {
	var ()
	return &DeleteConversationsMessagingIntegrationsLineIntegrationIDParams{

		Context: ctx,
	}
}

// NewDeleteConversationsMessagingIntegrationsLineIntegrationIDParamsWithHTTPClient creates a new DeleteConversationsMessagingIntegrationsLineIntegrationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteConversationsMessagingIntegrationsLineIntegrationIDParamsWithHTTPClient(client *http.Client) *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams {
	var ()
	return &DeleteConversationsMessagingIntegrationsLineIntegrationIDParams{
		HTTPClient: client,
	}
}

/*DeleteConversationsMessagingIntegrationsLineIntegrationIDParams contains all the parameters to send to the API endpoint
for the delete conversations messaging integrations line integration Id operation typically these are written to a http.Request
*/
type DeleteConversationsMessagingIntegrationsLineIntegrationIDParams struct {

	/*IntegrationID
	  Integration ID

	*/
	IntegrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete conversations messaging integrations line integration Id params
func (o *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams) WithTimeout(timeout time.Duration) *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete conversations messaging integrations line integration Id params
func (o *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete conversations messaging integrations line integration Id params
func (o *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams) WithContext(ctx context.Context) *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete conversations messaging integrations line integration Id params
func (o *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete conversations messaging integrations line integration Id params
func (o *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams) WithHTTPClient(client *http.Client) *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete conversations messaging integrations line integration Id params
func (o *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntegrationID adds the integrationID to the delete conversations messaging integrations line integration Id params
func (o *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams) WithIntegrationID(integrationID string) *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the delete conversations messaging integrations line integration Id params
func (o *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConversationsMessagingIntegrationsLineIntegrationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
