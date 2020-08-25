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

// NewDeleteOutboundRulesetParams creates a new DeleteOutboundRulesetParams object
// with the default values initialized.
func NewDeleteOutboundRulesetParams() *DeleteOutboundRulesetParams {
	var ()
	return &DeleteOutboundRulesetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOutboundRulesetParamsWithTimeout creates a new DeleteOutboundRulesetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOutboundRulesetParamsWithTimeout(timeout time.Duration) *DeleteOutboundRulesetParams {
	var ()
	return &DeleteOutboundRulesetParams{

		timeout: timeout,
	}
}

// NewDeleteOutboundRulesetParamsWithContext creates a new DeleteOutboundRulesetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOutboundRulesetParamsWithContext(ctx context.Context) *DeleteOutboundRulesetParams {
	var ()
	return &DeleteOutboundRulesetParams{

		Context: ctx,
	}
}

// NewDeleteOutboundRulesetParamsWithHTTPClient creates a new DeleteOutboundRulesetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOutboundRulesetParamsWithHTTPClient(client *http.Client) *DeleteOutboundRulesetParams {
	var ()
	return &DeleteOutboundRulesetParams{
		HTTPClient: client,
	}
}

/*DeleteOutboundRulesetParams contains all the parameters to send to the API endpoint
for the delete outbound ruleset operation typically these are written to a http.Request
*/
type DeleteOutboundRulesetParams struct {

	/*RuleSetID
	  Rule Set ID

	*/
	RuleSetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete outbound ruleset params
func (o *DeleteOutboundRulesetParams) WithTimeout(timeout time.Duration) *DeleteOutboundRulesetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete outbound ruleset params
func (o *DeleteOutboundRulesetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete outbound ruleset params
func (o *DeleteOutboundRulesetParams) WithContext(ctx context.Context) *DeleteOutboundRulesetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete outbound ruleset params
func (o *DeleteOutboundRulesetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete outbound ruleset params
func (o *DeleteOutboundRulesetParams) WithHTTPClient(client *http.Client) *DeleteOutboundRulesetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete outbound ruleset params
func (o *DeleteOutboundRulesetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleSetID adds the ruleSetID to the delete outbound ruleset params
func (o *DeleteOutboundRulesetParams) WithRuleSetID(ruleSetID string) *DeleteOutboundRulesetParams {
	o.SetRuleSetID(ruleSetID)
	return o
}

// SetRuleSetID adds the ruleSetId to the delete outbound ruleset params
func (o *DeleteOutboundRulesetParams) SetRuleSetID(ruleSetID string) {
	o.RuleSetID = ruleSetID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOutboundRulesetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ruleSetId
	if err := r.SetPathParam("ruleSetId", o.RuleSetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}