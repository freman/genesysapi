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

// NewDeleteFlowParams creates a new DeleteFlowParams object
// with the default values initialized.
func NewDeleteFlowParams() *DeleteFlowParams {
	var ()
	return &DeleteFlowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFlowParamsWithTimeout creates a new DeleteFlowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFlowParamsWithTimeout(timeout time.Duration) *DeleteFlowParams {
	var ()
	return &DeleteFlowParams{

		timeout: timeout,
	}
}

// NewDeleteFlowParamsWithContext creates a new DeleteFlowParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFlowParamsWithContext(ctx context.Context) *DeleteFlowParams {
	var ()
	return &DeleteFlowParams{

		Context: ctx,
	}
}

// NewDeleteFlowParamsWithHTTPClient creates a new DeleteFlowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFlowParamsWithHTTPClient(client *http.Client) *DeleteFlowParams {
	var ()
	return &DeleteFlowParams{
		HTTPClient: client,
	}
}

/*DeleteFlowParams contains all the parameters to send to the API endpoint
for the delete flow operation typically these are written to a http.Request
*/
type DeleteFlowParams struct {

	/*FlowID
	  Flow ID

	*/
	FlowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete flow params
func (o *DeleteFlowParams) WithTimeout(timeout time.Duration) *DeleteFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete flow params
func (o *DeleteFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete flow params
func (o *DeleteFlowParams) WithContext(ctx context.Context) *DeleteFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete flow params
func (o *DeleteFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete flow params
func (o *DeleteFlowParams) WithHTTPClient(client *http.Client) *DeleteFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete flow params
func (o *DeleteFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFlowID adds the flowID to the delete flow params
func (o *DeleteFlowParams) WithFlowID(flowID string) *DeleteFlowParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the delete flow params
func (o *DeleteFlowParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param flowId
	if err := r.SetPathParam("flowId", o.FlowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
