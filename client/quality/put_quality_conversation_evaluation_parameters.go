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

	"github.com/freman/genesysapi/models"
)

// NewPutQualityConversationEvaluationParams creates a new PutQualityConversationEvaluationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutQualityConversationEvaluationParams() *PutQualityConversationEvaluationParams {
	return &PutQualityConversationEvaluationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutQualityConversationEvaluationParamsWithTimeout creates a new PutQualityConversationEvaluationParams object
// with the ability to set a timeout on a request.
func NewPutQualityConversationEvaluationParamsWithTimeout(timeout time.Duration) *PutQualityConversationEvaluationParams {
	return &PutQualityConversationEvaluationParams{
		timeout: timeout,
	}
}

// NewPutQualityConversationEvaluationParamsWithContext creates a new PutQualityConversationEvaluationParams object
// with the ability to set a context for a request.
func NewPutQualityConversationEvaluationParamsWithContext(ctx context.Context) *PutQualityConversationEvaluationParams {
	return &PutQualityConversationEvaluationParams{
		Context: ctx,
	}
}

// NewPutQualityConversationEvaluationParamsWithHTTPClient creates a new PutQualityConversationEvaluationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutQualityConversationEvaluationParamsWithHTTPClient(client *http.Client) *PutQualityConversationEvaluationParams {
	return &PutQualityConversationEvaluationParams{
		HTTPClient: client,
	}
}

/*
PutQualityConversationEvaluationParams contains all the parameters to send to the API endpoint

	for the put quality conversation evaluation operation.

	Typically these are written to a http.Request.
*/
type PutQualityConversationEvaluationParams struct {

	/* Body.

	   evaluation
	*/
	Body *models.Evaluation

	/* ConversationID.

	   conversationId
	*/
	ConversationID string

	/* EvaluationID.

	   evaluationId
	*/
	EvaluationID string

	/* Expand.

	   evaluatorId, evaluationForm
	*/
	Expand *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put quality conversation evaluation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutQualityConversationEvaluationParams) WithDefaults() *PutQualityConversationEvaluationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put quality conversation evaluation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutQualityConversationEvaluationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) WithTimeout(timeout time.Duration) *PutQualityConversationEvaluationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) WithContext(ctx context.Context) *PutQualityConversationEvaluationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) WithHTTPClient(client *http.Client) *PutQualityConversationEvaluationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) WithBody(body *models.Evaluation) *PutQualityConversationEvaluationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) SetBody(body *models.Evaluation) {
	o.Body = body
}

// WithConversationID adds the conversationID to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) WithConversationID(conversationID string) *PutQualityConversationEvaluationParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithEvaluationID adds the evaluationID to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) WithEvaluationID(evaluationID string) *PutQualityConversationEvaluationParams {
	o.SetEvaluationID(evaluationID)
	return o
}

// SetEvaluationID adds the evaluationId to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) SetEvaluationID(evaluationID string) {
	o.EvaluationID = evaluationID
}

// WithExpand adds the expand to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) WithExpand(expand *string) *PutQualityConversationEvaluationParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the put quality conversation evaluation params
func (o *PutQualityConversationEvaluationParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *PutQualityConversationEvaluationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param conversationId
	if err := r.SetPathParam("conversationId", o.ConversationID); err != nil {
		return err
	}

	// path param evaluationId
	if err := r.SetPathParam("evaluationId", o.EvaluationID); err != nil {
		return err
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
