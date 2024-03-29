// Code generated by go-swagger; DO NOT EDIT.

package quality

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

// NewGetQualityFormsEvaluationsBulkContextsParams creates a new GetQualityFormsEvaluationsBulkContextsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetQualityFormsEvaluationsBulkContextsParams() *GetQualityFormsEvaluationsBulkContextsParams {
	return &GetQualityFormsEvaluationsBulkContextsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetQualityFormsEvaluationsBulkContextsParamsWithTimeout creates a new GetQualityFormsEvaluationsBulkContextsParams object
// with the ability to set a timeout on a request.
func NewGetQualityFormsEvaluationsBulkContextsParamsWithTimeout(timeout time.Duration) *GetQualityFormsEvaluationsBulkContextsParams {
	return &GetQualityFormsEvaluationsBulkContextsParams{
		timeout: timeout,
	}
}

// NewGetQualityFormsEvaluationsBulkContextsParamsWithContext creates a new GetQualityFormsEvaluationsBulkContextsParams object
// with the ability to set a context for a request.
func NewGetQualityFormsEvaluationsBulkContextsParamsWithContext(ctx context.Context) *GetQualityFormsEvaluationsBulkContextsParams {
	return &GetQualityFormsEvaluationsBulkContextsParams{
		Context: ctx,
	}
}

// NewGetQualityFormsEvaluationsBulkContextsParamsWithHTTPClient creates a new GetQualityFormsEvaluationsBulkContextsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetQualityFormsEvaluationsBulkContextsParamsWithHTTPClient(client *http.Client) *GetQualityFormsEvaluationsBulkContextsParams {
	return &GetQualityFormsEvaluationsBulkContextsParams{
		HTTPClient: client,
	}
}

/*
GetQualityFormsEvaluationsBulkContextsParams contains all the parameters to send to the API endpoint

	for the get quality forms evaluations bulk contexts operation.

	Typically these are written to a http.Request.
*/
type GetQualityFormsEvaluationsBulkContextsParams struct {

	/* ContextID.

	   A comma-delimited list of valid evaluation form context ids
	*/
	ContextID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get quality forms evaluations bulk contexts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQualityFormsEvaluationsBulkContextsParams) WithDefaults() *GetQualityFormsEvaluationsBulkContextsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get quality forms evaluations bulk contexts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQualityFormsEvaluationsBulkContextsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get quality forms evaluations bulk contexts params
func (o *GetQualityFormsEvaluationsBulkContextsParams) WithTimeout(timeout time.Duration) *GetQualityFormsEvaluationsBulkContextsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get quality forms evaluations bulk contexts params
func (o *GetQualityFormsEvaluationsBulkContextsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get quality forms evaluations bulk contexts params
func (o *GetQualityFormsEvaluationsBulkContextsParams) WithContext(ctx context.Context) *GetQualityFormsEvaluationsBulkContextsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get quality forms evaluations bulk contexts params
func (o *GetQualityFormsEvaluationsBulkContextsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get quality forms evaluations bulk contexts params
func (o *GetQualityFormsEvaluationsBulkContextsParams) WithHTTPClient(client *http.Client) *GetQualityFormsEvaluationsBulkContextsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get quality forms evaluations bulk contexts params
func (o *GetQualityFormsEvaluationsBulkContextsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContextID adds the contextID to the get quality forms evaluations bulk contexts params
func (o *GetQualityFormsEvaluationsBulkContextsParams) WithContextID(contextID []string) *GetQualityFormsEvaluationsBulkContextsParams {
	o.SetContextID(contextID)
	return o
}

// SetContextID adds the contextId to the get quality forms evaluations bulk contexts params
func (o *GetQualityFormsEvaluationsBulkContextsParams) SetContextID(contextID []string) {
	o.ContextID = contextID
}

// WriteToRequest writes these params to a swagger request
func (o *GetQualityFormsEvaluationsBulkContextsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContextID != nil {

		// binding items for contextId
		joinedContextID := o.bindParamContextID(reg)

		// query array param contextId
		if err := r.SetQueryParam("contextId", joinedContextID...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetQualityFormsEvaluationsBulkContexts binds the parameter contextId
func (o *GetQualityFormsEvaluationsBulkContextsParams) bindParamContextID(formats strfmt.Registry) []string {
	contextIDIR := o.ContextID

	var contextIDIC []string
	for _, contextIDIIR := range contextIDIR { // explode []string

		contextIDIIV := contextIDIIR // string as string
		contextIDIC = append(contextIDIC, contextIDIIV)
	}

	// items.CollectionFormat: "multi"
	contextIDIS := swag.JoinByFormat(contextIDIC, "multi")

	return contextIDIS
}
