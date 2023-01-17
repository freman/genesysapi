// Code generated by go-swagger; DO NOT EDIT.

package mobile_devices

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

// NewPostMobiledevicesParams creates a new PostMobiledevicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostMobiledevicesParams() *PostMobiledevicesParams {
	return &PostMobiledevicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostMobiledevicesParamsWithTimeout creates a new PostMobiledevicesParams object
// with the ability to set a timeout on a request.
func NewPostMobiledevicesParamsWithTimeout(timeout time.Duration) *PostMobiledevicesParams {
	return &PostMobiledevicesParams{
		timeout: timeout,
	}
}

// NewPostMobiledevicesParamsWithContext creates a new PostMobiledevicesParams object
// with the ability to set a context for a request.
func NewPostMobiledevicesParamsWithContext(ctx context.Context) *PostMobiledevicesParams {
	return &PostMobiledevicesParams{
		Context: ctx,
	}
}

// NewPostMobiledevicesParamsWithHTTPClient creates a new PostMobiledevicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostMobiledevicesParamsWithHTTPClient(client *http.Client) *PostMobiledevicesParams {
	return &PostMobiledevicesParams{
		HTTPClient: client,
	}
}

/*
PostMobiledevicesParams contains all the parameters to send to the API endpoint

	for the post mobiledevices operation.

	Typically these are written to a http.Request.
*/
type PostMobiledevicesParams struct {

	/* Body.

	   Device
	*/
	Body *models.UserDevice

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post mobiledevices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostMobiledevicesParams) WithDefaults() *PostMobiledevicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post mobiledevices params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostMobiledevicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post mobiledevices params
func (o *PostMobiledevicesParams) WithTimeout(timeout time.Duration) *PostMobiledevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post mobiledevices params
func (o *PostMobiledevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post mobiledevices params
func (o *PostMobiledevicesParams) WithContext(ctx context.Context) *PostMobiledevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post mobiledevices params
func (o *PostMobiledevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post mobiledevices params
func (o *PostMobiledevicesParams) WithHTTPClient(client *http.Client) *PostMobiledevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post mobiledevices params
func (o *PostMobiledevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post mobiledevices params
func (o *PostMobiledevicesParams) WithBody(body *models.UserDevice) *PostMobiledevicesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post mobiledevices params
func (o *PostMobiledevicesParams) SetBody(body *models.UserDevice) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostMobiledevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
