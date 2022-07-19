// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewGetRoutingPredictorsKeyperformanceindicatorsParams creates a new GetRoutingPredictorsKeyperformanceindicatorsParams object
// with the default values initialized.
func NewGetRoutingPredictorsKeyperformanceindicatorsParams() *GetRoutingPredictorsKeyperformanceindicatorsParams {
	var ()
	return &GetRoutingPredictorsKeyperformanceindicatorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingPredictorsKeyperformanceindicatorsParamsWithTimeout creates a new GetRoutingPredictorsKeyperformanceindicatorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoutingPredictorsKeyperformanceindicatorsParamsWithTimeout(timeout time.Duration) *GetRoutingPredictorsKeyperformanceindicatorsParams {
	var ()
	return &GetRoutingPredictorsKeyperformanceindicatorsParams{

		timeout: timeout,
	}
}

// NewGetRoutingPredictorsKeyperformanceindicatorsParamsWithContext creates a new GetRoutingPredictorsKeyperformanceindicatorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoutingPredictorsKeyperformanceindicatorsParamsWithContext(ctx context.Context) *GetRoutingPredictorsKeyperformanceindicatorsParams {
	var ()
	return &GetRoutingPredictorsKeyperformanceindicatorsParams{

		Context: ctx,
	}
}

// NewGetRoutingPredictorsKeyperformanceindicatorsParamsWithHTTPClient creates a new GetRoutingPredictorsKeyperformanceindicatorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoutingPredictorsKeyperformanceindicatorsParamsWithHTTPClient(client *http.Client) *GetRoutingPredictorsKeyperformanceindicatorsParams {
	var ()
	return &GetRoutingPredictorsKeyperformanceindicatorsParams{
		HTTPClient: client,
	}
}

/*GetRoutingPredictorsKeyperformanceindicatorsParams contains all the parameters to send to the API endpoint
for the get routing predictors keyperformanceindicators operation typically these are written to a http.Request
*/
type GetRoutingPredictorsKeyperformanceindicatorsParams struct {

	/*KpiGroup
	  The Group of Key Performance Indicators to return

	*/
	KpiGroup *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routing predictors keyperformanceindicators params
func (o *GetRoutingPredictorsKeyperformanceindicatorsParams) WithTimeout(timeout time.Duration) *GetRoutingPredictorsKeyperformanceindicatorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing predictors keyperformanceindicators params
func (o *GetRoutingPredictorsKeyperformanceindicatorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing predictors keyperformanceindicators params
func (o *GetRoutingPredictorsKeyperformanceindicatorsParams) WithContext(ctx context.Context) *GetRoutingPredictorsKeyperformanceindicatorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing predictors keyperformanceindicators params
func (o *GetRoutingPredictorsKeyperformanceindicatorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing predictors keyperformanceindicators params
func (o *GetRoutingPredictorsKeyperformanceindicatorsParams) WithHTTPClient(client *http.Client) *GetRoutingPredictorsKeyperformanceindicatorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing predictors keyperformanceindicators params
func (o *GetRoutingPredictorsKeyperformanceindicatorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKpiGroup adds the kpiGroup to the get routing predictors keyperformanceindicators params
func (o *GetRoutingPredictorsKeyperformanceindicatorsParams) WithKpiGroup(kpiGroup *string) *GetRoutingPredictorsKeyperformanceindicatorsParams {
	o.SetKpiGroup(kpiGroup)
	return o
}

// SetKpiGroup adds the kpiGroup to the get routing predictors keyperformanceindicators params
func (o *GetRoutingPredictorsKeyperformanceindicatorsParams) SetKpiGroup(kpiGroup *string) {
	o.KpiGroup = kpiGroup
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingPredictorsKeyperformanceindicatorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.KpiGroup != nil {

		// query param kpiGroup
		var qrKpiGroup string
		if o.KpiGroup != nil {
			qrKpiGroup = *o.KpiGroup
		}
		qKpiGroup := qrKpiGroup
		if qKpiGroup != "" {
			if err := r.SetQueryParam("kpiGroup", qKpiGroup); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
