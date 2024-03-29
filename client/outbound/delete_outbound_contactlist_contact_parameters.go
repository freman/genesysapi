// Code generated by go-swagger; DO NOT EDIT.

package outbound

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

// NewDeleteOutboundContactlistContactParams creates a new DeleteOutboundContactlistContactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOutboundContactlistContactParams() *DeleteOutboundContactlistContactParams {
	return &DeleteOutboundContactlistContactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOutboundContactlistContactParamsWithTimeout creates a new DeleteOutboundContactlistContactParams object
// with the ability to set a timeout on a request.
func NewDeleteOutboundContactlistContactParamsWithTimeout(timeout time.Duration) *DeleteOutboundContactlistContactParams {
	return &DeleteOutboundContactlistContactParams{
		timeout: timeout,
	}
}

// NewDeleteOutboundContactlistContactParamsWithContext creates a new DeleteOutboundContactlistContactParams object
// with the ability to set a context for a request.
func NewDeleteOutboundContactlistContactParamsWithContext(ctx context.Context) *DeleteOutboundContactlistContactParams {
	return &DeleteOutboundContactlistContactParams{
		Context: ctx,
	}
}

// NewDeleteOutboundContactlistContactParamsWithHTTPClient creates a new DeleteOutboundContactlistContactParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOutboundContactlistContactParamsWithHTTPClient(client *http.Client) *DeleteOutboundContactlistContactParams {
	return &DeleteOutboundContactlistContactParams{
		HTTPClient: client,
	}
}

/*
DeleteOutboundContactlistContactParams contains all the parameters to send to the API endpoint

	for the delete outbound contactlist contact operation.

	Typically these are written to a http.Request.
*/
type DeleteOutboundContactlistContactParams struct {

	/* ContactID.

	   Contact ID
	*/
	ContactID string

	/* ContactListID.

	   Contact List ID
	*/
	ContactListID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete outbound contactlist contact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOutboundContactlistContactParams) WithDefaults() *DeleteOutboundContactlistContactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete outbound contactlist contact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOutboundContactlistContactParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete outbound contactlist contact params
func (o *DeleteOutboundContactlistContactParams) WithTimeout(timeout time.Duration) *DeleteOutboundContactlistContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete outbound contactlist contact params
func (o *DeleteOutboundContactlistContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete outbound contactlist contact params
func (o *DeleteOutboundContactlistContactParams) WithContext(ctx context.Context) *DeleteOutboundContactlistContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete outbound contactlist contact params
func (o *DeleteOutboundContactlistContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete outbound contactlist contact params
func (o *DeleteOutboundContactlistContactParams) WithHTTPClient(client *http.Client) *DeleteOutboundContactlistContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete outbound contactlist contact params
func (o *DeleteOutboundContactlistContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the delete outbound contactlist contact params
func (o *DeleteOutboundContactlistContactParams) WithContactID(contactID string) *DeleteOutboundContactlistContactParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the delete outbound contactlist contact params
func (o *DeleteOutboundContactlistContactParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithContactListID adds the contactListID to the delete outbound contactlist contact params
func (o *DeleteOutboundContactlistContactParams) WithContactListID(contactListID string) *DeleteOutboundContactlistContactParams {
	o.SetContactListID(contactListID)
	return o
}

// SetContactListID adds the contactListId to the delete outbound contactlist contact params
func (o *DeleteOutboundContactlistContactParams) SetContactListID(contactListID string) {
	o.ContactListID = contactListID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOutboundContactlistContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contactId
	if err := r.SetPathParam("contactId", o.ContactID); err != nil {
		return err
	}

	// path param contactListId
	if err := r.SetPathParam("contactListId", o.ContactListID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
