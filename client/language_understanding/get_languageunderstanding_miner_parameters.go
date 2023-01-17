// Code generated by go-swagger; DO NOT EDIT.

package language_understanding

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

// NewGetLanguageunderstandingMinerParams creates a new GetLanguageunderstandingMinerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLanguageunderstandingMinerParams() *GetLanguageunderstandingMinerParams {
	return &GetLanguageunderstandingMinerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLanguageunderstandingMinerParamsWithTimeout creates a new GetLanguageunderstandingMinerParams object
// with the ability to set a timeout on a request.
func NewGetLanguageunderstandingMinerParamsWithTimeout(timeout time.Duration) *GetLanguageunderstandingMinerParams {
	return &GetLanguageunderstandingMinerParams{
		timeout: timeout,
	}
}

// NewGetLanguageunderstandingMinerParamsWithContext creates a new GetLanguageunderstandingMinerParams object
// with the ability to set a context for a request.
func NewGetLanguageunderstandingMinerParamsWithContext(ctx context.Context) *GetLanguageunderstandingMinerParams {
	return &GetLanguageunderstandingMinerParams{
		Context: ctx,
	}
}

// NewGetLanguageunderstandingMinerParamsWithHTTPClient creates a new GetLanguageunderstandingMinerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLanguageunderstandingMinerParamsWithHTTPClient(client *http.Client) *GetLanguageunderstandingMinerParams {
	return &GetLanguageunderstandingMinerParams{
		HTTPClient: client,
	}
}

/*
GetLanguageunderstandingMinerParams contains all the parameters to send to the API endpoint

	for the get languageunderstanding miner operation.

	Typically these are written to a http.Request.
*/
type GetLanguageunderstandingMinerParams struct {

	/* MinerID.

	   Miner ID
	*/
	MinerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get languageunderstanding miner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLanguageunderstandingMinerParams) WithDefaults() *GetLanguageunderstandingMinerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get languageunderstanding miner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLanguageunderstandingMinerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get languageunderstanding miner params
func (o *GetLanguageunderstandingMinerParams) WithTimeout(timeout time.Duration) *GetLanguageunderstandingMinerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get languageunderstanding miner params
func (o *GetLanguageunderstandingMinerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get languageunderstanding miner params
func (o *GetLanguageunderstandingMinerParams) WithContext(ctx context.Context) *GetLanguageunderstandingMinerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get languageunderstanding miner params
func (o *GetLanguageunderstandingMinerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get languageunderstanding miner params
func (o *GetLanguageunderstandingMinerParams) WithHTTPClient(client *http.Client) *GetLanguageunderstandingMinerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get languageunderstanding miner params
func (o *GetLanguageunderstandingMinerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMinerID adds the minerID to the get languageunderstanding miner params
func (o *GetLanguageunderstandingMinerParams) WithMinerID(minerID string) *GetLanguageunderstandingMinerParams {
	o.SetMinerID(minerID)
	return o
}

// SetMinerID adds the minerId to the get languageunderstanding miner params
func (o *GetLanguageunderstandingMinerParams) SetMinerID(minerID string) {
	o.MinerID = minerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLanguageunderstandingMinerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param minerId
	if err := r.SetPathParam("minerId", o.MinerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
