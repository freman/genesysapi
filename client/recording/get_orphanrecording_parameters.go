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

// NewGetOrphanrecordingParams creates a new GetOrphanrecordingParams object
// with the default values initialized.
func NewGetOrphanrecordingParams() *GetOrphanrecordingParams {
	var ()
	return &GetOrphanrecordingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrphanrecordingParamsWithTimeout creates a new GetOrphanrecordingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrphanrecordingParamsWithTimeout(timeout time.Duration) *GetOrphanrecordingParams {
	var ()
	return &GetOrphanrecordingParams{

		timeout: timeout,
	}
}

// NewGetOrphanrecordingParamsWithContext creates a new GetOrphanrecordingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrphanrecordingParamsWithContext(ctx context.Context) *GetOrphanrecordingParams {
	var ()
	return &GetOrphanrecordingParams{

		Context: ctx,
	}
}

// NewGetOrphanrecordingParamsWithHTTPClient creates a new GetOrphanrecordingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrphanrecordingParamsWithHTTPClient(client *http.Client) *GetOrphanrecordingParams {
	var ()
	return &GetOrphanrecordingParams{
		HTTPClient: client,
	}
}

/*GetOrphanrecordingParams contains all the parameters to send to the API endpoint
for the get orphanrecording operation typically these are written to a http.Request
*/
type GetOrphanrecordingParams struct {

	/*OrphanID
	  Orphan ID

	*/
	OrphanID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get orphanrecording params
func (o *GetOrphanrecordingParams) WithTimeout(timeout time.Duration) *GetOrphanrecordingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orphanrecording params
func (o *GetOrphanrecordingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orphanrecording params
func (o *GetOrphanrecordingParams) WithContext(ctx context.Context) *GetOrphanrecordingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orphanrecording params
func (o *GetOrphanrecordingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orphanrecording params
func (o *GetOrphanrecordingParams) WithHTTPClient(client *http.Client) *GetOrphanrecordingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orphanrecording params
func (o *GetOrphanrecordingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrphanID adds the orphanID to the get orphanrecording params
func (o *GetOrphanrecordingParams) WithOrphanID(orphanID string) *GetOrphanrecordingParams {
	o.SetOrphanID(orphanID)
	return o
}

// SetOrphanID adds the orphanId to the get orphanrecording params
func (o *GetOrphanrecordingParams) SetOrphanID(orphanID string) {
	o.OrphanID = orphanID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrphanrecordingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orphanId
	if err := r.SetPathParam("orphanId", o.OrphanID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
