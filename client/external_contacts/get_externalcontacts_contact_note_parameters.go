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

// NewGetExternalcontactsContactNoteParams creates a new GetExternalcontactsContactNoteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExternalcontactsContactNoteParams() *GetExternalcontactsContactNoteParams {
	return &GetExternalcontactsContactNoteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalcontactsContactNoteParamsWithTimeout creates a new GetExternalcontactsContactNoteParams object
// with the ability to set a timeout on a request.
func NewGetExternalcontactsContactNoteParamsWithTimeout(timeout time.Duration) *GetExternalcontactsContactNoteParams {
	return &GetExternalcontactsContactNoteParams{
		timeout: timeout,
	}
}

// NewGetExternalcontactsContactNoteParamsWithContext creates a new GetExternalcontactsContactNoteParams object
// with the ability to set a context for a request.
func NewGetExternalcontactsContactNoteParamsWithContext(ctx context.Context) *GetExternalcontactsContactNoteParams {
	return &GetExternalcontactsContactNoteParams{
		Context: ctx,
	}
}

// NewGetExternalcontactsContactNoteParamsWithHTTPClient creates a new GetExternalcontactsContactNoteParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExternalcontactsContactNoteParamsWithHTTPClient(client *http.Client) *GetExternalcontactsContactNoteParams {
	return &GetExternalcontactsContactNoteParams{
		HTTPClient: client,
	}
}

/*
GetExternalcontactsContactNoteParams contains all the parameters to send to the API endpoint

	for the get externalcontacts contact note operation.

	Typically these are written to a http.Request.
*/
type GetExternalcontactsContactNoteParams struct {

	/* ContactID.

	   ExternalContact Id
	*/
	ContactID string

	/* Expand.

	   which fields, if any, to expand
	*/
	Expand []string

	/* NoteID.

	   Note Id
	*/
	NoteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get externalcontacts contact note params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalcontactsContactNoteParams) WithDefaults() *GetExternalcontactsContactNoteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get externalcontacts contact note params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalcontactsContactNoteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get externalcontacts contact note params
func (o *GetExternalcontactsContactNoteParams) WithTimeout(timeout time.Duration) *GetExternalcontactsContactNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get externalcontacts contact note params
func (o *GetExternalcontactsContactNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get externalcontacts contact note params
func (o *GetExternalcontactsContactNoteParams) WithContext(ctx context.Context) *GetExternalcontactsContactNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get externalcontacts contact note params
func (o *GetExternalcontactsContactNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get externalcontacts contact note params
func (o *GetExternalcontactsContactNoteParams) WithHTTPClient(client *http.Client) *GetExternalcontactsContactNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get externalcontacts contact note params
func (o *GetExternalcontactsContactNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactID adds the contactID to the get externalcontacts contact note params
func (o *GetExternalcontactsContactNoteParams) WithContactID(contactID string) *GetExternalcontactsContactNoteParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the get externalcontacts contact note params
func (o *GetExternalcontactsContactNoteParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithExpand adds the expand to the get externalcontacts contact note params
func (o *GetExternalcontactsContactNoteParams) WithExpand(expand []string) *GetExternalcontactsContactNoteParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get externalcontacts contact note params
func (o *GetExternalcontactsContactNoteParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithNoteID adds the noteID to the get externalcontacts contact note params
func (o *GetExternalcontactsContactNoteParams) WithNoteID(noteID string) *GetExternalcontactsContactNoteParams {
	o.SetNoteID(noteID)
	return o
}

// SetNoteID adds the noteId to the get externalcontacts contact note params
func (o *GetExternalcontactsContactNoteParams) SetNoteID(noteID string) {
	o.NoteID = noteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalcontactsContactNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contactId
	if err := r.SetPathParam("contactId", o.ContactID); err != nil {
		return err
	}

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	// path param noteId
	if err := r.SetPathParam("noteId", o.NoteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetExternalcontactsContactNote binds the parameter expand
func (o *GetExternalcontactsContactNoteParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}
