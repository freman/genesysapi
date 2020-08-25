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
)

// NewGetAlertingInteractionstatsRulesParams creates a new GetAlertingInteractionstatsRulesParams object
// with the default values initialized.
func NewGetAlertingInteractionstatsRulesParams() *GetAlertingInteractionstatsRulesParams {
	var ()
	return &GetAlertingInteractionstatsRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertingInteractionstatsRulesParamsWithTimeout creates a new GetAlertingInteractionstatsRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAlertingInteractionstatsRulesParamsWithTimeout(timeout time.Duration) *GetAlertingInteractionstatsRulesParams {
	var ()
	return &GetAlertingInteractionstatsRulesParams{

		timeout: timeout,
	}
}

// NewGetAlertingInteractionstatsRulesParamsWithContext creates a new GetAlertingInteractionstatsRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAlertingInteractionstatsRulesParamsWithContext(ctx context.Context) *GetAlertingInteractionstatsRulesParams {
	var ()
	return &GetAlertingInteractionstatsRulesParams{

		Context: ctx,
	}
}

// NewGetAlertingInteractionstatsRulesParamsWithHTTPClient creates a new GetAlertingInteractionstatsRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAlertingInteractionstatsRulesParamsWithHTTPClient(client *http.Client) *GetAlertingInteractionstatsRulesParams {
	var ()
	return &GetAlertingInteractionstatsRulesParams{
		HTTPClient: client,
	}
}

/*GetAlertingInteractionstatsRulesParams contains all the parameters to send to the API endpoint
for the get alerting interactionstats rules operation typically these are written to a http.Request
*/
type GetAlertingInteractionstatsRulesParams struct {

	/*Expand
	  Which fields, if any, to expand

	*/
	Expand []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get alerting interactionstats rules params
func (o *GetAlertingInteractionstatsRulesParams) WithTimeout(timeout time.Duration) *GetAlertingInteractionstatsRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alerting interactionstats rules params
func (o *GetAlertingInteractionstatsRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alerting interactionstats rules params
func (o *GetAlertingInteractionstatsRulesParams) WithContext(ctx context.Context) *GetAlertingInteractionstatsRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alerting interactionstats rules params
func (o *GetAlertingInteractionstatsRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alerting interactionstats rules params
func (o *GetAlertingInteractionstatsRulesParams) WithHTTPClient(client *http.Client) *GetAlertingInteractionstatsRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alerting interactionstats rules params
func (o *GetAlertingInteractionstatsRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get alerting interactionstats rules params
func (o *GetAlertingInteractionstatsRulesParams) WithExpand(expand []string) *GetAlertingInteractionstatsRulesParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get alerting interactionstats rules params
func (o *GetAlertingInteractionstatsRulesParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertingInteractionstatsRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}