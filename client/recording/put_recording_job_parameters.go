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

// NewPutRecordingJobParams creates a new PutRecordingJobParams object
// with the default values initialized.
func NewPutRecordingJobParams() *PutRecordingJobParams {
	var ()
	return &PutRecordingJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRecordingJobParamsWithTimeout creates a new PutRecordingJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRecordingJobParamsWithTimeout(timeout time.Duration) *PutRecordingJobParams {
	var ()
	return &PutRecordingJobParams{

		timeout: timeout,
	}
}

// NewPutRecordingJobParamsWithContext creates a new PutRecordingJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRecordingJobParamsWithContext(ctx context.Context) *PutRecordingJobParams {
	var ()
	return &PutRecordingJobParams{

		Context: ctx,
	}
}

// NewPutRecordingJobParamsWithHTTPClient creates a new PutRecordingJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRecordingJobParamsWithHTTPClient(client *http.Client) *PutRecordingJobParams {
	var ()
	return &PutRecordingJobParams{
		HTTPClient: client,
	}
}

/*PutRecordingJobParams contains all the parameters to send to the API endpoint
for the put recording job operation typically these are written to a http.Request
*/
type PutRecordingJobParams struct {

	/*Body
	  query

	*/
	Body *models.ExecuteRecordingJobsQuery
	/*JobID
	  jobId

	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put recording job params
func (o *PutRecordingJobParams) WithTimeout(timeout time.Duration) *PutRecordingJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put recording job params
func (o *PutRecordingJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put recording job params
func (o *PutRecordingJobParams) WithContext(ctx context.Context) *PutRecordingJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put recording job params
func (o *PutRecordingJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put recording job params
func (o *PutRecordingJobParams) WithHTTPClient(client *http.Client) *PutRecordingJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put recording job params
func (o *PutRecordingJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put recording job params
func (o *PutRecordingJobParams) WithBody(body *models.ExecuteRecordingJobsQuery) *PutRecordingJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put recording job params
func (o *PutRecordingJobParams) SetBody(body *models.ExecuteRecordingJobsQuery) {
	o.Body = body
}

// WithJobID adds the jobID to the put recording job params
func (o *PutRecordingJobParams) WithJobID(jobID string) *PutRecordingJobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the put recording job params
func (o *PutRecordingJobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *PutRecordingJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param jobId
	if err := r.SetPathParam("jobId", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}