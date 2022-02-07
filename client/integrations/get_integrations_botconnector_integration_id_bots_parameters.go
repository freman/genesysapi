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

// NewGetIntegrationsBotconnectorIntegrationIDBotsParams creates a new GetIntegrationsBotconnectorIntegrationIDBotsParams object
// with the default values initialized.
func NewGetIntegrationsBotconnectorIntegrationIDBotsParams() *GetIntegrationsBotconnectorIntegrationIDBotsParams {
	var ()
	return &GetIntegrationsBotconnectorIntegrationIDBotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsParamsWithTimeout creates a new GetIntegrationsBotconnectorIntegrationIDBotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIntegrationsBotconnectorIntegrationIDBotsParamsWithTimeout(timeout time.Duration) *GetIntegrationsBotconnectorIntegrationIDBotsParams {
	var ()
	return &GetIntegrationsBotconnectorIntegrationIDBotsParams{

		timeout: timeout,
	}
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsParamsWithContext creates a new GetIntegrationsBotconnectorIntegrationIDBotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIntegrationsBotconnectorIntegrationIDBotsParamsWithContext(ctx context.Context) *GetIntegrationsBotconnectorIntegrationIDBotsParams {
	var ()
	return &GetIntegrationsBotconnectorIntegrationIDBotsParams{

		Context: ctx,
	}
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsParamsWithHTTPClient creates a new GetIntegrationsBotconnectorIntegrationIDBotsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIntegrationsBotconnectorIntegrationIDBotsParamsWithHTTPClient(client *http.Client) *GetIntegrationsBotconnectorIntegrationIDBotsParams {
	var ()
	return &GetIntegrationsBotconnectorIntegrationIDBotsParams{
		HTTPClient: client,
	}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsParams contains all the parameters to send to the API endpoint
for the get integrations botconnector integration Id bots operation typically these are written to a http.Request
*/
type GetIntegrationsBotconnectorIntegrationIDBotsParams struct {

	/*IntegrationID
	  The integration ID for this group of bots

	*/
	IntegrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get integrations botconnector integration Id bots params
func (o *GetIntegrationsBotconnectorIntegrationIDBotsParams) WithTimeout(timeout time.Duration) *GetIntegrationsBotconnectorIntegrationIDBotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrations botconnector integration Id bots params
func (o *GetIntegrationsBotconnectorIntegrationIDBotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrations botconnector integration Id bots params
func (o *GetIntegrationsBotconnectorIntegrationIDBotsParams) WithContext(ctx context.Context) *GetIntegrationsBotconnectorIntegrationIDBotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrations botconnector integration Id bots params
func (o *GetIntegrationsBotconnectorIntegrationIDBotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrations botconnector integration Id bots params
func (o *GetIntegrationsBotconnectorIntegrationIDBotsParams) WithHTTPClient(client *http.Client) *GetIntegrationsBotconnectorIntegrationIDBotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrations botconnector integration Id bots params
func (o *GetIntegrationsBotconnectorIntegrationIDBotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntegrationID adds the integrationID to the get integrations botconnector integration Id bots params
func (o *GetIntegrationsBotconnectorIntegrationIDBotsParams) WithIntegrationID(integrationID string) *GetIntegrationsBotconnectorIntegrationIDBotsParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the get integrations botconnector integration Id bots params
func (o *GetIntegrationsBotconnectorIntegrationIDBotsParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationsBotconnectorIntegrationIDBotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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