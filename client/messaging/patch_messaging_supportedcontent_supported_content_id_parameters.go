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

	"github.com/freman/genesysapi/models"
)

// NewPatchMessagingSupportedcontentSupportedContentIDParams creates a new PatchMessagingSupportedcontentSupportedContentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchMessagingSupportedcontentSupportedContentIDParams() *PatchMessagingSupportedcontentSupportedContentIDParams {
	return &PatchMessagingSupportedcontentSupportedContentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchMessagingSupportedcontentSupportedContentIDParamsWithTimeout creates a new PatchMessagingSupportedcontentSupportedContentIDParams object
// with the ability to set a timeout on a request.
func NewPatchMessagingSupportedcontentSupportedContentIDParamsWithTimeout(timeout time.Duration) *PatchMessagingSupportedcontentSupportedContentIDParams {
	return &PatchMessagingSupportedcontentSupportedContentIDParams{
		timeout: timeout,
	}
}

// NewPatchMessagingSupportedcontentSupportedContentIDParamsWithContext creates a new PatchMessagingSupportedcontentSupportedContentIDParams object
// with the ability to set a context for a request.
func NewPatchMessagingSupportedcontentSupportedContentIDParamsWithContext(ctx context.Context) *PatchMessagingSupportedcontentSupportedContentIDParams {
	return &PatchMessagingSupportedcontentSupportedContentIDParams{
		Context: ctx,
	}
}

// NewPatchMessagingSupportedcontentSupportedContentIDParamsWithHTTPClient creates a new PatchMessagingSupportedcontentSupportedContentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchMessagingSupportedcontentSupportedContentIDParamsWithHTTPClient(client *http.Client) *PatchMessagingSupportedcontentSupportedContentIDParams {
	return &PatchMessagingSupportedcontentSupportedContentIDParams{
		HTTPClient: client,
	}
}

/*
PatchMessagingSupportedcontentSupportedContentIDParams contains all the parameters to send to the API endpoint

	for the patch messaging supportedcontent supported content Id operation.

	Typically these are written to a http.Request.
*/
type PatchMessagingSupportedcontentSupportedContentIDParams struct {

	/* Body.

	   SupportedContent
	*/
	Body *models.SupportedContent

	/* SupportedContentID.

	   Supported Content ID
	*/
	SupportedContentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch messaging supportedcontent supported content Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) WithDefaults() *PatchMessagingSupportedcontentSupportedContentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch messaging supportedcontent supported content Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch messaging supportedcontent supported content Id params
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) WithTimeout(timeout time.Duration) *PatchMessagingSupportedcontentSupportedContentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch messaging supportedcontent supported content Id params
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch messaging supportedcontent supported content Id params
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) WithContext(ctx context.Context) *PatchMessagingSupportedcontentSupportedContentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch messaging supportedcontent supported content Id params
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch messaging supportedcontent supported content Id params
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) WithHTTPClient(client *http.Client) *PatchMessagingSupportedcontentSupportedContentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch messaging supportedcontent supported content Id params
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch messaging supportedcontent supported content Id params
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) WithBody(body *models.SupportedContent) *PatchMessagingSupportedcontentSupportedContentIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch messaging supportedcontent supported content Id params
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) SetBody(body *models.SupportedContent) {
	o.Body = body
}

// WithSupportedContentID adds the supportedContentID to the patch messaging supportedcontent supported content Id params
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) WithSupportedContentID(supportedContentID string) *PatchMessagingSupportedcontentSupportedContentIDParams {
	o.SetSupportedContentID(supportedContentID)
	return o
}

// SetSupportedContentID adds the supportedContentId to the patch messaging supportedcontent supported content Id params
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) SetSupportedContentID(supportedContentID string) {
	o.SupportedContentID = supportedContentID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchMessagingSupportedcontentSupportedContentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param supportedContentId
	if err := r.SetPathParam("supportedContentId", o.SupportedContentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
