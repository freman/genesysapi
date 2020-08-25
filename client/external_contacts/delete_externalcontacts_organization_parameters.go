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
)

// NewDeleteExternalcontactsOrganizationParams creates a new DeleteExternalcontactsOrganizationParams object
// with the default values initialized.
func NewDeleteExternalcontactsOrganizationParams() *DeleteExternalcontactsOrganizationParams {
	var ()
	return &DeleteExternalcontactsOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteExternalcontactsOrganizationParamsWithTimeout creates a new DeleteExternalcontactsOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteExternalcontactsOrganizationParamsWithTimeout(timeout time.Duration) *DeleteExternalcontactsOrganizationParams {
	var ()
	return &DeleteExternalcontactsOrganizationParams{

		timeout: timeout,
	}
}

// NewDeleteExternalcontactsOrganizationParamsWithContext creates a new DeleteExternalcontactsOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteExternalcontactsOrganizationParamsWithContext(ctx context.Context) *DeleteExternalcontactsOrganizationParams {
	var ()
	return &DeleteExternalcontactsOrganizationParams{

		Context: ctx,
	}
}

// NewDeleteExternalcontactsOrganizationParamsWithHTTPClient creates a new DeleteExternalcontactsOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteExternalcontactsOrganizationParamsWithHTTPClient(client *http.Client) *DeleteExternalcontactsOrganizationParams {
	var ()
	return &DeleteExternalcontactsOrganizationParams{
		HTTPClient: client,
	}
}

/*DeleteExternalcontactsOrganizationParams contains all the parameters to send to the API endpoint
for the delete externalcontacts organization operation typically these are written to a http.Request
*/
type DeleteExternalcontactsOrganizationParams struct {

	/*ExternalOrganizationID
	  External Organization ID

	*/
	ExternalOrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete externalcontacts organization params
func (o *DeleteExternalcontactsOrganizationParams) WithTimeout(timeout time.Duration) *DeleteExternalcontactsOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete externalcontacts organization params
func (o *DeleteExternalcontactsOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete externalcontacts organization params
func (o *DeleteExternalcontactsOrganizationParams) WithContext(ctx context.Context) *DeleteExternalcontactsOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete externalcontacts organization params
func (o *DeleteExternalcontactsOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete externalcontacts organization params
func (o *DeleteExternalcontactsOrganizationParams) WithHTTPClient(client *http.Client) *DeleteExternalcontactsOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete externalcontacts organization params
func (o *DeleteExternalcontactsOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExternalOrganizationID adds the externalOrganizationID to the delete externalcontacts organization params
func (o *DeleteExternalcontactsOrganizationParams) WithExternalOrganizationID(externalOrganizationID string) *DeleteExternalcontactsOrganizationParams {
	o.SetExternalOrganizationID(externalOrganizationID)
	return o
}

// SetExternalOrganizationID adds the externalOrganizationId to the delete externalcontacts organization params
func (o *DeleteExternalcontactsOrganizationParams) SetExternalOrganizationID(externalOrganizationID string) {
	o.ExternalOrganizationID = externalOrganizationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteExternalcontactsOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param externalOrganizationId
	if err := r.SetPathParam("externalOrganizationId", o.ExternalOrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}