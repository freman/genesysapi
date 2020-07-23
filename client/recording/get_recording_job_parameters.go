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
)

// NewGetRecordingJobParams creates a new GetRecordingJobParams object
// with the default values initialized.
func NewGetRecordingJobParams() *GetRecordingJobParams {
	var ()
	return &GetRecordingJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecordingJobParamsWithTimeout creates a new GetRecordingJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecordingJobParamsWithTimeout(timeout time.Duration) *GetRecordingJobParams {
	var ()
	return &GetRecordingJobParams{

		timeout: timeout,
	}
}

// NewGetRecordingJobParamsWithContext creates a new GetRecordingJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecordingJobParamsWithContext(ctx context.Context) *GetRecordingJobParams {
	var ()
	return &GetRecordingJobParams{

		Context: ctx,
	}
}

// NewGetRecordingJobParamsWithHTTPClient creates a new GetRecordingJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecordingJobParamsWithHTTPClient(client *http.Client) *GetRecordingJobParams {
	var ()
	return &GetRecordingJobParams{
		HTTPClient: client,
	}
}

/*GetRecordingJobParams contains all the parameters to send to the API endpoint
for the get recording job operation typically these are written to a http.Request
*/
type GetRecordingJobParams struct {

	/*JobID
	  jobId

	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recording job params
func (o *GetRecordingJobParams) WithTimeout(timeout time.Duration) *GetRecordingJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recording job params
func (o *GetRecordingJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recording job params
func (o *GetRecordingJobParams) WithContext(ctx context.Context) *GetRecordingJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recording job params
func (o *GetRecordingJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recording job params
func (o *GetRecordingJobParams) WithHTTPClient(client *http.Client) *GetRecordingJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recording job params
func (o *GetRecordingJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the get recording job params
func (o *GetRecordingJobParams) WithJobID(jobID string) *GetRecordingJobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get recording job params
func (o *GetRecordingJobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecordingJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param jobId
	if err := r.SetPathParam("jobId", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
