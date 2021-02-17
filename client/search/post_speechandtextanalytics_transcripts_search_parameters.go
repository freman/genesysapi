// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewPostSpeechandtextanalyticsTranscriptsSearchParams creates a new PostSpeechandtextanalyticsTranscriptsSearchParams object
// with the default values initialized.
func NewPostSpeechandtextanalyticsTranscriptsSearchParams() *PostSpeechandtextanalyticsTranscriptsSearchParams {
	var ()
	return &PostSpeechandtextanalyticsTranscriptsSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSpeechandtextanalyticsTranscriptsSearchParamsWithTimeout creates a new PostSpeechandtextanalyticsTranscriptsSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSpeechandtextanalyticsTranscriptsSearchParamsWithTimeout(timeout time.Duration) *PostSpeechandtextanalyticsTranscriptsSearchParams {
	var ()
	return &PostSpeechandtextanalyticsTranscriptsSearchParams{

		timeout: timeout,
	}
}

// NewPostSpeechandtextanalyticsTranscriptsSearchParamsWithContext creates a new PostSpeechandtextanalyticsTranscriptsSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSpeechandtextanalyticsTranscriptsSearchParamsWithContext(ctx context.Context) *PostSpeechandtextanalyticsTranscriptsSearchParams {
	var ()
	return &PostSpeechandtextanalyticsTranscriptsSearchParams{

		Context: ctx,
	}
}

// NewPostSpeechandtextanalyticsTranscriptsSearchParamsWithHTTPClient creates a new PostSpeechandtextanalyticsTranscriptsSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSpeechandtextanalyticsTranscriptsSearchParamsWithHTTPClient(client *http.Client) *PostSpeechandtextanalyticsTranscriptsSearchParams {
	var ()
	return &PostSpeechandtextanalyticsTranscriptsSearchParams{
		HTTPClient: client,
	}
}

/*PostSpeechandtextanalyticsTranscriptsSearchParams contains all the parameters to send to the API endpoint
for the post speechandtextanalytics transcripts search operation typically these are written to a http.Request
*/
type PostSpeechandtextanalyticsTranscriptsSearchParams struct {

	/*Body
	  Search request options

	*/
	Body *models.TranscriptSearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post speechandtextanalytics transcripts search params
func (o *PostSpeechandtextanalyticsTranscriptsSearchParams) WithTimeout(timeout time.Duration) *PostSpeechandtextanalyticsTranscriptsSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post speechandtextanalytics transcripts search params
func (o *PostSpeechandtextanalyticsTranscriptsSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post speechandtextanalytics transcripts search params
func (o *PostSpeechandtextanalyticsTranscriptsSearchParams) WithContext(ctx context.Context) *PostSpeechandtextanalyticsTranscriptsSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post speechandtextanalytics transcripts search params
func (o *PostSpeechandtextanalyticsTranscriptsSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post speechandtextanalytics transcripts search params
func (o *PostSpeechandtextanalyticsTranscriptsSearchParams) WithHTTPClient(client *http.Client) *PostSpeechandtextanalyticsTranscriptsSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post speechandtextanalytics transcripts search params
func (o *PostSpeechandtextanalyticsTranscriptsSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post speechandtextanalytics transcripts search params
func (o *PostSpeechandtextanalyticsTranscriptsSearchParams) WithBody(body *models.TranscriptSearchRequest) *PostSpeechandtextanalyticsTranscriptsSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post speechandtextanalytics transcripts search params
func (o *PostSpeechandtextanalyticsTranscriptsSearchParams) SetBody(body *models.TranscriptSearchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSpeechandtextanalyticsTranscriptsSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
