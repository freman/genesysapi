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

// NewPostOutboundDnclistExportParams creates a new PostOutboundDnclistExportParams object
// with the default values initialized.
func NewPostOutboundDnclistExportParams() *PostOutboundDnclistExportParams {
	var ()
	return &PostOutboundDnclistExportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOutboundDnclistExportParamsWithTimeout creates a new PostOutboundDnclistExportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOutboundDnclistExportParamsWithTimeout(timeout time.Duration) *PostOutboundDnclistExportParams {
	var ()
	return &PostOutboundDnclistExportParams{

		timeout: timeout,
	}
}

// NewPostOutboundDnclistExportParamsWithContext creates a new PostOutboundDnclistExportParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOutboundDnclistExportParamsWithContext(ctx context.Context) *PostOutboundDnclistExportParams {
	var ()
	return &PostOutboundDnclistExportParams{

		Context: ctx,
	}
}

// NewPostOutboundDnclistExportParamsWithHTTPClient creates a new PostOutboundDnclistExportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOutboundDnclistExportParamsWithHTTPClient(client *http.Client) *PostOutboundDnclistExportParams {
	var ()
	return &PostOutboundDnclistExportParams{
		HTTPClient: client,
	}
}

/*PostOutboundDnclistExportParams contains all the parameters to send to the API endpoint
for the post outbound dnclist export operation typically these are written to a http.Request
*/
type PostOutboundDnclistExportParams struct {

	/*DncListID
	  DncList ID

	*/
	DncListID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post outbound dnclist export params
func (o *PostOutboundDnclistExportParams) WithTimeout(timeout time.Duration) *PostOutboundDnclistExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post outbound dnclist export params
func (o *PostOutboundDnclistExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post outbound dnclist export params
func (o *PostOutboundDnclistExportParams) WithContext(ctx context.Context) *PostOutboundDnclistExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post outbound dnclist export params
func (o *PostOutboundDnclistExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post outbound dnclist export params
func (o *PostOutboundDnclistExportParams) WithHTTPClient(client *http.Client) *PostOutboundDnclistExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post outbound dnclist export params
func (o *PostOutboundDnclistExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDncListID adds the dncListID to the post outbound dnclist export params
func (o *PostOutboundDnclistExportParams) WithDncListID(dncListID string) *PostOutboundDnclistExportParams {
	o.SetDncListID(dncListID)
	return o
}

// SetDncListID adds the dncListId to the post outbound dnclist export params
func (o *PostOutboundDnclistExportParams) SetDncListID(dncListID string) {
	o.DncListID = dncListID
}

// WriteToRequest writes these params to a swagger request
func (o *PostOutboundDnclistExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dncListId
	if err := r.SetPathParam("dncListId", o.DncListID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}