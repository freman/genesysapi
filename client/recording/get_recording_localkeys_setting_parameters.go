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
)

// NewGetRecordingLocalkeysSettingParams creates a new GetRecordingLocalkeysSettingParams object
// with the default values initialized.
func NewGetRecordingLocalkeysSettingParams() *GetRecordingLocalkeysSettingParams {
	var ()
	return &GetRecordingLocalkeysSettingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecordingLocalkeysSettingParamsWithTimeout creates a new GetRecordingLocalkeysSettingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecordingLocalkeysSettingParamsWithTimeout(timeout time.Duration) *GetRecordingLocalkeysSettingParams {
	var ()
	return &GetRecordingLocalkeysSettingParams{

		timeout: timeout,
	}
}

// NewGetRecordingLocalkeysSettingParamsWithContext creates a new GetRecordingLocalkeysSettingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecordingLocalkeysSettingParamsWithContext(ctx context.Context) *GetRecordingLocalkeysSettingParams {
	var ()
	return &GetRecordingLocalkeysSettingParams{

		Context: ctx,
	}
}

// NewGetRecordingLocalkeysSettingParamsWithHTTPClient creates a new GetRecordingLocalkeysSettingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecordingLocalkeysSettingParamsWithHTTPClient(client *http.Client) *GetRecordingLocalkeysSettingParams {
	var ()
	return &GetRecordingLocalkeysSettingParams{
		HTTPClient: client,
	}
}

/*GetRecordingLocalkeysSettingParams contains all the parameters to send to the API endpoint
for the get recording localkeys setting operation typically these are written to a http.Request
*/
type GetRecordingLocalkeysSettingParams struct {

	/*SettingsID
	  Settings Id

	*/
	SettingsID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recording localkeys setting params
func (o *GetRecordingLocalkeysSettingParams) WithTimeout(timeout time.Duration) *GetRecordingLocalkeysSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recording localkeys setting params
func (o *GetRecordingLocalkeysSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recording localkeys setting params
func (o *GetRecordingLocalkeysSettingParams) WithContext(ctx context.Context) *GetRecordingLocalkeysSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recording localkeys setting params
func (o *GetRecordingLocalkeysSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recording localkeys setting params
func (o *GetRecordingLocalkeysSettingParams) WithHTTPClient(client *http.Client) *GetRecordingLocalkeysSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recording localkeys setting params
func (o *GetRecordingLocalkeysSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettingsID adds the settingsID to the get recording localkeys setting params
func (o *GetRecordingLocalkeysSettingParams) WithSettingsID(settingsID string) *GetRecordingLocalkeysSettingParams {
	o.SetSettingsID(settingsID)
	return o
}

// SetSettingsID adds the settingsId to the get recording localkeys setting params
func (o *GetRecordingLocalkeysSettingParams) SetSettingsID(settingsID string) {
	o.SettingsID = settingsID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecordingLocalkeysSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param settingsId
	if err := r.SetPathParam("settingsId", o.SettingsID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}