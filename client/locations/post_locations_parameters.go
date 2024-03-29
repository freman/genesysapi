// Code generated by go-swagger; DO NOT EDIT.

package locations

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

// NewPostLocationsParams creates a new PostLocationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLocationsParams() *PostLocationsParams {
	return &PostLocationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLocationsParamsWithTimeout creates a new PostLocationsParams object
// with the ability to set a timeout on a request.
func NewPostLocationsParamsWithTimeout(timeout time.Duration) *PostLocationsParams {
	return &PostLocationsParams{
		timeout: timeout,
	}
}

// NewPostLocationsParamsWithContext creates a new PostLocationsParams object
// with the ability to set a context for a request.
func NewPostLocationsParamsWithContext(ctx context.Context) *PostLocationsParams {
	return &PostLocationsParams{
		Context: ctx,
	}
}

// NewPostLocationsParamsWithHTTPClient creates a new PostLocationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLocationsParamsWithHTTPClient(client *http.Client) *PostLocationsParams {
	return &PostLocationsParams{
		HTTPClient: client,
	}
}

/*
PostLocationsParams contains all the parameters to send to the API endpoint

	for the post locations operation.

	Typically these are written to a http.Request.
*/
type PostLocationsParams struct {

	/* Body.

	   Location
	*/
	Body *models.LocationCreateDefinition

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post locations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLocationsParams) WithDefaults() *PostLocationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post locations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLocationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post locations params
func (o *PostLocationsParams) WithTimeout(timeout time.Duration) *PostLocationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post locations params
func (o *PostLocationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post locations params
func (o *PostLocationsParams) WithContext(ctx context.Context) *PostLocationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post locations params
func (o *PostLocationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post locations params
func (o *PostLocationsParams) WithHTTPClient(client *http.Client) *PostLocationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post locations params
func (o *PostLocationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post locations params
func (o *PostLocationsParams) WithBody(body *models.LocationCreateDefinition) *PostLocationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post locations params
func (o *PostLocationsParams) SetBody(body *models.LocationCreateDefinition) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostLocationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
