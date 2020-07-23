// Code generated by go-swagger; DO NOT EDIT.

package stations

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

// NewPatchStationsSettingsParams creates a new PatchStationsSettingsParams object
// with the default values initialized.
func NewPatchStationsSettingsParams() *PatchStationsSettingsParams {
	var ()
	return &PatchStationsSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchStationsSettingsParamsWithTimeout creates a new PatchStationsSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchStationsSettingsParamsWithTimeout(timeout time.Duration) *PatchStationsSettingsParams {
	var ()
	return &PatchStationsSettingsParams{

		timeout: timeout,
	}
}

// NewPatchStationsSettingsParamsWithContext creates a new PatchStationsSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchStationsSettingsParamsWithContext(ctx context.Context) *PatchStationsSettingsParams {
	var ()
	return &PatchStationsSettingsParams{

		Context: ctx,
	}
}

// NewPatchStationsSettingsParamsWithHTTPClient creates a new PatchStationsSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchStationsSettingsParamsWithHTTPClient(client *http.Client) *PatchStationsSettingsParams {
	var ()
	return &PatchStationsSettingsParams{
		HTTPClient: client,
	}
}

/*PatchStationsSettingsParams contains all the parameters to send to the API endpoint
for the patch stations settings operation typically these are written to a http.Request
*/
type PatchStationsSettingsParams struct {

	/*Body
	  Station settings

	*/
	Body *models.StationSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch stations settings params
func (o *PatchStationsSettingsParams) WithTimeout(timeout time.Duration) *PatchStationsSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch stations settings params
func (o *PatchStationsSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch stations settings params
func (o *PatchStationsSettingsParams) WithContext(ctx context.Context) *PatchStationsSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch stations settings params
func (o *PatchStationsSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch stations settings params
func (o *PatchStationsSettingsParams) WithHTTPClient(client *http.Client) *PatchStationsSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch stations settings params
func (o *PatchStationsSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch stations settings params
func (o *PatchStationsSettingsParams) WithBody(body *models.StationSettings) *PatchStationsSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch stations settings params
func (o *PatchStationsSettingsParams) SetBody(body *models.StationSettings) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchStationsSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
