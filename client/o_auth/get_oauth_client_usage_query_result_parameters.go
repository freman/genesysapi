// Code generated by go-swagger; DO NOT EDIT.

package o_auth

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

// NewGetOauthClientUsageQueryResultParams creates a new GetOauthClientUsageQueryResultParams object
// with the default values initialized.
func NewGetOauthClientUsageQueryResultParams() *GetOauthClientUsageQueryResultParams {
	var ()
	return &GetOauthClientUsageQueryResultParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOauthClientUsageQueryResultParamsWithTimeout creates a new GetOauthClientUsageQueryResultParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOauthClientUsageQueryResultParamsWithTimeout(timeout time.Duration) *GetOauthClientUsageQueryResultParams {
	var ()
	return &GetOauthClientUsageQueryResultParams{

		timeout: timeout,
	}
}

// NewGetOauthClientUsageQueryResultParamsWithContext creates a new GetOauthClientUsageQueryResultParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOauthClientUsageQueryResultParamsWithContext(ctx context.Context) *GetOauthClientUsageQueryResultParams {
	var ()
	return &GetOauthClientUsageQueryResultParams{

		Context: ctx,
	}
}

// NewGetOauthClientUsageQueryResultParamsWithHTTPClient creates a new GetOauthClientUsageQueryResultParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOauthClientUsageQueryResultParamsWithHTTPClient(client *http.Client) *GetOauthClientUsageQueryResultParams {
	var ()
	return &GetOauthClientUsageQueryResultParams{
		HTTPClient: client,
	}
}

/*GetOauthClientUsageQueryResultParams contains all the parameters to send to the API endpoint
for the get oauth client usage query result operation typically these are written to a http.Request
*/
type GetOauthClientUsageQueryResultParams struct {

	/*ClientID
	  Client ID

	*/
	ClientID string
	/*ExecutionID
	  ID of the query execution

	*/
	ExecutionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get oauth client usage query result params
func (o *GetOauthClientUsageQueryResultParams) WithTimeout(timeout time.Duration) *GetOauthClientUsageQueryResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oauth client usage query result params
func (o *GetOauthClientUsageQueryResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oauth client usage query result params
func (o *GetOauthClientUsageQueryResultParams) WithContext(ctx context.Context) *GetOauthClientUsageQueryResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oauth client usage query result params
func (o *GetOauthClientUsageQueryResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get oauth client usage query result params
func (o *GetOauthClientUsageQueryResultParams) WithHTTPClient(client *http.Client) *GetOauthClientUsageQueryResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get oauth client usage query result params
func (o *GetOauthClientUsageQueryResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the get oauth client usage query result params
func (o *GetOauthClientUsageQueryResultParams) WithClientID(clientID string) *GetOauthClientUsageQueryResultParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the get oauth client usage query result params
func (o *GetOauthClientUsageQueryResultParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithExecutionID adds the executionID to the get oauth client usage query result params
func (o *GetOauthClientUsageQueryResultParams) WithExecutionID(executionID string) *GetOauthClientUsageQueryResultParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the get oauth client usage query result params
func (o *GetOauthClientUsageQueryResultParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOauthClientUsageQueryResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clientId
	if err := r.SetPathParam("clientId", o.ClientID); err != nil {
		return err
	}

	// path param executionId
	if err := r.SetPathParam("executionId", o.ExecutionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
