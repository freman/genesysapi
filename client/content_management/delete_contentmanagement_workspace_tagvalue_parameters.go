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
)

// NewDeleteContentmanagementWorkspaceTagvalueParams creates a new DeleteContentmanagementWorkspaceTagvalueParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteContentmanagementWorkspaceTagvalueParams() *DeleteContentmanagementWorkspaceTagvalueParams {
	return &DeleteContentmanagementWorkspaceTagvalueParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContentmanagementWorkspaceTagvalueParamsWithTimeout creates a new DeleteContentmanagementWorkspaceTagvalueParams object
// with the ability to set a timeout on a request.
func NewDeleteContentmanagementWorkspaceTagvalueParamsWithTimeout(timeout time.Duration) *DeleteContentmanagementWorkspaceTagvalueParams {
	return &DeleteContentmanagementWorkspaceTagvalueParams{
		timeout: timeout,
	}
}

// NewDeleteContentmanagementWorkspaceTagvalueParamsWithContext creates a new DeleteContentmanagementWorkspaceTagvalueParams object
// with the ability to set a context for a request.
func NewDeleteContentmanagementWorkspaceTagvalueParamsWithContext(ctx context.Context) *DeleteContentmanagementWorkspaceTagvalueParams {
	return &DeleteContentmanagementWorkspaceTagvalueParams{
		Context: ctx,
	}
}

// NewDeleteContentmanagementWorkspaceTagvalueParamsWithHTTPClient creates a new DeleteContentmanagementWorkspaceTagvalueParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteContentmanagementWorkspaceTagvalueParamsWithHTTPClient(client *http.Client) *DeleteContentmanagementWorkspaceTagvalueParams {
	return &DeleteContentmanagementWorkspaceTagvalueParams{
		HTTPClient: client,
	}
}

/*
DeleteContentmanagementWorkspaceTagvalueParams contains all the parameters to send to the API endpoint

	for the delete contentmanagement workspace tagvalue operation.

	Typically these are written to a http.Request.
*/
type DeleteContentmanagementWorkspaceTagvalueParams struct {

	/* TagID.

	   Tag ID
	*/
	TagID string

	/* WorkspaceID.

	   Workspace ID
	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete contentmanagement workspace tagvalue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContentmanagementWorkspaceTagvalueParams) WithDefaults() *DeleteContentmanagementWorkspaceTagvalueParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete contentmanagement workspace tagvalue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteContentmanagementWorkspaceTagvalueParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete contentmanagement workspace tagvalue params
func (o *DeleteContentmanagementWorkspaceTagvalueParams) WithTimeout(timeout time.Duration) *DeleteContentmanagementWorkspaceTagvalueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete contentmanagement workspace tagvalue params
func (o *DeleteContentmanagementWorkspaceTagvalueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete contentmanagement workspace tagvalue params
func (o *DeleteContentmanagementWorkspaceTagvalueParams) WithContext(ctx context.Context) *DeleteContentmanagementWorkspaceTagvalueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete contentmanagement workspace tagvalue params
func (o *DeleteContentmanagementWorkspaceTagvalueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete contentmanagement workspace tagvalue params
func (o *DeleteContentmanagementWorkspaceTagvalueParams) WithHTTPClient(client *http.Client) *DeleteContentmanagementWorkspaceTagvalueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete contentmanagement workspace tagvalue params
func (o *DeleteContentmanagementWorkspaceTagvalueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagID adds the tagID to the delete contentmanagement workspace tagvalue params
func (o *DeleteContentmanagementWorkspaceTagvalueParams) WithTagID(tagID string) *DeleteContentmanagementWorkspaceTagvalueParams {
	o.SetTagID(tagID)
	return o
}

// SetTagID adds the tagId to the delete contentmanagement workspace tagvalue params
func (o *DeleteContentmanagementWorkspaceTagvalueParams) SetTagID(tagID string) {
	o.TagID = tagID
}

// WithWorkspaceID adds the workspaceID to the delete contentmanagement workspace tagvalue params
func (o *DeleteContentmanagementWorkspaceTagvalueParams) WithWorkspaceID(workspaceID string) *DeleteContentmanagementWorkspaceTagvalueParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the delete contentmanagement workspace tagvalue params
func (o *DeleteContentmanagementWorkspaceTagvalueParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContentmanagementWorkspaceTagvalueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
