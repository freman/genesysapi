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

// NewPostConversationsCallbacksBulkDisconnectParams creates a new PostConversationsCallbacksBulkDisconnectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostConversationsCallbacksBulkDisconnectParams() *PostConversationsCallbacksBulkDisconnectParams {
	return &PostConversationsCallbacksBulkDisconnectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostConversationsCallbacksBulkDisconnectParamsWithTimeout creates a new PostConversationsCallbacksBulkDisconnectParams object
// with the ability to set a timeout on a request.
func NewPostConversationsCallbacksBulkDisconnectParamsWithTimeout(timeout time.Duration) *PostConversationsCallbacksBulkDisconnectParams {
	return &PostConversationsCallbacksBulkDisconnectParams{
		timeout: timeout,
	}
}

// NewPostConversationsCallbacksBulkDisconnectParamsWithContext creates a new PostConversationsCallbacksBulkDisconnectParams object
// with the ability to set a context for a request.
func NewPostConversationsCallbacksBulkDisconnectParamsWithContext(ctx context.Context) *PostConversationsCallbacksBulkDisconnectParams {
	return &PostConversationsCallbacksBulkDisconnectParams{
		Context: ctx,
	}
}

// NewPostConversationsCallbacksBulkDisconnectParamsWithHTTPClient creates a new PostConversationsCallbacksBulkDisconnectParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostConversationsCallbacksBulkDisconnectParamsWithHTTPClient(client *http.Client) *PostConversationsCallbacksBulkDisconnectParams {
	return &PostConversationsCallbacksBulkDisconnectParams{
		HTTPClient: client,
	}
}

/*
PostConversationsCallbacksBulkDisconnectParams contains all the parameters to send to the API endpoint

	for the post conversations callbacks bulk disconnect operation.

	Typically these are written to a http.Request.
*/
type PostConversationsCallbacksBulkDisconnectParams struct {

	/* Body.

	   BulkCallbackDisconnectRequest
	*/
	Body *models.BulkCallbackDisconnectRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post conversations callbacks bulk disconnect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationsCallbacksBulkDisconnectParams) WithDefaults() *PostConversationsCallbacksBulkDisconnectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post conversations callbacks bulk disconnect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConversationsCallbacksBulkDisconnectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post conversations callbacks bulk disconnect params
func (o *PostConversationsCallbacksBulkDisconnectParams) WithTimeout(timeout time.Duration) *PostConversationsCallbacksBulkDisconnectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post conversations callbacks bulk disconnect params
func (o *PostConversationsCallbacksBulkDisconnectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post conversations callbacks bulk disconnect params
func (o *PostConversationsCallbacksBulkDisconnectParams) WithContext(ctx context.Context) *PostConversationsCallbacksBulkDisconnectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post conversations callbacks bulk disconnect params
func (o *PostConversationsCallbacksBulkDisconnectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post conversations callbacks bulk disconnect params
func (o *PostConversationsCallbacksBulkDisconnectParams) WithHTTPClient(client *http.Client) *PostConversationsCallbacksBulkDisconnectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post conversations callbacks bulk disconnect params
func (o *PostConversationsCallbacksBulkDisconnectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post conversations callbacks bulk disconnect params
func (o *PostConversationsCallbacksBulkDisconnectParams) WithBody(body *models.BulkCallbackDisconnectRequest) *PostConversationsCallbacksBulkDisconnectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post conversations callbacks bulk disconnect params
func (o *PostConversationsCallbacksBulkDisconnectParams) SetBody(body *models.BulkCallbackDisconnectRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostConversationsCallbacksBulkDisconnectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
