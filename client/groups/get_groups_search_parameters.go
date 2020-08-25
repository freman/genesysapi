// Code generated by go-swagger; DO NOT EDIT.

package groups

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
)

// NewGetGroupsSearchParams creates a new GetGroupsSearchParams object
// with the default values initialized.
func NewGetGroupsSearchParams() *GetGroupsSearchParams {
	var ()
	return &GetGroupsSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupsSearchParamsWithTimeout creates a new GetGroupsSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupsSearchParamsWithTimeout(timeout time.Duration) *GetGroupsSearchParams {
	var ()
	return &GetGroupsSearchParams{

		timeout: timeout,
	}
}

// NewGetGroupsSearchParamsWithContext creates a new GetGroupsSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupsSearchParamsWithContext(ctx context.Context) *GetGroupsSearchParams {
	var ()
	return &GetGroupsSearchParams{

		Context: ctx,
	}
}

// NewGetGroupsSearchParamsWithHTTPClient creates a new GetGroupsSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupsSearchParamsWithHTTPClient(client *http.Client) *GetGroupsSearchParams {
	var ()
	return &GetGroupsSearchParams{
		HTTPClient: client,
	}
}

/*GetGroupsSearchParams contains all the parameters to send to the API endpoint
for the get groups search operation typically these are written to a http.Request
*/
type GetGroupsSearchParams struct {

	/*Expand
	  expand

	*/
	Expand []string
	/*Q64
	  q64

	*/
	Q64 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get groups search params
func (o *GetGroupsSearchParams) WithTimeout(timeout time.Duration) *GetGroupsSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get groups search params
func (o *GetGroupsSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get groups search params
func (o *GetGroupsSearchParams) WithContext(ctx context.Context) *GetGroupsSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get groups search params
func (o *GetGroupsSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get groups search params
func (o *GetGroupsSearchParams) WithHTTPClient(client *http.Client) *GetGroupsSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get groups search params
func (o *GetGroupsSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get groups search params
func (o *GetGroupsSearchParams) WithExpand(expand []string) *GetGroupsSearchParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get groups search params
func (o *GetGroupsSearchParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithQ64 adds the q64 to the get groups search params
func (o *GetGroupsSearchParams) WithQ64(q64 string) *GetGroupsSearchParams {
	o.SetQ64(q64)
	return o
}

// SetQ64 adds the q64 to the get groups search params
func (o *GetGroupsSearchParams) SetQ64(q64 string) {
	o.Q64 = q64
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupsSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	// query param q64
	qrQ64 := o.Q64
	qQ64 := qrQ64
	if qQ64 != "" {
		if err := r.SetQueryParam("q64", qQ64); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}