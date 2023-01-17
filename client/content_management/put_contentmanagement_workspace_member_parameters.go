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

	"github.com/freman/genesysapi/models"
)

// NewPutContentmanagementWorkspaceMemberParams creates a new PutContentmanagementWorkspaceMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutContentmanagementWorkspaceMemberParams() *PutContentmanagementWorkspaceMemberParams {
	return &PutContentmanagementWorkspaceMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutContentmanagementWorkspaceMemberParamsWithTimeout creates a new PutContentmanagementWorkspaceMemberParams object
// with the ability to set a timeout on a request.
func NewPutContentmanagementWorkspaceMemberParamsWithTimeout(timeout time.Duration) *PutContentmanagementWorkspaceMemberParams {
	return &PutContentmanagementWorkspaceMemberParams{
		timeout: timeout,
	}
}

// NewPutContentmanagementWorkspaceMemberParamsWithContext creates a new PutContentmanagementWorkspaceMemberParams object
// with the ability to set a context for a request.
func NewPutContentmanagementWorkspaceMemberParamsWithContext(ctx context.Context) *PutContentmanagementWorkspaceMemberParams {
	return &PutContentmanagementWorkspaceMemberParams{
		Context: ctx,
	}
}

// NewPutContentmanagementWorkspaceMemberParamsWithHTTPClient creates a new PutContentmanagementWorkspaceMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutContentmanagementWorkspaceMemberParamsWithHTTPClient(client *http.Client) *PutContentmanagementWorkspaceMemberParams {
	return &PutContentmanagementWorkspaceMemberParams{
		HTTPClient: client,
	}
}

/*
PutContentmanagementWorkspaceMemberParams contains all the parameters to send to the API endpoint

	for the put contentmanagement workspace member operation.

	Typically these are written to a http.Request.
*/
type PutContentmanagementWorkspaceMemberParams struct {

	/* Body.

	   Workspace Member
	*/
	Body *models.WorkspaceMember

	/* MemberID.

	   Member ID
	*/
	MemberID string

	/* WorkspaceID.

	   Workspace ID
	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put contentmanagement workspace member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutContentmanagementWorkspaceMemberParams) WithDefaults() *PutContentmanagementWorkspaceMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put contentmanagement workspace member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutContentmanagementWorkspaceMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put contentmanagement workspace member params
func (o *PutContentmanagementWorkspaceMemberParams) WithTimeout(timeout time.Duration) *PutContentmanagementWorkspaceMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put contentmanagement workspace member params
func (o *PutContentmanagementWorkspaceMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put contentmanagement workspace member params
func (o *PutContentmanagementWorkspaceMemberParams) WithContext(ctx context.Context) *PutContentmanagementWorkspaceMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put contentmanagement workspace member params
func (o *PutContentmanagementWorkspaceMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put contentmanagement workspace member params
func (o *PutContentmanagementWorkspaceMemberParams) WithHTTPClient(client *http.Client) *PutContentmanagementWorkspaceMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put contentmanagement workspace member params
func (o *PutContentmanagementWorkspaceMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put contentmanagement workspace member params
func (o *PutContentmanagementWorkspaceMemberParams) WithBody(body *models.WorkspaceMember) *PutContentmanagementWorkspaceMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put contentmanagement workspace member params
func (o *PutContentmanagementWorkspaceMemberParams) SetBody(body *models.WorkspaceMember) {
	o.Body = body
}

// WithMemberID adds the memberID to the put contentmanagement workspace member params
func (o *PutContentmanagementWorkspaceMemberParams) WithMemberID(memberID string) *PutContentmanagementWorkspaceMemberParams {
	o.SetMemberID(memberID)
	return o
}

// SetMemberID adds the memberId to the put contentmanagement workspace member params
func (o *PutContentmanagementWorkspaceMemberParams) SetMemberID(memberID string) {
	o.MemberID = memberID
}

// WithWorkspaceID adds the workspaceID to the put contentmanagement workspace member params
func (o *PutContentmanagementWorkspaceMemberParams) WithWorkspaceID(workspaceID string) *PutContentmanagementWorkspaceMemberParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the put contentmanagement workspace member params
func (o *PutContentmanagementWorkspaceMemberParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PutContentmanagementWorkspaceMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param memberId
	if err := r.SetPathParam("memberId", o.MemberID); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", o.WorkspaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
