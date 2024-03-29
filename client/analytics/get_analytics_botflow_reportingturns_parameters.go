// Code generated by go-swagger; DO NOT EDIT.

package analytics

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

// NewGetAnalyticsBotflowReportingturnsParams creates a new GetAnalyticsBotflowReportingturnsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAnalyticsBotflowReportingturnsParams() *GetAnalyticsBotflowReportingturnsParams {
	return &GetAnalyticsBotflowReportingturnsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnalyticsBotflowReportingturnsParamsWithTimeout creates a new GetAnalyticsBotflowReportingturnsParams object
// with the ability to set a timeout on a request.
func NewGetAnalyticsBotflowReportingturnsParamsWithTimeout(timeout time.Duration) *GetAnalyticsBotflowReportingturnsParams {
	return &GetAnalyticsBotflowReportingturnsParams{
		timeout: timeout,
	}
}

// NewGetAnalyticsBotflowReportingturnsParamsWithContext creates a new GetAnalyticsBotflowReportingturnsParams object
// with the ability to set a context for a request.
func NewGetAnalyticsBotflowReportingturnsParamsWithContext(ctx context.Context) *GetAnalyticsBotflowReportingturnsParams {
	return &GetAnalyticsBotflowReportingturnsParams{
		Context: ctx,
	}
}

// NewGetAnalyticsBotflowReportingturnsParamsWithHTTPClient creates a new GetAnalyticsBotflowReportingturnsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAnalyticsBotflowReportingturnsParamsWithHTTPClient(client *http.Client) *GetAnalyticsBotflowReportingturnsParams {
	return &GetAnalyticsBotflowReportingturnsParams{
		HTTPClient: client,
	}
}

/*
GetAnalyticsBotflowReportingturnsParams contains all the parameters to send to the API endpoint

	for the get analytics botflow reportingturns operation.

	Typically these are written to a http.Request.
*/
type GetAnalyticsBotflowReportingturnsParams struct {

	/* ActionID.

	   Optional action ID to get the reporting turns associated to a particular flow action
	*/
	ActionID *string

	/* After.

	   The cursor that points to the ID of the last item in the list of entities that has been returned.
	*/
	After *string

	/* AskActionResults.

	   Optional case-insensitive comma separated list of ask action results to filter the reporting turns.
	*/
	AskActionResults *string

	/* BotFlowID.

	   ID of the bot flow.
	*/
	BotFlowID string

	/* Language.

	   Optional language code to get the reporting turns for a particular language
	*/
	Language *string

	/* PageSize.

	   Max number of entities to return. Maximum of 250

	   Default: "50"
	*/
	PageSize *string

	/* SessionID.

	   Optional session ID to get the reporting turns for a particular session
	*/
	SessionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get analytics botflow reportingturns params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnalyticsBotflowReportingturnsParams) WithDefaults() *GetAnalyticsBotflowReportingturnsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get analytics botflow reportingturns params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnalyticsBotflowReportingturnsParams) SetDefaults() {
	var (
		pageSizeDefault = string("50")
	)

	val := GetAnalyticsBotflowReportingturnsParams{
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) WithTimeout(timeout time.Duration) *GetAnalyticsBotflowReportingturnsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) WithContext(ctx context.Context) *GetAnalyticsBotflowReportingturnsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) WithHTTPClient(client *http.Client) *GetAnalyticsBotflowReportingturnsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) WithActionID(actionID *string) *GetAnalyticsBotflowReportingturnsParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) SetActionID(actionID *string) {
	o.ActionID = actionID
}

// WithAfter adds the after to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) WithAfter(after *string) *GetAnalyticsBotflowReportingturnsParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) SetAfter(after *string) {
	o.After = after
}

// WithAskActionResults adds the askActionResults to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) WithAskActionResults(askActionResults *string) *GetAnalyticsBotflowReportingturnsParams {
	o.SetAskActionResults(askActionResults)
	return o
}

// SetAskActionResults adds the askActionResults to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) SetAskActionResults(askActionResults *string) {
	o.AskActionResults = askActionResults
}

// WithBotFlowID adds the botFlowID to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) WithBotFlowID(botFlowID string) *GetAnalyticsBotflowReportingturnsParams {
	o.SetBotFlowID(botFlowID)
	return o
}

// SetBotFlowID adds the botFlowId to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) SetBotFlowID(botFlowID string) {
	o.BotFlowID = botFlowID
}

// WithLanguage adds the language to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) WithLanguage(language *string) *GetAnalyticsBotflowReportingturnsParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) SetLanguage(language *string) {
	o.Language = language
}

// WithPageSize adds the pageSize to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) WithPageSize(pageSize *string) *GetAnalyticsBotflowReportingturnsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WithSessionID adds the sessionID to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) WithSessionID(sessionID *string) *GetAnalyticsBotflowReportingturnsParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the get analytics botflow reportingturns params
func (o *GetAnalyticsBotflowReportingturnsParams) SetSessionID(sessionID *string) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnalyticsBotflowReportingturnsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActionID != nil {

		// query param actionId
		var qrActionID string

		if o.ActionID != nil {
			qrActionID = *o.ActionID
		}
		qActionID := qrActionID
		if qActionID != "" {

			if err := r.SetQueryParam("actionId", qActionID); err != nil {
				return err
			}
		}
	}

	if o.After != nil {

		// query param after
		var qrAfter string

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if o.AskActionResults != nil {

		// query param askActionResults
		var qrAskActionResults string

		if o.AskActionResults != nil {
			qrAskActionResults = *o.AskActionResults
		}
		qAskActionResults := qrAskActionResults
		if qAskActionResults != "" {

			if err := r.SetQueryParam("askActionResults", qAskActionResults); err != nil {
				return err
			}
		}
	}

	// path param botFlowId
	if err := r.SetPathParam("botFlowId", o.BotFlowID); err != nil {
		return err
	}

	if o.Language != nil {

		// query param language
		var qrLanguage string

		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {

			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize string

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := qrPageSize
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.SessionID != nil {

		// query param sessionId
		var qrSessionID string

		if o.SessionID != nil {
			qrSessionID = *o.SessionID
		}
		qSessionID := qrSessionID
		if qSessionID != "" {

			if err := r.SetQueryParam("sessionId", qSessionID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
