// Code generated by go-swagger; DO NOT EDIT.

package license

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

// NewPostLicenseOrganizationParams creates a new PostLicenseOrganizationParams object
// with the default values initialized.
func NewPostLicenseOrganizationParams() *PostLicenseOrganizationParams {
	var ()
	return &PostLicenseOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLicenseOrganizationParamsWithTimeout creates a new PostLicenseOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLicenseOrganizationParamsWithTimeout(timeout time.Duration) *PostLicenseOrganizationParams {
	var ()
	return &PostLicenseOrganizationParams{

		timeout: timeout,
	}
}

// NewPostLicenseOrganizationParamsWithContext creates a new PostLicenseOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLicenseOrganizationParamsWithContext(ctx context.Context) *PostLicenseOrganizationParams {
	var ()
	return &PostLicenseOrganizationParams{

		Context: ctx,
	}
}

// NewPostLicenseOrganizationParamsWithHTTPClient creates a new PostLicenseOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLicenseOrganizationParamsWithHTTPClient(client *http.Client) *PostLicenseOrganizationParams {
	var ()
	return &PostLicenseOrganizationParams{
		HTTPClient: client,
	}
}

/*PostLicenseOrganizationParams contains all the parameters to send to the API endpoint
for the post license organization operation typically these are written to a http.Request
*/
type PostLicenseOrganizationParams struct {

	/*Body
	  The license assignments to update.

	*/
	Body *models.LicenseBatchAssignmentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post license organization params
func (o *PostLicenseOrganizationParams) WithTimeout(timeout time.Duration) *PostLicenseOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post license organization params
func (o *PostLicenseOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post license organization params
func (o *PostLicenseOrganizationParams) WithContext(ctx context.Context) *PostLicenseOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post license organization params
func (o *PostLicenseOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post license organization params
func (o *PostLicenseOrganizationParams) WithHTTPClient(client *http.Client) *PostLicenseOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post license organization params
func (o *PostLicenseOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post license organization params
func (o *PostLicenseOrganizationParams) WithBody(body *models.LicenseBatchAssignmentRequest) *PostLicenseOrganizationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post license organization params
func (o *PostLicenseOrganizationParams) SetBody(body *models.LicenseBatchAssignmentRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostLicenseOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}