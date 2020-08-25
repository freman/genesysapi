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

// NewGetVoicemailMeMailboxParams creates a new GetVoicemailMeMailboxParams object
// with the default values initialized.
func NewGetVoicemailMeMailboxParams() *GetVoicemailMeMailboxParams {

	return &GetVoicemailMeMailboxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVoicemailMeMailboxParamsWithTimeout creates a new GetVoicemailMeMailboxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVoicemailMeMailboxParamsWithTimeout(timeout time.Duration) *GetVoicemailMeMailboxParams {

	return &GetVoicemailMeMailboxParams{

		timeout: timeout,
	}
}

// NewGetVoicemailMeMailboxParamsWithContext creates a new GetVoicemailMeMailboxParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVoicemailMeMailboxParamsWithContext(ctx context.Context) *GetVoicemailMeMailboxParams {

	return &GetVoicemailMeMailboxParams{

		Context: ctx,
	}
}

// NewGetVoicemailMeMailboxParamsWithHTTPClient creates a new GetVoicemailMeMailboxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVoicemailMeMailboxParamsWithHTTPClient(client *http.Client) *GetVoicemailMeMailboxParams {

	return &GetVoicemailMeMailboxParams{
		HTTPClient: client,
	}
}

/*GetVoicemailMeMailboxParams contains all the parameters to send to the API endpoint
for the get voicemail me mailbox operation typically these are written to a http.Request
*/
type GetVoicemailMeMailboxParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get voicemail me mailbox params
func (o *GetVoicemailMeMailboxParams) WithTimeout(timeout time.Duration) *GetVoicemailMeMailboxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get voicemail me mailbox params
func (o *GetVoicemailMeMailboxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get voicemail me mailbox params
func (o *GetVoicemailMeMailboxParams) WithContext(ctx context.Context) *GetVoicemailMeMailboxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get voicemail me mailbox params
func (o *GetVoicemailMeMailboxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get voicemail me mailbox params
func (o *GetVoicemailMeMailboxParams) WithHTTPClient(client *http.Client) *GetVoicemailMeMailboxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get voicemail me mailbox params
func (o *GetVoicemailMeMailboxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVoicemailMeMailboxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}