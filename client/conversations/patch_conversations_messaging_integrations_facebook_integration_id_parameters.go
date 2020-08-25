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

	"github.com/freman/genesysapi/models"
)

// NewPatchConversationsMessagingIntegrationsFacebookIntegrationIDParams creates a new PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams object
// with the default values initialized.
func NewPatchConversationsMessagingIntegrationsFacebookIntegrationIDParams() *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams {
	var ()
	return &PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConversationsMessagingIntegrationsFacebookIntegrationIDParamsWithTimeout creates a new PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchConversationsMessagingIntegrationsFacebookIntegrationIDParamsWithTimeout(timeout time.Duration) *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams {
	var ()
	return &PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams{

		timeout: timeout,
	}
}

// NewPatchConversationsMessagingIntegrationsFacebookIntegrationIDParamsWithContext creates a new PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchConversationsMessagingIntegrationsFacebookIntegrationIDParamsWithContext(ctx context.Context) *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams {
	var ()
	return &PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams{

		Context: ctx,
	}
}

// NewPatchConversationsMessagingIntegrationsFacebookIntegrationIDParamsWithHTTPClient creates a new PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchConversationsMessagingIntegrationsFacebookIntegrationIDParamsWithHTTPClient(client *http.Client) *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams {
	var ()
	return &PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams{
		HTTPClient: client,
	}
}

/*PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams contains all the parameters to send to the API endpoint
for the patch conversations messaging integrations facebook integration Id operation typically these are written to a http.Request
*/
type PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams struct {

	/*Body
	  FacebookIntegrationUpdateRequest

	*/
	Body *models.FacebookIntegrationUpdateRequest
	/*IntegrationID
	  Integration ID

	*/
	IntegrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch conversations messaging integrations facebook integration Id params
func (o *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams) WithTimeout(timeout time.Duration) *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch conversations messaging integrations facebook integration Id params
func (o *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch conversations messaging integrations facebook integration Id params
func (o *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams) WithContext(ctx context.Context) *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch conversations messaging integrations facebook integration Id params
func (o *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch conversations messaging integrations facebook integration Id params
func (o *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams) WithHTTPClient(client *http.Client) *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch conversations messaging integrations facebook integration Id params
func (o *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch conversations messaging integrations facebook integration Id params
func (o *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams) WithBody(body *models.FacebookIntegrationUpdateRequest) *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch conversations messaging integrations facebook integration Id params
func (o *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams) SetBody(body *models.FacebookIntegrationUpdateRequest) {
	o.Body = body
}

// WithIntegrationID adds the integrationID to the patch conversations messaging integrations facebook integration Id params
func (o *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams) WithIntegrationID(integrationID string) *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the patch conversations messaging integrations facebook integration Id params
func (o *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConversationsMessagingIntegrationsFacebookIntegrationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
