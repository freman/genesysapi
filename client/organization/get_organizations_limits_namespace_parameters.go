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

// NewGetOrganizationsLimitsNamespaceParams creates a new GetOrganizationsLimitsNamespaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationsLimitsNamespaceParams() *GetOrganizationsLimitsNamespaceParams {
	return &GetOrganizationsLimitsNamespaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsLimitsNamespaceParamsWithTimeout creates a new GetOrganizationsLimitsNamespaceParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationsLimitsNamespaceParamsWithTimeout(timeout time.Duration) *GetOrganizationsLimitsNamespaceParams {
	return &GetOrganizationsLimitsNamespaceParams{
		timeout: timeout,
	}
}

// NewGetOrganizationsLimitsNamespaceParamsWithContext creates a new GetOrganizationsLimitsNamespaceParams object
// with the ability to set a context for a request.
func NewGetOrganizationsLimitsNamespaceParamsWithContext(ctx context.Context) *GetOrganizationsLimitsNamespaceParams {
	return &GetOrganizationsLimitsNamespaceParams{
		Context: ctx,
	}
}

// NewGetOrganizationsLimitsNamespaceParamsWithHTTPClient creates a new GetOrganizationsLimitsNamespaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationsLimitsNamespaceParamsWithHTTPClient(client *http.Client) *GetOrganizationsLimitsNamespaceParams {
	return &GetOrganizationsLimitsNamespaceParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationsLimitsNamespaceParams contains all the parameters to send to the API endpoint

	for the get organizations limits namespace operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationsLimitsNamespaceParams struct {

	/* NamespaceName.

	   The namespace to fetch limits for
	*/
	NamespaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organizations limits namespace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsLimitsNamespaceParams) WithDefaults() *GetOrganizationsLimitsNamespaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organizations limits namespace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsLimitsNamespaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organizations limits namespace params
func (o *GetOrganizationsLimitsNamespaceParams) WithTimeout(timeout time.Duration) *GetOrganizationsLimitsNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations limits namespace params
func (o *GetOrganizationsLimitsNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations limits namespace params
func (o *GetOrganizationsLimitsNamespaceParams) WithContext(ctx context.Context) *GetOrganizationsLimitsNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations limits namespace params
func (o *GetOrganizationsLimitsNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizations limits namespace params
func (o *GetOrganizationsLimitsNamespaceParams) WithHTTPClient(client *http.Client) *GetOrganizationsLimitsNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizations limits namespace params
func (o *GetOrganizationsLimitsNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceName adds the namespaceName to the get organizations limits namespace params
func (o *GetOrganizationsLimitsNamespaceParams) WithNamespaceName(namespaceName string) *GetOrganizationsLimitsNamespaceParams {
	o.SetNamespaceName(namespaceName)
	return o
}

// SetNamespaceName adds the namespaceName to the get organizations limits namespace params
func (o *GetOrganizationsLimitsNamespaceParams) SetNamespaceName(namespaceName string) {
	o.NamespaceName = namespaceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsLimitsNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespaceName
	if err := r.SetPathParam("namespaceName", o.NamespaceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
