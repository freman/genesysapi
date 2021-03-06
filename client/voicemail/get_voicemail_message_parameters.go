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
	"github.com/go-openapi/swag"
)

// NewGetVoicemailMessageParams creates a new GetVoicemailMessageParams object
// with the default values initialized.
func NewGetVoicemailMessageParams() *GetVoicemailMessageParams {
	var ()
	return &GetVoicemailMessageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVoicemailMessageParamsWithTimeout creates a new GetVoicemailMessageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVoicemailMessageParamsWithTimeout(timeout time.Duration) *GetVoicemailMessageParams {
	var ()
	return &GetVoicemailMessageParams{

		timeout: timeout,
	}
}

// NewGetVoicemailMessageParamsWithContext creates a new GetVoicemailMessageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVoicemailMessageParamsWithContext(ctx context.Context) *GetVoicemailMessageParams {
	var ()
	return &GetVoicemailMessageParams{

		Context: ctx,
	}
}

// NewGetVoicemailMessageParamsWithHTTPClient creates a new GetVoicemailMessageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVoicemailMessageParamsWithHTTPClient(client *http.Client) *GetVoicemailMessageParams {
	var ()
	return &GetVoicemailMessageParams{
		HTTPClient: client,
	}
}

/*GetVoicemailMessageParams contains all the parameters to send to the API endpoint
for the get voicemail message operation typically these are written to a http.Request
*/
type GetVoicemailMessageParams struct {

	/*Expand
	  If the caller is a known user, which fields, if any, to expand

	*/
	Expand []string
	/*MessageID
	  Message ID

	*/
	MessageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get voicemail message params
func (o *GetVoicemailMessageParams) WithTimeout(timeout time.Duration) *GetVoicemailMessageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get voicemail message params
func (o *GetVoicemailMessageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get voicemail message params
func (o *GetVoicemailMessageParams) WithContext(ctx context.Context) *GetVoicemailMessageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get voicemail message params
func (o *GetVoicemailMessageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get voicemail message params
func (o *GetVoicemailMessageParams) WithHTTPClient(client *http.Client) *GetVoicemailMessageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get voicemail message params
func (o *GetVoicemailMessageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get voicemail message params
func (o *GetVoicemailMessageParams) WithExpand(expand []string) *GetVoicemailMessageParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get voicemail message params
func (o *GetVoicemailMessageParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithMessageID adds the messageID to the get voicemail message params
func (o *GetVoicemailMessageParams) WithMessageID(messageID string) *GetVoicemailMessageParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the get voicemail message params
func (o *GetVoicemailMessageParams) SetMessageID(messageID string) {
	o.MessageID = messageID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVoicemailMessageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	// path param messageId
	if err := r.SetPathParam("messageId", o.MessageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
