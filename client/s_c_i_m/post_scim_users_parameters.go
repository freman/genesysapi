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

	"github.com/freman/genesysapi/models"
)

// NewPostScimUsersParams creates a new PostScimUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostScimUsersParams() *PostScimUsersParams {
	return &PostScimUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostScimUsersParamsWithTimeout creates a new PostScimUsersParams object
// with the ability to set a timeout on a request.
func NewPostScimUsersParamsWithTimeout(timeout time.Duration) *PostScimUsersParams {
	return &PostScimUsersParams{
		timeout: timeout,
	}
}

// NewPostScimUsersParamsWithContext creates a new PostScimUsersParams object
// with the ability to set a context for a request.
func NewPostScimUsersParamsWithContext(ctx context.Context) *PostScimUsersParams {
	return &PostScimUsersParams{
		Context: ctx,
	}
}

// NewPostScimUsersParamsWithHTTPClient creates a new PostScimUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostScimUsersParamsWithHTTPClient(client *http.Client) *PostScimUsersParams {
	return &PostScimUsersParams{
		HTTPClient: client,
	}
}

/*
PostScimUsersParams contains all the parameters to send to the API endpoint

	for the post scim users operation.

	Typically these are written to a http.Request.
*/
type PostScimUsersParams struct {

	/* Body.

	   The information used to create a user.
	*/
	Body *models.ScimV2CreateUser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post scim users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScimUsersParams) WithDefaults() *PostScimUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post scim users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScimUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post scim users params
func (o *PostScimUsersParams) WithTimeout(timeout time.Duration) *PostScimUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post scim users params
func (o *PostScimUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post scim users params
func (o *PostScimUsersParams) WithContext(ctx context.Context) *PostScimUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post scim users params
func (o *PostScimUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post scim users params
func (o *PostScimUsersParams) WithHTTPClient(client *http.Client) *PostScimUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post scim users params
func (o *PostScimUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post scim users params
func (o *PostScimUsersParams) WithBody(body *models.ScimV2CreateUser) *PostScimUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post scim users params
func (o *PostScimUsersParams) SetBody(body *models.ScimV2CreateUser) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostScimUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
