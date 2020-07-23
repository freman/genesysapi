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
)

// NewPostFlowsDatatableRowsParams creates a new PostFlowsDatatableRowsParams object
// with the default values initialized.
func NewPostFlowsDatatableRowsParams() *PostFlowsDatatableRowsParams {
	var ()
	return &PostFlowsDatatableRowsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostFlowsDatatableRowsParamsWithTimeout creates a new PostFlowsDatatableRowsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostFlowsDatatableRowsParamsWithTimeout(timeout time.Duration) *PostFlowsDatatableRowsParams {
	var ()
	return &PostFlowsDatatableRowsParams{

		timeout: timeout,
	}
}

// NewPostFlowsDatatableRowsParamsWithContext creates a new PostFlowsDatatableRowsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostFlowsDatatableRowsParamsWithContext(ctx context.Context) *PostFlowsDatatableRowsParams {
	var ()
	return &PostFlowsDatatableRowsParams{

		Context: ctx,
	}
}

// NewPostFlowsDatatableRowsParamsWithHTTPClient creates a new PostFlowsDatatableRowsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostFlowsDatatableRowsParamsWithHTTPClient(client *http.Client) *PostFlowsDatatableRowsParams {
	var ()
	return &PostFlowsDatatableRowsParams{
		HTTPClient: client,
	}
}

/*PostFlowsDatatableRowsParams contains all the parameters to send to the API endpoint
for the post flows datatable rows operation typically these are written to a http.Request
*/
type PostFlowsDatatableRowsParams struct {

	/*DataTableRow*/
	DataTableRow map[string]interface{}
	/*DatatableID
	  id of datatable

	*/
	DatatableID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post flows datatable rows params
func (o *PostFlowsDatatableRowsParams) WithTimeout(timeout time.Duration) *PostFlowsDatatableRowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post flows datatable rows params
func (o *PostFlowsDatatableRowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post flows datatable rows params
func (o *PostFlowsDatatableRowsParams) WithContext(ctx context.Context) *PostFlowsDatatableRowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post flows datatable rows params
func (o *PostFlowsDatatableRowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post flows datatable rows params
func (o *PostFlowsDatatableRowsParams) WithHTTPClient(client *http.Client) *PostFlowsDatatableRowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post flows datatable rows params
func (o *PostFlowsDatatableRowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataTableRow adds the dataTableRow to the post flows datatable rows params
func (o *PostFlowsDatatableRowsParams) WithDataTableRow(dataTableRow map[string]interface{}) *PostFlowsDatatableRowsParams {
	o.SetDataTableRow(dataTableRow)
	return o
}

// SetDataTableRow adds the dataTableRow to the post flows datatable rows params
func (o *PostFlowsDatatableRowsParams) SetDataTableRow(dataTableRow map[string]interface{}) {
	o.DataTableRow = dataTableRow
}

// WithDatatableID adds the datatableID to the post flows datatable rows params
func (o *PostFlowsDatatableRowsParams) WithDatatableID(datatableID string) *PostFlowsDatatableRowsParams {
	o.SetDatatableID(datatableID)
	return o
}

// SetDatatableID adds the datatableId to the post flows datatable rows params
func (o *PostFlowsDatatableRowsParams) SetDatatableID(datatableID string) {
	o.DatatableID = datatableID
}

// WriteToRequest writes these params to a swagger request
func (o *PostFlowsDatatableRowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DataTableRow != nil {
		if err := r.SetBodyParam(o.DataTableRow); err != nil {
			return err
		}
	}

	// path param datatableId
	if err := r.SetPathParam("datatableId", o.DatatableID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
