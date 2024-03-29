// Code generated by go-swagger; DO NOT EDIT.

package journey

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

// NewGetJourneyOutcomesPredictorsParams creates a new GetJourneyOutcomesPredictorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetJourneyOutcomesPredictorsParams() *GetJourneyOutcomesPredictorsParams {
	return &GetJourneyOutcomesPredictorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetJourneyOutcomesPredictorsParamsWithTimeout creates a new GetJourneyOutcomesPredictorsParams object
// with the ability to set a timeout on a request.
func NewGetJourneyOutcomesPredictorsParamsWithTimeout(timeout time.Duration) *GetJourneyOutcomesPredictorsParams {
	return &GetJourneyOutcomesPredictorsParams{
		timeout: timeout,
	}
}

// NewGetJourneyOutcomesPredictorsParamsWithContext creates a new GetJourneyOutcomesPredictorsParams object
// with the ability to set a context for a request.
func NewGetJourneyOutcomesPredictorsParamsWithContext(ctx context.Context) *GetJourneyOutcomesPredictorsParams {
	return &GetJourneyOutcomesPredictorsParams{
		Context: ctx,
	}
}

// NewGetJourneyOutcomesPredictorsParamsWithHTTPClient creates a new GetJourneyOutcomesPredictorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetJourneyOutcomesPredictorsParamsWithHTTPClient(client *http.Client) *GetJourneyOutcomesPredictorsParams {
	return &GetJourneyOutcomesPredictorsParams{
		HTTPClient: client,
	}
}

/*
GetJourneyOutcomesPredictorsParams contains all the parameters to send to the API endpoint

	for the get journey outcomes predictors operation.

	Typically these are written to a http.Request.
*/
type GetJourneyOutcomesPredictorsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get journey outcomes predictors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJourneyOutcomesPredictorsParams) WithDefaults() *GetJourneyOutcomesPredictorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get journey outcomes predictors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJourneyOutcomesPredictorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get journey outcomes predictors params
func (o *GetJourneyOutcomesPredictorsParams) WithTimeout(timeout time.Duration) *GetJourneyOutcomesPredictorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get journey outcomes predictors params
func (o *GetJourneyOutcomesPredictorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get journey outcomes predictors params
func (o *GetJourneyOutcomesPredictorsParams) WithContext(ctx context.Context) *GetJourneyOutcomesPredictorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get journey outcomes predictors params
func (o *GetJourneyOutcomesPredictorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get journey outcomes predictors params
func (o *GetJourneyOutcomesPredictorsParams) WithHTTPClient(client *http.Client) *GetJourneyOutcomesPredictorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get journey outcomes predictors params
func (o *GetJourneyOutcomesPredictorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetJourneyOutcomesPredictorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
