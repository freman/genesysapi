// Code generated by go-swagger; DO NOT EDIT.

package geolocation

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

// NewPatchGeolocationsSettingsParams creates a new PatchGeolocationsSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchGeolocationsSettingsParams() *PatchGeolocationsSettingsParams {
	return &PatchGeolocationsSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchGeolocationsSettingsParamsWithTimeout creates a new PatchGeolocationsSettingsParams object
// with the ability to set a timeout on a request.
func NewPatchGeolocationsSettingsParamsWithTimeout(timeout time.Duration) *PatchGeolocationsSettingsParams {
	return &PatchGeolocationsSettingsParams{
		timeout: timeout,
	}
}

// NewPatchGeolocationsSettingsParamsWithContext creates a new PatchGeolocationsSettingsParams object
// with the ability to set a context for a request.
func NewPatchGeolocationsSettingsParamsWithContext(ctx context.Context) *PatchGeolocationsSettingsParams {
	return &PatchGeolocationsSettingsParams{
		Context: ctx,
	}
}

// NewPatchGeolocationsSettingsParamsWithHTTPClient creates a new PatchGeolocationsSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchGeolocationsSettingsParamsWithHTTPClient(client *http.Client) *PatchGeolocationsSettingsParams {
	return &PatchGeolocationsSettingsParams{
		HTTPClient: client,
	}
}

/*
PatchGeolocationsSettingsParams contains all the parameters to send to the API endpoint

	for the patch geolocations settings operation.

	Typically these are written to a http.Request.
*/
type PatchGeolocationsSettingsParams struct {

	/* Body.

	   Geolocation settings
	*/
	Body *models.GeolocationSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch geolocations settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchGeolocationsSettingsParams) WithDefaults() *PatchGeolocationsSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch geolocations settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchGeolocationsSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch geolocations settings params
func (o *PatchGeolocationsSettingsParams) WithTimeout(timeout time.Duration) *PatchGeolocationsSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch geolocations settings params
func (o *PatchGeolocationsSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch geolocations settings params
func (o *PatchGeolocationsSettingsParams) WithContext(ctx context.Context) *PatchGeolocationsSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch geolocations settings params
func (o *PatchGeolocationsSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch geolocations settings params
func (o *PatchGeolocationsSettingsParams) WithHTTPClient(client *http.Client) *PatchGeolocationsSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch geolocations settings params
func (o *PatchGeolocationsSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch geolocations settings params
func (o *PatchGeolocationsSettingsParams) WithBody(body *models.GeolocationSettings) *PatchGeolocationsSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch geolocations settings params
func (o *PatchGeolocationsSettingsParams) SetBody(body *models.GeolocationSettings) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchGeolocationsSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
