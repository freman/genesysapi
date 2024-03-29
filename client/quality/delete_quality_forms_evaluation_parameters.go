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

// NewDeleteQualityFormsEvaluationParams creates a new DeleteQualityFormsEvaluationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteQualityFormsEvaluationParams() *DeleteQualityFormsEvaluationParams {
	return &DeleteQualityFormsEvaluationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteQualityFormsEvaluationParamsWithTimeout creates a new DeleteQualityFormsEvaluationParams object
// with the ability to set a timeout on a request.
func NewDeleteQualityFormsEvaluationParamsWithTimeout(timeout time.Duration) *DeleteQualityFormsEvaluationParams {
	return &DeleteQualityFormsEvaluationParams{
		timeout: timeout,
	}
}

// NewDeleteQualityFormsEvaluationParamsWithContext creates a new DeleteQualityFormsEvaluationParams object
// with the ability to set a context for a request.
func NewDeleteQualityFormsEvaluationParamsWithContext(ctx context.Context) *DeleteQualityFormsEvaluationParams {
	return &DeleteQualityFormsEvaluationParams{
		Context: ctx,
	}
}

// NewDeleteQualityFormsEvaluationParamsWithHTTPClient creates a new DeleteQualityFormsEvaluationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteQualityFormsEvaluationParamsWithHTTPClient(client *http.Client) *DeleteQualityFormsEvaluationParams {
	return &DeleteQualityFormsEvaluationParams{
		HTTPClient: client,
	}
}

/*
DeleteQualityFormsEvaluationParams contains all the parameters to send to the API endpoint

	for the delete quality forms evaluation operation.

	Typically these are written to a http.Request.
*/
type DeleteQualityFormsEvaluationParams struct {

	/* FormID.

	   Form ID
	*/
	FormID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete quality forms evaluation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteQualityFormsEvaluationParams) WithDefaults() *DeleteQualityFormsEvaluationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete quality forms evaluation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteQualityFormsEvaluationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete quality forms evaluation params
func (o *DeleteQualityFormsEvaluationParams) WithTimeout(timeout time.Duration) *DeleteQualityFormsEvaluationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete quality forms evaluation params
func (o *DeleteQualityFormsEvaluationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete quality forms evaluation params
func (o *DeleteQualityFormsEvaluationParams) WithContext(ctx context.Context) *DeleteQualityFormsEvaluationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete quality forms evaluation params
func (o *DeleteQualityFormsEvaluationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete quality forms evaluation params
func (o *DeleteQualityFormsEvaluationParams) WithHTTPClient(client *http.Client) *DeleteQualityFormsEvaluationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete quality forms evaluation params
func (o *DeleteQualityFormsEvaluationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormID adds the formID to the delete quality forms evaluation params
func (o *DeleteQualityFormsEvaluationParams) WithFormID(formID string) *DeleteQualityFormsEvaluationParams {
	o.SetFormID(formID)
	return o
}

// SetFormID adds the formId to the delete quality forms evaluation params
func (o *DeleteQualityFormsEvaluationParams) SetFormID(formID string) {
	o.FormID = formID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteQualityFormsEvaluationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param formId
	if err := r.SetPathParam("formId", o.FormID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
