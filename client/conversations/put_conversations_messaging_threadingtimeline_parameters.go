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

// NewPutConversationsMessagingThreadingtimelineParams creates a new PutConversationsMessagingThreadingtimelineParams object
// with the default values initialized.
func NewPutConversationsMessagingThreadingtimelineParams() *PutConversationsMessagingThreadingtimelineParams {
	var ()
	return &PutConversationsMessagingThreadingtimelineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutConversationsMessagingThreadingtimelineParamsWithTimeout creates a new PutConversationsMessagingThreadingtimelineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutConversationsMessagingThreadingtimelineParamsWithTimeout(timeout time.Duration) *PutConversationsMessagingThreadingtimelineParams {
	var ()
	return &PutConversationsMessagingThreadingtimelineParams{

		timeout: timeout,
	}
}

// NewPutConversationsMessagingThreadingtimelineParamsWithContext creates a new PutConversationsMessagingThreadingtimelineParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutConversationsMessagingThreadingtimelineParamsWithContext(ctx context.Context) *PutConversationsMessagingThreadingtimelineParams {
	var ()
	return &PutConversationsMessagingThreadingtimelineParams{

		Context: ctx,
	}
}

// NewPutConversationsMessagingThreadingtimelineParamsWithHTTPClient creates a new PutConversationsMessagingThreadingtimelineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutConversationsMessagingThreadingtimelineParamsWithHTTPClient(client *http.Client) *PutConversationsMessagingThreadingtimelineParams {
	var ()
	return &PutConversationsMessagingThreadingtimelineParams{
		HTTPClient: client,
	}
}

/*PutConversationsMessagingThreadingtimelineParams contains all the parameters to send to the API endpoint
for the put conversations messaging threadingtimeline operation typically these are written to a http.Request
*/
type PutConversationsMessagingThreadingtimelineParams struct {

	/*Body
	  ConversationThreadingWindowRequest

	*/
	Body *models.ConversationThreadingWindow

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put conversations messaging threadingtimeline params
func (o *PutConversationsMessagingThreadingtimelineParams) WithTimeout(timeout time.Duration) *PutConversationsMessagingThreadingtimelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put conversations messaging threadingtimeline params
func (o *PutConversationsMessagingThreadingtimelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put conversations messaging threadingtimeline params
func (o *PutConversationsMessagingThreadingtimelineParams) WithContext(ctx context.Context) *PutConversationsMessagingThreadingtimelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put conversations messaging threadingtimeline params
func (o *PutConversationsMessagingThreadingtimelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put conversations messaging threadingtimeline params
func (o *PutConversationsMessagingThreadingtimelineParams) WithHTTPClient(client *http.Client) *PutConversationsMessagingThreadingtimelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put conversations messaging threadingtimeline params
func (o *PutConversationsMessagingThreadingtimelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put conversations messaging threadingtimeline params
func (o *PutConversationsMessagingThreadingtimelineParams) WithBody(body *models.ConversationThreadingWindow) *PutConversationsMessagingThreadingtimelineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put conversations messaging threadingtimeline params
func (o *PutConversationsMessagingThreadingtimelineParams) SetBody(body *models.ConversationThreadingWindow) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutConversationsMessagingThreadingtimelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
