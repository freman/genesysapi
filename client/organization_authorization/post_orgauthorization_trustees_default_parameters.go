// Code generated by go-swagger; DO NOT EDIT.

package organization_authorization

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

// NewPostOrgauthorizationTrusteesDefaultParams creates a new PostOrgauthorizationTrusteesDefaultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOrgauthorizationTrusteesDefaultParams() *PostOrgauthorizationTrusteesDefaultParams {
	return &PostOrgauthorizationTrusteesDefaultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOrgauthorizationTrusteesDefaultParamsWithTimeout creates a new PostOrgauthorizationTrusteesDefaultParams object
// with the ability to set a timeout on a request.
func NewPostOrgauthorizationTrusteesDefaultParamsWithTimeout(timeout time.Duration) *PostOrgauthorizationTrusteesDefaultParams {
	return &PostOrgauthorizationTrusteesDefaultParams{
		timeout: timeout,
	}
}

// NewPostOrgauthorizationTrusteesDefaultParamsWithContext creates a new PostOrgauthorizationTrusteesDefaultParams object
// with the ability to set a context for a request.
func NewPostOrgauthorizationTrusteesDefaultParamsWithContext(ctx context.Context) *PostOrgauthorizationTrusteesDefaultParams {
	return &PostOrgauthorizationTrusteesDefaultParams{
		Context: ctx,
	}
}

// NewPostOrgauthorizationTrusteesDefaultParamsWithHTTPClient creates a new PostOrgauthorizationTrusteesDefaultParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOrgauthorizationTrusteesDefaultParamsWithHTTPClient(client *http.Client) *PostOrgauthorizationTrusteesDefaultParams {
	return &PostOrgauthorizationTrusteesDefaultParams{
		HTTPClient: client,
	}
}

/*
PostOrgauthorizationTrusteesDefaultParams contains all the parameters to send to the API endpoint

	for the post orgauthorization trustees default operation.

	Typically these are written to a http.Request.
*/
type PostOrgauthorizationTrusteesDefaultParams struct {

	/* AssignDefaultRole.

	   Assign Admin role to default pairing with Customer Care
	*/
	AssignDefaultRole *bool

	/* AutoExpire.

	   Automatically expire pairing after 30 days
	*/
	AutoExpire *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post orgauthorization trustees default params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOrgauthorizationTrusteesDefaultParams) WithDefaults() *PostOrgauthorizationTrusteesDefaultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post orgauthorization trustees default params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOrgauthorizationTrusteesDefaultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post orgauthorization trustees default params
func (o *PostOrgauthorizationTrusteesDefaultParams) WithTimeout(timeout time.Duration) *PostOrgauthorizationTrusteesDefaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post orgauthorization trustees default params
func (o *PostOrgauthorizationTrusteesDefaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post orgauthorization trustees default params
func (o *PostOrgauthorizationTrusteesDefaultParams) WithContext(ctx context.Context) *PostOrgauthorizationTrusteesDefaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post orgauthorization trustees default params
func (o *PostOrgauthorizationTrusteesDefaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post orgauthorization trustees default params
func (o *PostOrgauthorizationTrusteesDefaultParams) WithHTTPClient(client *http.Client) *PostOrgauthorizationTrusteesDefaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post orgauthorization trustees default params
func (o *PostOrgauthorizationTrusteesDefaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssignDefaultRole adds the assignDefaultRole to the post orgauthorization trustees default params
func (o *PostOrgauthorizationTrusteesDefaultParams) WithAssignDefaultRole(assignDefaultRole *bool) *PostOrgauthorizationTrusteesDefaultParams {
	o.SetAssignDefaultRole(assignDefaultRole)
	return o
}

// SetAssignDefaultRole adds the assignDefaultRole to the post orgauthorization trustees default params
func (o *PostOrgauthorizationTrusteesDefaultParams) SetAssignDefaultRole(assignDefaultRole *bool) {
	o.AssignDefaultRole = assignDefaultRole
}

// WithAutoExpire adds the autoExpire to the post orgauthorization trustees default params
func (o *PostOrgauthorizationTrusteesDefaultParams) WithAutoExpire(autoExpire *bool) *PostOrgauthorizationTrusteesDefaultParams {
	o.SetAutoExpire(autoExpire)
	return o
}

// SetAutoExpire adds the autoExpire to the post orgauthorization trustees default params
func (o *PostOrgauthorizationTrusteesDefaultParams) SetAutoExpire(autoExpire *bool) {
	o.AutoExpire = autoExpire
}

// WriteToRequest writes these params to a swagger request
func (o *PostOrgauthorizationTrusteesDefaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AssignDefaultRole != nil {

		// query param assignDefaultRole
		var qrAssignDefaultRole bool

		if o.AssignDefaultRole != nil {
			qrAssignDefaultRole = *o.AssignDefaultRole
		}
		qAssignDefaultRole := swag.FormatBool(qrAssignDefaultRole)
		if qAssignDefaultRole != "" {

			if err := r.SetQueryParam("assignDefaultRole", qAssignDefaultRole); err != nil {
				return err
			}
		}
	}

	if o.AutoExpire != nil {

		// query param autoExpire
		var qrAutoExpire bool

		if o.AutoExpire != nil {
			qrAutoExpire = *o.AutoExpire
		}
		qAutoExpire := swag.FormatBool(qrAutoExpire)
		if qAutoExpire != "" {

			if err := r.SetQueryParam("autoExpire", qAutoExpire); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
