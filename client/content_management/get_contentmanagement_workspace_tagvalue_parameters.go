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

// NewGetContentmanagementWorkspaceTagvalueParams creates a new GetContentmanagementWorkspaceTagvalueParams object
// with the default values initialized.
func NewGetContentmanagementWorkspaceTagvalueParams() *GetContentmanagementWorkspaceTagvalueParams {
	var ()
	return &GetContentmanagementWorkspaceTagvalueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentmanagementWorkspaceTagvalueParamsWithTimeout creates a new GetContentmanagementWorkspaceTagvalueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContentmanagementWorkspaceTagvalueParamsWithTimeout(timeout time.Duration) *GetContentmanagementWorkspaceTagvalueParams {
	var ()
	return &GetContentmanagementWorkspaceTagvalueParams{

		timeout: timeout,
	}
}

// NewGetContentmanagementWorkspaceTagvalueParamsWithContext creates a new GetContentmanagementWorkspaceTagvalueParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContentmanagementWorkspaceTagvalueParamsWithContext(ctx context.Context) *GetContentmanagementWorkspaceTagvalueParams {
	var ()
	return &GetContentmanagementWorkspaceTagvalueParams{

		Context: ctx,
	}
}

// NewGetContentmanagementWorkspaceTagvalueParamsWithHTTPClient creates a new GetContentmanagementWorkspaceTagvalueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContentmanagementWorkspaceTagvalueParamsWithHTTPClient(client *http.Client) *GetContentmanagementWorkspaceTagvalueParams {
	var ()
	return &GetContentmanagementWorkspaceTagvalueParams{
		HTTPClient: client,
	}
}

/*GetContentmanagementWorkspaceTagvalueParams contains all the parameters to send to the API endpoint
for the get contentmanagement workspace tagvalue operation typically these are written to a http.Request
*/
type GetContentmanagementWorkspaceTagvalueParams struct {

	/*Expand
	  Which fields, if any, to expand.

	*/
	Expand []string
	/*TagID
	  Tag ID

	*/
	TagID string
	/*WorkspaceID
	  Workspace ID

	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get contentmanagement workspace tagvalue params
func (o *GetContentmanagementWorkspaceTagvalueParams) WithTimeout(timeout time.Duration) *GetContentmanagementWorkspaceTagvalueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contentmanagement workspace tagvalue params
func (o *GetContentmanagementWorkspaceTagvalueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contentmanagement workspace tagvalue params
func (o *GetContentmanagementWorkspaceTagvalueParams) WithContext(ctx context.Context) *GetContentmanagementWorkspaceTagvalueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contentmanagement workspace tagvalue params
func (o *GetContentmanagementWorkspaceTagvalueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contentmanagement workspace tagvalue params
func (o *GetContentmanagementWorkspaceTagvalueParams) WithHTTPClient(client *http.Client) *GetContentmanagementWorkspaceTagvalueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contentmanagement workspace tagvalue params
func (o *GetContentmanagementWorkspaceTagvalueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get contentmanagement workspace tagvalue params
func (o *GetContentmanagementWorkspaceTagvalueParams) WithExpand(expand []string) *GetContentmanagementWorkspaceTagvalueParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get contentmanagement workspace tagvalue params
func (o *GetContentmanagementWorkspaceTagvalueParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithTagID adds the tagID to the get contentmanagement workspace tagvalue params
func (o *GetContentmanagementWorkspaceTagvalueParams) WithTagID(tagID string) *GetContentmanagementWorkspaceTagvalueParams {
	o.SetTagID(tagID)
	return o
}

// SetTagID adds the tagId to the get contentmanagement workspace tagvalue params
func (o *GetContentmanagementWorkspaceTagvalueParams) SetTagID(tagID string) {
	o.TagID = tagID
}

// WithWorkspaceID adds the workspaceID to the get contentmanagement workspace tagvalue params
func (o *GetContentmanagementWorkspaceTagvalueParams) WithWorkspaceID(workspaceID string) *GetContentmanagementWorkspaceTagvalueParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get contentmanagement workspace tagvalue params
func (o *GetContentmanagementWorkspaceTagvalueParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentmanagementWorkspaceTagvalueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tagId
	if err := r.SetPathParam("tagId", o.TagID); err != nil {
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