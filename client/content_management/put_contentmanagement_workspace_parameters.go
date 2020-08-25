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

// NewPutContentmanagementWorkspaceParams creates a new PutContentmanagementWorkspaceParams object
// with the default values initialized.
func NewPutContentmanagementWorkspaceParams() *PutContentmanagementWorkspaceParams {
	var ()
	return &PutContentmanagementWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutContentmanagementWorkspaceParamsWithTimeout creates a new PutContentmanagementWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutContentmanagementWorkspaceParamsWithTimeout(timeout time.Duration) *PutContentmanagementWorkspaceParams {
	var ()
	return &PutContentmanagementWorkspaceParams{

		timeout: timeout,
	}
}

// NewPutContentmanagementWorkspaceParamsWithContext creates a new PutContentmanagementWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutContentmanagementWorkspaceParamsWithContext(ctx context.Context) *PutContentmanagementWorkspaceParams {
	var ()
	return &PutContentmanagementWorkspaceParams{

		Context: ctx,
	}
}

// NewPutContentmanagementWorkspaceParamsWithHTTPClient creates a new PutContentmanagementWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutContentmanagementWorkspaceParamsWithHTTPClient(client *http.Client) *PutContentmanagementWorkspaceParams {
	var ()
	return &PutContentmanagementWorkspaceParams{
		HTTPClient: client,
	}
}

/*PutContentmanagementWorkspaceParams contains all the parameters to send to the API endpoint
for the put contentmanagement workspace operation typically these are written to a http.Request
*/
type PutContentmanagementWorkspaceParams struct {

	/*Body
	  Workspace

	*/
	Body *models.Workspace
	/*WorkspaceID
	  Workspace ID

	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put contentmanagement workspace params
func (o *PutContentmanagementWorkspaceParams) WithTimeout(timeout time.Duration) *PutContentmanagementWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put contentmanagement workspace params
func (o *PutContentmanagementWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put contentmanagement workspace params
func (o *PutContentmanagementWorkspaceParams) WithContext(ctx context.Context) *PutContentmanagementWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put contentmanagement workspace params
func (o *PutContentmanagementWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put contentmanagement workspace params
func (o *PutContentmanagementWorkspaceParams) WithHTTPClient(client *http.Client) *PutContentmanagementWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put contentmanagement workspace params
func (o *PutContentmanagementWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put contentmanagement workspace params
func (o *PutContentmanagementWorkspaceParams) WithBody(body *models.Workspace) *PutContentmanagementWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put contentmanagement workspace params
func (o *PutContentmanagementWorkspaceParams) SetBody(body *models.Workspace) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the put contentmanagement workspace params
func (o *PutContentmanagementWorkspaceParams) WithWorkspaceID(workspaceID string) *PutContentmanagementWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the put contentmanagement workspace params
func (o *PutContentmanagementWorkspaceParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PutContentmanagementWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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