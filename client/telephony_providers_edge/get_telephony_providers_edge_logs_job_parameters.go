// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

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

// NewGetTelephonyProvidersEdgeLogsJobParams creates a new GetTelephonyProvidersEdgeLogsJobParams object
// with the default values initialized.
func NewGetTelephonyProvidersEdgeLogsJobParams() *GetTelephonyProvidersEdgeLogsJobParams {
	var ()
	return &GetTelephonyProvidersEdgeLogsJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgeLogsJobParamsWithTimeout creates a new GetTelephonyProvidersEdgeLogsJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTelephonyProvidersEdgeLogsJobParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgeLogsJobParams {
	var ()
	return &GetTelephonyProvidersEdgeLogsJobParams{

		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgeLogsJobParamsWithContext creates a new GetTelephonyProvidersEdgeLogsJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTelephonyProvidersEdgeLogsJobParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgeLogsJobParams {
	var ()
	return &GetTelephonyProvidersEdgeLogsJobParams{

		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgeLogsJobParamsWithHTTPClient creates a new GetTelephonyProvidersEdgeLogsJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTelephonyProvidersEdgeLogsJobParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgeLogsJobParams {
	var ()
	return &GetTelephonyProvidersEdgeLogsJobParams{
		HTTPClient: client,
	}
}

/*GetTelephonyProvidersEdgeLogsJobParams contains all the parameters to send to the API endpoint
for the get telephony providers edge logs job operation typically these are written to a http.Request
*/
type GetTelephonyProvidersEdgeLogsJobParams struct {

	/*EdgeID
	  Edge ID

	*/
	EdgeID string
	/*JobID
	  Job ID

	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get telephony providers edge logs job params
func (o *GetTelephonyProvidersEdgeLogsJobParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgeLogsJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edge logs job params
func (o *GetTelephonyProvidersEdgeLogsJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edge logs job params
func (o *GetTelephonyProvidersEdgeLogsJobParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgeLogsJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edge logs job params
func (o *GetTelephonyProvidersEdgeLogsJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edge logs job params
func (o *GetTelephonyProvidersEdgeLogsJobParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgeLogsJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edge logs job params
func (o *GetTelephonyProvidersEdgeLogsJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEdgeID adds the edgeID to the get telephony providers edge logs job params
func (o *GetTelephonyProvidersEdgeLogsJobParams) WithEdgeID(edgeID string) *GetTelephonyProvidersEdgeLogsJobParams {
	o.SetEdgeID(edgeID)
	return o
}

// SetEdgeID adds the edgeId to the get telephony providers edge logs job params
func (o *GetTelephonyProvidersEdgeLogsJobParams) SetEdgeID(edgeID string) {
	o.EdgeID = edgeID
}

// WithJobID adds the jobID to the get telephony providers edge logs job params
func (o *GetTelephonyProvidersEdgeLogsJobParams) WithJobID(jobID string) *GetTelephonyProvidersEdgeLogsJobParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get telephony providers edge logs job params
func (o *GetTelephonyProvidersEdgeLogsJobParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgeLogsJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param edgeId
	if err := r.SetPathParam("edgeId", o.EdgeID); err != nil {
		return err
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