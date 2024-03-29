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
	"github.com/go-openapi/swag"
)

// NewGetOutboundDnclistsDivisionviewParams creates a new GetOutboundDnclistsDivisionviewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOutboundDnclistsDivisionviewParams() *GetOutboundDnclistsDivisionviewParams {
	return &GetOutboundDnclistsDivisionviewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundDnclistsDivisionviewParamsWithTimeout creates a new GetOutboundDnclistsDivisionviewParams object
// with the ability to set a timeout on a request.
func NewGetOutboundDnclistsDivisionviewParamsWithTimeout(timeout time.Duration) *GetOutboundDnclistsDivisionviewParams {
	return &GetOutboundDnclistsDivisionviewParams{
		timeout: timeout,
	}
}

// NewGetOutboundDnclistsDivisionviewParamsWithContext creates a new GetOutboundDnclistsDivisionviewParams object
// with the ability to set a context for a request.
func NewGetOutboundDnclistsDivisionviewParamsWithContext(ctx context.Context) *GetOutboundDnclistsDivisionviewParams {
	return &GetOutboundDnclistsDivisionviewParams{
		Context: ctx,
	}
}

// NewGetOutboundDnclistsDivisionviewParamsWithHTTPClient creates a new GetOutboundDnclistsDivisionviewParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOutboundDnclistsDivisionviewParamsWithHTTPClient(client *http.Client) *GetOutboundDnclistsDivisionviewParams {
	return &GetOutboundDnclistsDivisionviewParams{
		HTTPClient: client,
	}
}

/*
GetOutboundDnclistsDivisionviewParams contains all the parameters to send to the API endpoint

	for the get outbound dnclists divisionview operation.

	Typically these are written to a http.Request.
*/
type GetOutboundDnclistsDivisionviewParams struct {

	/* DncListID.

	   Dnclist ID
	*/
	DncListID string

	/* IncludeImportStatus.

	   Include import status
	*/
	IncludeImportStatus *bool

	/* IncludeSize.

	   Include size
	*/
	IncludeSize *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get outbound dnclists divisionview params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundDnclistsDivisionviewParams) WithDefaults() *GetOutboundDnclistsDivisionviewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get outbound dnclists divisionview params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundDnclistsDivisionviewParams) SetDefaults() {
	var (
		includeImportStatusDefault = bool(false)

		includeSizeDefault = bool(false)
	)

	val := GetOutboundDnclistsDivisionviewParams{
		IncludeImportStatus: &includeImportStatusDefault,
		IncludeSize:         &includeSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get outbound dnclists divisionview params
func (o *GetOutboundDnclistsDivisionviewParams) WithTimeout(timeout time.Duration) *GetOutboundDnclistsDivisionviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound dnclists divisionview params
func (o *GetOutboundDnclistsDivisionviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound dnclists divisionview params
func (o *GetOutboundDnclistsDivisionviewParams) WithContext(ctx context.Context) *GetOutboundDnclistsDivisionviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound dnclists divisionview params
func (o *GetOutboundDnclistsDivisionviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound dnclists divisionview params
func (o *GetOutboundDnclistsDivisionviewParams) WithHTTPClient(client *http.Client) *GetOutboundDnclistsDivisionviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound dnclists divisionview params
func (o *GetOutboundDnclistsDivisionviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDncListID adds the dncListID to the get outbound dnclists divisionview params
func (o *GetOutboundDnclistsDivisionviewParams) WithDncListID(dncListID string) *GetOutboundDnclistsDivisionviewParams {
	o.SetDncListID(dncListID)
	return o
}

// SetDncListID adds the dncListId to the get outbound dnclists divisionview params
func (o *GetOutboundDnclistsDivisionviewParams) SetDncListID(dncListID string) {
	o.DncListID = dncListID
}

// WithIncludeImportStatus adds the includeImportStatus to the get outbound dnclists divisionview params
func (o *GetOutboundDnclistsDivisionviewParams) WithIncludeImportStatus(includeImportStatus *bool) *GetOutboundDnclistsDivisionviewParams {
	o.SetIncludeImportStatus(includeImportStatus)
	return o
}

// SetIncludeImportStatus adds the includeImportStatus to the get outbound dnclists divisionview params
func (o *GetOutboundDnclistsDivisionviewParams) SetIncludeImportStatus(includeImportStatus *bool) {
	o.IncludeImportStatus = includeImportStatus
}

// WithIncludeSize adds the includeSize to the get outbound dnclists divisionview params
func (o *GetOutboundDnclistsDivisionviewParams) WithIncludeSize(includeSize *bool) *GetOutboundDnclistsDivisionviewParams {
	o.SetIncludeSize(includeSize)
	return o
}

// SetIncludeSize adds the includeSize to the get outbound dnclists divisionview params
func (o *GetOutboundDnclistsDivisionviewParams) SetIncludeSize(includeSize *bool) {
	o.IncludeSize = includeSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundDnclistsDivisionviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dncListId
	if err := r.SetPathParam("dncListId", o.DncListID); err != nil {
		return err
	}

	if o.IncludeImportStatus != nil {

		// query param includeImportStatus
		var qrIncludeImportStatus bool

		if o.IncludeImportStatus != nil {
			qrIncludeImportStatus = *o.IncludeImportStatus
		}
		qIncludeImportStatus := swag.FormatBool(qrIncludeImportStatus)
		if qIncludeImportStatus != "" {

			if err := r.SetQueryParam("includeImportStatus", qIncludeImportStatus); err != nil {
				return err
			}
		}
	}

	if o.IncludeSize != nil {

		// query param includeSize
		var qrIncludeSize bool

		if o.IncludeSize != nil {
			qrIncludeSize = *o.IncludeSize
		}
		qIncludeSize := swag.FormatBool(qrIncludeSize)
		if qIncludeSize != "" {

			if err := r.SetQueryParam("includeSize", qIncludeSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
