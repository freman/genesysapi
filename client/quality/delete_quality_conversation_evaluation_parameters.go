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
)

// NewDeleteQualityConversationEvaluationParams creates a new DeleteQualityConversationEvaluationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteQualityConversationEvaluationParams() *DeleteQualityConversationEvaluationParams {
	return &DeleteQualityConversationEvaluationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteQualityConversationEvaluationParamsWithTimeout creates a new DeleteQualityConversationEvaluationParams object
// with the ability to set a timeout on a request.
func NewDeleteQualityConversationEvaluationParamsWithTimeout(timeout time.Duration) *DeleteQualityConversationEvaluationParams {
	return &DeleteQualityConversationEvaluationParams{
		timeout: timeout,
	}
}

// NewDeleteQualityConversationEvaluationParamsWithContext creates a new DeleteQualityConversationEvaluationParams object
// with the ability to set a context for a request.
func NewDeleteQualityConversationEvaluationParamsWithContext(ctx context.Context) *DeleteQualityConversationEvaluationParams {
	return &DeleteQualityConversationEvaluationParams{
		Context: ctx,
	}
}

// NewDeleteQualityConversationEvaluationParamsWithHTTPClient creates a new DeleteQualityConversationEvaluationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteQualityConversationEvaluationParamsWithHTTPClient(client *http.Client) *DeleteQualityConversationEvaluationParams {
	return &DeleteQualityConversationEvaluationParams{
		HTTPClient: client,
	}
}

/*
DeleteQualityConversationEvaluationParams contains all the parameters to send to the API endpoint

	for the delete quality conversation evaluation operation.

	Typically these are written to a http.Request.
*/
type DeleteQualityConversationEvaluationParams struct {

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

// WithDefaults hydrates default values in the delete quality conversation evaluation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteQualityConversationEvaluationParams) WithDefaults() *DeleteQualityConversationEvaluationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete quality conversation evaluation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteQualityConversationEvaluationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete quality conversation evaluation params
func (o *DeleteQualityConversationEvaluationParams) WithTimeout(timeout time.Duration) *DeleteQualityConversationEvaluationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete quality conversation evaluation params
func (o *DeleteQualityConversationEvaluationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete quality conversation evaluation params
func (o *DeleteQualityConversationEvaluationParams) WithContext(ctx context.Context) *DeleteQualityConversationEvaluationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete quality conversation evaluation params
func (o *DeleteQualityConversationEvaluationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete quality conversation evaluation params
func (o *DeleteQualityConversationEvaluationParams) WithHTTPClient(client *http.Client) *DeleteQualityConversationEvaluationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete quality conversation evaluation params
func (o *DeleteQualityConversationEvaluationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the delete quality conversation evaluation params
func (o *DeleteQualityConversationEvaluationParams) WithConversationID(conversationID string) *DeleteQualityConversationEvaluationParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the delete quality conversation evaluation params
func (o *DeleteQualityConversationEvaluationParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithEvaluationID adds the evaluationID to the delete quality conversation evaluation params
func (o *DeleteQualityConversationEvaluationParams) WithEvaluationID(evaluationID string) *DeleteQualityConversationEvaluationParams {
	o.SetEvaluationID(evaluationID)
	return o
}

// SetEvaluationID adds the evaluationId to the delete quality conversation evaluation params
func (o *DeleteQualityConversationEvaluationParams) SetEvaluationID(evaluationID string) {
	o.EvaluationID = evaluationID
}

// WithExpand adds the expand to the delete quality conversation evaluation params
func (o *DeleteQualityConversationEvaluationParams) WithExpand(expand *string) *DeleteQualityConversationEvaluationParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the delete quality conversation evaluation params
func (o *DeleteQualityConversationEvaluationParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteQualityConversationEvaluationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
