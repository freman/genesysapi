// Code generated by go-swagger; DO NOT EDIT.

package voicemail

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

// NewGetVoicemailMailboxParams creates a new GetVoicemailMailboxParams object
// with the default values initialized.
func NewGetVoicemailMailboxParams() *GetVoicemailMailboxParams {

	return &GetVoicemailMailboxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVoicemailMailboxParamsWithTimeout creates a new GetVoicemailMailboxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVoicemailMailboxParamsWithTimeout(timeout time.Duration) *GetVoicemailMailboxParams {

	return &GetVoicemailMailboxParams{

		timeout: timeout,
	}
}

// NewGetVoicemailMailboxParamsWithContext creates a new GetVoicemailMailboxParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVoicemailMailboxParamsWithContext(ctx context.Context) *GetVoicemailMailboxParams {

	return &GetVoicemailMailboxParams{

		Context: ctx,
	}
}

// NewGetVoicemailMailboxParamsWithHTTPClient creates a new GetVoicemailMailboxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVoicemailMailboxParamsWithHTTPClient(client *http.Client) *GetVoicemailMailboxParams {

	return &GetVoicemailMailboxParams{
		HTTPClient: client,
	}
}

/*GetVoicemailMailboxParams contains all the parameters to send to the API endpoint
for the get voicemail mailbox operation typically these are written to a http.Request
*/
type GetVoicemailMailboxParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get voicemail mailbox params
func (o *GetVoicemailMailboxParams) WithTimeout(timeout time.Duration) *GetVoicemailMailboxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get voicemail mailbox params
func (o *GetVoicemailMailboxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get voicemail mailbox params
func (o *GetVoicemailMailboxParams) WithContext(ctx context.Context) *GetVoicemailMailboxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get voicemail mailbox params
func (o *GetVoicemailMailboxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get voicemail mailbox params
func (o *GetVoicemailMailboxParams) WithHTTPClient(client *http.Client) *GetVoicemailMailboxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get voicemail mailbox params
func (o *GetVoicemailMailboxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVoicemailMailboxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
