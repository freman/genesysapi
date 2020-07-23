// Code generated by go-swagger; DO NOT EDIT.

package quality

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

// NewPostQualityCalibrationsParams creates a new PostQualityCalibrationsParams object
// with the default values initialized.
func NewPostQualityCalibrationsParams() *PostQualityCalibrationsParams {
	var ()
	return &PostQualityCalibrationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostQualityCalibrationsParamsWithTimeout creates a new PostQualityCalibrationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostQualityCalibrationsParamsWithTimeout(timeout time.Duration) *PostQualityCalibrationsParams {
	var ()
	return &PostQualityCalibrationsParams{

		timeout: timeout,
	}
}

// NewPostQualityCalibrationsParamsWithContext creates a new PostQualityCalibrationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostQualityCalibrationsParamsWithContext(ctx context.Context) *PostQualityCalibrationsParams {
	var ()
	return &PostQualityCalibrationsParams{

		Context: ctx,
	}
}

// NewPostQualityCalibrationsParamsWithHTTPClient creates a new PostQualityCalibrationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostQualityCalibrationsParamsWithHTTPClient(client *http.Client) *PostQualityCalibrationsParams {
	var ()
	return &PostQualityCalibrationsParams{
		HTTPClient: client,
	}
}

/*PostQualityCalibrationsParams contains all the parameters to send to the API endpoint
for the post quality calibrations operation typically these are written to a http.Request
*/
type PostQualityCalibrationsParams struct {

	/*Body
	  calibration

	*/
	Body *models.CalibrationCreate
	/*Expand
	  calibratorId

	*/
	Expand *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post quality calibrations params
func (o *PostQualityCalibrationsParams) WithTimeout(timeout time.Duration) *PostQualityCalibrationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post quality calibrations params
func (o *PostQualityCalibrationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post quality calibrations params
func (o *PostQualityCalibrationsParams) WithContext(ctx context.Context) *PostQualityCalibrationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post quality calibrations params
func (o *PostQualityCalibrationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post quality calibrations params
func (o *PostQualityCalibrationsParams) WithHTTPClient(client *http.Client) *PostQualityCalibrationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post quality calibrations params
func (o *PostQualityCalibrationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post quality calibrations params
func (o *PostQualityCalibrationsParams) WithBody(body *models.CalibrationCreate) *PostQualityCalibrationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post quality calibrations params
func (o *PostQualityCalibrationsParams) SetBody(body *models.CalibrationCreate) {
	o.Body = body
}

// WithExpand adds the expand to the post quality calibrations params
func (o *PostQualityCalibrationsParams) WithExpand(expand *string) *PostQualityCalibrationsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the post quality calibrations params
func (o *PostQualityCalibrationsParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *PostQualityCalibrationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Expand != nil {

		// query param expand
		var qrExpand string
		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {
			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
