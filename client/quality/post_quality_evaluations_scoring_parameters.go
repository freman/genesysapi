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

// NewPostQualityEvaluationsScoringParams creates a new PostQualityEvaluationsScoringParams object
// with the default values initialized.
func NewPostQualityEvaluationsScoringParams() *PostQualityEvaluationsScoringParams {
	var ()
	return &PostQualityEvaluationsScoringParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostQualityEvaluationsScoringParamsWithTimeout creates a new PostQualityEvaluationsScoringParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostQualityEvaluationsScoringParamsWithTimeout(timeout time.Duration) *PostQualityEvaluationsScoringParams {
	var ()
	return &PostQualityEvaluationsScoringParams{

		timeout: timeout,
	}
}

// NewPostQualityEvaluationsScoringParamsWithContext creates a new PostQualityEvaluationsScoringParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostQualityEvaluationsScoringParamsWithContext(ctx context.Context) *PostQualityEvaluationsScoringParams {
	var ()
	return &PostQualityEvaluationsScoringParams{

		Context: ctx,
	}
}

// NewPostQualityEvaluationsScoringParamsWithHTTPClient creates a new PostQualityEvaluationsScoringParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostQualityEvaluationsScoringParamsWithHTTPClient(client *http.Client) *PostQualityEvaluationsScoringParams {
	var ()
	return &PostQualityEvaluationsScoringParams{
		HTTPClient: client,
	}
}

/*PostQualityEvaluationsScoringParams contains all the parameters to send to the API endpoint
for the post quality evaluations scoring operation typically these are written to a http.Request
*/
type PostQualityEvaluationsScoringParams struct {

	/*Body
	  evaluationAndScoringSet

	*/
	Body *models.EvaluationFormAndScoringSet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post quality evaluations scoring params
func (o *PostQualityEvaluationsScoringParams) WithTimeout(timeout time.Duration) *PostQualityEvaluationsScoringParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post quality evaluations scoring params
func (o *PostQualityEvaluationsScoringParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post quality evaluations scoring params
func (o *PostQualityEvaluationsScoringParams) WithContext(ctx context.Context) *PostQualityEvaluationsScoringParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post quality evaluations scoring params
func (o *PostQualityEvaluationsScoringParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post quality evaluations scoring params
func (o *PostQualityEvaluationsScoringParams) WithHTTPClient(client *http.Client) *PostQualityEvaluationsScoringParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post quality evaluations scoring params
func (o *PostQualityEvaluationsScoringParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post quality evaluations scoring params
func (o *PostQualityEvaluationsScoringParams) WithBody(body *models.EvaluationFormAndScoringSet) *PostQualityEvaluationsScoringParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post quality evaluations scoring params
func (o *PostQualityEvaluationsScoringParams) SetBody(body *models.EvaluationFormAndScoringSet) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostQualityEvaluationsScoringParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}