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

// NewPostConversationsEmailsParams creates a new PostConversationsEmailsParams object
// with the default values initialized.
func NewPostConversationsEmailsParams() *PostConversationsEmailsParams {
	var ()
	return &PostConversationsEmailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationsEmailsParamsWithTimeout creates a new PostConversationsEmailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostConversationsEmailsParamsWithTimeout(timeout time.Duration) *PostConversationsEmailsParams {
	var ()
	return &PostConversationsEmailsParams{

		timeout: timeout,
	}
}

// NewPostConversationsEmailsParamsWithContext creates a new PostConversationsEmailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostConversationsEmailsParamsWithContext(ctx context.Context) *PostConversationsEmailsParams {
	var ()
	return &PostConversationsEmailsParams{

		Context: ctx,
	}
}

// NewPostConversationsEmailsParamsWithHTTPClient creates a new PostConversationsEmailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostConversationsEmailsParamsWithHTTPClient(client *http.Client) *PostConversationsEmailsParams {
	var ()
	return &PostConversationsEmailsParams{
		HTTPClient: client,
	}
}

/*PostConversationsEmailsParams contains all the parameters to send to the API endpoint
for the post conversations emails operation typically these are written to a http.Request
*/
type PostConversationsEmailsParams struct {

	/*Body
	  Create email request

	*/
	Body *models.CreateEmailRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post conversations emails params
func (o *PostConversationsEmailsParams) WithTimeout(timeout time.Duration) *PostConversationsEmailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversations emails params
func (o *PostConversationsEmailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversations emails params
func (o *PostConversationsEmailsParams) WithContext(ctx context.Context) *PostConversationsEmailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversations emails params
func (o *PostConversationsEmailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversations emails params
func (o *PostConversationsEmailsParams) WithHTTPClient(client *http.Client) *PostConversationsEmailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversations emails params
func (o *PostConversationsEmailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversations emails params
func (o *PostConversationsEmailsParams) WithBody(body *models.CreateEmailRequest) *PostConversationsEmailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversations emails params
func (o *PostConversationsEmailsParams) SetBody(body *models.CreateEmailRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationsEmailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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