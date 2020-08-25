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

// NewGetOutboundRulesetParams creates a new GetOutboundRulesetParams object
// with the default values initialized.
func NewGetOutboundRulesetParams() *GetOutboundRulesetParams {
	var ()
	return &GetOutboundRulesetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundRulesetParamsWithTimeout creates a new GetOutboundRulesetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutboundRulesetParamsWithTimeout(timeout time.Duration) *GetOutboundRulesetParams {
	var ()
	return &GetOutboundRulesetParams{

		timeout: timeout,
	}
}

// NewGetOutboundRulesetParamsWithContext creates a new GetOutboundRulesetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutboundRulesetParamsWithContext(ctx context.Context) *GetOutboundRulesetParams {
	var ()
	return &GetOutboundRulesetParams{

		Context: ctx,
	}
}

// NewGetOutboundRulesetParamsWithHTTPClient creates a new GetOutboundRulesetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutboundRulesetParamsWithHTTPClient(client *http.Client) *GetOutboundRulesetParams {
	var ()
	return &GetOutboundRulesetParams{
		HTTPClient: client,
	}
}

/*GetOutboundRulesetParams contains all the parameters to send to the API endpoint
for the get outbound ruleset operation typically these are written to a http.Request
*/
type GetOutboundRulesetParams struct {

	/*RuleSetID
	  Rule Set ID

	*/
	RuleSetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get outbound ruleset params
func (o *GetOutboundRulesetParams) WithTimeout(timeout time.Duration) *GetOutboundRulesetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound ruleset params
func (o *GetOutboundRulesetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound ruleset params
func (o *GetOutboundRulesetParams) WithContext(ctx context.Context) *GetOutboundRulesetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound ruleset params
func (o *GetOutboundRulesetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound ruleset params
func (o *GetOutboundRulesetParams) WithHTTPClient(client *http.Client) *GetOutboundRulesetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound ruleset params
func (o *GetOutboundRulesetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleSetID adds the ruleSetID to the get outbound ruleset params
func (o *GetOutboundRulesetParams) WithRuleSetID(ruleSetID string) *GetOutboundRulesetParams {
	o.SetRuleSetID(ruleSetID)
	return o
}

// SetRuleSetID adds the ruleSetId to the get outbound ruleset params
func (o *GetOutboundRulesetParams) SetRuleSetID(ruleSetID string) {
	o.RuleSetID = ruleSetID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundRulesetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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