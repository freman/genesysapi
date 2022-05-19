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

// NewGetRecordingKeyconfigurationParams creates a new GetRecordingKeyconfigurationParams object
// with the default values initialized.
func NewGetRecordingKeyconfigurationParams() *GetRecordingKeyconfigurationParams {
	var ()
	return &GetRecordingKeyconfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecordingKeyconfigurationParamsWithTimeout creates a new GetRecordingKeyconfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecordingKeyconfigurationParamsWithTimeout(timeout time.Duration) *GetRecordingKeyconfigurationParams {
	var ()
	return &GetRecordingKeyconfigurationParams{

		timeout: timeout,
	}
}

// NewGetRecordingKeyconfigurationParamsWithContext creates a new GetRecordingKeyconfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecordingKeyconfigurationParamsWithContext(ctx context.Context) *GetRecordingKeyconfigurationParams {
	var ()
	return &GetRecordingKeyconfigurationParams{

		Context: ctx,
	}
}

// NewGetRecordingKeyconfigurationParamsWithHTTPClient creates a new GetRecordingKeyconfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecordingKeyconfigurationParamsWithHTTPClient(client *http.Client) *GetRecordingKeyconfigurationParams {
	var ()
	return &GetRecordingKeyconfigurationParams{
		HTTPClient: client,
	}
}

/*GetRecordingKeyconfigurationParams contains all the parameters to send to the API endpoint
for the get recording keyconfiguration operation typically these are written to a http.Request
*/
type GetRecordingKeyconfigurationParams struct {

	/*KeyConfigurationID
	  Key Configurations Id

	*/
	KeyConfigurationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recording keyconfiguration params
func (o *GetRecordingKeyconfigurationParams) WithTimeout(timeout time.Duration) *GetRecordingKeyconfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recording keyconfiguration params
func (o *GetRecordingKeyconfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recording keyconfiguration params
func (o *GetRecordingKeyconfigurationParams) WithContext(ctx context.Context) *GetRecordingKeyconfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recording keyconfiguration params
func (o *GetRecordingKeyconfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recording keyconfiguration params
func (o *GetRecordingKeyconfigurationParams) WithHTTPClient(client *http.Client) *GetRecordingKeyconfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recording keyconfiguration params
func (o *GetRecordingKeyconfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyConfigurationID adds the keyConfigurationID to the get recording keyconfiguration params
func (o *GetRecordingKeyconfigurationParams) WithKeyConfigurationID(keyConfigurationID string) *GetRecordingKeyconfigurationParams {
	o.SetKeyConfigurationID(keyConfigurationID)
	return o
}

// SetKeyConfigurationID adds the keyConfigurationId to the get recording keyconfiguration params
func (o *GetRecordingKeyconfigurationParams) SetKeyConfigurationID(keyConfigurationID string) {
	o.KeyConfigurationID = keyConfigurationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecordingKeyconfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param keyConfigurationId
	if err := r.SetPathParam("keyConfigurationId", o.KeyConfigurationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
