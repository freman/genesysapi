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

// NewDeleteFlowsParams creates a new DeleteFlowsParams object
// with the default values initialized.
func NewDeleteFlowsParams() *DeleteFlowsParams {
	var ()
	return &DeleteFlowsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFlowsParamsWithTimeout creates a new DeleteFlowsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFlowsParamsWithTimeout(timeout time.Duration) *DeleteFlowsParams {
	var ()
	return &DeleteFlowsParams{

		timeout: timeout,
	}
}

// NewDeleteFlowsParamsWithContext creates a new DeleteFlowsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFlowsParamsWithContext(ctx context.Context) *DeleteFlowsParams {
	var ()
	return &DeleteFlowsParams{

		Context: ctx,
	}
}

// NewDeleteFlowsParamsWithHTTPClient creates a new DeleteFlowsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFlowsParamsWithHTTPClient(client *http.Client) *DeleteFlowsParams {
	var ()
	return &DeleteFlowsParams{
		HTTPClient: client,
	}
}

/*DeleteFlowsParams contains all the parameters to send to the API endpoint
for the delete flows operation typically these are written to a http.Request
*/
type DeleteFlowsParams struct {

	/*ID
	  List of Flow IDs

	*/
	ID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete flows params
func (o *DeleteFlowsParams) WithTimeout(timeout time.Duration) *DeleteFlowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete flows params
func (o *DeleteFlowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete flows params
func (o *DeleteFlowsParams) WithContext(ctx context.Context) *DeleteFlowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete flows params
func (o *DeleteFlowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete flows params
func (o *DeleteFlowsParams) WithHTTPClient(client *http.Client) *DeleteFlowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete flows params
func (o *DeleteFlowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete flows params
func (o *DeleteFlowsParams) WithID(id []string) *DeleteFlowsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete flows params
func (o *DeleteFlowsParams) SetID(id []string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFlowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
