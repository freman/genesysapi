// Code generated by go-swagger; DO NOT EDIT.

package alerting

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

	"github.com/freman/genesysapi/models"
)

// NewPutAlertingInteractionstatsRuleParams creates a new PutAlertingInteractionstatsRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutAlertingInteractionstatsRuleParams() *PutAlertingInteractionstatsRuleParams {
	return &PutAlertingInteractionstatsRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutAlertingInteractionstatsRuleParamsWithTimeout creates a new PutAlertingInteractionstatsRuleParams object
// with the ability to set a timeout on a request.
func NewPutAlertingInteractionstatsRuleParamsWithTimeout(timeout time.Duration) *PutAlertingInteractionstatsRuleParams {
	return &PutAlertingInteractionstatsRuleParams{
		timeout: timeout,
	}
}

// NewPutAlertingInteractionstatsRuleParamsWithContext creates a new PutAlertingInteractionstatsRuleParams object
// with the ability to set a context for a request.
func NewPutAlertingInteractionstatsRuleParamsWithContext(ctx context.Context) *PutAlertingInteractionstatsRuleParams {
	return &PutAlertingInteractionstatsRuleParams{
		Context: ctx,
	}
}

// NewPutAlertingInteractionstatsRuleParamsWithHTTPClient creates a new PutAlertingInteractionstatsRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutAlertingInteractionstatsRuleParamsWithHTTPClient(client *http.Client) *PutAlertingInteractionstatsRuleParams {
	return &PutAlertingInteractionstatsRuleParams{
		HTTPClient: client,
	}
}

/*
PutAlertingInteractionstatsRuleParams contains all the parameters to send to the API endpoint

	for the put alerting interactionstats rule operation.

	Typically these are written to a http.Request.
*/
type PutAlertingInteractionstatsRuleParams struct {

	/* Body.

	   AlertingRule
	*/
	Body *models.InteractionStatsRule

	/* Expand.

	   Which fields, if any, to expand
	*/
	Expand []string

	/* RuleID.

	   Rule ID
	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put alerting interactionstats rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAlertingInteractionstatsRuleParams) WithDefaults() *PutAlertingInteractionstatsRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put alerting interactionstats rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAlertingInteractionstatsRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put alerting interactionstats rule params
func (o *PutAlertingInteractionstatsRuleParams) WithTimeout(timeout time.Duration) *PutAlertingInteractionstatsRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put alerting interactionstats rule params
func (o *PutAlertingInteractionstatsRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put alerting interactionstats rule params
func (o *PutAlertingInteractionstatsRuleParams) WithContext(ctx context.Context) *PutAlertingInteractionstatsRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put alerting interactionstats rule params
func (o *PutAlertingInteractionstatsRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put alerting interactionstats rule params
func (o *PutAlertingInteractionstatsRuleParams) WithHTTPClient(client *http.Client) *PutAlertingInteractionstatsRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put alerting interactionstats rule params
func (o *PutAlertingInteractionstatsRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put alerting interactionstats rule params
func (o *PutAlertingInteractionstatsRuleParams) WithBody(body *models.InteractionStatsRule) *PutAlertingInteractionstatsRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put alerting interactionstats rule params
func (o *PutAlertingInteractionstatsRuleParams) SetBody(body *models.InteractionStatsRule) {
	o.Body = body
}

// WithExpand adds the expand to the put alerting interactionstats rule params
func (o *PutAlertingInteractionstatsRuleParams) WithExpand(expand []string) *PutAlertingInteractionstatsRuleParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the put alerting interactionstats rule params
func (o *PutAlertingInteractionstatsRuleParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithRuleID adds the ruleID to the put alerting interactionstats rule params
func (o *PutAlertingInteractionstatsRuleParams) WithRuleID(ruleID string) *PutAlertingInteractionstatsRuleParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the put alerting interactionstats rule params
func (o *PutAlertingInteractionstatsRuleParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAlertingInteractionstatsRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPutAlertingInteractionstatsRule binds the parameter expand
func (o *PutAlertingInteractionstatsRuleParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}
