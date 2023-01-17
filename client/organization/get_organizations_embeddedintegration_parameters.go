// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// NewGetOrganizationsEmbeddedintegrationParams creates a new GetOrganizationsEmbeddedintegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationsEmbeddedintegrationParams() *GetOrganizationsEmbeddedintegrationParams {
	return &GetOrganizationsEmbeddedintegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsEmbeddedintegrationParamsWithTimeout creates a new GetOrganizationsEmbeddedintegrationParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationsEmbeddedintegrationParamsWithTimeout(timeout time.Duration) *GetOrganizationsEmbeddedintegrationParams {
	return &GetOrganizationsEmbeddedintegrationParams{
		timeout: timeout,
	}
}

// NewGetOrganizationsEmbeddedintegrationParamsWithContext creates a new GetOrganizationsEmbeddedintegrationParams object
// with the ability to set a context for a request.
func NewGetOrganizationsEmbeddedintegrationParamsWithContext(ctx context.Context) *GetOrganizationsEmbeddedintegrationParams {
	return &GetOrganizationsEmbeddedintegrationParams{
		Context: ctx,
	}
}

// NewGetOrganizationsEmbeddedintegrationParamsWithHTTPClient creates a new GetOrganizationsEmbeddedintegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationsEmbeddedintegrationParamsWithHTTPClient(client *http.Client) *GetOrganizationsEmbeddedintegrationParams {
	return &GetOrganizationsEmbeddedintegrationParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationsEmbeddedintegrationParams contains all the parameters to send to the API endpoint

	for the get organizations embeddedintegration operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationsEmbeddedintegrationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organizations embeddedintegration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsEmbeddedintegrationParams) WithDefaults() *GetOrganizationsEmbeddedintegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organizations embeddedintegration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsEmbeddedintegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organizations embeddedintegration params
func (o *GetOrganizationsEmbeddedintegrationParams) WithTimeout(timeout time.Duration) *GetOrganizationsEmbeddedintegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations embeddedintegration params
func (o *GetOrganizationsEmbeddedintegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations embeddedintegration params
func (o *GetOrganizationsEmbeddedintegrationParams) WithContext(ctx context.Context) *GetOrganizationsEmbeddedintegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations embeddedintegration params
func (o *GetOrganizationsEmbeddedintegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizations embeddedintegration params
func (o *GetOrganizationsEmbeddedintegrationParams) WithHTTPClient(client *http.Client) *GetOrganizationsEmbeddedintegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizations embeddedintegration params
func (o *GetOrganizationsEmbeddedintegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsEmbeddedintegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
