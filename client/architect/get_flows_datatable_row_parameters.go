// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewGetFlowsDatatableRowParams creates a new GetFlowsDatatableRowParams object
// with the default values initialized.
func NewGetFlowsDatatableRowParams() *GetFlowsDatatableRowParams {
	var (
		showbriefDefault = bool(true)
	)
	return &GetFlowsDatatableRowParams{
		Showbrief: &showbriefDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowsDatatableRowParamsWithTimeout creates a new GetFlowsDatatableRowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFlowsDatatableRowParamsWithTimeout(timeout time.Duration) *GetFlowsDatatableRowParams {
	var (
		showbriefDefault = bool(true)
	)
	return &GetFlowsDatatableRowParams{
		Showbrief: &showbriefDefault,

		timeout: timeout,
	}
}

// NewGetFlowsDatatableRowParamsWithContext creates a new GetFlowsDatatableRowParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFlowsDatatableRowParamsWithContext(ctx context.Context) *GetFlowsDatatableRowParams {
	var (
		showbriefDefault = bool(true)
	)
	return &GetFlowsDatatableRowParams{
		Showbrief: &showbriefDefault,

		Context: ctx,
	}
}

// NewGetFlowsDatatableRowParamsWithHTTPClient creates a new GetFlowsDatatableRowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFlowsDatatableRowParamsWithHTTPClient(client *http.Client) *GetFlowsDatatableRowParams {
	var (
		showbriefDefault = bool(true)
	)
	return &GetFlowsDatatableRowParams{
		Showbrief:  &showbriefDefault,
		HTTPClient: client,
	}
}

/*GetFlowsDatatableRowParams contains all the parameters to send to the API endpoint
for the get flows datatable row operation typically these are written to a http.Request
*/
type GetFlowsDatatableRowParams struct {

	/*DatatableID
	  id of datatable

	*/
	DatatableID string
	/*RowID
	  The key for the row

	*/
	RowID string
	/*Showbrief
	  if true returns just the key field for the row

	*/
	Showbrief *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get flows datatable row params
func (o *GetFlowsDatatableRowParams) WithTimeout(timeout time.Duration) *GetFlowsDatatableRowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flows datatable row params
func (o *GetFlowsDatatableRowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flows datatable row params
func (o *GetFlowsDatatableRowParams) WithContext(ctx context.Context) *GetFlowsDatatableRowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flows datatable row params
func (o *GetFlowsDatatableRowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flows datatable row params
func (o *GetFlowsDatatableRowParams) WithHTTPClient(client *http.Client) *GetFlowsDatatableRowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flows datatable row params
func (o *GetFlowsDatatableRowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatatableID adds the datatableID to the get flows datatable row params
func (o *GetFlowsDatatableRowParams) WithDatatableID(datatableID string) *GetFlowsDatatableRowParams {
	o.SetDatatableID(datatableID)
	return o
}

// SetDatatableID adds the datatableId to the get flows datatable row params
func (o *GetFlowsDatatableRowParams) SetDatatableID(datatableID string) {
	o.DatatableID = datatableID
}

// WithRowID adds the rowID to the get flows datatable row params
func (o *GetFlowsDatatableRowParams) WithRowID(rowID string) *GetFlowsDatatableRowParams {
	o.SetRowID(rowID)
	return o
}

// SetRowID adds the rowId to the get flows datatable row params
func (o *GetFlowsDatatableRowParams) SetRowID(rowID string) {
	o.RowID = rowID
}

// WithShowbrief adds the showbrief to the get flows datatable row params
func (o *GetFlowsDatatableRowParams) WithShowbrief(showbrief *bool) *GetFlowsDatatableRowParams {
	o.SetShowbrief(showbrief)
	return o
}

// SetShowbrief adds the showbrief to the get flows datatable row params
func (o *GetFlowsDatatableRowParams) SetShowbrief(showbrief *bool) {
	o.Showbrief = showbrief
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowsDatatableRowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datatableId
	if err := r.SetPathParam("datatableId", o.DatatableID); err != nil {
		return err
	}

	// path param rowId
	if err := r.SetPathParam("rowId", o.RowID); err != nil {
		return err
	}

	if o.Showbrief != nil {

		// query param showbrief
		var qrShowbrief bool
		if o.Showbrief != nil {
			qrShowbrief = *o.Showbrief
		}
		qShowbrief := swag.FormatBool(qrShowbrief)
		if qShowbrief != "" {
			if err := r.SetQueryParam("showbrief", qShowbrief); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}