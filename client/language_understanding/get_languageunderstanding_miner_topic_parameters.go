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

// NewGetLanguageunderstandingMinerTopicParams creates a new GetLanguageunderstandingMinerTopicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLanguageunderstandingMinerTopicParams() *GetLanguageunderstandingMinerTopicParams {
	return &GetLanguageunderstandingMinerTopicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLanguageunderstandingMinerTopicParamsWithTimeout creates a new GetLanguageunderstandingMinerTopicParams object
// with the ability to set a timeout on a request.
func NewGetLanguageunderstandingMinerTopicParamsWithTimeout(timeout time.Duration) *GetLanguageunderstandingMinerTopicParams {
	return &GetLanguageunderstandingMinerTopicParams{
		timeout: timeout,
	}
}

// NewGetLanguageunderstandingMinerTopicParamsWithContext creates a new GetLanguageunderstandingMinerTopicParams object
// with the ability to set a context for a request.
func NewGetLanguageunderstandingMinerTopicParamsWithContext(ctx context.Context) *GetLanguageunderstandingMinerTopicParams {
	return &GetLanguageunderstandingMinerTopicParams{
		Context: ctx,
	}
}

// NewGetLanguageunderstandingMinerTopicParamsWithHTTPClient creates a new GetLanguageunderstandingMinerTopicParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLanguageunderstandingMinerTopicParamsWithHTTPClient(client *http.Client) *GetLanguageunderstandingMinerTopicParams {
	return &GetLanguageunderstandingMinerTopicParams{
		HTTPClient: client,
	}
}

/*
GetLanguageunderstandingMinerTopicParams contains all the parameters to send to the API endpoint

	for the get languageunderstanding miner topic operation.

	Typically these are written to a http.Request.
*/
type GetLanguageunderstandingMinerTopicParams struct {

	/* Expand.

	   Option to fetch phrases
	*/
	Expand *string

	/* MinerID.

	   Miner ID
	*/
	MinerID string

	/* TopicID.

	   The ID of the topic to be retrieved.
	*/
	TopicID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get languageunderstanding miner topic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLanguageunderstandingMinerTopicParams) WithDefaults() *GetLanguageunderstandingMinerTopicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get languageunderstanding miner topic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLanguageunderstandingMinerTopicParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get languageunderstanding miner topic params
func (o *GetLanguageunderstandingMinerTopicParams) WithTimeout(timeout time.Duration) *GetLanguageunderstandingMinerTopicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get languageunderstanding miner topic params
func (o *GetLanguageunderstandingMinerTopicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get languageunderstanding miner topic params
func (o *GetLanguageunderstandingMinerTopicParams) WithContext(ctx context.Context) *GetLanguageunderstandingMinerTopicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get languageunderstanding miner topic params
func (o *GetLanguageunderstandingMinerTopicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get languageunderstanding miner topic params
func (o *GetLanguageunderstandingMinerTopicParams) WithHTTPClient(client *http.Client) *GetLanguageunderstandingMinerTopicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get languageunderstanding miner topic params
func (o *GetLanguageunderstandingMinerTopicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get languageunderstanding miner topic params
func (o *GetLanguageunderstandingMinerTopicParams) WithExpand(expand *string) *GetLanguageunderstandingMinerTopicParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get languageunderstanding miner topic params
func (o *GetLanguageunderstandingMinerTopicParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithMinerID adds the minerID to the get languageunderstanding miner topic params
func (o *GetLanguageunderstandingMinerTopicParams) WithMinerID(minerID string) *GetLanguageunderstandingMinerTopicParams {
	o.SetMinerID(minerID)
	return o
}

// SetMinerID adds the minerId to the get languageunderstanding miner topic params
func (o *GetLanguageunderstandingMinerTopicParams) SetMinerID(minerID string) {
	o.MinerID = minerID
}

// WithTopicID adds the topicID to the get languageunderstanding miner topic params
func (o *GetLanguageunderstandingMinerTopicParams) WithTopicID(topicID string) *GetLanguageunderstandingMinerTopicParams {
	o.SetTopicID(topicID)
	return o
}

// SetTopicID adds the topicId to the get languageunderstanding miner topic params
func (o *GetLanguageunderstandingMinerTopicParams) SetTopicID(topicID string) {
	o.TopicID = topicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLanguageunderstandingMinerTopicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// query param expand
		var qrExpand string

		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {

			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}
	}

	// path param minerId
	if err := r.SetPathParam("minerId", o.MinerID); err != nil {
		return err
	}

	// path param topicId
	if err := r.SetPathParam("topicId", o.TopicID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
