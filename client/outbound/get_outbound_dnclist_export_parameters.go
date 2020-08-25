// Code generated by go-swagger; DO NOT EDIT.

package outbound

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

// NewGetOutboundDnclistExportParams creates a new GetOutboundDnclistExportParams object
// with the default values initialized.
func NewGetOutboundDnclistExportParams() *GetOutboundDnclistExportParams {
	var (
		downloadDefault = string("false")
	)
	return &GetOutboundDnclistExportParams{
		Download: &downloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundDnclistExportParamsWithTimeout creates a new GetOutboundDnclistExportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutboundDnclistExportParamsWithTimeout(timeout time.Duration) *GetOutboundDnclistExportParams {
	var (
		downloadDefault = string("false")
	)
	return &GetOutboundDnclistExportParams{
		Download: &downloadDefault,

		timeout: timeout,
	}
}

// NewGetOutboundDnclistExportParamsWithContext creates a new GetOutboundDnclistExportParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutboundDnclistExportParamsWithContext(ctx context.Context) *GetOutboundDnclistExportParams {
	var (
		downloadDefault = string("false")
	)
	return &GetOutboundDnclistExportParams{
		Download: &downloadDefault,

		Context: ctx,
	}
}

// NewGetOutboundDnclistExportParamsWithHTTPClient creates a new GetOutboundDnclistExportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutboundDnclistExportParamsWithHTTPClient(client *http.Client) *GetOutboundDnclistExportParams {
	var (
		downloadDefault = string("false")
	)
	return &GetOutboundDnclistExportParams{
		Download:   &downloadDefault,
		HTTPClient: client,
	}
}

/*GetOutboundDnclistExportParams contains all the parameters to send to the API endpoint
for the get outbound dnclist export operation typically these are written to a http.Request
*/
type GetOutboundDnclistExportParams struct {

	/*DncListID
	  DncList ID

	*/
	DncListID string
	/*Download
	  Redirect to download uri

	*/
	Download *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get outbound dnclist export params
func (o *GetOutboundDnclistExportParams) WithTimeout(timeout time.Duration) *GetOutboundDnclistExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound dnclist export params
func (o *GetOutboundDnclistExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound dnclist export params
func (o *GetOutboundDnclistExportParams) WithContext(ctx context.Context) *GetOutboundDnclistExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound dnclist export params
func (o *GetOutboundDnclistExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound dnclist export params
func (o *GetOutboundDnclistExportParams) WithHTTPClient(client *http.Client) *GetOutboundDnclistExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound dnclist export params
func (o *GetOutboundDnclistExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDncListID adds the dncListID to the get outbound dnclist export params
func (o *GetOutboundDnclistExportParams) WithDncListID(dncListID string) *GetOutboundDnclistExportParams {
	o.SetDncListID(dncListID)
	return o
}

// SetDncListID adds the dncListId to the get outbound dnclist export params
func (o *GetOutboundDnclistExportParams) SetDncListID(dncListID string) {
	o.DncListID = dncListID
}

// WithDownload adds the download to the get outbound dnclist export params
func (o *GetOutboundDnclistExportParams) WithDownload(download *string) *GetOutboundDnclistExportParams {
	o.SetDownload(download)
	return o
}

// SetDownload adds the download to the get outbound dnclist export params
func (o *GetOutboundDnclistExportParams) SetDownload(download *string) {
	o.Download = download
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundDnclistExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dncListId
	if err := r.SetPathParam("dncListId", o.DncListID); err != nil {
		return err
	}

	if o.Download != nil {

		// query param download
		var qrDownload string
		if o.Download != nil {
			qrDownload = *o.Download
		}
		qDownload := qrDownload
		if qDownload != "" {
			if err := r.SetQueryParam("download", qDownload); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}