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

// NewDeleteFlowsDatatableParams creates a new DeleteFlowsDatatableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteFlowsDatatableParams() *DeleteFlowsDatatableParams {
	return &DeleteFlowsDatatableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFlowsDatatableParamsWithTimeout creates a new DeleteFlowsDatatableParams object
// with the ability to set a timeout on a request.
func NewDeleteFlowsDatatableParamsWithTimeout(timeout time.Duration) *DeleteFlowsDatatableParams {
	return &DeleteFlowsDatatableParams{
		timeout: timeout,
	}
}

// NewDeleteFlowsDatatableParamsWithContext creates a new DeleteFlowsDatatableParams object
// with the ability to set a context for a request.
func NewDeleteFlowsDatatableParamsWithContext(ctx context.Context) *DeleteFlowsDatatableParams {
	return &DeleteFlowsDatatableParams{
		Context: ctx,
	}
}

// NewDeleteFlowsDatatableParamsWithHTTPClient creates a new DeleteFlowsDatatableParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteFlowsDatatableParamsWithHTTPClient(client *http.Client) *DeleteFlowsDatatableParams {
	return &DeleteFlowsDatatableParams{
		HTTPClient: client,
	}
}

/*
DeleteFlowsDatatableParams contains all the parameters to send to the API endpoint

	for the delete flows datatable operation.

	Typically these are written to a http.Request.
*/
type DeleteFlowsDatatableParams struct {

	/* DatatableID.

	   id of datatable
	*/
	DatatableID string

	/* Force.

	   force delete, even if in use
	*/
	Force *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete flows datatable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFlowsDatatableParams) WithDefaults() *DeleteFlowsDatatableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete flows datatable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteFlowsDatatableParams) SetDefaults() {
	var (
		forceDefault = bool(false)
	)

	val := DeleteFlowsDatatableParams{
		Force: &forceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete flows datatable params
func (o *DeleteFlowsDatatableParams) WithTimeout(timeout time.Duration) *DeleteFlowsDatatableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete flows datatable params
func (o *DeleteFlowsDatatableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete flows datatable params
func (o *DeleteFlowsDatatableParams) WithContext(ctx context.Context) *DeleteFlowsDatatableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete flows datatable params
func (o *DeleteFlowsDatatableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete flows datatable params
func (o *DeleteFlowsDatatableParams) WithHTTPClient(client *http.Client) *DeleteFlowsDatatableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete flows datatable params
func (o *DeleteFlowsDatatableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatatableID adds the datatableID to the delete flows datatable params
func (o *DeleteFlowsDatatableParams) WithDatatableID(datatableID string) *DeleteFlowsDatatableParams {
	o.SetDatatableID(datatableID)
	return o
}

// SetDatatableID adds the datatableId to the delete flows datatable params
func (o *DeleteFlowsDatatableParams) SetDatatableID(datatableID string) {
	o.DatatableID = datatableID
}

// WithForce adds the force to the delete flows datatable params
func (o *DeleteFlowsDatatableParams) WithForce(force *bool) *DeleteFlowsDatatableParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the delete flows datatable params
func (o *DeleteFlowsDatatableParams) SetForce(force *bool) {
	o.Force = force
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFlowsDatatableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datatableId
	if err := r.SetPathParam("datatableId", o.DatatableID); err != nil {
		return err
	}

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
