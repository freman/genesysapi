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

	"github.com/freman/genesysapi/models"
)

// NewPostExternalcontactsOrganizationNotesParams creates a new PostExternalcontactsOrganizationNotesParams object
// with the default values initialized.
func NewPostExternalcontactsOrganizationNotesParams() *PostExternalcontactsOrganizationNotesParams {
	var ()
	return &PostExternalcontactsOrganizationNotesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostExternalcontactsOrganizationNotesParamsWithTimeout creates a new PostExternalcontactsOrganizationNotesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostExternalcontactsOrganizationNotesParamsWithTimeout(timeout time.Duration) *PostExternalcontactsOrganizationNotesParams {
	var ()
	return &PostExternalcontactsOrganizationNotesParams{

		timeout: timeout,
	}
}

// NewPostExternalcontactsOrganizationNotesParamsWithContext creates a new PostExternalcontactsOrganizationNotesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostExternalcontactsOrganizationNotesParamsWithContext(ctx context.Context) *PostExternalcontactsOrganizationNotesParams {
	var ()
	return &PostExternalcontactsOrganizationNotesParams{

		Context: ctx,
	}
}

// NewPostExternalcontactsOrganizationNotesParamsWithHTTPClient creates a new PostExternalcontactsOrganizationNotesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostExternalcontactsOrganizationNotesParamsWithHTTPClient(client *http.Client) *PostExternalcontactsOrganizationNotesParams {
	var ()
	return &PostExternalcontactsOrganizationNotesParams{
		HTTPClient: client,
	}
}

/*PostExternalcontactsOrganizationNotesParams contains all the parameters to send to the API endpoint
for the post externalcontacts organization notes operation typically these are written to a http.Request
*/
type PostExternalcontactsOrganizationNotesParams struct {

	/*Body
	  ExternalContact

	*/
	Body *models.Note
	/*ExternalOrganizationID
	  External Organization Id

	*/
	ExternalOrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post externalcontacts organization notes params
func (o *PostExternalcontactsOrganizationNotesParams) WithTimeout(timeout time.Duration) *PostExternalcontactsOrganizationNotesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post externalcontacts organization notes params
func (o *PostExternalcontactsOrganizationNotesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post externalcontacts organization notes params
func (o *PostExternalcontactsOrganizationNotesParams) WithContext(ctx context.Context) *PostExternalcontactsOrganizationNotesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post externalcontacts organization notes params
func (o *PostExternalcontactsOrganizationNotesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post externalcontacts organization notes params
func (o *PostExternalcontactsOrganizationNotesParams) WithHTTPClient(client *http.Client) *PostExternalcontactsOrganizationNotesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post externalcontacts organization notes params
func (o *PostExternalcontactsOrganizationNotesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post externalcontacts organization notes params
func (o *PostExternalcontactsOrganizationNotesParams) WithBody(body *models.Note) *PostExternalcontactsOrganizationNotesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post externalcontacts organization notes params
func (o *PostExternalcontactsOrganizationNotesParams) SetBody(body *models.Note) {
	o.Body = body
}

// WithExternalOrganizationID adds the externalOrganizationID to the post externalcontacts organization notes params
func (o *PostExternalcontactsOrganizationNotesParams) WithExternalOrganizationID(externalOrganizationID string) *PostExternalcontactsOrganizationNotesParams {
	o.SetExternalOrganizationID(externalOrganizationID)
	return o
}

// SetExternalOrganizationID adds the externalOrganizationId to the post externalcontacts organization notes params
func (o *PostExternalcontactsOrganizationNotesParams) SetExternalOrganizationID(externalOrganizationID string) {
	o.ExternalOrganizationID = externalOrganizationID
}

// WriteToRequest writes these params to a swagger request
func (o *PostExternalcontactsOrganizationNotesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param externalOrganizationId
	if err := r.SetPathParam("externalOrganizationId", o.ExternalOrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
