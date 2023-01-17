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

// NewPostConversationsMessagingSupportedcontentParams creates a new PostConversationsMessagingSupportedcontentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostConversationsMessagingSupportedcontentParams() *PostConversationsMessagingSupportedcontentParams {
	return &PostConversationsMessagingSupportedcontentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationsMessagingSupportedcontentParamsWithTimeout creates a new PostConversationsMessagingSupportedcontentParams object
// with the ability to set a timeout on a request.
func NewPostConversationsMessagingSupportedcontentParamsWithTimeout(timeout time.Duration) *PostConversationsMessagingSupportedcontentParams {
	return &PostConversationsMessagingSupportedcontentParams{
		timeout: timeout,
	}
}

// NewPostConversationsMessagingSupportedcontentParamsWithContext creates a new PostConversationsMessagingSupportedcontentParams object
// with the ability to set a context for a request.
func NewPostConversationsMessagingSupportedcontentParamsWithContext(ctx context.Context) *PostConversationsMessagingSupportedcontentParams {
	return &PostConversationsMessagingSupportedcontentParams{
		Context: ctx,
	}
}

// NewPostConversationsMessagingSupportedcontentParamsWithHTTPClient creates a new PostConversationsMessagingSupportedcontentParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostConversationsMessagingSupportedcontentParamsWithHTTPClient(client *http.Client) *PostConversationsMessagingSupportedcontentParams {
	return &PostConversationsMessagingSupportedcontentParams{
		HTTPClient: client,
	}
}

/*
PostConversationsMessagingSupportedcontentParams contains all the parameters to send to the API endpoint

	for the post conversations messaging supportedcontent operation.

	Typically these are written to a http.Request.
*/
type PostConversationsMessagingSupportedcontentParams struct {

	/* Body.

	   SupportedContent
	*/
	Body *models.SupportedContent

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post conversations messaging supportedcontent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationsMessagingSupportedcontentParams) WithDefaults() *PostConversationsMessagingSupportedcontentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post conversations messaging supportedcontent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationsMessagingSupportedcontentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post conversations messaging supportedcontent params
func (o *PostConversationsMessagingSupportedcontentParams) WithTimeout(timeout time.Duration) *PostConversationsMessagingSupportedcontentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversations messaging supportedcontent params
func (o *PostConversationsMessagingSupportedcontentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversations messaging supportedcontent params
func (o *PostConversationsMessagingSupportedcontentParams) WithContext(ctx context.Context) *PostConversationsMessagingSupportedcontentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversations messaging supportedcontent params
func (o *PostConversationsMessagingSupportedcontentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversations messaging supportedcontent params
func (o *PostConversationsMessagingSupportedcontentParams) WithHTTPClient(client *http.Client) *PostConversationsMessagingSupportedcontentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversations messaging supportedcontent params
func (o *PostConversationsMessagingSupportedcontentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversations messaging supportedcontent params
func (o *PostConversationsMessagingSupportedcontentParams) WithBody(body *models.SupportedContent) *PostConversationsMessagingSupportedcontentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversations messaging supportedcontent params
func (o *PostConversationsMessagingSupportedcontentParams) SetBody(body *models.SupportedContent) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationsMessagingSupportedcontentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
