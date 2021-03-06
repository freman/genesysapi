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

// NewPutOrphanrecordingParams creates a new PutOrphanrecordingParams object
// with the default values initialized.
func NewPutOrphanrecordingParams() *PutOrphanrecordingParams {
	var ()
	return &PutOrphanrecordingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutOrphanrecordingParamsWithTimeout creates a new PutOrphanrecordingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutOrphanrecordingParamsWithTimeout(timeout time.Duration) *PutOrphanrecordingParams {
	var ()
	return &PutOrphanrecordingParams{

		timeout: timeout,
	}
}

// NewPutOrphanrecordingParamsWithContext creates a new PutOrphanrecordingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutOrphanrecordingParamsWithContext(ctx context.Context) *PutOrphanrecordingParams {
	var ()
	return &PutOrphanrecordingParams{

		Context: ctx,
	}
}

// NewPutOrphanrecordingParamsWithHTTPClient creates a new PutOrphanrecordingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutOrphanrecordingParamsWithHTTPClient(client *http.Client) *PutOrphanrecordingParams {
	var ()
	return &PutOrphanrecordingParams{
		HTTPClient: client,
	}
}

/*PutOrphanrecordingParams contains all the parameters to send to the API endpoint
for the put orphanrecording operation typically these are written to a http.Request
*/
type PutOrphanrecordingParams struct {

	/*Body*/
	Body *models.OrphanUpdateRequest
	/*OrphanID
	  Orphan ID

	*/
	OrphanID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put orphanrecording params
func (o *PutOrphanrecordingParams) WithTimeout(timeout time.Duration) *PutOrphanrecordingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put orphanrecording params
func (o *PutOrphanrecordingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put orphanrecording params
func (o *PutOrphanrecordingParams) WithContext(ctx context.Context) *PutOrphanrecordingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put orphanrecording params
func (o *PutOrphanrecordingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put orphanrecording params
func (o *PutOrphanrecordingParams) WithHTTPClient(client *http.Client) *PutOrphanrecordingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put orphanrecording params
func (o *PutOrphanrecordingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put orphanrecording params
func (o *PutOrphanrecordingParams) WithBody(body *models.OrphanUpdateRequest) *PutOrphanrecordingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put orphanrecording params
func (o *PutOrphanrecordingParams) SetBody(body *models.OrphanUpdateRequest) {
	o.Body = body
}

// WithOrphanID adds the orphanID to the put orphanrecording params
func (o *PutOrphanrecordingParams) WithOrphanID(orphanID string) *PutOrphanrecordingParams {
	o.SetOrphanID(orphanID)
	return o
}

// SetOrphanID adds the orphanId to the put orphanrecording params
func (o *PutOrphanrecordingParams) SetOrphanID(orphanID string) {
	o.OrphanID = orphanID
}

// WriteToRequest writes these params to a swagger request
func (o *PutOrphanrecordingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param orphanId
	if err := r.SetPathParam("orphanId", o.OrphanID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
