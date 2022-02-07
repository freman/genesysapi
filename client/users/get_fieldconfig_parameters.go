// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetFieldconfigParams creates a new GetFieldconfigParams object
// with the default values initialized.
func NewGetFieldconfigParams() *GetFieldconfigParams {
	var ()
	return &GetFieldconfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFieldconfigParamsWithTimeout creates a new GetFieldconfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFieldconfigParamsWithTimeout(timeout time.Duration) *GetFieldconfigParams {
	var ()
	return &GetFieldconfigParams{

		timeout: timeout,
	}
}

// NewGetFieldconfigParamsWithContext creates a new GetFieldconfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFieldconfigParamsWithContext(ctx context.Context) *GetFieldconfigParams {
	var ()
	return &GetFieldconfigParams{

		Context: ctx,
	}
}

// NewGetFieldconfigParamsWithHTTPClient creates a new GetFieldconfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFieldconfigParamsWithHTTPClient(client *http.Client) *GetFieldconfigParams {
	var ()
	return &GetFieldconfigParams{
		HTTPClient: client,
	}
}

/*GetFieldconfigParams contains all the parameters to send to the API endpoint
for the get fieldconfig operation typically these are written to a http.Request
*/
type GetFieldconfigParams struct {

	/*Type
	  Field type

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fieldconfig params
func (o *GetFieldconfigParams) WithTimeout(timeout time.Duration) *GetFieldconfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fieldconfig params
func (o *GetFieldconfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fieldconfig params
func (o *GetFieldconfigParams) WithContext(ctx context.Context) *GetFieldconfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fieldconfig params
func (o *GetFieldconfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fieldconfig params
func (o *GetFieldconfigParams) WithHTTPClient(client *http.Client) *GetFieldconfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fieldconfig params
func (o *GetFieldconfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithType adds the typeVar to the get fieldconfig params
func (o *GetFieldconfigParams) WithType(typeVar string) *GetFieldconfigParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get fieldconfig params
func (o *GetFieldconfigParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetFieldconfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {
		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}