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

	"github.com/freman/genesysapi/models"
)

// NewPostArchitectIvrsParams creates a new PostArchitectIvrsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostArchitectIvrsParams() *PostArchitectIvrsParams {
	return &PostArchitectIvrsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostArchitectIvrsParamsWithTimeout creates a new PostArchitectIvrsParams object
// with the ability to set a timeout on a request.
func NewPostArchitectIvrsParamsWithTimeout(timeout time.Duration) *PostArchitectIvrsParams {
	return &PostArchitectIvrsParams{
		timeout: timeout,
	}
}

// NewPostArchitectIvrsParamsWithContext creates a new PostArchitectIvrsParams object
// with the ability to set a context for a request.
func NewPostArchitectIvrsParamsWithContext(ctx context.Context) *PostArchitectIvrsParams {
	return &PostArchitectIvrsParams{
		Context: ctx,
	}
}

// NewPostArchitectIvrsParamsWithHTTPClient creates a new PostArchitectIvrsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostArchitectIvrsParamsWithHTTPClient(client *http.Client) *PostArchitectIvrsParams {
	return &PostArchitectIvrsParams{
		HTTPClient: client,
	}
}

/*
PostArchitectIvrsParams contains all the parameters to send to the API endpoint

	for the post architect ivrs operation.

	Typically these are written to a http.Request.
*/
type PostArchitectIvrsParams struct {

	// Body.
	Body *models.IVR

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post architect ivrs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostArchitectIvrsParams) WithDefaults() *PostArchitectIvrsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post architect ivrs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostArchitectIvrsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post architect ivrs params
func (o *PostArchitectIvrsParams) WithTimeout(timeout time.Duration) *PostArchitectIvrsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post architect ivrs params
func (o *PostArchitectIvrsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post architect ivrs params
func (o *PostArchitectIvrsParams) WithContext(ctx context.Context) *PostArchitectIvrsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post architect ivrs params
func (o *PostArchitectIvrsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post architect ivrs params
func (o *PostArchitectIvrsParams) WithHTTPClient(client *http.Client) *PostArchitectIvrsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post architect ivrs params
func (o *PostArchitectIvrsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post architect ivrs params
func (o *PostArchitectIvrsParams) WithBody(body *models.IVR) *PostArchitectIvrsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post architect ivrs params
func (o *PostArchitectIvrsParams) SetBody(body *models.IVR) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostArchitectIvrsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
