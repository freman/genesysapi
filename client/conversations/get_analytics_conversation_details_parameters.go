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

// NewGetAnalyticsConversationDetailsParams creates a new GetAnalyticsConversationDetailsParams object
// with the default values initialized.
func NewGetAnalyticsConversationDetailsParams() *GetAnalyticsConversationDetailsParams {
	var ()
	return &GetAnalyticsConversationDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnalyticsConversationDetailsParamsWithTimeout creates a new GetAnalyticsConversationDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAnalyticsConversationDetailsParamsWithTimeout(timeout time.Duration) *GetAnalyticsConversationDetailsParams {
	var ()
	return &GetAnalyticsConversationDetailsParams{

		timeout: timeout,
	}
}

// NewGetAnalyticsConversationDetailsParamsWithContext creates a new GetAnalyticsConversationDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAnalyticsConversationDetailsParamsWithContext(ctx context.Context) *GetAnalyticsConversationDetailsParams {
	var ()
	return &GetAnalyticsConversationDetailsParams{

		Context: ctx,
	}
}

// NewGetAnalyticsConversationDetailsParamsWithHTTPClient creates a new GetAnalyticsConversationDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAnalyticsConversationDetailsParamsWithHTTPClient(client *http.Client) *GetAnalyticsConversationDetailsParams {
	var ()
	return &GetAnalyticsConversationDetailsParams{
		HTTPClient: client,
	}
}

/*GetAnalyticsConversationDetailsParams contains all the parameters to send to the API endpoint
for the get analytics conversation details operation typically these are written to a http.Request
*/
type GetAnalyticsConversationDetailsParams struct {

	/*ConversationID
	  conversationId

	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get analytics conversation details params
func (o *GetAnalyticsConversationDetailsParams) WithTimeout(timeout time.Duration) *GetAnalyticsConversationDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get analytics conversation details params
func (o *GetAnalyticsConversationDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get analytics conversation details params
func (o *GetAnalyticsConversationDetailsParams) WithContext(ctx context.Context) *GetAnalyticsConversationDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get analytics conversation details params
func (o *GetAnalyticsConversationDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get analytics conversation details params
func (o *GetAnalyticsConversationDetailsParams) WithHTTPClient(client *http.Client) *GetAnalyticsConversationDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get analytics conversation details params
func (o *GetAnalyticsConversationDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the get analytics conversation details params
func (o *GetAnalyticsConversationDetailsParams) WithConversationID(conversationID string) *GetAnalyticsConversationDetailsParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get analytics conversation details params
func (o *GetAnalyticsConversationDetailsParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnalyticsConversationDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param conversationId
	if err := r.SetPathParam("conversationId", o.ConversationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
