// Code generated by go-swagger; DO NOT EDIT.

package recording

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
	"github.com/go-openapi/swag"
)

// NewGetRecordingSettingsParams creates a new GetRecordingSettingsParams object
// with the default values initialized.
func NewGetRecordingSettingsParams() *GetRecordingSettingsParams {
	var (
		createDefaultDefault = bool(false)
	)
	return &GetRecordingSettingsParams{
		CreateDefault: &createDefaultDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecordingSettingsParamsWithTimeout creates a new GetRecordingSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecordingSettingsParamsWithTimeout(timeout time.Duration) *GetRecordingSettingsParams {
	var (
		createDefaultDefault = bool(false)
	)
	return &GetRecordingSettingsParams{
		CreateDefault: &createDefaultDefault,

		timeout: timeout,
	}
}

// NewGetRecordingSettingsParamsWithContext creates a new GetRecordingSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecordingSettingsParamsWithContext(ctx context.Context) *GetRecordingSettingsParams {
	var (
		createDefaultDefault = bool(false)
	)
	return &GetRecordingSettingsParams{
		CreateDefault: &createDefaultDefault,

		Context: ctx,
	}
}

// NewGetRecordingSettingsParamsWithHTTPClient creates a new GetRecordingSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecordingSettingsParamsWithHTTPClient(client *http.Client) *GetRecordingSettingsParams {
	var (
		createDefaultDefault = bool(false)
	)
	return &GetRecordingSettingsParams{
		CreateDefault: &createDefaultDefault,
		HTTPClient:    client,
	}
}

/*GetRecordingSettingsParams contains all the parameters to send to the API endpoint
for the get recording settings operation typically these are written to a http.Request
*/
type GetRecordingSettingsParams struct {

	/*CreateDefault
	  If no settings are found, a new one is created with default values

	*/
	CreateDefault *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recording settings params
func (o *GetRecordingSettingsParams) WithTimeout(timeout time.Duration) *GetRecordingSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recording settings params
func (o *GetRecordingSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recording settings params
func (o *GetRecordingSettingsParams) WithContext(ctx context.Context) *GetRecordingSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recording settings params
func (o *GetRecordingSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recording settings params
func (o *GetRecordingSettingsParams) WithHTTPClient(client *http.Client) *GetRecordingSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recording settings params
func (o *GetRecordingSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateDefault adds the createDefault to the get recording settings params
func (o *GetRecordingSettingsParams) WithCreateDefault(createDefault *bool) *GetRecordingSettingsParams {
	o.SetCreateDefault(createDefault)
	return o
}

// SetCreateDefault adds the createDefault to the get recording settings params
func (o *GetRecordingSettingsParams) SetCreateDefault(createDefault *bool) {
	o.CreateDefault = createDefault
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecordingSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateDefault != nil {

		// query param createDefault
		var qrCreateDefault bool
		if o.CreateDefault != nil {
			qrCreateDefault = *o.CreateDefault
		}
		qCreateDefault := swag.FormatBool(qrCreateDefault)
		if qCreateDefault != "" {
			if err := r.SetQueryParam("createDefault", qCreateDefault); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
