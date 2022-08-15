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

// NewPostConversationsKeyconfigurationsParams creates a new PostConversationsKeyconfigurationsParams object
// with the default values initialized.
func NewPostConversationsKeyconfigurationsParams() *PostConversationsKeyconfigurationsParams {
	var ()
	return &PostConversationsKeyconfigurationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationsKeyconfigurationsParamsWithTimeout creates a new PostConversationsKeyconfigurationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostConversationsKeyconfigurationsParamsWithTimeout(timeout time.Duration) *PostConversationsKeyconfigurationsParams {
	var ()
	return &PostConversationsKeyconfigurationsParams{

		timeout: timeout,
	}
}

// NewPostConversationsKeyconfigurationsParamsWithContext creates a new PostConversationsKeyconfigurationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostConversationsKeyconfigurationsParamsWithContext(ctx context.Context) *PostConversationsKeyconfigurationsParams {
	var ()
	return &PostConversationsKeyconfigurationsParams{

		Context: ctx,
	}
}

// NewPostConversationsKeyconfigurationsParamsWithHTTPClient creates a new PostConversationsKeyconfigurationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostConversationsKeyconfigurationsParamsWithHTTPClient(client *http.Client) *PostConversationsKeyconfigurationsParams {
	var ()
	return &PostConversationsKeyconfigurationsParams{
		HTTPClient: client,
	}
}

/*PostConversationsKeyconfigurationsParams contains all the parameters to send to the API endpoint
for the post conversations keyconfigurations operation typically these are written to a http.Request
*/
type PostConversationsKeyconfigurationsParams struct {

	/*Body
	  Encryption Configuration

	*/
	Body *models.ConversationEncryptionConfiguration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post conversations keyconfigurations params
func (o *PostConversationsKeyconfigurationsParams) WithTimeout(timeout time.Duration) *PostConversationsKeyconfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversations keyconfigurations params
func (o *PostConversationsKeyconfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversations keyconfigurations params
func (o *PostConversationsKeyconfigurationsParams) WithContext(ctx context.Context) *PostConversationsKeyconfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversations keyconfigurations params
func (o *PostConversationsKeyconfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversations keyconfigurations params
func (o *PostConversationsKeyconfigurationsParams) WithHTTPClient(client *http.Client) *PostConversationsKeyconfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversations keyconfigurations params
func (o *PostConversationsKeyconfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversations keyconfigurations params
func (o *PostConversationsKeyconfigurationsParams) WithBody(body *models.ConversationEncryptionConfiguration) *PostConversationsKeyconfigurationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversations keyconfigurations params
func (o *PostConversationsKeyconfigurationsParams) SetBody(body *models.ConversationEncryptionConfiguration) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationsKeyconfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
