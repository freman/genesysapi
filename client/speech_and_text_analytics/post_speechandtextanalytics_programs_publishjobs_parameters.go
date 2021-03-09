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

// NewPostSpeechandtextanalyticsProgramsPublishjobsParams creates a new PostSpeechandtextanalyticsProgramsPublishjobsParams object
// with the default values initialized.
func NewPostSpeechandtextanalyticsProgramsPublishjobsParams() *PostSpeechandtextanalyticsProgramsPublishjobsParams {
	var ()
	return &PostSpeechandtextanalyticsProgramsPublishjobsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSpeechandtextanalyticsProgramsPublishjobsParamsWithTimeout creates a new PostSpeechandtextanalyticsProgramsPublishjobsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSpeechandtextanalyticsProgramsPublishjobsParamsWithTimeout(timeout time.Duration) *PostSpeechandtextanalyticsProgramsPublishjobsParams {
	var ()
	return &PostSpeechandtextanalyticsProgramsPublishjobsParams{

		timeout: timeout,
	}
}

// NewPostSpeechandtextanalyticsProgramsPublishjobsParamsWithContext creates a new PostSpeechandtextanalyticsProgramsPublishjobsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSpeechandtextanalyticsProgramsPublishjobsParamsWithContext(ctx context.Context) *PostSpeechandtextanalyticsProgramsPublishjobsParams {
	var ()
	return &PostSpeechandtextanalyticsProgramsPublishjobsParams{

		Context: ctx,
	}
}

// NewPostSpeechandtextanalyticsProgramsPublishjobsParamsWithHTTPClient creates a new PostSpeechandtextanalyticsProgramsPublishjobsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSpeechandtextanalyticsProgramsPublishjobsParamsWithHTTPClient(client *http.Client) *PostSpeechandtextanalyticsProgramsPublishjobsParams {
	var ()
	return &PostSpeechandtextanalyticsProgramsPublishjobsParams{
		HTTPClient: client,
	}
}

/*PostSpeechandtextanalyticsProgramsPublishjobsParams contains all the parameters to send to the API endpoint
for the post speechandtextanalytics programs publishjobs operation typically these are written to a http.Request
*/
type PostSpeechandtextanalyticsProgramsPublishjobsParams struct {

	/*Body
	  The publish programs job to create

	*/
	Body *models.ProgramJobRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post speechandtextanalytics programs publishjobs params
func (o *PostSpeechandtextanalyticsProgramsPublishjobsParams) WithTimeout(timeout time.Duration) *PostSpeechandtextanalyticsProgramsPublishjobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post speechandtextanalytics programs publishjobs params
func (o *PostSpeechandtextanalyticsProgramsPublishjobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post speechandtextanalytics programs publishjobs params
func (o *PostSpeechandtextanalyticsProgramsPublishjobsParams) WithContext(ctx context.Context) *PostSpeechandtextanalyticsProgramsPublishjobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post speechandtextanalytics programs publishjobs params
func (o *PostSpeechandtextanalyticsProgramsPublishjobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post speechandtextanalytics programs publishjobs params
func (o *PostSpeechandtextanalyticsProgramsPublishjobsParams) WithHTTPClient(client *http.Client) *PostSpeechandtextanalyticsProgramsPublishjobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post speechandtextanalytics programs publishjobs params
func (o *PostSpeechandtextanalyticsProgramsPublishjobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post speechandtextanalytics programs publishjobs params
func (o *PostSpeechandtextanalyticsProgramsPublishjobsParams) WithBody(body *models.ProgramJobRequest) *PostSpeechandtextanalyticsProgramsPublishjobsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post speechandtextanalytics programs publishjobs params
func (o *PostSpeechandtextanalyticsProgramsPublishjobsParams) SetBody(body *models.ProgramJobRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSpeechandtextanalyticsProgramsPublishjobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
