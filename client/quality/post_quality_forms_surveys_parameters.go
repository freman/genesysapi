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

// NewPostQualityFormsSurveysParams creates a new PostQualityFormsSurveysParams object
// with the default values initialized.
func NewPostQualityFormsSurveysParams() *PostQualityFormsSurveysParams {
	var ()
	return &PostQualityFormsSurveysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostQualityFormsSurveysParamsWithTimeout creates a new PostQualityFormsSurveysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostQualityFormsSurveysParamsWithTimeout(timeout time.Duration) *PostQualityFormsSurveysParams {
	var ()
	return &PostQualityFormsSurveysParams{

		timeout: timeout,
	}
}

// NewPostQualityFormsSurveysParamsWithContext creates a new PostQualityFormsSurveysParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostQualityFormsSurveysParamsWithContext(ctx context.Context) *PostQualityFormsSurveysParams {
	var ()
	return &PostQualityFormsSurveysParams{

		Context: ctx,
	}
}

// NewPostQualityFormsSurveysParamsWithHTTPClient creates a new PostQualityFormsSurveysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostQualityFormsSurveysParamsWithHTTPClient(client *http.Client) *PostQualityFormsSurveysParams {
	var ()
	return &PostQualityFormsSurveysParams{
		HTTPClient: client,
	}
}

/*PostQualityFormsSurveysParams contains all the parameters to send to the API endpoint
for the post quality forms surveys operation typically these are written to a http.Request
*/
type PostQualityFormsSurveysParams struct {

	/*Body
	  Survey form

	*/
	Body *models.SurveyForm

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post quality forms surveys params
func (o *PostQualityFormsSurveysParams) WithTimeout(timeout time.Duration) *PostQualityFormsSurveysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post quality forms surveys params
func (o *PostQualityFormsSurveysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post quality forms surveys params
func (o *PostQualityFormsSurveysParams) WithContext(ctx context.Context) *PostQualityFormsSurveysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post quality forms surveys params
func (o *PostQualityFormsSurveysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post quality forms surveys params
func (o *PostQualityFormsSurveysParams) WithHTTPClient(client *http.Client) *PostQualityFormsSurveysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post quality forms surveys params
func (o *PostQualityFormsSurveysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post quality forms surveys params
func (o *PostQualityFormsSurveysParams) WithBody(body *models.SurveyForm) *PostQualityFormsSurveysParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post quality forms surveys params
func (o *PostQualityFormsSurveysParams) SetBody(body *models.SurveyForm) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostQualityFormsSurveysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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