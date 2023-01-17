// Code generated by go-swagger; DO NOT EDIT.

package messaging

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

// NewGetMessagingSupportedcontentSupportedContentIDParams creates a new GetMessagingSupportedcontentSupportedContentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMessagingSupportedcontentSupportedContentIDParams() *GetMessagingSupportedcontentSupportedContentIDParams {
	return &GetMessagingSupportedcontentSupportedContentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMessagingSupportedcontentSupportedContentIDParamsWithTimeout creates a new GetMessagingSupportedcontentSupportedContentIDParams object
// with the ability to set a timeout on a request.
func NewGetMessagingSupportedcontentSupportedContentIDParamsWithTimeout(timeout time.Duration) *GetMessagingSupportedcontentSupportedContentIDParams {
	return &GetMessagingSupportedcontentSupportedContentIDParams{
		timeout: timeout,
	}
}

// NewGetMessagingSupportedcontentSupportedContentIDParamsWithContext creates a new GetMessagingSupportedcontentSupportedContentIDParams object
// with the ability to set a context for a request.
func NewGetMessagingSupportedcontentSupportedContentIDParamsWithContext(ctx context.Context) *GetMessagingSupportedcontentSupportedContentIDParams {
	return &GetMessagingSupportedcontentSupportedContentIDParams{
		Context: ctx,
	}
}

// NewGetMessagingSupportedcontentSupportedContentIDParamsWithHTTPClient creates a new GetMessagingSupportedcontentSupportedContentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMessagingSupportedcontentSupportedContentIDParamsWithHTTPClient(client *http.Client) *GetMessagingSupportedcontentSupportedContentIDParams {
	return &GetMessagingSupportedcontentSupportedContentIDParams{
		HTTPClient: client,
	}
}

/*
GetMessagingSupportedcontentSupportedContentIDParams contains all the parameters to send to the API endpoint

	for the get messaging supportedcontent supported content Id operation.

	Typically these are written to a http.Request.
*/
type GetMessagingSupportedcontentSupportedContentIDParams struct {

	/* SupportedContentID.

	   Supported Content ID
	*/
	SupportedContentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get messaging supportedcontent supported content Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMessagingSupportedcontentSupportedContentIDParams) WithDefaults() *GetMessagingSupportedcontentSupportedContentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get messaging supportedcontent supported content Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMessagingSupportedcontentSupportedContentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get messaging supportedcontent supported content Id params
func (o *GetMessagingSupportedcontentSupportedContentIDParams) WithTimeout(timeout time.Duration) *GetMessagingSupportedcontentSupportedContentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get messaging supportedcontent supported content Id params
func (o *GetMessagingSupportedcontentSupportedContentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get messaging supportedcontent supported content Id params
func (o *GetMessagingSupportedcontentSupportedContentIDParams) WithContext(ctx context.Context) *GetMessagingSupportedcontentSupportedContentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get messaging supportedcontent supported content Id params
func (o *GetMessagingSupportedcontentSupportedContentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get messaging supportedcontent supported content Id params
func (o *GetMessagingSupportedcontentSupportedContentIDParams) WithHTTPClient(client *http.Client) *GetMessagingSupportedcontentSupportedContentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get messaging supportedcontent supported content Id params
func (o *GetMessagingSupportedcontentSupportedContentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSupportedContentID adds the supportedContentID to the get messaging supportedcontent supported content Id params
func (o *GetMessagingSupportedcontentSupportedContentIDParams) WithSupportedContentID(supportedContentID string) *GetMessagingSupportedcontentSupportedContentIDParams {
	o.SetSupportedContentID(supportedContentID)
	return o
}

// SetSupportedContentID adds the supportedContentId to the get messaging supportedcontent supported content Id params
func (o *GetMessagingSupportedcontentSupportedContentIDParams) SetSupportedContentID(supportedContentID string) {
	o.SupportedContentID = supportedContentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMessagingSupportedcontentSupportedContentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param supportedContentId
	if err := r.SetPathParam("supportedContentId", o.SupportedContentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
