// Code generated by go-swagger; DO NOT EDIT.

package gamification

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

// NewGetGamificationTemplatesParams creates a new GetGamificationTemplatesParams object
// with the default values initialized.
func NewGetGamificationTemplatesParams() *GetGamificationTemplatesParams {

	return &GetGamificationTemplatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGamificationTemplatesParamsWithTimeout creates a new GetGamificationTemplatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGamificationTemplatesParamsWithTimeout(timeout time.Duration) *GetGamificationTemplatesParams {

	return &GetGamificationTemplatesParams{

		timeout: timeout,
	}
}

// NewGetGamificationTemplatesParamsWithContext creates a new GetGamificationTemplatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGamificationTemplatesParamsWithContext(ctx context.Context) *GetGamificationTemplatesParams {

	return &GetGamificationTemplatesParams{

		Context: ctx,
	}
}

// NewGetGamificationTemplatesParamsWithHTTPClient creates a new GetGamificationTemplatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGamificationTemplatesParamsWithHTTPClient(client *http.Client) *GetGamificationTemplatesParams {

	return &GetGamificationTemplatesParams{
		HTTPClient: client,
	}
}

/*GetGamificationTemplatesParams contains all the parameters to send to the API endpoint
for the get gamification templates operation typically these are written to a http.Request
*/
type GetGamificationTemplatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gamification templates params
func (o *GetGamificationTemplatesParams) WithTimeout(timeout time.Duration) *GetGamificationTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gamification templates params
func (o *GetGamificationTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gamification templates params
func (o *GetGamificationTemplatesParams) WithContext(ctx context.Context) *GetGamificationTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gamification templates params
func (o *GetGamificationTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gamification templates params
func (o *GetGamificationTemplatesParams) WithHTTPClient(client *http.Client) *GetGamificationTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gamification templates params
func (o *GetGamificationTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetGamificationTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}