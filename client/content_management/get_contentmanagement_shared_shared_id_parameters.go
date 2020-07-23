// Code generated by go-swagger; DO NOT EDIT.

package content_management

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

// NewGetContentmanagementSharedSharedIDParams creates a new GetContentmanagementSharedSharedIDParams object
// with the default values initialized.
func NewGetContentmanagementSharedSharedIDParams() *GetContentmanagementSharedSharedIDParams {
	var (
		dispositionDefault = string("attachment")
		redirectDefault    = bool(true)
	)
	return &GetContentmanagementSharedSharedIDParams{
		Disposition: &dispositionDefault,
		Redirect:    &redirectDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentmanagementSharedSharedIDParamsWithTimeout creates a new GetContentmanagementSharedSharedIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContentmanagementSharedSharedIDParamsWithTimeout(timeout time.Duration) *GetContentmanagementSharedSharedIDParams {
	var (
		dispositionDefault = string("attachment")
		redirectDefault    = bool(true)
	)
	return &GetContentmanagementSharedSharedIDParams{
		Disposition: &dispositionDefault,
		Redirect:    &redirectDefault,

		timeout: timeout,
	}
}

// NewGetContentmanagementSharedSharedIDParamsWithContext creates a new GetContentmanagementSharedSharedIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContentmanagementSharedSharedIDParamsWithContext(ctx context.Context) *GetContentmanagementSharedSharedIDParams {
	var (
		dispositionDefault = string("attachment")
		redirectDefault    = bool(true)
	)
	return &GetContentmanagementSharedSharedIDParams{
		Disposition: &dispositionDefault,
		Redirect:    &redirectDefault,

		Context: ctx,
	}
}

// NewGetContentmanagementSharedSharedIDParamsWithHTTPClient creates a new GetContentmanagementSharedSharedIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContentmanagementSharedSharedIDParamsWithHTTPClient(client *http.Client) *GetContentmanagementSharedSharedIDParams {
	var (
		dispositionDefault = string("attachment")
		redirectDefault    = bool(true)
	)
	return &GetContentmanagementSharedSharedIDParams{
		Disposition: &dispositionDefault,
		Redirect:    &redirectDefault,
		HTTPClient:  client,
	}
}

/*GetContentmanagementSharedSharedIDParams contains all the parameters to send to the API endpoint
for the get contentmanagement shared shared Id operation typically these are written to a http.Request
*/
type GetContentmanagementSharedSharedIDParams struct {

	/*ContentType
	  The requested format for the specified document. If supported, the document will be returned in that format. Example contentType=audio/wav

	*/
	ContentType *string
	/*Disposition
	  Request how the share content will be downloaded: attached as a file or inline. Default is attachment.

	*/
	Disposition *string
	/*Expand
	  Expand some document fields

	*/
	Expand *string
	/*Redirect
	  Turn on or off redirect

	*/
	Redirect *bool
	/*SharedID
	  Shared ID

	*/
	SharedID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) WithTimeout(timeout time.Duration) *GetContentmanagementSharedSharedIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) WithContext(ctx context.Context) *GetContentmanagementSharedSharedIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) WithHTTPClient(client *http.Client) *GetContentmanagementSharedSharedIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) WithContentType(contentType *string) *GetContentmanagementSharedSharedIDParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithDisposition adds the disposition to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) WithDisposition(disposition *string) *GetContentmanagementSharedSharedIDParams {
	o.SetDisposition(disposition)
	return o
}

// SetDisposition adds the disposition to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) SetDisposition(disposition *string) {
	o.Disposition = disposition
}

// WithExpand adds the expand to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) WithExpand(expand *string) *GetContentmanagementSharedSharedIDParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithRedirect adds the redirect to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) WithRedirect(redirect *bool) *GetContentmanagementSharedSharedIDParams {
	o.SetRedirect(redirect)
	return o
}

// SetRedirect adds the redirect to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) SetRedirect(redirect *bool) {
	o.Redirect = redirect
}

// WithSharedID adds the sharedID to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) WithSharedID(sharedID string) *GetContentmanagementSharedSharedIDParams {
	o.SetSharedID(sharedID)
	return o
}

// SetSharedID adds the sharedId to the get contentmanagement shared shared Id params
func (o *GetContentmanagementSharedSharedIDParams) SetSharedID(sharedID string) {
	o.SharedID = sharedID
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentmanagementSharedSharedIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// query param contentType
		var qrContentType string
		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {
			if err := r.SetQueryParam("contentType", qContentType); err != nil {
				return err
			}
		}

	}

	if o.Disposition != nil {

		// query param disposition
		var qrDisposition string
		if o.Disposition != nil {
			qrDisposition = *o.Disposition
		}
		qDisposition := qrDisposition
		if qDisposition != "" {
			if err := r.SetQueryParam("disposition", qDisposition); err != nil {
				return err
			}
		}

	}

	if o.Expand != nil {

		// query param expand
		var qrExpand string
		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {
			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}

	}

	if o.Redirect != nil {

		// query param redirect
		var qrRedirect bool
		if o.Redirect != nil {
			qrRedirect = *o.Redirect
		}
		qRedirect := swag.FormatBool(qrRedirect)
		if qRedirect != "" {
			if err := r.SetQueryParam("redirect", qRedirect); err != nil {
				return err
			}
		}

	}

	// path param sharedId
	if err := r.SetPathParam("sharedId", o.SharedID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
