// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

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

// NewGetWorkforcemanagementAdhocmodelingjobParams creates a new GetWorkforcemanagementAdhocmodelingjobParams object
// with the default values initialized.
func NewGetWorkforcemanagementAdhocmodelingjobParams() *GetWorkforcemanagementAdhocmodelingjobParams {
	var ()
	return &GetWorkforcemanagementAdhocmodelingjobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkforcemanagementAdhocmodelingjobParamsWithTimeout creates a new GetWorkforcemanagementAdhocmodelingjobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkforcemanagementAdhocmodelingjobParamsWithTimeout(timeout time.Duration) *GetWorkforcemanagementAdhocmodelingjobParams {
	var ()
	return &GetWorkforcemanagementAdhocmodelingjobParams{

		timeout: timeout,
	}
}

// NewGetWorkforcemanagementAdhocmodelingjobParamsWithContext creates a new GetWorkforcemanagementAdhocmodelingjobParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkforcemanagementAdhocmodelingjobParamsWithContext(ctx context.Context) *GetWorkforcemanagementAdhocmodelingjobParams {
	var ()
	return &GetWorkforcemanagementAdhocmodelingjobParams{

		Context: ctx,
	}
}

// NewGetWorkforcemanagementAdhocmodelingjobParamsWithHTTPClient creates a new GetWorkforcemanagementAdhocmodelingjobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkforcemanagementAdhocmodelingjobParamsWithHTTPClient(client *http.Client) *GetWorkforcemanagementAdhocmodelingjobParams {
	var ()
	return &GetWorkforcemanagementAdhocmodelingjobParams{
		HTTPClient: client,
	}
}

/*GetWorkforcemanagementAdhocmodelingjobParams contains all the parameters to send to the API endpoint
for the get workforcemanagement adhocmodelingjob operation typically these are written to a http.Request
*/
type GetWorkforcemanagementAdhocmodelingjobParams struct {

	/*JobID
	  The id of the modeling job

	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workforcemanagement adhocmodelingjob params
func (o *GetWorkforcemanagementAdhocmodelingjobParams) WithTimeout(timeout time.Duration) *GetWorkforcemanagementAdhocmodelingjobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workforcemanagement adhocmodelingjob params
func (o *GetWorkforcemanagementAdhocmodelingjobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workforcemanagement adhocmodelingjob params
func (o *GetWorkforcemanagementAdhocmodelingjobParams) WithContext(ctx context.Context) *GetWorkforcemanagementAdhocmodelingjobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workforcemanagement adhocmodelingjob params
func (o *GetWorkforcemanagementAdhocmodelingjobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workforcemanagement adhocmodelingjob params
func (o *GetWorkforcemanagementAdhocmodelingjobParams) WithHTTPClient(client *http.Client) *GetWorkforcemanagementAdhocmodelingjobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workforcemanagement adhocmodelingjob params
func (o *GetWorkforcemanagementAdhocmodelingjobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the get workforcemanagement adhocmodelingjob params
func (o *GetWorkforcemanagementAdhocmodelingjobParams) WithJobID(jobID string) *GetWorkforcemanagementAdhocmodelingjobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get workforcemanagement adhocmodelingjob params
func (o *GetWorkforcemanagementAdhocmodelingjobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkforcemanagementAdhocmodelingjobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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