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

// NewDeleteOutboundDigitalrulesetParams creates a new DeleteOutboundDigitalrulesetParams object
// with the default values initialized.
func NewDeleteOutboundDigitalrulesetParams() *DeleteOutboundDigitalrulesetParams {
	var ()
	return &DeleteOutboundDigitalrulesetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOutboundDigitalrulesetParamsWithTimeout creates a new DeleteOutboundDigitalrulesetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOutboundDigitalrulesetParamsWithTimeout(timeout time.Duration) *DeleteOutboundDigitalrulesetParams {
	var ()
	return &DeleteOutboundDigitalrulesetParams{

		timeout: timeout,
	}
}

// NewDeleteOutboundDigitalrulesetParamsWithContext creates a new DeleteOutboundDigitalrulesetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOutboundDigitalrulesetParamsWithContext(ctx context.Context) *DeleteOutboundDigitalrulesetParams {
	var ()
	return &DeleteOutboundDigitalrulesetParams{

		Context: ctx,
	}
}

// NewDeleteOutboundDigitalrulesetParamsWithHTTPClient creates a new DeleteOutboundDigitalrulesetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOutboundDigitalrulesetParamsWithHTTPClient(client *http.Client) *DeleteOutboundDigitalrulesetParams {
	var ()
	return &DeleteOutboundDigitalrulesetParams{
		HTTPClient: client,
	}
}

/*DeleteOutboundDigitalrulesetParams contains all the parameters to send to the API endpoint
for the delete outbound digitalruleset operation typically these are written to a http.Request
*/
type DeleteOutboundDigitalrulesetParams struct {

	/*DigitalRuleSetID
	  The Digital Rule Set ID

	*/
	DigitalRuleSetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete outbound digitalruleset params
func (o *DeleteOutboundDigitalrulesetParams) WithTimeout(timeout time.Duration) *DeleteOutboundDigitalrulesetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete outbound digitalruleset params
func (o *DeleteOutboundDigitalrulesetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete outbound digitalruleset params
func (o *DeleteOutboundDigitalrulesetParams) WithContext(ctx context.Context) *DeleteOutboundDigitalrulesetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete outbound digitalruleset params
func (o *DeleteOutboundDigitalrulesetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete outbound digitalruleset params
func (o *DeleteOutboundDigitalrulesetParams) WithHTTPClient(client *http.Client) *DeleteOutboundDigitalrulesetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete outbound digitalruleset params
func (o *DeleteOutboundDigitalrulesetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDigitalRuleSetID adds the digitalRuleSetID to the delete outbound digitalruleset params
func (o *DeleteOutboundDigitalrulesetParams) WithDigitalRuleSetID(digitalRuleSetID string) *DeleteOutboundDigitalrulesetParams {
	o.SetDigitalRuleSetID(digitalRuleSetID)
	return o
}

// SetDigitalRuleSetID adds the digitalRuleSetId to the delete outbound digitalruleset params
func (o *DeleteOutboundDigitalrulesetParams) SetDigitalRuleSetID(digitalRuleSetID string) {
	o.DigitalRuleSetID = digitalRuleSetID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOutboundDigitalrulesetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param digitalRuleSetId
	if err := r.SetPathParam("digitalRuleSetId", o.DigitalRuleSetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
