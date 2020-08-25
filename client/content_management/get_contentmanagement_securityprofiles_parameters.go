// Code generated by go-swagger; DO NOT EDIT.

package content_management

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

// NewGetContentmanagementSecurityprofilesParams creates a new GetContentmanagementSecurityprofilesParams object
// with the default values initialized.
func NewGetContentmanagementSecurityprofilesParams() *GetContentmanagementSecurityprofilesParams {

	return &GetContentmanagementSecurityprofilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentmanagementSecurityprofilesParamsWithTimeout creates a new GetContentmanagementSecurityprofilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContentmanagementSecurityprofilesParamsWithTimeout(timeout time.Duration) *GetContentmanagementSecurityprofilesParams {

	return &GetContentmanagementSecurityprofilesParams{

		timeout: timeout,
	}
}

// NewGetContentmanagementSecurityprofilesParamsWithContext creates a new GetContentmanagementSecurityprofilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContentmanagementSecurityprofilesParamsWithContext(ctx context.Context) *GetContentmanagementSecurityprofilesParams {

	return &GetContentmanagementSecurityprofilesParams{

		Context: ctx,
	}
}

// NewGetContentmanagementSecurityprofilesParamsWithHTTPClient creates a new GetContentmanagementSecurityprofilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContentmanagementSecurityprofilesParamsWithHTTPClient(client *http.Client) *GetContentmanagementSecurityprofilesParams {

	return &GetContentmanagementSecurityprofilesParams{
		HTTPClient: client,
	}
}

/*GetContentmanagementSecurityprofilesParams contains all the parameters to send to the API endpoint
for the get contentmanagement securityprofiles operation typically these are written to a http.Request
*/
type GetContentmanagementSecurityprofilesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get contentmanagement securityprofiles params
func (o *GetContentmanagementSecurityprofilesParams) WithTimeout(timeout time.Duration) *GetContentmanagementSecurityprofilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contentmanagement securityprofiles params
func (o *GetContentmanagementSecurityprofilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contentmanagement securityprofiles params
func (o *GetContentmanagementSecurityprofilesParams) WithContext(ctx context.Context) *GetContentmanagementSecurityprofilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contentmanagement securityprofiles params
func (o *GetContentmanagementSecurityprofilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contentmanagement securityprofiles params
func (o *GetContentmanagementSecurityprofilesParams) WithHTTPClient(client *http.Client) *GetContentmanagementSecurityprofilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contentmanagement securityprofiles params
func (o *GetContentmanagementSecurityprofilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentmanagementSecurityprofilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}