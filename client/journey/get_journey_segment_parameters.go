// Code generated by go-swagger; DO NOT EDIT.

package journey

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

// NewGetJourneySegmentParams creates a new GetJourneySegmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetJourneySegmentParams() *GetJourneySegmentParams {
	return &GetJourneySegmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetJourneySegmentParamsWithTimeout creates a new GetJourneySegmentParams object
// with the ability to set a timeout on a request.
func NewGetJourneySegmentParamsWithTimeout(timeout time.Duration) *GetJourneySegmentParams {
	return &GetJourneySegmentParams{
		timeout: timeout,
	}
}

// NewGetJourneySegmentParamsWithContext creates a new GetJourneySegmentParams object
// with the ability to set a context for a request.
func NewGetJourneySegmentParamsWithContext(ctx context.Context) *GetJourneySegmentParams {
	return &GetJourneySegmentParams{
		Context: ctx,
	}
}

// NewGetJourneySegmentParamsWithHTTPClient creates a new GetJourneySegmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetJourneySegmentParamsWithHTTPClient(client *http.Client) *GetJourneySegmentParams {
	return &GetJourneySegmentParams{
		HTTPClient: client,
	}
}

/*
GetJourneySegmentParams contains all the parameters to send to the API endpoint

	for the get journey segment operation.

	Typically these are written to a http.Request.
*/
type GetJourneySegmentParams struct {

	/* SegmentID.

	   ID of the segment.
	*/
	SegmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get journey segment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJourneySegmentParams) WithDefaults() *GetJourneySegmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get journey segment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJourneySegmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get journey segment params
func (o *GetJourneySegmentParams) WithTimeout(timeout time.Duration) *GetJourneySegmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get journey segment params
func (o *GetJourneySegmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get journey segment params
func (o *GetJourneySegmentParams) WithContext(ctx context.Context) *GetJourneySegmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get journey segment params
func (o *GetJourneySegmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get journey segment params
func (o *GetJourneySegmentParams) WithHTTPClient(client *http.Client) *GetJourneySegmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get journey segment params
func (o *GetJourneySegmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSegmentID adds the segmentID to the get journey segment params
func (o *GetJourneySegmentParams) WithSegmentID(segmentID string) *GetJourneySegmentParams {
	o.SetSegmentID(segmentID)
	return o
}

// SetSegmentID adds the segmentId to the get journey segment params
func (o *GetJourneySegmentParams) SetSegmentID(segmentID string) {
	o.SegmentID = segmentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetJourneySegmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param segmentId
	if err := r.SetPathParam("segmentId", o.SegmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
