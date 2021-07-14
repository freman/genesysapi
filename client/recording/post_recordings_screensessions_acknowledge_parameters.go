// Code generated by go-swagger; DO NOT EDIT.

package recording

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

// NewPostRecordingsScreensessionsAcknowledgeParams creates a new PostRecordingsScreensessionsAcknowledgeParams object
// with the default values initialized.
func NewPostRecordingsScreensessionsAcknowledgeParams() *PostRecordingsScreensessionsAcknowledgeParams {
	var ()
	return &PostRecordingsScreensessionsAcknowledgeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRecordingsScreensessionsAcknowledgeParamsWithTimeout creates a new PostRecordingsScreensessionsAcknowledgeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRecordingsScreensessionsAcknowledgeParamsWithTimeout(timeout time.Duration) *PostRecordingsScreensessionsAcknowledgeParams {
	var ()
	return &PostRecordingsScreensessionsAcknowledgeParams{

		timeout: timeout,
	}
}

// NewPostRecordingsScreensessionsAcknowledgeParamsWithContext creates a new PostRecordingsScreensessionsAcknowledgeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRecordingsScreensessionsAcknowledgeParamsWithContext(ctx context.Context) *PostRecordingsScreensessionsAcknowledgeParams {
	var ()
	return &PostRecordingsScreensessionsAcknowledgeParams{

		Context: ctx,
	}
}

// NewPostRecordingsScreensessionsAcknowledgeParamsWithHTTPClient creates a new PostRecordingsScreensessionsAcknowledgeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRecordingsScreensessionsAcknowledgeParamsWithHTTPClient(client *http.Client) *PostRecordingsScreensessionsAcknowledgeParams {
	var ()
	return &PostRecordingsScreensessionsAcknowledgeParams{
		HTTPClient: client,
	}
}

/*PostRecordingsScreensessionsAcknowledgeParams contains all the parameters to send to the API endpoint
for the post recordings screensessions acknowledge operation typically these are written to a http.Request
*/
type PostRecordingsScreensessionsAcknowledgeParams struct {

	/*Body
	  AcknowledgeScreenRecordingRequest

	*/
	Body *models.AcknowledgeScreenRecordingRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post recordings screensessions acknowledge params
func (o *PostRecordingsScreensessionsAcknowledgeParams) WithTimeout(timeout time.Duration) *PostRecordingsScreensessionsAcknowledgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post recordings screensessions acknowledge params
func (o *PostRecordingsScreensessionsAcknowledgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post recordings screensessions acknowledge params
func (o *PostRecordingsScreensessionsAcknowledgeParams) WithContext(ctx context.Context) *PostRecordingsScreensessionsAcknowledgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post recordings screensessions acknowledge params
func (o *PostRecordingsScreensessionsAcknowledgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post recordings screensessions acknowledge params
func (o *PostRecordingsScreensessionsAcknowledgeParams) WithHTTPClient(client *http.Client) *PostRecordingsScreensessionsAcknowledgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post recordings screensessions acknowledge params
func (o *PostRecordingsScreensessionsAcknowledgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post recordings screensessions acknowledge params
func (o *PostRecordingsScreensessionsAcknowledgeParams) WithBody(body *models.AcknowledgeScreenRecordingRequest) *PostRecordingsScreensessionsAcknowledgeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post recordings screensessions acknowledge params
func (o *PostRecordingsScreensessionsAcknowledgeParams) SetBody(body *models.AcknowledgeScreenRecordingRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRecordingsScreensessionsAcknowledgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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