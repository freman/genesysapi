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

// NewPutSpeechandtextanalyticsTopicParams creates a new PutSpeechandtextanalyticsTopicParams object
// with the default values initialized.
func NewPutSpeechandtextanalyticsTopicParams() *PutSpeechandtextanalyticsTopicParams {
	var ()
	return &PutSpeechandtextanalyticsTopicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSpeechandtextanalyticsTopicParamsWithTimeout creates a new PutSpeechandtextanalyticsTopicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSpeechandtextanalyticsTopicParamsWithTimeout(timeout time.Duration) *PutSpeechandtextanalyticsTopicParams {
	var ()
	return &PutSpeechandtextanalyticsTopicParams{

		timeout: timeout,
	}
}

// NewPutSpeechandtextanalyticsTopicParamsWithContext creates a new PutSpeechandtextanalyticsTopicParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSpeechandtextanalyticsTopicParamsWithContext(ctx context.Context) *PutSpeechandtextanalyticsTopicParams {
	var ()
	return &PutSpeechandtextanalyticsTopicParams{

		Context: ctx,
	}
}

// NewPutSpeechandtextanalyticsTopicParamsWithHTTPClient creates a new PutSpeechandtextanalyticsTopicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSpeechandtextanalyticsTopicParamsWithHTTPClient(client *http.Client) *PutSpeechandtextanalyticsTopicParams {
	var ()
	return &PutSpeechandtextanalyticsTopicParams{
		HTTPClient: client,
	}
}

/*PutSpeechandtextanalyticsTopicParams contains all the parameters to send to the API endpoint
for the put speechandtextanalytics topic operation typically these are written to a http.Request
*/
type PutSpeechandtextanalyticsTopicParams struct {

	/*Body
	  The topic to update

	*/
	Body *models.TopicRequest
	/*TopicID
	  The id of the topic

	*/
	TopicID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put speechandtextanalytics topic params
func (o *PutSpeechandtextanalyticsTopicParams) WithTimeout(timeout time.Duration) *PutSpeechandtextanalyticsTopicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put speechandtextanalytics topic params
func (o *PutSpeechandtextanalyticsTopicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put speechandtextanalytics topic params
func (o *PutSpeechandtextanalyticsTopicParams) WithContext(ctx context.Context) *PutSpeechandtextanalyticsTopicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put speechandtextanalytics topic params
func (o *PutSpeechandtextanalyticsTopicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put speechandtextanalytics topic params
func (o *PutSpeechandtextanalyticsTopicParams) WithHTTPClient(client *http.Client) *PutSpeechandtextanalyticsTopicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put speechandtextanalytics topic params
func (o *PutSpeechandtextanalyticsTopicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put speechandtextanalytics topic params
func (o *PutSpeechandtextanalyticsTopicParams) WithBody(body *models.TopicRequest) *PutSpeechandtextanalyticsTopicParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put speechandtextanalytics topic params
func (o *PutSpeechandtextanalyticsTopicParams) SetBody(body *models.TopicRequest) {
	o.Body = body
}

// WithTopicID adds the topicID to the put speechandtextanalytics topic params
func (o *PutSpeechandtextanalyticsTopicParams) WithTopicID(topicID string) *PutSpeechandtextanalyticsTopicParams {
	o.SetTopicID(topicID)
	return o
}

// SetTopicID adds the topicId to the put speechandtextanalytics topic params
func (o *PutSpeechandtextanalyticsTopicParams) SetTopicID(topicID string) {
	o.TopicID = topicID
}

// WriteToRequest writes these params to a swagger request
func (o *PutSpeechandtextanalyticsTopicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param topicId
	if err := r.SetPathParam("topicId", o.TopicID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
