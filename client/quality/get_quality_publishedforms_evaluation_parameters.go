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

// NewGetQualityPublishedformsEvaluationParams creates a new GetQualityPublishedformsEvaluationParams object
// with the default values initialized.
func NewGetQualityPublishedformsEvaluationParams() *GetQualityPublishedformsEvaluationParams {
	var ()
	return &GetQualityPublishedformsEvaluationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetQualityPublishedformsEvaluationParamsWithTimeout creates a new GetQualityPublishedformsEvaluationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetQualityPublishedformsEvaluationParamsWithTimeout(timeout time.Duration) *GetQualityPublishedformsEvaluationParams {
	var ()
	return &GetQualityPublishedformsEvaluationParams{

		timeout: timeout,
	}
}

// NewGetQualityPublishedformsEvaluationParamsWithContext creates a new GetQualityPublishedformsEvaluationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetQualityPublishedformsEvaluationParamsWithContext(ctx context.Context) *GetQualityPublishedformsEvaluationParams {
	var ()
	return &GetQualityPublishedformsEvaluationParams{

		Context: ctx,
	}
}

// NewGetQualityPublishedformsEvaluationParamsWithHTTPClient creates a new GetQualityPublishedformsEvaluationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetQualityPublishedformsEvaluationParamsWithHTTPClient(client *http.Client) *GetQualityPublishedformsEvaluationParams {
	var ()
	return &GetQualityPublishedformsEvaluationParams{
		HTTPClient: client,
	}
}

/*GetQualityPublishedformsEvaluationParams contains all the parameters to send to the API endpoint
for the get quality publishedforms evaluation operation typically these are written to a http.Request
*/
type GetQualityPublishedformsEvaluationParams struct {

	/*FormID
	  Form ID

	*/
	FormID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get quality publishedforms evaluation params
func (o *GetQualityPublishedformsEvaluationParams) WithTimeout(timeout time.Duration) *GetQualityPublishedformsEvaluationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get quality publishedforms evaluation params
func (o *GetQualityPublishedformsEvaluationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get quality publishedforms evaluation params
func (o *GetQualityPublishedformsEvaluationParams) WithContext(ctx context.Context) *GetQualityPublishedformsEvaluationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get quality publishedforms evaluation params
func (o *GetQualityPublishedformsEvaluationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get quality publishedforms evaluation params
func (o *GetQualityPublishedformsEvaluationParams) WithHTTPClient(client *http.Client) *GetQualityPublishedformsEvaluationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get quality publishedforms evaluation params
func (o *GetQualityPublishedformsEvaluationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormID adds the formID to the get quality publishedforms evaluation params
func (o *GetQualityPublishedformsEvaluationParams) WithFormID(formID string) *GetQualityPublishedformsEvaluationParams {
	o.SetFormID(formID)
	return o
}

// SetFormID adds the formId to the get quality publishedforms evaluation params
func (o *GetQualityPublishedformsEvaluationParams) SetFormID(formID string) {
	o.FormID = formID
}

// WriteToRequest writes these params to a swagger request
func (o *GetQualityPublishedformsEvaluationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
