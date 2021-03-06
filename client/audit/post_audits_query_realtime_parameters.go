// Code generated by go-swagger; DO NOT EDIT.

package audit

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
	"github.com/go-openapi/swag"

	"github.com/freman/genesysapi/models"
)

// NewPostAuditsQueryRealtimeParams creates a new PostAuditsQueryRealtimeParams object
// with the default values initialized.
func NewPostAuditsQueryRealtimeParams() *PostAuditsQueryRealtimeParams {
	var ()
	return &PostAuditsQueryRealtimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuditsQueryRealtimeParamsWithTimeout creates a new PostAuditsQueryRealtimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAuditsQueryRealtimeParamsWithTimeout(timeout time.Duration) *PostAuditsQueryRealtimeParams {
	var ()
	return &PostAuditsQueryRealtimeParams{

		timeout: timeout,
	}
}

// NewPostAuditsQueryRealtimeParamsWithContext creates a new PostAuditsQueryRealtimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAuditsQueryRealtimeParamsWithContext(ctx context.Context) *PostAuditsQueryRealtimeParams {
	var ()
	return &PostAuditsQueryRealtimeParams{

		Context: ctx,
	}
}

// NewPostAuditsQueryRealtimeParamsWithHTTPClient creates a new PostAuditsQueryRealtimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAuditsQueryRealtimeParamsWithHTTPClient(client *http.Client) *PostAuditsQueryRealtimeParams {
	var ()
	return &PostAuditsQueryRealtimeParams{
		HTTPClient: client,
	}
}

/*PostAuditsQueryRealtimeParams contains all the parameters to send to the API endpoint
for the post audits query realtime operation typically these are written to a http.Request
*/
type PostAuditsQueryRealtimeParams struct {

	/*Body
	  query

	*/
	Body *models.AuditRealtimeQueryRequest
	/*Expand
	  Which fields, if any, to expand

	*/
	Expand []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post audits query realtime params
func (o *PostAuditsQueryRealtimeParams) WithTimeout(timeout time.Duration) *PostAuditsQueryRealtimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post audits query realtime params
func (o *PostAuditsQueryRealtimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post audits query realtime params
func (o *PostAuditsQueryRealtimeParams) WithContext(ctx context.Context) *PostAuditsQueryRealtimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post audits query realtime params
func (o *PostAuditsQueryRealtimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post audits query realtime params
func (o *PostAuditsQueryRealtimeParams) WithHTTPClient(client *http.Client) *PostAuditsQueryRealtimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post audits query realtime params
func (o *PostAuditsQueryRealtimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post audits query realtime params
func (o *PostAuditsQueryRealtimeParams) WithBody(body *models.AuditRealtimeQueryRequest) *PostAuditsQueryRealtimeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post audits query realtime params
func (o *PostAuditsQueryRealtimeParams) SetBody(body *models.AuditRealtimeQueryRequest) {
	o.Body = body
}

// WithExpand adds the expand to the post audits query realtime params
func (o *PostAuditsQueryRealtimeParams) WithExpand(expand []string) *PostAuditsQueryRealtimeParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the post audits query realtime params
func (o *PostAuditsQueryRealtimeParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuditsQueryRealtimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
