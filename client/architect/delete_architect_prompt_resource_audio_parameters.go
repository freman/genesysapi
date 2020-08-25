// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewDeleteArchitectPromptResourceAudioParams creates a new DeleteArchitectPromptResourceAudioParams object
// with the default values initialized.
func NewDeleteArchitectPromptResourceAudioParams() *DeleteArchitectPromptResourceAudioParams {
	var ()
	return &DeleteArchitectPromptResourceAudioParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteArchitectPromptResourceAudioParamsWithTimeout creates a new DeleteArchitectPromptResourceAudioParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteArchitectPromptResourceAudioParamsWithTimeout(timeout time.Duration) *DeleteArchitectPromptResourceAudioParams {
	var ()
	return &DeleteArchitectPromptResourceAudioParams{

		timeout: timeout,
	}
}

// NewDeleteArchitectPromptResourceAudioParamsWithContext creates a new DeleteArchitectPromptResourceAudioParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteArchitectPromptResourceAudioParamsWithContext(ctx context.Context) *DeleteArchitectPromptResourceAudioParams {
	var ()
	return &DeleteArchitectPromptResourceAudioParams{

		Context: ctx,
	}
}

// NewDeleteArchitectPromptResourceAudioParamsWithHTTPClient creates a new DeleteArchitectPromptResourceAudioParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteArchitectPromptResourceAudioParamsWithHTTPClient(client *http.Client) *DeleteArchitectPromptResourceAudioParams {
	var ()
	return &DeleteArchitectPromptResourceAudioParams{
		HTTPClient: client,
	}
}

/*DeleteArchitectPromptResourceAudioParams contains all the parameters to send to the API endpoint
for the delete architect prompt resource audio operation typically these are written to a http.Request
*/
type DeleteArchitectPromptResourceAudioParams struct {

	/*LanguageCode
	  Language

	*/
	LanguageCode string
	/*PromptID
	  Prompt ID

	*/
	PromptID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete architect prompt resource audio params
func (o *DeleteArchitectPromptResourceAudioParams) WithTimeout(timeout time.Duration) *DeleteArchitectPromptResourceAudioParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete architect prompt resource audio params
func (o *DeleteArchitectPromptResourceAudioParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete architect prompt resource audio params
func (o *DeleteArchitectPromptResourceAudioParams) WithContext(ctx context.Context) *DeleteArchitectPromptResourceAudioParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete architect prompt resource audio params
func (o *DeleteArchitectPromptResourceAudioParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete architect prompt resource audio params
func (o *DeleteArchitectPromptResourceAudioParams) WithHTTPClient(client *http.Client) *DeleteArchitectPromptResourceAudioParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete architect prompt resource audio params
func (o *DeleteArchitectPromptResourceAudioParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguageCode adds the languageCode to the delete architect prompt resource audio params
func (o *DeleteArchitectPromptResourceAudioParams) WithLanguageCode(languageCode string) *DeleteArchitectPromptResourceAudioParams {
	o.SetLanguageCode(languageCode)
	return o
}

// SetLanguageCode adds the languageCode to the delete architect prompt resource audio params
func (o *DeleteArchitectPromptResourceAudioParams) SetLanguageCode(languageCode string) {
	o.LanguageCode = languageCode
}

// WithPromptID adds the promptID to the delete architect prompt resource audio params
func (o *DeleteArchitectPromptResourceAudioParams) WithPromptID(promptID string) *DeleteArchitectPromptResourceAudioParams {
	o.SetPromptID(promptID)
	return o
}

// SetPromptID adds the promptId to the delete architect prompt resource audio params
func (o *DeleteArchitectPromptResourceAudioParams) SetPromptID(promptID string) {
	o.PromptID = promptID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteArchitectPromptResourceAudioParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param languageCode
	if err := r.SetPathParam("languageCode", o.LanguageCode); err != nil {
		return err
	}

	// path param promptId
	if err := r.SetPathParam("promptId", o.PromptID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}