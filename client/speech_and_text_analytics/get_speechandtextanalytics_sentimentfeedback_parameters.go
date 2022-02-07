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
)

// NewGetSpeechandtextanalyticsSentimentfeedbackParams creates a new GetSpeechandtextanalyticsSentimentfeedbackParams object
// with the default values initialized.
func NewGetSpeechandtextanalyticsSentimentfeedbackParams() *GetSpeechandtextanalyticsSentimentfeedbackParams {
	var ()
	return &GetSpeechandtextanalyticsSentimentfeedbackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpeechandtextanalyticsSentimentfeedbackParamsWithTimeout creates a new GetSpeechandtextanalyticsSentimentfeedbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSpeechandtextanalyticsSentimentfeedbackParamsWithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsSentimentfeedbackParams {
	var ()
	return &GetSpeechandtextanalyticsSentimentfeedbackParams{

		timeout: timeout,
	}
}

// NewGetSpeechandtextanalyticsSentimentfeedbackParamsWithContext creates a new GetSpeechandtextanalyticsSentimentfeedbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSpeechandtextanalyticsSentimentfeedbackParamsWithContext(ctx context.Context) *GetSpeechandtextanalyticsSentimentfeedbackParams {
	var ()
	return &GetSpeechandtextanalyticsSentimentfeedbackParams{

		Context: ctx,
	}
}

// NewGetSpeechandtextanalyticsSentimentfeedbackParamsWithHTTPClient creates a new GetSpeechandtextanalyticsSentimentfeedbackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSpeechandtextanalyticsSentimentfeedbackParamsWithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsSentimentfeedbackParams {
	var ()
	return &GetSpeechandtextanalyticsSentimentfeedbackParams{
		HTTPClient: client,
	}
}

/*GetSpeechandtextanalyticsSentimentfeedbackParams contains all the parameters to send to the API endpoint
for the get speechandtextanalytics sentimentfeedback operation typically these are written to a http.Request
*/
type GetSpeechandtextanalyticsSentimentfeedbackParams struct {

	/*Dialect
	  The key for filter the listing by dialect, dialect format is {language}-{country} where language follows ISO 639-1 standard and country follows ISO 3166-1 alpha 2 standard

	*/
	Dialect *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get speechandtextanalytics sentimentfeedback params
func (o *GetSpeechandtextanalyticsSentimentfeedbackParams) WithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsSentimentfeedbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get speechandtextanalytics sentimentfeedback params
func (o *GetSpeechandtextanalyticsSentimentfeedbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get speechandtextanalytics sentimentfeedback params
func (o *GetSpeechandtextanalyticsSentimentfeedbackParams) WithContext(ctx context.Context) *GetSpeechandtextanalyticsSentimentfeedbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get speechandtextanalytics sentimentfeedback params
func (o *GetSpeechandtextanalyticsSentimentfeedbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get speechandtextanalytics sentimentfeedback params
func (o *GetSpeechandtextanalyticsSentimentfeedbackParams) WithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsSentimentfeedbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get speechandtextanalytics sentimentfeedback params
func (o *GetSpeechandtextanalyticsSentimentfeedbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDialect adds the dialect to the get speechandtextanalytics sentimentfeedback params
func (o *GetSpeechandtextanalyticsSentimentfeedbackParams) WithDialect(dialect *string) *GetSpeechandtextanalyticsSentimentfeedbackParams {
	o.SetDialect(dialect)
	return o
}

// SetDialect adds the dialect to the get speechandtextanalytics sentimentfeedback params
func (o *GetSpeechandtextanalyticsSentimentfeedbackParams) SetDialect(dialect *string) {
	o.Dialect = dialect
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpeechandtextanalyticsSentimentfeedbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Dialect != nil {

		// query param dialect
		var qrDialect string
		if o.Dialect != nil {
			qrDialect = *o.Dialect
		}
		qDialect := qrDialect
		if qDialect != "" {
			if err := r.SetQueryParam("dialect", qDialect); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}