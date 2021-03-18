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

// NewPostSpeechandtextanalyticsTopicsParams creates a new PostSpeechandtextanalyticsTopicsParams object
// with the default values initialized.
func NewPostSpeechandtextanalyticsTopicsParams() *PostSpeechandtextanalyticsTopicsParams {
	var ()
	return &PostSpeechandtextanalyticsTopicsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSpeechandtextanalyticsTopicsParamsWithTimeout creates a new PostSpeechandtextanalyticsTopicsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSpeechandtextanalyticsTopicsParamsWithTimeout(timeout time.Duration) *PostSpeechandtextanalyticsTopicsParams {
	var ()
	return &PostSpeechandtextanalyticsTopicsParams{

		timeout: timeout,
	}
}

// NewPostSpeechandtextanalyticsTopicsParamsWithContext creates a new PostSpeechandtextanalyticsTopicsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSpeechandtextanalyticsTopicsParamsWithContext(ctx context.Context) *PostSpeechandtextanalyticsTopicsParams {
	var ()
	return &PostSpeechandtextanalyticsTopicsParams{

		Context: ctx,
	}
}

// NewPostSpeechandtextanalyticsTopicsParamsWithHTTPClient creates a new PostSpeechandtextanalyticsTopicsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSpeechandtextanalyticsTopicsParamsWithHTTPClient(client *http.Client) *PostSpeechandtextanalyticsTopicsParams {
	var ()
	return &PostSpeechandtextanalyticsTopicsParams{
		HTTPClient: client,
	}
}

/*PostSpeechandtextanalyticsTopicsParams contains all the parameters to send to the API endpoint
for the post speechandtextanalytics topics operation typically these are written to a http.Request
*/
type PostSpeechandtextanalyticsTopicsParams struct {

	/*Body
	  The topic to create

	*/
	Body *models.TopicRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post speechandtextanalytics topics params
func (o *PostSpeechandtextanalyticsTopicsParams) WithTimeout(timeout time.Duration) *PostSpeechandtextanalyticsTopicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post speechandtextanalytics topics params
func (o *PostSpeechandtextanalyticsTopicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post speechandtextanalytics topics params
func (o *PostSpeechandtextanalyticsTopicsParams) WithContext(ctx context.Context) *PostSpeechandtextanalyticsTopicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post speechandtextanalytics topics params
func (o *PostSpeechandtextanalyticsTopicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post speechandtextanalytics topics params
func (o *PostSpeechandtextanalyticsTopicsParams) WithHTTPClient(client *http.Client) *PostSpeechandtextanalyticsTopicsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post speechandtextanalytics topics params
func (o *PostSpeechandtextanalyticsTopicsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post speechandtextanalytics topics params
func (o *PostSpeechandtextanalyticsTopicsParams) WithBody(body *models.TopicRequest) *PostSpeechandtextanalyticsTopicsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post speechandtextanalytics topics params
func (o *PostSpeechandtextanalyticsTopicsParams) SetBody(body *models.TopicRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSpeechandtextanalyticsTopicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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