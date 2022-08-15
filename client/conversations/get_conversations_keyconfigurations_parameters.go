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

// NewGetConversationsKeyconfigurationsParams creates a new GetConversationsKeyconfigurationsParams object
// with the default values initialized.
func NewGetConversationsKeyconfigurationsParams() *GetConversationsKeyconfigurationsParams {

	return &GetConversationsKeyconfigurationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsKeyconfigurationsParamsWithTimeout creates a new GetConversationsKeyconfigurationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationsKeyconfigurationsParamsWithTimeout(timeout time.Duration) *GetConversationsKeyconfigurationsParams {

	return &GetConversationsKeyconfigurationsParams{

		timeout: timeout,
	}
}

// NewGetConversationsKeyconfigurationsParamsWithContext creates a new GetConversationsKeyconfigurationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationsKeyconfigurationsParamsWithContext(ctx context.Context) *GetConversationsKeyconfigurationsParams {

	return &GetConversationsKeyconfigurationsParams{

		Context: ctx,
	}
}

// NewGetConversationsKeyconfigurationsParamsWithHTTPClient creates a new GetConversationsKeyconfigurationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationsKeyconfigurationsParamsWithHTTPClient(client *http.Client) *GetConversationsKeyconfigurationsParams {

	return &GetConversationsKeyconfigurationsParams{
		HTTPClient: client,
	}
}

/*GetConversationsKeyconfigurationsParams contains all the parameters to send to the API endpoint
for the get conversations keyconfigurations operation typically these are written to a http.Request
*/
type GetConversationsKeyconfigurationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get conversations keyconfigurations params
func (o *GetConversationsKeyconfigurationsParams) WithTimeout(timeout time.Duration) *GetConversationsKeyconfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations keyconfigurations params
func (o *GetConversationsKeyconfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations keyconfigurations params
func (o *GetConversationsKeyconfigurationsParams) WithContext(ctx context.Context) *GetConversationsKeyconfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations keyconfigurations params
func (o *GetConversationsKeyconfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations keyconfigurations params
func (o *GetConversationsKeyconfigurationsParams) WithHTTPClient(client *http.Client) *GetConversationsKeyconfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations keyconfigurations params
func (o *GetConversationsKeyconfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsKeyconfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
