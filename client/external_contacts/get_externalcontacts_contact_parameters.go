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

// NewGetExternalcontactsContactParams creates a new GetExternalcontactsContactParams object
// with the default values initialized.
func NewGetExternalcontactsContactParams() *GetExternalcontactsContactParams {
	var ()
	return &GetExternalcontactsContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalcontactsContactParamsWithTimeout creates a new GetExternalcontactsContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetExternalcontactsContactParamsWithTimeout(timeout time.Duration) *GetExternalcontactsContactParams {
	var ()
	return &GetExternalcontactsContactParams{

		timeout: timeout,
	}
}

// NewGetExternalcontactsContactParamsWithContext creates a new GetExternalcontactsContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetExternalcontactsContactParamsWithContext(ctx context.Context) *GetExternalcontactsContactParams {
	var ()
	return &GetExternalcontactsContactParams{

		Context: ctx,
	}
}

// NewGetExternalcontactsContactParamsWithHTTPClient creates a new GetExternalcontactsContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetExternalcontactsContactParamsWithHTTPClient(client *http.Client) *GetExternalcontactsContactParams {
	var ()
	return &GetExternalcontactsContactParams{
		HTTPClient: client,
	}
}

/*GetExternalcontactsContactParams contains all the parameters to send to the API endpoint
for the get externalcontacts contact operation typically these are written to a http.Request
*/
type GetExternalcontactsContactParams struct {

	/*ContactID
	  ExternalContact ID

	*/
	ContactID string
	/*Expand
	  which fields, if any, to expand (externalOrganization,externalDataSources)

	*/
	Expand []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get externalcontacts contact params
func (o *GetExternalcontactsContactParams) WithTimeout(timeout time.Duration) *GetExternalcontactsContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get externalcontacts contact params
func (o *GetExternalcontactsContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get externalcontacts contact params
func (o *GetExternalcontactsContactParams) WithContext(ctx context.Context) *GetExternalcontactsContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get externalcontacts contact params
func (o *GetExternalcontactsContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get externalcontacts contact params
func (o *GetExternalcontactsContactParams) WithHTTPClient(client *http.Client) *GetExternalcontactsContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get externalcontacts contact params
func (o *GetExternalcontactsContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the get externalcontacts contact params
func (o *GetExternalcontactsContactParams) WithContactID(contactID string) *GetExternalcontactsContactParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the get externalcontacts contact params
func (o *GetExternalcontactsContactParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithExpand adds the expand to the get externalcontacts contact params
func (o *GetExternalcontactsContactParams) WithExpand(expand []string) *GetExternalcontactsContactParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get externalcontacts contact params
func (o *GetExternalcontactsContactParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalcontactsContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contactId
	if err := r.SetPathParam("contactId", o.ContactID); err != nil {
		return err
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
