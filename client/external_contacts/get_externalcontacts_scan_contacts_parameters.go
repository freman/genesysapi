// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

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

// NewGetExternalcontactsScanContactsParams creates a new GetExternalcontactsScanContactsParams object
// with the default values initialized.
func NewGetExternalcontactsScanContactsParams() *GetExternalcontactsScanContactsParams {
	var ()
	return &GetExternalcontactsScanContactsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalcontactsScanContactsParamsWithTimeout creates a new GetExternalcontactsScanContactsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetExternalcontactsScanContactsParamsWithTimeout(timeout time.Duration) *GetExternalcontactsScanContactsParams {
	var ()
	return &GetExternalcontactsScanContactsParams{

		timeout: timeout,
	}
}

// NewGetExternalcontactsScanContactsParamsWithContext creates a new GetExternalcontactsScanContactsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetExternalcontactsScanContactsParamsWithContext(ctx context.Context) *GetExternalcontactsScanContactsParams {
	var ()
	return &GetExternalcontactsScanContactsParams{

		Context: ctx,
	}
}

// NewGetExternalcontactsScanContactsParamsWithHTTPClient creates a new GetExternalcontactsScanContactsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetExternalcontactsScanContactsParamsWithHTTPClient(client *http.Client) *GetExternalcontactsScanContactsParams {
	var ()
	return &GetExternalcontactsScanContactsParams{
		HTTPClient: client,
	}
}

/*GetExternalcontactsScanContactsParams contains all the parameters to send to the API endpoint
for the get externalcontacts scan contacts operation typically these are written to a http.Request
*/
type GetExternalcontactsScanContactsParams struct {

	/*Cursor
	  Indicates where to resume query results (not required for first page), each page returns a new cursor with a 24h TTL

	*/
	Cursor *string
	/*Limit
	  The number of contacts per page; must be between 10 and 200, default is 100)

	*/
	Limit *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get externalcontacts scan contacts params
func (o *GetExternalcontactsScanContactsParams) WithTimeout(timeout time.Duration) *GetExternalcontactsScanContactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get externalcontacts scan contacts params
func (o *GetExternalcontactsScanContactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get externalcontacts scan contacts params
func (o *GetExternalcontactsScanContactsParams) WithContext(ctx context.Context) *GetExternalcontactsScanContactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get externalcontacts scan contacts params
func (o *GetExternalcontactsScanContactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get externalcontacts scan contacts params
func (o *GetExternalcontactsScanContactsParams) WithHTTPClient(client *http.Client) *GetExternalcontactsScanContactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get externalcontacts scan contacts params
func (o *GetExternalcontactsScanContactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the get externalcontacts scan contacts params
func (o *GetExternalcontactsScanContactsParams) WithCursor(cursor *string) *GetExternalcontactsScanContactsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get externalcontacts scan contacts params
func (o *GetExternalcontactsScanContactsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithLimit adds the limit to the get externalcontacts scan contacts params
func (o *GetExternalcontactsScanContactsParams) WithLimit(limit *int32) *GetExternalcontactsScanContactsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get externalcontacts scan contacts params
func (o *GetExternalcontactsScanContactsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalcontactsScanContactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
