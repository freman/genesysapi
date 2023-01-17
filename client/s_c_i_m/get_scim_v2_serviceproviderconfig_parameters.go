// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

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

// NewGetScimV2ServiceproviderconfigParams creates a new GetScimV2ServiceproviderconfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScimV2ServiceproviderconfigParams() *GetScimV2ServiceproviderconfigParams {
	return &GetScimV2ServiceproviderconfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScimV2ServiceproviderconfigParamsWithTimeout creates a new GetScimV2ServiceproviderconfigParams object
// with the ability to set a timeout on a request.
func NewGetScimV2ServiceproviderconfigParamsWithTimeout(timeout time.Duration) *GetScimV2ServiceproviderconfigParams {
	return &GetScimV2ServiceproviderconfigParams{
		timeout: timeout,
	}
}

// NewGetScimV2ServiceproviderconfigParamsWithContext creates a new GetScimV2ServiceproviderconfigParams object
// with the ability to set a context for a request.
func NewGetScimV2ServiceproviderconfigParamsWithContext(ctx context.Context) *GetScimV2ServiceproviderconfigParams {
	return &GetScimV2ServiceproviderconfigParams{
		Context: ctx,
	}
}

// NewGetScimV2ServiceproviderconfigParamsWithHTTPClient creates a new GetScimV2ServiceproviderconfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScimV2ServiceproviderconfigParamsWithHTTPClient(client *http.Client) *GetScimV2ServiceproviderconfigParams {
	return &GetScimV2ServiceproviderconfigParams{
		HTTPClient: client,
	}
}

/*
GetScimV2ServiceproviderconfigParams contains all the parameters to send to the API endpoint

	for the get scim v2 serviceproviderconfig operation.

	Typically these are written to a http.Request.
*/
type GetScimV2ServiceproviderconfigParams struct {

	/* IfNoneMatch.

	   The ETag of a resource in double quotes. Returned as header and meta.version with initial call to GET /api/v2/scim/v2/serviceproviderconfig. Example: "42". If the ETag is different from the version on the server, returns the current configuration of the resource. If the ETag is current, returns 304 Not Modified.
	*/
	IfNoneMatch *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get scim v2 serviceproviderconfig params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScimV2ServiceproviderconfigParams) WithDefaults() *GetScimV2ServiceproviderconfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scim v2 serviceproviderconfig params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScimV2ServiceproviderconfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get scim v2 serviceproviderconfig params
func (o *GetScimV2ServiceproviderconfigParams) WithTimeout(timeout time.Duration) *GetScimV2ServiceproviderconfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scim v2 serviceproviderconfig params
func (o *GetScimV2ServiceproviderconfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scim v2 serviceproviderconfig params
func (o *GetScimV2ServiceproviderconfigParams) WithContext(ctx context.Context) *GetScimV2ServiceproviderconfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scim v2 serviceproviderconfig params
func (o *GetScimV2ServiceproviderconfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scim v2 serviceproviderconfig params
func (o *GetScimV2ServiceproviderconfigParams) WithHTTPClient(client *http.Client) *GetScimV2ServiceproviderconfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scim v2 serviceproviderconfig params
func (o *GetScimV2ServiceproviderconfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get scim v2 serviceproviderconfig params
func (o *GetScimV2ServiceproviderconfigParams) WithIfNoneMatch(ifNoneMatch *string) *GetScimV2ServiceproviderconfigParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get scim v2 serviceproviderconfig params
func (o *GetScimV2ServiceproviderconfigParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WriteToRequest writes these params to a swagger request
func (o *GetScimV2ServiceproviderconfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfNoneMatch != nil {

		// header param If-None-Match
		if err := r.SetHeaderParam("If-None-Match", *o.IfNoneMatch); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
