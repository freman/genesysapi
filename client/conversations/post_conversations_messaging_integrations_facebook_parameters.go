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

// NewPostConversationsMessagingIntegrationsFacebookParams creates a new PostConversationsMessagingIntegrationsFacebookParams object
// with the default values initialized.
func NewPostConversationsMessagingIntegrationsFacebookParams() *PostConversationsMessagingIntegrationsFacebookParams {
	var ()
	return &PostConversationsMessagingIntegrationsFacebookParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationsMessagingIntegrationsFacebookParamsWithTimeout creates a new PostConversationsMessagingIntegrationsFacebookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostConversationsMessagingIntegrationsFacebookParamsWithTimeout(timeout time.Duration) *PostConversationsMessagingIntegrationsFacebookParams {
	var ()
	return &PostConversationsMessagingIntegrationsFacebookParams{

		timeout: timeout,
	}
}

// NewPostConversationsMessagingIntegrationsFacebookParamsWithContext creates a new PostConversationsMessagingIntegrationsFacebookParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostConversationsMessagingIntegrationsFacebookParamsWithContext(ctx context.Context) *PostConversationsMessagingIntegrationsFacebookParams {
	var ()
	return &PostConversationsMessagingIntegrationsFacebookParams{

		Context: ctx,
	}
}

// NewPostConversationsMessagingIntegrationsFacebookParamsWithHTTPClient creates a new PostConversationsMessagingIntegrationsFacebookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostConversationsMessagingIntegrationsFacebookParamsWithHTTPClient(client *http.Client) *PostConversationsMessagingIntegrationsFacebookParams {
	var ()
	return &PostConversationsMessagingIntegrationsFacebookParams{
		HTTPClient: client,
	}
}

/*PostConversationsMessagingIntegrationsFacebookParams contains all the parameters to send to the API endpoint
for the post conversations messaging integrations facebook operation typically these are written to a http.Request
*/
type PostConversationsMessagingIntegrationsFacebookParams struct {

	/*Body
	  FacebookIntegrationRequest

	*/
	Body *models.FacebookIntegrationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post conversations messaging integrations facebook params
func (o *PostConversationsMessagingIntegrationsFacebookParams) WithTimeout(timeout time.Duration) *PostConversationsMessagingIntegrationsFacebookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversations messaging integrations facebook params
func (o *PostConversationsMessagingIntegrationsFacebookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversations messaging integrations facebook params
func (o *PostConversationsMessagingIntegrationsFacebookParams) WithContext(ctx context.Context) *PostConversationsMessagingIntegrationsFacebookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversations messaging integrations facebook params
func (o *PostConversationsMessagingIntegrationsFacebookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversations messaging integrations facebook params
func (o *PostConversationsMessagingIntegrationsFacebookParams) WithHTTPClient(client *http.Client) *PostConversationsMessagingIntegrationsFacebookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversations messaging integrations facebook params
func (o *PostConversationsMessagingIntegrationsFacebookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversations messaging integrations facebook params
func (o *PostConversationsMessagingIntegrationsFacebookParams) WithBody(body *models.FacebookIntegrationRequest) *PostConversationsMessagingIntegrationsFacebookParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversations messaging integrations facebook params
func (o *PostConversationsMessagingIntegrationsFacebookParams) SetBody(body *models.FacebookIntegrationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationsMessagingIntegrationsFacebookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}