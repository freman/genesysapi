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

// NewPutFlowsDatatableRowParams creates a new PutFlowsDatatableRowParams object
// with the default values initialized.
func NewPutFlowsDatatableRowParams() *PutFlowsDatatableRowParams {
	var ()
	return &PutFlowsDatatableRowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutFlowsDatatableRowParamsWithTimeout creates a new PutFlowsDatatableRowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutFlowsDatatableRowParamsWithTimeout(timeout time.Duration) *PutFlowsDatatableRowParams {
	var ()
	return &PutFlowsDatatableRowParams{

		timeout: timeout,
	}
}

// NewPutFlowsDatatableRowParamsWithContext creates a new PutFlowsDatatableRowParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutFlowsDatatableRowParamsWithContext(ctx context.Context) *PutFlowsDatatableRowParams {
	var ()
	return &PutFlowsDatatableRowParams{

		Context: ctx,
	}
}

// NewPutFlowsDatatableRowParamsWithHTTPClient creates a new PutFlowsDatatableRowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutFlowsDatatableRowParamsWithHTTPClient(client *http.Client) *PutFlowsDatatableRowParams {
	var ()
	return &PutFlowsDatatableRowParams{
		HTTPClient: client,
	}
}

/*PutFlowsDatatableRowParams contains all the parameters to send to the API endpoint
for the put flows datatable row operation typically these are written to a http.Request
*/
type PutFlowsDatatableRowParams struct {

	/*Body
	  datatable row

	*/
	Body map[string]interface{}
	/*DatatableID
	  id of datatable

	*/
	DatatableID string
	/*RowID
	  the key for the row

	*/
	RowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put flows datatable row params
func (o *PutFlowsDatatableRowParams) WithTimeout(timeout time.Duration) *PutFlowsDatatableRowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put flows datatable row params
func (o *PutFlowsDatatableRowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put flows datatable row params
func (o *PutFlowsDatatableRowParams) WithContext(ctx context.Context) *PutFlowsDatatableRowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put flows datatable row params
func (o *PutFlowsDatatableRowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put flows datatable row params
func (o *PutFlowsDatatableRowParams) WithHTTPClient(client *http.Client) *PutFlowsDatatableRowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put flows datatable row params
func (o *PutFlowsDatatableRowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put flows datatable row params
func (o *PutFlowsDatatableRowParams) WithBody(body map[string]interface{}) *PutFlowsDatatableRowParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put flows datatable row params
func (o *PutFlowsDatatableRowParams) SetBody(body map[string]interface{}) {
	o.Body = body
}

// WithDatatableID adds the datatableID to the put flows datatable row params
func (o *PutFlowsDatatableRowParams) WithDatatableID(datatableID string) *PutFlowsDatatableRowParams {
	o.SetDatatableID(datatableID)
	return o
}

// SetDatatableID adds the datatableId to the put flows datatable row params
func (o *PutFlowsDatatableRowParams) SetDatatableID(datatableID string) {
	o.DatatableID = datatableID
}

// WithRowID adds the rowID to the put flows datatable row params
func (o *PutFlowsDatatableRowParams) WithRowID(rowID string) *PutFlowsDatatableRowParams {
	o.SetRowID(rowID)
	return o
}

// SetRowID adds the rowId to the put flows datatable row params
func (o *PutFlowsDatatableRowParams) SetRowID(rowID string) {
	o.RowID = rowID
}

// WriteToRequest writes these params to a swagger request
func (o *PutFlowsDatatableRowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param datatableId
	if err := r.SetPathParam("datatableId", o.DatatableID); err != nil {
		return err
	}

	// path param rowId
	if err := r.SetPathParam("rowId", o.RowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}