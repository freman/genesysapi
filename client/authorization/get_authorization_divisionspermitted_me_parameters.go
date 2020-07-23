// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// NewGetAuthorizationDivisionspermittedMeParams creates a new GetAuthorizationDivisionspermittedMeParams object
// with the default values initialized.
func NewGetAuthorizationDivisionspermittedMeParams() *GetAuthorizationDivisionspermittedMeParams {
	var ()
	return &GetAuthorizationDivisionspermittedMeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationDivisionspermittedMeParamsWithTimeout creates a new GetAuthorizationDivisionspermittedMeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthorizationDivisionspermittedMeParamsWithTimeout(timeout time.Duration) *GetAuthorizationDivisionspermittedMeParams {
	var ()
	return &GetAuthorizationDivisionspermittedMeParams{

		timeout: timeout,
	}
}

// NewGetAuthorizationDivisionspermittedMeParamsWithContext creates a new GetAuthorizationDivisionspermittedMeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthorizationDivisionspermittedMeParamsWithContext(ctx context.Context) *GetAuthorizationDivisionspermittedMeParams {
	var ()
	return &GetAuthorizationDivisionspermittedMeParams{

		Context: ctx,
	}
}

// NewGetAuthorizationDivisionspermittedMeParamsWithHTTPClient creates a new GetAuthorizationDivisionspermittedMeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthorizationDivisionspermittedMeParamsWithHTTPClient(client *http.Client) *GetAuthorizationDivisionspermittedMeParams {
	var ()
	return &GetAuthorizationDivisionspermittedMeParams{
		HTTPClient: client,
	}
}

/*GetAuthorizationDivisionspermittedMeParams contains all the parameters to send to the API endpoint
for the get authorization divisionspermitted me operation typically these are written to a http.Request
*/
type GetAuthorizationDivisionspermittedMeParams struct {

	/*Name
	  Search term to filter by division name

	*/
	Name *string
	/*Permission
	  The permission string, including the object to access, e.g. routing:queue:view

	*/
	Permission string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get authorization divisionspermitted me params
func (o *GetAuthorizationDivisionspermittedMeParams) WithTimeout(timeout time.Duration) *GetAuthorizationDivisionspermittedMeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization divisionspermitted me params
func (o *GetAuthorizationDivisionspermittedMeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization divisionspermitted me params
func (o *GetAuthorizationDivisionspermittedMeParams) WithContext(ctx context.Context) *GetAuthorizationDivisionspermittedMeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization divisionspermitted me params
func (o *GetAuthorizationDivisionspermittedMeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization divisionspermitted me params
func (o *GetAuthorizationDivisionspermittedMeParams) WithHTTPClient(client *http.Client) *GetAuthorizationDivisionspermittedMeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization divisionspermitted me params
func (o *GetAuthorizationDivisionspermittedMeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get authorization divisionspermitted me params
func (o *GetAuthorizationDivisionspermittedMeParams) WithName(name *string) *GetAuthorizationDivisionspermittedMeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get authorization divisionspermitted me params
func (o *GetAuthorizationDivisionspermittedMeParams) SetName(name *string) {
	o.Name = name
}

// WithPermission adds the permission to the get authorization divisionspermitted me params
func (o *GetAuthorizationDivisionspermittedMeParams) WithPermission(permission string) *GetAuthorizationDivisionspermittedMeParams {
	o.SetPermission(permission)
	return o
}

// SetPermission adds the permission to the get authorization divisionspermitted me params
func (o *GetAuthorizationDivisionspermittedMeParams) SetPermission(permission string) {
	o.Permission = permission
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationDivisionspermittedMeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	// query param permission
	qrPermission := o.Permission
	qPermission := qrPermission
	if qPermission != "" {
		if err := r.SetQueryParam("permission", qPermission); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
