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

// NewGetGamificationProfilesParams creates a new GetGamificationProfilesParams object
// with the default values initialized.
func NewGetGamificationProfilesParams() *GetGamificationProfilesParams {

	return &GetGamificationProfilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGamificationProfilesParamsWithTimeout creates a new GetGamificationProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGamificationProfilesParamsWithTimeout(timeout time.Duration) *GetGamificationProfilesParams {

	return &GetGamificationProfilesParams{

		timeout: timeout,
	}
}

// NewGetGamificationProfilesParamsWithContext creates a new GetGamificationProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGamificationProfilesParamsWithContext(ctx context.Context) *GetGamificationProfilesParams {

	return &GetGamificationProfilesParams{

		Context: ctx,
	}
}

// NewGetGamificationProfilesParamsWithHTTPClient creates a new GetGamificationProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGamificationProfilesParamsWithHTTPClient(client *http.Client) *GetGamificationProfilesParams {

	return &GetGamificationProfilesParams{
		HTTPClient: client,
	}
}

/*GetGamificationProfilesParams contains all the parameters to send to the API endpoint
for the get gamification profiles operation typically these are written to a http.Request
*/
type GetGamificationProfilesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gamification profiles params
func (o *GetGamificationProfilesParams) WithTimeout(timeout time.Duration) *GetGamificationProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gamification profiles params
func (o *GetGamificationProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gamification profiles params
func (o *GetGamificationProfilesParams) WithContext(ctx context.Context) *GetGamificationProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gamification profiles params
func (o *GetGamificationProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gamification profiles params
func (o *GetGamificationProfilesParams) WithHTTPClient(client *http.Client) *GetGamificationProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gamification profiles params
func (o *GetGamificationProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetGamificationProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
