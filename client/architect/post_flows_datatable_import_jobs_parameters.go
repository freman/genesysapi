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

	"github.com/freman/genesysapi/models"
)

// NewPostFlowsDatatableImportJobsParams creates a new PostFlowsDatatableImportJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostFlowsDatatableImportJobsParams() *PostFlowsDatatableImportJobsParams {
	return &PostFlowsDatatableImportJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostFlowsDatatableImportJobsParamsWithTimeout creates a new PostFlowsDatatableImportJobsParams object
// with the ability to set a timeout on a request.
func NewPostFlowsDatatableImportJobsParamsWithTimeout(timeout time.Duration) *PostFlowsDatatableImportJobsParams {
	return &PostFlowsDatatableImportJobsParams{
		timeout: timeout,
	}
}

// NewPostFlowsDatatableImportJobsParamsWithContext creates a new PostFlowsDatatableImportJobsParams object
// with the ability to set a context for a request.
func NewPostFlowsDatatableImportJobsParamsWithContext(ctx context.Context) *PostFlowsDatatableImportJobsParams {
	return &PostFlowsDatatableImportJobsParams{
		Context: ctx,
	}
}

// NewPostFlowsDatatableImportJobsParamsWithHTTPClient creates a new PostFlowsDatatableImportJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostFlowsDatatableImportJobsParamsWithHTTPClient(client *http.Client) *PostFlowsDatatableImportJobsParams {
	return &PostFlowsDatatableImportJobsParams{
		HTTPClient: client,
	}
}

/*
PostFlowsDatatableImportJobsParams contains all the parameters to send to the API endpoint

	for the post flows datatable import jobs operation.

	Typically these are written to a http.Request.
*/
type PostFlowsDatatableImportJobsParams struct {

	/* Body.

	   import job information
	*/
	Body *models.DataTableImportJob

	/* DatatableID.

	   id of datatable
	*/
	DatatableID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post flows datatable import jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostFlowsDatatableImportJobsParams) WithDefaults() *PostFlowsDatatableImportJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post flows datatable import jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostFlowsDatatableImportJobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post flows datatable import jobs params
func (o *PostFlowsDatatableImportJobsParams) WithTimeout(timeout time.Duration) *PostFlowsDatatableImportJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post flows datatable import jobs params
func (o *PostFlowsDatatableImportJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post flows datatable import jobs params
func (o *PostFlowsDatatableImportJobsParams) WithContext(ctx context.Context) *PostFlowsDatatableImportJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post flows datatable import jobs params
func (o *PostFlowsDatatableImportJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post flows datatable import jobs params
func (o *PostFlowsDatatableImportJobsParams) WithHTTPClient(client *http.Client) *PostFlowsDatatableImportJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post flows datatable import jobs params
func (o *PostFlowsDatatableImportJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post flows datatable import jobs params
func (o *PostFlowsDatatableImportJobsParams) WithBody(body *models.DataTableImportJob) *PostFlowsDatatableImportJobsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post flows datatable import jobs params
func (o *PostFlowsDatatableImportJobsParams) SetBody(body *models.DataTableImportJob) {
	o.Body = body
}

// WithDatatableID adds the datatableID to the post flows datatable import jobs params
func (o *PostFlowsDatatableImportJobsParams) WithDatatableID(datatableID string) *PostFlowsDatatableImportJobsParams {
	o.SetDatatableID(datatableID)
	return o
}

// SetDatatableID adds the datatableId to the post flows datatable import jobs params
func (o *PostFlowsDatatableImportJobsParams) SetDatatableID(datatableID string) {
	o.DatatableID = datatableID
}

// WriteToRequest writes these params to a swagger request
func (o *PostFlowsDatatableImportJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
