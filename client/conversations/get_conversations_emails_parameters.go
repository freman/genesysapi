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

// NewGetConversationsEmailsParams creates a new GetConversationsEmailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConversationsEmailsParams() *GetConversationsEmailsParams {
	return &GetConversationsEmailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsEmailsParamsWithTimeout creates a new GetConversationsEmailsParams object
// with the ability to set a timeout on a request.
func NewGetConversationsEmailsParamsWithTimeout(timeout time.Duration) *GetConversationsEmailsParams {
	return &GetConversationsEmailsParams{
		timeout: timeout,
	}
}

// NewGetConversationsEmailsParamsWithContext creates a new GetConversationsEmailsParams object
// with the ability to set a context for a request.
func NewGetConversationsEmailsParamsWithContext(ctx context.Context) *GetConversationsEmailsParams {
	return &GetConversationsEmailsParams{
		Context: ctx,
	}
}

// NewGetConversationsEmailsParamsWithHTTPClient creates a new GetConversationsEmailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConversationsEmailsParamsWithHTTPClient(client *http.Client) *GetConversationsEmailsParams {
	return &GetConversationsEmailsParams{
		HTTPClient: client,
	}
}

/*
GetConversationsEmailsParams contains all the parameters to send to the API endpoint

	for the get conversations emails operation.

	Typically these are written to a http.Request.
*/
type GetConversationsEmailsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get conversations emails params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsEmailsParams) WithDefaults() *GetConversationsEmailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get conversations emails params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsEmailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get conversations emails params
func (o *GetConversationsEmailsParams) WithTimeout(timeout time.Duration) *GetConversationsEmailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations emails params
func (o *GetConversationsEmailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations emails params
func (o *GetConversationsEmailsParams) WithContext(ctx context.Context) *GetConversationsEmailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations emails params
func (o *GetConversationsEmailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations emails params
func (o *GetConversationsEmailsParams) WithHTTPClient(client *http.Client) *GetConversationsEmailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations emails params
func (o *GetConversationsEmailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsEmailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
