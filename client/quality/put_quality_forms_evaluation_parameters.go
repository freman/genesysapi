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

// NewPutQualityFormsEvaluationParams creates a new PutQualityFormsEvaluationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutQualityFormsEvaluationParams() *PutQualityFormsEvaluationParams {
	return &PutQualityFormsEvaluationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutQualityFormsEvaluationParamsWithTimeout creates a new PutQualityFormsEvaluationParams object
// with the ability to set a timeout on a request.
func NewPutQualityFormsEvaluationParamsWithTimeout(timeout time.Duration) *PutQualityFormsEvaluationParams {
	return &PutQualityFormsEvaluationParams{
		timeout: timeout,
	}
}

// NewPutQualityFormsEvaluationParamsWithContext creates a new PutQualityFormsEvaluationParams object
// with the ability to set a context for a request.
func NewPutQualityFormsEvaluationParamsWithContext(ctx context.Context) *PutQualityFormsEvaluationParams {
	return &PutQualityFormsEvaluationParams{
		Context: ctx,
	}
}

// NewPutQualityFormsEvaluationParamsWithHTTPClient creates a new PutQualityFormsEvaluationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutQualityFormsEvaluationParamsWithHTTPClient(client *http.Client) *PutQualityFormsEvaluationParams {
	return &PutQualityFormsEvaluationParams{
		HTTPClient: client,
	}
}

/*
PutQualityFormsEvaluationParams contains all the parameters to send to the API endpoint

	for the put quality forms evaluation operation.

	Typically these are written to a http.Request.
*/
type PutQualityFormsEvaluationParams struct {

	/* Body.

	   Evaluation form
	*/
	Body *models.EvaluationForm

	/* FormID.

	   Form ID
	*/
	FormID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put quality forms evaluation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutQualityFormsEvaluationParams) WithDefaults() *PutQualityFormsEvaluationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put quality forms evaluation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutQualityFormsEvaluationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put quality forms evaluation params
func (o *PutQualityFormsEvaluationParams) WithTimeout(timeout time.Duration) *PutQualityFormsEvaluationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put quality forms evaluation params
func (o *PutQualityFormsEvaluationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put quality forms evaluation params
func (o *PutQualityFormsEvaluationParams) WithContext(ctx context.Context) *PutQualityFormsEvaluationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put quality forms evaluation params
func (o *PutQualityFormsEvaluationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put quality forms evaluation params
func (o *PutQualityFormsEvaluationParams) WithHTTPClient(client *http.Client) *PutQualityFormsEvaluationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put quality forms evaluation params
func (o *PutQualityFormsEvaluationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put quality forms evaluation params
func (o *PutQualityFormsEvaluationParams) WithBody(body *models.EvaluationForm) *PutQualityFormsEvaluationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put quality forms evaluation params
func (o *PutQualityFormsEvaluationParams) SetBody(body *models.EvaluationForm) {
	o.Body = body
}

// WithFormID adds the formID to the put quality forms evaluation params
func (o *PutQualityFormsEvaluationParams) WithFormID(formID string) *PutQualityFormsEvaluationParams {
	o.SetFormID(formID)
	return o
}

// SetFormID adds the formId to the put quality forms evaluation params
func (o *PutQualityFormsEvaluationParams) SetFormID(formID string) {
	o.FormID = formID
}

// WriteToRequest writes these params to a swagger request
func (o *PutQualityFormsEvaluationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param formId
	if err := r.SetPathParam("formId", o.FormID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
