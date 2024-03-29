// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetIntegrationsSpeechLexBotBotIDAliasesParams creates a new GetIntegrationsSpeechLexBotBotIDAliasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIntegrationsSpeechLexBotBotIDAliasesParams() *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	return &GetIntegrationsSpeechLexBotBotIDAliasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationsSpeechLexBotBotIDAliasesParamsWithTimeout creates a new GetIntegrationsSpeechLexBotBotIDAliasesParams object
// with the ability to set a timeout on a request.
func NewGetIntegrationsSpeechLexBotBotIDAliasesParamsWithTimeout(timeout time.Duration) *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	return &GetIntegrationsSpeechLexBotBotIDAliasesParams{
		timeout: timeout,
	}
}

// NewGetIntegrationsSpeechLexBotBotIDAliasesParamsWithContext creates a new GetIntegrationsSpeechLexBotBotIDAliasesParams object
// with the ability to set a context for a request.
func NewGetIntegrationsSpeechLexBotBotIDAliasesParamsWithContext(ctx context.Context) *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	return &GetIntegrationsSpeechLexBotBotIDAliasesParams{
		Context: ctx,
	}
}

// NewGetIntegrationsSpeechLexBotBotIDAliasesParamsWithHTTPClient creates a new GetIntegrationsSpeechLexBotBotIDAliasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIntegrationsSpeechLexBotBotIDAliasesParamsWithHTTPClient(client *http.Client) *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	return &GetIntegrationsSpeechLexBotBotIDAliasesParams{
		HTTPClient: client,
	}
}

/*
GetIntegrationsSpeechLexBotBotIDAliasesParams contains all the parameters to send to the API endpoint

	for the get integrations speech lex bot bot Id aliases operation.

	Typically these are written to a http.Request.
*/
type GetIntegrationsSpeechLexBotBotIDAliasesParams struct {

	/* BotID.

	   The bot ID
	*/
	BotID string

	/* Name.

	   Filter on alias name
	*/
	Name *string

	/* PageNumber.

	   Page number

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   Page size

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	/* Status.

	   Filter on alias status
	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get integrations speech lex bot bot Id aliases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) WithDefaults() *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get integrations speech lex bot bot Id aliases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetIntegrationsSpeechLexBotBotIDAliasesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) WithTimeout(timeout time.Duration) *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) WithContext(ctx context.Context) *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) WithHTTPClient(client *http.Client) *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBotID adds the botID to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) WithBotID(botID string) *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	o.SetBotID(botID)
	return o
}

// SetBotID adds the botId to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) SetBotID(botID string) {
	o.BotID = botID
}

// WithName adds the name to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) WithName(name *string) *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) WithPageNumber(pageNumber *int32) *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) WithPageSize(pageSize *int32) *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithStatus adds the status to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) WithStatus(status *string) *GetIntegrationsSpeechLexBotBotIDAliasesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get integrations speech lex bot bot Id aliases params
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationsSpeechLexBotBotIDAliasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param botId
	if err := r.SetPathParam("botId", o.BotID); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
