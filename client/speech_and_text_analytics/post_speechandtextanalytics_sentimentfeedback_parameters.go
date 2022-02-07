// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

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

// NewPostSpeechandtextanalyticsSentimentfeedbackParams creates a new PostSpeechandtextanalyticsSentimentfeedbackParams object
// with the default values initialized.
func NewPostSpeechandtextanalyticsSentimentfeedbackParams() *PostSpeechandtextanalyticsSentimentfeedbackParams {
	var ()
	return &PostSpeechandtextanalyticsSentimentfeedbackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSpeechandtextanalyticsSentimentfeedbackParamsWithTimeout creates a new PostSpeechandtextanalyticsSentimentfeedbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSpeechandtextanalyticsSentimentfeedbackParamsWithTimeout(timeout time.Duration) *PostSpeechandtextanalyticsSentimentfeedbackParams {
	var ()
	return &PostSpeechandtextanalyticsSentimentfeedbackParams{

		timeout: timeout,
	}
}

// NewPostSpeechandtextanalyticsSentimentfeedbackParamsWithContext creates a new PostSpeechandtextanalyticsSentimentfeedbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSpeechandtextanalyticsSentimentfeedbackParamsWithContext(ctx context.Context) *PostSpeechandtextanalyticsSentimentfeedbackParams {
	var ()
	return &PostSpeechandtextanalyticsSentimentfeedbackParams{

		Context: ctx,
	}
}

// NewPostSpeechandtextanalyticsSentimentfeedbackParamsWithHTTPClient creates a new PostSpeechandtextanalyticsSentimentfeedbackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSpeechandtextanalyticsSentimentfeedbackParamsWithHTTPClient(client *http.Client) *PostSpeechandtextanalyticsSentimentfeedbackParams {
	var ()
	return &PostSpeechandtextanalyticsSentimentfeedbackParams{
		HTTPClient: client,
	}
}

/*PostSpeechandtextanalyticsSentimentfeedbackParams contains all the parameters to send to the API endpoint
for the post speechandtextanalytics sentimentfeedback operation typically these are written to a http.Request
*/
type PostSpeechandtextanalyticsSentimentfeedbackParams struct {

	/*Body
	  The SentimentFeedback to create

	*/
	Body *models.SentimentFeedback

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post speechandtextanalytics sentimentfeedback params
func (o *PostSpeechandtextanalyticsSentimentfeedbackParams) WithTimeout(timeout time.Duration) *PostSpeechandtextanalyticsSentimentfeedbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post speechandtextanalytics sentimentfeedback params
func (o *PostSpeechandtextanalyticsSentimentfeedbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post speechandtextanalytics sentimentfeedback params
func (o *PostSpeechandtextanalyticsSentimentfeedbackParams) WithContext(ctx context.Context) *PostSpeechandtextanalyticsSentimentfeedbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post speechandtextanalytics sentimentfeedback params
func (o *PostSpeechandtextanalyticsSentimentfeedbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post speechandtextanalytics sentimentfeedback params
func (o *PostSpeechandtextanalyticsSentimentfeedbackParams) WithHTTPClient(client *http.Client) *PostSpeechandtextanalyticsSentimentfeedbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post speechandtextanalytics sentimentfeedback params
func (o *PostSpeechandtextanalyticsSentimentfeedbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post speechandtextanalytics sentimentfeedback params
func (o *PostSpeechandtextanalyticsSentimentfeedbackParams) WithBody(body *models.SentimentFeedback) *PostSpeechandtextanalyticsSentimentfeedbackParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post speechandtextanalytics sentimentfeedback params
func (o *PostSpeechandtextanalyticsSentimentfeedbackParams) SetBody(body *models.SentimentFeedback) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSpeechandtextanalyticsSentimentfeedbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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