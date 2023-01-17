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

// NewGetOutboundContactlistExportParams creates a new GetOutboundContactlistExportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOutboundContactlistExportParams() *GetOutboundContactlistExportParams {
	return &GetOutboundContactlistExportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundContactlistExportParamsWithTimeout creates a new GetOutboundContactlistExportParams object
// with the ability to set a timeout on a request.
func NewGetOutboundContactlistExportParamsWithTimeout(timeout time.Duration) *GetOutboundContactlistExportParams {
	return &GetOutboundContactlistExportParams{
		timeout: timeout,
	}
}

// NewGetOutboundContactlistExportParamsWithContext creates a new GetOutboundContactlistExportParams object
// with the ability to set a context for a request.
func NewGetOutboundContactlistExportParamsWithContext(ctx context.Context) *GetOutboundContactlistExportParams {
	return &GetOutboundContactlistExportParams{
		Context: ctx,
	}
}

// NewGetOutboundContactlistExportParamsWithHTTPClient creates a new GetOutboundContactlistExportParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOutboundContactlistExportParamsWithHTTPClient(client *http.Client) *GetOutboundContactlistExportParams {
	return &GetOutboundContactlistExportParams{
		HTTPClient: client,
	}
}

/*
GetOutboundContactlistExportParams contains all the parameters to send to the API endpoint

	for the get outbound contactlist export operation.

	Typically these are written to a http.Request.
*/
type GetOutboundContactlistExportParams struct {

	/* ContactListID.

	   ContactList ID
	*/
	ContactListID string

	/* Download.

	   Redirect to download uri

	   Default: "false"
	*/
	Download *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get outbound contactlist export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundContactlistExportParams) WithDefaults() *GetOutboundContactlistExportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get outbound contactlist export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundContactlistExportParams) SetDefaults() {
	var (
		downloadDefault = string("false")
	)

	val := GetOutboundContactlistExportParams{
		Download: &downloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get outbound contactlist export params
func (o *GetOutboundContactlistExportParams) WithTimeout(timeout time.Duration) *GetOutboundContactlistExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound contactlist export params
func (o *GetOutboundContactlistExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound contactlist export params
func (o *GetOutboundContactlistExportParams) WithContext(ctx context.Context) *GetOutboundContactlistExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound contactlist export params
func (o *GetOutboundContactlistExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound contactlist export params
func (o *GetOutboundContactlistExportParams) WithHTTPClient(client *http.Client) *GetOutboundContactlistExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound contactlist export params
func (o *GetOutboundContactlistExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactListID adds the contactListID to the get outbound contactlist export params
func (o *GetOutboundContactlistExportParams) WithContactListID(contactListID string) *GetOutboundContactlistExportParams {
	o.SetContactListID(contactListID)
	return o
}

// SetContactListID adds the contactListId to the get outbound contactlist export params
func (o *GetOutboundContactlistExportParams) SetContactListID(contactListID string) {
	o.ContactListID = contactListID
}

// WithDownload adds the download to the get outbound contactlist export params
func (o *GetOutboundContactlistExportParams) WithDownload(download *string) *GetOutboundContactlistExportParams {
	o.SetDownload(download)
	return o
}

// SetDownload adds the download to the get outbound contactlist export params
func (o *GetOutboundContactlistExportParams) SetDownload(download *string) {
	o.Download = download
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundContactlistExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contactListId
	if err := r.SetPathParam("contactListId", o.ContactListID); err != nil {
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
