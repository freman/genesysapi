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

// NewGetQualityFormsEvaluationParams creates a new GetQualityFormsEvaluationParams object
// with the default values initialized.
func NewGetQualityFormsEvaluationParams() *GetQualityFormsEvaluationParams {
	var ()
	return &GetQualityFormsEvaluationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetQualityFormsEvaluationParamsWithTimeout creates a new GetQualityFormsEvaluationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetQualityFormsEvaluationParamsWithTimeout(timeout time.Duration) *GetQualityFormsEvaluationParams {
	var ()
	return &GetQualityFormsEvaluationParams{

		timeout: timeout,
	}
}

// NewGetQualityFormsEvaluationParamsWithContext creates a new GetQualityFormsEvaluationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetQualityFormsEvaluationParamsWithContext(ctx context.Context) *GetQualityFormsEvaluationParams {
	var ()
	return &GetQualityFormsEvaluationParams{

		Context: ctx,
	}
}

// NewGetQualityFormsEvaluationParamsWithHTTPClient creates a new GetQualityFormsEvaluationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetQualityFormsEvaluationParamsWithHTTPClient(client *http.Client) *GetQualityFormsEvaluationParams {
	var ()
	return &GetQualityFormsEvaluationParams{
		HTTPClient: client,
	}
}

/*GetQualityFormsEvaluationParams contains all the parameters to send to the API endpoint
for the get quality forms evaluation operation typically these are written to a http.Request
*/
type GetQualityFormsEvaluationParams struct {

	/*FormID
	  Form ID

	*/
	FormID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get quality forms evaluation params
func (o *GetQualityFormsEvaluationParams) WithTimeout(timeout time.Duration) *GetQualityFormsEvaluationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get quality forms evaluation params
func (o *GetQualityFormsEvaluationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get quality forms evaluation params
func (o *GetQualityFormsEvaluationParams) WithContext(ctx context.Context) *GetQualityFormsEvaluationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get quality forms evaluation params
func (o *GetQualityFormsEvaluationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get quality forms evaluation params
func (o *GetQualityFormsEvaluationParams) WithHTTPClient(client *http.Client) *GetQualityFormsEvaluationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get quality forms evaluation params
func (o *GetQualityFormsEvaluationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormID adds the formID to the get quality forms evaluation params
func (o *GetQualityFormsEvaluationParams) WithFormID(formID string) *GetQualityFormsEvaluationParams {
	o.SetFormID(formID)
	return o
}

// SetFormID adds the formId to the get quality forms evaluation params
func (o *GetQualityFormsEvaluationParams) SetFormID(formID string) {
	o.FormID = formID
}

// WriteToRequest writes these params to a swagger request
func (o *GetQualityFormsEvaluationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
