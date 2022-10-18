// Code generated by go-swagger; DO NOT EDIT.

package presence

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

// NewGetPresenceSourceParams creates a new GetPresenceSourceParams object
// with the default values initialized.
func NewGetPresenceSourceParams() *GetPresenceSourceParams {
	var ()
	return &GetPresenceSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPresenceSourceParamsWithTimeout creates a new GetPresenceSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPresenceSourceParamsWithTimeout(timeout time.Duration) *GetPresenceSourceParams {
	var ()
	return &GetPresenceSourceParams{

		timeout: timeout,
	}
}

// NewGetPresenceSourceParamsWithContext creates a new GetPresenceSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPresenceSourceParamsWithContext(ctx context.Context) *GetPresenceSourceParams {
	var ()
	return &GetPresenceSourceParams{

		Context: ctx,
	}
}

// NewGetPresenceSourceParamsWithHTTPClient creates a new GetPresenceSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPresenceSourceParamsWithHTTPClient(client *http.Client) *GetPresenceSourceParams {
	var ()
	return &GetPresenceSourceParams{
		HTTPClient: client,
	}
}

/*GetPresenceSourceParams contains all the parameters to send to the API endpoint
for the get presence source operation typically these are written to a http.Request
*/
type GetPresenceSourceParams struct {

	/*SourceID
	  Presence Source ID

	*/
	SourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get presence source params
func (o *GetPresenceSourceParams) WithTimeout(timeout time.Duration) *GetPresenceSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get presence source params
func (o *GetPresenceSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get presence source params
func (o *GetPresenceSourceParams) WithContext(ctx context.Context) *GetPresenceSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get presence source params
func (o *GetPresenceSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get presence source params
func (o *GetPresenceSourceParams) WithHTTPClient(client *http.Client) *GetPresenceSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get presence source params
func (o *GetPresenceSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSourceID adds the sourceID to the get presence source params
func (o *GetPresenceSourceParams) WithSourceID(sourceID string) *GetPresenceSourceParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the get presence source params
func (o *GetPresenceSourceParams) SetSourceID(sourceID string) {
	o.SourceID = sourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPresenceSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sourceId
	if err := r.SetPathParam("sourceId", o.SourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
