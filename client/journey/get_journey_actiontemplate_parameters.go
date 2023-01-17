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

// NewGetJourneyActiontemplateParams creates a new GetJourneyActiontemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetJourneyActiontemplateParams() *GetJourneyActiontemplateParams {
	return &GetJourneyActiontemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetJourneyActiontemplateParamsWithTimeout creates a new GetJourneyActiontemplateParams object
// with the ability to set a timeout on a request.
func NewGetJourneyActiontemplateParamsWithTimeout(timeout time.Duration) *GetJourneyActiontemplateParams {
	return &GetJourneyActiontemplateParams{
		timeout: timeout,
	}
}

// NewGetJourneyActiontemplateParamsWithContext creates a new GetJourneyActiontemplateParams object
// with the ability to set a context for a request.
func NewGetJourneyActiontemplateParamsWithContext(ctx context.Context) *GetJourneyActiontemplateParams {
	return &GetJourneyActiontemplateParams{
		Context: ctx,
	}
}

// NewGetJourneyActiontemplateParamsWithHTTPClient creates a new GetJourneyActiontemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetJourneyActiontemplateParamsWithHTTPClient(client *http.Client) *GetJourneyActiontemplateParams {
	return &GetJourneyActiontemplateParams{
		HTTPClient: client,
	}
}

/*
GetJourneyActiontemplateParams contains all the parameters to send to the API endpoint

	for the get journey actiontemplate operation.

	Typically these are written to a http.Request.
*/
type GetJourneyActiontemplateParams struct {

	/* ActionTemplateID.

	   ID of the action template.
	*/
	ActionTemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get journey actiontemplate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJourneyActiontemplateParams) WithDefaults() *GetJourneyActiontemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get journey actiontemplate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJourneyActiontemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get journey actiontemplate params
func (o *GetJourneyActiontemplateParams) WithTimeout(timeout time.Duration) *GetJourneyActiontemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get journey actiontemplate params
func (o *GetJourneyActiontemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get journey actiontemplate params
func (o *GetJourneyActiontemplateParams) WithContext(ctx context.Context) *GetJourneyActiontemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get journey actiontemplate params
func (o *GetJourneyActiontemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get journey actiontemplate params
func (o *GetJourneyActiontemplateParams) WithHTTPClient(client *http.Client) *GetJourneyActiontemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get journey actiontemplate params
func (o *GetJourneyActiontemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionTemplateID adds the actionTemplateID to the get journey actiontemplate params
func (o *GetJourneyActiontemplateParams) WithActionTemplateID(actionTemplateID string) *GetJourneyActiontemplateParams {
	o.SetActionTemplateID(actionTemplateID)
	return o
}

// SetActionTemplateID adds the actionTemplateId to the get journey actiontemplate params
func (o *GetJourneyActiontemplateParams) SetActionTemplateID(actionTemplateID string) {
	o.ActionTemplateID = actionTemplateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetJourneyActiontemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionTemplateId
	if err := r.SetPathParam("actionTemplateId", o.ActionTemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
