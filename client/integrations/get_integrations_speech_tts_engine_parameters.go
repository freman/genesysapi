// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetIntegrationsSpeechTtsEngineParams creates a new GetIntegrationsSpeechTtsEngineParams object
// with the default values initialized.
func NewGetIntegrationsSpeechTtsEngineParams() *GetIntegrationsSpeechTtsEngineParams {
	var (
		includeVoicesDefault = bool(false)
	)
	return &GetIntegrationsSpeechTtsEngineParams{
		IncludeVoices: &includeVoicesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationsSpeechTtsEngineParamsWithTimeout creates a new GetIntegrationsSpeechTtsEngineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIntegrationsSpeechTtsEngineParamsWithTimeout(timeout time.Duration) *GetIntegrationsSpeechTtsEngineParams {
	var (
		includeVoicesDefault = bool(false)
	)
	return &GetIntegrationsSpeechTtsEngineParams{
		IncludeVoices: &includeVoicesDefault,

		timeout: timeout,
	}
}

// NewGetIntegrationsSpeechTtsEngineParamsWithContext creates a new GetIntegrationsSpeechTtsEngineParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIntegrationsSpeechTtsEngineParamsWithContext(ctx context.Context) *GetIntegrationsSpeechTtsEngineParams {
	var (
		includeVoicesDefault = bool(false)
	)
	return &GetIntegrationsSpeechTtsEngineParams{
		IncludeVoices: &includeVoicesDefault,

		Context: ctx,
	}
}

// NewGetIntegrationsSpeechTtsEngineParamsWithHTTPClient creates a new GetIntegrationsSpeechTtsEngineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIntegrationsSpeechTtsEngineParamsWithHTTPClient(client *http.Client) *GetIntegrationsSpeechTtsEngineParams {
	var (
		includeVoicesDefault = bool(false)
	)
	return &GetIntegrationsSpeechTtsEngineParams{
		IncludeVoices: &includeVoicesDefault,
		HTTPClient:    client,
	}
}

/*GetIntegrationsSpeechTtsEngineParams contains all the parameters to send to the API endpoint
for the get integrations speech tts engine operation typically these are written to a http.Request
*/
type GetIntegrationsSpeechTtsEngineParams struct {

	/*EngineID
	  The engine ID

	*/
	EngineID string
	/*IncludeVoices
	  Include voices for the engine

	*/
	IncludeVoices *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get integrations speech tts engine params
func (o *GetIntegrationsSpeechTtsEngineParams) WithTimeout(timeout time.Duration) *GetIntegrationsSpeechTtsEngineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrations speech tts engine params
func (o *GetIntegrationsSpeechTtsEngineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrations speech tts engine params
func (o *GetIntegrationsSpeechTtsEngineParams) WithContext(ctx context.Context) *GetIntegrationsSpeechTtsEngineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrations speech tts engine params
func (o *GetIntegrationsSpeechTtsEngineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrations speech tts engine params
func (o *GetIntegrationsSpeechTtsEngineParams) WithHTTPClient(client *http.Client) *GetIntegrationsSpeechTtsEngineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrations speech tts engine params
func (o *GetIntegrationsSpeechTtsEngineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEngineID adds the engineID to the get integrations speech tts engine params
func (o *GetIntegrationsSpeechTtsEngineParams) WithEngineID(engineID string) *GetIntegrationsSpeechTtsEngineParams {
	o.SetEngineID(engineID)
	return o
}

// SetEngineID adds the engineId to the get integrations speech tts engine params
func (o *GetIntegrationsSpeechTtsEngineParams) SetEngineID(engineID string) {
	o.EngineID = engineID
}

// WithIncludeVoices adds the includeVoices to the get integrations speech tts engine params
func (o *GetIntegrationsSpeechTtsEngineParams) WithIncludeVoices(includeVoices *bool) *GetIntegrationsSpeechTtsEngineParams {
	o.SetIncludeVoices(includeVoices)
	return o
}

// SetIncludeVoices adds the includeVoices to the get integrations speech tts engine params
func (o *GetIntegrationsSpeechTtsEngineParams) SetIncludeVoices(includeVoices *bool) {
	o.IncludeVoices = includeVoices
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationsSpeechTtsEngineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param engineId
	if err := r.SetPathParam("engineId", o.EngineID); err != nil {
		return err
	}

	if o.IncludeVoices != nil {

		// query param includeVoices
		var qrIncludeVoices bool
		if o.IncludeVoices != nil {
			qrIncludeVoices = *o.IncludeVoices
		}
		qIncludeVoices := swag.FormatBool(qrIncludeVoices)
		if qIncludeVoices != "" {
			if err := r.SetQueryParam("includeVoices", qIncludeVoices); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
