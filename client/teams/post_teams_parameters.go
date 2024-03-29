// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewPostTeamsParams creates a new PostTeamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostTeamsParams() *PostTeamsParams {
	return &PostTeamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostTeamsParamsWithTimeout creates a new PostTeamsParams object
// with the ability to set a timeout on a request.
func NewPostTeamsParamsWithTimeout(timeout time.Duration) *PostTeamsParams {
	return &PostTeamsParams{
		timeout: timeout,
	}
}

// NewPostTeamsParamsWithContext creates a new PostTeamsParams object
// with the ability to set a context for a request.
func NewPostTeamsParamsWithContext(ctx context.Context) *PostTeamsParams {
	return &PostTeamsParams{
		Context: ctx,
	}
}

// NewPostTeamsParamsWithHTTPClient creates a new PostTeamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostTeamsParamsWithHTTPClient(client *http.Client) *PostTeamsParams {
	return &PostTeamsParams{
		HTTPClient: client,
	}
}

/*
PostTeamsParams contains all the parameters to send to the API endpoint

	for the post teams operation.

	Typically these are written to a http.Request.
*/
type PostTeamsParams struct {

	/* Body.

	   Team
	*/
	Body *models.Team

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post teams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostTeamsParams) WithDefaults() *PostTeamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post teams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostTeamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post teams params
func (o *PostTeamsParams) WithTimeout(timeout time.Duration) *PostTeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post teams params
func (o *PostTeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post teams params
func (o *PostTeamsParams) WithContext(ctx context.Context) *PostTeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post teams params
func (o *PostTeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post teams params
func (o *PostTeamsParams) WithHTTPClient(client *http.Client) *PostTeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post teams params
func (o *PostTeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post teams params
func (o *PostTeamsParams) WithBody(body *models.Team) *PostTeamsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post teams params
func (o *PostTeamsParams) SetBody(body *models.Team) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
