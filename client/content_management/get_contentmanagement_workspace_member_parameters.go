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
	"github.com/go-openapi/swag"
)

// NewGetContentmanagementWorkspaceMemberParams creates a new GetContentmanagementWorkspaceMemberParams object
// with the default values initialized.
func NewGetContentmanagementWorkspaceMemberParams() *GetContentmanagementWorkspaceMemberParams {
	var ()
	return &GetContentmanagementWorkspaceMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentmanagementWorkspaceMemberParamsWithTimeout creates a new GetContentmanagementWorkspaceMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContentmanagementWorkspaceMemberParamsWithTimeout(timeout time.Duration) *GetContentmanagementWorkspaceMemberParams {
	var ()
	return &GetContentmanagementWorkspaceMemberParams{

		timeout: timeout,
	}
}

// NewGetContentmanagementWorkspaceMemberParamsWithContext creates a new GetContentmanagementWorkspaceMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContentmanagementWorkspaceMemberParamsWithContext(ctx context.Context) *GetContentmanagementWorkspaceMemberParams {
	var ()
	return &GetContentmanagementWorkspaceMemberParams{

		Context: ctx,
	}
}

// NewGetContentmanagementWorkspaceMemberParamsWithHTTPClient creates a new GetContentmanagementWorkspaceMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContentmanagementWorkspaceMemberParamsWithHTTPClient(client *http.Client) *GetContentmanagementWorkspaceMemberParams {
	var ()
	return &GetContentmanagementWorkspaceMemberParams{
		HTTPClient: client,
	}
}

/*GetContentmanagementWorkspaceMemberParams contains all the parameters to send to the API endpoint
for the get contentmanagement workspace member operation typically these are written to a http.Request
*/
type GetContentmanagementWorkspaceMemberParams struct {

	/*Expand
	  Which fields, if any, to expand.

	*/
	Expand []string
	/*MemberID
	  Member ID

	*/
	MemberID string
	/*WorkspaceID
	  Workspace ID

	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get contentmanagement workspace member params
func (o *GetContentmanagementWorkspaceMemberParams) WithTimeout(timeout time.Duration) *GetContentmanagementWorkspaceMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contentmanagement workspace member params
func (o *GetContentmanagementWorkspaceMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contentmanagement workspace member params
func (o *GetContentmanagementWorkspaceMemberParams) WithContext(ctx context.Context) *GetContentmanagementWorkspaceMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contentmanagement workspace member params
func (o *GetContentmanagementWorkspaceMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contentmanagement workspace member params
func (o *GetContentmanagementWorkspaceMemberParams) WithHTTPClient(client *http.Client) *GetContentmanagementWorkspaceMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contentmanagement workspace member params
func (o *GetContentmanagementWorkspaceMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get contentmanagement workspace member params
func (o *GetContentmanagementWorkspaceMemberParams) WithExpand(expand []string) *GetContentmanagementWorkspaceMemberParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get contentmanagement workspace member params
func (o *GetContentmanagementWorkspaceMemberParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithMemberID adds the memberID to the get contentmanagement workspace member params
func (o *GetContentmanagementWorkspaceMemberParams) WithMemberID(memberID string) *GetContentmanagementWorkspaceMemberParams {
	o.SetMemberID(memberID)
	return o
}

// SetMemberID adds the memberId to the get contentmanagement workspace member params
func (o *GetContentmanagementWorkspaceMemberParams) SetMemberID(memberID string) {
	o.MemberID = memberID
}

// WithWorkspaceID adds the workspaceID to the get contentmanagement workspace member params
func (o *GetContentmanagementWorkspaceMemberParams) WithWorkspaceID(workspaceID string) *GetContentmanagementWorkspaceMemberParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get contentmanagement workspace member params
func (o *GetContentmanagementWorkspaceMemberParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentmanagementWorkspaceMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "multi")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
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