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

// NewGetArchitectPromptParams creates a new GetArchitectPromptParams object
// with the default values initialized.
func NewGetArchitectPromptParams() *GetArchitectPromptParams {
	var ()
	return &GetArchitectPromptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectPromptParamsWithTimeout creates a new GetArchitectPromptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArchitectPromptParamsWithTimeout(timeout time.Duration) *GetArchitectPromptParams {
	var ()
	return &GetArchitectPromptParams{

		timeout: timeout,
	}
}

// NewGetArchitectPromptParamsWithContext creates a new GetArchitectPromptParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArchitectPromptParamsWithContext(ctx context.Context) *GetArchitectPromptParams {
	var ()
	return &GetArchitectPromptParams{

		Context: ctx,
	}
}

// NewGetArchitectPromptParamsWithHTTPClient creates a new GetArchitectPromptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArchitectPromptParamsWithHTTPClient(client *http.Client) *GetArchitectPromptParams {
	var ()
	return &GetArchitectPromptParams{
		HTTPClient: client,
	}
}

/*GetArchitectPromptParams contains all the parameters to send to the API endpoint
for the get architect prompt operation typically these are written to a http.Request
*/
type GetArchitectPromptParams struct {

	/*PromptID
	  Prompt ID

	*/
	PromptID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get architect prompt params
func (o *GetArchitectPromptParams) WithTimeout(timeout time.Duration) *GetArchitectPromptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect prompt params
func (o *GetArchitectPromptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect prompt params
func (o *GetArchitectPromptParams) WithContext(ctx context.Context) *GetArchitectPromptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect prompt params
func (o *GetArchitectPromptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect prompt params
func (o *GetArchitectPromptParams) WithHTTPClient(client *http.Client) *GetArchitectPromptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect prompt params
func (o *GetArchitectPromptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPromptID adds the promptID to the get architect prompt params
func (o *GetArchitectPromptParams) WithPromptID(promptID string) *GetArchitectPromptParams {
	o.SetPromptID(promptID)
	return o
}

// SetPromptID adds the promptId to the get architect prompt params
func (o *GetArchitectPromptParams) SetPromptID(promptID string) {
	o.PromptID = promptID
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectPromptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param promptId
	if err := r.SetPathParam("promptId", o.PromptID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
